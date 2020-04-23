package main

import (
	log "github.com/sirupsen/logrus"
	"math"
)

//LFSR - linear-feedback shift register
func LFSR5(initialState string, linearFeedback string) (int, []uint8) {
	log.Info("LFSR-5")
	a := make([]uint8, 0)
	b := make([]uint8, 0)
	с := make([]uint8, 0)
	result := make([]uint8, 0)
	for _, value := range initialState {
		a = append(a, uint8(value)-uint8('0'))
	}
	for _, value := range linearFeedback {
		с = append(с, uint8(value)-uint8('0'))
	}
	log.Info("initial state values")
	for index, value := range a {
		log.Info(index, value)
	}
	log.Info("linear-feedback values")
	for index, value := range с {
		log.Info(index, value)
	}
	var resultOfBooleanMultiplication uint8
	var SL uint8
	for i := range a {
		b = append(b, a[i])
	}

	for j := 0; j < 31; j++ {
		for index, value := range b {
			resultOfBooleanMultiplication = value * с[index]
			if index == 0 {
				SL = resultOfBooleanMultiplication
			} else {
				SL = SL ^ resultOfBooleanMultiplication
			}
		}
		log.Info("Elem:", b[4])
		result = append(result, b[4])
		if result[len(result)-1] == a[0] && len(result) > 9 {
			flag := true
			for j := 2; j <= 5; j++ {
				if result[len(result)-j] != a[j-1] {
					flag = false
					break
				}
			}
			if flag == true {
				return j - 4, result[:j-4]
			}

		}
		for i := 4; i > 0; i-- {
			b[i] = b[i-1]
		}
		b[0] = SL
	}
	return 31, result
}

func LFSR7(initialState string, linearFeedback string) (int, []uint8) {
	log.Info("LFSR-7")
	a := make([]uint8, 0)
	b := make([]uint8, 0)
	с := make([]uint8, 0)
	result := make([]uint8, 0)
	for _, value := range initialState {
		a = append(a, uint8(value)-uint8('0'))
	}
	for _, value := range linearFeedback {
		с = append(с, uint8(value)-uint8('0'))
	}
	log.Info("initial state values")
	for index, value := range a {
		log.Info(index, value)
	}
	log.Info("linear-feedback values")
	for index, value := range с {
		log.Info(index, value)
	}
	var resultOfBooleanMultiplication uint8
	var SL uint8
	for i := range a {
		b = append(b, a[i])
	}
	for j := 0; j < 127; j++ {
		for index, value := range b {
			resultOfBooleanMultiplication = value * с[index]
			if index == 0 {
				SL = resultOfBooleanMultiplication
			} else {
				SL = SL ^ resultOfBooleanMultiplication
			}
		}
		log.Info("Elem:", b[6])
		result = append(result, b[6])
		if result[len(result)-1] == a[0] && len(result) > 13 {
			flag := true
			for j := 2; j <= 7; j++ {
				if result[len(result)-j] != a[j-1] {
					flag = false
					break
				}
			}
			if flag == true {
				return j - 6, result[:j-6]
			}

		}
		for i := 6; i > 0; i-- {
			b[i] = b[i-1]
		}
		b[0] = SL
	}
	return 127, result
}
func LFSR8(initialState string, linearFeedback string) (int, []uint8) {
	log.Info("LFSR-8")
	a := make([]uint8, 0)
	b := make([]uint8, 0)
	с := make([]uint8, 0)
	result := make([]uint8, 0)
	for _, value := range initialState {
		a = append(a, uint8(value)-uint8('0'))
	}
	for _, value := range linearFeedback {
		с = append(с, uint8(value)-uint8('0'))
	}
	log.Info("initial state values")
	for index, value := range a {
		log.Info(index, value)
	}
	log.Info("linear-feedback values")
	for index, value := range с {
		log.Info(index, value)
	}
	var resultOfBooleanMultiplication uint8
	var SL uint8
	for i := range a {
		b = append(b, a[i])
	}
	for j := 0; j < 255; j++ {
		for index, value := range b {
			resultOfBooleanMultiplication = value * с[index]
			if index == 0 {
				SL = resultOfBooleanMultiplication
			} else {
				SL = SL ^ resultOfBooleanMultiplication
			}
		}
		log.Info("Elem:", b[7])
		result = append(result, b[7])
		if result[len(result)-1] == a[0] && len(result) > 15 {
			flag := true
			for j := 2; j <= 8; j++ {
				if result[len(result)-j] != a[j-1] {
					flag = false
					break
				}
			}
			if flag == true {
				return j - 7, result[:j-7]
			}

		}
		for i := 7; i > 0; i-- {
			b[i] = b[i-1]
		}
		b[0] = SL
	}
	return 255, result
}
func MakeGreaterSequence(seq1 []uint8, seq2 []uint8, seq3 []uint8) ([]uint8, []uint8, []uint8) {
	for len(seq1) < 10000 {
		for j := range seq1 {
			seq1 = append(seq1, seq1[j])
		}
	}
	for len(seq2) < 10000 {
		for j := range seq2 {
			seq2 = append(seq2, seq2[j])
		}
	}
	for len(seq3) < 10000 {
		for j := range seq3 {
			seq3 = append(seq3, seq3[j])
		}
	}
	seq1 = seq1[:10000]
	seq2 = seq2[:10000]
	seq3 = seq3[:10000]
	return seq1, seq2, seq3
}
func GeffeGenerator(seq1 []uint8, seq2 []uint8, seq3 []uint8) []uint8 {
	g := make([]uint8, 10000, 10000)
	for t := 0; t < 10000; t++ {
		g[t] = (seq1[t] * seq2[t]) ^ ((seq1[t] ^ uint8(1)) * seq3[t])
	}
	return g
}
func Statistics(g []uint8) {
	amountOfNulls := 0
	amountOfOnes := 0
	for t := 0; t < 10000; t++ {
		if g[t] == 0 {
			amountOfNulls++
		} else {
			amountOfOnes++
		}
	}
	log.Info("amount of nulls: ", amountOfNulls)
	log.Info("amount of ones: ", amountOfOnes)
}
func Ri(g []uint8) [5]int {
	sum := 0
	var r [5]int
	for i := 1; i <= 5; i++ {
		for t := 0; t < 10000-i; t++ {
			sum += Teta(g[t] ^ g[t+1])
		}
		r[i-1] = sum
	}
	return r
}
func Teta(x uint8) int {
	return int(math.Pow(-1, float64(x)))
}
func main() {
	period1, seq1 := LFSR5("10010", "01101")
	log.Info("Period of LFSR5:", period1)
	log.Info("Sequence 1:", seq1)
	period2, seq2 := LFSR7("0010000", "0101101")
	log.Info("Period of LFSR7:", period2)
	log.Info("Sequence 2:", seq2)
	period3, seq3 := LFSR8("01100011", "11110011")
	log.Info("Period of LFSR8:", period3)
	log.Info("Sequence 3:", seq3)
	seq1, seq2, seq3 = MakeGreaterSequence(seq1, seq2, seq3)
	g := GeffeGenerator(seq1, seq2, seq3)
	Statistics(g)
	r := Ri(g)
	log.Info("R[i]:", r)
}
