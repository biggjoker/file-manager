package ws

import (
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
	cmd := exec.Command("cmd.exe")

	// todo timeout handle
	for {


		messageType, p, err := ws.ReadMessage()
		if err != nil {
			zlog.Error("ws error", err)
			return
		}
		cmd.Stdin.Read(p)
		err = cmd.Run()
		out, err := cmd.Output()

		if err != nil {
			zlog.Error("ws error", err)
			return
		}
		if err != nil {
			if err := ws.WriteMessage(messageType, out); err != nil {
				zlog.Error("ws error", err)
				return
			}
		}
	}

}

