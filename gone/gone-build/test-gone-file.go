
package main



func twoWords(firstWord string) ( string, string ) {
    return firstWord, " World!"
}

func getGreeting() string {
    var firstWord string = "Hello"
    return firstWord
}

func lambdaTester(firstWord: string, secondWord: string) string {
    concat := func(x: string, y: string) : string {return x + y}
    result := concat(firstWord, secondWord)
    return result
}
func main() {
    
    exclaimedGreeting := getGreeting()
    noSpaces:=exclaimedGreeting
    println(noSpaces)

}
