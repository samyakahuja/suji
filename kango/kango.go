package kango

import (
	"fmt"
	"math/rand"
	"strings"
)

var KangoDigits = []string{
	"零",
	"一",
	"二",
	"三",
	"四",
	"五",
	"六",
	"七",
	"八",
	"九",
	"十",
}

type Magnitude struct {
	Name  string
	Zeros int
}

var MajorMagnitudes = []Magnitude{
	Magnitude{"京", 1e16},
	Magnitude{"兆", 1e12},
	Magnitude{"億", 1e8},
	Magnitude{"万", 1e4},
	Magnitude{"", 1},
}

var MinorMagnitudes = []Magnitude{
	Magnitude{"千", 1000},
	Magnitude{"百", 100},
	Magnitude{"十", 10},
}

func New(seed int64) *Generator {
	return &Generator{
		Rand: rand.New(rand.NewSource(seed)),
		Max:  10,
	}
}

func SpellKango(n int) string {
	if n < 10 {
		return KangoDigits[n]
	}
	var value string
	numbers := []string{}

	for _, mag := range MajorMagnitudes {
		value, n = SpellChunk(n, mag.Name, mag.Zeros)
		numbers = append(numbers, value)
	}
	return strings.Join(numbers, "")
}

func SpellChunk(n int, mag string, scale int) (string, int) {
	if n == 0 {
		return "", 0
	}

	remainder := n % scale
	n = n / scale

	if n == 0 {
		return "", remainder
	}

	var value string
	numbers := []string{}

	for _, mag := range MinorMagnitudes {
		value, n = SpellDigit(n, mag.Name, mag.Zeros)
		numbers = append(numbers, value)
	}

	if n > 0 {
		if mag == "" || n > 1 {
			numbers = append(numbers, KangoDigits[n])
		}
	}

	numbers = append(numbers, mag)
	if remainder > 0 {
		numbers = append(numbers, " ")
	}

	return strings.Join(numbers, ""), remainder
}

func SpellDigit(n int, mag string, scale int) (string, int) {
	digit := n / scale
	remainder := n % scale

	if digit == 0 {
		return "", remainder
	}

	var value string
	var padding string

	if digit > 0 {
		if digit > 1 {
			value = KangoDigits[digit]
		}
		if remainder > 0 {
			padding = " "
		}
	}

	return fmt.Sprintf("%s%s%s", value, mag, padding), remainder
}
