package main

import (
	"task_manager/data"
	"task_manager/router"
)

func main() {
	data.ConnectToMongo()
	r := router.SetupRouter()
	r.Run(":8080")
}
