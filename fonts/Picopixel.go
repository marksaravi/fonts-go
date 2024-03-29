package fonts

var Picopixel BitmapFont = BitmapFont{scale: 1, Bitmap: []byte{
	0xE8, 0xB4, 0x57, 0xD5, 0xF5, 0x00, 0x4E, 0x3E, 0x80, 0xA5, 0x4A, 0x4A,
	0x5A, 0x50, 0xC0, 0x6A, 0x40, 0x95, 0x80, 0xAA, 0x80, 0x5D, 0x00, 0x60,
	0xE0, 0x80, 0x25, 0x48, 0x56, 0xD4, 0x75, 0x40, 0xC5, 0x4E, 0xC5, 0x1C,
	0x97, 0x92, 0xF3, 0x1C, 0x53, 0x54, 0xE5, 0x48, 0x55, 0x54, 0x55, 0x94,
	0xA0, 0x46, 0x64, 0xE3, 0x80, 0x98, 0xC5, 0x04, 0x56, 0xC6, 0x57, 0xDA,
	0xD7, 0x5C, 0x72, 0x46, 0xD6, 0xDC, 0xF3, 0xCE, 0xF3, 0x48, 0x72, 0xD4,
	0xB7, 0xDA, 0xF8, 0x24, 0xD4, 0xBB, 0x5A, 0x92, 0x4E, 0x8E, 0xEB, 0x58,
	0x80, 0x9D, 0xB9, 0x90, 0x56, 0xD4, 0xD7, 0x48, 0x56, 0xD4, 0x40, 0xD7,
	0x5A, 0x71, 0x1C, 0xE9, 0x24, 0xB6, 0xD4, 0xB6, 0xA4, 0x8C, 0x6B, 0x55,
	0x00, 0xB5, 0x5A, 0xB5, 0x24, 0xE5, 0x4E, 0xEA, 0xC0, 0x91, 0x12, 0xD5,
	0xC0, 0x54, 0xF0, 0x90, 0xC7, 0xF0, 0x93, 0x5E, 0x71, 0x80, 0x25, 0xDE,
	0x5E, 0x30, 0x6E, 0x80, 0x77, 0x9C, 0x93, 0x5A, 0xB8, 0x45, 0x60, 0x92,
	0xEA, 0xAA, 0x40, 0xD5, 0x6A, 0xD6, 0x80, 0x55, 0x00, 0xD7, 0x40, 0x75,
	0x90, 0xE8, 0x71, 0xE0, 0xBA, 0x40, 0xB5, 0x80, 0xB5, 0x00, 0x8D, 0x54,
	0xAA, 0x80, 0xAC, 0xE0, 0xE5, 0x70, 0x6A, 0x26, 0xFC, 0xC8, 0xAC, 0x5A,
},
	Glyphs: []Glyph{{BitmapOffset: 0, Width: 0, Height: 0, XAdvance: 2, XOffset: 0, YOffset: 1}, // 0x20 ' '
		{BitmapOffset: 0, Width: 1, Height: 5, XAdvance: 2, XOffset: 0, YOffset: -4},   // 0x21 '!'
		{BitmapOffset: 1, Width: 3, Height: 2, XAdvance: 4, XOffset: 0, YOffset: -4},   // 0x22 '"'
		{BitmapOffset: 2, Width: 5, Height: 5, XAdvance: 6, XOffset: 0, YOffset: -4},   // 0x23 '#'
		{BitmapOffset: 6, Width: 3, Height: 6, XAdvance: 4, XOffset: 0, YOffset: -4},   // 0x24 '$'
		{BitmapOffset: 9, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4},   // 0x25 '%'
		{BitmapOffset: 11, Width: 4, Height: 5, XAdvance: 5, XOffset: 0, YOffset: -4},  // 0x26 '&'
		{BitmapOffset: 14, Width: 1, Height: 2, XAdvance: 2, XOffset: 0, YOffset: -4},  // 0x27 '''
		{BitmapOffset: 15, Width: 2, Height: 5, XAdvance: 3, XOffset: 0, YOffset: -4},  // 0x28 '('
		{BitmapOffset: 17, Width: 2, Height: 5, XAdvance: 3, XOffset: 0, YOffset: -4},  // 0x29 ')'
		{BitmapOffset: 19, Width: 3, Height: 3, XAdvance: 4, XOffset: 0, YOffset: -3},  // 0x2A '*'
		{BitmapOffset: 21, Width: 3, Height: 3, XAdvance: 4, XOffset: 0, YOffset: -3},  // 0x2B '+'
		{BitmapOffset: 23, Width: 2, Height: 2, XAdvance: 3, XOffset: 0, YOffset: 0},   // 0x2C ','
		{BitmapOffset: 24, Width: 3, Height: 1, XAdvance: 4, XOffset: 0, YOffset: -2},  // 0x2D '-'
		{BitmapOffset: 25, Width: 1, Height: 1, XAdvance: 2, XOffset: 0, YOffset: 0},   // 0x2E '.'
		{BitmapOffset: 26, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4},  // 0x2F '/'
		{BitmapOffset: 28, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4},  // 0x30 '0'
		{BitmapOffset: 30, Width: 2, Height: 5, XAdvance: 3, XOffset: 0, YOffset: -4},  // 0x31 '1'
		{BitmapOffset: 32, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4},  // 0x32 '2'
		{BitmapOffset: 34, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4},  // 0x33 '3'
		{BitmapOffset: 36, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4},  // 0x34 '4'
		{BitmapOffset: 38, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4},  // 0x35 '5'
		{BitmapOffset: 40, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4},  // 0x36 '6'
		{BitmapOffset: 42, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4},  // 0x37 '7'
		{BitmapOffset: 44, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4},  // 0x38 '8'
		{BitmapOffset: 46, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4},  // 0x39 '9'
		{BitmapOffset: 48, Width: 1, Height: 3, XAdvance: 2, XOffset: 0, YOffset: -3},  // 0x3A ':'
		{BitmapOffset: 49, Width: 2, Height: 4, XAdvance: 3, XOffset: 0, YOffset: -3},  // 0x3B ';'
		{BitmapOffset: 50, Width: 2, Height: 3, XAdvance: 3, XOffset: 0, YOffset: -3},  // 0x3C '<'
		{BitmapOffset: 51, Width: 3, Height: 3, XAdvance: 4, XOffset: 0, YOffset: -3},  // 0x3D '='
		{BitmapOffset: 53, Width: 2, Height: 3, XAdvance: 3, XOffset: 0, YOffset: -3},  // 0x3E '>'
		{BitmapOffset: 54, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4},  // 0x3F '?'
		{BitmapOffset: 56, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4},  // 0x40 '@'
		{BitmapOffset: 58, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4},  // 0x41 'A'
		{BitmapOffset: 60, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4},  // 0x42 'B'
		{BitmapOffset: 62, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4},  // 0x43 'C'
		{BitmapOffset: 64, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4},  // 0x44 'D'
		{BitmapOffset: 66, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4},  // 0x45 'E'
		{BitmapOffset: 68, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4},  // 0x46 'F'
		{BitmapOffset: 70, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4},  // 0x47 'G'
		{BitmapOffset: 72, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4},  // 0x48 'H'
		{BitmapOffset: 74, Width: 1, Height: 5, XAdvance: 2, XOffset: 0, YOffset: -4},  // 0x49 'I'
		{BitmapOffset: 75, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4},  // 0x4A 'J'
		{BitmapOffset: 77, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4},  // 0x4B 'K'
		{BitmapOffset: 79, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4},  // 0x4C 'L'
		{BitmapOffset: 81, Width: 5, Height: 5, XAdvance: 6, XOffset: 0, YOffset: -4},  // 0x4D 'M'
		{BitmapOffset: 85, Width: 4, Height: 5, XAdvance: 5, XOffset: 0, YOffset: -4},  // 0x4E 'N'
		{BitmapOffset: 88, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4},  // 0x4F 'O'
		{BitmapOffset: 90, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4},  // 0x50 'P'
		{BitmapOffset: 92, Width: 3, Height: 6, XAdvance: 4, XOffset: 0, YOffset: -4},  // 0x51 'Q'
		{BitmapOffset: 95, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4},  // 0x52 'R'
		{BitmapOffset: 97, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4},  // 0x53 'S'
		{BitmapOffset: 99, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4},  // 0x54 'T'
		{BitmapOffset: 101, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4}, // 0x55 'U'
		{BitmapOffset: 103, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4}, // 0x56 'V'
		{BitmapOffset: 105, Width: 5, Height: 5, XAdvance: 6, XOffset: 0, YOffset: -4}, // 0x57 'W'
		{BitmapOffset: 109, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4}, // 0x58 'X'
		{BitmapOffset: 111, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4}, // 0x59 'Y'
		{BitmapOffset: 113, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4}, // 0x5A 'Z'
		{BitmapOffset: 115, Width: 2, Height: 5, XAdvance: 3, XOffset: 0, YOffset: -4}, // 0x5B '['
		{BitmapOffset: 117, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4}, // 0x5C '\'
		{BitmapOffset: 119, Width: 2, Height: 5, XAdvance: 3, XOffset: 0, YOffset: -4}, // 0x5D ']'
		{BitmapOffset: 121, Width: 3, Height: 2, XAdvance: 4, XOffset: 0, YOffset: -4}, // 0x5E '^'
		{BitmapOffset: 122, Width: 4, Height: 1, XAdvance: 4, XOffset: 0, YOffset: 1},  // 0x5F '_'
		{BitmapOffset: 123, Width: 2, Height: 2, XAdvance: 3, XOffset: 0, YOffset: -4}, // 0x60 '`'
		{BitmapOffset: 124, Width: 3, Height: 4, XAdvance: 4, XOffset: 0, YOffset: -3}, // 0x61 'a'
		{BitmapOffset: 126, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4}, // 0x62 'b'
		{BitmapOffset: 128, Width: 3, Height: 3, XAdvance: 4, XOffset: 0, YOffset: -2}, // 0x63 'c'
		{BitmapOffset: 130, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4}, // 0x64 'd'
		{BitmapOffset: 132, Width: 3, Height: 4, XAdvance: 4, XOffset: 0, YOffset: -3}, // 0x65 'e'
		{BitmapOffset: 134, Width: 2, Height: 5, XAdvance: 3, XOffset: 0, YOffset: -4}, // 0x66 'f'
		{BitmapOffset: 136, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -3}, // 0x67 'g'
		{BitmapOffset: 138, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4}, // 0x68 'h'
		{BitmapOffset: 140, Width: 1, Height: 5, XAdvance: 2, XOffset: 0, YOffset: -4}, // 0x69 'i'
		{BitmapOffset: 141, Width: 2, Height: 6, XAdvance: 3, XOffset: 0, YOffset: -4}, // 0x6A 'j'
		{BitmapOffset: 143, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4}, // 0x6B 'k'
		{BitmapOffset: 145, Width: 2, Height: 5, XAdvance: 3, XOffset: 0, YOffset: -4}, // 0x6C 'l'
		{BitmapOffset: 147, Width: 5, Height: 3, XAdvance: 6, XOffset: 0, YOffset: -2}, // 0x6D 'm'
		{BitmapOffset: 149, Width: 3, Height: 3, XAdvance: 4, XOffset: 0, YOffset: -2}, // 0x6E 'n'
		{BitmapOffset: 151, Width: 3, Height: 3, XAdvance: 4, XOffset: 0, YOffset: -2}, // 0x6F 'o'
		{BitmapOffset: 153, Width: 3, Height: 4, XAdvance: 4, XOffset: 0, YOffset: -2}, // 0x70 'p'
		{BitmapOffset: 155, Width: 3, Height: 4, XAdvance: 4, XOffset: 0, YOffset: -2}, // 0x71 'q'
		{BitmapOffset: 157, Width: 2, Height: 3, XAdvance: 3, XOffset: 0, YOffset: -2}, // 0x72 'r'
		{BitmapOffset: 158, Width: 3, Height: 4, XAdvance: 4, XOffset: 0, YOffset: -3}, // 0x73 's'
		{BitmapOffset: 160, Width: 2, Height: 5, XAdvance: 3, XOffset: 0, YOffset: -4}, // 0x74 't'
		{BitmapOffset: 162, Width: 3, Height: 3, XAdvance: 4, XOffset: 0, YOffset: -2}, // 0x75 'u'
		{BitmapOffset: 164, Width: 3, Height: 3, XAdvance: 4, XOffset: 0, YOffset: -2}, // 0x76 'v'
		{BitmapOffset: 166, Width: 5, Height: 3, XAdvance: 6, XOffset: 0, YOffset: -2}, // 0x77 'w'
		{BitmapOffset: 168, Width: 3, Height: 3, XAdvance: 4, XOffset: 0, YOffset: -2}, // 0x78 'x'
		{BitmapOffset: 170, Width: 3, Height: 4, XAdvance: 4, XOffset: 0, YOffset: -2}, // 0x79 'y'
		{BitmapOffset: 172, Width: 3, Height: 4, XAdvance: 4, XOffset: 0, YOffset: -3}, // 0x7A 'z'
		{BitmapOffset: 174, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4}, // 0x7B '{'
		{BitmapOffset: 176, Width: 1, Height: 6, XAdvance: 2, XOffset: 0, YOffset: -4}, // 0x7C '|'
		{BitmapOffset: 177, Width: 3, Height: 5, XAdvance: 4, XOffset: 0, YOffset: -4}, // 0x7D '}'
		{BitmapOffset: 179, Width: 4, Height: 2, XAdvance: 5, XOffset: 0, YOffset: -3}, // 0x7E '~'
	},
}
