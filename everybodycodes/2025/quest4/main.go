package main

import (
	"aoc/internal/conv"
	"fmt"
	"math"
	"strings"
)

func main() {
	partI()
	partII()
	partIII()
}

func partI() {
	input := `1000
995
976
957
930
925
900
871
846
825
821
803
784
763
742
714
698
676
665
656
629
614
589
584
576
566
553
530
507
485
484
461
450
436
416
408
384
372
351
348
335
311
299
289
281
258
239
212
207
195`

	gears := conv.ToIntSlice(conv.SplitNewline(input))

	ratio := float64(1)
	l := len(gears) - 1
	for i := range l {
		ratio *= float64(gears[i]) / float64(gears[i+1])
	}

	result := int(ratio * 2025)
	fmt.Println(result)
}

func partII() {
	input := `966
957
953
944
923
910
894
867
852
833
819
806
782
765
745
727
716
711
691
686
662
639
633
614
601
580
578
577
561
557
543
533
507
491
485
474
467
455
433
429
406
401
378
361
354
332
330
314
305
277`

	gears := conv.ToIntSlice(conv.SplitNewline(input))

	firstTeeth := float64(gears[0])
	lastTeeth := float64(gears[len(gears)-1])

	result := int(math.Ceil(1e13 * lastTeeth / firstTeeth))
	fmt.Println(result)

}

func partIII() {
	input := `631
626|626
622|622
604|604
593|2372
581|581
563|1689
552|552
537|1611
532|532
524|1048
521|521
504|2016
496|496
490|1960
479|479
475|950
463|463
446|446
432|432
413|413
408|408
405|1215
389|389
378|1134
370|370
366|1098
350|350
347|1388
338|338
324|972
314|314
298|596
285|285
271|271
265|265
255|510
242|242
240|240
225|225
221|663
206|206
190|570
187|187
168|672
164|164
149|298
130|130
127|127
88
`

	lines := conv.SplitNewline(input)

	type Gear struct {
		left  int
		right int
	}

	var gears []Gear
	for _, line := range lines {
		if contains := false; !contains {
			splitted := strings.Split(line, "|")
			if len(splitted) == 2 {
				left := conv.MustAtoi(splitted[0])
				right := conv.MustAtoi(splitted[1])
				gears = append(gears, Gear{left: left, right: right})
			} else {
				teeth := conv.MustAtoi(line)
				gears = append(gears, Gear{left: teeth, right: teeth})
			}
		}
	}

	turns := float64(100)
	l := len(gears) - 1
	for i := range l {
		currentRight := float64(gears[i].right)
		nextLeft := float64(gears[i+1].left)

		turns = turns * currentRight / nextLeft
	}

	result := int(turns)
	fmt.Println(result)
}
