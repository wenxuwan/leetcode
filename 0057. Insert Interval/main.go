/*
56 题的变形，其实该题目最多也就算中等难度。
这里就用最简单的方法，将新的区间插入到原先
的切片里面，然后同样的方式来合并区间。
 */
package main

import "fmt"

type Interval struct {
 	   Start int
 	   End   int
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	intervalsResult := make([]Interval, 0, len(intervals))
	j := 0
	if len(intervals) == 0{
		intervalsResult = append(intervalsResult,newInterval)
		return intervalsResult
	}
	for j = 0; j < len(intervals); j++ {
		if (newInterval.Start <= intervals[j].Start) {
			tempinervals := append([]Interval{}, intervals[j:]...)
			//fmt.Print(tempinervals)
			intervals = append(intervals[:j], newInterval)
			intervals = append(intervals, tempinervals...)
			break
			//fmt.Print(intervals)
		}
	}
	if j == len(intervals){
		intervals = append(intervals,newInterval)
	}
	fmt.Print(intervals)
	for i := 0; i < len(intervals) - 1;i++{
		if intervals[i].End >= intervals[i+1].Start {
			intervals[i+1].Start = intervals[i].Start
			if intervals[i].End > intervals[i+1].End {
				intervals[i+1].End = intervals[i].End
			}
		} else {
			intervalsResult = append(intervalsResult, intervals[i])
		}
	}
	intervalsResult = append(intervalsResult, intervals[len(intervals) - 1])
	return intervalsResult
}

func main(){
	testData := []Interval{{1,5}}
	newData :=Interval{2,7}
	fmt.Print(insert(testData,newData))
}