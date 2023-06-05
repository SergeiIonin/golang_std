package main

import "fmt"

func main() {
	input0 := "()"
	input1 := "()[]{}"
	input2 := "(]"
	input3 := "(){}}{"
	input4 := ")(){}"
	input5 := "(){}("
	input6 := "([]"
	input7 := ""
	input8 := "("
	input9 := "([{}])"
	input10 := "[({])}"
	input11 := "{([])}"
	input12 := "{([]){}}"
	input13 := "({{{{}}}))"
	input14 := "{{{{}}})"
	input15 := "(()])}[}[}[]][}}[}{})][[(]({])])}}(])){)((){"
	input16 := "))"

	res15 := isValid(input15)
	res1 := isValid(input1)
	res13 := isValid(input13)
	res11 := isValid(input11)
	res4 := isValid(input4)
	res0 := isValid(input0)
	res2 := isValid(input2)
	res3 := isValid(input3)
	res5 := isValid(input5)
	res6 := isValid(input6)
	res7 := isValid(input7)
	res8 := isValid(input8)
	res9 := isValid(input9)
	res10 := isValid(input10)
	res12 := isValid(input12)
	res14 := isValid(input14)
	res16 := isValid(input16)

	fmt.Println(res0)  // true
	fmt.Println(res1)  // true
	fmt.Println(res2)  // false
	fmt.Println(res3)  // false
	fmt.Println(res4)  // false
	fmt.Println(res5)  // false
	fmt.Println(res6)  // false
	fmt.Println(res7)  // false
	fmt.Println(res8)  // false
	fmt.Println(res9)  // true
	fmt.Println(res10) // false
	fmt.Println(res11) // true
	fmt.Println(res12) // true
	fmt.Println(res13) // false
	fmt.Println(res14) // false
	fmt.Println(res15) // false
	fmt.Println(res16) // false

}

func isValid(s string) bool {
	runes := []rune(s)

	if len(runes)%2 != 0 || len(runes) == 0 {
		return false
	}

	visited := make(map[int]rune, len(runes)/2)

	for i := range runes {
		_, ok := visited[i]
		if ok {
			continue
		}
		closingBrackets := make([]rune, 0)
		opening := isOpening(runes[i])
		if opening {
			closingBrackets = append(closingBrackets, getClosingBracket(runes[i]))
		}
		j := i + 1
		for j < len(runes) {
			next := runes[j]
			visited[j] = next
			if isOpening(next) {
				closingBrackets = append(closingBrackets, getClosingBracket(next))
			} else {
				var close rune
				if len(closingBrackets) == 0 {
					return false
				} else {
					close = closingBrackets[len(closingBrackets)-1]
				}
				if next == close {
					closingBrackets = closingBrackets[:len(closingBrackets)-1]
				} else {
					return false
				}
			}
			if len(closingBrackets) == 0 {
				break
			}
			j++
			if j == len(runes) {
				return false
			}
		}
	}
	return true
}

func isOpening(bracket rune) bool {
	return bracket == '[' || bracket == '{' || bracket == '('
}

func getClosingBracket(open rune) rune {
	var res rune
	switch open {
	case '{':
		res = '}'
	case '[':
		res = ']'
	case '(':
		res = ')'
	}
	return res
}
