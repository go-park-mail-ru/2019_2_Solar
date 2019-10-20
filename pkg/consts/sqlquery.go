package consts

const (
	SelectUserIdByEmailSQLQuery                = "SELECT u.id from sunrise.users as u where u.email = $1"
	SelectUserIdUsernameEmailByUsernameOrEmail = "SELECT u.id, u.username, u.email from sunrise.users as u where u.username = $1 OR u.email = $2"
	ReadUserIdByUsernameSQLQuery               = "SELECT u.id from sunrise.users as u where u.username = $1"
	SelectUserUsernameByUsername               = "SELECT u.username from sunrise.users as u where u.username = $1"
	SelectAllUsers                             = "SELECT * from sunrise.users"
	UpdateUserByID                             = "UPDATE sunrise_db.sunrise.users SET username = $1, name = $2, 	surname = $3," +
		" hashpassword = $4,email = $5, age = $6, status = $7 where id = $8"
	UpdateUserAvatarDirByID = "UPDATE sunrise_db.sunrise.users SET avatardir = $1 where id = $2"
	InsertRegistrationQuery = "INSERT INTO sunrise.users (username, email, hashpassword)	values ($1,$2,$3) RETURNING id"
	InsertSessionQuery      = "INSERT INTO sunrise.usersessions (userid, cookiesvalue, cookiesexpiration)	values ($1,$2,$3) RETURNING id"

	ReadUserByCookieValueSQLQuery = "SELECT U.id, U.username, U.name, U.surname, U.hashpassword, U.email, U.age, U.status," +
		" U.avatardir, U.isactive from sunrise.Users as U JOIN sunrise.usersessions as s on U.id = s.userid " +
		"where s.cookiesvalue = $1"
	ReadCookiesExpirationByCookieValueSQLQuery = "SELECT s.cookiesvalue, s.cookiesexpiration from sunrise.usersessions" +
		" as s where s.cookiesvalue = $1"
	ReadUserByEmailSQLQuery = "SELECT U.id, U.username, U.name, U.surname, U.hashpassword, U.email, U.age, U.status," +
		" U.avatardir, U.isactive from sunrise.Users as U where U.email = $1"

	DeleteSessionByKey = "DELETE FROM sunrise.usersessions as s WHERE s.cookiesvalue = $1"
)
