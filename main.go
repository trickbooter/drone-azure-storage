package main

import (
	"fmt"
	"github.com/drone/drone-plugin-go/plugin"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	buildCommit string
)

type AzureBlobxfer struct {
	StorageAccountKey  string `json:"account_key"`
	StorageAccountName string `json:"storage_account"`
	Container          string `json:"container"`
	Source             string `json:"source"`
}

func main() {
	fmt.Printf("Drone Azure Storage Plugin built from %s\n", buildCommit)

	workspace := plugin.Workspace{}
	vargs := AzureBlobxfer{}

	plugin.Param("workspace", &workspace)
	plugin.Param("vargs", &vargs)
	plugin.MustParse()

	if len(vargs.StorageAccountKey) == 0 {
		fmt.Println("storage_account must be defined in your .drone.yml")
		return
	}

	if len(vargs.Container) == 0 {
		fmt.Println("container must be defined in your .drome.yml")
		return
	}

	cmd := command(vargs, workspace)
	trace(cmd)

	// Append StorageAccountKey to the cmd after trace to avoid exposing the key
	cmd.Args = append(cmd.Args, "--storageaccountkey", vargs.StorageAccountKey)
	cmd.Env = os.Environ()
	cmd.Dir = workspace.Path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Failed to upload %s to %s/%s: %s", vargs.Source, vargs.StorageAccountName, vargs.Container, err)
		os.Exit(1)
	}
}

func command(s AzureBlobxfer, w plugin.Workspace) *exec.Cmd {

	source := filepath.Join(w.Path, s.Source)
	segments := strconv.Itoa(stripCount(source))

	args := []string{
		"--strip-components",
		segments,
		s.StorageAccountName,
		s.Container,
		source,
	}
	return exec.Command("blobxfer", args...)
}

func stripCount(path string) int {
	c := strings.Count(filepath.Clean(path), "/")
	// if we have a valid path that is a file and has more than 0 leading segments
	// subtract one segment to account for the file
	if finfo, err := os.Stat(path); err == nil && !finfo.IsDir() && c > 0 {
		return c - 1
	} else {
		return c
	}
}

// trace writes each command to standard error (preceded by a ‘$ ’) before it
// is executed. Used for debugging your build.
func trace(cmd *exec.Cmd) {
	fmt.Println("$", strings.Join(cmd.Args, " "))
}
