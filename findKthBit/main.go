package main
import "fmt"

func findKthBit(n int, k int) byte {
	s := []byte{'0'}
	for i := 2; i <= n; i++ {
		s = append(s, '1')
		n := len(s)
		// minus one to exclude '1'
		for j := n - 2; j >= 0; j-- {
			var b byte
			if s[j] == '1' {
				b = '0'
			} else {
				b = '1'
			}
			s = append(s, b)
		}
	}
	return s[k]
}


func main() {
	fmt.Println(findKthBit(3,1))
}