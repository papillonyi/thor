package main

import (
	"fmt"
	"github.com/papillonyi/thor/api"
	"github.com/papillonyi/thor/model"
	"github.com/papillonyi/thor/mongo"
	"github.com/papillonyi/thor/pkg/setting"
	"log"
	"net/http"
)

func init() {
	setting.Setup()
	model.Setup()
	mongo.Setup()

}

func main() {
	//model.Migrate()
	routersInit := api.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout

	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
