package pattern

// Паттерн Facade относится к структурным паттернам уровня объекта.
// Паттерн Facade предоставляет высокоуровневый унифицированный интерфейс в виде набора имен методов к набору
// взаимосвязанных классов или объектов некоторой подсистемы, что облегчает ее использование.
// Разбиение сложной системы на подсистемы позволяет упростить процесс разработки, а также помогает максимально снизить
// зависимости одной подсистемы от другой. Однако использовать такие подсистемы становиться довольно сложно. Один из способов
// решения этой проблемы является паттерн Facade. Наша задача, сделать простой, единый интерфейс, через который можно было бы взаимодействовать с подсистемами.

type GlobalSystem struct { // представляет собой сложную систему состояющую из нескольких подсистем
	a *System1
	b *System2
	c *System3
}

func NewGlobalSystem() *GlobalSystem {
	return &GlobalSystem{
		a: &System1{},
		b: &System2{},
		c: &System3{},
	}
}

func (g *GlobalSystem) Start() string {
	return g.a.Start() + g.b.Start() + g.c.Start() + "System totally started!"
}

type System1 struct {
}

func (r *System1) Start() string {
	return "Start System-1\n"
}

type System2 struct {
}

func (r *System2) Start() string {
	return "Start System-2\n"
}

type System3 struct {
}

func (n *System3) Start() string {
	return "Start System-3\n"
}
