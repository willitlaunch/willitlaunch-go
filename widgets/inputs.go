// Inputs widget
package widgets

type Slider struct {
	Style string
	Gid   int
	Wid   int
	Label string
	Value float32
	Min   float32
	Max   float32
}

func (s *Slider) Init() {
	s.Style = "slider"
}

type Toggle struct {
	Style string
	Gid   int
	Wid   int
	Label string
	Value bool
}

func (t *Toggle) Init() {
	t.Style = "toggle"
}

type Button struct {
	Style string
	Gid   int
	Wid   int
	Label string
	Value bool
}

func (t *Button) Init() {
	t.Style = "button"
}

type Selector struct {
	Style   string
	Gid     int
	Wid     int
	Label   string
	Value   string
	Options []string
}

func (s *Selector) Init() {
	s.Style = "selector"
}
