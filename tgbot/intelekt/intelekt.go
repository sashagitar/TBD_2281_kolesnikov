package intelekt

// import (
// 	"reflect"
// 	"strconv"
// 	"strings"
// 	"time"

// 	"github.com/sashagitar/TBD_2281_kolesnikov/buylist"
// 	"github.com/sashagitar/TBD_2281_kolesnikov/produkt"
// )

// const dateStrFormat = "2006-01-02 15:04"

// type Intelekt struct {
// 	answer          string
// 	msg_stru        string
// 	help            string
// 	buylist_comands string
// 	holodos_comads  string
// 	pr              *produkt.Produkt
// 	date            *time.Time
// 	buylistV        *buylist.Buylist
// 	id_user         int
// 	funkt           int
// 	status          int
// 	param           int
// 	pole            int
// 	sort_list       bool
// 	sort_holodos    bool
// 	sort_stat       bool
// 	get             bool
// 	err             error
// }

// func getComand(c int) string {
// 	if c == 0 {
// 		return "Команды: \n" +
// 			"		список - открыть список продуктов, которые надо купить, и перейти к его редактированию\n" +
// 			"		холодильник - открыть список продуктов в холодильнике и перейти к его редактированию\n" +
// 			"		статистика - вывод продуктов, которые были использованны или выброшенны\n" +
// 			"		назад - вернутся в меню"
// 	}
// 	if c == 1 {
// 		return "Команды:\n" +
// 			"		добавить - добавить продукт\n" +
// 			"		купил [n] - перевести продукт из списка покупок в холодильник\n" +
// 			"		сортировать - изменения порядка сотрировки по алфавиту//по сроку годности\n" +
// 			"		назад - вернутся в меню"
// 	}
// 	if c == 2 {
// 		return "Команды:\n" +
// 			"		добавить - добавить продукт\n" +
// 			"		выбросить [n] - перевести продукт в статус 'выброшен'\n" +
// 			"		использовать [n] - перевести продукт в статус 'использован'\n" +
// 			"		открыть [n] - открыть продукт и назначить новый срок хранения\n" +
// 			"		сортировать - изменения порядка сотрировки по алфавиту//по сроку годности\n" +
// 			"		назад - вернутся в меню"
// 	}
// 	return "err"
// }

// func getFunc(s string) int {
// 	if s == "список" {
// 		return 1
// 	}
// 	if s == "холодильник" {
// 		return 2
// 	}
// 	if s == "статистика" {
// 		return 3
// 	}
// 	return -1
// }

// func getPole(p int) string {
// 	switch p {
// 	case -1:
// 		return "имя товара: "
// 	case 0:
// 		return "имя товара: "
// 	case 1:
// 		return "введите вес в кг: "
// 	case 2:
// 		return "Введите дату уведомления\nФормат гггг-мм-дд чч:мм"

// 	}
// 	return ""
// }

// func (i *Intelekt) Create(id_user int) {
// 	i.answer = ""
// 	i.msg_stru = ""
// 	i.get = false
// 	i.funkt = -1
// 	i.pr = nil
// 	i.date = nil
// 	i.buylistV = buylist.Get()
// 	i.id_user = id_user
// 	i.status = -1
// 	i.param = -1
// 	i.pole = -1
// 	i.sort_list = false
// 	i.sort_holodos = false
// 	i.sort_stat = false
// }

// func (i *Intelekt) Status() (int, int, int) {
// 	return i.status, i.param, i.pole
// }

// func (i *Intelekt) backe() {
// 	i.status = 0
// 	i.param = -1
// 	i.pole = -1
// 	i.pr = nil
// 	i.date = nil
// 	i.answer = getComand(0)
// }

// func (i *Intelekt) sortL(what bool) (string, error) {
// 	var l *bool
// 	if what {
// 		l = &i.sort_list
// 	} else {
// 		l = &i.sort_holodos
// 	}
// 	if *l {
// 		*l = false
// 	} else {
// 		*l = true
// 	}
// 	i.answer, i.err = i.buylistV.GetList(i.id_user, false, *l)
// 	return i.answer + "тип сортировки изменён", i.err
// }

// func (i *Intelekt) newDate(msg_user string) (string, bool) {
// 	msg_pole := getPole(i.pole)
// 	if i.pole == 0 {
// 		date, err := time.Parse("2006-01-02 15:04", msg_user)
// 		if err == nil {
// 			i.msg_stru += "Дата уведомления: " + date.String() + "\n"
// 			i.answer = i.msg_stru + msg_pole
// 			i.param++
// 		} else {
// 			msg_pole = getPole(i.param - 1)
// 			i.answer = "Введенны не верные данные\n\n" + i.msg_stru + msg_pole
// 		}
// 		return i.answer, false
// 	}

