// Code generaTed by fileb0x at "2023-09-21 09:26:25.97826158 -0500 CDT m=+1.615598195" from config file "b0x.yml" DO NOT EDIT.
// modified(2023-09-21 08:15:37.676143479 -0500 CDT)
// original path: ../../assets/elastic/beats/auditbeat.png

package assets

import (
	"bytes"
	"compress/gzip"
	"io"

	"os"
)

// FileAssetsElasticBeatsAuditbeatPng is "assets/elastic/beats/auditbeat.png"
var FileAssetsElasticBeatsAuditbeatPng = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xea\x0c\xf0\x73\xe7\xe5\x92\xe2\x62\x60\x60\xe0\xf5\xf4\x70\x09\x62\x60\x60\x64\x00\x61\x26\x66\x06\x06\x06\x45\x19\x87\xaf\x0c\x0c\x0c\x2c\xe9\x8e\xbe\x8e\x0c\x0c\x1b\xfb\xb9\xff\x24\xb2\x32\x30\x30\x28\x24\x7b\x04\xf9\x32\x30\x54\xa9\x31\x30\x34\xb4\x30\x30\xfc\x62\x60\x60\x68\x78\xc1\xc0\x50\x6a\xc0\xc0\xf0\x2a\x81\x81\xc1\x6a\x06\x03\x83\x78\xc1\x9c\x5d\x81\x36\x0c\x0c\x0c\x3c\x01\x3e\x21\xae\x0c\x0c\x0c\xc1\x91\x69\x0c\xfb\x37\xff\xff\xff\xdf\xf1\x7a\xf3\x16\x90\x0d\x25\x41\x7e\xc1\x0c\x0e\xcf\x6e\xa4\x81\x38\x49\xde\xee\x2e\xcc\x82\x3c\x3e\x9f\x18\x18\x18\x38\x0b\x3c\x22\x8b\x19\x18\xb8\x85\x41\x98\x91\x61\xd6\x1c\x09\x06\x06\x06\xf6\x12\x4f\x5f\x57\xf6\x67\x1c\x02\x1c\x46\x22\x53\x25\xe7\x5f\x65\x60\x60\xc8\xf2\x74\x71\x0c\xa9\xb8\xf5\xf6\xf2\x46\x4e\x06\x05\x1e\x57\xc7\xe9\x25\x27\x7d\x34\x9b\x5a\x9d\x94\x6e\x9c\xda\x6b\xf6\x70\xd7\xfd\x59\x2a\x2f\xd7\x2c\x3f\xeb\xed\x72\x47\x9e\xa1\xe1\x7e\x43\x92\x2a\xf7\x89\xef\x20\x3f\x11\x07\x3e\x30\xa9\xe4\x3a\xfc\x65\x60\xa8\x61\xf5\x98\x94\xc6\x1c\xad\xd7\xbd\x36\xe5\x90\xc5\x36\x06\x06\x06\xd5\x12\xd7\x88\x92\x94\xc4\x92\x54\xab\xe4\xa2\xd4\xc4\x92\x54\x06\x23\x03\x23\x23\x5d\x03\x0b\x5d\x43\xb3\x10\x03\x0b\x2b\x63\x73\x2b\x03\x43\x6d\x03\x03\x2b\x03\x83\x0a\x81\xad\x1f\x51\x34\xe4\xe6\xa7\x64\xa6\x55\xa2\x68\x30\xb7\x32\x36\xb1\x32\x31\x82\x68\x28\x58\x51\x2e\xc3\xc0\xc0\xa0\x01\xd7\x50\x92\x99\x9b\x5a\x5c\x92\x98\x5b\x80\x66\x89\xa9\x81\x95\x91\x01\x44\xcf\x86\xe4\x8f\x89\x20\xf7\x7a\xba\xfa\xb9\xac\x73\x4a\x68\x02\x04\x00\x00\xff\xff\x10\x27\xd3\xdc\xcf\x01\x00\x00")

func init() {

	rb := bytes.NewReader(FileAssetsElasticBeatsAuditbeatPng)
	r, err := gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err := FS.OpenFile(CTX, "assets/elastic/beats/auditbeat.png", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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