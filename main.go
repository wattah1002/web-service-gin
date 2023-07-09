package main

import (
	"web-service-gin/db"
	"web-service-gin/router"
)

func main() {
	dbConn := db.Init()
	router.Router(dbConn)
}
