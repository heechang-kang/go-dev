// print, Variables

package main

import "fmt" //formatter, print things

func main() {
	fmt.Println("testing")

	///// Variable /////
	var name string = "heechangkang"
	fmt.Println(name)
	fmt.Printf("\n%T\n", name)
	const bornYear = 1997
	bornMonth, bornDate := 2, 24
	//bornYear = 1998
	fmt.Println(bornYear, bornMonth, bornDate)

	///// MAP /////
	emails := make(map[string]string)
	// emails := map[string]string{"a":"a@mail.com", "b":"b@mail.com"}
	emails["a"] = "a@mail.com"
	emails["b"] = "b@mail.com"

	fmt.Println(emails)
	delete(emails, "b")
	fmt.Println(emails)

	///// RANGE /////
	fmt.Println("\n\n")
	ids := []int{1, 2, 3, 4, 5, 6}

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("sum: ", sum)
	//

}
