package main

import "fmt"

var months = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	var arriveMonthAlice, arriveDayAlice int
	var leaveMonthAlice, leaveDayAlice int
	var arriveMonthBob, arriveDayBob int
	var leaveMonthBob, leaveDayBob int
	for i := range 2 {
		arriveMonthAlice = arriveMonthAlice*10 + int(arriveAlice[i]-'0')
		arriveMonthBob = arriveMonthBob*10 + int(arriveBob[i]-'0')
		leaveMonthAlice = leaveMonthAlice*10 + int(leaveAlice[i]-'0')
		leaveMonthBob = leaveMonthBob*10 + int(leaveBob[i]-'0')
	}

	for i := 3; i < 5; i++ {
		arriveDayAlice = arriveDayAlice*10 + int(arriveAlice[i]-'0')
		arriveDayBob = arriveDayBob*10 + int(arriveBob[i]-'0')
		leaveDayAlice = leaveDayAlice*10 + int(leaveAlice[i]-'0')
		leaveDayBob = leaveDayBob*10 + int(leaveBob[i]-'0')
	}
	maxMonthArrive, minMonthLeave := 0, 0
	maxDayArrive, minDayLeave := 0, 0
	//  max(arriveMonthAlice, arriveMonthBob), min(leaveMonthAlice, leaveMonthBob)
	//  max(arriveDayAlice, arriveDayBob), min(leaveDayAlice, leaveDayBob)
	if arriveMonthAlice > arriveMonthBob {
		maxMonthArrive = arriveMonthAlice
		maxDayArrive = arriveDayAlice
	} else if arriveMonthAlice < arriveMonthBob {
		maxMonthArrive = arriveMonthBob
		maxDayArrive = arriveDayBob
	} else {
		maxMonthArrive = arriveMonthBob
		maxDayArrive = max(arriveDayBob, arriveDayAlice)
	}

	if leaveMonthAlice < leaveMonthBob {
		minMonthLeave = leaveMonthAlice
		minDayLeave = leaveDayAlice
	} else if leaveMonthAlice > leaveMonthBob {
		minMonthLeave = leaveMonthBob
		minDayLeave = leaveDayBob
	} else {
		minMonthLeave = leaveMonthBob
		minDayLeave = min(leaveDayBob, leaveDayAlice)
	}

	if minMonthLeave-maxMonthArrive >= 0 {
		days := 0
		for i := maxMonthArrive - 1; i < minMonthLeave-1; i++ {
			days += months[i]
		}
		diff := days - maxDayArrive + minDayLeave + 1
		if diff > 0 {
			return diff
		}
		return 0
	}

	return 0
}

func main() {
	fmt.Println(countDaysTogether("09-01", "10-19", "06-19", "10-20"))
	fmt.Println(countDaysTogether("10-20", "12-22", "06-21", "07-05"))
	fmt.Println(countDaysTogether("08-15", "08-18", "08-16", "08-19"))
}
