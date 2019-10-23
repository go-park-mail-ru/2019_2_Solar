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
	SELECTBoardById = "SELECT b.id, b.owner_id, b.title, b.description, b.category, b.createdTime, b.isDeleted " +
		"FROM sunrise.board as b WHERE b.id = $1"

	INSERTPin = "INSERT INTO sunrise.pin (owner_id, author_id, board_id, title, description, pindir, createdTime) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id"
	SELECTPinById = "SELECT p.id, p.owner_id, p.author_id, p.board_id, p.title, p.description, p.pindir, p.createdTime, p.isDeleted " +
		"FROM sunrise.pin as p WHERE p.id = $1"
	SELECTPinsByBoardId = "SELECT p.id, p.owner_id, p.author_id, p.board_id, p.title, p.description, p.pindir, p.createdTime, p.isDeleted " +
		"FROM sunrise.pin as p WHERE p.board_id = $1"

	INSERTNotice = "INSERT INTO sunrise.notice (user_id, receiver_id, message, createdTime) VALUES ($1,$2,$3,$4) RETURNING id"
)
