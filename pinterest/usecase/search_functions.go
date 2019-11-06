package usecase

import (
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
)

func (USC *UseStruct) SearchPinsByTag(tag string) ([]models.PinForSearchResult, error) {
	var err error
	var params []interface{}
	params = append(params, tag)
	pins, err := USC.PRepository.SelectPinsByTag(consts.SELECTPinsByTag, params)
	if err != nil {
		return []models.PinForSearchResult{}, err
	}
	for _, pin := range pins {
		USC.Sanitizer.SanitPinForSearchResult(&pin)
	}
	return pins, nil
}
