package main

import (
	"fmt"
)

func pyramidTransition(bottom string, allowed []string) bool {
	var backtrack func(baseIdx int, base []byte, buildingBase []byte) bool
	mp := make(map[[2]byte][]byte)
	for _, str := range allowed {
		lr := [2]byte{str[0], str[1]}
		if _, ok := mp[lr]; !ok {
			mp[lr] = []byte{}
		}
		mp[lr] = append(mp[lr], str[2])
	}
	backtrack = func(baseIdx int, base []byte, buildingBase []byte) bool {
		if len(base) == 1 {
			return true
		}
		if baseIdx == len(base)-1 {
			return backtrack(0, buildingBase, []byte{})
		}
		lr := [2]byte{base[baseIdx], base[baseIdx+1]}
		for _, str := range mp[lr] {
			buildingBase = append(buildingBase, str)
			if backtrack(baseIdx+1, base, buildingBase) {
				return true
			}
			buildingBase = buildingBase[:len(buildingBase)-1]
		}
		return false
	}

	return backtrack(0, []byte(bottom), []byte{})
}

func main() {
	fmt.Println(pyramidTransition("DBCDA", []string{"DBD", "BCC", "CDD", "DAD", "DDA", "AAC", "CCA", "BCD"}))
	fmt.Println(pyramidTransition("AAAA", []string{"AAB", "AAC", "BCD", "BBE", "DEF"}))
	fmt.Println(pyramidTransition("BCD", []string{"BCC", "CDE", "CEA", "FFF"}))
}
