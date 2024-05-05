package main

import "fmt"

type Test02 struct {
	TestCase string
	Value    string
}

func main() {
	l := "L"
	r := "R"
	draw := "="

	list := []string{"210122", "000210", "221012", "012001"}
	result := []Test02{}
	for _, item := range list {
		itemCase := Test02{}
		newStr := ""
		for i := 0; i < len(item); i++ {
			if i == len(item)-1 {
				break
			} else {
				if item[i] > item[i+1] {
					newStr += l
				} else if item[i] < item[i+1] {
					newStr += r
				} else {
					newStr += draw
				}
			}
		}

		itemCase.TestCase = item
		itemCase.Value = newStr
		result = append(result, itemCase)

	}

	fmt.Printf("%+v", result)

}
