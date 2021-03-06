// Code generated by fileb0x at "2018-08-16 15:50:27.871562418 +0300 MSK m=+0.011123147" from config file "b0x.yaml" DO NOT EDIT.
// modification hash(8274e4bd2382466ad1ec17fdf7164458.d47d734c826bc110696e7160789432d0)

package static

import (
	"bytes"
	"compress/gzip"
	"context"
	"io"
	"net/http"
	"os"
	"path"

	"golang.org/x/net/webdav"
)

var (
	// CTX is a context for webdav vfs
	CTX = context.Background()

	// FS is a virtual memory file system
	FS = webdav.NewMemFS()

	// Handler is used to server files through a http handler
	Handler *webdav.Handler

	// HTTP is the http file system
	HTTP http.FileSystem = new(HTTPFS)
)

// HTTPFS implements http.FileSystem
type HTTPFS struct{}

// FileSwaggerJSON is "swagger.json"
var FileSwaggerJSON = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xec\x7d\x5b\x6f\x1c\x37\xb2\xff\xbb\x3f\x05\x31\xff\x3f\xe0\x04\xeb\xd1\x64\xbd\x8b\x05\x4e\xde\x1c\xcb\x49\x04\x5b\x3e\x03\x49\x4e\x16\x58\x05\x02\xd5\xcd\x99\xe1\xba\x87\xec\x90\x6c\xd9\x8a\xa0\xef\x7e\xd0\x64\xdf\xaf\xec\xdb\x4c\x4b\xae\x27\xcb\x12\x9b\x2c\xb2\x58\xbf\xba\x92\x7c\x78\x81\xd0\xc2\xe1\x4c\x06\x7b\x22\x17\x3f\xa2\xff\xbc\x40\x08\xa1\x05\xf6\x7d\x8f\x3a\x58\x51\xce\x56\xff\x95\x9c\x2d\x5e\x20\xf4\xc7\xab\xb0\xad\x2f\xb8\x1b\x38\x76\x6d\xe5\x17\xbc\xdd\x12\xb1\xf8\x11\x2d\x5e\x9f\xfc\xb0\xd0\xbf\xa3\x6c\xc3\x17\x3f\xa2\x07\xf3\xad\x4b\xa4\x23\xa8\x1f\x7e\x1b\xb6\x7a\x1f\xdc\x12\xf4\x66\x7d\x86\x24\x11\x77\xd4\x21\x88\xca\xe4\xc7\x0d\x17\xc8\xe1\x8c\x11\x27\x6c\x8d\xbe\x50\xb5\x43\x61\x7b\xc1\x88\x22\xf2\x44\xf7\x8e\xd0\x42\x51\xe5\x91\x6c\x5f\xf1\x1f\xee\x88\x90\xd1\x30\x3f\x9c\xfc\x70\xf2\xf7\x90\xcc\x47\x33\x25\xac\x76\x32\xa5\x69\xe5\x70\xb6\xa1\xdb\x3d\xf6\xd3\x5f\x22\xb4\xd8\x12\x95\xf9\x6f\x38\x12\xde\xa6\xab\x10\xfd\xee\xad\xfe\xf4\x1c\xfb\x8b\xe4\xb7\x7f\xbc\x4a\x3f\x91\xc1\x7e\x8f\xc5\x7d\x48\xc3\x2f\x44\x21\x33\x10\x0a\x47\x42\x1b\xc1\xf7\x08\x7b\x1e\x0a\x24\x11\x88\xe1\x3d\x91\x3e\x76\xd2\x89\xe9\xef\xb9\x4f\x84\x5e\xe8\x33\x37\xea\xe3\x92\x78\xc4\x51\xc4\x4d\x06\x96\xd9\xf6\x3e\x16\x78\x4f\x14\x11\x45\x42\x1f\x32\x3f\x87\x53\xb9\xf7\xf5\x9a\x49\x25\x28\xdb\x66\x7a\xd0\x7f\xdd\x70\xb1\xc7\xe1\xe4\x17\x41\x40\xdd\xe2\x5f\x43\x52\xc3\xbf\xfd\x7b\xf9\x49\x12\xb1\x3c\x3b\x2d\x36\xa0\x7a\xcd\x77\x04\xbb\x44\x14\xff\x26\xc8\x9f\x01\x15\x24\x9c\x8d\x12\x01\xc9\xfc\xf1\xf1\x55\x3d\xb9\x84\x05\xfb\xc2\x84\xf4\xef\xc3\xa5\x2b\x8c\x10\xee\x4f\x77\x4f\xd9\x22\xf7\xdb\x3f\x5e\x75\x99\x7f\x61\x86\x17\xdc\x23\xd3\xcf\xb1\x13\x49\x1f\xe3\xed\x32\x12\x5d\x95\x9b\x57\x10\xe9\x73\x26\x89\xcc\x89\x01\x42\x8b\xd7\x3f\xfc\x50\xf8\x55\x59\xb2\xb3\x5b\xdd\xa3\x52\xe5\xf7\xbb\xcc\x6c\xf8\x22\x99\xd2\xd9\x91\x3d\x2e\x0d\x80\xd0\xe2\xff\x0b\xb2\x09\xfb\xfe\x7f\x2b\x97\x6c\x28\xa3\xe1\x58\x72\x55\x96\x88\x0f\x54\xaa\x3c\xff\x1f\xeb\x98\xb0\x70\xc9\x06\x07\x9e\x6a\x9f\xcf\x8e\x08\x71\x8f\x88\x10\x5c\x0c\xa5\x78\x5b\x25\xc6\xa7\x86\x90\x9f\xb8\x7b\xdf\x40\xfb\x8b\x8a\x59\x2c\xbe\x2e\xf7\x44\xed\xb8\xbb\xbc\xa3\x92\xde\x52\x8f\x2a\x8d\x38\x7e\x70\xeb\x51\x27\xee\xcc\x7c\x1a\x7d\xb6\x58\x51\xb6\x15\x44\xe6\x99\x6b\x07\x79\x67\xe6\x4b\x2b\xc0\x4b\x46\x19\x0a\x77\x67\x09\xb9\x80\x76\x80\x76\xb3\x43\xbb\x74\x9f\x1f\x06\xeb\x12\x71\x00\xa8\x33\x50\x27\xe8\x1d\x56\xa4\x06\xeb\x32\xeb\xdf\x15\xec\xd2\xad\x67\x03\x77\xe9\x40\x7a\x1f\xb4\x60\x5b\xd2\xb7\x66\x22\xe0\x1a\xe0\xda\xb1\xe9\xe2\x5f\x58\x79\x40\x43\xcc\x9f\x01\x11\x39\x69\x9d\x1c\x45\x87\x42\x4f\xb2\xba\x00\x92\x95\xf6\x60\xfc\xc1\xc2\xe7\x72\x54\x24\x7c\x2b\x08\x56\x24\x05\xc3\x26\x18\x34\x6d\xab\x04\x01\x50\x10\x50\x70\x6c\xba\xe2\x91\x6f\x43\xa1\xaa\x1c\xad\xea\x2f\xbd\x81\xc7\x46\x6e\x2d\xc1\xf3\xef\xad\x80\x94\x88\x1b\x72\xb4\x4c\xb9\x87\x98\xc6\xb7\x6a\x5f\x26\xd8\xe9\x12\x8f\x28\x32\x26\x7a\x9e\xea\x1e\xbb\xb8\xc9\xe6\x8b\x50\x96\x3e\x56\x39\x1a\x00\xa4\x33\x07\xd2\xae\x58\xf0\xda\x1e\x0b\x24\x32\x1b\xd4\x5d\x80\xd4\xd6\x7b\x85\xab\x87\xe4\xe7\xc7\x43\x79\x88\xb6\xbe\x21\x08\x32\x58\x44\x47\xa7\x8b\x35\x13\xe4\x63\xb5\x3b\x6a\xf8\xad\x96\x3e\xb0\x79\x26\x70\x17\x83\x51\x51\xf1\x93\xef\x5a\x7b\x8b\xa6\x2d\x60\x23\x60\xe3\x73\xc5\xc6\x6f\xde\x79\x0d\xb4\x88\x83\xf3\xfa\xa4\x9d\x57\x2b\x30\x37\x6d\x01\xcc\x01\xcc\xc1\xd0\xed\xec\xd8\x83\x5f\xdf\xd1\xaf\x3f\x4e\x91\x9f\x45\x16\x38\xe9\x1c\xb2\xc0\x00\x82\x00\x82\x3d\x4a\x0b\x87\xa2\x16\x14\x0e\x8e\x9a\x28\xb6\x47\xcb\x28\x51\x9c\xf2\xb3\x3d\x53\x9c\x76\x0e\x48\x09\x48\x09\xbe\xff\xa8\xbe\x7f\x85\xe4\x1e\xc0\xf7\x4f\xc5\x7f\xac\xcc\xb5\xcd\x3c\xa0\x08\xdc\xc6\x54\x5e\x3d\x24\x3f\x3f\x1e\xcc\x6c\xb6\x35\x98\x41\x05\x80\x0a\xf8\x76\x54\x80\x1d\x39\x89\xb8\xce\xde\x76\x07\x90\x9f\x49\xaa\xce\x1e\xa6\xa3\x54\x9d\x1d\x52\x9b\xc6\x00\xd6\x00\xd6\x00\xd6\x07\x01\xeb\x67\xe9\x3e\xbc\xee\xe2\x3e\x8c\x94\x3b\x04\xcd\xd2\xae\x59\x6c\x53\x87\xf6\xca\x25\x4a\x1d\xda\x29\x17\xd3\x18\x94\x0b\x28\x17\x50\x2e\xb3\xf4\x04\x3a\xe1\xf6\x37\x9d\xcb\xec\x11\x9f\x71\x89\xef\xf1\xfb\x3d\x61\xaa\x7b\x2e\xf3\x34\xf9\xd6\x2a\x2a\x93\x19\xca\x26\x99\x99\xf6\x0e\xd9\x4c\x80\x65\x80\xe5\x32\x39\x47\x3e\x62\x5b\x94\xe7\xa1\x28\x99\x0a\x3c\x24\x4f\xc7\x48\x9e\x76\x80\xe7\x28\x7b\x9a\x72\xb4\x3d\x7b\x9a\xe9\x1d\xa0\x19\xa0\x19\xd2\xa7\xa3\xc6\x3f\xaa\x64\xf7\x00\xf9\xd3\x14\x00\xc6\xca\x9f\x5a\x4d\x04\x8a\x0d\xed\x2c\xf4\xd5\x43\xfa\x9f\xc7\xc3\x99\xeb\xd6\x86\x3a\x68\x02\xd0\x04\x60\xa4\xe7\xc9\x71\xab\xa4\x63\x4e\x69\xd4\x7a\x02\x01\xeb\x8f\x95\x47\xed\x00\xd5\x51\x22\xd5\x0e\xad\x4d\x63\x00\x6c\x00\x6c\x00\xec\x03\x01\xf6\xf3\x74\x25\x5e\x77\x71\x25\x46\xca\xa5\x82\x7a\x19\xf3\x20\x66\x07\x15\x13\xa5\x53\xed\x54\x8c\x69\x0c\x2a\x06\x54\x0c\xa8\x98\x99\xfa\x04\x9d\xc0\x1b\x0e\x87\x8e\x14\xaf\x59\xd1\x3d\xde\x66\x81\x79\x2a\x57\x40\x8f\x83\x28\xcb\x25\x5c\x1d\xce\x14\xa6\x8c\x88\x2e\xee\xc1\x99\xa6\x18\x00\x1c\x00\x1c\x00\xfc\x59\xfb\x08\x46\xf0\x8d\xb8\x83\x93\x00\xfa\xab\x52\x7f\xf9\xdc\xed\x5e\x23\xb4\xe6\x6e\xc7\x6c\x03\x0a\xc7\xe9\x56\x20\xb4\xe6\x2e\xd4\x08\x81\xa6\x02\x4d\xf5\xa4\xd3\x0f\xa1\xe0\x8f\x52\x48\xb4\xe6\x2e\x54\x10\x8d\x52\x12\x9a\x57\x00\x82\xe8\xc7\xfa\xe4\x01\x7c\x98\xac\xeb\x12\x0f\x8b\x1c\x1e\x74\x4b\x6f\x5c\xc4\x04\x83\x62\x00\xc5\x00\x8a\xe1\x1b\x70\x61\x12\x89\x07\x2f\xe6\xdb\xf6\x62\x08\x73\x7d\x4e\xfb\x9c\x6a\x78\x17\x7d\x69\xe5\xb6\x24\xc3\xd8\x38\x2c\x71\xcf\xe0\xab\x80\x4a\x7a\x86\x2a\x69\x7c\xdf\x20\x2f\x5d\x43\xe1\x2a\x16\x3f\xf0\x0d\x5a\x52\xc8\x76\xc7\x0b\xac\x71\x32\x3a\x5c\x10\x33\xb3\xfd\x68\x41\xd2\x33\x60\x24\x60\x24\x1c\x2c\x18\xd5\x4c\x2e\x4b\xed\x01\x8e\x15\xc4\xa2\x3f\xd6\xa1\x02\x8b\x49\x80\x71\x6c\x63\x1c\xaf\x1e\xe2\x1f\x1f\x0f\x63\x27\x5b\x5a\xc8\x80\xfc\x80\xfc\x10\xb0\xc9\x93\x43\xca\xb2\x31\x47\x5b\x1d\xb0\xfd\x00\x06\x7a\x30\x26\x3e\x47\x21\x77\x1b\x88\x36\x4d\x01\xa5\x01\xa5\x01\xa5\x0f\x80\xd2\xdf\xb8\xb7\x30\x52\x30\x1d\x34\xca\x78\xa7\x06\xac\x95\x4a\x74\x66\xc0\x46\xa9\x98\xa6\xa0\x54\x40\xa9\x80\x52\x99\xa1\xe9\xff\xda\x1e\xb0\xe1\xac\x40\xc7\x40\x0c\x65\x5b\x41\x64\x7e\xd9\xed\xa2\x2f\x67\xe6\x4b\xab\xe0\x4b\x32\x8a\x4d\x92\x32\xea\x18\x72\x94\x00\xc5\x90\xa3\xb4\x88\x7b\xe4\x85\x6b\x28\x58\x9d\xc5\xbd\x41\x8e\x72\x84\x1b\xd0\x6c\x41\x32\xca\x50\x46\xac\x6c\x4f\x50\xc6\xfd\x02\x3e\x02\x3e\x42\x7e\x72\xd4\x88\x43\x49\x64\x0f\x10\x70\x88\xe4\x7e\xac\xec\x64\xfb\x14\xc0\x26\xb6\xb1\x89\x57\x0f\xd1\x8f\x8f\xb3\xb3\x8e\x01\xf9\x01\xf9\x21\x48\x91\x27\x87\x96\x44\x63\x96\x66\x3a\xa0\xfb\x2c\x6e\x38\xb3\x45\xe8\xf8\x4e\x83\x76\xd3\x3c\x3a\xc3\x0c\x00\x0d\x00\x0d\x00\x3d\x35\x40\x7f\xd3\x7e\xc2\x48\x79\x49\xd0\x24\xa3\xa5\x25\x6d\xb5\x49\x94\x95\xb4\xd0\x26\xa6\x25\x68\x13\xd0\x26\xa0\x4d\x66\x67\xee\xbf\xb6\x86\x6a\xc8\x48\x76\x8c\xbe\x4c\x7a\xc9\x8b\xed\xcd\x2e\x70\x9f\x0b\xc0\x2e\xc0\x6e\x25\x39\x47\x7e\xf3\x09\xae\x68\x99\xc5\x15\x2d\x21\x90\xae\x1e\x7c\xee\x3e\x4e\x89\xd5\xed\x28\x0d\x08\x0d\x08\x0d\x08\x9d\x27\xc7\xe7\xee\x4c\x63\xe0\x15\x94\xf5\xc0\x6d\x80\xec\xc1\x2f\x59\xdb\x60\x70\x14\xaa\x68\x81\x61\xd3\x0a\x90\x18\x90\x18\x90\x78\x36\x48\xfc\xda\xca\x8c\x86\x97\xaa\xfb\xda\xbc\x2b\x8f\x6f\xa7\xb4\x7b\x91\xc7\xb7\xd2\x22\x44\xc1\xb7\x10\x19\x06\xd8\x3d\x1e\x5d\xe9\x56\x69\xa6\xfb\x93\xbf\x15\xd8\x7d\x32\xe4\xbe\xe5\x8c\x11\x47\xa3\xdb\x13\xa1\xf8\x92\x38\xcb\xdf\xc9\xed\x25\x77\x3e\x13\xb5\x7c\x4f\xee\x0f\x47\x38\x65\x8a\x6c\xcb\x9d\xa6\x94\x57\x1c\x1b\x2a\x90\x2d\x0d\xd9\xbf\x11\x21\x67\xb0\xe6\xcf\xdc\x96\x18\x4c\xc8\x86\x7b\x1e\xff\x62\x1b\x01\x1c\x3c\x9c\xc2\xd4\x3b\xd8\x60\xc9\x23\x25\x07\x1b\xd1\x17\xe4\x8e\xf2\x40\x4e\x13\x51\xfd\xbb\x45\x51\x41\x6c\x6f\x80\x1d\x68\x69\x07\x4a\xe2\x08\xd2\xe3\x5e\xcf\x4b\xfd\x9d\x95\x0d\x78\xf5\xe1\x12\x45\xc3\xd8\xa4\xab\x4c\xcf\x90\xb1\x02\x73\x10\x34\x67\x99\x1c\x37\x34\x2f\x8e\x98\xb2\xca\x0a\xf2\x50\x58\x34\x92\x0e\x99\xab\xc1\xe8\xbd\x8a\x76\x45\xf6\x25\x01\xab\x23\x7d\x96\x28\x1e\x9d\xe8\x33\xa3\xb5\x1f\xe8\x3b\xd5\xd4\x44\x7d\x03\x86\x03\x86\x7f\x33\x18\x7e\x98\x6a\xdd\xa2\xd4\x1e\xa0\x58\xd7\x88\xfe\x58\x67\xfa\x5a\x27\x00\x45\x65\x16\xa0\xaf\x3c\x39\x17\xc4\xbf\xfa\x70\x09\x70\x0f\x70\x0f\x70\x0f\x70\x0f\x70\x3f\x11\xdc\x3f\x98\x1f\x1e\xa7\x0c\xd5\xb4\xe3\x7d\x12\xa0\x01\xa4\x07\xa4\x87\xe0\x4c\x9e\x1c\x59\x94\x8c\x39\xd5\xab\x55\x13\x07\x58\x3e\x72\xd5\x9a\xd5\x79\x6d\x4b\x48\x8e\x8e\x6b\xb7\xa3\xb2\x69\x08\xc0\x0c\xc0\x0c\xc0\x3c\x31\x30\x1f\x38\x82\xff\xfc\xbc\x8f\xd7\xb6\xde\xc7\x48\x07\xc3\x41\x63\x8d\x75\x2e\xdc\x52\x6b\x45\xb5\xd6\xed\x5a\xcb\x34\x04\xad\x05\x5a\x0b\xb4\xd6\xcc\xdc\x09\x6b\x90\x86\xba\xeb\xae\xd1\x1c\x71\x47\x1d\xd2\xa7\xe0\x46\x7f\x68\x19\xc6\x31\x83\xd8\x95\xdb\xe8\xb6\x50\x6f\x03\x18\xfc\x0c\x31\x78\x8a\x30\x4a\x46\xb6\x86\xdb\xa6\xa6\x33\x28\x81\x19\xe1\x7e\x6a\x5b\x84\x4c\x72\x9b\xba\x79\x7b\x72\x33\xee\x17\xd0\x11\xd0\x11\x52\x9b\x23\x07\x17\xb4\x68\xfd\x4e\xd5\x6e\x1d\x8a\xd4\x81\x93\x9c\x7a\xf0\xf1\xb2\x9c\xd6\x73\x81\x7c\xa7\x85\x85\xbc\x7a\x88\x7e\x7a\x9c\x9b\xad\x0c\x9a\x00\x34\x01\xc4\x2a\x8a\xb1\x8a\xa2\x68\xcc\xd0\x68\x07\x88\x9f\x5b\x1a\xd4\x0e\xa7\x93\x3c\x68\xab\xc5\x1e\x27\x42\x01\xa7\x01\xa7\x01\xa7\x27\xc6\xe9\xe7\xec\x3e\xbc\xb6\x76\x1f\x46\x4b\x53\x82\x6e\x19\x3f\x61\x69\xa7\x5f\x92\x8c\x65\xab\x7e\x89\x53\x96\xa0\x5f\x40\xbf\x80\x7e\x99\x99\x1f\x60\x8f\xd9\x90\xb5\xec\x18\x93\xe1\x5e\x60\x88\x7c\x88\x7f\x7c\x5c\xb9\xc4\xf7\xf8\xfd\x9e\xb0\x1e\xd7\x07\x9c\x26\xdf\xda\x05\x69\xa2\x41\x51\x66\x4c\x9b\x80\x4d\x3a\xcc\x65\xd4\x03\xe4\x39\x01\xb7\x01\xb7\x2b\x70\x3b\x92\x8f\x99\x06\x70\x8a\x72\x3f\x14\x6b\x53\x60\x80\xdc\xeb\xc8\x6a\xe1\x30\x15\x2e\xb1\x42\xe8\x11\xbe\x07\x55\x00\xaa\x00\x54\xc1\x93\x55\x05\x50\x80\x33\xdf\xab\xb8\xbb\x06\x5c\x8a\x28\x6e\x1d\x7a\x91\x97\xe5\x4d\x0a\x00\x0e\x00\x0e\x00\x7e\x6c\x00\x7f\x6d\x0f\xe0\x10\x85\xe9\x65\x6e\x4b\x96\x31\xb7\x6b\x82\x30\xb6\x80\xdd\x21\x0e\x53\xc4\xec\xcc\xc8\xed\xb0\x9d\xf1\xb7\x00\xb9\x01\xb9\x01\xb9\x9f\x24\x72\x67\xa3\x30\xf0\x16\x64\xc7\xba\xc6\x3b\xee\x05\x7b\xd2\xfd\xc6\xae\xdf\xcc\x77\x96\x55\xed\x66\x94\xf6\xa2\xf6\xa8\x57\xc0\x60\xc0\x60\xa8\x69\x1f\xb5\x28\xa5\x28\xaf\x07\xa8\x64\x37\x42\x3f\x56\x21\x7b\xeb\x04\x00\xe6\x5b\x61\x7e\xf5\x60\xfe\x7d\xec\x63\x94\x5b\x22\x7e\x64\x90\xb7\x23\xbe\x69\x08\x88\x0f\x88\x0f\x56\x77\x15\x39\x77\x45\xc9\x98\x93\xcd\x1d\x61\x3b\xc4\x4a\xfa\xc0\x70\xf7\x24\xa4\x25\xf6\xfe\x42\x54\x04\xbc\x56\x99\x47\xd3\x2b\x24\x1c\x01\x7f\x9f\x21\xfe\x8e\x9f\xe2\xcb\x4a\xd6\x38\xd6\x2c\x24\xf8\x06\x43\x69\xa5\x49\x7b\x74\x4c\x05\x3c\x05\x3c\x05\x7b\x76\x36\xf6\xac\x2d\xb6\x43\x8c\x62\x0e\xe7\x2f\x2d\x81\x39\x3a\x7e\xd9\x1e\x68\x30\x0d\x01\x98\x01\x98\x01\x98\x27\x06\xe6\xe7\x17\xe6\xb6\x0e\x85\x8c\x74\xe0\x12\x54\x88\x55\x98\x5b\x2a\x2e\xf0\x96\x4c\x5b\xcd\x6d\xc6\xb0\xaa\xe1\x36\x4d\x21\x94\xf2\xb4\x34\xcc\x04\x65\xc0\x99\x3d\x33\xb8\x0a\x38\xb3\xa9\x00\x0f\x2a\x63\x04\x2f\xa2\x4f\x17\x19\x82\x12\xca\x17\x6f\x1c\x87\x48\xf9\x81\xdc\x11\x2f\x0b\x14\x35\xbb\x6d\xf1\x75\xb9\xe5\x4b\x1f\x3b\x9f\x0d\xb0\x2c\xb6\x54\x9d\x24\xaf\x55\x07\xfb\x13\x46\xd4\xca\xd9\xad\x3e\x07\xb7\x64\x89\x7d\xba\xba\x23\xcc\xe5\x62\xb5\xa5\x6a\x17\xdc\x9e\x38\x7c\xbf\xca\xb4\x36\xcd\x1c\x8f\x12\xa6\x56\xfe\xe7\xed\x6a\xcf\x5d\xe2\x2d\x72\x28\x66\x96\xeb\x67\xea\x91\x2a\xf2\xf8\xed\x7f\x89\x93\x6e\xa2\x85\x2f\x42\xdc\x51\xb4\xb0\x31\x17\x8e\xee\xe5\x66\x93\xef\xa6\x4d\xae\xcc\x6c\x63\xd1\x0a\x4d\xa4\x45\x25\x37\xa2\x2e\xfa\x75\x7b\x15\x36\x7c\x51\x64\xf7\xe3\xb1\x57\xfc\x1c\xfb\xf9\x04\x68\x4e\x36\x92\x36\x68\xb9\x44\xba\x0f\xb4\xe1\x02\x99\x65\x46\x7b\xec\xa7\x2c\xa9\x63\x55\x06\x72\x52\x90\x34\x38\x96\x59\x5a\x17\x2b\x1c\xaf\xce\x1f\xed\x6c\x36\x69\xf4\x1b\x5c\x94\xee\x92\x6c\x87\x0d\x75\x29\x64\xe8\x1e\x51\x86\x2e\x7e\x7e\xfb\x8f\x7f\xfc\xe3\x7f\x50\xa4\x47\x5e\xf5\x62\xa5\xa9\xcf\x71\xdf\xa8\xea\x6d\xa2\xe7\x52\xa0\xab\x1a\x2e\x92\xe5\x3d\xcd\x4c\xbf\xd0\x99\xc9\x2a\xb5\x4f\xd5\x34\x1c\x77\xa2\x26\x2d\x5d\x3b\xd1\xa8\xd9\xd8\x62\x96\xba\x04\x43\xfa\x36\x3d\x54\x0e\xc0\xbf\xb0\xdc\x53\xbc\xdd\x3a\xff\x5f\xfd\xf5\x2c\x25\xf9\x34\xbf\xf3\x6a\xa5\x39\x6c\x57\x27\xd1\x48\x6f\xdf\x56\xb1\xc6\xae\xab\x37\x31\xf6\xd6\x35\x42\x5a\x58\xd0\x19\x2d\x93\x49\x35\x58\xac\x93\x6e\x58\xb7\x50\x85\xe4\x47\x4f\x5d\x15\x76\x54\xb7\x15\xb1\x10\x38\xef\x05\x2e\xa8\x22\x7b\x59\x36\x69\x5a\xe0\xa5\x3e\x27\x9d\xc7\xb5\x64\xde\x73\xdb\xde\xa6\x79\x33\xcb\x4c\x9b\x12\xb7\xa2\x5f\xd3\x6c\x39\x7c\x5f\xad\x45\xf7\xe1\x9c\x5f\xd5\xab\x31\x8f\xee\xa9\x92\x1d\x14\x19\xdf\xef\x31\x73\x47\xd8\x01\xd5\xd2\xd6\xca\xf0\x68\xf8\x4a\x98\x8c\x8c\xa9\x89\x77\xa8\xe1\x4f\xd9\xd5\xef\xba\x4f\x33\x84\x13\x76\x37\x15\xc1\xef\xd8\x9d\x25\x91\xb9\x96\x59\xea\xcc\x26\xea\xab\x7d\xce\xf4\xd7\x95\x1d\x47\x9b\xcf\xca\xf2\xb8\x20\x92\x07\xc2\x21\x07\x55\xec\x3e\x17\x6a\xfa\xad\xb4\xe6\x42\x59\xf2\x68\xad\x09\xaa\xa4\xd5\x44\xb2\x6e\xf6\x3c\x60\x6a\x66\xdb\xdf\x34\x3e\x37\x94\xcd\x12\xa8\x35\x0b\x2c\xc0\x3a\x6c\x97\x07\xec\x70\x8b\x84\x58\x9d\x8c\x3c\x96\x83\x11\x76\x9c\xfb\xbf\xe0\x8a\x3b\xdc\xb3\xc7\xea\x09\x85\xa2\xae\x5b\xca\x14\xd9\x16\xa2\x38\x99\x10\x18\x65\xea\x5f\xff\x5c\x34\x6f\xef\x9a\x41\xe3\xd9\x5b\x81\xc5\xba\xb0\x56\xf3\xda\x6c\xbf\x95\x4e\x6d\xd4\x6d\x37\xd3\x32\xdc\x70\x51\x9c\xfa\xbb\x9c\x41\xf7\x3d\xd2\xd2\x4e\xdc\x51\x36\xa0\xee\xeb\x46\xe7\x0c\xac\xb7\x58\x38\xcd\xde\x5b\xec\x3c\xfc\xb8\x92\xdb\x19\x52\xfa\x77\x1e\x30\xb5\xce\x4c\xe6\x20\x0a\x43\x06\xb7\xc3\xc8\xbe\x0c\x6e\x0b\x44\xcf\x60\xe3\x66\x4f\xf7\x34\xec\xda\x4c\x33\xb4\x5c\xf6\xf7\x33\x3c\x7c\x4b\xbc\xde\x2b\xf8\x41\x7f\x5d\xcd\x9d\x52\xde\xa1\x23\x73\xa2\xef\xab\x63\x6d\x58\xd0\xcd\xe6\x86\xba\xfd\x03\x6e\xba\x87\xb3\xd3\x59\x31\x3f\x73\x9a\xb7\x9e\xf5\x69\xa3\xbc\x72\xcc\x9e\x24\xee\x89\x4a\x09\xb9\xb2\xc9\x8d\x11\xc4\xf7\xa8\x83\x3b\x38\x32\xd8\x51\xf4\xae\x76\x2b\xdc\x72\xee\x11\xcc\x1a\x98\xf5\xc6\x7c\x5f\xe7\x87\xc4\x34\x4f\x6d\x87\xd9\x3b\x20\x31\x45\xd5\x14\xcf\x3b\x3e\x39\xe3\x90\xa2\x76\x8e\x6e\xfc\xc0\xf3\x6e\xa2\xb7\xf5\x8e\xe3\x19\x6b\x37\x6b\x1d\x78\x9e\x79\xfb\x52\x42\x00\xb4\x3e\x00\x5a\x05\x1c\xd3\x18\xb4\x17\x05\x5c\x2a\xe8\xa3\xe8\x68\x76\x85\xd2\x28\x6c\xec\xf8\x82\x03\x74\x76\x8a\xbe\xe3\xcc\xbb\x47\x74\x93\x41\x57\x44\x25\xf2\xb1\x50\x88\x6f\x92\xcb\x14\xbe\xef\xb9\xdf\xe3\xa1\x72\x7a\x28\xa7\x44\xb1\x0a\x2c\x9d\xf6\xcc\x5d\x99\xe6\xab\x6a\xcd\xc9\x15\xf6\x6e\x1c\x3f\x68\x59\x04\xdd\x0e\xbd\x5d\x7f\x42\x81\xc4\x5b\x82\x6e\xef\x11\xf6\xbc\xd4\xf0\x95\xa1\xcc\xab\x1d\x95\x55\x81\xb3\x2e\x5c\x0d\xda\xd8\x7a\x15\x52\xf2\x76\xfd\xa9\x69\x3e\x7b\xb2\xe7\xba\x1c\xa1\x7d\x4a\x17\x6f\xce\xe7\x31\xa5\x73\x43\x73\x75\x7c\x81\x08\x69\xa8\xb6\x60\xfc\x6f\x51\xe3\x79\x5a\x32\x97\xc5\x2d\x5c\x6f\xcf\x98\xa6\xa1\x55\x13\x8e\x20\x18\x51\x44\x22\x23\x02\xa1\xb0\x75\x89\xd0\xd6\x5a\x22\x77\x98\x7a\xf8\xd6\x23\x37\x13\x83\xd1\x9b\x78\xa0\x66\x54\x12\x04\xbb\xf7\x37\x93\x03\x23\x76\xef\xdb\xe8\x38\x22\x34\x07\xec\x60\x6c\xf9\x94\x0e\xd5\x42\x93\x29\x52\x9b\x9c\x1e\x33\x4c\x05\x2d\x73\x92\xe1\xdf\x4a\x70\x54\x2f\xc4\x51\xdb\x3a\xdf\x04\x45\xd0\x16\x55\x01\xf6\x97\xe4\x1a\x88\xb4\x56\xbd\xf3\x46\xcd\xb6\x24\x64\xa1\x65\xad\x27\x38\x30\x0b\x59\x7d\x3b\xd5\xa8\xde\x55\xd5\x0d\x56\x8d\xc6\x77\x66\xea\xb3\x62\xde\x3b\xe6\xfa\x9c\x36\xba\xee\x71\x93\x3c\xbb\x48\xfc\xe1\x48\xc1\x6c\xec\xba\x82\x48\x49\xe4\xa2\x94\x5d\xb1\xf7\xd8\x93\x3e\x8e\xe3\x61\xbd\x49\xc6\x7f\x82\x2e\xf4\x44\x8e\xdf\x64\x7e\xd9\xa4\x99\xb7\x01\x09\xb7\xc1\x32\xdd\x22\xa9\x6d\x20\x9b\x6b\x57\x2d\xb3\x43\x01\x36\xe9\x67\xba\x94\x74\x04\x2e\xb6\x79\xe9\x98\x9e\x43\xb0\xe1\xae\x71\xf1\xef\xb4\x27\x40\xee\x97\x77\xd8\x0b\x08\xf2\x31\x15\xa1\x1b\x40\xd8\x1d\x15\x9c\x19\x63\x02\x0b\x1a\xda\x72\xbd\x23\x9e\xba\xeb\x52\xb0\xf3\xd8\x59\x3f\x43\x55\x6f\xcb\x46\x7f\x3d\x2b\xd5\x28\x9a\xaa\x73\xde\x09\x5d\x97\x23\x15\x66\x2e\x16\x2e\x92\x44\x50\xec\xd1\xbf\x42\xc6\xa2\x37\xeb\x33\x53\x7c\x7d\xcd\xce\x89\xd4\x7e\xfb\x72\x19\xfa\xec\x61\x73\x65\xfe\x84\xf6\xe6\x2f\x3f\x5e\xb3\xbf\xa1\xeb\x05\x65\x77\xd8\xa3\x2e\x0a\x24\x11\xe1\x9a\x5c\x2f\xcc\xef\xff\x0c\xb8\xc2\x88\x7c\x75\x08\x71\x89\x1b\xff\x56\xb7\x35\x3a\xc2\x8c\xb3\xb8\x66\x27\x27\x27\x44\x39\x27\x27\x27\xd7\xec\xec\x34\x1c\x2f\x60\xf4\xcf\x80\x44\xa3\x51\x97\x30\x45\x37\xd4\x31\x5f\x39\xdc\x25\xd7\xec\x94\x28\x4c\x3d\xed\xbc\x72\xdf\x94\xbe\xe9\xc8\x02\xf9\x5a\x20\x52\xa2\xcf\x94\xb9\xd8\x0c\xbe\xa1\xc4\x73\xd1\xcb\xd8\xfc\x7f\x89\xf6\x81\x54\xe8\x96\x20\xc6\xd9\xf2\x2f\x22\x38\xd2\x5b\x21\xa6\x95\x71\x85\x08\xe3\xc1\x76\x87\x14\xdd\xee\x94\x44\x8a\xa3\x0d\x21\x2e\xda\x72\x7f\x47\x44\xdc\x4e\x44\xf5\x1b\xe8\xe5\x2f\xdc\x7d\x89\x5c\x4e\xe4\x4b\x85\xc8\x57\x2a\x55\xd8\xe4\xe7\x70\xd4\x3c\xa9\x92\xe8\x78\x56\x5e\xda\xe4\x10\xc3\x51\x2f\xc7\x91\xac\x87\x88\x19\xd5\xc2\xa5\xd7\xdc\x32\xa6\x66\x56\xaa\x26\x18\xed\xda\xf5\xf1\x4e\x88\xba\xf0\x5e\xb4\x25\xfa\xe7\x5d\xa3\xef\x1b\x62\x87\x37\x3b\xa5\xfc\x89\x9c\x57\x13\xaf\xf9\xf5\xea\x6a\x7d\x38\xb4\x31\xe7\x31\x4a\xf0\x72\x76\xda\x0c\x30\x46\x8c\x05\xf1\x05\x91\xda\x2f\xca\x49\x74\xe6\x28\x50\xe7\x9d\x1e\x4a\xb3\xf5\x46\x78\x1f\x36\xae\xe6\x56\x87\xed\x74\x79\xc8\xb4\x65\xcd\x82\xbf\xcf\x4f\xbb\x62\xc9\xc3\x16\x85\x45\x0f\xd7\x4a\xab\xf0\xdc\x69\x9a\xda\xed\x58\x1f\x52\x3d\xcc\x1c\x2f\xdb\x76\xd5\x65\x69\x5b\xc5\x0f\x94\x9d\x9d\x36\xcc\x73\xda\xe3\x34\x55\x93\xf9\xb9\x88\x79\xc5\xc9\xa4\x2a\x21\x33\x99\xb4\x8a\x3b\xa3\x18\x0c\x7c\x36\xcc\x6e\xee\xb5\xe0\x55\xeb\x73\x56\xa8\xba\xec\x8c\x02\x1f\xdb\x0d\xc1\x4a\xb1\xbf\xc2\x5b\xeb\xcf\x8e\x6f\xc7\x9d\xb1\x6d\xe8\x95\x37\xec\xa3\xa8\x45\xde\x59\xa2\xd1\x67\x23\xc5\x37\x44\xe0\x11\xf9\x6c\x8e\x03\xc1\x09\x9e\x27\x94\xc0\xd6\x3b\x6f\x22\x4f\xfd\x22\xf0\x6c\x6b\x7c\x2f\xb2\x12\x30\x27\x64\x20\x6d\x21\x95\x5c\xbb\x4a\x94\x20\x43\x43\x2a\x49\x3f\x53\x31\x2a\x46\x41\xcb\x5a\x91\x62\xeb\x19\x70\xeb\x63\x85\x04\x16\x39\x95\xb4\x09\xb9\x94\x88\x6c\x6a\x1e\xe0\xdc\x3d\xf7\x1d\x21\x3d\x76\x50\x3b\xd5\x90\x39\x79\xd5\xd3\xc0\xa1\xec\x71\xe6\x27\x18\x3f\x1e\x50\xd6\x58\xe7\x69\x56\x96\x79\x16\xe6\x1a\x48\x22\xcc\x51\x72\x8f\x20\xfd\x81\x96\x4b\xb5\x23\xa8\xfa\x1e\x96\x51\xea\x44\xf7\xf8\xeb\x0d\xf9\xaa\x6e\xe2\x87\x87\xfb\xb9\xaa\xad\x25\x17\xe7\xf8\xeb\xbb\xaf\xaa\x74\xa9\x45\x91\x12\xca\x0e\x41\xc9\x19\x6b\xa7\x44\x09\xbc\xd9\x50\x67\x42\x2a\xae\xa2\x11\x0e\xab\x43\x75\xc7\x37\x1e\xdf\x52\x36\xac\xfb\x0f\xba\x8b\x9a\x6a\x86\x18\x5e\x3a\x1d\x77\x92\x07\x2c\x36\xce\x96\x1a\xc8\x09\x8b\x57\x3f\x49\x22\x0c\x1c\x5a\xea\xab\x4f\x32\x5f\xb8\x3a\x23\x6d\xd5\x66\x5c\xe4\x1b\xe6\xad\x8b\xf4\x6e\xcc\x81\xe6\x45\xda\xd1\x54\x2c\xab\xb0\x6e\x1b\x39\x96\x4e\x7b\x56\x6c\x5b\x5f\xb4\x97\x8d\x24\x6d\x72\x61\x07\xb4\x16\xe4\x82\x78\x04\x4b\x82\xe2\x3e\x7a\xf3\xeb\x4c\x7e\xd4\x77\xf0\x34\x96\x9e\x57\xca\x65\x34\x74\xc3\xe7\x76\x20\xdc\xd4\xf9\xa5\x12\xc7\xf4\xff\x6f\x3d\xcc\xb6\x2b\x49\xf6\x77\x31\x5c\x27\xdc\xcb\x1f\xac\x29\x31\x0e\xab\x5d\x28\x60\x91\xb1\x8d\x72\xb7\x9f\x75\x34\x06\x0b\x37\xa7\xc5\xef\xcb\xdd\x14\xfd\xfe\xf8\xf7\x7e\x26\xa3\xdc\x6e\x37\x0e\x3a\x21\x54\x7f\xa6\x29\x47\x64\xef\x33\x2e\xa6\x93\x86\x23\x4e\xd9\x29\x4f\x14\xc4\x37\x43\x14\x4e\x03\xce\x00\x3e\x42\x01\x91\x8a\xb0\xe8\xa2\x60\xa3\xc4\xce\xf3\x67\xe0\xe6\x73\x43\xd1\x9a\x37\x85\xc3\xd7\xdc\x2d\x9e\x64\x75\x07\xdd\x0d\x01\x47\x5e\x46\x8d\xc1\xf9\x1e\xbf\xef\x2d\xc5\xa6\x20\x0d\x4e\xac\x1c\xbe\x70\xa9\xcb\xf1\x88\x35\x77\x0f\x7b\x2e\x22\x2b\xe2\x4f\xfd\x40\xc4\xb4\x73\x29\x9d\x84\x98\x81\xf6\x49\x76\x4b\x23\xa8\xb7\x9c\x58\x18\x84\xf2\xfe\x0e\xcb\xfe\xd2\xb8\xd6\x5f\xd7\x79\xc7\x0a\x0b\x75\xe3\xf0\x80\xa9\xc9\x0a\xfe\xf5\x18\x6f\xf5\x10\x75\xb2\x2b\x54\xbb\x42\xf0\xb9\x8b\x74\xd3\x71\x35\xc2\x65\xd8\xe5\x1b\x35\xb7\x3d\xd7\xe6\xe1\xc6\x4d\x4a\xc6\xc4\x50\xaf\x36\xec\x62\xba\x0a\x50\xd7\xba\x00\xd4\x9d\x99\x0f\xdb\x7c\x55\x49\xf9\x86\x92\xb8\x9e\x13\xe5\x6e\x14\x81\xdb\x49\x9e\xd5\xed\x24\x75\x9b\xa5\x4c\x58\x71\xc3\x64\xa2\x1c\xfa\x2e\x9b\x78\x2e\x28\x5c\x35\xf4\xdd\xd5\xdb\x35\xe2\x02\x7d\x3a\x5d\x7f\x7f\xe8\xaa\x11\x0b\x61\x48\x6e\x85\xaa\x9f\x5f\xdc\xa4\x50\x48\x92\xa6\x8f\x42\xc3\x09\x33\x37\x34\x3a\x7a\x5f\x52\xe0\x07\x8b\x5c\x0d\x5d\xd6\x78\xb0\x28\x07\x68\x35\xf1\x42\x1a\x29\x43\xfb\x69\x2c\x9e\x5a\xc3\xcd\xca\x64\x0b\x8d\x35\xca\xd0\x39\x9d\x86\xb8\x19\x5a\x62\xf1\x96\x32\x27\xe5\x0a\x75\x36\x75\xfb\x2f\x6d\x1c\x55\x0b\x87\x23\x4a\xc4\xc8\x17\x94\xd4\xc6\xe6\x50\xf6\xc8\xf7\xa6\xcc\x6d\x99\x33\x59\x03\xdb\xc5\x4e\x3f\xc9\x2d\x79\x20\x89\x40\x26\x77\x6b\x79\x77\xe8\x41\x12\xc0\x71\x6d\x78\x6f\x0e\x7e\x8a\x3b\x98\x23\x13\xa5\x4d\x29\x9b\xa2\xca\x23\x59\x3e\xca\x5a\xd4\x36\x77\x07\x6a\xe0\xd6\xfc\x4c\xf2\x6b\x27\x6d\x60\xbd\xc3\xc2\xb5\xc7\x66\xdd\x7a\x84\x1b\x0a\x03\x49\xfa\xf7\x33\x03\x36\x06\x5e\x23\xc8\x05\x1e\xc9\xc6\xde\x45\xe0\x91\xbe\xca\x74\xc7\x8b\xef\x30\x74\xb9\x91\x6c\x57\x7c\x00\xbf\x8b\x04\xfd\xca\x65\x9d\x85\xd7\x10\xb2\x1f\xec\x8c\xe4\x22\xfa\xcd\xde\x48\x6d\xf0\x5f\x79\xb2\x25\x9c\xd7\x9e\x12\xfe\x70\x69\x42\x71\x33\xdb\x79\x2c\xbe\x92\xe4\x22\x7a\xdc\xa1\x71\x23\x96\x5a\x1b\x08\x89\x7e\x56\x1c\x89\x80\x25\x97\xa4\xa0\x70\x0b\x92\x41\x97\x22\x47\x0f\xa4\x4f\xe3\x72\x44\x51\xe2\x9a\x5b\x63\x85\xe0\xe2\x58\xe7\x59\xde\x99\xc1\xab\x63\xb6\x5c\xdd\x4c\xbb\x2e\x1f\xb9\x2a\x2f\xcd\xf1\xf7\xea\x65\x51\x00\x8b\xdb\x33\xf5\xca\x23\x59\x85\xd7\x08\xca\xa9\x8f\x8a\xd7\x08\x6a\x56\xc7\xae\x98\x7f\xc0\xa1\x2d\x78\xe5\x00\x5e\x39\x18\x88\x06\x6d\xf1\xd3\x4c\xab\x7c\xdc\xce\x20\xc4\xd0\x28\x6a\xd4\xcb\x54\xb6\x4b\xd1\x60\x68\x14\xa7\x72\xa2\x6f\x0e\x4c\x2a\xbe\x7e\xd4\xca\xaf\xaa\x0f\x9a\xdf\x5c\x40\x1b\xc1\xf7\x3a\x91\x95\x29\xd5\x1a\xe1\x94\x52\xcb\x2b\x0a\xb9\x17\xaa\x67\xb0\xc6\xb6\x15\xf9\x95\xed\x9b\x2a\xf3\x0f\xbb\xbe\xf9\x69\xcc\x61\x79\x8b\xc5\xc1\x0d\x31\xde\xb8\x92\x78\xc4\x3c\x00\x1c\x7c\x3a\x80\x52\x9f\xb0\x20\xc4\xe5\x7b\x3c\xa0\xf2\xf9\xd4\x7c\x5e\x5d\x6b\xe2\x1f\xcb\x3f\x39\x5b\xc3\x15\xa8\x33\xb8\x6a\xa7\xb2\x94\xaf\xef\x13\x17\x03\xef\x4b\x8d\x0f\x42\x4f\x7d\x59\xea\x6c\x34\x42\x4b\xaa\x38\xd1\x0a\xc9\xca\x8c\x99\x21\x56\x58\x6c\x89\xba\x81\x44\xf1\xa8\x89\xe2\x9a\x05\x9e\x66\x1e\x57\x7a\x84\xd9\x95\xe1\x46\xbb\xfb\x2a\xff\xb8\x65\xbb\xcd\xa3\x33\xda\x33\x4c\x62\x47\xf3\xf9\x9d\xaa\xdd\x1a\x0b\xbc\x6f\xb4\x8c\xf3\x4d\x8b\x1e\xa3\x99\xe7\x17\xaa\x76\xc8\xe8\x0b\x30\xf3\xc0\xcc\x9b\xde\xcc\xdb\x51\xd7\x25\xac\x65\x35\x77\xd4\x25\xc9\x16\xd5\x0e\x9b\x39\x65\xf6\xaa\xdf\x03\x0c\xbf\x9a\x31\xc1\xec\x04\xb3\x13\xcc\xce\x29\x6b\xc8\x8a\x3a\xa7\x35\x78\x53\x6e\x5e\xa9\xa6\x86\x47\x36\x4d\x37\x13\xef\xa1\x54\x2d\xdb\x06\x39\x23\xb2\x0e\xc6\x19\x69\xc7\x11\xf9\xd4\x39\x71\x44\x06\x0c\xb7\xf1\x22\x89\x6d\x62\x53\x0c\x1e\xcb\x25\x12\x01\x63\x94\x6d\x13\xa0\xe8\x6b\xc7\x29\xb2\xf7\xbd\xec\x9d\xe6\x15\xb6\x1d\x2b\x2a\x80\x76\xfb\xee\x56\x60\xe6\xf4\x3f\xc2\xf8\x93\xf9\xbc\xeb\x53\xa8\x07\x4e\xfe\xd5\xbf\x87\x3a\xfe\x5d\x18\x4f\xd5\x3a\x48\xf6\x57\xef\x9a\x93\xb8\x83\xea\xa2\x29\xd1\xbf\x9e\xf1\xd3\xc5\x87\x59\x22\x40\xf3\x15\xc3\x99\x56\xfa\x02\xda\x18\x13\x7a\xde\x30\x5c\x7f\xc7\xf3\x3c\xa5\x6c\x3e\x6c\xfa\x50\x7c\x9b\xb8\x8e\x53\xa6\x61\x8e\x59\x49\x2d\x62\x54\xa5\x38\xa0\xba\xa8\x5c\x19\x3e\x79\xd1\x77\x3e\x02\x31\xea\x98\x17\x6f\xce\x67\xc9\xec\xca\x22\xd5\x1a\x7e\xe7\x4a\x53\x75\xf2\x31\x63\xd0\xa7\xac\xef\xcf\xf3\xda\x9b\x62\xc6\x93\xd0\x2a\x3e\x56\x44\x06\x2d\x6f\x7e\xab\xb8\xaa\x66\x3e\xac\xbd\x2a\x6b\xa8\x3a\xce\xc6\x4d\x73\xb2\xfc\x65\x47\x9d\x9d\x29\x31\x76\x30\x0b\x0d\xb3\x21\xe5\xe2\xd3\xbd\x39\x39\x81\x51\xa2\x0f\x8a\xcb\x63\x9e\x0e\x97\xc3\xdf\x8d\x2f\xe0\xf9\x21\xad\xaf\xe7\x68\xbd\xb4\xfa\x9a\xd9\x76\x59\x84\xec\xee\xd1\xd4\x7a\x9b\xf1\x10\x93\xb9\x9b\x31\x89\x96\xfe\x66\x42\xcf\x2c\x19\x16\xa3\x9a\x35\xe7\x72\x1f\x64\x59\x98\x3c\x53\x86\x52\x16\xcc\x9f\x8b\x65\x0f\xe3\x29\x72\xd3\xbc\x3c\xdd\xc6\xc2\xb4\x55\x21\xce\x63\xfe\x60\xaa\xa3\x6a\x82\x3d\xf9\xf5\xae\x58\xeb\xd1\xef\xc4\xae\x99\xab\x39\xbb\x55\xba\xfe\xba\x30\xd7\x4c\xab\xfc\x5c\xcd\x23\x6e\xe9\x8d\x11\x48\xab\x31\xfb\x42\xff\x96\x67\xa8\x4b\x57\x40\xe9\xee\x3b\xa4\xc6\xf2\x1d\xf5\xd5\x0f\x15\xb7\xe1\x94\x54\x77\x7f\x83\x60\x9f\x7f\xc1\xe1\xf8\xfb\x3f\x3d\x3c\xa9\x83\x11\x2d\x47\x2e\x2b\x5a\x97\xce\x5c\x66\x4e\x92\xe9\x83\x8f\x3d\xf7\x85\x97\x3d\x35\xf9\xc7\x37\x77\x44\xd3\xac\xf4\x45\xf9\x41\xca\x6a\x96\xc4\x0d\xcb\xdc\x08\xf6\xb7\x44\x3f\xa5\x94\x3c\x6e\xd9\xfb\x1a\xdf\xae\x2f\xc1\x1f\xfe\x9d\xd1\x19\xf0\xad\xf2\x30\x6d\xbf\xe3\xaf\x37\x9e\x3e\xd3\x0a\x87\x60\x9b\x97\xfb\x17\xc1\x03\xbf\x49\x42\xe2\x36\xa1\x70\x6c\xf5\x0f\x7c\x53\x48\x91\x1f\xf7\x80\x32\xdc\x50\x3d\xda\x05\xd3\x24\x84\xbb\x49\x2f\xd8\xd5\x5b\xe9\x5c\x8f\x63\x69\xf9\x9e\x47\x44\x35\x51\x3c\xe8\x82\x26\x9b\xab\x1e\xc2\x41\x1a\x6e\x68\x3a\xc4\x4d\xcd\x66\x8c\x50\xf0\x86\xdc\xaf\xac\x47\x99\x59\x59\x68\xb2\x2f\xea\xa1\xff\xe8\xc5\x70\xc5\xbd\x6b\x01\x98\xa6\x65\x0a\x9b\x66\xaf\xce\x04\x34\x27\x80\xa5\x67\xae\x23\xcf\x4b\xe0\xd8\xc2\xf9\x5c\x24\x5e\x47\x6a\xb3\xfb\x60\x80\xf6\x7c\x22\x30\x3d\x23\xde\x59\x71\xad\x86\x5f\x03\x18\xb5\x2d\x8e\x3d\x0d\x9f\x2c\x39\x14\xad\xc4\xdc\x18\xf4\x2b\xc1\x2e\x11\xa7\xf9\xd3\xd5\x0d\x55\xd3\x3b\xdd\x5e\xdf\x90\xa3\x03\x2b\xff\x5e\x86\xbd\x2c\xd3\x47\x58\x30\x73\xe3\x5f\x9a\xfb\x90\xa3\x4f\x24\xfa\x8e\x30\x87\xbb\xc4\x0d\x6d\xbf\x5b\x2c\xc9\xbf\xfe\xf9\x7d\x5f\x7f\x8e\x66\x6f\xf9\x5c\xe4\x7d\xf5\x14\xab\x8f\xf3\x64\x4b\x6b\x3d\xe0\x8e\x4b\x45\xd9\x76\x19\x1a\x46\x82\x61\x0f\x15\x82\x47\x33\x78\x46\xa5\x2a\x99\xf8\x84\x03\x12\xed\xf7\xfb\x57\xde\xee\x6f\xae\x9b\x47\x0e\xdf\xfb\x58\xe9\x55\xba\x1b\x7a\xc9\xff\x4f\x01\xf5\xdc\x49\x93\x67\x95\xdb\xe1\x1c\xff\x97\x8b\x09\x1e\x07\x38\xa7\x6c\x92\x7e\xd7\x58\xd5\x57\x92\x0d\xe9\x57\x90\xc9\xee\xeb\x49\x5e\x91\xa8\xe6\xc6\x41\x1f\x45\x30\xc8\xdb\xb4\xdf\x0d\x34\x2f\x97\xe8\xce\xfc\xd4\xf5\xb9\xac\x83\x40\x69\x14\x4a\xda\xe7\xef\xd2\x6f\x62\x42\xfd\x5d\xfc\xd5\xc1\x0b\xec\x63\x87\xaa\xfb\xc9\xca\x69\xe2\xfe\x9f\xe4\x65\xf3\x75\x07\x62\xe0\x4a\x92\x27\xfc\xe4\x54\xf5\x55\xf4\x07\x38\xe3\xd4\x74\xa9\x7d\x94\x8a\x1c\xf8\x48\x89\xe9\xa4\x7e\x63\xc0\xfb\x58\x47\x36\xc3\x34\x26\xb7\x95\x1c\x64\x5a\xe5\x73\xb8\x46\x53\x0d\x3d\x96\x10\xf5\x32\x15\x9b\x22\xcd\x6b\xc7\xa2\x68\xaa\xb3\x62\xd2\x96\xa8\xf2\x65\x36\xa7\x64\x83\x03\x4f\xfd\xc4\xdd\xfb\xe6\x87\xc6\x75\x95\x9c\xc2\xcc\xc5\xc2\x45\x92\x08\x8a\x3d\xfa\x97\xae\x0e\x79\xb3\x3e\x33\x0f\x71\x5f\xb3\x73\x22\x65\x94\xa1\x77\x38\x0b\x9b\xab\xe8\x7d\xfb\xbd\xf9\xcb\x8f\xd7\xec\x6f\xe8\x7a\x41\xd9\x1d\xf6\xa8\xb9\xca\x33\x5c\xb2\xeb\x85\xf9\xfd\x9f\x01\x57\x18\x91\xaf\x0e\x21\x2e\x71\xe3\xdf\xea\xb6\x46\x4d\x9a\x71\x16\xd7\xec\xe4\xe4\x84\x28\xe7\xe4\xe4\xe4\x9a\x99\x17\xd0\x73\xaf\xe9\x53\x97\x30\x45\x37\xd4\x31\x5f\x85\xce\xe9\x35\x3b\x25\x0a\x53\x4f\x47\x23\xb8\x1f\x3d\x2a\x1e\x2e\x21\xf9\x5a\x20\xd2\xbc\x12\x8f\xcd\xe0\xfa\xb1\x71\xf4\x32\x4e\xeb\xbd\x44\xfb\x40\x2a\x74\x4b\x10\xe3\x6c\xf9\x17\x11\x1c\xe9\x47\xc9\x63\x5a\x19\x57\x88\x30\x1e\x6c\x77\x48\xd1\xed\x4e\x49\xa4\x38\xda\x10\xe2\xa2\x2d\xf7\x77\x44\xc4\xed\x92\x3b\x80\x5f\xfe\xc2\xdd\x97\xc8\xe5\x44\xbe\x54\x88\x7c\xa5\x52\x85\x4d\xd2\xc7\xd0\x13\x52\x25\xd1\x01\x94\xf4\x19\x74\x1f\xd3\x21\xd1\x2e\xd7\x2c\xc7\x91\xea\xfd\x22\x66\x54\xc3\xed\xa6\xf8\x56\x3c\x9a\xe4\xbd\xf8\x63\xd4\xc1\xe7\x05\xbe\xed\x71\xf8\xce\x11\x88\x77\x42\x18\x71\xc8\xac\x49\x4e\x32\xce\x4e\xed\xa7\xef\x37\x4c\x3a\x94\x90\xd2\x6f\x2b\xe9\x79\x4f\x99\x5b\xa0\x28\xfc\xb8\x9a\x1d\x16\x65\xcb\x6d\x76\xfa\x90\xb5\x2e\x31\x2d\xb4\x60\xa8\xe5\x44\x2f\x4b\x2b\x1f\x1f\x05\x3d\x3b\xb5\x98\x6d\x85\x89\x32\x78\x2a\x53\xee\xc5\x08\x2c\x7b\x5b\x5c\x91\xae\x68\xb2\x65\x6f\x76\x4a\xf9\x53\x3d\x3a\xa7\x47\xf8\xf5\xea\x6a\xdd\xaa\xa4\xb7\x84\x2d\x3d\x6e\x74\x89\x96\x15\x9f\x08\x5c\xa8\x87\xec\xb1\xa0\xe1\x68\x7a\xa4\x85\x8f\x05\xde\x13\x95\xb5\x31\x17\x6f\x39\x63\xc4\x09\x47\x31\x81\x5c\x9b\xec\x5d\x3a\xfb\xe2\x5f\x32\x75\x64\x51\xaf\x99\xd2\x3f\x13\xbb\x34\xa3\x54\x85\x67\x95\x08\x48\xf1\xa6\xc3\xdf\xc9\xed\x25\x77\x3e\x13\xf5\x9e\xdc\x8f\x49\xe0\x25\x71\x96\x49\xdf\xcb\xf7\xe4\x7e\x38\x9d\x52\xf7\x15\xc5\x70\xea\x69\x2d\xee\xa7\xfc\x5e\xaa\xa5\xd4\xf4\xbe\x2c\x3d\x12\xda\x9d\xda\x4f\xfe\x56\x60\x97\x8c\xb9\x9a\x51\x97\x83\xa8\x92\x44\x9c\x9d\xf6\x21\x2a\x08\x32\xc1\xfc\x84\xa4\x28\x83\x90\xd1\x43\x3d\x89\x4a\x5c\x79\x7b\xda\x0a\x34\x7c\x2c\x3d\x29\xde\x93\x94\x0b\xee\x55\x50\x41\x98\x7e\xb6\xf5\x3f\x79\x37\x32\x97\xcd\x70\xf7\xa9\x33\xff\xc7\xab\x8e\xe4\x87\xa3\xf6\xa2\x3c\xc1\x9d\xf8\x9a\xec\x0c\xec\x18\x25\x55\xef\x09\x18\x00\x2b\xe8\xb2\x85\x74\x76\x64\x8f\x6d\x2e\x53\x6c\x75\x42\x5e\x64\xa1\x58\x93\xfa\xe2\xf1\xff\x02\x00\x00\xff\xff\x3a\x4f\x56\x4c\xd8\xfe\x01\x00")

