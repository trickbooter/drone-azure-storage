package main

import (
	"github.com/drone/drone-plugin-go/plugin"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestCommandBuildCorrectly(t *testing.T) {
	vargs := AzureBlobxfer{}
	vargs.StorageAccountKey = "xyzabc"
	vargs.Container = "my-container"
	vargs.StorageAccountName = "my-storage-account"
	vargs.Source = "__source__"
	w := plugin.Workspace{Path: "/test/path"}
	s := filepath.Join(w.Path, vargs.Source)
	if !reflect.DeepEqual(command(vargs, w).Args, []string{
		"blobxfer",
		"--strip-components",
		strings.Count(s, "/"),
		"my-storage-account",
		"my-container",
		s,
	}) {
		t.Error("command not composed correctly")
	}
}
