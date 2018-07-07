/*
合并区间问题，首先要想到的肯定是要把区间按照Start排序，
之所以排序的原因是因为我们可以一次性找到能不能合并的区间
eg:
	[1,5],[6,7],[5,8]
如果直接遍历,[1,5],[6,7]无法合并，但[1,5][5,8]是可以合并的。
合并之后[1,8]又可以和[6,7]合并。所以我们必须保证有序，才容易
直接判断某一个区间是不是已经不能和其它区间合并了。

排好序之后就是简单的合并区间，intervals[i].End >= intervals[i+1].Start
是合并的前提，不满足的话就代表这个区间已经没重合区间了和其它的。
*/

package main

import "fmt"

type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	intervalsResult := make([]Interval,0,len(intervals))
	if len(intervals) == 0{
		return  intervalsResult
	}
	for i :=0;i < len(intervals);i++{
		for j := 0; j < len(intervals) - i - 1; j++{
			if intervals[j].Start > intervals[j + 1].Start{
				start := intervals[j + 1].Start
				end := intervals[j + 1].End
				intervals[j + 1].Start =  intervals[j].Start
				intervals[j + 1].End =  intervals[j].End
				intervals[j].Start = start
				intervals[j].End =end
			}
		}
	}
	fmt.Println(intervals)
	for i :=0;i < len(intervals) - 1;i++{
		if intervals[i].End >= intervals[i+1].Start{
			intervals[i + 1].Start = intervals[i].Start
			if intervals[i].End > intervals[i + 1].End{
				intervals[i + 1].End = intervals[i].End
			}
		}else{
			intervalsResult = append(intervalsResult, intervals[i])
		}
	}
	intervalsResult = append(intervalsResult, intervals[len(intervals) - 1])
	return intervalsResult
}

func main(){
	testData := []Interval{}
	fmt.Print(merge(testData))
}
