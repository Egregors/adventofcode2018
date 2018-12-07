package main

import (
	"log"
	"time"

	"github.com/Egregors/adventofcode2018/utils"
)

func main() {
	log.Println("D2")
	log.Printf("https://adventofcode.com/2018/day/2\n\n")

	data, err := utils.ReadFileLines("d2/input.txt")
	if err != nil {
		panic(err)
	}

	start := time.Now()
	res, err := solve(data)
	if err != nil {
		panic(err)
	}
	end := time.Since(start)

	log.Printf("Res p1: %d in %s", res, end)

	start = time.Now()
	res2, err := solve2(data)
	if err != nil {
		panic(err)
	}
	end = time.Since(start)
	log.Printf("Res p2: %s in %s", res2, end)

}

func eq(a, b map[string]int) bool {
	diff := 0
	for k, v := range a {
		if w, ok := b[k]; !ok || v != w {
			diff++
		}
	}
	if diff == 1 {
		return true
	}
	return false
}

func solve(data []string) (int, error) {
	var (
		two, three int
	)
	for i := 0; i < len(data); i++ {
		s := data[i]

		chars := make(map[string]int)
		groups := make(map[int]bool)
		groups[2], groups[3] = false, false

		for _, c := range s {
			chars[string(c)]++
		}

		for _, v := range chars {
			if v == 2 {
				groups[2] = true
			}
			if v == 3 {
				groups[3] = true
			}
		}

		if groups[2] {
			two++
		}

		if groups[3] {
			three++
		}

	}
	return two * three, nil
}

func solve2(data []string) (string, error) {

	IDs := make(map[string]map[string]int, len(data))

	for i := 0; i < len(data); i++ {
		s := data[i]
		ID := make(map[string]int)
		for _, ch := range s {
			ID[string(ch)]++
		}
		IDs[s] = ID
	}

	var goalIDs []string

	for _, leters := range IDs {
		for idd, leterss := range IDs {
			if eq(leters, leterss) {
				goalIDs = append(goalIDs, idd)
			}
		}
	}

	log.Println(goalIDs)

	var res string

	s1 := IDs[goalIDs[0]]
	s2 := IDs[goalIDs[1]]

	for k := range s1 {
		if _, ok := s2[k]; ok {
			res = res + k
		}
	}

	return res, nil
}
