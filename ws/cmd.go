package ws

import (
	"bytes"
	"encoding/json"
	"github.com/biggjoker/file-manager/g"
	"github.com/biggjoker/file-manager/protocol"
	"github.com/biggjoker/file-manager/zlog"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os/exec"
	"runtime"
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

	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	for {


		messageType, p, err := ws.ReadMessage()
		if err != nil {
			zlog.Error("ws error", err)
			return
		}
		var msg protocol.CmdInfo
		if err :=json.Unmarshal(p,&msg);err != nil {
			if errWri := ws.WriteMessage(messageType,[]byte(err.Error())); errWri != nil {
				zlog.Error("ws error", errWri)
				return
			}
		}



		var cmd *exec.Cmd
		if runtime.GOOS == "windows" {
			cmd = exec.Command("powershell", msg.Cmd+"\n")
		}else {
			cmd = exec.Command("bash", msg.Cmd+"\n")
		}
		cmd.Dir = g.GetFullPath(msg.Path)
		cmdResult := bytes.NewBuffer(nil)
		cmd.Stdout = cmdResult

		if err := cmd.Run();err == nil{
			resp := protocol.CmdRespone{
				cmdResult.String(),
			}
			respWrap ,_ := json.Marshal(resp)
			if err := ws.WriteMessage(messageType,respWrap); err != nil {
				zlog.Error("ws error", err)
				return
			}
		} else {
			if err := ws.WriteMessage(messageType,[]byte(err.Error())); err != nil {
				zlog.Error("ws error", err)
				return
			}
		}

	}

}