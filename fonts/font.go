package fonts

type Glyph struct {
	BitmapOffset int ///< Pointer into GFXfont->bitmap
	Width        int ///< Bitmap dimensions in pixels
	Height       int ///< Bitmap dimensions in pixels
	XAdvance     int ///< Distance to advance cursor (x axis)
	XOffset      int ///< X dist from cursor pos to UL corner
	YOffset      int ///< Y dist from cursor pos to UL corner
}
type BitmapFont struct {
	Bitmap    []byte
	Glyphs    []Glyph
}
type MatrixFont struct {
	Data          []byte
	Width         int
	Height        int
	LetterSpacing int
}

//16x21 A
// font := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0E, 0x00, 0x1F, 0x00, 0x1F, 0x00,
// 	0x1F, 0x00, 0x3B, 0x80, 0x3B, 0x80, 0x71, 0x80, 0x7F, 0xC0, 0x71, 0xC0, 0xE0, 0xE0, 0xE0, 0xE0,
// 	0xE0, 0xE0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
