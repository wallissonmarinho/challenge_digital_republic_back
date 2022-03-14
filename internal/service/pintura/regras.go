package pintura

import (
	"errors"
	"math"

	"gopkg.in/guregu/null.v4"
)

type regraValidarArea struct {
	alturaParede        null.Float
	alturaPorta         null.Float
	totalPortasEJanelas null.Float
	totalParede         null.Float
}

func (r *regraValidarArea) Validate() error {

	if math.Round((r.alturaParede.Float64-r.alturaPorta.Float64)*100)/100 <= float64(0.30) {
		return errors.New("Altura da parede deve ser 30 centimentros maior que a altura da porta")
	}

	if math.Round(r.totalPortasEJanelas.Float64*100)/100 > math.Round((r.totalParede.Float64/2)*100)/100 {
		return errors.New("O total de área das portas e janelas deve ser no máximo 50 por cento da área de parede")
	}

	if r.totalParede.Float64 < float64(1) {
		return errors.New("O total de área da parede não pode ser menor que 1")
	}

	if r.alturaParede.Float64 > float64(15) {
		return errors.New("O total de área da parede não pode ser maior que 15")
	}

	return nil
}
