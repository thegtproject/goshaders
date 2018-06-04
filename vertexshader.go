package main

var vertexShader = `
#version 330 core
in vec2 position;
in vec2 texture;

out vec2 Texture;

void main() {
	gl_Position = vec4(position, 0.0, 1.0);
	Texture = texture;
}
`

// var vertexShader = `
// #version 330

// uniform mat4 projection;
// uniform mat4 camera;
// uniform mat4 model;

// in vec3 vert;
// in vec2 vertTexCoord;
// in vec2 texture;
// in vec2 position;
// out vec2 fragTexCoord;
// out vec2 Texture;
// void main() {
//     fragTexCoord = vertTexCoord;
//     gl_Position = vec4(position, 0.0, 1.0);
// }
// `
