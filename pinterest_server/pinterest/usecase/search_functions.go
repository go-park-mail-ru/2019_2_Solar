package usecase

import (
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
)

func (USC *UseStruct) SearchPinsByTag(tag string) ([]models.PinDisplay, error) {
	pins, err := USC.PRepository.SelectPinsByTag(tag)
	if err != nil {
		return []models.PinDisplay{}, err
	}
	for _, pin := range pins {
		USC.Sanitizer.SanitPinDisplay(&pin)
	}
	return pins, nil
}
