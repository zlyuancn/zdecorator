/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2019/5/9
   Description :
-------------------------------------------------
*/

package zdecorator

import "reflect"

type Function interface{}
type DecoratorFunc func(d *SrcFunction) Function

type Decorator struct {
    root *SrcFunction
}

func New(fn Function) *Decorator {
    return &Decorator{
        root: &SrcFunction{
            rawfn:      fn,
            reflection: reflect.ValueOf(fn),
        },
    }
}

func (m *Decorator) AddDecorator(fns ...DecoratorFunc) *Decorator {
    for _, fn := range fns {
        outfn := fn(m.root)
        m.root = &SrcFunction{
            rawfn:      outfn,
            reflection: reflect.ValueOf(outfn),
        }
    }
    return m
}

func (m *Decorator) Call(arg ...interface{}) *Result {
    return m.root.Call(arg...)
}

func (m *Decorator) RawFn() interface{} {
    return m.root.rawfn
}
