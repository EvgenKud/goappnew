//go:build mage
// +build mage

package main

import (
	"fmt"
	"github.com/aevea/git/v2"
	"github.com/magefile/mage/sh"
	"github.com/mlctrez/goappnew/goapp/global"
	"os/exec"
	"strings"
)

const (
	entryPath     = "goapp/service/main/main.go"
	module        = "github.com/mlctrez/goappnew"
	wasmBinPath   = "goapp/web/app.wasm"
	serverBinName = "goappnew"
	serverBinPath = "temp/goappnew"
)

// Wasm build wasm binary.
func Wasm() error {
	fmt.Println("Building wasm...")
	version, commit := getBuildData()

	return sh.RunWithV(map[string]string{"GOARCH": "wasm", "GOOS": "js"},
		"go", "build", "-o", wasmBinPath,
		fmt.Sprintf("-ldflags=-X %s.goapp.Version=%s -X %s.goapp.Commit=%s",
			module, version, module, commit,
		),
		entryPath,
	)
}

// Server build server binary.
func Server() error {
	fmt.Println("Building server...")
	version, commit := getBuildData()

	return sh.RunV("go", "build", "-o", serverBinPath,
		fmt.Sprintf("-ldflags=-X %s.goapp.Version=%s -X %s.goapp.Commit=%s",
			module, version, module, commit,
		),
		entryPath,
	)
}

func Css() error {
	fmt.Println("Generate CSS...")

	return sh.RunV("go", "run", "cmd/css/main.go")
}

func Run() error {
	return sh.RunWith(map[string]string{"DEV": "1"}, serverBinPath)
}

func Clean() error {
	return sh.RunV("rm", "-Rf", "temp", wasmBinPath, global.AppCss)
}

func Kill() error {
	return sh.RunV("pkill", serverBinName)
}

func Build() error {
	if err := Clean(); err != nil {
		return err
	}
	if err := Wasm(); err != nil {
		return err
	}
	if err := Css(); err != nil {
		return err
	}

	return Server()
}

func Update() error {
	Clean()
	Build()
	Kill()
	cmd := exec.Command(serverBinPath)

	return cmd.Start()
}

func getBuildData() (string, string) {
	gt, err := git.OpenGit("./", nil)
	if err != nil {
		panic(err)
	}

	commit, err := gt.CurrentCommit()
	if err != nil {
		panic(err)
	}
	lCommit := commit.Hash.String()

	tg, err := gt.CurrentTag()
	if err != nil {
		panic(err)
	}
	tag := strings.ReplaceAll(tg.Name, "refs/tags/", "")

	fmt.Println("Build version:", tag, " commit:", lCommit)

	return tag, lCommit
}
