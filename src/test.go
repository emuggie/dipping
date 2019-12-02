package main

import (
	log_impl "github.com/emuggie/dipping/package/log/impl"

	str_impl "github.com/emuggie/dipping/package/strings/impl"

	"github.com/emuggie/dipping/package/log"

	"github.com/emuggie/dipping/plugin"
)

func main() {
	log_impl.Delegate()
	str_impl.Delegate()
	d := log.Import()
	d.Println(plugin.Open("/home/emuggie/dev/workspace/VisualStudioCode/golang/golang-dipping/src/test/test.so"))
}
