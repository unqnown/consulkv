package main

import (
	"github.com/unqnown/consulkv/app"
	"github.com/unqnown/consulkv/pkg/check"
)

func main() {
	check.Fatal(app.Run())
}
