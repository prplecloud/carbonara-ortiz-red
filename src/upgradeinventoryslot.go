package red

import "fmt"

func (c *Character) UpgradeInventorySlot() {
	if c.inventoryUpgrades < 3 {
		c.inventoryCapacity += 10
		c.inventoryUpgrades++
		fmt.Println("Inventory capacity increased to", c.inventoryCapacity)
	} else {
		fmt.Println("You can only upgrade your inventory capacity three times.")
	}
}

func (c *Character) MaxItem() bool {
	nbItem := 0
	for _, item := range c.inventory {
		nbItem += item
	}
    if nbItem >= c.inventoryCapacity {
        fmt.Println("Inventory is full, you cannot add more items.")
        return false

    } else {
        return true
    }
}