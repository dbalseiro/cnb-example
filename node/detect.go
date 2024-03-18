package node

import (
	"fmt"

	"github.com/paketo-buildpacks/packit"
)

func Detect() packit.DetectFunc {
	return func(dc packit.DetectContext) (packit.DetectResult, error) {
		return packit.DetectResult{}, fmt.Errorf("fail lol")
	}
}
