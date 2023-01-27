package main

import (
	"fmt"
	"sort"
	"strconv"
)

// Веруть int
func ReturnInt() int {
	return 1
}

// Вернуть float
func ReturnFloat() float32 {
	return 1.1
}

// Вернуть массив int
func ReturnIntArray() [3]int {
	return [3]int{1, 3, 4}
}

// Вернуть слайс int
func ReturnIntSlice() []int {
	return []int{1, 2, 3}
}

// Из Слайса int в string
func IntSliceToString(sl []int) string {
	res := ""
	for _, ch := range sl {
		res += strconv.Itoa(ch)
	}
	return res
}

// Объединить слайсы
func MergeSlices(slf []float32, sli []int) []int {
	res := make([]int, len(slf))
	for i, s := range slf {
		res[i] = int(s)
	}
	res = append(res, sli...)
	return res

}

// Вернуть отсортированные значения map по ключу
func GetMapValuesSortedByKey(mp map[int]string) []string {
	var keys []int
	var res []string
	for k := range mp {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	for _, k := range keys {
		res = append(res, mp[k])
	}
	return res
}

func main () {
	fmt.Println("Hello, world")
	fmt.Println("Return int: ", ReturnInt())
	fmt.Println("Return float : ", ReturnFloat())
	fmt.Println("Return int array : ", ReturnIntArray())
	fmt.Println("Return int slice : ", ReturnIntSlice())
	fmt.Println("Int slice to string : ", IntSliceToString([]int{1, 2, 3}))
	fmt.Println("Merge slices : ", MergeSlices([]float32{1.1, 2.2, 3.3}, []int{4 ,5}))
	fmt.Println("GET map values sorted by key : ", GetMapValuesSortedByKey(map[int]string{10: "Зима",30: "Лето", 20: "Весна", 40: "Осень"}))
}
