package main

import "fmt"

func angleClock(hour int, minutes int) float64 {
	anglePerMin := 6.0
	anglerPerHour := 30.0
	minuteAngle := float64(minutes) * anglePerMin
	hourAngleOffset := (float64(minutes) / 60.0) * anglerPerHour
	hourAngle := (float64(hour) * anglerPerHour) + hourAngleOffset
	diff := abs(hourAngle - minuteAngle)
	return min(360-diff, diff)
}

func abs(a float64) float64 {
	if a < 0.0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(angleClock(3, 30))
}
