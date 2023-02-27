package main

func main() {
	file := bar()
	defer file.Close()
}
