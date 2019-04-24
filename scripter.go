package js

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/jquery"
	"honnef.co/go/js/dom"
)

type Scripter struct {
	Object       *js.Object
	Global       *js.Object
	DOM          dom.Window
	JQuery       jquery.JQuery
	ErrorHandler func(err error)
}

func NewScripter(obj interface{}) *Scripter {
	return &Scripter{
		Object: js.MakeWrapper(obj),
		Global: js.Global,
		DOM:    dom.GetWindow(),
		JQuery: jquery.NewJQuery(),
	}
}

type Script func(s *Scripter) error

func (s *Scripter) RunScripts(scripts ...Script) {
	for _, script := range scripts {
		if err := script(s); err != nil {
			s.ErrorHandler(err)
		}
	}
}
