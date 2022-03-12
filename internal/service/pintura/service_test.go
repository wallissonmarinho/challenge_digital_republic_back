package pintura

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wallissonmarinho/challenge_digital_republic/internal/domain"
	"gopkg.in/guregu/null.v4"
)

func TestObterLatasdeTinta(t *testing.T) {
	t.Run("Teste deveria obter latas de tinta de acordo com a quantidade de litros de tinta informado", func(t *testing.T) {
		latas := domain.Lata{
			TotalLata18l:  null.IntFrom(1),
			TotalLata0_5l: null.IntFrom(2),
		}
		resp, err := ObterLatasdeTinta(null.FloatFrom(95))

		assert.Nil(t, err)
		assert.Equal(t, latas.TotalLata18l, resp.TotalLata18l)
		assert.Equal(t, latas.TotalLata0_5l, resp.TotalLata0_5l)
	})

	t.Run("Teste não deveria obter latas de tinta se não for informado a quantidade de tinta", func(t *testing.T) {
		latas := domain.Lata{}
		resp, err := ObterLatasdeTinta(null.FloatFrom(0))

		assert.NotNil(t, err)
		assert.Equal(t, latas, resp)
		assert.EqualError(t, err, "Necessário informar a area para obter a quantidade de latas de tinta")
	})
}

func TestCalcularArea(t *testing.T) {
	t.Run("Teste deveria obter metros quadrados da area calculada", func(t *testing.T) {
		paredes := []domain.Parede{
			{
				Altura:  null.FloatFrom(2.3),
				Largura: null.FloatFrom(4),
				Portas:  null.IntFrom(1),
				Janelas: null.IntFrom(1),
			},
			{
				Altura:  null.FloatFrom(2.3),
				Largura: null.FloatFrom(4),
				Portas:  null.IntFrom(1),
				Janelas: null.IntFrom(1),
			},
			{
				Altura:  null.FloatFrom(2.3),
				Largura: null.FloatFrom(3),
				Portas:  null.IntFrom(1),
				Janelas: null.IntFrom(1),
			},
			{
				Altura:  null.FloatFrom(2.3),
				Largura: null.FloatFrom(3),
			},
		}

		resp, err := CalcularArea(paredes)

		assert.Nil(t, err)
		assert.Equal(t, resp, null.FloatFrom(20.44))
	})

	t.Run("Teste não deveria obter metros quadrados da area calculada caso a porta não for 30 centimetros ou mais que a parede", func(t *testing.T) {
		paredes := []domain.Parede{
			{
				Altura:  null.FloatFrom(2.3),
				Largura: null.FloatFrom(4),
				Portas:  null.IntFrom(1),
				Janelas: null.IntFrom(1),
			},
			{
				Altura:  null.FloatFrom(2.3),
				Largura: null.FloatFrom(4),
				Portas:  null.IntFrom(1),
				Janelas: null.IntFrom(1),
			},
			{
				Altura:  null.FloatFrom(2.2),
				Largura: null.FloatFrom(3),
				Portas:  null.IntFrom(1),
				Janelas: null.IntFrom(1),
			},
			{
				Altura:  null.FloatFrom(2.3),
				Largura: null.FloatFrom(3),
			},
		}

		_, err := CalcularArea(paredes)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "Altura da parede deve ser 30 centimentros maior que a altura da porta")
	})

	t.Run("Teste não deveria obter metros quadrados da area calculada caso a area total de portas e janelas for maior que 50 por cento do que a area total da parede", func(t *testing.T) {
		paredes := []domain.Parede{
			{
				Altura:  null.FloatFrom(2.3),
				Largura: null.FloatFrom(4),
				Portas:  null.IntFrom(1),
				Janelas: null.IntFrom(1),
			},
			{
				Altura:  null.FloatFrom(2.3),
				Largura: null.FloatFrom(4),
				Portas:  null.IntFrom(1),
				Janelas: null.IntFrom(1),
			},
			{
				Altura:  null.FloatFrom(2.2),
				Largura: null.FloatFrom(3),
				Portas:  null.IntFrom(2),
				Janelas: null.IntFrom(1),
			},
			{
				Altura:  null.FloatFrom(2.3),
				Largura: null.FloatFrom(3),
			},
		}

		_, err := CalcularArea(paredes)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "O total de área das portas e janelas deve ser no máximo 50 por cento da área de parede")
	})
}
