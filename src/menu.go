package red

import (
	"fmt"
	"os"
)

func Menu1(c1 *Character, m1 *Shop, g1 *Goblin) {
	fmt.Println("-------------------------------------")
	fmt.Println("|	     MAIN MENU 		    |")
	fmt.Println("|1. Display Character Informations  |")
	fmt.Println("|2. Display Inventory		    |")
	fmt.Println("|3. Display Shop		    |")
	fmt.Println("|4. Black-Smith	    		    |")
	fmt.Println("|5. Training Fight		    |")
	fmt.Println("|6. Bonus : Who are they	    |")
	fmt.Println("|7. Quit Game			    |")
	fmt.Println("-------------------------------------")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	fmt.Scan(&choice)

	switch choice {
	case 1:
		c1.DisplayInfo()
	case 2:
		c1.AccessInventory()
	case 3:
		m1.DisplayShop(c1)
	case 4:
		c1.BlackSmith()
	case 5:
		TrainingFight(c1, g1)
	case 6:
		fmt.Println("ABBA")
		fmt.Println("Steven Spielberg")
	case 7:
		os.Exit(0)
		///Quit
	default:
		fmt.Println("------------------------------------")
		fmt.Println("Choice is not valid. Please choose a valid option")
	}
}
