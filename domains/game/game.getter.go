package game

func (g Game) State() State {
	return g.state
}

func (g Game) ID() string {
	return g.id.String()
}

func (g Game) Manager() string {
	return g.manager.Name()
}
