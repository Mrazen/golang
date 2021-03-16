package main

import "fmt"

type Battery struct {
	levelCharge string
}

func (b Battery) String() string {
	return b.levelCharge + "123"
}
func printLevel(s fmt.Stringer) {
	fmt.Println(s.String())
}

func main() {
	var text string
	fmt.Scan(&text)
	batteryForTest := Battery{
		levelCharge: text,
	}
	printLevel(batteryForTest)

}
