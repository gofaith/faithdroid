package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

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
		e = os.Mkdir("faithdroid", 0755)
		if e != nil {
			fmt.Println(`os.Mkdir ./faithdroid/ error :`, e)
			return
		}
		fs, e := listDir(fileToolkit.Getrpath(gopath) + "github.com/gofaith/faithdroid/")
		if e != nil {
			fmt.Println(`ls error :`, e)
			return
		}
		for _, v := range fs {
			c = exec.Command("cp", "-r", fileToolkit.Getrpath(gopath)+"github.com/gofaith/faithdroid/"+v, "faithdroid")
			c.Stdin = os.Stdin
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
			e = c.Run()
			if e != nil {
				fmt.Println(`command.Run error :`, e)
				continue
			}
		}
		return
	case "update":
		fs, e := listDir(fileToolkit.Getrpath(gopath) + "github.com/gofaith/faithdroid/")
		if e != nil {
			fmt.Println(`ls error :`, e)
			return
		}
		for _, v := range fs {
			c = exec.Command("cp", "-r", fileToolkit.Getrpath(gopath)+"github.com/gofaith/faithdroid/"+v, ".")
			c.Stdin = os.Stdin
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
			e = c.Run()
			if e != nil {
				fmt.Println(`command.Run error :`, e)
				continue
			}
		}
		return
	default:
		fmt.Println("command not found")
		return
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
func listDir(dir string) ([]string, error) {
	c := exec.Command("ls", dir)
	buf := bytes.NewBufferString("")
	c.Stdout = buf
	e := c.Run()
	if e != nil {
		return nil, e
	}
	ss := strings.Split(buf.String(), "\n")
	return ss[:len(ss)-2], nil
}
