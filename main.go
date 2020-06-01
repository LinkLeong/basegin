package main

import (
	"../basegin/pak/logging"
	"../basegin/pak/setting"
	"../basegin/router"
	"fmt"
	"net/http"
)

func main() {
	//model.Setup()
	setting.Setup()
	logging.Setup()
	//gredis.Setup()
	r := router.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServeSetting.HttpPort),
		Handler:        r,
		ReadTimeout:    setting.ServeSetting.ReadTimeout,
		WriteTimeout:   setting.ServeSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
