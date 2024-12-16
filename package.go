package main
import "strconv"
import "fmt"
func Welcome(name string) string {
	return "Welcome, " + name + "!"
}
func HappyBirthday(name string, age int) string {
	return "Happy birthday " + name + "! You are now " + strconv.Itoa(age) + " years old!" 
}
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	tableNumber := fmt.Sprintf("%03d", table)
	distanceFormatted := fmt.Sprintf("%.1f", distance)
	return fmt.Sprintf("Welcome to my party, %s!\nYou have been assigned to table %s. Your table is %s, exactly %s meters from here.\nYou will be sitting next to %s.", 
		name, tableNumber, direction, distanceFormatted, neighbor)
}