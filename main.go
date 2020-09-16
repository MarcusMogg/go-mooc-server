package main

import (
	"fmt"
	"server/global"
	"server/initialize"
	"server/service"
)

func main() {
	initialize.Mysql()
	initialize.DBTables()
	go service.Transcoding()
	runServer()
}

func runServer() {
	Router := initialize.Router()
	Router.Static("source", "./source")

	address := fmt.Sprintf(":%d", global.GCONFIG.Addr)
	Router.Run(address)

}
