package main

import "fmt"

type Codec struct {
	ids    []int
	id     int
	UrlMap map[string]string
}

var ENCODE_CHARS = []byte{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
}

func Constructor() Codec {
	fmt.Println(len(ENCODE_CHARS))
	return Codec{
		ids:    make([]int, 0),
		id:     1,
		UrlMap: make(map[string]string),
	}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	currId := this.id
	this.ids = append(this.ids, this.id+1)
	bs := make([]byte, 0)
	for currId > 0 {
		bs = append(bs, ENCODE_CHARS[currId%62])
		currId /= 62
	}
	shortUrl := string(bs)
	this.UrlMap[shortUrl] = longUrl
	this.id++
	return shortUrl
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	return this.UrlMap[shortUrl]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */

func main() {
	codec := Constructor()
	fmt.Println(codec.encode("https://leetcode.com/problems/design-tinyurl"))
	fmt.Println(codec.encode("https://leetcode.com/problems/design-tinyurl"))
	fmt.Println(codec.encode("https://leetcode.com/problems/design-tinyurl"))
	fmt.Println(codec.encode("https://leetcode.com/problems/design-tinyurl"))
	fmt.Println(codec.encode("https://leetcode.com/problems/design-tinyurl"))
	fmt.Println(codec.encode("https://leetcode.com/problems/design-tinyurl"))
	fmt.Println(codec.encode("https://leetcode.com/problems/design-tinyurl"))
	fmt.Println(codec.encode("https://leetcode.com/problems/design-tinyurl"))
	fmt.Println(codec.encode("https://leetcode.com/problems/design-tinyurl"))
	fmt.Println(codec.encode("https://leetcode.com/problems/design-tinyurl"))
	fmt.Println(codec.encode("https://leetcode.com/problems/design-tinyurl"))
	fmt.Println(codec.encode("https://leetcode.com/problems/design-tinyurl"))
	fmt.Println(codec.encode("https://leetcode.com/problems/design-tinyurl"))
	fmt.Println(codec.encode("https://leetcode.com/problems/design-tinyurl"))
	fmt.Println(codec.encode("https://leetcode.com/problems/design-tinyurl"))
	fmt.Println(codec.encode("https://leetcode.com/problems/design-tinyurl"))
	fmt.Println(codec.decode("1"))

}
