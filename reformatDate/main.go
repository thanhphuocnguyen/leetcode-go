package main

import (
	"fmt"
	"strings"
)

func reformatDate(date string) string {
	monthMp := map[string]string{
		"Jan": "01",
		"Feb": "02",
		"Mar": "03",
		"Apr": "04",
		"May": "05",
		"Jun": "06",
		"Jul": "07",
		"Aug": "08",
		"Sep": "09",
		"Oct": "10",
		"Nov": "11",
		"Dec": "12",
	}
	paths := strings.Split(date, " ")
	day := paths[0][:len(paths[0])-2]
	if len(day) < 2 {
		day = "0" + day
	}
	return paths[2] + "-" + monthMp[paths[1]] + "-" + day
}

func main() {
	fmt.Println(reformatDate("6th Jun 1933"))
	fmt.Println(reformatDate("20th Oct 2052"))
}
