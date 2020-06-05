package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/biggjoker/file-manager/app"
	"github.com/biggjoker/file-manager/g"
	"github.com/biggjoker/file-manager/zlog"
	"github.com/gin-gonic/gin"
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
