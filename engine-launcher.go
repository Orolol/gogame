package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"github.com/pebbe/zmq4"
	"github.com/google/uuid"
)

//InitializePlayerDefaultValue init player
func InitializePlayerDefaultValue(idPlayer int) PlayerInGame {
	army := PlayerArmy{
		NbSoldier:  1000,
		NbLigtTank: 100,
		NbHvyTank:  50,
		NbArt:      50,
		NbAirSup:   0,
		NbAirBomb:  0}

	policy := PlayerModifierPolicy{
		RecruitmentPolicy: 1}

	var player = PlayerInGame{
		PlayerID:       idPlayer,
		ModifierPolicy: policy,
		Army:           army,
		nick:           "Player " + strconv.Itoa(idPlayer),
		NbPop:          10000}

	return player
}

//PlayerAction player action
type PlayerAction func(player *PlayerInGame, value float32) *PlayerInGame

//PASetRecruitementPolicy change recruitement policy to the value
func PASetRecruitementPolicy(player *PlayerInGame, value float32) *PlayerInGame {
	player.ModifierPolicy.RecruitmentPolicy = value
	return player
}

func createGame(idp1 int, idp2 int, conf GameConf, queue chan GameMsg) Game {
	var gameID = uuid.New()
	fmt.Println("Creating game")
	// Go grab player profile in base
	var mockP1 = InitializePlayerDefaultValue(1)
	var mockP2 = InitializePlayerDefaultValue(2)
	var listPlayer = map[int]PlayerInGame{mockP1.PlayerID: mockP1, mockP2.PlayerID: mockP2}
	var game = Game{
		GameID:      gameID,
		CurrentTurn: 0,
		ListPlayers: listPlayer,
		Conf:        conf,
		Queue:       queue}

	return game
}

func GameEvent (queue chan GameMsg, game Game, player1, player2 *PlayerInGame) {
	fmt.Println("Running game event")
	for msg := range queue {
		fmt.Println(msg.Text)
		player, ok := game.ListPlayers[msg.PlayerID]
		if ok {
			fmt.Println("FOUND")
			if player1.PlayerID == player.PlayerID {
				player1 = msg.Action(player1, msg.value)
			} else {
				player2 = msg.Action(player2, msg.value)
			}
			fmt.Println(player)
			fmt.Println(game.ListPlayers)
		} else {
			fmt.Println("Erreur de destinataire", msg)
		}
	}
}

func runGame(game Game, queue chan GameMsg, queueGameOut chan Game) {
	var player1, player2 PlayerInGame
	for key, value := range game.ListPlayers {
		if key == 1 {
			player1 = value
		} else if key == 2 {
			player2 = value
		}
	}

	go GameEvent (queue, game, &player1, &player2)

	fmt.Println("Start game ", player1.nick, " vs ", player2.nick)




	for game.CurrentTurn < 5 {
		timer1 := time.NewTimer(time.Second)
		//Read the stack
		//Resolve combat
		var preFightP1 = player1
		var preFightP2 = player2

		player2 = AlgoDamageRepartition(player2, AlgoDamageDealt(preFightP1))
		player1 = AlgoDamageRepartition(player1, AlgoDamageDealt(preFightP2))

		player2.Army.NbSoldier += AlgoReinforcement(player2)
		player1.Army.NbSoldier += AlgoReinforcement(player1)

		player2.NbPop -= AlgoReinforcement(player2)
		player1.NbPop -= AlgoReinforcement(player1)

		<-timer1.C
		game.CurrentTurn++
		queueGameOut <- game

		fmt.Println("Next turn ", game.CurrentTurn)
	}
	fmt.Println("End game")
}

//AlgoDamageDealt Calculate dmg dealt
func AlgoDamageDealt(player PlayerInGame) float32 {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var rollp1 = r1.Float32()
	var dmg = player.Army.NbSoldier * 0.05 * rollp1
	fmt.Println("Calculating dmg for ", player, " :: ", dmg)
	return dmg
}

//AlgoReinforcement calc reinforcement
func AlgoReinforcement(player PlayerInGame) float32 {
	reinforcement := player.NbPop * 0.001 * player.ModifierPolicy.RecruitmentPolicy
	fmt.Println("Calculating renf for ", player, " :: ", reinforcement)
	return reinforcement
}

//AlgoDamageRepartition Calculate loses
func AlgoDamageRepartition(player PlayerInGame, dmgIncoming float32) PlayerInGame {
	player.Army.NbSoldier -= dmgIncoming
	return player
}

func GameManagerF(queueGameOut chan Game) {
	var conf = GameConf{
		GameType:  "test",
		NbPlayers: 2}
	var GameList = make(map[uuid.UUID]Game)
	GameManager := GameManager{GameList: GameList}
	queue := make(chan GameMsg)
	game := createGame(11, 22, conf, queue)
	fmt.Println(game.GameID)
	go runGame(game, queue, queueGameOut)

	GameManager.GameList[game.GameID] = game

}

func ZMQReader() {

}
func ZMQPusher() {

}

func main() {

	push, err := zmq4.NewSocket(zmq4.REQ)
	if err != nil {
		panic(err)
	}
	push.Connect("tcp://127.0.0.1:31337")
	fmt.Printf("PUSH Queue connected\n")
	pull, err := zmq4.NewSocket(zmq4.REP)
	if err != nil {
		panic(err)
	}
	time.Sleep(1 * time.Second)
	pull.Connect("tcp://127.0.0.1:31338")
	fmt.Printf("PULL Queue connected\n")
	res, err := push.SendMessage("Hello World", zmq4.DONTWAIT)
	if err != nil {
		panic(err)
	}
	fmt.Printf("HW sent\n")
	time.Sleep(1* time.Second)
	sz, err := pull.RecvMessage(zmq4.DONTWAIT)
	if err != nil {
		panic(err)
	}
	fmt.Printf("HW rcv\n")
	for _,m := range sz{
		fmt.Printf(m)
	}

	fmt.Printf("We received a message of size ", res)


	var queueGameOut chan Game

	go GameManagerF(queueGameOut)

	var lGmsg []GameMsg
	lGmsg = append(lGmsg, GameMsg{Action: PASetRecruitementPolicy, PlayerID: 1, Text: "Change rec value to 5", value: 5.0})

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
