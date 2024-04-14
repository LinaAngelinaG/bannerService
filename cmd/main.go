package main

import (
	"bannerService/config"
	"bannerService/config/inj"
	httpChi "bannerService/internal/banner/httpChi"
	"bannerService/internal/server"
)

func main() {
	defer httpChi.HandlePanic()
	cfg := config.Load()
	inj.Init(cfg)
	defer server.StopContext()
	server.New(cfg).Run()
}
