package js

import (
	"github.com/autom8ter/objectify"
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
	Util         *objectify.Handler
}

func NewScripter(obj interface{}) *Scripter {
	o := objectify.Default()
	return &Scripter{
		Object: js.MakeWrapper(obj),
		Global: js.Global,
		DOM:    dom.GetWindow(),
		JQuery: jquery.NewJQuery(),
		ErrorHandler: func(err error) {
			if err != nil {
				o.Entry().Warnln(err.Error())
			}
		},
		Util: o,
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
