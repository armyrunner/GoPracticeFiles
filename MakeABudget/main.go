package main

import (
	"fmt"
	"os"

	bp "personal-budget/budgetProgram"
	"time"
)

var months = []time.Month{
	time.January,
	time.February,
	time.March,
	time.April,
	time.May,
	time.June,
	time.July,
	time.August,
	time.September,
	time.October,
	time.November,
	time.December,
}

func createBudgetMenu() {
	i := 1
	fmt.Println("Which month do you wnat to create a budget for?")

	for _, month := range months{
		fmt.Printf("%v. %v\n",i , month )
		i++
	}

}

func transactionMenu(){
	fmt.Println("What do you want to do?")
}

func printBudgetMenu(){
	fmt.Println("What do you want to print?")
}

func printMenu() {

	menu := "1) Create A Budget\n" +
			"2) Transactions\n" +
			"3) Print Budget\n" +
			"4) Exit\n "

	fmt.Print(menu)
}

func main() {

	fmt.Println( "Please make a selection!")
	fmt.Println(" ")
	choice := -1
	
	printMenu()
	fmt.Scanln(&choice)

	for choice != 4 {
		switch choice {
		case 1:
			createBudgetMenu()
		case 2:
			transactionMenu()
		case 3:
			printBudgetMenu()
		default:
			fmt.Println("Invalid Choice, Please try agian.")
		}
		printMenu()
		fmt.Scanln(&choice)
	
	}
	os.Exit(3)


	// bu, _ := bp.CreateBudget(time.January, 1000)
	// bu.AddItem("bananas", 10.0)

	// fmt.Println("Items in January:", len(bu.Items))
	// fmt.Printf("Current cost for January: $%.2f \n", bu.CurrentCost())

	// bp.CreateBudget(time.February, 1000)

	// bu2 := bp.GetBudget(time.February)
	// bu2.AddItem("bananas", 10.0)
	// bu2.AddItem("coffee", 3.99)
	// bu2.AddItem("gym", 50.0)
	// bu2.RemoveItem("coffee")
	// fmt.Println("Items in February:", len(bu2.Items))
	// fmt.Printf("Current cost for February: $%.2f \n", bu2.CurrentCost())

	// fmt.Println("Resetting Budget Report...")
	// bp.InitializeReport()

	// for _, month := range months {
	// 	_, err := bp.CreateBudget(month, 100.00)
	// 	if err == nil {
	// 		fmt.Println("Budget created for", month)
	// 	} else {
	// 		fmt.Println("Error creating budget:", err)
	// 	}
	// }

	// _, err := bp.CreateBudget(time.December, 100.00)
	// fmt.Println(err)
}
