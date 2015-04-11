// Outputs widget
package widgets

type Dial struct {
	Style string
	Gid   int
	Wid   int
	Label string
	Value float32
	Min   float32
	Max   float32
}

type Bar Dial
type Number Dial

func (d *Dial) Init() {
	d.Style = "dial"
}

func (b *Bar) Init() {
	b.Style = "bar"
}

func (n *Number) Init() {
	n.Style = "number"
}

type Bool struct {
	Style string
	Gid   int
	Wid   int
	Label string
	Value bool
}

func (b *Bool) Init() {
	b.Style = "bool"
}

type Bools struct {
	Style  string
	Gid    int
	Wid    int
	Label  string
	Values []bool
}

func (b *Bools) Init() {
	b.Style = "bools"
}

type Dials struct {
	Style  string
	Gid    int
	Wid    int
	Label  string
	Values []float32
	Mins   []float32
	Maxes  []float32
}

type Bars Dials
type Numbers Dials

func (d *Dials) Init() {
	d.Style = "dials"
}

func (b *Bars) Init() {
	b.Style = "bars"
}

func (n *Numbers) Init() {
	n.Style = "numbers"
}

type Map struct {
	Style   string
	Gid     int
	Wid     int
	Label   string
	Values  [][2]float32
	Types   []string
	MapType string
}

func (m *Map) Init() {
	m.Style = "map"
}
