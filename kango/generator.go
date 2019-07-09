package kango

import "math/rand"

type Generator struct {
	Rand *rand.Rand
	Max  int
}

func (gen *Generator) SetMax(max int) {
	gen.Max = max
}

func (gen *Generator) Int() int {
	return gen.Rand.Intn(gen.Max + 1)
}

func (gen *Generator) GetNums(n int) []int {
	if n == 0 {
		n = 10
	}

	// get unique numbers
	candidates := map[int]bool{}
	for len(candidates) < n {
		candidates[gen.Int()] = true
	}

	var numbers []int
	for k, _ := range candidates {
		numbers = append(numbers, k)
	}
	return numbers
}
