package main

import (
	"fmt"
	"sort"
)

func total(params []int) int {
	totals := 0
	for i := 0; i < len(params); i++ {

		temp := params[i]
		totals += temp
	}
	return totals
}

func average(params []int) float64 {
	n := len(params)
	sum := 0

	for i := 0; i < n; i++ {
		sum += (params[i])
	}

	avg := (float64(sum)) / (float64(n))
	return avg
}

func findMinAndMax(a []int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

func main() {
	// fmt.Println("OK")
	data := [...]int{1, 2, 33, 44, 55, 33, 44, 22, 44, 33, 2, 77, 66, 1, 2, 3, 4, 5, 6, 7, 89, 3, 3, 8, 9, 75, 4, 3, 2, 2, 1, 3}
	fmt.Println("Data :")
	fmt.Println(data)
	fmt.Printf("\n")
	// fmt.Println(len(data))
	fmt.Println("Bagi Data Tersebut menjadi 3 bagian :")
	length := len(data) / 3
	// fmt.Println(length)
	// length = math.Round(float64(length))
	// fmt.Println(length)
	newData1 := data[:int(length)]
	newData2 := data[int(length) : int(length)*2]
	newData3 := data[int(length)*2:]

	fmt.Println(newData1)
	fmt.Println(newData2)
	fmt.Println(newData3)
	fmt.Printf("\n")

	// Tampilkan data setiap kelompok terurut dari besar ke kecil
	fmt.Println("Tampilkan data setiap kelompok terurut dari besar ke kecil :")
	sort.Sort(sort.Reverse(sort.IntSlice(newData1)))
	sort.Sort(sort.Reverse(sort.IntSlice(newData2)))
	sort.Sort(sort.Reverse(sort.IntSlice(newData3)))
	fmt.Println(newData1)
	fmt.Println(newData2)
	fmt.Println(newData3)
	fmt.Printf("\n")

	// Tampilkan total setiap kelompok data
	fmt.Println("Tampilkan total setiap kelompok data :")
	fmt.Println(total(newData1))
	fmt.Println(total(newData2))
	fmt.Println(total(newData3))
	fmt.Printf("\n")

	// Tampilkan rata rata setiap kelompok data
	fmt.Println("Tampilkan rata rata setiap kelompok data :")
	fmt.Println(average(newData1))
	fmt.Println(average(newData2))
	fmt.Println(average(newData3))
	fmt.Printf("\n")

	// Carilah nilai tertinggi dan terendah dari setiap kelompok data
	fmt.Println("Carilah nilai tertinggi dan terendah dari setiap kelompok data :")
	min, max := findMinAndMax(newData1)
	fmt.Println("Data 1: ", newData1)
	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
	fmt.Printf("\n")

	min, max = findMinAndMax(newData2)
	fmt.Println("Data 2: ", newData2)
	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
	fmt.Printf("\n")

	min, max = findMinAndMax(newData3)
	fmt.Println("Data 3: ", newData3)
	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
}
