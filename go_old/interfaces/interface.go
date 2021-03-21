// package inter

// import "fmt"

// type Battery struct {
// 	levelCharge string
// }

// func (b Battery) String() string {
// 	level := ""
// 	for _, el := range b.levelCharge {
// 		//fmt.Printf("%T", el-'0')
// 		if el-'0' == 0 {
// 			level += " "
// 		}
// 	}
// 	for _, el := range b.levelCharge {
// 		if el-'0' == 1 {
// 			level += "X"
// 		}
// 	}
// 	//fmt.Print(level)
// 	return "[" + level + "]"
// }
// func printLevel(s fmt.Stringer) {
// 	fmt.Println(s.String())
// }

// //
// func main() {
// 	var text string
// 	fmt.Scan(&text)
// 	batteryForTest := Battery{
// 		levelCharge: text,
// 	}
// 	printLevel(batteryForTest)

// }
