package main

import (
	"log"
)

func sendTestExample() {
	log.Println(NewClient("<USERNAME>", "<SECRET>", false).SendMessage(&SMS{
		From:    "Github",
		To:      "+46…",
		Message: "Hej! 👋",
	}))
}
