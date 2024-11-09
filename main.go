package main

import (
	"fmt"
)

func checkWin(slot [][]string, multipliers map[string]uint) []uint {
	result := []uint{}

	for _, row := range slot {
		Wsym := row[0]
		for _, sym := range row[1:] {
			if sym != Wsym {
				Wsym = ""
				break
			}
		}
		if Wsym == "" {
			continue
		}
		result = append(result, multipliers[Wsym])
	}
	return result
}

func main() {
	symbols := map[string]uint{
		"A": 4,
		"B": 7,
		"C": 13,
		"D": 20,
	}
	multipliers := map[string]uint{
		"A": 20,
		"B": 15,
		"C": 10,
		"D": 5,
	}
	balance := uint(100)
	symbolArray := GenerateSymbolArray(symbols)
	getName()
	for balance > 0 {
		bet := getBet(balance)
		if bet == 0 {
			break
		}
		balance -= bet
		slot := MakeBandit(symbolArray, 3, 3)
		printSlot(slot)
		prizes := checkWin(slot, multipliers)
		for i, prize := range prizes {
			balance += (prize * bet)
			if prize > 0 {
				fmt.Printf("You have won $%d, (%dx) in line #%d\n", (prize * bet), prize, i+1)
			}
		}
	}
	fmt.Printf("You have exited the game with balance: $%d.\n", balance)
}
