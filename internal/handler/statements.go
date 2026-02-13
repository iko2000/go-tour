package examples


import fmt "fmt"


func ExperimentStatements() {
	fmt.Println("Hello, World!");
	num := 5;


	val := checker(num);
	fmt.Println(val)
		for i := 0; i < 10; i++{
		fmt.Println(i);
	}

	lastNames := []string{"BUSH", "TRUMP", "OBAMA"};

	for index, name := range lastNames {
		fmt.Printf("%d: %s", index, name);
	}


}


// If Statements.


func checker (num int) bool {

	if num == 5{
		return true;
	} else {
		return false;
	}

}


// For Loops.


// Swtich Statements.