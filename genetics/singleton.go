package genetics

import (
	"fmt"
	"os"
)

const (
	path = "data/dataset6-sorted.txt"
)

var (
	GroundData  [ChromosomeSize]int
	PerfectData PerfectHeritageData
)

func init() {
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	GroundData, err = ReadInts(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	PerfectData = InitData(&GroundData)
}

type PerfectHeritageData struct {
	Sum    int
	Values [3]int
}

func InitData(data *[ChromosomeSize]int) PerfectHeritageData {
	perfect := PerfectHeritageData{}
	for i := 0; i < ChromosomeSize; i++ {
		perfect.Sum += (*data)[i]
	}

	perfect.Values[0] = int(float64(perfect.Sum) * 0.4)
	perfect.Values[1] = int(float64(perfect.Sum) * 0.4)
	perfect.Values[2] = int(float64(perfect.Sum) * 0.2)

	return perfect
}
