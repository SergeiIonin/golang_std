package main

import (
	"errors"
	"fmt"
)

func main() {
	u := User{30, "Bob"}
	if err := u.Validate(); err != nil {
		fmt.Printf("Error validating customer: %v \n", err) // will return Error validating customer: <nil>
	}
	if err := u.Validate2(); err != nil {
		fmt.Printf("Error validating customer: %v", err) // will return Error validating customer: <nil>
	} else {
		fmt.Printf("User is valid")
	}
}

type User struct {
	Age int
	Name string
}

func (u User) Validate() error {
    var m *MultiError
 
    if u.Age < 0 {
        m = &MultiError{}
        m.Add(errors.New("age is negative"))
    }
    if u.Name == "" {
        if m == nil {
            m = &MultiError{}
        }
        m.Add(errors.New("name is nil"))
    }
	
	// In this case m (nil-pointer) is wrapped into non-nil interface (Error). Hence we'll always receive non-nil error calling Validate()
    return m
}

func (u User) Validate2() error {
	var m *MultiError
 
    if u.Age < 0 {
        m = &MultiError{}
        m.Add(errors.New("age is negative"))
    }
    if u.Name == "" {
        if m == nil {
            m = &MultiError{}
        }
        m.Add(errors.New("name is nil"))
    }
	
	// explicitly return m in case of an error and nil otherwise
	if m != nil {
		return m
	}
    return nil
}