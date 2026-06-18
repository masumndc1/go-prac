package main

import "fmt"

func processData(i any) {
	// The syntax 'i.(type)' extracts the underlying type of the interface
	// 'v' captures the actual value casted to that concrete type
	switch v := i.(type) {
	case int:
		fmt.Printf("Integer discovered! Multiplied by 2: %d\n", v*2)
	case string:
		fmt.Printf("String discovered! Length: %d, Value: %q\n", len(v), v)
	case bool:
		fmt.Printf("Boolean discovered! Direct value: %t\n", v)
	case []int:
		fmt.Printf("Slice of ints discovered! Item count: %d\n", len(v))
	default:
		fmt.Printf("Unknown or unhandled type: %T\n", v)
	}
}

func identifyType(i any) {
	switch v := i.(type) {
	case int:
		fmt.Printf("It's an integer: %d\n", v)
	case string:
		fmt.Printf("It's a string: %s\n", v)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	operatingSystem := "linux"
	age := 25
	value := 1

	switch operatingSystem {
	case "windows":
		fmt.Println("Windows OS")
	case "linux":
		fmt.Println("Linux OS") // Output: Linux OS
	case "macOS":
		fmt.Println("Mac OS")
	default:
		fmt.Println("Unknown OS") // Runs if no match is found
	}

	switch day := "Saturday"; day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's a weekday.")
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!") // Output: It's the weekend!
	}

	switch {
	case age < 18:
		fmt.Println("Minor")
	case age >= 18 && age < 65:
		fmt.Println("Adult") // Output: Adult
	default:
		fmt.Println("Senior")
	}

	switch value {
	case 1:
		fmt.Println("Case 1 matched.") // Prints first
		fallthrough
	case 2:
		fmt.Println("Case 2 also executed via fallthrough.") // Prints second
	case 3:
		fmt.Println("This will NOT print.")
	}

	// Pass different data types into the exact same function
	processData(42)          // Matches 'int'
	processData("Hello Go")  // Matches 'string'
	processData(true)        // Matches 'bool'
	processData([]int{1, 2}) // Matches '[]int'
	processData(3.14159)     // Matches 'default' (float64)

}
