package consts

const (
	ReadUserIdByEmailSQLQuery         = "SELECT u.id from sunrise.users as u where u.email = $1"
	ReadUserIdByUsernameEmailSQLQuery = "SELECT u.id, u.username, u.email from sunrise.users as u where u.username = $1 AND u.email = $2"
	ReadUserIdByUsernameSQLQuery      = "SELECT u.id from sunrise.users as u where u.username = $1"
	SelectUserUsernameByUsername = "SELECT u.username from sunrise.users as u where u.username = $1"
	SelectAllUsers = "SELECT * from sunrise.users"

	InsertRegistrationQuery           = "INSERT INTO sunrise.users (username, email, hashpassword)	values ($1,$2,$3) RETURNING id"
	InsertSessionQuery                = "INSERT INTO sunrise.usersessions (userid, cookiesvalue, cookiesexpiration)	values ($1,$2,$3) RETURNING id"

	ReadUserByCookieValueSQLQuery = "SELECT U.id, U.username, U.name, U.surname, U.hashpassword, U.email, U.age, U.status," +
		" U.avatardir, U.isactive from sunrise.Users as U JOIN sunrise.usersessions as s on U.id = s.userid " +
		"where s.cookiesvalue = $1"
	ReadCookiesExpirationByCookieValueSQLQuery = "SELECT s.cookiesvalue, s.cookiesexpiration from sunrise.usersessions" +
		" as s where s.cookiesvalue = $1"
	ReadUserByEmailSQLQuery = "SELECT U.id, U.username, U.name, U.surname, U.hashpassword, U.email, U.age, U.status," +
		" U.avatardir, U.isactive from sunrise.Users as U where U.email = $1"

	DeleteSessionByKey = "DELETE FROM sunrise.usersessions as s WHERE s.cookiesvalue = $1"
)
