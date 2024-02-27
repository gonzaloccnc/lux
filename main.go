package main

import (
	"fmt"
	"log"

	sc "github.com/gonzaloccnc/lux/components/singlechoice"
)

func main() {
	var items []string = make([]string, 20)

	for i := 0; i < 20; i++ {
		items[i] = fmt.Sprintf("Item %d", i)
	}
	s := sc.New(items, "Select one item")

	choice, err := s.Run()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("the choice is: %s", choice)

	s2 := sc.New(items, "Select one item")

	choice2, err := s2.Run()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("the choice is: %s", choice2)
}
