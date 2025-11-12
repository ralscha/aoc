package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"fmt"
	"strings"
)

func main() {
	partI()
	partII()
	partIII()
}

func partI() {
	input := `Aznarel,Aznyn,Vaelardith,Thymnyn,Thymnarel,Vaelnarel,Azardith,Thymardith,Vaelnyn

a > r,b
V > a
m > n,a
i > t
l > n,a
y > n,m
r > d,e
z > a,b
A > z
d > i
n > a,y
t > h
T > h
e > l
h > b`

	names, rules := parseInput(input)

	for _, name := range names {
		if isValidName(name, rules) {
			fmt.Println(name)
			return
		}
	}
}

func isValidName(name string, rules map[rune]*container.Set[rune]) bool {
	runes := []rune(name)
	for i := 0; i < len(runes)-1; i++ {
		from := runes[i]
		to := runes[i+1]

		if rules[from] == nil || !rules[from].Contains(to) {
			return false
		}
	}
	return true
}

func parseInput(input string) ([]string, map[rune]*container.Set[rune]) {
	lines := conv.SplitNewline(input)
	names := strings.Split(lines[0], ",")

	rules := make(map[rune]*container.Set[rune])
	for i := 2; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}
		parts := strings.Split(line, " > ")
		from := rune(parts[0][0])
		tos := strings.Split(parts[1], ",")

		if rules[from] == nil {
			rules[from] = container.NewSet[rune]()
		}
		for _, to := range tos {
			rules[from].Add(rune(to[0]))
		}
	}

	return names, rules
}

func partII() {
	input := `Phyrris,Gorathisis,Gorathonar,Azmarris,Nyrixbel,Tirris,Arakisis,Azmaririn,Gorathlar,Nyrixirin,Arakonar,Azmarzor,Elarris,Nythzor,Brynvash,Elararis,Tiraris,Arakbel,Azmarbel,Elaronar,Shaelonar,Tirphor,Tirbel,Tirisis,Nyrixaris,Phyronar,Shaellar,Arakvash,Nythisis,Nyrixonar,Phyrzor,Nythris,Tironar,Brynonar,Brynzor,Gorathzor,Shaelisis,Nytharis,Tiririn,Nyrixzor,Brynlar,Brynisis,Elarisis,Brynirin,Arakris,Araklar,Azmarphor,Nythbel,Elarbel,Shaelphor,Gorathphor,Goratharis,Nyrixris,Tirvash,Azmarlar,Arakaris,Nyrixphor,Elaririn,Phyrlar,Nythonar,Nythlar,Azmaronar,Nythirin,Gorathris,Brynaris,Brynphor,Nyrixvash,Phyrbel,Shaelirin,Tirzor,Brynris,Arakzor,Gorathvash,Brynbel,Phyrisis,Elarvash,Arakirin,Nythvash,Gorathirin,Shaelris,Shaelaris,Nyrixisis,Elarlar,Nythphor,Tirlar,Nyrixlar,Arakphor,Phyraris,Azmarisis,Phyririn,Shaelbel,Phyrphor,Shaelzor,Azmarvash,Elarzor,Phyrvash,Azmararis,Shaelvash,Elarphor,Gorathbel

	i > s,r,n,v,x
	s > i,h
	p > h
	o > r,n,v
	b > e
	x > b,z,i,r,o,l,a,p,v
	B > r
	z > v,o
	a > k,r,s,e,t
	v > a
	r > a,b,z,i,r,o,l,p,v
	h > o,v,b,z,i,r,l,a,p
	t > h
	y > r,n,v
	A > r,z
	T > i
	S > h
	E > l
	n > a,b,z,i,r,o,l,p,v
	N > y
	m > a
	e > l
	P > h
	G > o
	l > a,v,b,z,i,r,o,l,p
	k > i,b,z,r,o,l,a,p,v`

	names, rules := parseInput(input)

	sum := 0
	for idx, name := range names {
		if isValidName(name, rules) {
			sum += idx + 1
		}
	}
	fmt.Println(sum)
}

func partIII() {
	input := `Ny,Nyl,Nyth,Nyss,Nyrix,Xil,Sil,Kron,Tharil,Norak,Mal,Orah,Quor,Harn,Urak,Selk,Rael,Ild,Pylar,Ynd

r > o,i,e,y,v,a,n,z,p,t,q
o > s,r,n,v
p > y
u > o,v
k > i,n,z,p,t,q,e
Y > n
i > d,s,x,v,l
z > y,r
l > i,n,z,p,t,q,e,k,v,a
n > d,a,i,n,z,p,t,q,e,v
S > i,e
K > r
t > h
d > p,r,i,n,z,t,q,e
P > y
a > r,k,v,h
y > r,v,t,n
M > a
e > l,n,v
N > y,o
q > u
X > i
U > r
h > y,i,n,z,p,t,q,e,v
T > h
s > s,i,n,z,p,t,q,e
x > i,n,z,p,t,q,e
H > a
O > r
Q > u
R > a
I > l`

	prefixes, rules := parseInput(input)

	uniqueNames := container.NewSet[string]()

	for _, prefix := range prefixes {
		if !isValidName(prefix, rules) {
			continue
		}
		generateNames(prefix, rules, uniqueNames, 7, 11)
	}

	fmt.Println(uniqueNames.Len())
}

func generateNames(current string, rules map[rune]*container.Set[rune], uniqueNames *container.Set[string], minLen, maxLen int) {
	if len(current) >= minLen {
		uniqueNames.Add(current)
	}

	if len(current) >= maxLen {
		return
	}

	lastChar := rune(current[len(current)-1])

	if nextChars, ok := rules[lastChar]; ok {
		for _, nextChar := range nextChars.Values() {
			generateNames(current+string(nextChar), rules, uniqueNames, minLen, maxLen)
		}
	}
}
