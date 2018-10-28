package middleware

type MidData struct {
	key   string
	value string
}

type Middlehandler func(stack []MidData) []MidData

var Middleware []Middlehandler

func ExecuteMiddleware() []MidData {
	var stack []MidData
	for _, middle := range Middleware {
		stack = append(stack, middle(stack)...)
	}
	return stack
}
