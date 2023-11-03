/*****************************************************************************
*                    
*  Author:           Jarette Greene
*  Email:            jkgreene0406@my.msutexas.edu / jarettegreene09@gmail.com
*  Label:            P03
*  Title:			 Image Ascii Art           
*  Course:           CMPS 4143
*  Semester:         Fall 2023
* 
*  Description:
*	
*		This Package gets the RGB values of every pixel in an image 
*
* 
*  Usage:
*    - import to main.go file 
* 
*  Files           
*       N/A 
*****************************************************************************/
package Img_colors
//necessary packages
import (
	// fmt : allows for input and output from and to the console
	"fmt"
	// image : decodes the rgb values of an image
	"image"
	_ "image/png" // import this package to decode PNGs

	//os : allows to check for errors when opening a file
	"os"
)

//Colors gets an image from the system and displays all the RGB value of each pixel of the image 
func Colors() {
	// Opens image and checks for errors when opeing image 
	reader, err := os.Open("downloaded_image.jpg")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	//closes image at the end of fucntion call
	defer reader.Close()

	//allows for the RGB values of the image to be decoded and checks for errors 
	img, _, err := image.Decode(reader)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// finding the boundaries of the image (height and width )
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	// loop that goes through every pixel of the image and displays them to the console
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			color := img.At(x, y)
			r, g, b, _ := color.RGBA()
			fmt.Printf("Pixel at (%d, %d) - R: %d, G: %d, B: %d\n", x, y, r>>8, g>>8, b>>8)
		}
	}
}
