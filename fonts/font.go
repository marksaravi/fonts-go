package fonts

import "fmt"

type FontInfo struct {
	LineHeight       int
	YOffsetAboveLine int
	YOffsetBelowLine int
}

type Glyph struct {
	BitmapOffset int ///< Pointer into GFXfont->bitmap
	Width        int ///< Bitmap dimensions in pixels
	Height       int ///< Bitmap dimensions in pixels
	XAdvance     int ///< Distance to advance cursor (x axis)
	XOffset      int ///< X dist from cursor pos to UL corner
	YOffset      int ///< Y dist from cursor pos to UL corner
}
type BitmapFont struct {
	Bitmap []byte
	Glyphs []Glyph
	Scale  int
}
type MatrixFont struct {
	Data          []byte
	Width         int
	Height        int
	LetterSpacing int
	Scale         int
}

//16x21 A
// font := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0E, 0x00, 0x1F, 0x00, 0x1F, 0x00,
// 	0x1F, 0x00, 0x3B, 0x80, 0x3B, 0x80, 0x71, 0x80, 0x7F, 0xC0, 0x71, 0xC0, 0xE0, 0xE0, 0xE0, 0xE0,
// 	0xE0, 0xE0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

func (f *BitmapFont) GetInfo() FontInfo {
	YOffsetAboveLine := 100000
	YOffsetBelowLine := -100000
	for _, g := range f.Glyphs {
		if YOffsetAboveLine > g.YOffset {
			YOffsetAboveLine = g.YOffset
		}
		y := g.Height + g.YOffset
		if YOffsetBelowLine < y {
			YOffsetBelowLine = y
		}
	}

	LineHeight := YOffsetBelowLine - YOffsetAboveLine
	YOffsetAboveLine = -YOffsetAboveLine
	return FontInfo{
		LineHeight:       LineHeight,
		YOffsetAboveLine: YOffsetAboveLine,
		YOffsetBelowLine: YOffsetBelowLine,
	}
}

func (f *BitmapFont) SetScale(scale int) error {
	if scale < 1 || scale > 5 {
		return fmt.Errorf("scale must be >=1 and <=5")
	}
	f.Scale = scale
	return nil
}
