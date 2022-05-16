package main

import (
	"flag"
	"fmt"
	"os"
)

const horse = `
      ./|,,/|
     <   o o)
    <\ (    |
   <\\  |\  |
  <\\\  |(__)    %s
 <\\\\  |
`

func usage() {
	help := `
Usage: horsesay <message>

Causes a horse to say the specified message.`
	fmt.Fprintln(flag.CommandLine.Output(), help)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Fprintln(flag.CommandLine.Output(), "Gotta have something to say.")
		usage()
		return
	}
	if len(os.Args) > 2 {
		fmt.Fprintln(flag.CommandLine.Output(), "Too many arguments.")
		usage()
		return
	}

	fmt.Printf(horse, os.Args[1])
}
