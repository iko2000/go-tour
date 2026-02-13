package examples	

import (
	"fmt"
	"strings"
)

func TypeHandler() {
	fmt.Println("Hello, World!");

	name := "James";
	age := 25;

	fmt.Printf("My name is %s and I am %d years old.\n", name, age);

	fmt.Println("Hello, " + name + "!");


	uppername := basicFunction(name);
	fmt.Printf("My name in uppercase is %s.\n", uppername);
	result, err := advancedFunction(age);
	if err != nil {
		fmt.Println("Error:", err);
	} else {
		fmt.Println(result);
	}
}


func basicFunction(name string) string {
  return strings.ToUpper(name);
}

func advancedFunction(id int) (string, error) {
  
	 if(id < 0) {
		return "", fmt.Errorf("ID cannot be negative");
	 }
	 return fmt.Sprintf("Your ID is %d", id), nil;
}