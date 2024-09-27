package main

import (
	"strings"
)

type MultiError struct {
    errs []string
}
 
func (m *MultiError) Add(err error) {
    m.errs = append(m.errs, err.Error())
}
 
func (m *MultiError) Error() string {
    return strings.Join(m.errs, ";")
}