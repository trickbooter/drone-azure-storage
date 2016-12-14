package main

import (
	"github.com/drone/drone-plugin-go/plugin"
	"path/filepath"
	"reflect"
	"strconv"
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
	src := filepath.Join(w.Path, vargs.Source)
	seg := strconv.Itoa(strings.Count(s, "/"))
	if !reflect.DeepEqual(command(vargs, w).Args, []string{
		"blobxfer",
		"--strip-components",
		seg,
		"my-storage-account",
		"my-container",
		src,
	}) {
		t.Error("command not composed correctly")
	}
}
