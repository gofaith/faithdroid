package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/StevenZack/tools/fileToolkit"

	"github.com/StevenZack/tools/strToolkit"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("not enough args")
		fmt.Println("examples:")
		fmt.Println("1. faithdroid build_386")
		fmt.Println("2. faithdroid build_arm64")
		fmt.Println("3. faithdroid build_arm")
		fmt.Println("4. faithdroid build_arm*")
		fmt.Println("\n5. faithdroid new")
		fmt.Println("6. faithdroid update")
		return
	}
	gopath := fileToolkit.Getrpath(os.Getenv("GOPATH")) + "src"
	w, e := os.Getwd()
	if e != nil {
		fmt.Println("somewhere got error:", e)
		return
	}
	if !strToolkit.StartsWith(w, gopath) {
		fmt.Println("current path is not in GOPATH")
		return
	}
	curPath := w[len(fileToolkit.Getrpath(gopath)):]
	var c *exec.Cmd
	switch args[0] {
	case "build_386":
		c = exec.Command("gomobile", "bind", "--target=android/386", "-o", "example-android-project/faithdroid/faithdroid.aar", curPath)
	case "build_arm64":
		c = exec.Command("gomobile", "bind", "--target=android/arm64", "-o", "example-android-project/faithdroid/faithdroid.aar", curPath)
	case "build_arm":
		c = exec.Command("gomobile", "bind", "--target=android/arm", "-o", "example-android-project/faithdroid/faithdroid.aar", curPath)
	case "build_arm*":
		c = exec.Command("gomobile", "bind", "--target=android/arm,android/arm64", "-o", "example-android-project/faithdroid/faithdroid.aar", curPath)
	case "new":
		c = exec.Command("copy", "-r", fileToolkit.Getrpath(gopath)+"github.com/gofaith/faithdroid", ".")
	case "update":
		c = exec.Command("copy", "-r", fileToolkit.Getrpath(gopath)+"github.com/gofaith/faithdroid/*", ".")
	}
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	e = c.Run()
	if e != nil {
		fmt.Println(`command.Run() error :`, e)
		return
	}
}
