package main

import "fmt"

var activateSelectedShaderbool = false

func activateSelectedShader() {
	activateSelectedShaderbool = false
	activeShaderIndex = selectedShader
	if shaders[selectedShader].obj == nil {
		fmt.Println("this shader has not compiled correctly. press r to try to recompile")
		return
	}
	shader = shaders[selectedShader].obj
	updateVertexShader()
	termprint()
}

var selectedShader = 0

func selectShader(index int) {
	selectedShader = index
	termprint()
}

func nextShader() {
	if selectedShader+1 >= len(shaders) {
		return
	}
	selectShader(selectedShader + 1)
}

func prevShader() {
	if selectedShader-1 < 0 {
		return
	}
	selectShader(selectedShader - 1)
}

func selectShaderByName(name string) {
	for i, se := range shaders {
		if se.name == name+shaderFileExt {
			selectShader(i)
		}
	}
	return
}
