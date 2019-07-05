/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2019/5/17
   Description :
-------------------------------------------------
*/

package zdecorator

type FuncResult struct {
    result []interface{}
}

func (m *FuncResult) Value() interface{} {
    return m.result[0]
}

func (m *FuncResult) ValueNum1() interface{} {
    return m.result[0]
}

func (m *FuncResult) ValueNum2() (interface{}, interface{}) {
    return m.result[0], m.result[1]
}

func (m *FuncResult) ValueNum3() (interface{}, interface{}, interface{}) {
    return m.result[0], m.result[1], m.result[2]
}

func (m *FuncResult) ValueOf(index int) interface{} {
    return m.result[index]
}

func (m *FuncResult) ValueAll() []interface{} {
    return m.result
}

func (m *FuncResult) ValueSlice(start, end int) []interface{} {
    return m.result[start:end]
}

func (m *FuncResult) ValueAndError() (interface{}, error) {
    return m.result[0], m.result[1].(error)
}

func (m *FuncResult) ValueAndBool() (interface{}, bool) {
    return m.result[0], m.result[1].(bool)
}

func (m *FuncResult) ValueLen() int {
    return len(m.result)
}
