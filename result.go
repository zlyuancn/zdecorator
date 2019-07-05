/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2019/5/17
   Description :
-------------------------------------------------
*/

package zdecorator

import "reflect"

type Result struct {
    result []interface{}
}

func makeResult(values []reflect.Value) *Result {
    result := make([]interface{}, len(values))
    for i, value := range values {
        result[i] = value.Interface()
    }
    return &Result{result: result}
}

func (m *Result) Value() interface{} {
    return m.result[0]
}

func (m *Result) ValueNum1() interface{} {
    return m.result[0]
}

func (m *Result) ValueNum2() (interface{}, interface{}) {
    return m.result[0], m.result[1]
}

func (m *Result) ValueNum3() (interface{}, interface{}, interface{}) {
    return m.result[0], m.result[1], m.result[2]
}

func (m *Result) ValueOf(index int) interface{} {
    return m.result[index]
}

func (m *Result) ValueAll() []interface{} {
    return m.result
}

func (m *Result) ValueSlice(start, end int) []interface{} {
    return m.result[start:end]
}

func (m *Result) ValueAndError() (interface{}, error) {
    return m.result[0], m.result[1].(error)
}

func (m *Result) ValueAndBool() (interface{}, bool) {
    return m.result[0], m.result[1].(bool)
}

func (m *Result) ValueLen() int {
    return len(m.result)
}
