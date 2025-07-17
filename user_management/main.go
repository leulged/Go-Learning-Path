package main

import (
	"user_management/router"
	"user_management/config"
)

func main(){
	config.ConnectToMongo()
	r := router.SetupRouter()
	r.Run(":8080")
}