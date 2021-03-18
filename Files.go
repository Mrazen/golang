package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	var el, res int64

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		el, _ = strconv.ParseInt(scanner.Text(), 10, 64)
		res += el
	}
	w := bufio.NewWriter(os.Stdout)
	n, err := w.WriteString(strconv.FormatInt(res, 10))
	if err != nil && err != io.EOF {
		fmt.Print(n)
	}
}
