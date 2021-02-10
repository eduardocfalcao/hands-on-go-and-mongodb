package main

import (
	"fmt"
	"strconv"

	"github.com/eduardocfalcao/hands-on-go-and-mongodb/db"
	"github.com/eduardocfalcao/hands-on-go-and-mongodb/service"
	"github.com/urfave/negroni"
)

func main() {

	db, err := db.GetDB()
	if err != nil {
		panic(err)
	}

	fmt.Println("Db Initialized: ", db)

	router := service.InitRouter()

	server := negroni.Classic()
	server.UseHandler(router)

	port := 33001
	addr := fmt.Sprintf(":%s", strconv.Itoa(port))

	server.Run(addr)
}
