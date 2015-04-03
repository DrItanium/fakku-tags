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
		return
	}
	for i := 0; i < len(tags); i++ {
		curr := tags[i]
		if verboseFlag {
			url, err0 := curr.Url()
			if err0 != nil {
				fmt.Println(err0)
				return
			}
			fmt.Printf("%s - %s - %s\n", curr.Name, curr.Description, url)
		} else {
			fmt.Println(curr.Name)
		}
	}
}
