package jsonstatemachine

type TokenType int

const (
	OBJECT TokenType = iota
	ARRAY
	STRING
	NUMBER
	TRUE
	FALSE
	NULL
)
