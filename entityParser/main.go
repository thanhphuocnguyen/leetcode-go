package main

import "fmt"

func entityParser(text string) string {
	mp := map[string]byte{
		"&quot;":  '"',
		"&apos;":  '\'',
		"&amp;":   '&',
		"&gt;":    '>',
		"&lt;":    '<',
		"&frasl;": '/',
	}
	bs := make([]byte, 0, len(text))
	tempSign := make([]byte, 0, len(text))

	for i := 0; i < len(text); i++ {
		if len(tempSign) > 0 || text[i] == '&' {
			if len(tempSign) > 0 && text[i] == '&' {
				bs = append(bs, tempSign...)
				tempSign = tempSign[:0]
			}
			tempSign = append(tempSign, text[i])
			if text[i] == ';' {
				if val, ok := mp[string(tempSign)]; ok {
					bs = append(bs, val)
				} else {
					bs = append(bs, tempSign...)
				}
				tempSign = tempSign[:0]
			}
		} else {
			bs = append(bs, text[i])
		}
	}
	if len(tempSign) > 0 {
		bs = append(bs, tempSign...)
	}
	return string(bs)
}

func main() {
	fmt.Println(entityParser("&a&gt;"))
	fmt.Println(entityParser("&&gt;"))
	fmt.Println(entityParser("&amp; is an HTML entity but &ambassador; is not."))
}
