package main

import "fmt"

func main() {

	// var declares 1 or more variables.
	var a = "initial"
	fmt.Println(a)

	// You can declare multiple variables at once.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go will infer the type of initialized variables.
	var d = true
	fmt.Println(d)

	// Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.
	var e int
	fmt.Println(e)

	// The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "short" in this case.
	f := "short"
	fmt.Println(f)

	// 声明并赋值运算符(:=)左边至少得有一个变量是没有定义的。
	// 左边变量有没有定义的标准是在当前名字域范围内，即当前语句块，或者当前函数；而不管是否在外层名字域范围内已经定义。
	var g int = 0
	fmt.Printf("&g=%p\n", &g)
	fmt.Printf("g=%b\n", g)
	g, h := foo()
	fmt.Printf("&g=%p, &h=%p\n", &g, &h)
}

func foo() (int, int) {
	return 100, 200
}
