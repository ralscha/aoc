package main

import (
	"aoc/internal/stringutil"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	partI()
	partII()
	partIII()
}

func partI() {
	input := `WORDS:LOR,LL,SI,OR,ID,UA,IM

LOREM IPSUM DOLOR SIT AMET, CONSECTETUR ADIPISCING ELIT, SED DO EIUSMOD TEMPOR INCIDIDUNT UT LABORE ET DOLORE MAGNA ALIQUA. UT ENIM AD MINIM VENIAM, QUIS NOSTRUD EXERCITATION ULLAMCO LABORIS NISI UT ALIQUIP EX EA COMMODO CONSEQUAT. DUIS AUTE IRURE DOLOR IN REPREHENDERIT IN VOLUPTATE VELIT ESSE CILLUM DOLORE EU FUGIAT NULLA PARIATUR. EXCEPTEUR SINT OCCAECAT CUPIDATAT NON PROIDENT, SUNT IN CULPA QUI OFFICIA DESERUNT MOLLIT ANIM ID EST LABORUM.
`
	lines := strings.Split(input, "\n")
	wordsLine := strings.Split(lines[0], ":")[1]
	words := strings.Split(wordsLine, ",")
	text := lines[2]

	wordCount := 0
	for _, word := range words {
		wordCount += strings.Count(text, word)
	}
	fmt.Println(wordCount)
}

func partII() {
	input, err := os.ReadFile("everybodycodes/quest2/partII.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")
	wordsLine := strings.Split(lines[0], ":")[1]
	words := strings.Split(wordsLine, ",")
	textLines := lines[2:]
	for i, l := range words {
		words[i] = strings.TrimRight(l, "\r\n")
	}

	wordCount := 0
	for _, l := range textLines {
		found := make([]bool, len(l))
		for _, word := range words {
			indices := stringutil.FindAllOccurrences(l, word)
			for _, i := range indices {
				for j := 0; j < len(word); j++ {
					found[i+j] = true
				}
			}
			reversed := stringutil.Reverse(word)
			indices = stringutil.FindAllOccurrences(l, reversed)
			for _, i := range indices {
				for j := 0; j < len(reversed); j++ {
					found[i+j] = true
				}
			}
		}
		for _, f := range found {
			if f {
				wordCount++
			}
		}
	}
	fmt.Println(wordCount)

}

func partIII() {
	input, err := os.ReadFile("everybodycodes/quest2/partIII.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")
	wordsLine := strings.Split(lines[0], ":")[1]
	words := strings.Split(wordsLine, ",")
	textLines := lines[2:]
	for i, l := range words {
		words[i] = strings.TrimRight(l, "\r\n")
	}
	for i, l := range textLines {
		textLines[i] = strings.TrimRight(l, "\r\n")
	}

	wordCount := 0

	found := make([][]bool, len(textLines))
	for i := range found {
		found[i] = make([]bool, len(textLines[i]))
	}
	for row, line := range textLines {
		for rowa := range len(line) {
			newLine := line[rowa:]
			newLine += line[:rowa]
			for _, word := range words {
				indices := stringutil.FindAllOccurrences(newLine, word)
				for _, i := range indices {
					for j := 0; j < len(word); j++ {
						found[row][(i+j+rowa)%len(newLine)] = true
					}
				}
				reversed := stringutil.Reverse(word)
				indices = stringutil.FindAllOccurrences(newLine, reversed)
				for _, i := range indices {
					for j := 0; j < len(reversed); j++ {
						found[row][(i+j+rowa)%len(newLine)] = true
					}
				}
			}
		}
	}

	rotated := make([]string, len(textLines[0]))
	for i := range rotated {
		rotated[i] = ""
	}
	for _, l := range textLines {
		for i, c := range l {
			rotated[i] += string(c)
		}
	}

	for col, l := range rotated {
		for _, word := range words {
			indices := stringutil.FindAllOccurrences(l, word)
			for _, i := range indices {
				for j := 0; j < len(word); j++ {
					found[i+j][col] = true
				}
			}

			reversed := stringutil.Reverse(word)
			indices = stringutil.FindAllOccurrences(l, reversed)
			for _, i := range indices {
				for j := 0; j < len(reversed); j++ {
					found[i+j][col] = true
				}
			}
		}
	}

	for _, f := range found {
		for _, f2 := range f {
			if f2 {
				wordCount++
			}
		}
	}

	fmt.Println(wordCount)
}
