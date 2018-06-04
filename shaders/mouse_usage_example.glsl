uniform vec4 u_mouse;
uniform vec2 u_resolution;

void main() {
    vec2 mouse_norm = vec2( u_mouse.x/u_resolution.x, 1.0 - u_mouse.y/u_resolution.y );
    vec3 color = vec3(mouse_norm.x, mouse_norm.y, u_mouse.z);
    gl_FragColor = vec4(color-1.0*u_mouse.w*0.5, 1.0);
}