# Adding Images
This project uses the [fileb0x](https://github.com/UnnoTed/fileb0x) library to convert PNG images into binary go data that can be embedded in the code and read out of memory during rutime.
To add an image to the project you'll want to add the image, add the go function to render it, and add generate the bin-data go file: 
* Add your image as a PNG to `assets/<dir>/<file>`  (this needs to be a PNG)
* Add the go function to render to the area it belongs under `assets/<dir>/<file.go>`
* Run generate to convert the images into [embedded binary go data](https://github.com/UnnoTed/fileb0x): `cd nodes/assets; go generate`
* This will unfortunately update all the generated bin data timestamps, which makes it look like a giant git commit - just commit the one you added. 

