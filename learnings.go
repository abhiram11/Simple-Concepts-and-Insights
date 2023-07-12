package main

import "fmt"

func main() {
	c := make(chan int, 3)
	go func() {
		for i := 0; i < 3; i++ {
			c <- i
			fmt.Printf("number %v inserted into channel\n", i)
		}
		fmt.Println("Leng: ", len(c))

	}()

	for i := 0; i < 3; i++ {
		fmt.Printf("Value popped from channel: %v \n", <-c)
	}

}

// func main() {
// 	msg := make(chan string, 2)
// 	msg2 := make(chan string, 1)

// 	go func() {
// 		msg2 <- "poof2"
// 		fmt.Println("val func 2: ", <-msg2)
// 	}()

// 	go func() {
// 		msg <- "poof"
// 		fmt.Println("val1: ", <-msg) //<-msg, <-msg)

// 		time.Sleep(time.Second)
// 		msg <- "spoof"
// 		fmt.Println("val2: ", <-msg) //<-msg, <-msg)
// 	}()

// 	time.Sleep(time.Second * 2)

// 	//, <-msg)

// }

// func f(from string) {
// 	for i := 0; i < 2; i++ {
// 		fmt.Println(from, " : ", i)
// 	}
// }

// func main() {
// 	f("direct") // running directly, without goroutines

// 	go f("goroutiness") // just typing this won't wait for this function to finish, add timer or waitgroup.

// 	go f("dict")
// 	go f("map")
// 	go f("trap")

// 	time.Sleep(time.Second)
// 	fmt.Print("Done!!")
// }

// func main() {
// 	// fmt.Println(add(10, 32))
// 	first, second := add(10, 22) // even if "First" is not required by you, you still have to allot it a variable.
// 	fmt.Println(first, second)

// }

// func add(a int, b int) (int, int) {
// 	return (a + b), (a - b) // returning multiple values..
// }

// func main() {
// 	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
// 	sum := 0

// 	for _, num := range nums {
// 		fmt.Println(num)
// 		sum += num
// 	}

// 	fmt.Println("sum : ", sum)
// }

// func main() {
// 	// learning Zero Values
// 	m := make(map[string]bool)
// 	m["k1"] = true
// 	m["k2"] = false
// 	m["4"] = true
// 	fmt.Println(m)

// 	fmt.Println(m["1"])

// }

// func main() {
// 	a := []int{1, 2, 3, 4, 5}
// 	// fmt.Printf("%T", a)
// 	b := append(a, 36)
// 	// fmt.Println("a: ", a)
// 	// fmt.Println("b: ", b)
// 	fmt.Println(a[2:5])
// 	fmt.Println(b[:3])

// }

// func main() {

// 	switch time.Now().Weekday() {
// 	case time.Saturday, time.Sunday:
// 		fmt.Println("It's the weekend")
// 	default:
// 		fmt.Println("It's a weekday")
// 	}
// }

// func main() {
// 	i := 0
// 	for {
// 		fmt.Println("i: ", i)
// 		i += 1
// 		if i == 2 {
// 			break
// 		}
// 	}

// 	for j := 0; j <= 3; j++ {
// 		fmt.Println("j: ", j)
// 	}

// }

// func main() {
// 	var a = "Initial"

// 	var b, c = 1, 2

// 	fmt.Printf("%T %T %T", a, b, c)

// }
