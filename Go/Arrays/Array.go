package main

/*func main() {
	var array [5]int

	for i := 0; i < 5; i++ {
		fmt.Printf("Enter number %d: ", i)
		fmt.Scanln(&array[i])

		if i == 4 {
			fmt.Println()
			for i := 0; i < 5; i++ {
				fmt.Printf("%d\n", array[i])
			}
		}
	}
}*/

//Problema 02
/*func main() {
	array := [5]int{1, 2, 3, 4, 5}
	j := 4

	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", array[j])
		j--
	}
}*/

// Problema 03
/*func main() {
	array := [5]int{1, 2, 3, 4, 5}
	sum := 0

	for i := 0; i < 5; i++ {
		sum += array[i]
	}

	fmt.Printf("Average value: %.2f", float64(sum/5))
}
*/

// Problema 05
/*func main() {
	arrayOne := [5]int{1, 2, 3, 4, 5}
	arrayTwo := [5]int{6, 7, 8, 9, 10}
	var number int
	var sum int

	fmt.Print("Enter a number from 1-5: ")
	fmt.Scanln(&number)

	for i := 0; i < 5; i++ {
		if number == arrayOne[i] {
			sum += number
		}
	}

	fmt.Print("Enter a number from 6-10: ")
	fmt.Scanln(&number)

	for i := 0; i < 5; i++ {
		if number == arrayTwo[i] {
			sum += number
		}
	}

	fmt.Printf("Sum: %d", sum)
}
*/

//