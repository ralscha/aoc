package mathx

func StringPermutations(arr []string) [][]string {
	if len(arr) == 0 {
		return [][]string{}
	}
	if len(arr) == 1 {
		return [][]string{arr}
	}
	var result [][]string
	for i, v := range arr {
		remaining := make([]string, 0, len(arr)-1)
		remaining = append(remaining, arr[:i]...)
		remaining = append(remaining, arr[i+1:]...)
		for _, p := range StringPermutations(remaining) {
			result = append(result, append([]string{v}, p...))
		}
	}
	return result
}
