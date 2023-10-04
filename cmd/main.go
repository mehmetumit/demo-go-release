package main

import (
	"demo/pkg/version"
	"fmt"
)

func DemoAdd(firstNum int, secondNum int)int{
    return firstNum + secondNum
}
func DemoSubtract(firstNum int, secondNum int)int{
    return firstNum - secondNum
}
func main(){
    fmt.Printf("Version: %s \nCommit: %s\n", version.Version, version.GitCommit)
}
