package main

import (
	"fmt"
	"github.com/LightAlykard/TempExch/pkg/version"
)

func main() {
	fmt.Printf("version: %s\n", version.Version.Version)
	fmt.Printf("Commit: %s\n", version.Version.Commit)
	fmt.Printf("BuildTime: %s\n", version.Version.BuildTime)
}
