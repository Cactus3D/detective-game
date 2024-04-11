package locations

import (
	"container/heap"
	"testing"
)

func TestLen(t *testing.T) {
	var testsTable = []struct {
		input    priorityQueue
		expected int
	}{
		{
			priorityQueue{},
			0,
		},
		{
			priorityQueue{&queueItem{}, &queueItem{}},
			2,
		},
	}

	for _, test := range testsTable {
		got := test.input.Len()
		if got != test.expected {
			t.Errorf("Len = %d; extected %d", got, test.expected)
		}
	}
}

func TestLess(t *testing.T) {
	pq := priorityQueue{&queueItem{vertex: &Vertex{}, distance: 10}, &queueItem{vertex: &Vertex{}, distance: 5}}
	isLess := pq.Less(0, 1)
	if isLess {
		t.Errorf("Distance %d less %d are %t", pq[0].distance, pq[1].distance, isLess)
	}
	isLess = pq.Less(1, 0)
	if !isLess {
		t.Errorf("Distance %d less %d are %t", pq[0].distance, pq[1].distance, isLess)
	}
}

func TestSwap(t *testing.T) {
	i1, i2 := &queueItem{}, &queueItem{}
	pq := priorityQueue{i1, i2}
	pq.Swap(0, 1)
	if pq[0] != i2 {
		t.Errorf("Not swap has done")
	}
}

func TestPush(t *testing.T) {
	pq := priorityQueue{}
	qi := &queueItem{index: 10}

	heap.Init(&pq)
	heap.Push(&pq, qi)
	if pq.Len() != 1 {
		t.Errorf("Len of array = %d, expected 1", pq.Len())
	} else if pq[0].index != 0 {
		t.Errorf("Index of first element = %d, expected 0", pq[0].index)
	}
}

func TestPop(t *testing.T) {
	pq := priorityQueue{&queueItem{}, &queueItem{}, &queueItem{}}
	heap.Init(&pq)
	qi := heap.Pop(&pq).(*queueItem)

	if pq.Len() != 2 {
		t.Errorf("Len of array = %d, expected 2", pq.Len())
	}
	if qi.index != -1 {
		t.Errorf("Len of array = %d, expected 2", pq.Len())
	}
}

func TestUpdate(t *testing.T) {
	pq := priorityQueue{}
	heap.Init(&pq)
	for i := 1; i < 4; i++ {
		pq.Push(&queueItem{
			distance: i * 10,
		})
	}

	qi := pq[2]

	pq.update(qi, &Vertex{}, 0)

	gotDst := qi.distance
	gotIdx := qi.index
	gotVrtx := qi.vertex

	if gotDst != 0 {
		t.Errorf("Destination = %d, expected 0", gotDst)
	} else if gotIdx != 0 {
		t.Errorf("Index = %d, expected 0", gotIdx)
	} else if gotVrtx == nil {
		t.Errorf("Vertex must not be nil")
	}

}
