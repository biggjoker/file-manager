package ws

import (
	"bytes"
	"fmt"
	"github.com/biggjoker/file-manager/zlog"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os/exec"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ExecCmd(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}

	defer ws.Close()

	var outInfo bytes.Buffer
	// todo timeout handle
	for {


		messageType, p, err := ws.ReadMessage()
		fmt.Println(p)
		if err != nil {
			zlog.Error("ws error", err)
			return
		}

		cmd := exec.Command("cmd.exe","echo","11111")
		cmd.Stdout = &outInfo
		err = cmd.Run()

		if err != nil {
			zlog.Error("ws error", err)
			return
		}
		if err == nil {
			if err := ws.WriteMessage(messageType, outInfo.Bytes()); err != nil {
				zlog.Error("ws error", err)
				return
			}
		}
	}

}

