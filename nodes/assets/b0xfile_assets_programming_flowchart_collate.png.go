// Code generaTed by fileb0x at "2023-09-21 08:22:37.792398288 -0500 CDT m=+0.824893128" from config file "b0x.yml" DO NOT EDIT.
// modified(2023-09-21 08:15:37.724142836 -0500 CDT)
// original path: ../../assets/programming/flowchart/collate.png

package assets

import (
	"bytes"
	"compress/gzip"
	"io"

	"os"
)

// FileAssetsProgrammingFlowchartCollatePng is "assets/programming/flowchart/collate.png"
var FileAssetsProgrammingFlowchartCollatePng = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x95\x79\x54\x52\x89\x17\xc7\x9f\x22\x62\x63\xd1\x36\x3f\x9a\x19\x2d\x5b\x26\xab\x51\x0a\x75\x1a\xd3\x72\xd2\xd2\x71\xcb\x25\xa3\x32\x34\x97\xdc\x35\xc3\x63\x5a\x59\x2a\xab\xa0\x02\x02\x62\xb9\x61\x52\x6a\x56\xa3\x18\xee\x69\x2a\x66\x5a\x02\x6a\xa9\x69\x4d\x0b\x2d\x90\x0b\x99\x0d\x3a\x95\x08\xbc\xdf\x31\xce\xe9\xf4\xdb\xfe\x98\xdf\xf7\x9c\x77\xce\xbd\xef\xbd\xcf\xbd\xdf\x7b\xdf\x79\xef\x65\xf9\x7a\xbb\x2e\xf9\xe6\x87\x6f\x00\x00\x58\xe2\xee\xe6\xec\x07\x00\x7a\xc0\xc2\x61\x04\x01\x00\x20\xae\xca\x1f\x0d\x00\x80\x41\x94\x93\x97\x13\x00\x08\x98\xc6\xea\x50\x28\x00\x00\x6b\xc3\xdc\xfc\xbc\x00\xe0\xac\x39\x00\xe0\xc9\x00\xa0\x02\x00\x00\x3f\x0e\x00\xc9\x28\x00\x98\x0c\x01\x00\xfb\x02\x00\x58\x95\xc0\x6d\xd9\xbf\x0b\x00\xf4\x5f\xf9\xee\x43\xbb\x00\x00\x40\x26\x93\x33\x32\x32\x48\x24\x12\x91\x48\xa4\x52\xa9\x7c\x3e\xff\xfa\xf5\xeb\x61\x61\x61\x91\x91\x91\x14\x0a\xe5\xdc\xb9\x73\x78\x3c\x3e\x39\x39\x99\x40\x20\xe4\xe6\xe6\x86\x84\x84\xe4\xe5\xe5\xb1\xd9\x6c\x2e\x97\x1b\x17\x17\x87\x46\xa3\x0b\x0a\x0a\xe8\x74\x7a\x5b\x5b\x5b\x54\x54\x54\x6c\x6c\xac\xb5\xb5\x75\x51\x51\x51\x66\x66\x66\x4c\x4c\x4c\x40\x40\xc0\x85\x0b\x17\x98\x4c\x66\x45\x45\x45\x44\x44\x44\x76\x76\x76\x56\x56\x56\x6a\x6a\xaa\xad\xad\xad\x42\xa1\xa0\xd1\x68\x0c\x06\x43\x20\x10\xe8\xe9\xe9\x81\x20\xc8\x62\xb1\x46\x47\x47\x17\x2f\x5e\x9c\x93\x93\x63\x64\x64\x24\x14\x0a\x39\x1c\x4e\x65\x65\x25\x8f\xc7\x4b\x4f\x4f\x4f\x4a\x4a\xc2\xe1\x70\x09\x09\x09\x3a\x9f\xa4\xcf\x22\x12\x89\x64\x32\x99\xf4\x1f\x22\x7e\x96\x2e\xa6\x50\x28\x5f\xa7\x04\x02\xe1\xbf\xe2\xba\x93\x78\x3c\x9e\xc5\x62\x91\xc9\xe4\x2f\x48\x5e\x5e\xde\x97\x82\x1c\x0e\x47\x87\xe8\x52\x26\x93\xa9\x8b\x19\x0c\x06\x89\x44\x22\x93\xc9\x39\x9f\xc5\x66\xb3\x75\x54\x6e\x6e\x2e\x87\xc3\xd1\x35\xc5\xe3\xf1\x44\x22\x31\x3d\x3d\xfd\x6b\x87\x5f\x1b\xd3\x15\xf9\x7a\x04\x1c\x0e\x47\x20\x10\x74\xb1\xce\x36\x85\x42\x21\x91\x48\x34\x1a\xed\x0b\x48\xa7\xd3\x49\x24\x92\x6e\xb1\xd9\xd9\xd9\xba\xab\xba\x41\x74\xec\xbf\x09\x8f\xc7\xeb\x9e\x35\x89\x44\xc2\xe1\x70\x3a\xdb\x5f\x36\xf3\xb9\x31\x29\x83\x9a\xc1\xc8\xcb\xc9\xb9\xb0\x30\x1d\x25\x93\x72\xb9\xa9\xa2\x7d\xbc\x47\xa8\xb8\x27\x54\xdc\x6d\x1f\xeb\xbe\xd4\x54\x41\xce\xc8\xa0\xb1\xe9\x35\xfd\x0d\x62\xed\x70\x1f\x38\x22\xd6\x0c\xf3\x25\x75\xd9\x2c\x5a\x36\x93\x56\xd3\x5f\x2f\x52\x0f\x49\xc0\x87\xa2\xf9\xa1\x6a\x71\x2d\x89\x48\x14\x69\x86\xf8\xe2\x3a\x22\x91\xc8\xef\xab\x17\x69\x86\xab\xc5\x75\xd5\x92\x3a\x09\x38\x92\x9d\x93\x5d\xd3\x57\x5f\xd5\x2b\xa0\x66\x51\x29\x54\x4a\x16\x23\x2b\x93\x96\x29\x78\xd0\x78\xe3\x41\x23\x8d\x45\x17\x3c\x68\x12\x6b\x86\x6a\x06\x1a\x6a\x87\x9a\x6f\xdc\x6f\x14\x0c\x36\x0b\x06\x9b\x6e\x0c\x34\xf2\xfb\xeb\x6b\x07\x9b\xb3\x59\x74\x3a\x87\x21\xd1\x3e\xa4\xe7\x32\x78\xf5\x65\x2d\x2f\x3b\x5b\x5f\x77\xb5\xbe\xbe\x5d\x5a\x5f\x46\xce\x20\xb3\x0a\xd9\xcc\x02\x36\x33\x9f\x45\x24\x10\x41\x10\x94\xae\x17\x10\x00\x00\xa8\x49\xf2\xf3\x3e\x00\xfc\xff\x82\x21\x4c\xbe\x35\x30\x0c\x1d\x9a\x51\xfe\xd9\xeb\xbd\xb2\x53\xb8\xfc\x1f\x7d\xb7\x11\xd0\xc0\x57\x32\xb4\x41\xc8\x6b\xf9\x11\x43\xc8\x31\xd9\x11\xa8\x7e\xf8\xab\x00\x48\x10\xf4\xf5\x51\x48\x08\xc6\x20\x2c\x14\x62\x00\x81\x44\x40\xf4\x21\x18\x48\x18\xf4\xd8\xd4\x44\xb0\xe1\xf1\x77\xf1\x10\xc3\x98\x31\x79\x34\x14\xf6\x26\xd2\x28\x16\x1a\x69\x00\x8b\x8b\x36\x88\x32\x80\x42\xc7\xa1\x46\x50\x83\x43\x2b\xbb\xba\xef\xf4\x2c\x85\x78\x8e\xe8\x2d\xdf\xb2\x65\xf3\x52\xd1\x8f\xbe\x7d\x0b\xaf\xfc\x31\x4f\x57\x67\x64\xdf\x1f\x4e\x0e\x0b\x3e\x92\xdc\xbd\x5c\x60\x32\xe3\x25\x7a\x88\x4d\xa9\x8a\x84\x93\x00\x00\x33\x73\x77\x76\x42\x9f\xf9\x63\xaa\xe4\x23\xee\x5b\x97\xef\x7a\xb0\x58\x84\x0b\xdd\xae\x7c\x83\x93\x93\x53\x37\x1f\x9f\x79\x54\xb0\xaa\xdc\xe8\x72\xde\x32\xa3\xa3\xb1\x8e\xdd\xcf\xf5\x8d\x88\x5c\xbf\x15\xe7\x3a\x89\x53\xcb\xee\xac\x19\xdc\xba\xc8\x79\x9b\xfe\xb1\x2e\xe0\x01\xcf\xff\xc6\x3c\xac\xb9\x8e\xf1\x4c\x36\xf9\x7c\x6b\xab\xcd\xc7\x54\xdc\x3d\xe6\xc7\xc7\x63\xbd\x92\x99\xfe\xb6\xe9\xa7\x25\x03\xb0\xbf\xb5\x2b\xeb\x35\x9d\x3e\xce\xaf\x7e\xbe\x3c\xde\x28\x79\xf4\x29\x80\x7b\x5a\x0b\x85\x9f\x7c\xe9\x42\x35\x9d\x6c\xec\x2f\x30\x5a\x15\x9c\x38\x81\x17\x9b\xca\x95\x16\xfc\xe3\x29\x05\x4f\xec\xc3\xf6\xf0\x5b\xec\x65\x3b\x46\x03\x9f\x88\xe2\x4b\xac\x96\xd7\xb4\x06\xa5\xdc\x4b\xce\xe1\xdd\xfa\x20\x34\xc9\x7c\x62\x1f\x66\x82\x8a\x7e\xbf\x4c\x45\x9f\xae\x92\xbe\x00\x14\xd2\x8a\x1a\x87\xd9\xdf\x5e\x1d\x85\x18\x84\x3e\x7a\x5f\x5c\x4e\x7d\xfa\xe9\xf9\xbc\x59\x6a\xb4\xad\x59\x9a\x3a\x6b\x76\x70\xc3\x89\x76\x79\x42\x53\x44\x8f\xb9\x22\xf3\x2d\xbb\xec\x2a\x64\x63\xb1\x87\xf8\xf8\x95\x5b\x97\xfd\xce\xbe\x7c\x10\x96\x58\x57\xf7\xf4\xc2\xf7\x54\xc7\xb8\xd0\xe3\x17\x65\xc2\x96\x84\x83\xc9\x8c\x92\xbc\xf8\xe7\x2b\xf8\x4b\x32\x7a\x18\x03\x1b\xd7\x3c\xd7\xab\x85\xba\xbe\x39\x00\xdb\xf3\xf2\xba\x4f\xfc\x43\xf0\xc8\x8e\xf5\x13\xa5\xa8\x07\xaf\x95\xd7\x20\xf1\x3d\xa2\xd6\xb5\x81\x1f\x6f\xa6\x19\xa2\xb1\x15\xaa\x4d\x5e\x6b\xde\xa4\x69\x7e\x0d\x0e\x08\x15\xa1\x53\x56\x1d\xaf\x8b\x5f\x82\xc6\x5e\x52\xb9\xfa\x84\xbd\x00\xb7\x2d\x8f\x76\x3f\x51\x01\x1f\xa6\x58\x88\xa6\x86\x77\x1f\xdb\xb7\x75\xa5\x6a\x23\xec\x8f\xe5\x37\x11\xc3\x14\xc5\x5b\xb7\xdb\x53\x0f\x03\xda\x59\xfd\x8e\x9c\x1e\x59\xf6\xf1\xf1\x7c\xd3\xe8\xe9\xa6\x1a\x44\xf4\xdc\xee\x33\x97\x47\x67\xa4\xe1\x4e\x54\xf5\x44\xe2\x8a\x77\x6e\xb6\xb0\x59\x75\xd4\x7a\x70\x31\x3c\x49\x52\xba\x6e\xa7\x2f\xfb\x0a\x64\x9f\xbc\x83\xed\x7d\x7b\x5f\x05\xe2\xb8\x97\x58\xaf\x51\x29\x75\xe3\xd9\x14\x6e\xdf\x7b\x0d\x1a\xb3\x63\xad\x3c\x18\xe9\x39\x87\x4d\x7c\xb1\xd9\xfa\x5c\xcb\x2d\x43\xcb\x51\xa2\xf5\x76\xb0\x52\xc0\x01\xcd\x61\x2d\xf5\x45\x07\x78\xdc\xec\x6a\x3b\xdf\xc6\x94\xdf\xb9\x20\xa7\x50\x84\xe3\x40\xed\x62\x91\x57\xeb\x6a\x8d\x1b\x6d\x8a\x16\xee\xe8\x6f\x08\x32\xce\x53\x87\xa1\x40\xaf\x65\xe2\x06\x47\x07\xd0\xba\x00\xfc\x4d\xdf\x95\x53\x38\x12\x4b\xc8\x05\xcb\x0a\x25\xa0\x39\x5c\x72\x30\xb6\x33\x07\x9d\xeb\x2b\x02\x51\x5d\x73\xd8\x44\xe1\xe8\xf6\xdd\x72\x8c\xe3\xcc\xee\xdb\x49\xda\x62\xcb\x7b\x85\x65\x23\xc4\x72\xe9\xea\x93\x2c\xd0\x1c\x86\xe9\xbb\xb8\x6e\x9a\x18\x7e\xd2\xb1\x0f\xec\xe4\x8e\x7f\x9c\x6b\x61\x1d\xf1\xa0\x86\x07\xf8\x1d\xb5\xae\x06\x3b\x4f\x69\x4d\xd8\xde\xdd\xb1\x48\xb7\x37\xce\x77\x3c\xb5\xc5\xa6\xe3\x96\x32\x1f\x1f\xfe\xde\x3d\xf3\x66\x5c\x59\x47\x2f\xc1\x7b\x75\xc3\x89\x94\x64\x7d\xcd\x2e\x79\x47\x2f\x11\x8e\x69\xda\x75\x02\xf9\xc2\x12\xf4\x82\x7f\x72\xf5\xf6\xf2\x37\x1e\xd6\xde\x3d\x93\xae\x1c\x70\x9c\xbf\xf3\x2c\x71\x44\x25\xbd\xb6\x3c\xda\xf5\xca\x9d\xe9\x0b\x16\x9a\xa5\xed\x4a\xa9\xcc\x51\xdc\xcc\x39\xb0\x1d\xbc\xcb\x18\xb3\x00\xbd\xe0\xf5\xf1\x2b\xcf\x25\xc6\x72\x71\xca\x58\xa0\x64\xf5\xdd\x2a\x68\xac\xe6\xb5\x15\xe8\x05\x37\xda\x50\xe6\xe3\x7d\xb3\x1f\xb7\xa1\xba\x6c\x54\xbf\x7d\x46\xea\x36\xdc\x3b\x71\xca\x01\xbc\xd9\xb7\x9b\x04\xb9\xfb\x73\x31\xb7\xd6\x0e\xdc\xd7\x45\x3c\x11\x8b\xf4\xdd\x54\x69\xf9\x0c\x27\xe8\x25\xbc\x9b\xf9\xed\xfb\x16\xb9\xa4\x7f\x77\xf5\xc9\xaa\x53\x34\xca\x29\xf1\xe3\x12\x1c\x66\x53\x82\x69\xe1\xf6\xd4\xab\x7e\xc6\xe5\x6a\x31\x4b\x6d\x6b\x77\x7f\xeb\xfc\x2f\x15\x7b\xfd\xfe\xac\x78\x0a\x36\xa7\xf9\x9e\x32\x9f\xc8\xaf\x1c\x3b\x2d\xa9\xb4\x7b\xff\x70\x74\xb0\x74\x1d\x67\x5a\xd2\x8f\x82\xa4\xb6\x57\xbd\xf3\x68\xe2\x4e\x96\xa2\xf6\xcd\x2a\x35\x87\xa9\x46\xb3\x05\xa7\x47\x9f\xa1\x41\x87\x8f\x91\xbe\x5a\x93\xb5\xb2\xf6\x79\x98\x2f\x64\xcd\x99\xf7\x97\x2e\xde\x8d\x7a\xfe\xcd\xb2\x6e\xaa\xf1\xaa\xfb\xcc\x43\xe4\x3d\xc6\xab\x1c\xd6\x5f\x64\xc5\xc0\x33\x5f\xbd\xba\xbd\x31\x57\x8b\x6d\x3f\x3f\x22\xcc\x75\x85\xcd\xf9\xbc\x95\x7f\xd8\xb0\x3a\x5c\xe3\x99\xef\xae\xda\xb0\x71\xa7\x85\xc7\xf8\x0f\x9b\xcf\xd0\x93\x4e\xce\x17\xc2\xf9\xda\x6e\x31\xc2\x4f\x23\x3d\x9f\x7f\xe2\x7a\x65\x5c\x48\xca\x8d\x38\x43\xa4\xde\xfe\x54\xf5\x22\x77\xdc\xd5\xc4\xd1\x77\x45\x45\x4a\x87\x5a\x88\x05\x71\x53\x04\x1d\x82\x06\x57\xa8\x22\xaf\x3a\xa5\xef\x0d\xd4\xa6\x39\x3c\x83\x66\x58\x68\x36\x45\x51\xa0\x68\x50\x5f\x05\x62\x7d\x4d\xf4\x07\x7f\xe0\xc1\xab\x30\x5e\xde\xbd\x06\xb0\x72\x8d\xcf\x5f\xae\x93\xb8\xfb\x01\x7b\xfc\xe0\xc3\x1e\xad\x3c\x5e\x26\x57\xd4\x01\x39\xab\x9d\xb6\x5e\x0c\xf3\x50\xc8\xb6\xb3\x0e\x5b\x72\xdf\x7c\x98\xb1\xab\x09\xee\x07\x31\x5b\x40\x86\xd7\xfe\x36\xc4\xec\x8b\xb4\x45\xd1\x73\xd2\xe6\x27\x4f\x11\xdb\xa4\xbf\x8f\x10\x73\x02\x77\xf2\xe4\xc1\xcd\x97\x1e\x27\xec\xab\x81\xde\x9e\x50\x63\x1c\xff\xda\xbd\xa5\x04\x67\x61\x72\x79\x31\xac\x02\xe3\xc8\x49\x38\x68\x65\x47\x3a\x0c\x9b\x55\x63\xe9\xf6\x6b\xb6\x56\x0f\xc6\x11\x9e\xf8\x4c\xe6\x69\x2b\x26\xe7\x6c\x7b\xce\x04\x56\x43\xa1\xdd\xfa\xf9\x9e\x8e\x7d\x60\xb2\x97\x52\x2a\xc3\x27\xfe\x14\xd4\x1d\xb1\x62\x72\x0c\xec\xd9\x31\x5a\x3e\x8c\xf1\xfd\xee\x6d\xb2\x39\x16\x71\xa7\x17\x8c\x6e\x58\xbb\x39\xe1\x94\x39\xbd\x18\xdc\x1f\xdb\xa9\x3c\x96\x88\xf0\x81\x47\xcf\x05\x0f\x16\xbc\x17\x9d\x99\x60\xd6\x42\x52\x5c\xca\x9b\x9e\xf8\xc4\x57\x6a\x0d\xe6\xb0\x18\xc8\x41\x0f\xa4\xff\xd4\x58\xe4\xa2\xe8\xb9\xe0\x0e\x9b\x6b\xc6\x33\x18\xa4\x7f\xaf\x27\xd5\x59\x81\x43\x9b\x07\x81\x67\x57\xfc\xd4\xc0\x8b\x32\x5b\x63\xe9\x66\x0d\x16\xa3\xf7\xb6\x18\x1f\x32\x80\x9a\xcd\xc3\xaf\xaa\x36\x2c\x9a\x91\xca\x84\xe6\xd5\x08\xd4\x34\xf0\x0c\xec\x53\x99\xb9\xcb\x1c\x4b\x0e\x19\xab\x5a\xaa\xd7\xa6\x69\x3c\xb4\xd3\xde\x2e\x1d\x4f\xd7\xb1\xe3\x9c\x58\x53\x18\xde\xe4\xd2\x87\xd3\x1e\x73\x0a\xe7\xe6\xab\xd0\xb3\x5a\x30\xfc\xe5\xeb\x21\x0f\x24\xcb\xbb\x1b\x99\xe8\xdf\x20\xdc\xa3\x2d\x36\x7d\x54\x36\xa2\x7f\x88\xe3\xbb\x13\x7c\x57\x77\x05\x74\x44\x6e\x3e\x36\xa4\x47\x9b\xc3\x8a\xf4\x1a\x02\x82\x9c\x13\x9f\x1e\x0a\xf3\xea\xb0\x07\xeb\xaf\x07\xd7\x40\xc5\x24\x64\x41\x80\xfe\xa0\x36\x7f\x06\x17\x3c\x39\x27\xfd\x11\xf6\x88\xcb\xfd\xa5\x10\x77\xe5\x25\xfe\x57\xbb\x29\x71\x2e\xb4\xef\x4d\x69\x4c\xb5\x5e\x4a\xa2\x08\x8d\xbc\x3f\x2f\x3d\x34\x89\xb3\x81\xcd\xf6\x6e\xb1\x81\x3b\x97\xe0\x30\x63\x42\x59\x3d\xf3\x89\x28\xf7\x21\x76\xaf\xb6\xd8\xf4\xbe\x79\xa5\x43\xa1\xbf\xfe\x95\xf9\x6d\x1f\x84\x68\xa4\xdb\x8c\xda\xa4\x4b\xaf\x0d\x31\xfb\xa1\x05\xea\xfe\xb1\xe3\x40\x10\x44\x29\x4b\x6a\x18\x6d\x95\x2f\x7a\x53\x3a\x12\x2a\x43\x2c\xd9\xaf\x92\x5e\x2a\x35\x49\x17\xcc\xd3\xce\xc7\x7d\xd0\x7a\x48\x57\x8f\x39\x9c\x97\xbf\x94\xae\x5a\xd6\x3b\x94\x81\x06\x7b\xa4\x1f\x40\x2d\xf5\xd9\x5b\x70\x66\xdb\xcd\xbd\x4a\xc9\xe3\x13\x62\x7a\x98\xb0\x6e\x60\xc9\xf9\x8e\xc0\xb4\x2e\x29\xf8\x68\xa0\x43\x99\xff\xfb\x36\xae\x74\xa0\xf5\xb4\xa2\x77\x66\xdb\x66\x33\x42\xa6\x7a\x0c\x58\x2e\x6f\x08\xbe\x75\x5f\x6c\x78\x4b\xa6\x5e\x6a\x59\x95\xf2\xb8\xd5\x56\x91\x71\x38\x68\xf5\xa7\xa2\xae\xaa\x23\xf5\x8d\xa5\x0b\xdf\x7b\x56\x72\xbf\x02\x7b\x6b\x62\x33\xb6\xad\xad\x11\x87\x6b\x4a\xee\x31\x57\x04\xae\xdf\xa1\xb6\xd8\xf9\x58\x73\xab\x15\xf2\x77\x7f\xdb\xa6\xbb\xb4\xb0\x09\xe9\xd6\xed\x5b\x42\x5b\x02\x01\x00\xd8\x98\xe4\xe2\x9f\x14\x1e\x9a\x14\x61\x1f\x96\x18\x11\x9a\x14\x01\x58\xa3\xac\x51\x48\x2b\x2b\xa4\x95\x0d\x1a\x65\x65\x6f\x6d\x65\x8f\xb2\xb1\x40\xa1\xec\x51\x28\xcf\xca\xba\x96\x7f\x01\xe2\xb1\xe1\x31\x91\x29\xff\x1b\xb0\x57\x19\x1b\x2e\x74\x74\x77\xf1\x76\xe6\xef\x09\x21\xfe\x33\x00\x00\xff\xff\x2e\x0a\xe1\x73\xbf\x0b\x00\x00")

func init() {

	rb := bytes.NewReader(FileAssetsProgrammingFlowchartCollatePng)
	r, err := gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err := FS.OpenFile(CTX, "assets/programming/flowchart/collate.png", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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
