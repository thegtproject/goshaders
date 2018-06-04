package main

import (
	"image"
	"time"

	"github.com/faiface/glhf"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

var (
	width, height = 1000, 800
	vertexFormat  = glhf.AttrFormat{
		{Name: "position", Type: glhf.Vec2},
		{Name: "texture", Type: glhf.Vec2},
	}
	uniformsFormat = glhf.AttrFormat{
		glhf.Attr{
			Name: "u_time",
			Type: glhf.Float,
		},
		glhf.Attr{
			Name: "u_resolution",
			Type: glhf.Vec2,
		},
		glhf.Attr{
			Name: "u_mouse",
			Type: glhf.Vec4,
		},
		glhf.Attr{
			Name: "u_players",
			Type: glhf.Vec4,
		},
	}

	fwidth, fheight   = float32(width), float32(height)
	activeShaderIndex int

	resolution      = mgl32.Vec2{fwidth, fheight}
	mouse           = mgl32.Vec4{}
	elapsed         float32
	playerpositions = mgl32.Vec4{-0.1, -0.1, 0.1, 0.1}

	start = time.Now()

	nrgba = image.NewNRGBA(image.Rect(0, 0, width, height))

	shader, shader2 *glhf.Shader
	texture         *glhf.Texture
	slice           *glhf.VertexSlice
	win             *glfw.Window
)

const (
	shaderDirectory    = "shaders/"
	shaderStringPrefix = "#version 330 core\n"
)
