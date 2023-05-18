package pattern

import "fmt"

type Handler interface {
	DoSomething(str string)
	Next(h Handler)
}

type ConcreteHandler1 struct {
	h Handler
}

func (h *ConcreteHandler1) DoSomething(str string) {
	fmt.Println("Enter to handler1")
	if (str == "1") {
		fmt.Println("Hello from DoSomething from handler-1")
	} else if h.h != nil {
		h.Next(h.h)
		h.h.DoSomething(str)
	}
	return
}

func (h *ConcreteHandler1) Next(handler Handler) {
	h.h = handler
}

type ConcreteHandler2 struct {
	h Handler
}

func (h *ConcreteHandler2) DoSomething(str string) {
	fmt.Println("Enter to handler2")
	if (str == "2") {
		fmt.Println("Hello from DoSomething from handler-2")
	} else if h.h != nil {
		h.Next(h.h)
		h.h.DoSomething(str)
	}
	return
}

func (h *ConcreteHandler2) Next(handler Handler) {
	h.h = handler
}

type ConcreteHandler3 struct {
	h Handler
}

func (h *ConcreteHandler3) DoSomething(str string) {
	fmt.Println("Enter to handler3")
	if (str == "3") {
		fmt.Println("Hello from DoSomething from handler-3")
	} else if h.h != nil {
		h.Next(h.h)
		h.h.DoSomething(str)
	}
	return
}

func (h *ConcreteHandler3) Next(handler Handler) {
	h.h = handler
}

// func main() {
// 	handlers := &ConcreteHandler1{ // Enter to handler1 Enter to handler2 Hello from DoSomething from handler-2
// 		h: &ConcreteHandler2{
// 			h: &ConcreteHandler3{},
// 		},
// 	}
// 	handlers.DoSomething("2")
// }


