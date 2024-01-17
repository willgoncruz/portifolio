package main 

import (
	"context"
	"log"
	"os"

  "portifolio/pages"
)

func main() {
  args := os.Args
  if len(args) != 2 {
    log.Fatal("wrong usage of script, must receive path to file.\n Example: go run static.go index.html")
    return
  }

  path := args[1]

	f, err := os.Create(path)
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	  return
  }

	err = pages.Portifolio("Will").Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}
}
