package main

import (
	"github.com/muyisz/cutey-ani/data"
	"github.com/muyisz/cutey-ani/router"
)

func main() {
	db := data.CreatDB()
	db.InitDatabase()
	rou := router.InitRouter(db)
	rou.Run()
	db.CloseDatabase()
}
