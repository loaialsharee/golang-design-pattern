package main

func main() {
	s := GetPrinter()
	for i := 0; i < 5; i++ {
		s.PrintSentence("This sentence is printed using a singleton!")
	}
}
