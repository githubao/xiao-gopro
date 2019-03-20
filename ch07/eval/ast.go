// 类型定义
// author: baoqiang
// time: 2019/3/19 下午8:15
package eval

type Expr interface {
	Eval(env Env) float64
	Check(vars map[Var]bool) error
}

type Var string

type literal float64

type unary struct {
	op rune
	x  Expr
}

type binary struct {
	op   rune
	x, y Expr
}

type call struct {
	fn   string
	args []Expr
}
