package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/simplereach/timeutils"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatalf("Usage: %s <utc date>", os.Args[0])
	}
	t, err := timeutils.ParseDateString(strings.Join(os.Args[1:], " "))
	if err != nil {
		log.Fatalf("Error parsing input date: %s\n", err)
	}
	fmt.Printf("UTC: %s\nLocal: %s\n", t.UTC(), t)
}
