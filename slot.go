package main

import (
	"fmt"
	"math/rand"
)

func GenerateSymbolArray(symbols map[string]uint) []string {
	symbolArray := []string{}
	for sym, count := range symbols {
		for i := uint(0); i < count; i++ {
			symbolArray = append(symbolArray, sym)
		}
	}
	return symbolArray
}

func GenerateRandomNumber(min int, max int) int {
	num := rand.Intn(int(max-min+1)) + int(min)
	return num
}

func MakeBandit(symbolArray []string, rows int, cols int) [][]string {
	slot := [][]string{}
	for i := 0; i < 3; i++ {
		slot = append(slot, []string{})
	}
	for col := 0; col < cols; col++ {
		visited := map[int]bool{}
		for row := 0; row < rows; row++ {
			for true {
				k := GenerateRandomNumber(0, len(symbolArray)-1)
				_, exists := visited[k]
				if !exists {
					slot[row] = append(slot[row], symbolArray[k])
					visited[k] = true
					break
				}
			}
		}

	}
	return slot
}

func printSlot(slot [][]string) {
	for _, row := range slot {
		for j, sym := range row {
			fmt.Print(sym)
			if j != len(row)-1 {
				fmt.Printf(" | ")
			}
		}
		fmt.Println()
	}
}
