package main

import (
	"os"

	"github.com/go-git/go-git/v5"
)

func main() {

	r, err := git.PlainClone(".", false, &git.CloneOptions{
		URL:      "https://github.com/semyon-dev/stepik-go/tree/master/work_with_files/task_csv_1",
		Progress: os.Stdout,
	})
}
