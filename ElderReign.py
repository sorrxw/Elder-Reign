population = 100 # initialize the player to start with a population of 100
gold_amount = 10000 # initialize the player to start with 10000 gold
turn = 0

def main():
    print("Elder Reign")
    print("by Sorrxw \n")
    
    selections()
    options()
    
    
def selections():
    king_name = input("Enter the name of your king: ")
    kingdom_name = input("Enter the name of your kingdom: ")
    
def options():
    print("Welcome to your kingdom, " + str(king_name))
    print("\n" + str(kingdom_name) + "'s stats: ")
    print("Population: " + int(population))
    print("Gold: " + int(gold_amount))

# function for gold increasing
def gold_increase():
    if (population < 1000):
        gold_amount+=100
       
# function for turn advancing / going forward in time
def turn_advance():
    turn+=1
    
main()