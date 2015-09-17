package gps

type TokenIdent int8

var (
	TokenIdentWasPrefixed TokenIdent = -1
	TokenIdentError       TokenIdent = 0
	TokenIdentIoEOF       TokenIdent = 1
	TokenIdentSlashR      TokenIdent = 2
	TokenIdentSlashN      TokenIdent = 3
	TokenIdentLine        TokenIdent = 4
)

type Token struct {
	Error error
	Runes []rune
	Ident TokenIdent
}
