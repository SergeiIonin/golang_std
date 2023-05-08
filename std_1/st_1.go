package main

import (
	"fmt"
	"math"
	"time"
)
import "rsc.io/quote"

func main() {
	fmt.Println("hola mundo!")

	fmt.Println("---VALUES---")
	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	fmt.Println("---VARIABLES---")
	var a = "initial"
	fmt.Println(a)

	var b, c = 1, 2 // declare 2 vars at once
	fmt.Println(b)
	fmt.Println(c)

	var emptyInt int
	fmt.Println(emptyInt) // if the type is known, unitialized var is zero-valued

	f := "apple" // shorthand for var f = "apple", it's only available inside funcs
	fmt.Println(f)
	//var unknown // won't compile

	fmt.Println("---USE OTHER MODULES---")
	fmt.Println(quote.Go())

	fmt.Println("---CONSTANTS---")
	const s string = "constant"

	const n = 500000000 // A const statement can appear anywhere a var statement can.
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d)) // A numeric constant has no type until it’s given one, such as by an explicit conversion

	fmt.Println(math.Sin(n)) // A number can be given a type by using it in a context that requires one, such as a
	// variable assignment or function call

	fmt.Println("---FOR---")
	i := 1
	for i <= 3 {
		fmt.Println("i = ", i)
		i++
	}

	for j := 7; j <= 9; j++ {
		fmt.Println("j = ", j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for i := 0; i <= 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("---IF/ELSE---")
	if 7%2 == 0 { // () aren't required, but the output should be in the {}
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 { // else is not required
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 { // A statement can precede conditionals; any variables declared in this statement are
		// available in the current and all subsequent branches.
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
	// there's no ternary operator in GO

	fmt.Println("---SWITCH---")
	k := 2
	fmt.Print("Write ", k, " as ")
	switch k {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: // we can pass multiple expressions
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	//switch without an expression is an alternate way to express if/else logic. Here we also show how
	//the case expressions can be non-constants.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	//A type switch compares types instead of values. You can use this to discover the type of
	//an interface value. In this example, the variable t will have the type corresponding to its clause.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Println("Don't know the type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("str")
}
