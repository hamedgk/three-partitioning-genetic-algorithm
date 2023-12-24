package genetics

import (
	"bufio"
	"io"
	"strconv"
)

func Abs(x int) int{
	if x < 0{
		return -x
	}
	return x
}

func ReadInts(r io.Reader) ([ChromosomeSize]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result [ChromosomeSize]int
	var i = 0
	for scanner.Scan() || i < ChromosomeSize {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result[i] = x
		i++
	}
	return result, scanner.Err()
}
