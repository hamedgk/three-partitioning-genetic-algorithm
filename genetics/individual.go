package genetics

type Individual struct{
	Chromosome Chromosome
	OverallValue int
}

func NewIndividual() Individual{
	indie := Individual{}
	indie.Chromosome = RandomChromosome()

	indie.OverallValue = 0
	for i := 0; i < 3; i++ {
		indie.OverallValue += Abs(indie.Chromosome.Values[i] - PerfectData.Values[i])
	}
	return indie
}

func (indie *Individual) Mutation(){
	indie.Chromosome.Mutation()
	indie.OverallValue = 0
	for i := 0; i < 3; i++ {
		indie.OverallValue += Abs(indie.Chromosome.Values[i] - PerfectData.Values[i])
	}
}