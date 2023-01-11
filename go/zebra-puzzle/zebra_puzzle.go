package zebra

type Solution struct {
	DrinksWater string
	OwnsZebra   string
}

const (
	red = iota
	green
	ivory
	yellow
	blue
)

const (
	English = iota
	Spanish
	Ukrainian
	Norwegian
	Japanese
)

var nationality = map[int]string{
	English:   "Englishman",
	Spanish:   "Spaniard",
	Ukrainian: "Ukrainian",
	Norwegian: "Norwegian",
	Japanese:  "Japanese",
}

const (
	dog = iota
	snails
	fox
	horse
	zebra
)

const (
	coffee = iota
	tea
	milk
	orange_juice
	water
)

const (
	Old_Gold = iota
	Kools
	Chesterfields
	Lucky_Strikes
	Parliaments
)

type House struct {
	number, color, resident, pet, drink, smokes int
}

type Street []House

func SolvePuzzle() (rez Solution) {
	combinations := permutations(5)
	for _, nationalitySet := range combinations {
		//#1
		if nationalitySet[0] != Norwegian {
			continue
		}
		for _, smokesSet := range combinations {
			//#14
			if !match(nationalitySet, Japanese, smokesSet, Parliaments) {
				continue
			}
			for _, colorsSet := range combinations {
				//#2 #8 #6 #15
				if !match(colorsSet, red, nationalitySet, English) ||
					!match(smokesSet, Kools, colorsSet, yellow) ||
					!rightOf(colorsSet, green, colorsSet, ivory) ||
					!nextTo(nationalitySet, Norwegian, colorsSet, blue) {
					continue
				}
				for _, drinksSet := range combinations {
					//#4 #5 #13 #9
					if !(drinksSet[2] == milk) ||
						!match(drinksSet, coffee, colorsSet, green) ||
						!match(nationalitySet, Ukrainian, drinksSet, tea) ||
						!match(smokesSet, Lucky_Strikes, drinksSet, orange_juice) {
						continue
					}
					for _, petsSet := range combinations {
						//#3 #7 #11 #12
						if !match(nationalitySet, Spanish, petsSet, dog) ||
							!match(smokesSet, Old_Gold, petsSet, snails) ||
							!nextTo(smokesSet, Chesterfields, petsSet, fox) ||
							!nextTo(smokesSet, Kools, petsSet, horse) {
							continue
						}
						for i, drink := range drinksSet {
							if drink == water {
								rez.DrinksWater = nationality[nationalitySet[i]]
							}
						}
						for i, pet := range petsSet {
							if pet == zebra {
								rez.OwnsZebra = nationality[nationalitySet[i]]
							}
						}
					}
				}
			}
		}

	}

	return rez

}

func match(indexes1 []int, value1 int, indexes2 []int, value2 int) bool {
	for i, val := range indexes1 {
		if val == value1 && indexes2[i] == value2 {
			return true
		}
	}
	return false
}

func nextTo(indexes1 []int, value1 int, indexes2 []int, value2 int) bool {
	for i := 1; i < len(indexes1)-1; i++ {
		if indexes1[i] == value1 && indexes2[i-1] == value2 ||
			indexes1[i-1] == value1 && indexes2[i] == value2 {
			return true
		}
	}
	return false
}

func rightOf(indexes1 []int, value1 int, indexes2 []int, value2 int) bool {
	for i := 1; i < len(indexes1); i++ {
		if indexes1[i] == value1 && indexes2[i-1] == value2 {
			return true
		}
	}
	return false
}

// Simple algorithm to generate all permutations of n size set
func permutations(n int) (rez [][]int) {
	rez = make([][]int, 0)

	p := make([]int, n)
	used := make([]bool, n)
	for i := 0; i < n; i++ {
		used[i] = true
		p[0] = i
		gen(1, n, p, used, &rez)
		used[i] = false
	}
	return rez
}

func gen(m, n int, p []int, used []bool, rez *[][]int) {
	if m >= n {
		tmp := make([]int, len(p))
		copy(tmp, p)
		*rez = append(*rez, tmp)
	} else {
		for i := 0; i < n; i++ {
			if !used[i] {
				used[i] = true
				p[m] = i
				gen(m+1, n, p, used, rez)
				used[i] = false
			}
		}
	}
}
