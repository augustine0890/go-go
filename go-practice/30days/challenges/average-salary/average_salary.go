package main

import "fmt"

func average(salary []int) float64 {
	minS, maxS, sumS := salary[0], salary[0], 0

	for _, s := range salary {
		sumS += s
		if s < minS {
			minS = s
		}
		if s > maxS {
			maxS = s
		}
	}
	avgS := float64(sumS-minS-maxS) / float64(len(salary)-2)
	return avgS
}

func main() {
	res1 := average([]int{4000, 3000, 1000, 2000})
	fmt.Printf("Result: %f\n", res1)
	res2 := average([]int{1000, 2000, 3000})
	fmt.Printf("Result: %f\n", res2)
	res3 := average([]int{6000, 5000, 4000, 3000, 2000, 1000})
	fmt.Printf("Result: %f\n", res3)
	res4 := average([]int{8000, 9000, 2000, 3000, 6000, 1000})
	fmt.Printf("Result: %f\n", res4)
}
