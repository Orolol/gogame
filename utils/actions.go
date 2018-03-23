package utils

var actions []PlayerActionOrder

func GetActions() []PlayerActionOrder {
	return actions
}

func GetAction(name string) PlayerActionOrder {
	var ret PlayerActionOrder
	for _, x := range actions {
		if x.ActionName == name {
			ret = x
		}
	}
	return ret
}

func SetBaseValueActions() {

	//ACTIONS
	actions = append(actions, PlayerActionOrder{
		Name:        "Civ Fact -> Lght Fact",
		ActionName:  "actionCivConvertFactoryToLightTankFact",
		Description: "Convert Civilian Factory to light Tank factory (Cost 1M) ",
		Cooldown:    10,
		Costs:       []Cost{Cost{Type: "money", Value: 1000000}},
	})

	actions = append(actions, PlayerActionOrder{
		Name:        "Civ Fact -> Hvy Fact",
		ActionName:  "actionCivConvertFactoryToHvyTankFact",
		Description: "Convert Civilian Factory to Heavy Tank factory (Cost 1M) ",
		Cooldown:    10,
		Costs:       []Cost{Cost{Type: "money", Value: 1000000}},
	})
	actions = append(actions, PlayerActionOrder{
		Name:        "War Propaganda",
		ActionName:  "actionWarPropaganda",
		Description: "Boost morale by 15% (cost 10M) ",
		Cooldown:    10,
		Costs:       []Cost{Cost{Type: "money", Value: 10000000}},
	})
	actions = append(actions, PlayerActionOrder{
		Name:        "Emergency recruitment",
		ActionName:  "emergencyRecruitment",
		Description: "Recruit immediatly 10% of your manpower. (cost 50M and 10% morale)",
		Cooldown:    30,
		Costs:       []Cost{Cost{Type: "money", Value: 50000000}, Cost{Type: "morale", Value: 10}},
	})
	actions = append(actions, PlayerActionOrder{
		Name:        "Purge of weak elements",
		ActionName:  "purgeSoldier",
		Description: "Executes possible traitors and weak soldier to improve morale and discipline by 15%. (cost 15% soldier and 10M)",
		Cooldown:    30,
		Costs:       []Cost{Cost{Type: "money", Value: 10000000}},
	})
	actions = append(actions, PlayerActionOrder{
		Name:        "Send spys to assassinate key scientits",
		ActionName:  "assassinateScientist",
		Description: "Send spys to assassinate key scientits (cost 45M)",
		Cooldown:    30,
		Effects:     []Effect{Effect{ModifierType: "Civilian", ModifierName: "NbScientist", Operator: "-", Value: 30, Target: "Opponent"}},
		Costs:       []Cost{Cost{Type: "money", Value: 45000000}},
	})
	actions = append(actions, PlayerActionOrder{
		Name:        "CBTEST",
		ActionName:  "CBTEST",
		Description: "CBTEST",
		Cooldown:    3,
		Costs:       []Cost{Cost{Type: "money", Value: 0}},
		Effects: []Effect{
			Effect{
				ModifierType: "Modifier",
				ModifierName: "TEST",
				Operator:     "turn+",
				Value:        3,
				Target:       "Player",
				Callbacks: []CallbackEffect{
					CallbackEffect{
						Constraints: []Constraint{Constraint{Type: "ModifierTurn", Operator: ">", Key: "TEST"}},
						Effects:     []Effect{Effect{ModifierType: "Army", ModifierName: "Morale", Operator: "+", Value: 30, Target: "Player"}},
					},
				},
			},
			Effect{ModifierType: "Army", ModifierName: "Morale", Operator: "-", Value: 30, Target: "Player"},
		},
	})
	actions = append(actions, PlayerActionOrder{
		Name:       "Buy foreign tanks",
		ActionName: "buyForeignTanks",
		Constraints: []Constraint{
			Constraint{Type: "tech", Value: "technoIndusT1N1"},
			Constraint{Type: "tech", Value: "technoIndusT1N2"},
		},
		Description: "Get 150 light tank and 50 heavy one (cost 100M) ",
		Cooldown:    60,
		Costs:       []Cost{Cost{Type: "money", Value: 10000000}},
	})

}
