package main

import "fmt"


type TVState interface {
	state()
}


type on struct{

}

func (o *on) state() { 
	fmt.Println("TV is ON!")
}

type off struct{

}

func (o *off) state() { 
	fmt.Println("TV is OFF!")
}


type stateContext struct {
	currentTvState TVState
}

func getContext() *stateContext {
	return &stateContext{
		currentTvState: &off{},
	}
}

func (sc *stateContext) setState(state TVState) {
	sc.currentTvState = state
}

func (sc *stateContext) getState() {
	sc.currentTvState.state()
}

// func main() {
// 	tvContext := getContext() 
// 	tvContext.getState()     
// 	tvContext.setState(&on{})
// 	tvContext.getState()      

// }