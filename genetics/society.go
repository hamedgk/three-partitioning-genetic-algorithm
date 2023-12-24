package genetics

import (
	"math/rand"
)

const (
	Population = 50
)

type Society struct {
	Population [Population]Individual
	//this is used for normalization
	MaxOverallValue   int
	SumOfNormalValues int
}

func NewSociety() Society {
	society := Society{}
	max := MinInt
	for index := range society.Population {
		society.Population[index] = NewIndividual()
		if society.Population[index].OverallValue > max {
			max = society.Population[index].OverallValue
		}
	}

	society.MaxOverallValue = max

	for index := range society.Population {
		value := society.Population[index].OverallValue
		society.SumOfNormalValues += (society.MaxOverallValue - value) + 100
	}
	return society
}

func (society *Society) Update() {
	max := MinInt
	for index := range society.Population {
		if society.Population[index].OverallValue > max {
			max = society.Population[index].OverallValue
		}
	}

	society.MaxOverallValue = max

	society.SumOfNormalValues = 0
	for index := range society.Population {
		value := society.Population[index].OverallValue
		society.SumOfNormalValues += (society.MaxOverallValue - value) + 100
	}
}

func (society *Society) RandomSelection() int {
	random := rand.Intn(society.SumOfNormalValues)
	var s int
	for index := range society.Population {
		value := society.Population[index].OverallValue
		s += (society.MaxOverallValue - value) + 100
		if random <= s {
			return index
		}
	}

	l := len(society.Population)
	return l - 1
}

func (society *Society) Reproduce(male, female int) Individual {
	var indie = Individual{}
	mc, fc := &society.Population[male].Chromosome, &society.Population[female].Chromosome
	indie.Chromosome = CrossOver(mc, fc)

	indie.OverallValue = 0
	for i := 0; i < 3; i++ {
		indie.OverallValue += Abs(indie.Chromosome.Values[i] - PerfectData.Values[i])
	}
	return indie
}
