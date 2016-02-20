package main

import "fmt"

func perm(s string) []string {
	perms := []string{}
	if len(s) == 1 {
		return []string{s}
	} else {
		for i := 0; i < len(s); i++ {
			for _, perm := range perm(s[:i] + s[i+1:]) {
				perms = append(perms, string(s[i]) + perm)
			}
		}
		return perms
	}
}

func main() {
	fmt.Println(perm("a"))
	fmt.Println(perm("ab"))
	fmt.Println(perm("abc"))
	fmt.Println(perm("aac"))
}
