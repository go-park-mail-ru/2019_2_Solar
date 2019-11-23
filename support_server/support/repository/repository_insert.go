package repository

import (
	"github.com/go-park-mail-ru/2019_2_Solar/support_server/pkg/consts"
	"github.com/go-park-mail-ru/2019_2_Solar/support_server/pkg/models"
	"time"
)

func (RS *ReposStruct) InsertAdminSession(adminID uint64, cookieValue string, cookieExpires time.Time) (uint64, error) {
	var id uint64
	err := RS.DataBase.QueryRow(consts.INSERTAdminSession, adminID, cookieValue, cookieExpires).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (RS *ReposStruct) InsertSupportChatMessage(message models.NewChatMessage, createdTime time.Time) (uint64, error) {
	var id uint64
	err := RS.DataBase.QueryRow(consts.INSERTSupportChatMessage, message.IdSender, message.UserNameRecipient, message.Message, createdTime).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
