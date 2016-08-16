/*
Project:Learn_go
Time:2016/8/16-17:57
author:Dio
*/
package main

import "fmt"

func main() {
	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0
	for key, value := range map1 {
		fmt.Printf("key is: %d - value is: %f\n", key, value)
	}
	/* out
	key is: 1 - value is: 1.000000
	key is: 2 - value is: 2.000000
	key is: 3 - value is: 3.000000
	key is: 4 - value is: 4.000000
	*/

	/*
		8.1
	*/
	capitals := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}
	for key := range capitals {
		fmt.Println("Map item: Capital of", key, "is", capitals[key])
	}
	/** OUT
	Map item: Capital of Italy is Rome
	Map item: Capital of Japan is Tokyo
	Map item: Capital of France is Paris
	*/
}
