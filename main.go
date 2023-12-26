package main

import (
	"fmt"

	"github.com/tushargarg0987/greetbot-discord/bot"
	"github.com/tushargarg0987/greetbot-discord/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
