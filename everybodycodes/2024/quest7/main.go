package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type chariot struct {
	name    string
	essence int
}

func simulateRace(name string, actionStr string, segments int) chariot {
	actions := []rune(strings.ReplaceAll(actionStr, ",", ""))
	power := 10
	essence := 0

	for i := range segments {
		action := actions[i%len(actions)]

		switch action {
		case '+':
			power++
		case '-':
			if power > 0 {
				power--
			}
		}

		essence += power
	}

	return chariot{
		name:    name,
		essence: essence,
	}
}

func simulateRaceWithTrack(name string, actionStr string, loops int, track []rune) chariot {
	actions := []rune(strings.ReplaceAll(actionStr, ",", ""))
	power := 10
	essence := 0
	trackPos := 0
	currentLoop := 0
	pos := 0
	for currentLoop < loops {
		action := actions[pos%len(actions)]
		pos++
		trackPos = (trackPos + 1) % len(track)

		if track[trackPos] == '-' {
			if power > 0 {
				power--
			}
		} else if track[trackPos] == '+' {
			power++
		} else {
			if track[trackPos] == 'S' {
				currentLoop++
			}
			switch action {
			case '+':
				power++
			case '-':
				if power > 0 {
					power--
				}
			}
		}
		essence += power
	}

	return chariot{
		name:    name,
		essence: essence,
	}
}

func partI() {
	input, err := os.ReadFile("everybodycodes/quest7/partI.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(input), "\n")
	for i, l := range lines {
		lines[i] = strings.TrimRight(l, "\r\n")
	}
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	chariots := make([]chariot, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}

		chariots[i] = simulateRace(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]), 10)
	}

	slices.SortFunc(chariots, func(a, b chariot) int {
		return b.essence - a.essence
	})

	var ranking strings.Builder
	for _, p := range chariots {
		ranking.WriteString(p.name)
	}
	fmt.Println("Part I", ranking.String())
}

func convertTrack(track string) []rune {
	lines := strings.Split(track, "\n")
	runePattern := make([][]rune, len(lines))
	for i, line := range lines {
		runePattern[i] = []rune(line)
	}

	rows := len(runePattern)
	cols := len(runePattern[0])
	result := make([]rune, 0)

	top := 0
	bottom := rows - 1
	left := 0
	right := cols - 1

	for i := left; i <= right; i++ {
		result = append(result, runePattern[top][i])
	}
	top++

	for i := top; i <= bottom; i++ {
		result = append(result, runePattern[i][right])
	}
	right--

	for i := right; i >= left; i-- {
		result = append(result, runePattern[bottom][i])
	}
	bottom--

	for i := bottom; i >= top; i-- {
		result = append(result, runePattern[i][left])
	}

	return result
}

func partII() {
	input, err := os.ReadFile("everybodycodes/quest7/partII.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(input), "\n")
	for i, l := range lines {
		lines[i] = strings.TrimRight(l, "\r\n")
	}
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	track := `S-=++=-==++=++=-=+=-=+=+=--=-=++=-==++=-+=-=+=-=+=+=++=-+==++=++=-=-=--
-                                                                     -
=                                                                     =
+                                                                     +
=                                                                     +
+                                                                     =
=                                                                     =
-                                                                     -
--==++++==+=+++-=+=-=+=-+-=+-=+-=+=-=+=--=+++=++=+++==++==--=+=++==+++-`
	trackRunes := convertTrack(track)

	chariots := make([]chariot, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}

		chariots[i] = simulateRaceWithTrack(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]), 10, trackRunes)
	}

	slices.SortFunc(chariots, func(a, b chariot) int {
		return b.essence - a.essence
	})

	var ranking strings.Builder
	for _, p := range chariots {
		ranking.WriteString(p.name)
	}
	fmt.Println("Part II", ranking.String())
}

type direction int

const (
	right direction = iota
	down
	left
	up
)

