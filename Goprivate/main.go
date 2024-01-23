package main

import (
	"fmt"

	"github.com/denisgariglio/fcutils-secret/pkg/events"
)

func main() {

	ed := events.NewEventDispatcher()
	fmt.Println(ed)

}
