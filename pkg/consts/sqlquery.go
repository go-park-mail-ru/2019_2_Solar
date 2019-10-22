package consts

const (
	SELECTUserIdByEmail                        = "SELECT u.id from sunrise.users as u where u.email = $1"
	SELECTUserIdUsernameEmailByUsernameOrEmail = "SELECT u.id, u.username, u.email from sunrise.users as u where u.username = $1 OR u.email = $2"
	SELECTAllUsers                             = "SELECT * from sunrise.users"
	UPDATEUserByID                             = "UPDATE sunrise_db.sunrise.users SET username = $1, name = $2, 	surname = $3," +
		" hashpassword = $4,email = $5, age = $6, status = $7 where id = $8"
	UPDATEUserAvatarDirByID = "UPDATE sunrise_db.sunrise.users SET avatardir = $1 where id = $2"
	INSERTRegistration      = "INSERT INTO sunrise.users (username, email, hashpassword)	values ($1,$2,$3) RETURNING id"
	INSERTSession           = "INSERT INTO sunrise.usersessions (userid, cookiesvalue, cookiesexpiration)	values ($1,$2,$3) RETURNING id"

	SELECTUserByCookieValue = "SELECT U.id, U.username, U.name, U.surname, U.hashpassword, U.email, U.age, U.status," +
		" U.avatardir, U.isactive from sunrise.Users as U JOIN sunrise.usersessions as s on U.id = s.userid " +
		"where s.cookiesvalue = $1"
	SELECTCookiesExpirationByCookieValue = "SELECT s.cookiesvalue, s.cookiesexpiration from sunrise.usersessions" +
		" as s where s.cookiesvalue = $1"
	SELECTUserByEmail = "SELECT U.id, U.username, U.name, U.surname, U.hashpassword, U.email, U.age, U.status," +
		" U.avatardir, U.isactive from sunrise.Users as U where U.email = $1"

	DELETESessionByKey = "DELETE FROM sunrise.usersessions as s WHERE s.cookiesvalue = $1"

	SELECTCategoryByName = "SELECT c.name FROM sunrise.category as c WHERE c.name = $1"
	INSERTBoard = "INSERT INTO sunrise.board (owner_id, title, description, category, createdTime) VALUES ($1,$2,$3,$4,$5) RETURNING id"

	INSERTPin = "INSERT INTO sunrise.pin (owner_id, author_id, board_id, title, description, pindir, createdTime) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id"
)
