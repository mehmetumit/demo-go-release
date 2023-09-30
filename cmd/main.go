package main

import (
	"demo/pkg/version"
	"fmt"
)

func main(){
    fmt.Printf("Version: %s \nCommit: %s\n", version.Version, version.GitCommit)
    fmt.Println("Demo")
}
