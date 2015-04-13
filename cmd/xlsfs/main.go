package main

import (
	"flag"
	"fmt"
	"os"
)

const usageMsg = `
Xlsfs starts 9P2000 file server
for the specified <filename.xlsx> document.

The root of filesystem consists of dirs with spreadsheet names.
Each spreadsheet represented as a dir.
The following structure depends on [TBD].
`

func main() {
	addr := flag.String("addr", "localhost:5640", "service listen address")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <filename.xlsx>", os.Args[0])
		fmt.Fprintf(os.Stderr, usageMsg)
		fmt.Fprintf(os.Stderr, "\nOptions:\n")
		flag.PrintDefaults()
		os.Exit(2)
	}
	flag.Parse()

	// [TBD]
	fmt.Println("Addr has value:", *addr)

	os.Exit(0)
}
