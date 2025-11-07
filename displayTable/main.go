package main

import (
	"fmt"
	"sort"
	"strconv"
)

func displayTable(orders [][]string) [][]string {
	foodTableMp := make(map[string]map[string]int)
	tableList := make([]string, 0)
	foodMp := make(map[string]bool, 0)
	foodList := make([]string, 0)

	for _, ord := range orders {
		table, food := ord[1], ord[2]
		if _, ok := foodMp[food]; !ok {
			foodList = append(foodList, food)
		}
		foodMp[food] = true
		val, ok := foodTableMp[table]
		if ok {
			val[food]++
		} else {
			foodTableMp[table] = map[string]int{}
			tableList = append(tableList, table)
			foodTableMp[table][food]++
		}
	}
	sort.Strings(foodList)
	sort.Slice(tableList, func(i, j int) bool {
		numI, _ := strconv.Atoi(tableList[i])
		numJ, _ := strconv.Atoi(tableList[j])
		return numI < numJ
	})
	clear(foodMp)
	ans := make([][]string, len(foodTableMp)+1)
	ans[0] = make([]string, len(foodList)+1)
	ans[0][0] = "Table"

	for i, food := range foodList {
		ans[0][i+1] = food
	}

	for i, table := range tableList {
		// plus one for table name
		ans[i+1] = make([]string, len(foodList)+1)
		ans[i+1][0] = table
		currTable := foodTableMp[table]
		for j, food := range foodList {
			ans[i+1][j+1] = strconv.Itoa(currTable[food])
		}
	}
	return ans
}

func main() {
	//[]
	fmt.Println(displayTable([][]string{{"David", "3", "Ceviche"}, {"Corina", "10", "Beef Burrito"}, {"David", "3", "Fried Chicken"}, {"Carla", "5", "Water"}, {"Carla", "5", "Ceviche"}, {"Rous", "3", "Ceviche"}}))
}
