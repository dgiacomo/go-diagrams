// Code generaTed by fileb0x at "2023-09-21 09:26:26.106395788 -0500 CDT m=+1.743732403" from config file "b0x.yml" DO NOT EDIT.
// modified(2023-09-21 08:15:37.676143479 -0500 CDT)
// original path: ../../assets/elastic/beats/winlogbeat.png

package assets

import (
	"bytes"
	"compress/gzip"
	"io"

	"os"
)

// FileAssetsElasticBeatsWinlogbeatPng is "assets/elastic/beats/winlogbeat.png"
var FileAssetsElasticBeatsWinlogbeatPng = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xea\x0c\xf0\x73\xe7\xe5\x92\xe2\x62\x60\x60\xe0\xf5\xf4\x70\x09\x62\x60\x60\x64\x00\x61\x26\x66\x06\x06\x06\x45\x19\x87\xaf\x0c\x0c\x0c\x2c\xe9\x8e\xbe\x8e\x0c\x0c\x1b\xfb\xb9\xff\x24\xb2\x32\x30\x30\x28\x24\x7b\x04\xf9\x32\x30\x54\xa9\x31\x30\x34\xb4\x30\x30\xfc\x62\x60\x60\x68\x78\xc1\xc0\x50\x6a\xc0\xc0\xf0\x2a\x81\x81\xc1\x6a\x06\x03\x83\x78\xc1\x9c\x5d\x81\x36\x0c\x0c\x0c\x3c\x01\x3e\x21\xae\x0c\x0c\x0c\xc1\x91\x69\x0c\xfb\x37\xff\xff\xff\xdf\xf1\x7a\xf3\x16\x90\x0d\x25\x41\x7e\xc1\x0c\x0e\xcf\x6e\xa4\x81\x38\x49\xde\xee\x2e\xcc\x82\x3c\x3e\x9f\x18\x18\x18\x38\x0b\x3c\x22\x8b\x19\x18\xb8\x85\x41\x98\x91\x61\xd6\x1c\x09\x06\x06\x06\xf6\x12\x4f\x5f\x57\xf6\x67\x1c\x02\x1c\xc6\x16\xfb\x6e\x1d\x2a\x67\x60\x60\xa8\xf3\x74\x71\x0c\xa9\xb8\xf5\xf6\x9a\x21\x57\x83\x81\x80\xeb\xc1\xbd\x4b\xa6\xdf\x09\xee\x4a\x48\x54\xd0\x4e\x12\x16\x7e\xbe\x35\x72\xd9\xef\x09\x1c\x9e\x41\xb7\xaf\xbd\xad\x38\x79\xe6\x37\x33\x43\x82\xe9\x82\x88\x93\x9c\x5f\xbe\x83\xfc\x85\x0c\xfe\xf0\xad\x58\xdd\xd8\xce\xc8\xd0\x70\xb2\xa1\x46\x9c\x61\x83\x8c\x41\x5d\x2f\x03\xc3\x03\xed\x03\x3f\xee\x5b\xa1\xa9\x34\xe0\xbd\xf0\x92\xe9\x0a\x3b\xc3\x81\xa0\x06\x15\xd7\x37\x93\x4b\x1f\xa5\xcc\x67\xfb\x19\xcd\xc0\xc0\xa0\x5a\xe2\x1a\x51\x92\x92\x58\x92\x6a\x95\x5c\x94\x9a\x58\x92\xca\x60\x64\x60\x64\xa4\x6b\x60\xa1\x6b\x68\x16\x62\x60\x61\x65\x6c\x66\x65\x68\xa0\x6d\x60\x60\x65\x60\xf0\x97\xff\xea\x53\x14\x0d\xb9\xf9\x29\x99\x69\x95\x28\x1a\xcc\xad\x8c\x8d\xad\x4c\xcc\x20\x1a\xd2\xac\x3d\xca\x18\x18\x18\x34\xe0\x1a\x4a\x32\x73\x53\x8b\x4b\x12\x73\x0b\xd0\x2c\x31\x35\xb4\x32\x85\xea\x31\xdb\xb2\xac\x06\xe4\x5e\x4f\x57\x3f\x97\x75\x4e\x09\x4d\x80\x00\x00\x00\xff\xff\xe2\x96\x9c\xdd\xe3\x01\x00\x00")

func init() {

	rb := bytes.NewReader(FileAssetsElasticBeatsWinlogbeatPng)
	r, err := gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err := FS.OpenFile(CTX, "assets/elastic/beats/winlogbeat.png", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}
