package main

import (
	"github.com/NYTimes/marvin"
	hello "github.com/xxy/go-test"
	"google.golang.org/appengine"
)

func main() {
	marvin.Init(hello.NewTestService())
	// marvin.RouterSelect("httprouter")
	appengine.Main()
}
