package main

import (
	"log"
	"os"
)

var localClient *Elks46

func init() {
	localClient = NewClient(os.Getenv("46_USERNAME"), os.Getenv("46_SECRET"), true)
}

func sendTestExample() {
	log.Println(localClient.SendMessage(&SMS{
		From:    "Github",
		To:      "+46…",
		Message: "Hej! 👋",
	}))
}
