// Star Nest by Pablo RomÃ¡n Andrioli
 
// This content is under the MIT License.
 
#define iterations 17
#define formuparam 0.53
 
#define volsteps 20
#define stepsize 0.1
 
#define zoom   0.800
#define tile   0.850
#define speed  0.010
 
#define brightness 0.0015
#define darkmatter 0.300
#define distfading 0.730
#define saturation 0.850

uniform vec2  u_resolution;
uniform float u_time;
uniform vec4  u_mouse;

void main()
{
    //get coords and direction
    vec2 uv=gl_FragCoord.xy/u_resolution.xy-.5;
    uv.y*=u_resolution.y/u_resolution.x;
    vec3 dir=vec3(uv*zoom,1.);
    float time=u_time*speed+.25;
 
    //mouse rotation
    float a1=.5+u_mouse.x/u_resolution.x*2.;
    float a2=.8+u_mouse.y/u_resolution.y*2.;
    mat2 rot1=mat2(cos(a1),sin(a1),-sin(a1),cos(a1));
    mat2 rot2=mat2(cos(a2),sin(a2),-sin(a2),cos(a2));
    dir.xz*=rot1;
    dir.xy*=rot2;
    vec3 from=vec3(1.,.5,0.5);
    from+=vec3(time*2.,time,-2.);
    from.xz*=rot1;
    from.xy*=rot2;
   
    //volumetric rendering
    float s=0.1,fade=1.;
    vec3 v=vec3(0.);
    dir = dir * .5;
    vec3 ptile = vec3(tile);
    vec3 ptile_s = vec3(tile*2.);
    for (int r=0; r<volsteps; r++) {
        vec3 p=s*dir+from;
        p = abs(ptile-mod(p,ptile_s)); // tiling fold
        float pa,a=pa=0.;
        float dpp = dot(p,p);
        for (int i=0; i<iterations; i++) {
            p=abs(p)/dpp-formuparam; // the magic formula
            dpp = dot(p,p);
            float lp = sqrt(dpp);
            a+=abs(lp-pa); // absolute sum of average change
            pa=lp;
        }
        p = vec3(0.0f,0.0f,0.0f);
            pa = 0.0f;
        float aa = a*a;
        float dm=max(0.,-0.001*aa+darkmatter); //dark matter (reversed order, MAD operations are faster (multiply then add))
        a*=aa; // add contrast
        if (r>6) fade*=1.-dm; // dark matter, don't render near
        //v+=vec3(dm,dm*.5,0.);
        v+=fade;
        float ss = s*s;
        v+=vec3(s,ss,ss*ss)*a*brightness*fade; // coloring based on distance
        fade*=distfading; // distance fading
        s+=stepsize;
    }
    v=mix(vec3(length(v)),v,saturation); //color adjust
    gl_FragColor = vec4(v*.01,1.);
}