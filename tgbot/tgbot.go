package tgbot

import (
	"log"

	//tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/sashagitar/TBD_2281_kolesnikov/buylist"
	"github.com/sashagitar/TBD_2281_kolesnikov/sqlmy"
	"github.com/sashagitar/TBD_2281_kolesnikov/tgbot/comands"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var connDB = "postgres://postgres:pass@localhost:5432/test"

type com interface {
	GetAnswer(command string, msg string) (string, error)
}

type Bot struct {
	users_bot_tupoi map[int]*comands.Params
	store           buylist.BD
	bot             *tgbotapi.BotAPI
	// users_bots map[int]*intelekt.Intelekt
}

// Создание бота
func Create(api string) (*Bot, error) {
	// Подключение бота
	bot, err := tgbotapi.NewBotAPI(api)
	if err != nil {
		log.Panic(err)
	}
	// Создание объекта бота
	b := Bot{}
	b.bot = bot

	// Выбор интелекта
	// b.users_bots = make(map[int]*intelekt.Intelekt, 0)
	b.users_bot_tupoi = make(map[int]*comands.Params, 0)

	// Побдключение к базе данных
	b.store, err = sqlmy.Connect(connDB)
	if err != nil {
		return nil, err
	}
	log.Printf("Authorized on account %s\n", bot.Self.UserName)

	return &b, nil
}

// Отправка сообщения пользователю
func (b *Bot) sendAnswer(id_user int, answer *string) {
	msg := tgbotapi.NewMessage(int64(id_user), *answer)
	b.bot.Send(msg)
}

// Запуск бота
func (b *Bot) Run() {
	// Запуск ожидания сообщений от пользователя
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {

		if update.Message == nil { // If we got a message
			continue
		}

		// Получение id  пользователя,команды и сообщения
		command := update.Message.Command()
		msg_user := update.Message.Text
		id_user := update.Message.From.ID

		// Если пользователь написал впервые - создаём для него отдеьный объект интелекта
		if b.users_bot_tupoi[id_user] == nil {
			b.users_bot_tupoi[id_user] = comands.Create(id_user, b.store)
		}

		// Отправляем сообщение интелекту
		answer, err := b.users_bot_tupoi[id_user].GetAnswer(command, msg_user)

		// if b.users_bots[id_user] == nil {
		// 	b.users_bots[id_user] = &intelekt.Intelekt{}
		// 	b.users_bots[id_user].Create(id_user)
		// }
		// answer, mode, err := b.users_bots[id_user].GetAnsver(msg_user, command)

		log.Printf("[%d] com %s, msg %s, %s", id_user, command, msg_user, err)
		// s, pa, po := users_bots[id_user].Status()
		// log.Printf("status %d, param %d, pole %d", s, pa, po)

		b.sendAnswer(id_user, &answer)

		// if mode == -1 {
		// 	break
		// }
	}
}