// 	if i.pole == 2 {
// 		i.answer = ""
// 		return i.answer, true
// 	}
// 	return "err", false
// }

// func (i *Intelekt) createProdukt(msg_user string, mode bool) (string, bool) {
// 	msg_pole := getPole(i.pole)
// 	if i.pole == 0 {
// 		i.msg_stru = ""
// 		i.pr = &produkt.Produkt{
// 			Bought:     mode,
// 			Used:       false,
// 			Thrown_out: false,
// 		}
// 		i.answer = "Для добавления продукта введите его параметры\n" + msg_pole
// 		i.pole++
// 		return i.answer, false
// 	}

// 	if i.pole == 1 {
// 		i.pr.Name = msg_user
// 		i.msg_stru = "Имя: " + msg_user + "\n"
// 		i.answer = i.msg_stru + msg_pole
// 		i.pole++
// 		return i.answer, false
// 	}

// 	if i.pole == 2 {
// 		weight, err := strconv.ParseFloat(msg_user, int(reflect.TypeOf(i.pr.Weight).Size()))
// 		if err != nil {
// 			msg_pole = getPole(i.pole - 1)
// 			i.answer = "введено не число\n\n" + i.msg_stru + msg_pole
// 		} else {
// 			i.pr.Weight = weight
// 			i.msg_stru += "Вес: " + msg_user + "\n"
// 			i.answer = i.msg_stru + msg_pole
// 			i.pole++
// 		}
// 		return i.answer, false
// 	}

// 	if i.pole == 3 {
// 		date, err := time.Parse("2006-01-02 15:04", msg_user)
// 		if err == nil {
// 			t := produkt.Timer{}
// 			if mode {
// 				t.SetNotification(&date)
// 			} else {
// 				t.SetTimer(&date)
// 			}
// 			i.pr.Timer = t
// 			i.msg_stru += "Дата уведомления: " + msg_user + "\n"
// 			i.answer = i.msg_stru + msg_pole
// 			i.pole++
// 		} else {
// 			msg_pole = getPole(i.pole - 1)
// 			i.answer = "Введенны не верные данные\n\n" + i.msg_stru + msg_pole
// 		}
// 		return i.answer, false
// 	}

// 	if i.pole == 4 {
// 		if msg_user == "д" || msg_user == "Д" {
// 			i.pole = -1
// 			return i.answer, true
// 		}
// 		if msg_user == "н" || msg_user == "Н" {
// 			i.pole = 0
// 			return i.answer, false
// 		}
// 	}
// 	return i.answer, false
// }

// func (i *Intelekt) list(msg_user string) (string, error) {
// 	if i.param == -1 {
// 		if msg_user == "добавить" {
// 			i.param = 0
// 			i.pole = 0
// 		}

// 		if msg_user == "сортировать" {
// 			i.answer, i.err = i.sortL(true)
// 			return i.answer, i.err
// 		}

// 		words := strings.Fields(msg_user)
// 		if words[0] == "купил" {
// 			if len(words) != 2 {
// 				return "Введите номер продукта в списке", nil
// 			}
// 			id, err := strconv.Atoi(words[1])
// 			if err != nil || id == 0 {
// 				return "Введено не число", err
// 			}
// 			f, err := i.buylistV.MoveInHolodos(id, i.id_user)
// 			if err != nil {
// 				return "err", err
// 			}
// 			if f {
// 				i.answer, i.err = i.buylistV.GetList(i.id_user, false, i.sort_list)
// 				i.answer += "Продукт перемещён в холодильник"
// 			} else {
// 				i.answer = "такого продукта нет"
// 			}
// 			return i.answer, i.err
// 		}
// 	}

// 	if i.param == 0 {
// 		i.answer, i.get = i.createProdukt(msg_user, false)
// 		if i.get {
// 			i.answer = "Произошла ошибка повторите попытку"
// 			f, err := i.buylistV.AddForList(i.pr, i.id_user)
// 			if err != nil {
// 				return "err", err
// 			}
// 			if f {
// 				i.answer = "Продукт добавлен"
// 				i.param = -1
// 				i.pr = nil
// 			} else {
// 				i.backe()
// 			}
// 			i.get = false
// 		}
// 		return i.answer, nil
// 	}
// 	return "err", nil
// }

