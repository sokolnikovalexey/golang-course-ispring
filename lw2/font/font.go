package font

import(
	"fmt"
)

type Font struct {
	fontFamily string
	size int
}

func (m *Font) String() string {
	return fmt.Sprintf(`{font-family: "%v"; font-size: %vpt;}`, m.fontFamily, m.size)
}

func (m* Font) SetFamily(s string) {
	if len(s) > 0	{
		m.fontFamily = s;
	}
}
func (m* Font) SetSize(i int) {
	if i > 4 && i < 145 {
		m.size = i;
	}
}
func (m *Font) Family() string {
	return m.fontFamily
}
func (m* Font) Size() int {
	return m.size
}

func New(fontFamily string, size int) *Font {
	f := &Font{}
	f.SetFamily(fontFamily)
	f.SetSize(size)
	return f
}