package main

import "fmt"

func getName() string {
	name := ""
	fmt.Println("Welcome to the game!")
	fmt.Print("Enter your name: ")
	_, err := fmt.Scan(&name)
	if err != nil {
		return ""
	}
	fmt.Printf("Welcome %s, Let's Play\n", name)
	return name
}

func getBet(balance uint) uint {
	var bet uint
	for true {
		fmt.Printf("Enter your bet(Balance = $%d): ", balance)
		fmt.Scan(&bet)
		if bet > balance {
			fmt.Printf("You don't have enough balance to make a bet of $%d\n", bet)
			continue
		} else {
			break
		}
	}
	return bet
}
