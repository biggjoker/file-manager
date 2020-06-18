package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/biggjoker/file-manager/app"
	"github.com/biggjoker/file-manager/g"
	"github.com/biggjoker/file-manager/zlog"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func startSignal(pid int) {
	sigs := make(chan os.Signal, 1)
	zlog.Info(pid, "register signal notify")
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	for {
		s := <-sigs
		zlog.Info("recv ", s)

		switch s {
		case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			zlog.Info(pid, "exit")
			os.Exit(0)
		}
	}
}

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	g.ParseConfig(*cfg)
	out := zlog.InitLog(g.Config().Debug)

	cmd := exec.Command("cmd.exe","echo","11111")
	stdout, err := cmd.StdoutPipe()
	if  err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(stdout)

	line,err :=reader.ReadString('\n')
	for err == nil {
		fmt.Println(line)
		line,err =reader.ReadString('\n')
	}


	//ws.CreateWs()

	gin.DefaultWriter = out
	gin.DefaultErrorWriter = out
	routes := gin.Default()
	routes.Static("/static", "./assets/dist/static")
	routes.StaticFile("/", "./assets/dist/index.html")

	app.Routes(routes)
	go routes.Run(g.Config().HttpAddr)
	zlog.Info("process start")

	startSignal(os.Getpid())
}