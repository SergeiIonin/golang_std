package main

import "fmt"

func main() {
	res := gcdOfStrings("ABCABC", "ABC")
	fmt.Printf("res = %v\n", res) // ABC
}

func gcdOfStrings(str1 string, str2 string) string {
    checkDiv := func(str string, div string) bool {
        if len(str) % len(div) != 0 {
            return false
        }
        factor := len(div)
        runes := []rune(str)
        for i := 0; i < len(runes); i += factor {
            s := string(runes[i:i+factor])
            fmt.Printf("s = %s, div = %s, factor = %d \n", s, div, factor)
            if string(runes[i:i+factor]) != div {
                return false
            }
        }
        fmt.Printf("%s is a divisor of %s \n", div, str)
        return true
    }
    runes1 := []rune(str1)
    var gcd string
    for i := range runes1 {
        cand := string(runes1[:i+1])
        if len(cand) > len(str2) {
            break
        }
        fmt.Printf("cand = %s \n", cand)
        if checkDiv(str1, cand) && checkDiv(str2, cand) {
            gcd = cand
        }
    }
    return gcd
}