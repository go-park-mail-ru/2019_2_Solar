package repository

import "github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"

func CombineInsertRegistrationQuery(username, email, password string) string {
	return consts.InsertRegistrationQuery + username + consts.CommaQuery + email + consts.CommaQuery + password + consts.EndInsertQuery
}

func CombineInsertSessionQuery(userid, cookiesvalue, cookiesexpiration string) string {
	return consts.InsertSessionQuery + userid + consts.CommaQuery + cookiesvalue + consts.CommaQuery + cookiesexpiration + consts.EndInsertQuery
}
