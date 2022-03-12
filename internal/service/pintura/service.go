package pintura

import (
	"context"
	"errors"
	"math"

	"github.com/go-kit/log"
	"github.com/wallissonmarinho/challenge_digital_republic/internal/domain"
	"gopkg.in/guregu/null.v4"
)

type PinturaService interface {
	ObterQuantidadeDeLatasdeTinta(ctx context.Context, Paredes []domain.Parede) (domain.Lata, error)
}

type pinturaService struct {
	logger log.Logger
}

func NewPinturaService(logger log.Logger) *pinturaService {
	return &pinturaService{
		logger: logger,
	}
}

func (s *pinturaService) ObterQuantidadeDeLatasdeTinta(ctx context.Context, Paredes []domain.Parede) (domain.Lata, error) {
	area, err := CalcularArea(Paredes)
	if err != nil {
		return domain.Lata{}, err
	}

	latas, err := ObterLatasdeTinta(area)
	if err != nil {
		return domain.Lata{}, err
	}

	return latas, nil
}

func CalcularArea(Paredes []domain.Parede) (null.Float, error) {
	area := null.FloatFrom(0)
	paredes := null.FloatFrom(0)
	portasEJanelas := null.FloatFrom(0)
	janelaLargura := null.FloatFrom(2)
	janelaAltura := null.FloatFrom(1.2)
	portaLargura := null.FloatFrom(0.8)
	portaAltura := null.FloatFrom(1.9)

	for _, parede := range Paredes {
		paredes = null.FloatFrom(paredes.Float64 + (parede.Largura.Float64 * parede.Altura.Float64))
		portasEJanelas = null.FloatFrom(
			portasEJanelas.Float64 +
				(janelaLargura.Float64*janelaAltura.Float64)*float64(parede.Janelas.Int64) +
				(portaLargura.Float64*portaAltura.Float64)*float64(parede.Portas.Int64))

		regra := &regraValidarArea{
			alturaParede:        parede.Altura,
			alturaPorta:         portaAltura,
			totalPortasEJanelas: portasEJanelas,
			totalParede:         paredes,
		}

		err := regra.Validate()
		if err != nil {
			return null.FloatFrom(0), err
		}
	}

	area = null.FloatFrom(math.Round((paredes.Float64-portasEJanelas.Float64)*100) / 100)

	return area, nil
}

func ObterLatasdeTinta(metros null.Float) (domain.Lata, error) {

	latas := domain.Lata{}

	if metros.Float64 == float64(0) || !metros.Valid {
		return domain.Lata{}, errors.New("Necessário informar a area para obter a quantidade de latas de tinta")
	}

	Litros := null.FloatFrom(math.Round((metros.Float64/5)*100) / 100)

	latas.Litros = Litros

	for indice, tamanho := range domain.Latas() {

		for Litros.Float64 >= tamanho || (len(domain.Latas())-1 == indice && Litros.Float64 > 0) {
			if tamanho == float64(18) {
				latas.TotalLata18l = null.IntFrom(latas.TotalLata18l.Int64 + 1)
				Litros = null.FloatFrom(math.Round((Litros.Float64-tamanho)*100) / 100)
			}
			if tamanho == float64(3.6) {
				latas.TotalLata3_6l = null.IntFrom(latas.TotalLata3_6l.Int64 + 1)
				Litros = null.FloatFrom(math.Round((Litros.Float64-tamanho)*100) / 100)
			}
			if tamanho == float64(2.5) {
				latas.TotalLata2_5l = null.IntFrom(latas.TotalLata2_5l.Int64 + 1)
				Litros = null.FloatFrom(math.Round((Litros.Float64-tamanho)*100) / 100)
			}
			if tamanho == float64(0.5) {
				latas.TotalLata0_5l = null.IntFrom(latas.TotalLata0_5l.Int64 + 1)
				Litros = null.FloatFrom(math.Round((Litros.Float64-tamanho)*100) / 100)
			}
			continue
		}
	}
	return latas, nil
}
