package main

import "fmt"

func collision(speed []int, pos int) int{
	var count int
	// for i:= 0; i<pos;i++{
	// 	if speed[i] > speed[pos]{
	// 		count++
	// 	}
	// }
	// for i:= len(speed)-1;i>pos;i--{
	// 	if speed[i] < speed[pos]{
	// 		count++
	// 	}
	// }

	for i,particle := range speed{
		if i > pos && particle < speed[pos]{
			count++
		}else if i < pos && particle > speed[pos]{
			count++
		}
	}
	return count
}

func main(){
	fmt.Println(collision([]int{6,6,1,6,3,4,6,8}, 2))
}