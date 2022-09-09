package main

import (
	"fmt"
)

func remove1(slice []int) []int {
	res := []int{}

	for _, v := range slice {
		if v != 0 {
			res = append(res, v)
		}
	}

	return res
}

func remove2(slice []int) []int {
	count := 0
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] == 0 {
			slice = append(slice[:i], slice[i+1:]...)
			i--
		}
		count++
		if count == len(slice) {
			break
		}
	}

	return slice
}

func monotone(slice []int) bool {
	prev := slice[0]

	//Предполагаем что возрастает || равен
	if slice[0] <= slice[len(slice)-1] {
		for _, v := range slice {
			if v < prev {
				return false
			}
			prev = v
		}
	} else { //Убывает
		for _, v := range slice {
			if v > prev {
				return false
			}
			prev = v
		}
	}

	return true
}

func monotone2(slice []int) bool {
	up, down := true, true
	for i := 0; i < len(slice)-1; i++ {
		up, down = up && (slice[i] >= slice[i+1]), down && (slice[i] < slice[i+1])
		if !up && !down{
			return false
		}
	}
	return true
}

func main() {

	data := []int{10, 2, 2}
	fmt.Println(monotone2(data))
}
