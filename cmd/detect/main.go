package main

import (
	"github.com/dbalseiro/cnb-example/node"
	"github.com/paketo-buildpacks/packit"
)

func main() {
	packit.Detect(node.Detect())
}
