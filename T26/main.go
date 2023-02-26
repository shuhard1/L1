package main

import "fmt"

func uniq(s string) bool {
	var uniqChars []rune
	uniqChars = append(uniqChars, rune(s[0]))
	var res bool = true
	for i, r := range s {
		for _, c := range uniqChars {
			if r == c && i != 0 {
				res = false
			} else if i != 0 {
				uniqChars = append(uniqChars, rune(s[i]))
			}
		}
	}
	return res
}

func main() {
	fmt.Println(uniq("abcd"))
	fmt.Println(uniq("abCdefAaf"))
	fmt.Println(uniq("aabcd"))
}
