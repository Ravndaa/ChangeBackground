package main

import (
	"fmt"
	"os"
	"ravndal/SetPaper/cmd"
)

func main() {

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

}

/***

All magic should be credited: https://www.reddit.com/r/golang/comments/55guj9/need_help_with_my_code/



****/
