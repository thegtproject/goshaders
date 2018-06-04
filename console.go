package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const welcomeStr = "goshaders | left and right to cycle, space to activate, esc to exit"

func termclear() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		return
	}
	const clearScreenSeq = "\x1b\x5b\x48"
	fmt.Print(clearScreenSeq)
}

func termprint() {
	var s string
	if selectedShader == activeShaderIndex {
		s = "active"
	} else {
		s = "not active"
	}
	win.SetTitle(shaders[selectedShader].name + " (" + s + ")")
}
