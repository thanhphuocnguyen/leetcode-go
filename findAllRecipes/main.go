package main

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	suppliesMap := make(map[string]bool)
	for _, s := range supplies {
		suppliesMap[s] = true
	}
	recipeToIdx := make(map[string]int)
	graph := make(map[string][]string)
	indegree := make([]int, len(recipes))
	for i, rp := range recipes {
		recipeToIdx[rp] = i
		for _, ing := range ingredients[i] {
			if !suppliesMap[ing] {
				if graph[ing] == nil {
					graph[ing] = []string{rp}
				} else {
					graph[ing] = append(graph[ing], rp)
				}
				indegree[i]++
			}
		}
	}
	q := make([]int, 0)
	for i := range recipes {
		if indegree[i] == 0 {
			q = append(q, i)
		}
	}
	ans := make([]string, 0)
	for len(q) > 0 {
		recipeIdx := q[0]
		q = q[1:]
		recipe := recipes[recipeIdx]
		ans = append(ans, recipe)
		if _, ok := graph[recipe]; !ok {
			continue
		}
		for _, dependRecipe := range graph[recipe] {
			indegree[recipeToIdx[dependRecipe]]--
			if indegree[recipeToIdx[dependRecipe]] == 0 {
				q = append(q, recipeToIdx[dependRecipe])
			}
		}
	}
	return ans
}

func main() {
	println(findAllRecipes([]string{"bread", "sandwich", "burger"}, [][]string{{"yeast", "flour"}, {"bread", "meat"}, {"sandwich", "meat", "bread"}}, []string{"yeast", "flour", "meat"}))
}
