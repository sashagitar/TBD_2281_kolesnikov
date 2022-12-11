package comands

import (
	"strconv"
	"strings"
	"time"

	"github.com/sashagitar/TBD_2281_kolesnikov/buylist"
	"github.com/sashagitar/TBD_2281_kolesnikov/produkt"
)

type Params struct {
	id_user  int
	buylistV *buylist.Buylist
	sort     bool
}

const help = `/addForList имя вес гггг-мм-дд чч:мм(уведомление если требуется) - добавление продукта в список покупок
/addForHolodos имя вес гггг-мм-дд чч:мм(дата истечения срока годности) - добавление продукта в холодильник
/moveProdToHolod [id] - перемещение продукта из списка в холодильник
/open [id] гггг-мм-дд чч:мм(новая дата истечения срока годности) - открыть продукт
/use [id] - использование продукта
/drop [id] - выбросить продукт
/statsuse - статистика использованных ранее продуктов
/stats гггг-мм-дд чч:мм(дата начала) гггг-мм-дд чч:мм(дата окончания) - количество использованных и выброшенных продуктов за указанный период`

const dateStrFormat = "2006-01-02 15:04"

func (p *Params) Create(id_user int) {
	p.id_user = id_user
	p.buylistV = buylist.Get()

}

func (p Params) addProdukt(msg string, mode bool) (string, error) {
	words := strings.Fields(msg)
	if mode && len(words) != 5 {
		return "слишком мало аргументов", nil
	}
	if len(words) > 3 {
		weight, err := strconv.ParseFloat(words[2], 64)
		if err != nil {
			return "", err
		}
		pr := produkt.Produkt{
			Name:       words[1],
			Weight:     weight,
			Bought:     mode,
			Used:       false,
			Thrown_out: false,
		}
		if len(words) == 5 {
			s := time.Now()
			f, err := time.Parse(dateStrFormat, words[3]+" "+words[4])
			if err != nil {
				return "Дата введена не в правильном формате", nil
			}
			t := produkt.Timer{
				Notif:  true,
				Start:  &s,
				Finish: &f,
			}
			pr.Timer = t
		}
		var f bool
		if mode {
			f, err = p.buylistV.AddForList(&pr, p.id_user)
		} else {
			f, err = p.buylistV.AddForHolodos(&pr, p.id_user)
		}
		if err != nil {
			return "ошибка бд", err
		}
		if f {
			list, err := p.buylistV.GetList(p.id_user, false, p.sort)
			if err == nil {
				return "Произошла ошибка", nil
			}
			return "Продукт добавлен\n\n" + list, nil
		}
		return "Я не знаю, как ты так умудлился", nil
	}
	return "err", nil
}

func (p Params) moveProdToHolod(msg string) (string, error) {
	words := strings.Fields(msg)
	if len(words) != 2 {
		return "Неверное количество аргументов", nil
	}
	id, err := strconv.Atoi(words[1])
	if err != nil {
		return "Введено не число", err
	}
	f, err := p.buylistV.MoveInHolodos(id, p.id_user)
	if err != nil {
		return "Произошла ошибка", err
	}
	if f {
		list, err := p.buylistV.GetList(p.id_user, true, p.sort)
		if err != nil {
			return "Произошла ошибка", err
		}
		return "продукт перемещён\n" + list, nil
	}
	return "такого продукта нет", nil
}

func (p Params) open(msg string) (string, error) {
	words := strings.Fields(msg)
	if len(words) != 4 {
		return "неверное количество аргументов", nil
	}
	id, err := strconv.Atoi(words[1])
	if err != nil {
		return "Введено не число", err
	}
	t, err := time.Parse(dateStrFormat, words[2]+" "+words[3])
	if err != nil {
		return "Дата введена не в правильном формате", nil
	}
	pr := p.buylistV.Holodos.Get(id)
	if pr == nil {
		return "Нет такого продукта", nil
	}
	pr.SetTimer(&t)
	return "данные обновлены", nil
}

func (p Params) useOrDrop(msg string, mode bool) (string, error) {
	words := strings.Fields(msg)
	if len(words) != 2 {
		return "неверное количество аргументов", nil
	}
	id, err := strconv.Atoi(words[1])
	if err != nil {
		return "Введено не число", err
	}
	var f bool
	if mode {
		f, err = p.buylistV.Used(id, p.id_user)
	} else {
		f, err = p.buylistV.Trash(id, p.id_user)
	}
	if err != nil {
		return "такого продукта нет", nil
	}
	if f {
		return "Статус продукта обновлён", nil
	}
	return "Не, ну ты молодец, я не знаю как это у тебя получилось", nil

}

func (p Params) statsuse(msg string) (string, error) {
	return p.buylistV.GetHistory(p.id_user)
}

func (p Params) stats(msg string) (string, error) {
	words := strings.Fields(msg)
	if len(words) != 5 {
		return "неверное количество аргументов", nil
	}

	ts, err := time.Parse(dateStrFormat, words[1]+" "+words[2])
	if err != nil {
		return "Дата введена не в правильном формате", nil
	}
	tf, err := time.Parse(dateStrFormat, words[3]+" "+words[4])
	if err != nil {
		return "Дата введена не в правильном формате", nil
	}
	use, drop, err := p.buylistV.GetStats(p.id_user, ts, tf)
	if err != nil {
		return "ошибка бд", err
	}
	return " Выброшено " + strconv.Itoa(drop) + " использованно " + strconv.Itoa(use), nil
}

func (p Params) GetAnswer(command string, msg string) (string, error) {
	if command == "start" {
		return help, nil
	}
	if command == "addForList" {
		return p.addProdukt(msg, false)
	}

	if command == "addForHolodos" {
		return p.addProdukt(msg, true)
	}

	if command == "moveProdToHolod" {
		return p.moveProdToHolod(msg)
	}

	if command == "open" {
		return p.open(msg)
	}

	if command == "use" {
		return p.useOrDrop(msg, true)
	}

	if command == "drop" {
		return p.useOrDrop(msg, true)
	}

	if command == "statsuse" {
		return p.statsuse(msg)
	}

	if command == "stats" {
		return p.stats(msg)
	}
	return help, nil
}