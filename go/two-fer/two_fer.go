//Package twofer implements the solution to
//the twofer exercism problem
package twofer

import "fmt"

/*
ShareWith returns a customizeable string with the name variable inserted
*/
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
