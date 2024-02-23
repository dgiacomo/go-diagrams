# Add New Entity, Image and Binary Go Representation of Image
* add PNG image to `assets/<dir>/<file>`  (this needs to be a PNG)
* add go function to render image to `assets/<dir>/<file.go>`
* run generate to convert the images into [embedded binary go data](https://github.com/UnnoTed/fileb0x): `cd nodes/assets; go generate`
