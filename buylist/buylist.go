package buylist

import (
	"time"

	"github.com/sashagitar/TBD_2281_kolesnikov/list"
	"github.com/sashagitar/TBD_2281_kolesnikov/produkt"
	"github.com/sashagitar/TBD_2281_kolesnikov/sqlmy"
)

type BD interface {
	AddProdukt(id_user int, p *sqlmy.ProduktDB) (bool, error)
	Trash(id int, id_user int) (bool, error)
	Used(id int, id_user int) (bool, error)
	SetFinish(id_user int, id int, date time.Time) (bool, error)
	MoveInHolodos(id_user int, id int) (bool, error)
	GetList(id_user int, bought bool, sort bool) (*[]sqlmy.ProduktDB, error)
	GetUseList(id_user int) (*[]sqlmy.ProduktDB, error)
	GetStats(id_user int, ts time.Time, tf time.Time) (int, int, error)
}

type Buylist struct {
	holodos list.List
	list    list.List
	store   BD
}

func Get(db BD) *Buylist {
	v := Buylist{}
	v.create(db)
	return &v
}

func (b *Buylist) Clear(m bool) {
	if m {
		b.holodos.Clear()
	} else {
		b.list.Clear()
	}
}

func (b *Buylist) GetList(id_user int, bought bool, sort bool) (string, error) {
	s := "Список продуктов:\n"
	b.Clear(bought)
	prdb, err := b.store.GetList(id_user, bought, sort)
	if err != nil {
		return "", err
	}
	for _, v := range *prdb {
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
	prdb, err := b.store.GetUseList(id_user)
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
	use, drop, err := b.store.GetStats(id_user, ts, tf)
	if err != nil {
		return -1, -1, err
	}
	return use, drop, err
}

func (b *Buylist) create(db BD) {
	b.holodos.Create()
	b.list.Create()
	b.store = db
}

func (b *Buylist) AddForList(p *produkt.Produkt, id_user int) (bool, error) {
	if id_user > 0 {
		pdb := produkt.ParseProdukt(*p, id_user)
		f, err := b.store.AddProdukt(id_user, &pdb)
		if err != nil || !f {
			return false, err
		}
	}
	b.list.Add(p)
	return true, nil
}

func (b *Buylist) AddForHolodos(p *produkt.Produkt, id_user int) (bool, error) {
	if id_user > 0 {
		pdb := produkt.ParseProdukt(*p, id_user)
		f, err := b.store.AddProdukt(id_user, &pdb)
		if err != nil || !f {
			return false, err
		}
	}
	b.holodos.Add(p)
	return true, nil
}

func (b *Buylist) MoveInHolodos(id int, id_user int) (bool, error) {
	p := b.list.Get(id)
	if p != nil {
		f, err := b.store.MoveInHolodos(id_user, p.Id_bd)
		if err != nil || !f {
			return false, err
		}
		p.Bought = true
		b.holodos.Add(p)
		return true, nil
	}
	return false, nil
}

func (b *Buylist) Trash(id int, id_user int) (bool, error) {
	p := b.holodos.Get(id)
	if p != nil {
		f, err := b.store.Trash(p.Id_bd, id_user)
		if err != nil || !f {
			return false, err
		}
		p.Thrown_out = true
		return true, nil
	}
	return false, nil
}

func (b *Buylist) Used(id int, id_user int) (bool, error) {
	p := b.holodos.Get(id)
	if p != nil {
		f, err := b.store.Used(p.Id_bd, id_user)
		if err != nil || !f {
			return false, err
		}
		p.Used = true
		return true, nil
	}
	return false, nil
}

func (b *Buylist) OpenProdukt(id int, id_user int, date *time.Time) (bool, error) {
	p := b.holodos.Get(id)
	if p != nil {
		f, err := b.store.SetFinish(id_user, p.Id_bd, *date)
		if err != nil || !f {
			return false, err
		}
		p.SetTimer(&produkt.Timer{Finish: date})
		return true, nil
	}
	return false, nil
}
