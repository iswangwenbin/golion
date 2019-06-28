package errorsx

const (
	CodeBadArgument     = 100
	CodeJsonDecodeError = 101
)

var codeText = map[int]string{
	CodeBadArgument:     "Bad Argument",
	CodeJsonDecodeError: "Json Decode Error",
}

func CodeText(code int) string {
	return codeText[code]
}
