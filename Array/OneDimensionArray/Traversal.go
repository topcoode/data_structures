package array

import "fmt"

func Traversal_array() {
	//Printing index / Traversal
	myArray := []string{"MySQL", "MongoDB", "Oracle", "Elasticsearch", "SQLite"}

	for index, element := range myArray {
		fmt.Println("Index:", index, "Element:", element)
	}

}