func init() {
	err := CTX.Err()
	if err != nil {
		panic(err)
	}

	var f webdav.File

	var rb *bytes.Reader
	var r *gzip.Reader

	rb = bytes.NewReader(FileSwaggerJSON)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "swagger.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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

	Handler = &webdav.Handler{
		FileSystem: FS,
		LockSystem: webdav.NewMemLS(),
	}

}

// Open a file
func (hfs *HTTPFS) Open(path string) (http.File, error) {

	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// ReadFile is adapTed from ioutil
func ReadFile(path string) ([]byte, error) {
	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(make([]byte, 0, bytes.MinRead))

	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = buf.ReadFrom(f)
	return buf.Bytes(), err
}

// WriteFile is adapTed from ioutil
func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := FS.OpenFile(CTX, filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

// WalkDirs looks for files in the given dir and returns a list of files in it
// usage for all files in the b0x: WalkDirs("", false)
func WalkDirs(name string, includeDirsInList bool, files ...string) ([]string, error) {
	f, err := FS.OpenFile(CTX, name, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}

	fileInfos, err := f.Readdir(0)
	if err != nil {
		return nil, err
	}

	err = f.Close()
	if err != nil {
		return nil, err
	}

	for _, info := range fileInfos {
		filename := path.Join(name, info.Name())

		if includeDirsInList || !info.IsDir() {
			files = append(files, filename)
		}

		if info.IsDir() {
			files, err = WalkDirs(filename, includeDirsInList, files...)
			if err != nil {
				return nil, err
			}
		}
	}

	return files, nil
}
