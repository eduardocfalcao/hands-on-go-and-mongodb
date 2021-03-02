package main

import (
	"fmt"
	"strconv"

	"github.com/eduardocfalcao/hands-on-go-and-mongodb/config"
	"github.com/eduardocfalcao/hands-on-go-and-mongodb/db"
	"github.com/eduardocfalcao/hands-on-go-and-mongodb/logger"
	"github.com/eduardocfalcao/hands-on-go-and-mongodb/service"
	"github.com/urfave/negroni"
)

func main() {
	logger.Init()
	config.Load()

	l := logger.Get()

	db, err := db.GetDB()
	if err != nil {
		panic(err)
	}

	l.Infof("Db Initialized: ", db)

	router := service.InitRouter()

	server := negroni.Classic()
	server.UseHandler(router)

	port := config.AppPort()
	addr := fmt.Sprintf(":%s", strconv.Itoa(port))

	server.Run(addr)
}
