package random

import (
	"fmt"
	"testing"
	"time"
)

func TestGenericRandomGenerator(t *testing.T){
	/* read the list until head is not nil */
	g := NewRandomGenerator()
	for i := 0; i < 10; i++ {
		num, _ := g.NextNumber()
		//printGeneratorParameters(g)
		fmt.Printf(" %d ", num)
	}
	fmt.Println()
}

func TestRandomGeneratorWithSeed(t *testing.T){
	//Using System tick count as seed
	g := NewGeneratorWithSeed(time.Now().UTC().Nanosecond())
	nums, _ := g.NextNumbersWithinMax(100, 20)
	printNumberList(nums)
	fmt.Println()
}

func TestRandomGeneratorWithSeedForAdditive(t *testing.T){
	//Using System tick count as seed
	g := NewGeneratorWithSeedAndAlgorithm(time.Now().UTC().Nanosecond(), RandomGeneratorAlgorithm(Additive))
	nums, _ := g.NextNumbersWithinMax(100, 65)
	printNumberList(nums)
	fmt.Println()
}

func printNumberList(nums []int) {
	for _, num := range nums {
		fmt.Printf("%d ", num)
	}
}

func printGeneratorParameters(g *Generator) {
	fmt.Printf("algorithm	= %s\n", g.algorithmType)
	fmt.Printf("seed      = %d\n", g.xSeed)
	fmt.Printf("multipler = %d\n",g.multiplier)
	fmt.Printf("modulus   = %d\n",g.modulus)
	fmt.Printf("increment = %d\n", g.increment)
	fmt.Printf("lastNumber = %d\n", g.lastNumber)
}