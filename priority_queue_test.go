package gods

import (
	"container/heap"
	"fmt"
	"strconv"
	"testing"
)

type TestStruct struct {
	A int64
	B float64
	C string
}

func TestPriorityQueue(t *testing.T) {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	var tao [10]TestStruct

	for i := 0; i < 10; i++ {
		tao[i].A = int64(1 + i*10)
		tao[i].B = float64(i) * float64(1.1)
		tao[i].C = "hello " + strconv.FormatInt(int64(i), 10)
		item := &Item{
			Value:    &tao[i],
			Priority: 1000 - tao[i].A,
		}
		heap.Push(&pq, item)
		pq.Update(item, item.Value, item.Priority)
	}

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fh := item.Value.(*TestStruct)
		fmt.Println(strconv.FormatInt(item.Priority, 10) + ":" + fh.C)
	}

}
