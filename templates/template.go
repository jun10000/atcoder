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

func (input *Input) Int() int {
	input.Scanner.Scan()
	ret, err := strconv.Atoi(input.Scanner.Text())
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

func main() {
	//inputs := NewInput()
	//fmt.Println()
}
