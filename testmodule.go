package testmod

import "fmt"

// Hi returns a friendly greeting v.1
func Hi(name string) string {
	return fmt.Sprintf("Hi, %s!", name)
}
