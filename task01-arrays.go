package homework

func average(input [15]float32) (result float32) {
	var total float32
	for _, element := range input {
		total += element
	}
	return total / float32(len(input))
}
