package main

// Library imports
import (
	"fmt"; // print debugging
	"os"; // file management
	"image"; // image processing
	_ "image/jpeg"; //jpeg image file initialization
	"image/color";
	"math";
	"github.com/nfnt/resize";
)

func main() {
	// Get the JPEG image file.
	fmt.Print("Enter the filename of the source JPEG image: ")
	var fileName string
	
	fmt.Scanln(&fileName)

	//Prefix our source-images file
	fileName = "source-images/" + fileName

	fmt.Println("File name is: " + fileName)

	// Read the JPEG image file.
	f, err := os.Open(fileName)
	
	// Check the file was opened
	if err != nil {
		panic(err)
	}

	// Close the file when able
	defer f.Close()


	// Decode the image file
	img, format, err := image.Decode(f)

	// Check that the decoded file was JPEG
	if format != "jpeg" {
		fmt.Println("Only jpeg images are supported")
	}

	if err != nil {
		panic(err)
	}

	img = resize.Resize(200, 0, img, resize.Lanczos3)
	size := img.Bounds().Size()

	// loop through all the pixels
	for x := 0; x < size.X; x++ {
			for y := 0; y < size.Y; y++ {
					pix := img.At(y,x)
					originalColor := color.RGBAModel.Convert(pix).(color.RGBA)

					r := float64(originalColor.R) * 1 //0.21
					g := float64(originalColor.G) * 1 //0.72
					b := float64(originalColor.B) * 1 //0.07

					// gray scale average
					grey := uint8((r + g + b) / 3)
					var grayRamp string
					grayRamp = ".'`^\",:;Il!i><~+_-?][}{1)(|/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$"
					rampLength := len(grayRamp) 

					var outputCharacter string
					var indexOfChar = int(math.Ceil(float64(rampLength - 1) * float64(grey) / 255.0))
					outputCharacter = grayRamp[indexOfChar:indexOfChar+1]
					fmt.Print(outputCharacter)
			}
			fmt.Println()
	}
}
