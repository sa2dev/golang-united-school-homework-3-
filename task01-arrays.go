package homework

func average(input [15]float32) (result float32) {
	var sum float32 = 0.0
	for _, v := range input {
		sum += v
	}
	return sum / 15.0
}
