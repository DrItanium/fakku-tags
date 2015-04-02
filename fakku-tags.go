// list all known tags
package main

import (
	"flag"
	"fmt"
	"github.com/DrItanium/fakku"
)

var verboseFlag bool

func init() {
	const usage = "print descriptions of the tags"
	flag.BoolVar(&verboseFlag, "verbose", false, usage)
	flag.BoolVar(&verboseFlag, "v", false, usage+" (shorthand)")
}
func main() {
	flag.Parse()
	tags, err := fakku.Tags()

	if err != nil {
		fmt.Println(err)
		fmt.Println("Something bad happened! Perhaps Fakku is down?")
	}
	for i := 0; i < len(tags); i++ {
		curr := tags[i]
		if verboseFlag {
			fmt.Printf("%s - %s - %s\n", curr.Name, curr.Description, curr.Url)
		} else {
			fmt.Println(curr.Name)
		}
	}
}
