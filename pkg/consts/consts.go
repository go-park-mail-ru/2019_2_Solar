package consts

const (
	LetterBytes        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"                                    // 62 possibilities
	LetterIdxBits      = 6                                                                                                   // 6 bits to represent 64 possibilities / indexes
	LetterIdxMask      = 1<<LetterIdxBits - 1                                                                                // All 1-bits, as many as letterIdxBits
	LoggerFormat       = "${time_rfc3339}, method = ${method}, uri = ${uri}, status = ${status}, remote_ip = ${remote_ip}\n" //Logger format for echo logger middleware
	HostAddress        = "127.0.0.1:8080"
	NumberOfPinsOnPage = 10
)
