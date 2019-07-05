/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2019/5/24
   Description :
-------------------------------------------------
*/

package zdecorator

import "reflect"

type SrcFunction struct {
    rawfn         Function
    reflection reflect.Value
}

func (m *SrcFunction) Call(args ...interface{}) *Result {
    argsValue := make([]reflect.Value, len(args))
    for i, arg := range args {
        argsValue[i] = reflect.ValueOf(arg)
    }

    result := m.reflection.Call(argsValue)
    return makeResult(result)
}

func (m *SrcFunction) RawFn() interface{} {
    return m.rawfn
}
