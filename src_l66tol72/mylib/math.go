/*
mylib is my original library
*/
package mylib

//lesson59
// Averate returns the agerage of a series of numbers
func Average(s []int) int {
	// total number
	total := 0
	for _, i := range s {
		total += i
	}
	return int(total / len(s))
}
