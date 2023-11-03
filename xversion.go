package xversion

import "fmt"

var GitCommit string

func init() {
	fmt.Println("版本:", GitCommit)
}
