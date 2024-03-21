population = 100 # initialize the player to start with a population of 100
gold_amount = 10000 # initialize the player to start with 10000 gold
housing_level = 1 # initialize the player to start with housing level 1

def main():
    global turn
    turn = 0
    
    print("Elder Reign")
    print("by Sorrxw \n")
    
    selections()
    options()
    
    
def selections():
    global king_name # make the player's name globally accessible
    global kingdom_name # make the player's kingdom name globally accessible
    king_name = input("Enter the name of your king: ")
    kingdom_name = input("Enter the name of your kingdom: ")
    
def options():
    print("\nWelcome your majesty, what would you like to do today?")
    print("\n1. Advance a turn")
    print("2. Check kingdom statistics\n")
    choice = int(input("Enter your choice: "))
    
    if (choice == 1):
        turn_advance()
    elif (choice == 2):
        show_stats()
    
def show_stats():
    print("\nWelcome to your kingdom, " + str(king_name))
    print("\n" + str(kingdom_name) + "'s stats: ")
    print("Population: " + str(population))
    print("Gold: " + str(gold_amount))
    print("Housing Level: " + str(housing_level))

# function for gold increasing
def gold_increase():
    if (population < 1000):
        gold_amount+=100
        
def population_increase():
    if (housing_level == 1):
        population += 10
       
# function for turn advancing / going forward in time
def turn_advance():
    turn+=1
    printer("It is now turn " + str(turn))
    
main()