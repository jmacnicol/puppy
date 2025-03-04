package puppy

import (
	"github.com/jmacnicol/dog"
)

// don't forget to tag this
func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func version1_2() string {
	return "I'm from version 1.2.0"
}
