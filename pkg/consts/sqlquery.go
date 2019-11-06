package consts

const (
	SELECTUserIDByEmail                        = "SELECT u.id from sunrise.user as u where u.email = $1"
	SELECTUserIDUsernameEmailByUsernameOrEmail = "SELECT u.id, u.username, u.email from sunrise.user as u where u.username = $1 OR u.email = $2"
	SELECTAllUsers                             = "SELECT * from sunrise.user"
	UPDATEUserByID                             = "UPDATE sunrise_db.sunrise.user SET username = $1, name = $2, 	surname = $3," +

		" hashpassword = $4,email = $5, age = $6, status = $7 where id = $8"
	UPDATEUserAvatarDirByID = "UPDATE sunrise_db.sunrise.user SET avatardir = $1 where id = $2"
	INSERTRegistration      = "INSERT INTO sunrise.user (username, email, hashpassword, salt, created_time)	values ($1,$2,$3,$4,$5) RETURNING id"
	INSERTSession           = "INSERT INTO sunrise.usersession (userid, cookiesvalue, cookiesexpiration)	values ($1,$2,$3) RETURNING id"

	SELECTUserByCookieValue = "SELECT U.id, U.username, U.name, U.surname, U.hashpassword, U.email, U.age, U.status," +
		" U.avatardir, U.isactive, U.salt, U.created_time from sunrise.User as U JOIN sunrise.usersession as s on U.id = s.userid " +
		"where s.cookiesvalue = $1"
	SELECTCookiesExpirationByCookieValue = "SELECT s.cookiesvalue, s.cookiesexpiration from sunrise.usersession" +
		" as s where s.cookiesvalue = $1"
	SELECTUserByUsername = "SELECT U.id, U.username, U.name, U.surname, U.hashpassword, U.email, U.age, U.status," +
		" U.avatardir, U.isactive, U.salt, U.created_time from sunrise.User as U where U.username = $1"
	SELECTUserByEmail = "SELECT U.id, U.username, U.name, U.surname, U.hashpassword, U.email, U.age, U.status," +
		" U.avatardir, U.isactive, U.salt, U.created_time from sunrise.User as U where U.email = $1"

	DELETESessionByKey = "DELETE FROM sunrise.usersession as s WHERE s.cookiesvalue = $1"

	SELECTCategoryByName = "SELECT c.name FROM sunrise.category as c WHERE c.name = $1"
	INSERTBoard          = "INSERT INTO sunrise.board (owner_id, title, description, category, createdTime) VALUES ($1,$2,$3,$4,$5) RETURNING id"
	SELECTBoardByID      = "SELECT b.id, b.owner_id, b.title, b.description, b.category, b.createdTime, b.isDeleted " +
		"FROM sunrise.board as b WHERE b.id = $1"

	INSERTPin = "INSERT INTO sunrise.pin (owner_id, author_id, board_id, title, description, pindir, createdTime)" +
		" VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id"
	SELECTPinByID = "SELECT p.id, p.owner_id, p.author_id, p.board_id, p.title," +
		" p.description, p.pindir, p.createdTime, p.isDeleted " +
		"FROM sunrise.pin as p WHERE p.id = $1"
	SELECTPinsByBoardID = "SELECT p.id, p.owner_id, p.author_id, p.board_id, p.title," +
		" p.description, p.pindir, p.createdTime, p.isDeleted " +
		"FROM sunrise.pin as p WHERE p.board_id = $1"

	SELECTNewPinsByNumber = "SELECT p.id, p.pindir FROM (select id, pindir, isdeleted, ROW_NUMBER() OVER (ORDER BY createdtime) " +
		"from sunrise.pin WHERE isdeleted = false) as p WHERE p.ROW_NUMBER BETWEEN 0 AND $1;"
	SELECTMyPinsByNumber = "SELECT p.id, p.pindir FROM (select id, pindir, isdeleted, ROW_NUMBER() OVER (ORDER BY createdtime) " +
		"from sunrise.pin WHERE owner_id = $2 AND isdeleted = false) as p WHERE p.ROW_NUMBER BETWEEN 0 AND $1;"
	SELECTSubscribePinsByNumber = "SELECT p.id, p.pindir FROM (select id, pindir, isdeleted, ROW_NUMBER() OVER (ORDER BY createdtime) " +
		"from sunrise.pin join sunrise.subscribe as s on s.subscriber_id = $2 " +
		"AND s.followee_id = pin.owner_id " +
		"AND isdeleted = false) as p " +
		"WHERE p.ROW_NUMBER BETWEEN 0 AND $1;"
	SELECTComments = "SELECT c.text, u.username, c.created_time FROM sunrise.comment as c join sunrise.pin as p on p.id = $1 " +
		"join sunrise.user as u on u.id = p.owner_id where c.pin_id = $1"

	INSERTNotice          = "INSERT INTO sunrise.notice (user_id, receiver_id, message, createdTime) VALUES ($1,$2,$3,$4) RETURNING id"
	INSERTComment         = "INSERT INTO sunrise.comment (pin_id, text, author_id, created_time) VALUES ($1,$2,$3,$4) RETURNING id"
	INSERTSubscribeByName = "INSERT INTO sunrise.subscribe (subscriber_id, followee_id) " +
		"select $1, u.id from sunrise.user as u " +
		"where u.username = $2 " +
		"RETURNING id;"

	INSERTChatMessage = "INSERT INTO sunrise.chat_message (sender_id, receiver_id, text, send_time) " +
		"SELECT $1, u.id, $3, $4 from sunrise.user as u where u.username = $2 RETURNING id"

	DELETESubscribeByName = "DELETE FROM sunrise.subscribe as s WHERE s.subscriber_id = $1 and s.followee_id IN " +
		"(select u.id from sunrise.user as u " +
		"where u.username = $2);"

	SELECTPinsByTag = "SELECT DISTINCT p.id, p.pindir, p.title FROM sunrise.pin as p " +
		"JOIN sunrise.pinandtag as pt ON p.id = pt.pin_id " +
		"WHERE pt.tag_name = $1 AND p.isdeleted = false;"

	SELECTSessionByCookieValue = "SELECT s.id, s.userid FROM sunrise.usersession as s " +
		"WHERE s.cookiesvalue = $1;"
)
