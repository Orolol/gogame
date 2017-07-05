# Installation

You will need the following :
- Go https://golang.org/doc/install
- ZMQ http://zeromq.org/intro:get-the-software
- libZMQ https://github.com/zeromq/libzmq
- libsodium https://github.com/jedisct1/libsodium
- czmq https://github.com/zeromq/czmq


There's 3 separate modules :  
- Utils : Contain models and function used by both API and engine
- gogame : the game engine. Read from ZMQ (to get game creation msg and game commands), push to ZMQ (to sedn new game state after each clock tick)
- API : Api RESTful, read and push with ZMQ, expose an API for the front (see https://github.com/orolol/gogame-front), and expose a websocket connection for FromChanToZMQ

install Utils firt :
- cd utils
- go get
- go build
- go install

install engine
- cd ../gogame
- go get
- go build
- ./gogame (launch the engine)

install Api
- cd ../Api
- go get
- go build
- ./Api

Port used whold be 31338, 31337 for ZMQ, 5001 for web socket
