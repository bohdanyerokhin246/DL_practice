package main

import fmt "fmt"

type Volume interface {
	volume() float32
}
type Square interface {
	square() float32
}

type Prisms struct {
	S float32
	H float32
}
type Pyramids struct {
	S float32
	H float32
}
type Cylinders struct {
	Pi float32
	R  float32
	H  float32
}
type Cones struct {
	Pi float32
	R  float32
	H  float32
}
type Circles struct {
	Pi float32
	R  float32
}
type Layers struct {
	Pi float32
	R  float32
}

func (prism Prisms) volume() float32 {
	return prism.S * prism.H
}
func (pyramid Pyramids) volume() float32 {
	return 0.33 * pyramid.S * pyramid.H
}
func (cylinder Cylinders) volume() float32 {
	return cylinder.Pi * cylinder.R * cylinder.R * cylinder.H
}
func (cone Cones) volume() float32 {
	return 0.33 * cone.Pi * cone.R * cone.R * cone.H
}
func (circle Circles) square() float32 {
	return circle.Pi * circle.R * circle.R
}
func (layer Layers) square() float32 {
	return 4 * layer.Pi * layer.R * layer.R
}

func getVolume(volume Volume) float32 {
	return volume.volume()
}
func getSquare(square Square) float32 {
	return square.square()
}

func main() {
	prism := Prisms{S: 12, H: 6}
	pyramid := Pyramids{S: 12, H: 6}
	cylinder := Cylinders{Pi: 3.14, R: 10, H: 6}
	cone := Cones{Pi: 3.14, R: 10, H: 6}
	circle := Circles{Pi: 3.14, R: 10}
	layer := Layers{Pi: 3.14, R: 10}

	fmt.Println("Prism volume =", getVolume(prism))
	fmt.Println("Pyramid volume =", getVolume(pyramid))
	fmt.Println("Cylinder volume =", getVolume(cylinder))
	fmt.Println("Cone volume =", getVolume(cone))
	fmt.Println("Circle square =", getSquare(circle))
	fmt.Println("Layer square =", getSquare(layer))
}