// func (i *Intelekt) holodos(msg_user string) (string, error) {
// 	if i.param == -1 {
// 		if msg_user == "добавить" {
// 			i.param = 0
// 		}
// 		if msg_user == "сортировать" {
// 			i.answer, i.err = i.sortL(true)
// 			return i.answer, i.err
// 		}

// 		words := strings.Fields(msg_user)
// 		if len(words) != 2 {
// 			return "Введите номер продукта в списке", nil
// 		}
// 		id, err := strconv.Atoi(words[1])
// 		if err != nil || id == 0 {
// 			return "Введено не число", err
// 		}
// 		if words[0] == "выбросить" {
// 			f, err := i.buylistV.Trash(id, i.id_user)
// 			if err != nil {
// 				return "err", nil
// 			}
// 			if f {
// 				i.answer, i.err = i.buylistV.GetList(i.id_user, false, i.sort_list)
// 				i.answer += "Продукт перемещён выброшен"
// 			} else {
// 				i.answer = "такого продукта нет"
// 			}
// 			return i.answer, i.err
// 		}
// 		if words[0] == "использовать" {
// 			f, err := i.buylistV.Used(id, i.id_user)
// 			if err != nil {
// 				return "err", err
// 			}
// 			if f {
// 				i.answer, i.err = i.buylistV.GetList(i.id_user, false, i.sort_list)
// 				i.answer += "Продукт использован"
// 			} else {
// 				i.answer = "такого продукта нет"
// 			}
// 			return i.answer, i.err
// 		}
// 		if words[0] == "открыть" {
// 			i.answer, i.get = i.newDate(msg_user)
// 			if i.get {
// 				i.answer = "Произошла ошибка повторите попытку"
// 				i.pr = i.buylistV.Holodos.Get(id)
// 				if i.pr != nil {
// 					if i.pr.SetFinish(i.date, i.id_user) {
// 						i.answer = "Срок годности обновлён"
// 						i.param = -1
// 						i.date = nil
// 					}
// 					i.pr = nil
// 				}
// 			}
// 			return i.answer, nil
// 		}
// 	}

// 	if i.param == 0 {
// 		answer, get := i.createProdukt(msg_user, true)
// 		if get {
// 			i.answer = "Произошла ошибка повторите попытку"
// 			f, err := i.buylistV.AddForHolodos(i.pr, i.id_user)
// 			if err != nil {
// 				return "err", err
// 			}
// 			if f {
// 				i.answer = "Продукт добавлен"
// 				i.param = -1
// 				i.pr = nil
// 			} else {
// 				i.backe()
// 			}
// 		}
// 		return answer, nil
// 	}
// 	return "err", nil
// }

// func (i *Intelekt) statistiсs(id_user int) (string, error) {
// 	return i.buylistV.GetHistory(id_user)
// }

// func (i *Intelekt) GetAnsver(msg_user string, comand string) (string, int, error) {
// 	i.funkt = getFunc(msg_user)

// 	if i.status >= 0 {
// 		if msg_user == "назад" {
// 			i.backe()
// 			return i.answer, 0, nil
// 		}

// 		if i.status == 1 {
// 			i.answer, i.err = i.list(msg_user)
// 			return i.answer, 0, i.err
// 		}

// 		if i.status == 2 {
// 			i.answer, i.err = i.holodos(msg_user)
// 			return i.answer, 0, i.err
// 		}

// 		if i.funkt > 0 && i.status == 0 {
// 			if i.funkt == 1 {
// 				i.status = 1
// 				i.answer, i.err = i.buylistV.GetList(i.id_user, false, i.sort_list)
// 				i.answer += "\n" + getComand(i.funkt)
// 				return i.answer, 0, i.err
// 			}

// 			if i.funkt == 2 {
// 				i.status = 2
// 				i.answer, i.err = i.buylistV.GetList(i.id_user, true, i.sort_list)
// 				i.answer += "\n" + getComand(i.funkt)
// 				return i.answer, 0, i.err
// 			}

// 			if i.funkt == 3 {
// 				i.status = 3
// 				i.answer, i.err = i.statistiсs(i.id_user)
// 				return i.answer, 0, i.err
// 			}
// 		}

// 		if comand == "kill" {
// 			i.answer = "бот умер"
// 			return i.answer, -1, nil
// 		}

// 	}

// 	if comand == "start" {
// 		i.status = 0
// 		i.answer = getComand(0)
// 		return i.answer, 0, nil
// 	}

// 	return "err", 0, nil
// }
