package usecase

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
)

func (USC UseStruct) SetJSONData(data interface{}, token string, infMsg string) models.OutJSON {
	user, ok := data.(models.User)
	if ok {
		outJSON := models.OutJSON{
			CSRFToken: token,
			BodyJSON: models.DataJSON{
				UserJSON: user,
				InfoJSON: infMsg,
			},
		}
		return outJSON
	}
	if anotherUser, ok := data.(models.AnotherUser); ok {
		outJSON := models.OutJSON{
			CSRFToken: token,
			BodyJSON: models.DataJSON{
				UserJSON: anotherUser,
				InfoJSON:  infMsg,
			},
		}
		return outJSON
	}
	if users, ok := data.([]models.User); ok {
		outJSON := models.OutJSON{
			CSRFToken: token,
			BodyJSON: models.DataJSON{
				UsersJSON: users,
				InfoJSON:  infMsg,
			},
		}
		return outJSON
	}
	if anotherUsers, ok := data.([]models.AnotherUser); ok {
		outJSON := models.OutJSON{
			CSRFToken: token,
			BodyJSON: models.DataJSON{
				UsersJSON: anotherUsers,
				InfoJSON:  infMsg,
			},
		}
		return outJSON
	}
	outJSON := models.OutJSON{
		CSRFToken: token,
		BodyJSON: models.DataJSON{
			InfoJSON: infMsg,
		},
	}
	return outJSON
}

func (USC UseStruct) SetResponseError(encoder *json.Encoder, msg string, err error) error {
	data := USC.SetJSONData(nil, "",  msg)
	if err := encoder.Encode(data); err != nil {
		return err
	}
	return nil
}
