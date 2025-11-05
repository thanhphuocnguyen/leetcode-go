package main

func kthCharacter(k int) byte {
	str := backtrack("a", k)
	return str[k-1]
}

func backtrack(str string, k int) string {
	if len(str) >= k {
		return str
	}
	charArr := make([]byte, len(str))
	for i := range str {
		charArr[i] = str[i] + 1
	}
	return backtrack(str+string(charArr), k)
}

func main() {
	// Example usage
	k := 5
	result := kthCharacter(k)
	println("The", k, "th character is:", string(result))
}

// The function kthCharacter returns the k-th character in a generated string
