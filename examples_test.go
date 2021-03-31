package gomodexample2_test

import (
	"embed"
	"fmt"
	"io/fs"
	
	"github.com/parro-it/gomodexample2"
)

//go:embed fixtures
var fixtureRootFS embed.FS
var fixtureFS, _ = fs.Sub(fixtureRootFS, "fixtures")

// This example show how to use gomodexample2.Func()
func ExampleFunc() {
	fmt.Println(gomodexample2.Func())
	// Output: 42
}
