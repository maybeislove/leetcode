package main

import "fmt"

func main() {
	fmt.Printf("%v", sortArrayByParityII([]int{4, 2, 5, 7}))
}

func sortArrayByParityII(nums []int) []int {
	//result := []int{}
	for i, num := range nums {
		odd := false
		even := false

		if i%2 == 0 {
			even = true
		} else {
			odd = true
		}

		if odd && num%2 != 0 {
			continue
			//result = append(result, num)
		} else if even && num%2 == 0 {
			continue
			//result = append(result, num)
		} else if odd {
			for j := i + 1; j < len(nums); j++ {
				if nums[j]%2 != 0 {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
		} else {
			for j := i + 1; j < len(nums); j++ {
				if nums[j]%2 == 0 {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
		}
	}
	return nums
}
