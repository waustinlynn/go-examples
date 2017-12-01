package main

import (
	"fmt"
)

func quicksort(a []int, p, r int){
	if(p < r){
		q := partition(a, p, r)
		fmt.Println("qs", p, r,q, a)
		quicksort(a, p, q)
		quicksort(a, q+1, r)
	}
}

func partition(a []int, p, r int) int {
	pivot := a[p]
	i := p
	j := r
	temp := 0
	for {
		for a[i] < pivot && a[i] != pivot {
			i++
		}
		for a[j] > pivot && a[j] != pivot {
			j--
		}
		if(i < j){
			temp = a[i]
			a[i] = a[j]
			a[j] = temp
		}else{
			return j
		}
		//fmt.Println(i, j, a)
	}
}

func main(){
	a := []int {63,52,25,37,14,17,8,6}
	quicksort(a, 0, len(a) -1)
}

