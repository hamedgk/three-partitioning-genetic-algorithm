package main

import (
	"fmt"
	"genetic/genetics"
	"math/rand"
)

func main() {
	bestChild := genetics.Individual{OverallValue: genetics.MaxInt}
	society := genetics.NewSociety()

	for {
		newSociety := genetics.Society{}
		for i := 0; i < genetics.Population; i++ {
			male := society.RandomSelection()
			female := society.RandomSelection()
			child := society.Reproduce(male, female)

			smallChance := rand.Intn(1000)
			if smallChance < 50 {
				child.Mutation()
			}
			if child.OverallValue < bestChild.OverallValue {
				bestChild = child
				fmt.Println(bestChild.OverallValue)
			}
			newSociety.Population[i] = child
		}
		newSociety.Update()
		society = newSociety
	}

}
