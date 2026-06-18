package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	slashCnt := 0
	paths := make([]string, 0)
	tempPath := []byte{}
	for i, rn := range path {
		if rn == '/' {
			slashCnt++
			if len(tempPath) > 0 {
				if tempPath[0] == '.' {
					dotCnt := 1
					for j := 1; j < len(tempPath); j++ {
						if tempPath[j] == '.' {
							dotCnt++
						} else {
							dotCnt = 0
							break
						}
					}
					if dotCnt == 0 || dotCnt > 2 {
						paths = append(paths, string(tempPath))
					} else if dotCnt == 2 {
						if len(paths) > 0 {
							paths = paths[:len(paths)-1]
						}
					}
				} else {
					paths = append(paths, string(tempPath))
				}
				tempPath = []byte{}
			}
		} else {
			tempPath = append(tempPath, path[i])
		}
	}
	if len(tempPath) > 0 {
		if tempPath[0] == '.' {
			dotCnt := 1
			for j := 1; j < len(tempPath); j++ {
				if tempPath[j] == '.' {
					dotCnt++
				} else {
					dotCnt = 0
					break
				}
			}
			if dotCnt == 0 || dotCnt > 2 {
				paths = append(paths, string(tempPath))
			} else if dotCnt == 2 {
				if len(paths) > 0 {
					paths = paths[:len(paths)-1]
				}
			}
		} else {
			paths = append(paths, string(tempPath))
		}
		tempPath = []byte{}
	}
	return "/" + strings.Join(paths, "/")
}

func main() {
	fmt.Println(simplifyPath("/."))
	fmt.Println(simplifyPath("/../"))
	fmt.Println(simplifyPath("/.../a/../b/c/../d/./"))
	fmt.Println(simplifyPath("/../..ga/b/.f..d/..../e.baaeeh./.a"))
	fmt.Println(simplifyPath("/home/"))
	fmt.Println(simplifyPath("/home//foo/"))
	fmt.Println(simplifyPath("/home/user/Documents/../Pictures"))
}
