package main

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	adj := make([][]byte, 26)
	for i := range 26 {
		adj[i] = make([]byte, 0)
	}
	for i := range s1 {
		c1, c2 := s1[i]-'a', s2[i]-'a'
		adj[c1] = append(adj[c1], c2)
		adj[c2] = append(adj[c2], c1)
	}
	ans := make([]byte, len(baseStr))
	for i := range baseStr {
		visited := make([]bool, 26)
		ans[i] = (dfs(adj, baseStr[i]-'a', visited) + 'a')
	}
	return string(ans)
}

func dfs(adj [][]byte, curr byte, visited []bool) byte {
	visited[curr] = true
	rs := curr
	for _, nei := range adj[curr] {
		if !visited[nei] {
			rs = min(curr, dfs(adj, nei, visited))
		}
	}
	return rs
}

func main() {
	println(smallestEquivalentString("leetcode", "programs", "sourcecode"))
}
