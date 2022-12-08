package conv

import "strings"

func SplitNewline(s string) []string {
	splitted := strings.Split(s, "\n")
	if len(splitted) > 0 && splitted[len(splitted)-1] == "" {
		splitted = splitted[:len(splitted)-1]
	}
	return splitted
}
