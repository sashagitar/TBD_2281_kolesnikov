package buylist

import (
	"time"

	"github.com/sashagitar/TBD_2281_kolesnikov/list"
	"github.com/sashagitar/TBD_2281_kolesnikov/produkt"
	"github.com/sashagitar/TBD_2281_kolesnikov/sqlmy"
)

type store interface {
	GetList(id_user int, bought bool, sort bool) ([]sqlmy.ProduktDB, error)
}

type Buylist struct {
	Holodos list.List
	List    list.List
	Store   store
}

func (b *Buylist) Clear(m bool) {
	if m {
		b.Holodos.Clear()
	} else {
		b.List.Clear()
	}
}

func (b *Buylist) GetList(id_user int, bought bool, sort bool) (string, error) {
	s := "Список продуктов:\n"
	b.Clear(bought)
	prdb, err := b.Store.GetList(id_user, bought, sort)
	if err != nil {
		return "", err
	}
	for _, v := range prdb {
		p := produkt.ParseProduktDB(v)
		if bought {
			b.AddForHolodos(p, -1)
		} else {
			b.AddForList(p, -1)
		}
		s += p.String() + "\n"
	}

	return s + "\n", err
}

func (b *Buylist) GetHistory(id_user int) (string, error) {
	s := "Список продуктов:\n"
	prdb, err := sqlmy.GetUseList(id_user)
	if err != nil {
		return "", err
	}
	for _, v := range *prdb {
		p := produkt.ParseProduktDB(v)
		s += p.String() + "\n"
	}

	return s + "\n", err
}

func (b *Buylist) GetStats(id_user int, ts time.Time, tf time.Time) (int, int, error) {
	use, drop, err := sqlmy.GetStats(id_user, ts, tf)
	if err != nil {
		return -1, -1, err
	}
	return use, drop, err
}

func (b *Buylist) Create() {
	b.Holodos.Create()
	b.List.Create()
}

func (b *Buylist) AddForList(p *produkt.Produkt, id_user int) (bool, error) {
	if id_user > 0 {
		pdb := produkt.ParseProdukt(*p, id_user)
		f, err := sqlmy.AddProdukt(id_user, &pdb)
		if err != nil || !f {
			return false, err
		}
	}
	b.List.Add(p)
	return true, nil
}

func (b *Buylist) AddForHolodos(p *produkt.Produkt, id_user int) (bool, error) {
	if id_user > 0 {
		pdb := produkt.ParseProdukt(*p, id_user)
		f, err := sqlmy.AddProdukt(id_user, &pdb)
		if err != nil || !f {
			return false, err
		}
	}
	b.Holodos.Add(p)
	return true, nil
}

func (b *Buylist) MoveInHolodos(id int, id_user int) (bool, error) {
	p := b.List.Get(id)
	if p != nil {
		f, err := sqlmy.MoveInHolodos(id_user, p.Id_bd)
		if err != nil || !f {
			return false, err
		}
		p.Bought = true
		b.Holodos.Add(p)
		return true, nil
	}
	return false, nil
}

func (b *Buylist) Trash(id int, id_user int) (bool, error) {
	p := b.Holodos.Get(id)
	if p != nil {
		f, err := sqlmy.Trash(p.Id_bd, id_user)
		if err != nil || !f {
			return false, err
		}
		p.Thrown_out = true
		return true, nil
	}
	return false, nil
}

func (b *Buylist) Used(id int, id_user int) (bool, error) {
	p := b.Holodos.Get(id)
	if p != nil {
		f, err := sqlmy.Used(p.Id_bd, id_user)
		if err != nil || !f {
			return false, err
		}
		p.Used = true
		return true, nil
	}
	return false, nil
}

func Get() *Buylist {
	v := Buylist{}
	v.Create()
	return &v
}
