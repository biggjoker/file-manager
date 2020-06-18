package ws

import (
	"flag"
	"github.com/biggjoker/file-manager/g"
	"net/http"
)

var (
	addr *string
)

func CreateWs() {
	go func() {
		addr = flag.String("addr", "127.0.0.1"+g.Config().WsAddr, "http service address")
		http.HandleFunc("/file-ws/cmd", ExecCmd)
		http.ListenAndServe(*addr, nil)
	}()
}
