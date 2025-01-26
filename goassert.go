package goassert

/*
If the condition is not true, then panic with the given message. You can use fmt.Sprintf() to create more complex messages

Usage example:
func AddMoney(value int) {
	Assert(value > 0, fmt.Sprintf("%v is invalid", value))
	...
}
*/
func Assert(condition bool, message string) {
	if !condition {
		panic(message)
	}
}
