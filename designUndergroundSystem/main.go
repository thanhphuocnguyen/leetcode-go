package main

import "fmt"

type CheckinInfo struct {
	station string
	time    int
}

type StationsStat struct {
	total int
	cnt   int
}

type UndergroundSystem struct {
	pending map[int]CheckinInfo
	table   map[string]map[string]*StationsStat
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		table:   map[string]map[string]*StationsStat{},
		pending: map[int]CheckinInfo{},
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.pending[id] = CheckinInfo{station: stationName, time: t}
	if _, ok := this.table[stationName]; !ok {
		this.table[stationName] = make(map[string]*StationsStat)
	}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	if _, ok := this.table[this.pending[id].station][stationName]; !ok {
		this.table[this.pending[id].station][stationName] = &StationsStat{}
	}
	this.table[this.pending[id].station][stationName].total += t - this.pending[id].time
	this.table[this.pending[id].station][stationName].cnt++
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	return float64(this.table[startStation][endStation].total) / float64(this.table[startStation][endStation].cnt)
}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */

func main() {
	system := Constructor()
	system.CheckIn(45, "Leyton", 3)
	system.CheckIn(32, "Paradise", 8)
	system.CheckIn(27, "Leyton", 10)
	system.CheckOut(45, "Waterloo", 15)
	system.CheckOut(27, "Waterloo", 20)
	system.CheckOut(32, "Cambridge", 22)
	fmt.Println(system.GetAverageTime("Paradise", "Cambridge"))
	fmt.Println(system.GetAverageTime("Leyton", "Waterloo"))
	system.CheckIn(10, "Leyton", 24)
	fmt.Println(system.GetAverageTime("Leyton", "Waterloo"))
	system.CheckOut(10, "Waterloo", 38)
	fmt.Println(system.GetAverageTime("Leyton", "Waterloo"))
}
