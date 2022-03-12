package domain

import (
	"sort"

	"gopkg.in/guregu/null.v4"
)

type Parede struct {
	Altura  null.Float `json:"altura"`
	Largura null.Float `json:"largura"`
	Portas  null.Int   `json:"portas"`
	Janelas null.Int   `json:"janelas"`
}

type Lata struct {
	TotalLata18l  null.Int   `json:"total_lata18l"`
	TotalLata3_6l null.Int   `json:"total_lata3_6l"`
	TotalLata2_5l null.Int   `json:"total_lata2_5l"`
	TotalLata0_5l null.Int   `json:"total_lata0_5l"`
	Litros        null.Float `json:"litros"`
}

func Latas() [4]float64 {

	latas := [4]float64{18, 3.6, 2.5, 0.5}

	sort.SliceStable(latas[:], func(i, j int) bool {
		return latas[i] > latas[j]
	})

	return latas
}
