package main
import ("fmt")

var population int = 100 // initialize the player to start with a population of 100
var gold_amount int = 10000 // initialize the player to start with 10000 gold
var housing_level int = 1 // initialize the player to start with housing level 1
var turn int = 1 // initialize the player to start on turn 1
var king_name string
var kingdom_name string


func main() {
  fmt.Println("Elder Reign")
  fmt.Println("by Sorrxw")
  
  selections() // call the function to allow for the player to set their name
  options()
}

func selections() {
    fmt.Print("\nEnter the name of your king: ")
    fmt.Scan(&king_name)
    
    fmt.Print("Enter the name of your kingdom: ")
    fmt.Scan(&kingdom_name)
}

func options() {
    
    var options_choice int // variable to collect user input for their selection
    
    fmt.Println("\nWelcome your majesty, what would you like to do today?")
    fmt.Println("1. Advance a turn")
    fmt.Println("2. Check kingdom statistics")
    fmt.Println("3. Upgrade buildings")
    fmt.Print("Enter your choice: ")
    fmt.Scan(&options_choice)
    
    switch (options_choice) {
        case 1:
            turn_advance()
        case 2:
            show_stats()
        case 3:
            upgrade_buildings()
    }
}

func show_stats() {
    fmt.Println("\nWelcome to your kingdom king,", king_name + "!")
    fmt.Println(kingdom_name + "'s stats:")
    fmt.Println("Population:", population)
    fmt.Println("Gold:", gold_amount)
    fmt.Println("Housing Level:", housing_level)
    options() // call the options function so that the user automatically is returned to the main screen
}

func turn_advance() {
    turn = turn + 1 // move to the next sequential turn
    fmt.Println("\nIt is now turn", turn) // tell the user what turn it is
    population_increase() // call the population increase function
    gold_increase() // call the gold increase function
    options() // call the options function so that the user automatically is returned to the main screen
}

func gold_increase() {
    
    if (population < 1000) {
        gold_amount = gold_amount + 100
    }
    
    if (population >= 1000 && population < 5000) {
        gold_amount = gold_amount + 250
    }
    
    if (population >= 5000 && population < 10000) {
        gold_amount = gold_amount + 500
    }
    
    if (population >= 10000) {
        gold_amount = gold_amount + 1000
    }
    
}

func population_increase() {
    
    switch (housing_level) {
        case 1:
            population = population + 10
        case 2:
            population = population + 50
        case 3:
            population = population + 100
    }
    
}

func upgrade_buildings() { // WORK IN PROGRESS
    fmt.Println("Buildings to upgrade: ")
    fmt.Println("Upgrade housing: ")
}