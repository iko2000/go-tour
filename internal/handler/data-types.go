package examples

import fmt "fmt"
import strings "strings"


func DataTypes() {

	var nm int = 5;
	fmt.Println(nm);

	name := "Iviko";

	fmt.Println(strings.ToUpper(name));


	var nums [5]int;

	fmt.Println(nums);

	names := []string{"Iviko"};

	fmt.Println(names);

	x := 35;

	p := &x;

	fmt.Println(*p);



}