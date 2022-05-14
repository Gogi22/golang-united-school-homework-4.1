package reverse_string

func ReverseString(input string) (output string) {
	temp := make([]rune, 0)

	for _, v := range input {
		temp = append(temp, v)
	}

	for i := 0; i < len(temp)/2; i++ {
		temp[i], temp[len(temp)-i-1] = temp[len(temp)-i-1], temp[i]
	}

	output = string(temp)
	return
}
