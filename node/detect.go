package node

import (
	"os"
	"path/filepath"

	"github.com/paketo-buildpacks/packit"
)

func Detect() packit.DetectFunc {
	return func(dc packit.DetectContext) (packit.DetectResult, error) {
		if err := shouldDetect(dc.WorkingDir); err != nil {
			return packit.DetectResult{}, err
		}
		return packit.DetectResult{Plan: packit.BuildPlan{
			Provides: []packit.BuildPlanProvision{{Name: "node"}},
			Requires: []packit.BuildPlanRequirement{{Name: "node"}},
		}}, nil
	}
}

func shouldDetect(workingDir string) error {
	path := filepath.Join(workingDir, "app.js")
	_, err := os.Stat(path)
	return err
}
