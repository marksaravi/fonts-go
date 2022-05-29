Author: Mark Saravi  
# fonts-go
This repo is a translation of fonts data from **C/C++** [Adafruit GFX Library](https://github.com/adafruit/Adafruit-GFX-Library) to GO.

sample code:  
(This is a psudo code to show how to interpret the fonts data)  
```go
func drawBitmapChar(char byte) {  
    font:=fonts.FreeMono18pt7b  
	glyph := font.Glyphs[char-0x20]  
	for h := 0; h < glyph.Height; h++ {  
		for w := 0; w < glyph.Width; w++ {  
			bitIndex := h*glyph.Width + w  
			shift := byte(bitIndex) % 8  
			d := font.Bitmap[glyph.BitmapOffset+bitIndex/8]  
			mask := byte(0b10000000) >> shift  
			bit := d & mask  
			if bit != 0 {  
				x := float64(cursorX + w + glyph.XOffset)  
				y := float64(cursorY + lineHeight + h + glyph.YOffset)  
				Pixel(x, y, fontColor)  
			}  
		}  
	}  
	xforward := glyph.XAdvance  
	cursorX += xforward  
	if cursorX+xforward >= ScreenWidth() {  
		nextLine()  
	}  
}  
```