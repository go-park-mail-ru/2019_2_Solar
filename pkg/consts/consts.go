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

	HostAddress        = "127.0.0.1:8080"
	NumberOfPinsOnPage = 10
)
