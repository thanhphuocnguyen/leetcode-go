package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func countMentions(numberOfUsers int, events [][]string) []int {
	// evt msg ["MESSAGE", "timestampi", "mentions_stringi"]
	// offline msg ["OFFLINE", "timestampi", "idi"]

	// sort events by the timestamps
	sort.Slice(events, func(i, j int) bool {
		if events[i][1] == events[j][1] {
			return events[i][0] == "OFFLINE"
		}
		ti, _ := strconv.Atoi(events[i][1])
		tj, _ := strconv.Atoi(events[j][1])
		return ti < tj
	})

	// create ans with cap eq events length
	mentions := make([]int, numberOfUsers)

	// maintain onl/off users
	onlUsers := make(map[int]bool)
	for i := range numberOfUsers {
		onlUsers[i] = true
	}

	offUsers := make(map[int]int)
	all := 0

	for _, evt := range events {
		tp, timeStr := evt[0], evt[1]
		time, _ := strconv.Atoi(timeStr)

		switch tp {
		case "MESSAGE":
			// release users
			for id, offtime := range offUsers {
				if time >= offtime+60 {
					onlUsers[id] = true
					delete(offUsers, id)
				}
			}
			switch evt[2] {
			case "ALL":
				{
					all++
					continue
				}
			case "HERE":
				for id := range onlUsers {
					mentions[id]++
				}
			default:
				idsStr := strings.Split(evt[2], " ")
				for _, idStr := range idsStr {
					id, _ := strconv.Atoi(idStr[2:])
					mentions[id]++
				}
			}
		case "OFFLINE":
			id, _ := strconv.Atoi(evt[2])
			delete(onlUsers, id)
			offUsers[id] = time
		}
	}

	for i := range mentions {
		mentions[i] += all
	}
	return mentions
}

func main() {
	fmt.Println(countMentions(15, [][]string{{"MESSAGE", "174", "HERE"}, {"OFFLINE", "426", "0"}, {"MESSAGE", "98", "ALL"}, {"MESSAGE", "383", "ALL"}, {"MESSAGE", "178", "id13 id1 id6 id0 id8 id6 id0"}, {"OFFLINE", "474", "10"}, {"MESSAGE", "190", "ALL"}, {"MESSAGE", "190", "HERE"}, {"MESSAGE", "366", "ALL"}, {"OFFLINE", "113", "4"}, {"MESSAGE", "130", "HERE"}, {"OFFLINE", "299", "10"}, {"OFFLINE", "352", "8"}, {"MESSAGE", "167", "id12 id13 id2 id10"}, {"MESSAGE", "120", "ALL"}, {"MESSAGE", "175", "ALL"}, {"OFFLINE", "150", "2"}, {"MESSAGE", "124", "ALL"}, {"OFFLINE", "70", "13"}}))
	fmt.Println(countMentions(3, [][]string{{"MESSAGE", "5", "HERE"}, {"OFFLINE", "10", "0"}, {"MESSAGE", "15", "HERE"}, {"OFFLINE", "18", "2"}, {"MESSAGE", "20", "HERE"}}))
	fmt.Println(countMentions(3, [][]string{{"MESSAGE", "2", "HERE"}, {"OFFLINE", "2", "1"}, {"OFFLINE", "1", "0"}, {"MESSAGE", "61", "HERE"}}))
	fmt.Println(countMentions(2, [][]string{{"OFFLINE", "10", "0"}, {"MESSAGE", "12", "HERE"}}))
	fmt.Println(countMentions(2, [][]string{{"MESSAGE", "10", "id1 id0"}, {"OFFLINE", "11", "0"}, {"MESSAGE", "12", "ALL"}}))
	fmt.Println(countMentions(2, [][]string{{"MESSAGE", "10", "id1 id0"}, {"OFFLINE", "11", "0"}, {"MESSAGE", "71", "HERE"}}))
}
