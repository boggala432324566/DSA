package main

import (
	"fmt"
	"math"
	"strconv"
)

const MaxCount int = math.MaxInt64

func (q *Queue) SmallestNumberDividesDigits0ans9(num int64) {
	var list []string
	q.EnQueue("9")
	for i := 0; i < MaxCount; i++ {
		ele := q.Peek()
		q.DeQueue()
		list = append(list, ele)
		q.EnQueue(ele + "0") //If it is a int use q.EnQueue((ele*10)+0)
		q.EnQueue(ele + "9") //If it is a int use q.EnQueue((ele*10)+9)
	}
	for _, v := range list {
		intVal, _ := strconv.ParseInt(v, 10, 64)
		if intVal%num == 0 {
			fmt.Printf("Smallest number is:%d", intVal)
			break
		}
	}
}
