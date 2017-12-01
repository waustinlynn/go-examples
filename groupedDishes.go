package main

import (
	"fmt"
	"sort"
)
func groupingDishes(dishes [][]string) [][]string {
	groupedDishes := make(map[string][]string)
	for _, arr := range dishes {
		for i := 1; i < len(arr); i++ {
			groupedDishes[arr[i]] = append(groupedDishes[arr[i]], arr[0])
			// if _, ok := groupedDishes[arr[i]]; ok {
				
			// }else{
			// 	groupedDishes[arr[i]] = append
			// }
		}
	}
	finalSlice := [][]string{}
	//fmt.Println(groupedDishes)
	for key, item := range groupedDishes {
		if(len(item) < 2) {
			continue
		}
		sort.Slice(item, func(i,j int) bool {return item[i] < item[j]})
		newArr := []string{key}
		newArr = append(newArr, item...)
		//fmt.Println(key, item)
		finalSlice = append(finalSlice, newArr)
	}
	sort.Slice(finalSlice, func(i,j int) bool {
		return len(finalSlice[i]) < len(finalSlice[j])
		//return finalSlice[i] < finalSlice[j]
	})
	return finalSlice	
}
	
func main() {
	dishes := [][]string{}
	first := []string {"Salad", "Tomato", "Cucumber", "Salad", "Sauce"}
	dishes = append(dishes, first)
	first = []string {"Pizza", "Tomato", "Sausage", "Sauce", "Dough"}
	dishes = append(dishes, first)
	first = []string {"Quesadilla", "Chicken", "Cheese", "Sauce"}
	dishes = append(dishes, first)
	first = []string{"Sandwich", "Salad", "Bread", "Tomato", "Cheese"}
	dishes = append(dishes, first)
	
	fmt.Println(groupingDishes(dishes))
}