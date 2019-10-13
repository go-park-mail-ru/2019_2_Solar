package consts

const (
	LetterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" // 62 possibilities
	LetterIdxBits = 6                                                                // 6 bits to represent 64 possibilities / indexes
	LetterIdxMask = 1<<LetterIdxBits - 1                                             // All 1-bits, as many as letterIdxBits
)
