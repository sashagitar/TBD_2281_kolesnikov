package produkt

import (
	"strconv"
	"time"

	"github.com/sashagitar/TBD_2281_kolesnikov/sqlmy"
)

type Timer struct {
	Notif    bool
	Start    *time.Time
	Finish   *time.Time
	Interval *time.Time
}

type Produkt struct {
	Id         int
	Id_bd      int
	Name       string
	Weight     float64
	Bought     bool
	Used       bool
	Thrown_out bool
	Timer      *Timer
}

func ParseProdukt(p Produkt, id_user int) sqlmy.ProduktDB {
	return sqlmy.ProduktDB{
		Id:          p.Id_bd,
		Id_user:     id_user,
		Name:        p.Name,
		Weight:      p.Weight,
		Bought:      p.Bought,
		Used:        p.Used,
		Thrown_out:  p.Thrown_out,
		Date_start:  *p.Timer.Start,
		Date_finish: *p.Timer.Finish,
	}

}

func ParseProduktDB(pdb sqlmy.ProduktDB) *Produkt {
	ti := Timer{
		Notif:  pdb.Bought,
		Start:  &pdb.Date_start,
		Finish: &pdb.Date_finish,
	}
	p := Produkt{
		Id_bd:      pdb.Id,
		Name:       pdb.Name,
		Weight:     pdb.Weight,
		Bought:     pdb.Bought,
		Used:       pdb.Used,
		Thrown_out: pdb.Thrown_out,
		Timer:      &ti,
	}
	return &p
}

func (t *Timer) SetStart() {
	ti := time.Now()
	t.Start = &ti
}

func (t *Timer) SetFinish(d *time.Time, id_user int) bool {
	t.Finish = d
	return true
}

func (t *Timer) SetInterval(d *time.Time) {
	t.Interval = d
}

func (t *Timer) SetNotification(d *time.Time) {
	t.Finish = d
	t.Notif = true
}

func (t *Timer) SetTimer(finish *time.Time) {
	t.SetStart()
	t.Finish = finish
	t.Notif = false
}

func (t *Timer) String() string {
	if t.Notif {
		return "напомнить через: " + (t.Finish.Sub(time.Now())).String()
	} else {
		return "свежее ещё " + (t.Finish.Sub(time.Now())).String()
	}
}

func (p *Produkt) SetTimer(t *Timer) {
	p.Timer = t
}

func (p *Produkt) GetTimer() *Timer {
	return p.Timer
}

func (p *Produkt) String() string {
	return strconv.Itoa(p.Id) + " " + p.Name + p.Timer.String()
}
