package main

import (
	"fmt"
	"strings"
)

// Complete the containsAll function. It takes a string s, and a variadic list of strings substrings.
// It returns true if all of the substrings are present somewhere in the original string.
func containsAll(s string, substrings ...string) bool {
	for _, str := range substrings {
		if strings.Contains(s, str) {
			continue
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(containsAll("hello world", "hello"))
	fmt.Println(containsAll("hello world", "hello", " worlds"))
	fmt.Println(containsAll("lorem", "ipsum"))
	fmt.Println(containsAll("lorem ipsum dolor", "ipsum", "dolor"))
	fmt.Println(containsAll("Boot.dev uses Golang for its backend", "for", "the", "backend"))
	fmt.Println(containsAll("Boot.dev uses Golang for its backend", "Golang"))
	fmt.Println(containsAll("Did you ever hear the Tragedy of Darth Plagueis the wise? I thought not. It's not a story the Jedi would tell you. It's a Sith legend. Darth Plagueis was a Dark Lord of the Sith, so powerful and so wise he could use the Force to influence the midichlorians to create life... He had such a knowledge of the dark side that he could even keep the ones he cared about from dying. The dark side of the Force is a pathway to many abilities some consider to be unnatural. He became so powerful... the only thing he was afraid of was losing his power, which eventually, of course, he did. Unfortunately, he taught his apprentice everything he knew, then his apprentice killed him in his sleep. It's ironic he could save others from death, but not himself", "ironic", "save"))
	fmt.Println(containsAll("Did you ever hear the Tragedy of Darth Plagueis the wise? I thought not. It's not a story the Jedi would tell you. It's a Sith legend. Darth Plagueis was a Dark Lord of the Sith, so powerful and so wise he could use the Force to influence the midichlorians to create life... He had such a knowledge of the dark side that he could even keep the ones he cared about from dying. The dark side of the Force is a pathway to many abilities some consider to be unnatural. He became so powerful... the only thing he was afraid of was losing his power, which eventually, of course, he did. Unfortunately, he taught his apprentice everything he knew, then his apprentice killed him in his sleep. It's ironic he could save others from death, but not himself", "foobar"))
}
