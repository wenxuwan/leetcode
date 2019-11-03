/*
该题目的做法就是找到任务数最多的任务，然后能执行的时候优先执行。
*/
package main

import "fmt"

func leastInterval(tasks []byte, n int) int {
	taskNum := make(map[byte]int)
	task := make(map[byte]bool)
	var runningTask byte
	var maxTaskNum int
	result := 0
	for _, value := range tasks{
		taskNum[value]++
	}
	for key, _ := range taskNum {
		task[key] = true
	}
	for len(taskNum) != 0{
		for key, _ := range task {
			task[key] = true
		}
		for i := 0; i <= n; i++{
			maxTaskNum = -1
			runningTask = 'w'
			for key, value := range(taskNum){
				if value > maxTaskNum && task[key] == true{
					maxTaskNum = value
					runningTask = key
				}
			}
			if maxTaskNum != -1{
				taskNum[runningTask]--
				task[runningTask] = false
				if taskNum[runningTask] == 0{
					delete(taskNum, runningTask)
				}
				result++
				//fmt.Println("runing:", runningTask)
				if len(taskNum) == 0{
					break
				}
			}else{
				result++
				//fmt.Println("waiting:", runningTask)
			}
		}
		//fmt.Println(taskNum,task)
	}
	//fmt.Println(task)
	return result
}

func main(){
	tasks := []byte{'A','A','A','B','B','B'}
	n := 2
	fmt.Println(leastInterval(tasks, n))
}
