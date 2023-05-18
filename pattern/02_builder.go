package pattern



// Паттерн Builder относится к порождающим паттернам уровня объекта.
// Паттерн Builder определяет процесс поэтапного построения сложного продукта. После того как будет построена последняя его часть, продукт можно использовать.
// Также с помощью данного паттерна можно реализовать создание какого то базового вида (в моем случае создать, какую то пиццу, где все ингридиенты будут по умолчанию)

type Pizza struct {
	size int
	toppings []string
	price int
}

type Builder interface {
	SetPizzaSize(size int)
	SetPizzaTopings(toppings []string)
	SetPizzaPrice(price int)
	GetPizza() Pizza
}

type PizzaBuilder struct {
	pizza *Pizza
}

func NewPizzaBuilder() *PizzaBuilder {
	return &PizzaBuilder{}
}

func (p *PizzaBuilder) SetPizzaSize(size int) {
	p.pizza.size = size
}

func (p *PizzaBuilder) SetPizzaTopings(toppings []string) {
	p.pizza.toppings = append(p.pizza.toppings, toppings...)
}

func (p *PizzaBuilder) SetPizzaPrice(price int) {
	p.pizza.price = price
}

func (p *PizzaBuilder) GetPizza() Pizza {
	return Pizza{
		size: p.pizza.size,
		toppings: p.pizza.toppings,
		price: p.pizza.price,
	}
}

// func main() {
// 	director := NewPizzaBuilder()
// 	director.SetPizzaSize(30)
// 	toppings := []string{"cheese", "smthing"}
// 	director.SetPizzaTopings(toppings)
// 	director.SetPizzaPrice(36)
// 	fmt.Println(director.GetPizza())
// }
