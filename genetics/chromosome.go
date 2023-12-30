package genetics

import (
	"math/rand"
)

const (
	ChromosomeSize = 100

	MaxUint = ^uint(0)
	MinUint = 0
	MaxInt  = int(MaxUint >> 1)
	MinInt  = -MaxInt - 1
)

type Chromosome struct {
	Distribution [ChromosomeSize]uint8
	Counts       [3]int
	Values       [3]int
}

func (chromosome *Chromosome) Mutation() {
	random := rand.Intn(ChromosomeSize)
	randomSibling := rand.Intn(3)

	chosenSibling := chromosome.Distribution[random]
	chosenGround := GroundData[random]

	chromosome.Values[chosenSibling] -= chosenGround
	chromosome.Counts[chosenSibling]--

	chromosome.Values[randomSibling] += chosenGround
	chromosome.Counts[randomSibling]++

	chromosome.Distribution[random] = uint8(randomSibling)
}

func CrossOver(male, female *Chromosome) Chromosome {
	var chrom = Chromosome{}
	for i := 0; i < ChromosomeSize/2; i++ {
		chrom.Distribution[i] = male.Distribution[i]
	}
	for i := ChromosomeSize / 2; i < ChromosomeSize; i++ {
		chrom.Distribution[i] = female.Distribution[i]
	}

	//for i := 0; i<ChromosomeSize; i += 2{
	//	chrom.Distribution[i] = male.Distribution[i]
	//	chrom.Distribution[i+1] = female.Distribution[i+1]
	//}

	for i := 0; i < ChromosomeSize; i++ {
		sibling := chrom.Distribution[i]
		switch sibling {
		case 0:
			chrom.Counts[0]++
			chrom.Values[0] += GroundData[i]
		case 1:
			chrom.Counts[1]++
			chrom.Values[1] += GroundData[i]
		case 2:
			chrom.Counts[2]++
			chrom.Values[2] += GroundData[i]
		default:
			panic("invalid sibling...")
		}
	}
	return chrom
}

func RandomChromosome() Chromosome {
	var chrom = Chromosome{}
	for i := 0; i < ChromosomeSize; i++ {
		randomSibling := rand.Intn(100)

		switch {
		case randomSibling > -1 && randomSibling < 41:
			chrom.Counts[0]++
			chrom.Values[0] += GroundData[i]
			chrom.Distribution[i] = 0
		case randomSibling > 40 && randomSibling < 81:
			chrom.Counts[1]++
			chrom.Values[1] += GroundData[i]
			chrom.Distribution[i] = 1
		case randomSibling > 80:
			chrom.Counts[2]++
			chrom.Values[2] += GroundData[i]
			chrom.Distribution[i] = 2
		default:
			panic("invalid random number...")
		}
	}
	return chrom
}
