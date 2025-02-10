package event

import (
	"fmt"

	"github.com/yahyaammar-dev/pacebe/types"
)

func NewListener() {
	// register all listeners

	// Register a listener for the "user.created" types.Event
	Register("user.created", func(e types.Event) {
		user, ok := e.Payload.(string)
		if !ok {
			fmt.Println("Invalid payload")
			return
		}
		fmt.Printf("Sending welcome email to %s\n", user)
	})

	// Register another listener for the same types.Event
	Register("user.created", func(e types.Event) {
		user, ok := e.Payload.(string)
		if !ok {
			fmt.Println("Invalid payload")
			return
		}
		fmt.Printf("Logging user creation for %s\n", user)
	})

}
