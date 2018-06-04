package main

import (
	"time"

	"github.com/go-gl/glfw/v3.1/glfw"
)

func loop() {
	termprint()
	for !win.ShouldClose() {
		elapsed = float32(time.Since(start).Seconds())
		shader.Begin()
		shader.SetUniformAttr(0, elapsed)
		shader.SetUniformAttr(1, resolution)
		shader.SetUniformAttr(2, mouse)
		shader.SetUniformAttr(3, playerpositions)
		texture.Begin()
		slice.Begin()
		slice.Draw()
		slice.End()
		texture.End()
		shader.End()

		win.SwapBuffers()
		glfw.PollEvents()

		if activateSelectedShaderbool {
			activateSelectedShader()
		}
	}
}
