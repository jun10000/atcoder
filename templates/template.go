package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

type Input struct {
	Scanner *bufio.Scanner
}

func NewInput() *Input {
	ret := &Input{
		Scanner: bufio.NewScanner(os.Stdin),
	}
	ret.Scanner.Split(bufio.ScanWords)
	ret.Scanner.Buffer([]byte{}, math.MaxInt32)
	return ret
}

func (input *Input) Text() string {
	input.Scanner.Scan()
	return input.Scanner.Text()
}

func (input *Input) Int() int {
	ret, err := strconv.Atoi(input.Text())
	if err != nil {
		panic(err)
	}
	return ret
}

func (input *Input) IntSlice(num int) []int {
	ret := make([]int, num)
	for i := range num {
		ret[i] = input.Int()
	}
	return ret
}

func (input *Input) Float() float64 {
	ret, err := strconv.ParseFloat(input.Text(), 64)
	if err != nil {
		panic(err)
	}
	return ret
}

func (input *Input) FloatSlice(num int) []float64 {
	ret := make([]float64, num)
	for i := range num {
		ret[i] = input.Float()
	}
	return ret
}

func AbsInt(i int) int {
	if i >= 0 {
		return i
	} else {
		return -i
	}
}

func MinInt(i, j int) int {
	if i <= j {
		return i
	} else {
		return j
	}
}

func MaxInt(i, j int) int {
	if i <= j {
		return j
	} else {
		return i
	}
}

func Pow2Int(i int) int {
	return i * i
}

func main() {
	// inputs := NewInput()
	// fmt.Println()
}
