package usecase

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
)

func (USC UsecaseStruct) SetJSONData(data interface{}, infMsg string) models.OutJSON {
	user, ok := data.(models.User)
	if ok {
		outJSON := models.OutJSON{
			BodyJSON: models.DataJSON{
				UserJSON: user,
				InfoJSON: infMsg,
			},
		}
		return outJSON
	}
	if users, ok := data.([]models.User); ok {
		outJSON := models.OutJSON{
			BodyJSON: models.DataJSON{
				UsersJSON: users,
				InfoJSON:  infMsg,
			},
		}
		return outJSON
	}
	outJSON := models.OutJSON{
		BodyJSON: models.DataJSON{
			InfoJSON: infMsg,
		},
	}
	return outJSON
}

func (USC UsecaseStruct) SetResponseError(encoder *json.Encoder, msg string, err error) error {
	data := USC.SetJSONData(nil, msg)
	if err := encoder.Encode(data); err != nil {
		return err
	}
	return nil
}
