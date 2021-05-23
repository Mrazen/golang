package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

//02-01-2006 15:04:05
//2020-05-15 08:00:00
func main() {
	const now = 1589570165

	timeStamp := time.Unix(now, 0).UTC()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()

	elapsedTime := strings.Replace(s, "мин.", "m", 1)
	elapsedTime = strings.Replace(elapsedTime, "сек.", "s", 1)
	elapsedTime = strings.Replace(elapsedTime, " ", "", -1)

	dur, err := time.ParseDuration(elapsedTime)
	if err != nil {
		panic(err)
	}
	//dur.Round(time.Hour).Hours()
	unitedDate := timeStamp.Add(dur)

	fmt.Println(unitedDate.Format(time.UnixDate))

}
