package partyrobot
import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	x := fmt.Sprintf("Welcome to my party, %s!",name)
    return x
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	x := fmt.Sprintf("Happy birthday %s! You are now %d years old!",name,age)
    return x
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor string, direction string, distance float64) string {
	x := Welcome(name)+fmt.Sprintf("\nYou have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.", table,direction,distance)+ fmt.Sprintf("\nYou will be sitting next to %s.",neighbor)
    return x
}
