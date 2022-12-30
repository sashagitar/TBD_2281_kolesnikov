package tgbot

import (
	"log"

	//tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/sashagitar/TBD_2281_kolesnikov/sqlmy"
	"github.com/sashagitar/TBD_2281_kolesnikov/tgbot/comands"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// var users_bots map[int]*intelekt.Intelekt
var users_bot_tupoi map[int]*comands.Params

func Create(api string) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(api)
	// users_bots = make(map[int]*intelekt.Intelekt, 0)
	users_bot_tupoi = make(map[int]*comands.Params, 0)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s\n", bot.Self.UserName)
	return bot
}

func sendAnswer(bot *tgbotapi.BotAPI, id_user int, answer *string) {
	msg := tgbotapi.NewMessage(int64(id_user), *answer)
	bot.Send(msg)
}

var connDB = "postgres://postgres:pass@localhost:5432/test"

func Run(bot *tgbotapi.BotAPI) {

	if err := sqlmy.Connect(connDB); err != nil {
		panic(err)
	}
	// make migration

	// Read migrations from /home/migrations and connect to a local postgres database.
	m, err := migrate.New("file://migrations", connDB)
	if err != nil {
		panic(err)
	}

	// Migrate all the way up ...
	if err := m.Up(); err != nil {
		panic(err)
	}
	// err = m.Force(1)
	// if err := m.Down(); err != nil {
	// 	log.Println(err)
	// }
	log.Println("migration is done")

	//11 is migrations version number, you may use your latest version
	if err != nil {
		log.Println(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {

		if update.Message == nil { // If we got a message
			continue
		}

		command := update.Message.Command()
		msg_user := update.Message.Text
		id_user := update.Message.From.ID

		if users_bot_tupoi[id_user] == nil {
			users_bot_tupoi[id_user] = &comands.Params{}
			users_bot_tupoi[id_user].Create(id_user)
		}
		answer, err := users_bot_tupoi[id_user].GetAnswer(command, msg_user)

		// if users_bots[id_user] == nil {
		// 	users_bots[id_user] = &intelekt.Intelekt{}
		// 	users_bots[id_user].Create(id_user)
		// }
		// answer, mode, err := users_bots[id_user].GetAnsver(msg_user, command)

		log.Printf("[%d] com %s, msg %s, %s", id_user, command, msg_user, err)
		// s, pa, po := users_bots[id_user].Status()
		// log.Printf("status %d, param %d, pole %d", s, pa, po)

		sendAnswer(bot, id_user, &answer)

		// if mode == -1 {
		// 	break
		// }
	}
}
