package app

import "github.com/rizalarfiyan/be-petang/app/handler"

type Router interface {
	BaseRoute(handler handler.BaseHandler)
}
