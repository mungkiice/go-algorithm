package main

import "fmt"

func jobOffers(scores, lowerLimits, upperLimits []int) []int {
	var result []int
	for i := 0; i< len(lowerLimits); i++{
		var count int
		for _,score := range scores {
			if score >= lowerLimits[i] && score <= upperLimits[i] {
				count++
			} 
		}
		result = append(result, count)
	}
	return result
}

func main(){
	fmt.Println(jobOffers([]int{1,3,5,6,8}, []int{2,2}, []int{6,8}))
}