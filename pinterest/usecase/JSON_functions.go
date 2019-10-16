package usecase

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
)

func SetJsonData(data interface{}, infMsg string) models.OutJSON {
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

func SetResponseError(encoder *json.Encoder, msg string, err error) {
	//log.Printf("%s: %s", msg, err)
	data := SetJsonData(nil, msg)
	encoder.Encode(data)
}