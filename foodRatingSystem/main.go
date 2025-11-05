package main

import (
	"container/heap"
)

type Item struct {
	Name   string
	Rating int
}
type PriorityQueue []*Item
type FoodRatings struct {
	FoodCuisine map[string]string
	FoodCRating map[string]int
	CuisineFood map[string]*PriorityQueue
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	if pq[i].Rating != pq[j].Rating {
		return pq[i].Rating > pq[j].Rating
	} else {
		return pq[i].Name < pq[j].Name
	}
}

func (pq PriorityQueue) Top() *Item {
	if pq.Len() > 0 {
		return pq[0]
	}
	return nil
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // don't stop the GC from reclaiming the item eventually
	*pq = old[0 : n-1]
	return item
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	ctor := FoodRatings{
		FoodCuisine: map[string]string{},
		FoodCRating: map[string]int{},
		CuisineFood: map[string]*PriorityQueue{},
	}
	for i, food := range foods {
		ctor.FoodCRating[food] = ratings[i]
		ctor.FoodCuisine[food] = cuisines[i]
	}

	for i, cuisine := range cuisines {
		var pq PriorityQueue
		if _, ok := ctor.CuisineFood[cuisine]; !ok {
			pq = make(PriorityQueue, 0)
			ctor.CuisineFood[cuisine] = &pq
		}
		heap.Push(ctor.CuisineFood[cuisine], &Item{
			Name:   foods[i],
			Rating: ratings[i],
		})

	}
	return ctor
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	this.FoodCRating[food] = newRating
	heap.Push(this.CuisineFood[this.FoodCuisine[food]], &Item{
		Name:   food,
		Rating: newRating,
	})

}

func (this *FoodRatings) HighestRated(cuisine string) string {
	top := *this.CuisineFood[cuisine].Top()
	for this.FoodCRating[top.Name] != this.CuisineFood[cuisine].Top().Rating {
		heap.Pop(this.CuisineFood[cuisine])
		top = *this.CuisineFood[cuisine].Top()
	}
	return top.Name
}

/**
 * Your FoodRatings object will be instantiated and called as such:
 * obj := Constructor(foods, cuisines, ratings);
 * obj.ChangeRating(food,newRating);
 * param_2 := obj.HighestRated(cuisine);
 */

func main() {
	foodRatingSystem := Constructor(
		[]string{"wayfnkzb", "eqiyfjefmx", "pfe", "atlqeka", "iqtf", "akbtinchc", "pchazissxl",
			"qfehnltip", "vyavqiaha", "gpkieuuj", "cmmevexdu", "vy", "skrqgtqcol", "xsydpmavlp", "cprtdfwmcl", "iqmcqgb",
			"mzwtzlvjud", "dqgys", "emdmcsv", "ebrckogq", "txnsrgoz", "bogmosr", "usmxuzwwwg", "kxwhpirwqd", "nmuukr",
			"wnhajrqpux", "stczw", "sywjkh", "hmodhglhuy", "fbrefccega", "sgcggzyjbe", "dcfzm", "inuzyvkre", "uvbeprr",
			"nbbgq", "yohckl", "mulurwuvrd", "dwlqqoxz", "ejzec", "rd", "rgaxj", "wsjswolesx", "dmddhdnhm", "woguqmkc", "hhmkj",
			"dcveeeor", "holkfujdgy", "lg", "uxbmcubn", "qsdouuk", "orqn", "uyyksfql", "gxdhpvy", "bhwxoepw", "ynmtwtb", "cfcqsiitm",
			"ezwsqzcyy", "toafgd", "izqhhhkdwe", "itcxa", "senrcaz", "uvnogrjcr", "qeeclian", "zslurjd", "bzlvx", "aibtvlpryg", "oyqxgjyrnh",
			"lmmwtkgyyo", "mynkq", "rhxfybuakp", "vvadaflbt", "opcpi", "ytqmebeo", "qevszdh", "itkvzv", "qqtjye", "aqjwrbyr",
			"potiiecni", "smqifbhwzm", "hyhza", "cvfi"},
		[]string{"ikzdtynn", "tm", "tvmsdpsup", "tvmsdpsup", "ikzdtynn", "tm", "tm", "jbatpoby", "vrchucermi", "ylfyymwb", "tvmsdpsup",
			"ikzdtynn", "jbatpoby", "ikzdtynn", "dm", "jbatpoby", "vrchucermi", "tvmsdpsup", "tm", "dm", "dm", "tm", "jbatpoby", "ylfyymwb",
			"ylfyymwb", "vrchucermi", "jbatpoby", "vrchucermi", "tvmsdpsup", "ikzdtynn", "tvmsdpsup", "ikzdtynn", "tm", "dm", "jbatpoby",
			"jbatpoby", "vrchucermi", "ylfyymwb", "tm", "tvmsdpsup", "vrchucermi", "ikzdtynn", "ylfyymwb", "ylfyymwb", "ylfyymwb", "tm",
			"tvmsdpsup", "tm", "tvmsdpsup", "dm", "dm", "ylfyymwb", "dm", "ylfyymwb", "jbatpoby", "tvmsdpsup", "ikzdtynn", "ikzdtynn",
			"ylfyymwb", "dm", "tm", "vrchucermi", "dm", "jbatpoby", "ylfyymwb", "tvmsdpsup", "dm", "ylfyymwb", "jbatpoby", "dm", "tvmsdpsup",
			"ikzdtynn", "tm", "ikzdtynn", "ylfyymwb", "ikzdtynn", "tm", "tm", "dm", "ikzdtynn", "ylfyymwb"},
		[]int{577, 115, 150, 549, 727, 490, 585, 511, 700, 724, 641, 539, 282, 439, 229, 97, 524, 511, 27, 713, 359, 232, 895,
			769, 326, 186, 342, 913, 915, 523, 360, 903, 310, 355, 383, 173, 129, 6, 579, 7, 49, 168, 711, 925, 524,
			882, 393, 848, 165, 652, 692, 39, 370, 252, 472, 472, 875, 844, 562, 436, 901, 329, 597, 645, 208, 444, 219,
			873, 924, 597, 210, 774, 805, 331, 919, 548, 145, 162, 537, 818, 848},
	)
	foodRatingSystem.HighestRated("korean")
	foodRatingSystem.HighestRated("japanese")
	foodRatingSystem.ChangeRating("sushi", 16)
	foodRatingSystem.HighestRated("japanese")
	foodRatingSystem.ChangeRating("ramen", 16)
	foodRatingSystem.HighestRated("japanese")
}
