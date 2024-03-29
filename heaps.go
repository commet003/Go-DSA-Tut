package main

import "fmt"

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Insert(key int){
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)	
}

func (h *MaxHeap) Extract() int{
	extracted := h.array[0]
	lastElement := len(h.array) - 1
	if len(h.array) == 0 {
		fmt.Println("Cannot exstract element because the array is empty.")
		return -1
	}
	h.array[0] = h.array[lastElement]
	h.array = h.array[:lastElement]

	h.maxHeapifyDown(0)
	return extracted
}

func (h *MaxHeap) maxHeapifyDown(index int){
	lastIndex := len(h.array) - 1
	l,r := left(index), right(index) 
	childToCompare := 0

	for l <= lastIndex{
		if l == lastIndex{
			childToCompare = l
		}else if h.array[l] > h.array[r]{
			childToCompare = l
		}else{
			childToCompare = r
		}
		
		if h.array[index] < h.array[childToCompare]{
			h.swap(index,childToCompare)
			index = childToCompare
			l,r = left(index), right(index)
		}else{
			return
		}
	}
}

func (h *MaxHeap) maxHeapifyUp(index int){
	for h.array[parent(index)] < h.array[index]{
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func parent(parentIndex int) int{
	return(parentIndex-1)/2
}

func left(leftChildIndex int) int{
	return 2*leftChildIndex + 1
}

func right(rightChildIndex int) int{
	return 2*rightChildIndex + 2
}

func (h *MaxHeap) swap(i1,i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}


func main(){
	m := &MaxHeap{}
	buildHeap := []int{10,20,30,15,7,2,50,41,23}

	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
