package consts

const (
	// 62 possibilities
	LetterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 6 bits to represent 64 possibilities / indexes
	LetterIdxBits = 6

	// All 1-bits, as many as letterIdxBits
	LetterIdxMask = 1<<LetterIdxBits - 1

	//Logger format for echo logger middleware
	LoggerFormat = "${time_rfc3339}, method = ${method}, uri = ${uri}," +
		" status = ${status}, remote_ip = ${remote_ip}\n"

	HostAddress        = "0.0.0.0:8080"
	ConnStr            = "host=db user=postgres password=7396 dbname=sunrise_db sslmode=disable"//"host=my_postgres user=postgres password=7396 dbname=sunrise_db sslmode=disable"
	NumberOfPinsOnPage = 10
)
