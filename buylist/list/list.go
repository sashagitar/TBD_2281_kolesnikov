package list

import (
	"fmt"

	"github.com/sashagitar/TBD_2281_kolesnikov/buylist/produkt"
)

type List struct {
	list map[int]*produkt.Produkt
}

func (l *List) Create() {
	l.list = make(map[int]*produkt.Produkt)
}

func (l *List) Add(p *produkt.Produkt) {
	p.Id = len(l.list) + 1
	l.list[len(l.list)+1] = p
}

func (l *List) Delete(id int) {
	delete(l.list, id)
}

func (l *List) Clear() {
	l.Clear()
}

func (l *List) Get(id int) *produkt.Produkt {
	p := l.list[id]
	if p != nil {
		l.Delete(id)
		return p
	}
	return nil
}

func (l *List) Print() {
	var str string
	for _, v := range l.list {
		if !v.Thrown_out && !v.Used {
			str += v.Name + "\n"
		}
	}
	fmt.Print(str)
}
