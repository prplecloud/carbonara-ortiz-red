package red

import ("fmt"
		"os"
)

func menu1(c1 *Character, m1 *Shop) {

	var choice int

fmt.Println("-------------------------------------")
fmt.Println("|	     MAIN MENU 		    |")
fmt.Println("|1. Display Character Informations  |")
fmt.Println("|2. Display Inventory		    |")
fmt.Println("|3. Verify HP		   	    |")
fmt.Println("|4. Take a Hp Potion	  	    |")
fmt.Println("|5. Take a Poison Potion	    |")
fmt.Println("|6. Display Shop		    |")
fmt.Println("|7. Use SpellBook : FireBall	    |")
fmt.Println("|8. Black-Smith	    		    |")
fmt.Println("|9. Show equiped Equipment	    		    |")
fmt.Println("|15. Quit Game			    |")
fmt.Println("-------------------------------------")
fmt.Println("")
fmt.Println("")
fmt.Println("")

fmt.Scan(&choice)

	switch choice {
	case 1:
		c1.DisplayInfo()
		///DisplaycharacInfo
	case 2:
		c1.AccessInventory()
		///DisplayInventory
	case 3:
		c1.Death()
		///Vérifier les HP 
	case 4:
		c1.Takepot()
		///Prendre une potion de HP
	case 5:
		c1.PoisonPot()
	case 6:
		m1.DisplayShop(c1)
	case 7:
		c1.SpellBook("FireBall")
		///Utilise SpellBook
	case 8:
		c1.BlackSmith()
		///Montre Forgeron
	case 9:
		///Show equipment
		
	case 15 :
		os.Exit(0)
		///Quit
	default :
	fmt.Println("------------------------------------")
		fmt.Println("Choice is not valid. Please choose a valid option")
	}
}