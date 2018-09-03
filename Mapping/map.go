package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	accountBalances := make(map[string]float32)

	accountBalances["0x1"] = 10
	accountBalances["0x2"] = 20
	accountBalances["0x3"] = 30
	accountBalances["0x4"] = 40
	/* accessing the map using key, Deleting a Value from Key
	fmt.Println(accountBalances)

	Balance_0x1 := accountBalances["0x1"]
	fmt.Println(Balance_0x1)

	delete(accountBalances, "0x1")

	fmt.Println(accountBalances)

	Balance_0x1 = accountBalances["0x1"]
	fmt.Println(Balance_0x1)
	*/

	/* Ittrating on a map, this is quite userful when writting a contract over ethereum */
	i := 0
	for k, v := range accountBalances {

		if k == "0x2" {
			fmt.Println("Index :", i, "-->", k, ":", v)
		}

		i++
		fmt.Println("Index :", i)
	}

}
