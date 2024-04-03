package main

func twoWords(firstWord string) (string, string) {
	return firstWord, " World!"
}

func getGreeting() string {
	var firstWord string = "Hello"
	return firstWord
}

func lambdaTester(firstWord string, secondWord string) string {
	concat := func(x string, y string) string { return x + y }
	result := concat(firstWord, secondWord)
	return result
}

func questionMarkTester(codeWord string) string {
	if codeWord == "hello" {
		return "codeWord"
	}
	return "nil"
}

func questionMark(value interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return value
}
func main() {

	exclaimedGreeting := getGreeting()
	noSpaces := exclaimedGreeting
	a := questionMarkTester("hello")

	println(noSpaces, a)
}
