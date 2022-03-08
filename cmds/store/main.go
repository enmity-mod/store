package main

import (
	"flag"

	"enmity.app/store/internals/store"
)

var (
	name string
	dir  string
)

func init() {
	flag.StringVar(&name, "name", "Your plugin repo", "Your repo's name.")
	flag.StringVar(&dir, "dir", "", "Your repository's folder.")
	flag.Parse()
}

func main() {
	store.GenerateStore(&name, &dir)
}
