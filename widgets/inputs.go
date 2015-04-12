// Inputs widget
package widgets

type Slider struct {
	WidgetBase
	Value float32
	Min   float32
	Max   float32
}

func (s *Slider) Init() {
	s.Style = "slider"
}

type Toggle struct {
	WidgetBase
	Value bool
}

func (t *Toggle) Init() {
	t.Style = "toggle"
}

type Button struct {
	WidgetBase
	Value bool
}

func (t *Button) Init() {
	t.Style = "button"
}

type Selector struct {
	WidgetBase
	Value   string
	Options []string
}

func (s *Selector) Init() {
	s.Style = "selector"
}
