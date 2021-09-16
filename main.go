package main

import "fmt"


func KidsWithCandies(candies []int, extraCandies int)[]bool{

	max := 0
	extra := extraCandies
	results := make([]bool, len(candies))

	for _,value := range candies{
		if value > max{
			max = value
		}
	}

	for i,value := range candies{
		if value + extra >= max{
			results[i]=true
		}
	}

	return results

}

func main(){
	var candiesArray = []int{2,3,4,8,5,1}
	var extraCandies int = 3
	fmt.Println(KidsWithCandies(candiesArray,extraCandies))

}