func partIII() {
	track := `S+= +=-== +=++=     =+=+=--=    =-= ++=     +=-  =+=++=-+==+ =++=-=-=--
- + +   + =   =     =      =   == = - -     - =  =         =-=        -
= + + +-- =-= ==-==-= --++ +  == == = +     - =  =    ==++=    =++=-=++
+ + + =     +         =  + + == == ++ =     = =  ==   =   = =++=
= = + + +== +==     =++ == =+=  =  +  +==-=++ =   =++ --= + =
+ ==- = + =   = =+= =   =       ++--          +     =   = = =--= ==++==
=     ==- ==+-- = = = ++= +=--      ==+ ==--= +--+=-= ==- ==   =+=    =
-               = = = =   +  +  ==+ = = +   =        ++    =          -
-               = + + =   +  -  = + = = +   =        +     =          -
--==++++==+=+++-= =-= =-+-=  =+-= =-= =--   +=++=+++==     -=+=++==+++-`

	trackLines := strings.Split(track, "\n")
	trackRunesGrid := make([][]rune, len(trackLines))
	for i, line := range trackLines {
		trackRunesGrid[i] = []rune(line)
	}
	maxLen := 0
	for _, line := range trackRunesGrid {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}
	for i, line := range trackRunesGrid {
		if len(line) < maxLen {
			fill := make([]rune, maxLen-len(line))
			for j := range fill {
				fill[j] = ' '
			}
			trackRunesGrid[i] = append(line, fill...)
		}
	}

	trackRunes := make([]rune, 0)
	trackRunes = append(trackRunes, trackRunesGrid[0][0])
	width := len(trackRunesGrid[0])
	height := len(trackRunesGrid)
	row := 0
	col := 1
	dir := right
	for trackRunesGrid[row][col] != 'S' {
		trackRunes = append(trackRunes, trackRunesGrid[row][col])
		switch dir {
		case right:
			if col+1 < width && trackRunesGrid[row][col+1] != ' ' {
				col++
			} else {
				if row+1 < height && trackRunesGrid[row+1][col] != ' ' {
					dir = down
					row++
				} else {
					dir = up
					row--
				}
			}
		case down:
			if row+1 < height && trackRunesGrid[row+1][col] != ' ' {
				row++
			} else {
				if col-1 >= 0 && trackRunesGrid[row][col-1] != ' ' {
					dir = left
					col--
				} else {
					dir = right
					col++
				}
			}
		case left:
			if col-1 >= 0 && trackRunesGrid[row][col-1] != ' ' {
				col--
			} else {
				if row-1 >= 0 && trackRunesGrid[row-1][col] != ' ' {
					dir = up
					row--
				} else {
					dir = down
					row++
				}
			}
		case up:
			if row-1 >= 0 && trackRunesGrid[row-1][col] != ' ' {
				row--
			} else {
				if col+1 < width && trackRunesGrid[row][col+1] != ' ' {
					dir = right
					col++
				} else {
					dir = left
					col--
				}
			}
		}
	}

	rivalName := "A"
	rivalActions := "=,-,-,+,+,=,+,+,=,-,+"
	rivalResult := simulateRaceWithTrack(rivalName, rivalActions, 2024, trackRunes)
	winningPlans := 0

	const totalLength = 11
	combinations := make([]string, 0)
	var generate func(current []rune, plus, minus, equals int)
	generate = func(current []rune, plus, minus, equals int) {
		if len(current) == totalLength {
			if plus == 0 && minus == 0 && equals == 0 {
				combinations = append(combinations, string(current))
			}
			return
		}

		if plus > 0 {
			generate(append(current, '+'), plus-1, minus, equals)
		}

		if minus > 0 {
			generate(append(current, '-'), plus, minus-1, equals)
		}

		if equals > 0 {
			generate(append(current, '='), plus, minus, equals-1)
		}
	}

	generate(make([]rune, 0), 5, 3, 3)

	for _, combination := range combinations {
		result := simulateRaceWithTrack("B", combination, 2024, trackRunes)
		if result.essence > rivalResult.essence {
			winningPlans++
		}
	}

	fmt.Println("Part III", winningPlans)
}

func main() {
	partI()
	partII()
	partIII()
}
