package main

import "fmt"

func uniq(s string) bool {
	//сюда будут записываться уникальные символы
	var uniqChars []rune
	uniqChars = append(uniqChars, rune(s[0]))

	//пробегаемся по строке s
	for i, r := range s {
		//пробегаемся по слайсу uniqChars
		for _, c := range uniqChars {
			//проверяем данный символ с каждым символом из uniqChars
			if r == c && i != 0 {
				//если есть совпадение, то возращаем false
				return false
			} else if i != 0 {
				//если совпадений нет, то добавляем уникальный символ в uniqChars
				uniqChars = append(uniqChars, rune(s[i]))
			}
		}
	}
	return true
}

func main() {
	fmt.Println(uniq("abcd"))
	fmt.Println(uniq("abCdefAaf"))
	fmt.Println(uniq("aabcd"))
}
