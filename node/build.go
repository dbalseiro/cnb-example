package node

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/paketo-buildpacks/packit"
)

func Build() packit.BuildFunc {
	return func(ctx packit.BuildContext) (packit.BuildResult, error) {
		file, err := os.Open(filepath.Join(ctx.CNBPath, "buildpack.toml"))
		if err != nil {
			return packit.BuildResult{}, err
		}

		apiKey := os.Getenv("API_KEY")
		fmt.Printf("API KEY = %s\n", apiKey)

		var m struct {
			Metadata struct {
				Dependencies []struct {
					URI string `toml:"uri"`
				} `toml:"dependencies"`
			} `toml:"metadata"`
		}

		_, err = toml.DecodeReader(file, &m)
		if err != nil {
			return packit.BuildResult{}, err
		}

		uri := m.Metadata.Dependencies[0].URI
		fmt.Printf("URI -> %s\n", uri)

		nodeLayer, err := ctx.Layers.Get("node")
		if err != nil {
			return packit.BuildResult{}, err
		}

		nodeLayer, err = nodeLayer.Reset()
		if err != nil {
			return packit.BuildResult{}, err
		}

		nodeLayer.Launch = true

		dldir, err := os.MkdirTemp("", "dldir")
		if err != nil {
			return packit.BuildResult{}, err
		}
		defer os.RemoveAll(dldir)

		fmt.Println("Downloading dependency . . .")
		tarfile := filepath.Join(dldir, "node.tar.xz")
		err = exec.Command("curl", uri, "-o", tarfile).Run()
		if err != nil {
			return packit.BuildResult{}, err
		}

		fmt.Println("Untaring dependency . . .")
		err = exec.Command("tar", "-xf", tarfile, "--strip-components=1", "-C", nodeLayer.Path).Run()
		if err != nil {
			return packit.BuildResult{}, err
		}

		return packit.BuildResult{
			Layers: []packit.Layer{nodeLayer},
		}, nil
	}
}
