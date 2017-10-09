# ICE!

A game about covering as much ground as you can!

Made for Open Jam 17, theme: "Leave A Mark"

Control penguins using WASD and arrow keys

All made in golang using the [Pixel](https://www.github.com/faiface/pixel) library

## Requirements

* Golang

Install golang for your distro

* [PixelGL](https://godoc.org/github.com/faiface/pixel/pixelgl) backend uses OpenGL to render
graphics. Because of that, OpenGL development libraries are needed for compilation. The dependencies
are same as for [GLFW](https://github.com/go-gl/glfw).

- On macOS, you need Xcode or Command Line Tools for Xcode (`xcode-select --install`) for required
  headers and libraries.
- On Ubuntu/Debian-like Linux distributions, you need `libgl1-mesa-dev` and `xorg-dev` packages.
- On CentOS/Fedora-like Linux distributions, you need `libX11-devel libXcursor-devel libXrandr-devel
  libXinerama-devel mesa-libGL-devel libXi-devel` packages.
- See [here](http://www.glfw.org/docs/latest/compile.html#compile_deps) for full details.

**The combination of Go 1.8, macOS and latest XCode seems to be problematic** as mentioned in issue
[#7](https://github.com/faiface/pixel/issues/7). This issue is probably not related to Pixel.
**Upgrading to Go 1.8.1 fixes the issue.**

The above was taken from the [Pixel README.md](https://github.com/faiface/pixel/blob/master/README.md)

* Use `go get github.com/faiface/pixel` !
* There might be certain pieces of libraries missing like `golang.org/x/image/colornames` -- just
  grab them easily with `go get golang.org/x/image/colornames` for example.

## Running ICE!

Just build it using `go build` in the source directory, then just run the program! It's easy as
that.

If any problems arise, feel free to make an Issue!
