package evergreen

import (
	"github.com/nu7hatch/gouuid"
	"log"
)

func newUuid() (result string, err error) {
	u4, err := uuid.NewV4()

	if err != nil {
		log.Printf("UUID generation error: %v", err)
	}

	result = u4.String()
	return
}

// This empty gif definition below is borrowed from nginx's code.
// http://nginx.org/
var emptyGif = []byte{
	// header
	'G', 'I', 'F', '8', '9', 'a',

	// logical screen descriptor
	// logical screen width
	0x01, 0x00,
	// logical screen height
	0x01, 0x00,
	// global 1-bit color table
	0x80,
	// background color #1
	0x01,
	// no aspect ratio
	0x00,

	// global color table
	// #0: black
	0x00, 0x00, 0x00,
	// #1: white
	0xff, 0xff, 0xff,

	// graphic control extension
	// extension introducer
	0x21,
	// graphic control label
	0xf9,
	// block size
	0x04,
	// transparent color is given,
	0x01,
	// no disposal specified,
	// user input is not expected
	// delay time
	0x00, 0x00,
	// transparent color #1
	0x01,
	// block terminator
	0x00,

	// image descriptor
	// image separator
	0x2c,
	// image left position
	0x00, 0x00,
	// image top position
	0x00, 0x00,
	// image width
	0x01, 0x00,
	// image height
	0x01, 0x00,
	// no local color table, no interlaced
	0x00,

	// table based image data
	// LZW minimum code size,
	0x02,
	// must be at least 2-bit
	// block size
	0x02,
	// compressed bytes 01_001_100, 0000000_1
	0x4c, 0x01,
	// 100: clear code
	// 001: 1
	// 101: end of information code
	// block terminator
	0x00,
	// trailer
	0x3B,
}

func newEmptyGif() (result []byte) {
	return emptyGif
}
