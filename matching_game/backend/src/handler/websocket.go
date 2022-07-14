package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/SuperTikuwa/matching_game/models"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

const (
	WriteBufferSize = 1024
	ReadBufferSize  = 1024
)

var connections = make(map[string][]*websocket.Conn)

func GameSocket(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  ReadBufferSize,
		WriteBufferSize: WriteBufferSize,
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	gameID := mux.Vars(r)["gameID"]
	index := len(connections[gameID])
	connections[gameID] = append(connections[gameID], conn)

	for {
		gameStatus := models.GameStatus{}
		_, msg, err := conn.ReadMessage()
		if err != nil {
			connections[gameID] = append(connections[gameID][:index], connections[gameID][index+1:]...)
			return
		}

		err = json.Unmarshal(msg, &gameStatus)
		if err != nil {
			log.Println("unmarshal error:", err)
			continue
		}

		broadCast(connections[gameID], string(msg))
	}
}

func broadCast(connections []*websocket.Conn, msg string) {
	for _, conn := range connections {
		err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			return
		}
	}
}
