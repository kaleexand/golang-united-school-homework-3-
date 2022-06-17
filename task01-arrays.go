package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	if len(input) == 0 {
		return float32(0)
	}
	sum := float32(0)
	cup := 0
	for i := 0; i < len(input); i++ {
		if input[i] != 0 {
			cup++
		}
		sum += input[i]
	}
	return sum / float32(cup)
}

//func main() {
//	t1 := [15]float32{1, 2, 3, 4, 5, 6}
//	res := average(t1)
//	fmt.Printf("avg: %f\n", res)
//}
