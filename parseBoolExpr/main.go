package main

import "fmt"

func main() {
	fmt.Println(parseBoolExpr("&(|(f))"))
}

// Top down
func parseBoolExpr(expression string) bool {
	idx := 0
	return parseExpr(&expression, &idx)
}

func parseExpr(expr *string, idx *int) bool {
	// extract current character
	currentChar := (*expr)[*idx]
	*idx++
	if currentChar == 't' {
		return true
	}
	if currentChar == 'f' {
		return false
	}

	if currentChar == '!' {
		*idx++ // skip '('
		rs := !parseExpr(expr, idx)
		*idx++ // skip ')'
		return rs
	}

	values := []bool{}
	// skip '('
	*idx++
	for (*expr)[*idx] != ')' {
		// skip ','
		if (*expr)[*idx] != ',' {
			// call recursively to parse the sub-expression
			values = append(values, parseExpr(expr, idx))
		} else {
			*idx++
		}
	}

	*idx++ // skip ')'

	if currentChar == '&' {
		// if all values are true, return true
		for _, v := range values {
			if !v {
				return false
			}
		}
		return true
	}

	if currentChar == '|' {
		// if any value is true, return true
		for _, v := range values {
			if v {
				return true
			}
		}
		return false
	}

	return false
}
