package controllers

type ControlController struct {
}

func (c *ControlController) Init() {
}

func (c *ControlController) Tick() {
}

func (c *ControlController) Update(event Event) {
}

func (c *ControlController) GetInitJSON() []byte {
	var json []byte
	return json
}

func (c *ControlController) GetTickJSON() []byte {
	var json []byte
	return json
}
