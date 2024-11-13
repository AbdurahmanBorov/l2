package pattern

import "fmt"

type Handler interface {
	SetNext(handler Handler)
	Handle(request string)
}

type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SetNext(handler Handler) {
	h.next = handler
}

func (h *BaseHandler) Handle(request string) {
	if h.next != nil {
		h.next.Handle(request)
	}
}

type AuthHandler struct {
	BaseHandler
}

func (h *AuthHandler) Handle(request string) {
	if request == "Auth" {
		fmt.Println("Аутентификация выполнена")
	} else {
		fmt.Println("Переход к следующему обработчику")
		h.BaseHandler.Handle(request)
	}
}

func ChainOfResp() {
	auth := &AuthHandler{}
	auth.Handle("Auth")
}
