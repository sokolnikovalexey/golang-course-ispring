package font

import(
	"fmt"
)

type Font struct {
	family string
	size   int
}

func (f *Font) String() string {
	return fmt.Sprintf(`{font-family: "%v"; font-size: %vpt;}`, f.family, f.size)
}

func (f* Font) SetFamily(s string) {
	if len(s) > 0	{
		f.family = s;
	}
}
func (f* Font) SetSize(i int) {
	if i > MIN_FONT_SIZE && i < MAX_FONT_SIZE {
		f.size = i;
	}
}
func (f *Font) Family() string {
	return f.family
}
func (f* Font) Size() int {
	return f.size
}

const (
	MIN_FONT_SIZE  = 4
	MAX_FONT_SIZE = 145
	DEFAULT_FONT_SIZE = 8
	DEFAULT_FONT_FAMILY = "Arial"

)


func New(fontFamily string, size int) *Font {

	if size <= MIN_FONT_SIZE || size >= MAX_FONT_SIZE {
		size = DEFAULT_FONT_SIZE
	}
	if len(fontFamily) < 1	{
		fontFamily = DEFAULT_FONT_FAMILY
	}

	f := &Font{}
	f.SetFamily(fontFamily)
	f.SetSize(size)
	return f
}