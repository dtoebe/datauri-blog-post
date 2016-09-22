package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/rand"
	"time"
)

func main() {

	//Seed the std random number generator
	rand.Seed(time.Now().UTC().UnixNano())

	// We will leave out creating the png image file since we
	// don't need the file

	// Set the bounds of the image basically 100px width / height
	imgRect := image.Rect(0, 0, 100, 100)
	// This creates the image in memory, and sets all pixels to a gray value.
	// And it sets the size based on the Rect we just created
	img := image.NewGray(imgRect)

	// Ok a lot going on here. Let's start creates a nil mask layer on the image
	// Tt takes the image we created in memory, and then the bounding Rect
	// Next we add a reference to a uniform color (white) over the mask
	// The image.ZP means the mask we are drawing starts at the 0,0 point
	// draw.Src refrences the source mask
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)

	// These nested loop reference the y,x coordinates.
	// 100 is reference to the image with / height in pixels
	// 10 is reference to the size of the random blocks to be generated
	for y := 0; y < 100; y += 10 {
		for x := 0; x < 100; x += 10 {
			// Here we create a color to be set as a black block
			fill := &image.Uniform{color.Black}
			// rand.Intn(10) creates a random non-negative number between 0,10
			// Then checks if it's even
			if rand.Intn(10)%2 == 0 {
				fill = &image.Uniform{color.White}
			}
			// This draw.Draw takes a new Rect that is 10px,10px and sets
			// color based on fill
			draw.Draw(img, image.Rect(x, y, x+10, y+10), fill, image.ZP, draw.Src)
		}
	}

	// Here we allocate and create a buffer that can easily be
	// turned into a []byte
	out := new(bytes.Buffer)

	// We now encode the image we created to the buffer
	err := png.Encode(out, img)
	if err != nil {
		// Handle errors
	}

	// This now takes a []byte of the buffer and base64 encodes it to a string
	// Never needing to create the image file all done in memory
	base64Img := base64.StdEncoding.EncodeToString(out.Bytes())

	//And now you can see the magic happen.
	// Go ahead and run it then take the output and copy/paste it in your
	// browser's URL bar, and you'll see your image.
	fmt.Println("data:image/png;base64,", base64Img)
}
