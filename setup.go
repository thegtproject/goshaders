package main

import (
	"fmt"

	"github.com/faiface/glhf"
	"github.com/faiface/mainthread"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/pkg/errors"
)

func main() {
	defer glfw.Terminate()

	err := glfw.Init()
	if err != nil {
		panic(errors.Wrap(err, "failed to initialize GLFW"))
	}
	generateShaderFileEntries()
	mainthread.Run(setup)
}

var mt = mainthread.Call

func setup() {
	mt(createWindow)
	mt(centerWindow)
	mt(compileFragmentShaders)
	createCallbacks()
	mt(activateSelectedShader)
	mt(createBlankTexture)
	mt(loop)
	fmt.Println()
}

func createWindow() {
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.False)
	glfw.WindowHint(glfw.Resizable, glfw.False)
	var err error
	win, err = glfw.CreateWindow(width, height, "goshaders", nil, nil)
	if err != nil {
		panic(err)
	}
	win.MakeContextCurrent()
	glhf.Init()
}

func centerWindow() {
	winWidth, winHeight := win.GetSize()
	win.SetPos(
		glfw.GetPrimaryMonitor().GetVideoMode().Width/2-winWidth/2,
		glfw.GetPrimaryMonitor().GetVideoMode().Height/2-winHeight/2,
	)
}

func compileFragmentShaders() {
	for _, se := range shaders {
		se.compile()
	}
}

func createCallbacks() {
	win.SetCursorPosCallback(
		func(_ *glfw.Window, xpos float64, ypos float64) {
			mouse[0], mouse[1] = float32(xpos), float32(ypos)
		},
	)
	win.SetMouseButtonCallback(
		func(_ *glfw.Window, button glfw.MouseButton, action glfw.Action, _ glfw.ModifierKey) {
			if !(button == glfw.MouseButton1 || button == glfw.MouseButton2) {
				return
			}
			switch action {
			case glfw.Press:
				mouse[int(button)+2] = 1.0
			case glfw.Release:
				mouse[int(button)+2] = 0.0
			}
		},
	)
	win.SetKeyCallback(
		func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
			if action == glfw.Press {
				handleKeyPress(key)
			}
		},
	)
}

func createBlankTexture() {
	texture = glhf.NewTexture(
		nrgba.Bounds().Dx(),
		nrgba.Bounds().Dy(),
		true,
		nrgba.Pix,
	)
}
func updateVertexShader() {
	slice = glhf.MakeVertexSlice(shader, 6, 6)
	slice.Begin()
	slice.SetVertexData([]float32{
		-1, -1, 0, 1,
		+1, -1, 1, 1,
		+1, +1, 1, 0,

		-1, -1, 0, 1,
		+1, +1, 1, 0,
		-1, +1, 0, 0,
	})
	slice.End()
}

func handleKeyPress(key glfw.Key) {
	switch key {
	case glfw.KeyRight:
		nextShader()
	case glfw.KeyLeft:
		prevShader()
	case glfw.KeySpace:
		activateSelectedShaderbool = true
	case glfw.KeyR:
		reloadcompile()
	case glfw.KeyEscape:
		fallthrough
	case glfw.KeyQ:
		win.SetShouldClose(true)
	}
}

func reloadcompile() {
	if err := shaders[selectedShader].reload(); err != nil {
		fmt.Println("\r\nerror", err)
		return
	}
	shaders[selectedShader].compile()
	activateSelectedShaderbool = true
}
