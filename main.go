package main

import (
	"fmt"
)

func readersWithPublication(m map[string]map[string][]string) {
	countOfReaders := 0
	for k1 := range m {
		for _, v2 := range m[k1] {
			if len(v2) != 0 {
				countOfReaders++
				break
			}
		}
	}
	fmt.Println("Общее число читателей с изданиями на руках:", countOfReaders)
}

func valueOfPublication(m map[string]map[string][]string) {
	countOfPublication := 0
	iter := 0
	for k1 := range m {
		for _, v2 := range m[k1] {
			iter++
			countOfPublication += len(v2)
			if iter%2 == 0 && iter != 0 {
				fmt.Println("У читателя", k1, "изданий на руках:", countOfPublication)
				countOfPublication = 0
			}
		}
	}
}

func main() {
	m := map[string]map[string][]string{}

	m["anastasia"] = map[string][]string{
		"books":      {"mumu", "pushkin"},
		"periodical": {"vokrugsveta", "gazeta"},
	}

	m["sergey"] = map[string][]string{
		"books":      {"yama", "shatuny"},
		"periodical": {"rybalka", "playboy"},
	}

	m["gitler"] = map[string][]string{
		"books":      {},
		"periodical": {},
	}

	m["stalin"] = map[string][]string{
		"books":      {"bible"},
		"periodical": {},
	}

	readersWithPublication(m) // 3

	valueOfPublication(m)

	/*У читателя gitler изданий на руках: 0
	У читателя stalin изданий на руках: 1
	У читателя anastasia изданий на руках: 4
	У читателя sergey изданий на руках: 4*/

}
