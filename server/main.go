package main

import (
	"github.com/NYTimes/marvin"
	hello "github.com/xxy/helloworld"
	"google.golang.org/appengine"
)

func main() {
	marvin.Init(hello.NewStatsService())
	appengine.Main()
}
