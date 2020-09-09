package main

import (
	"fmt"
	"server/global"
	"server/initialize"
)

func main() {
	initialize.Mysql()
	initialize.DBTables()
	runServer()
}

func runServer() {
	Router := initialize.Router()
	Router.Static("source", "./resourse")

	address := fmt.Sprintf(":%d", global.GCONFIG.Addr)
	Router.Run(address)
}
