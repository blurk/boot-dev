package main

func fizzbuzz() {
	for i := 3; i < 101; i++ {
		if i%15 == 0 {
			println("fizzbuzz")
		} else if i%3 == 0 {
			println("fizz")
		} else if i%5 == 0 {
			println("buzz")
		} else {
			println(i)
		}
	}
}

// don't touch below this line

func main() {
	fizzbuzz()
}
