package main

import (
	"encoding/binary"
	"math"
	"time"
	"unsafe"

	mat "bitbucket.org/zombiezen/math3/mat32"
	"github.com/adam000/Go-SDL2/sdl"
	"github.com/adam000/shader"
	"github.com/go-gl/gl"
)

const basicVertexShaderFileName = "shaders/basic.vs"

// {{{
func getGlError(context string) {
	switch gl.GetError() {
	case gl.NO_ERROR:
	case gl.INVALID_ENUM:
		panic("Invalid enum at " + context)
	case gl.INVALID_VALUE:
		panic("Invalid value at " + context)
	case gl.INVALID_OPERATION:
		panic("Invalid operation at " + context)
	case gl.INVALID_FRAMEBUFFER_OPERATION:
		panic("Invalid framebuffer operation at " + context)
	case gl.OUT_OF_MEMORY:
		panic("Out of memory at " + context)
	case gl.STACK_UNDERFLOW:
		panic("Stack underflow at " + context)
	case gl.STACK_OVERFLOW:
		panic("Stack overflow at " + context)
	}
} // }}}

func convertMat32(mat *mat.Matrix) *[16]float32 {
	return (*[16]float32)(unsafe.Pointer(mat))
}

func makeSymmetricProjectionMatrix() mat.Matrix {
	// 80 degrees in radians
	fovy := 1.396
	aspect := 4 / 3.0
	nearPlane := 0.01
	farPlane := 100.0

	screenRange := math.Tan(fovy/2) * nearPlane
	right := screenRange * aspect
	top := screenRange

	var mat mat.Matrix
	mat[0][0] = float32(nearPlane / right)
	mat[1][1] = float32(nearPlane / top)
	mat[2][2] = float32(-(farPlane + nearPlane) / (farPlane - nearPlane))
	mat[2][3] = -1
	mat[3][2] = float32((-2 * farPlane * nearPlane) / (farPlane - nearPlane))

	return mat
}

// Draw something! in SDL using OpenGL with shaders / retained mode
func main() {
	const (
		width  = 1024
		height = 768
	)

	triVertexes := [][3]gl.GLfloat{
		{2.0, 0.0, -5.0},
		{0.0, 4.0, -5.0},
		{-2.0, 0.0, -5.0},
	}

	triColors := [][3]gl.GLfloat{
		{1.0, 0.0, 0.0},
		{0.0, 1.0, 0.0},
		{0.0, 0.0, 1.0},
	}

	// SDL Initialization
	sdl.Init(sdl.InitEverything)
	defer sdl.Quit()

	window, err := sdl.NewWindow(
		"Hello world!",
		sdl.Rect(sdl.WindowPosCentered, sdl.WindowPosCentered, width, height),
		sdl.WindowOpenGL)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	context, err := sdl.NewGLContext(window)
	if err != nil {
		panic(err)
	}
	defer context.Destroy()

	// OpenGL initialization
	if err := gl.Init(); err != 0 {
		panic("Problem in GL initialization")
	}

	// Initialize what color and depth to clear to
	gl.ClearColor(0, 0, 0, 0)
	gl.ClearDepth(1.0)
	gl.DepthFunc(gl.LEQUAL)
	gl.Enable(gl.DEPTH_TEST)

	// Shader initialization
	vertex, err := shader.NewShader(basicVertexShaderFileName, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	// Shader Program initialization
	program, err := shader.CompileProgramFromShaders(vertex)
	if err != nil {
		panic(err)
	}

	program.Use()

	posAttrib := program.Attrib("aPosition")
	colorAttrib := program.Attrib("aColor")
	projUniform := program.Uniform("uProjMatrix")
	viewUniform := program.Uniform("uViewMatrix")
	modelUniform := program.Uniform("uModelMatrix")

	// Buffer initialization
	posBuffer := gl.GenBuffer()
	posBuffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, binary.Size(triVertexes), triVertexes, gl.STATIC_DRAW)

	colorBuffer := gl.GenBuffer()
	colorBuffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, binary.Size(triColors), triColors, gl.STATIC_DRAW)

	// Draw - this portion suitable for a loop
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	program.Use()

	// Set uniforms - projection and view only need to be set once, model is
	// done per-object
	projection := makeSymmetricProjectionMatrix()
	projUniform.UniformMatrix4f(false, convertMat32(&projection))

	// Note: The view / model matrices are not really needed for this exercise,
	// but are useful in other applications.
	viewUniform.UniformMatrix4f(false, convertMat32(&mat.Identity))
	modelUniform.UniformMatrix4f(false, convertMat32(&mat.Identity))

	posAttrib.EnableArray()
	posBuffer.Bind(gl.ARRAY_BUFFER)
	posAttrib.AttribPointer(3, gl.FLOAT, false, 0, nil)

	colorAttrib.EnableArray()
	colorBuffer.Bind(gl.ARRAY_BUFFER)
	colorAttrib.AttribPointer(3, gl.FLOAT, false, 0, nil)

	gl.DrawArrays(gl.TRIANGLES, 0, 3)

	posAttrib.DisableArray()
	colorAttrib.DisableArray()

	gl.ProgramUnuse()
	window.GLSwap()

	time.Sleep(time.Second * 7)
}
