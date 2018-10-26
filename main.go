package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	baz, err := initializeBaz(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(baz.X)
}
