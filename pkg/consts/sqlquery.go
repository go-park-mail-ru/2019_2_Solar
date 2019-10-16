package consts

const (
	FindEmailSQLQuery       = "SELECT u.id from testschema.users as u where u.email = '$1'"
	FindUsernameSQLQuery    = "SELECT u.id from testschema.users as u where u.username = '$1'"
	InsertRegistrationQuery = "INSERT INTO testschema.users (username, email, hashpassword)	values ('"
	InsertSessionQuery      = "INSERT INTO testschema.users (username, email, hashpassword)	values ('$1','$2','$3')"

	QueryReadUserByCookie = "SELECT U.id, U.username, U.name, U.surname, U.hashpassword, U.email, U.age, U.status," +
		" U.avatardir, U.isactive from testschema.Users as U JOIN testschema.usersessions as s on U.id = s.userid " +
		"where s.cookiesvalue = "
	QueryCookiesExpiration = "SELECT s.cookiesvalue, s.cookiesexpiration from testschema.usersessions" +
		" as s where s.cookiesvalue = "
	QueryReadUserByEmail = "SELECT U.id, U.username, U.name, U.surname, U.hashpassword, U.email, U.age, U.status," +
		" U.avatardir, U.isactive from testschema.Users as U where U.email = "
)
