package main

import (
	"fmt"

	"github.com/sashagitar/TBD_2281_kolesnikov/tgbot"
)

const api_tg = "5701189684:AAFFPRctBBKqjl-yRzo7sIy-hsg8cGApz_4"

func main() {
	fmt.Printf("%s\n", api_tg)

	// Создание бота
	bot, err := tgbot.Create(api_tg)
	if err != nil {
		panic(err)
	}
	// Запуск бота
	bot.Run()
}
