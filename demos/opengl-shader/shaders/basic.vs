uniform mat4 uProjMatrix;
uniform mat4 uViewMatrix;
uniform mat4 uModelMatrix;

attribute vec3 aPosition;
attribute vec3 aColor;

void main() {
    vec4 vPosition = vec4(aPosition.x, aPosition.y, aPosition.z, 1.0);
    vPosition = uViewMatrix * vPosition;
    gl_Position = uProjMatrix * vPosition;

    gl_FrontColor = vec4(aColor.r, aColor.g, aColor.b, 1.0);
}
