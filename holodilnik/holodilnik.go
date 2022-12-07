package holodilnik

import (
	"github.com/sashagitar/TBD_2281_kolesnikov/tree/go_lang/buy_list/produkt"
)

type Holodos struct {
	list map[string]produkt.Produkt
}

func (h *Holodos) Create() {
	h.list = make(map[string]produkt.Produkt)
}

func (h *Holodos) Add(produkt.Produkt) {
	h.list[produkt.Produkt.Name] = produkt.Produkt
}

func (h *Holodos) Delete(name string) {
	h.list[name]
}
