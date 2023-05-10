package main

import "fmt"

type MyError struct{}

func (me MyError) Error() string {
	return "My error"
}

type Empty struct {
	Field string
}
type EmptyInterface interface {
	Foo() string
}

func main() {
	fmt.Println("returnError() == nil ?") // true
	fmt.Println(returnError() == nil)
	fmt.Println("returnPtrError() == nil ?") // true
	fmt.Println(returnPtrError() == nil)
	fmt.Println("returnCustomError() == nil ?") // false
	fmt.Println(returnCustomError() == nil)
	fmt.Println("returnPtrCustomError() == nil ?") // false
	fmt.Println(returnPtrCustomError() == nil)
	fmt.Println("returnMyError() == nil ?") // true
	fmt.Println(returnMyError() == nil)

	var e Empty
	fmt.Println("e == nil ?") // false, struct will be {}
	fmt.Println(e)

	var ei EmptyInterface
	fmt.Println("ei == nil ?") // true
	fmt.Println(ei == nil)
}

func returnError() error {
	var err error
	return err
}

func returnPtrError() *error {
	var err *error
	return err
}

func returnCustomError() error {
	var err MyError
	return err
}

// looking smart but it won't be nil!
func returnPtrCustomError() error {
	var err *MyError
	return err
}

func returnMyError() *MyError {
	return nil
}
