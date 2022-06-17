package homework

import "sort"

//import (
//	"fmt"
//	"sort"
//)

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	res := []string{}
	keys := []int{}
	for k := range input {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for i := 0; i < len(keys); i++ {
		res = append(res, input[keys[i]])
	}
	return res
}

//func main() {
//	t3 := map[int]string{2: "a", 0: "b", 1: "c"}
//	res := sortMapValues(t3)
//	fmt.Println(res)
//}

//Input -> {2: "a", 0: "b", 1: "c"}
//Output -> ["b", "c", "a"]
//Input -> {10: "aa", 0: "bb", 500: "cc"}
//Output -> ["bb", "aa", "cc"]
