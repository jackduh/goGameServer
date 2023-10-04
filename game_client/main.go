package main

import (
	"encoding/json"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
	"github.com/jackduh/goGameServer/types"
)

const wsServerEndPoint = "ws://localhost:4000/ws"

type GameClient struct {
	conn     *websocket.Conn
	clientID int
	username string
}

func newGameClient(conn *websocket.Conn, username string) *GameClient {
	return &GameClient{
		clientID: rand.Intn(math.MaxInt),
		username: username,
		conn:     conn,
	}
}

func (c *GameClient) login() error {
	b, err := json.Marshal(types.Login{
		ClientID: c.clientID,
		Username: c.username,
	})
	if err != nil {
		return err
	}
	msg := types.WSMessage{
		Type: "Login",
		Data: b,
	}
	return c.conn.WriteJSON(msg)
}
func main() {
	dialer := websocket.Dialer{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, _, err := dialer.Dial(wsServerEndPoint, nil)

	if err != nil {
		log.Fatal(err)
	}

	c := newGameClient(conn, "Jack")

	if err := c.login(); err != nil {
		log.Fatal(err)
	}

	for {
		x := rand.Intn(1000)
		y := rand.Intn(1000)

		state := types.PlayerState{
			Health:   100,
			Position: types.Position{X: x, Y: y},
		}

		b, err := json.Marshal(state)
		if err != nil {
			log.Fatal(err)
		}

		msg := types.WSMessage{
			Type: "playerState",
			Data: b,
		}
		if err := conn.WriteJSON(msg); err != nil {
			log.Fatal(err)
		}

		time.Sleep(time.Millisecond * 100)
	}
}
