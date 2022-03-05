package main

import "fmt"

func main() {
	money := 550
	water := 400
	milk := 540
	coffeeBeans := 120
	disposableCups := 9

	var action string
	for action != "exit" {
		action = askForAction()
		processTheAction(action, &water, &milk, &coffeeBeans, &disposableCups, &money)
		fmt.Println()
	}
}

func displaySupplies(water, milk, coffeeBeans, disposableCups, money int) {
	fmt.Println()
	fmt.Println("The coffee machine has:")
	fmt.Println(water, "of water")
	fmt.Println(milk, "of milk")
	fmt.Println(coffeeBeans, "of coffee beans")
	fmt.Println(disposableCups, "of disposable cups")
	fmt.Println(money, "of money")
}

func askForAction() string {
	var action string
	fmt.Print("Write action (buy, fill, take, remaining, exit):\n> ")
	fmt.Scan(&action)
	return action
}

func processTheAction(action string, water, milk, coffeeBeans, disposableCups, money *int) {
	switch action {
	case "buy":
		processTheBuyAction(water, milk, coffeeBeans, disposableCups, money)
	case "fill":
		processTheFillAction(water, milk, coffeeBeans, disposableCups)
	case "take":
		processTheTakeAction(money)
	case "remaining":
		displaySupplies(*water, *milk, *coffeeBeans, *disposableCups, *money)
	}
}

func processTheBuyAction(water, milk, coffeeBeans, disposableCups, money *int) {
	var coffeeType string
	var consumableWater, consumableMilk, consumableCoffee, profit int
	fmt.Print("\nWhat do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, 4 - strongBlack, back - to main menu:\n> ")
	fmt.Scan(&coffeeType)
	switch coffeeType {
	case "1":
		consumableWater, consumableMilk, consumableCoffee, profit = sellEspresso()
	case "2":
		consumableWater, consumableMilk, consumableCoffee, profit = sellLatte()
	case "3":
		consumableWater, consumableMilk, consumableCoffee, profit = sellCappuccino()
	case "4":
		consumableWater, consumableMilk, consumableCoffee, profit = sellStrongBlack()
	case "back":
		return
	}

	switch {
	case *water < consumableWater:
		fmt.Println("Sorry, not enough water!")
		return
	case *milk < consumableMilk:
		fmt.Println("Sorry, not enough milk!")
		return
	case *coffeeBeans < consumableCoffee:
		fmt.Println("Sorry, not enough coffee beans!")
		return
	case *disposableCups == 0:
		fmt.Println("Sorry, not enough disposable cups!")
		return
	}

	fmt.Println("I have enough resources, making you a coffee!")
	*water -= consumableWater
	*milk -= consumableMilk
	*coffeeBeans -= consumableCoffee
	*disposableCups -= 1
	*money += profit
}

func sellEspresso() (waterPerCup, milkPerCup, coffeePerCup, pricePerCup int) {
	waterPerCup = 250
	milkPerCup = 0
	coffeePerCup = 16
	pricePerCup = 4
	return
}

func sellLatte() (waterPerCup, milkPerCup, coffeePerCup, pricePerCup int) {
	waterPerCup = 350
	milkPerCup = 75
	coffeePerCup = 20
	pricePerCup = 7
	return
}

func sellCappuccino() (waterPerCup, milkPerCup, coffeePerCup, pricePerCup int) {
	waterPerCup = 200
	milkPerCup = 100
	coffeePerCup = 12
	pricePerCup = 6
	return
}

func sellStrongBlack() (waterPerCup, milkPerCup, coffeePerCup, pricePerCup int) {
	waterPerCup = 100
	milkPerCup = 0
	coffeePerCup = 20
	pricePerCup = 3
	return
}

func processTheFillAction(water, milk, coffeeBeans, disposableCups *int) {
	fmt.Println()
	incomingWater := takeWater()
	incomingMilk := takeMilk()
	incomingCoffee := takeCoffee()
	incomingCups := takeDisposableCups()

	*water += incomingWater
	*milk += incomingMilk
	*coffeeBeans += incomingCoffee
	*disposableCups += incomingCups
}

func takeWater() (water int) {
	fmt.Print("Write how many ml of water you want to add:\n> ")
	fmt.Scan(&water)
	return
}

func takeMilk() (milk int) {
	fmt.Print("Write how many ml of milk you want to add:\n> ")
	fmt.Scan(&milk)
	return
}

func takeCoffee() (coffee int) {
	fmt.Print("Write how many grams of coffee beans you want to add:\n> ")
	fmt.Scan(&coffee)
	return
}

func takeDisposableCups() (cups int) {
	fmt.Print("Write how many disposable coffee cups you want to add:\n> ")
	fmt.Scan(&cups)
	return
}

func processTheTakeAction(money *int) {
	fmt.Printf("\nI gave you $%d\n", *money)
	*money = 0
}
