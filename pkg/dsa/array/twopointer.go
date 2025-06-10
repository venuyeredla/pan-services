package array

// Sorted array
func removeDuplicates(input []int) {
	i := 0
	for j := 1; j < len(input); j++ {
		if input[j] != input[i] {
			i = i + 1
			input[i] = input[j]
		}
	}
	i++
	for ; i < len(input); i++ {
		input[i] = -1
	}
}

// {1, 2, 0, 0, 0, 3, 6};
// Method-1 : whenever zero element move all next elements to before then place zero at the end.
// Method-2 : Counting zeros and moving next nonzero elemnt
func MovallZeros(a []int) {
	widx := -1
	for i := 0; i < len(a); i++ {
		if a[i] != 0 {
			widx += 1
			a[widx] = a[i]
		}
	}
	for widx += 1; widx < len(a); widx++ {
		a[widx] = 0
	}
}
