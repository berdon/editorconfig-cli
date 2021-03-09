package main

import (
	"os"

	"github.com/amyboyd/editorconfig-cli/editorconfig"
)

func main() {
	app := editorconfig.CreateCliApp()

	app.Run(os.Args)
}
