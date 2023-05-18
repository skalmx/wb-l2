package pattern

type IComputer interface {
	SetPrice(price int)
	GetPrice() int
}

type Computer struct {
	name  string
	price int
}

type HyperXComputer struct {
	Computer
}

func NewHyperXComputer() *HyperXComputer {
	return &HyperXComputer{
		Computer: Computer{
			name:  "HyperPC",
			price: 10000000,
		},
	}
}

func (c *HyperXComputer) SetPrice(price int) {
	c.price = price
}

func (c *HyperXComputer) GetPrice() int {
	return c.price
}

type TrashComputer struct {
	Computer
}

func NewTrashComputer() *TrashComputer {
	return &TrashComputer{
		Computer: Computer{
			name:  "VERY OLD AND BAD PC",
			price: 100,
		},
	}
}

func (c *TrashComputer) SetPrice(price int) {
	c.price = price
}

func (c *TrashComputer) GetPrice() int {
	return c.price
}

func GetComputer(model string) IComputer {
	if model == "HyperPC" {
		return NewHyperXComputer()
	}
	if model == "Cheap" {
		return NewTrashComputer()
	}
	return nil
}

// func main() {
// 	h := GetComputer("HyperPC")
// 	t := GetComputer("Cheap")
// 	fmt.Println(h)
// 	fmt.Println(t)
// }
