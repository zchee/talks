package main

import (
	"fmt"

	"github.com/neovim/go-client/nvim/plugin"
)

func hello() error {
	msg := "Hello go-client!!!!!"
	var p plugin.Plugin
	return p.Nvim.Command(fmt.Sprintf("echo %s", msg))
}

func main() {
	plugin.Main(func(p *plugin.Plugin) error {
		p.HandleFunction(&plugin.FunctionOptions{Name: "Hello"}, hello)
		return nil
	})
}
