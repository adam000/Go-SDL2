Go-SDL2 is a collection of packages for Go bindings for the
[Simple DirectMedia Layer](http://libsdl.org) version 2, known as "SDL2".

These bindings aim to be very Go-idiomatic and intuitive, and will include
bindings for SDL packages such as SDL-ttf, SDL-image, and more.

Examples of how to use these packages are available
[here](https://github.com/adam000/go-sdl-gl-examples).

**These bindings are still experimental and the API is subject to change.**
There will eventually be an API freeze when the bindings are more
complete and tested.

*This software is provided free of charge. See LICENSE for details.*

# Installing

The packages are designed to be go-getable, but they require that the underlying
C library be installed first.  The cgo directives depend on the availability of
pkg-config for the SDL libraries.  The packages are separated by the SDL library
required.

* `github.com/adam000/Go-SDL2/sdl/...` requires SDL2 2.0.2 or higher.
* `github.com/adam000/Go-SDL2/image` requires SDL2\_image-2.0.0 or higher.
