package main

import (
	"fmt"
)

func main() {
	btn := MakeButton()

	handlerOne := make(chan string)
	handlerTwo := make(chan string)

	btn.AddEventListener("click", handlerOne)
	btn.AddEventListener("click", handlerTwo)


	go func () {
		for {
			msg := <- handlerOne
			fmt.Println("handler One : " +  msg)
		}
	}()

	go func () {
		for {
			msg := <- handlerTwo
			fmt.Println("handler Two : " +  msg)
		}
	}()



	btn.TriggerEvent("click", " button Clicked!")
	btn.RemoveEventListener("click", handlerTwo)
	btn.TriggerEvent("click", " button Clicked again1!")
	fmt.Scanln()
}



type Button struct {
	eventListeners map[string][] chan string
}

func MakeButton() *Button{
	result := new(Button)

	result.eventListeners = make(map[string][] chan string)

	return result
}

func (this *Button) AddEventListener(event string, responseChannel chan string) {
	if _, present := this.eventListeners[event]; present {
		this.eventListeners[event] = 
			append(this.eventListeners[event], responseChannel)
	} else {
		this.eventListeners[event] = []chan string{responseChannel}
	}
}

func (this *Button) RemoveEventListener(event string, listenerChannel chan string) {
	if _, present := this.eventListeners[event]; present {
		for idx, _ := range this.eventListeners[event] {
			if this.eventListeners[event][idx] == listenerChannel {
				this.eventListeners[event] = append(this.eventListeners[event][:idx],
				this.eventListeners[event][idx + 1:]...)
				break;
			}
		}
	}
}

func (this *Button) TriggerEvent (event string, response string) {
	if _, present :=  this.eventListeners[event]; present {
		for _, handler := range this.eventListeners[event] {
			go func (handler chan string) {
				handler <- response
			}(handler)
		}
	}
}


