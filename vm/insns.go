package vm

var insnsMap = map[string](func(int, int) int){
	"+": Add,
	"-": Subtract,
	"*": Multiply,
	"/": Divide,
}

func Add(left, right int) int {
	return left + right
}

func Subtract(left, right int) int {
	return left - right
}

func Multiply(left, right int) int {
	return left * right
}

func Divide(left, right int) int {
	return left / right
}
