package main

import "fmt"

func permute(res *[]string, output, input string, count *int) {
	if len(input) == 0 {
		*res = append(*res, output)
	} else {
		for i := range input {
			(*count)++
			permute(res, output+string(input[i]), input[:i]+input[i+1:], count)
		}
	}
}

func getPermutations(arg string) []string {
	res := []string{}
	count := 0
	permute(&res, "", arg, &count)
	fmt.Println(count)
	return res
}

func main() {
	fmt.Println(getPermutations("abc"))
}
