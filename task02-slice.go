package homework

//import "fmt"

func reverse(input []int64) (result []int64) {
	//Place your code here
	len_ := len(input) - 1
	с := []int64{}
	for i := 0; len_ >= 0; len_-- {
		с = append(с, input[len_])
		i += 1
	}
	return с
}

//
//func main() {
//	t2 := []int64{1, 2, 5, 15}
//	res := reverse(t2)
//	fmt.Println(res)
//}
