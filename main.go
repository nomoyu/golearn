package main

import "github.com/nomoyu/golearn/cmd"

func main() {
	defer cmd.Clean()
	cmd.Start()
}
