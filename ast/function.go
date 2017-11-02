package ast

import "github.com/arnodel/golua/ir"

type FunctionName struct {
	nameList []Name
	method   Name
}

func (n FunctionName) HWrite(w HWriter) {
	for i, name := range n.nameList {
		if i > 0 {
			w.Writef(".")
		}
		w.Writef(string(name))
	}
	if n.method != "" {
		w.Writef(":%s", n.method)
	}
}

type Function struct {
	ParList
	body Stat
}

func (f Function) HWrite(w HWriter) {
	w.Writef("(")
	for i, param := range f.params {
		w.Writef(string(param))
		if i < len(f.params)-1 || f.hasDots {
			w.Writef(", ")
		}
	}
	if f.hasDots {
		w.Writef("...")
	}
	w.Writef(")")
	w.Indent()
	w.Next()
	f.body.HWrite(w)
	w.Dedent()
}

func (f Function) CompileExp(c *Compiler) ir.Register {
	fc := NewCompiler(c)
	var recvRegs []ir.Register
	for _, p := range f.params {
		reg := fc.DeclareLocal(p)
		recvRegs = append(recvRegs, reg)
	}
	if !f.hasDots {
		fc.Emit(ir.Receive{Dst: recvRegs})
	} else {
		etc := fc.DeclareLocal(Name("..."))
		fc.Emit(ir.ReceiveEtc{Dst: recvRegs, Etc: etc})
	}
	f.body.CompileStat(fc)
	kidx := c.GetConstant(fc.code)
	reg := c.NewRegister()
	c.Emit(ir.MkClosure{
		Dst:      reg,
		Code:     kidx,
		Upvalues: fc.upvalues,
	})
	return reg
}

type FunctionStat struct {
	Function
	name FunctionName
}

func (s FunctionStat) HWrite(w HWriter) {
	w.Writef("function ")
	s.name.HWrite(w)
	s.Function.HWrite(w)
}

type LocalFunctionStat struct {
	Function
	name Name
}

func (s LocalFunctionStat) HWrite(w HWriter) {
	w.Writef("local function ")
	s.name.HWrite(w)
	s.Function.HWrite(w)
}

func (s LocalFunctionStat) CompileStat(c *Compiler) {
	// TODO
}

type ParList struct {
	params  []Name
	hasDots bool
}

func NewParList(params []Name, hasDots bool) (ParList, error) {
	return ParList{
		params:  params,
		hasDots: hasDots,
	}, nil
}

func NewFunctionName(names []Name, method Name) (FunctionName, error) {
	return FunctionName{
		nameList: names,
		method:   method,
	}, nil
}

func NewFunction(parList ParList, body Stat) (Function, error) {
	return Function{ParList: parList, body: body}, nil
}

func NewFunctionStat(name FunctionName, fx Function) (FunctionStat, error) {
	return FunctionStat{Function: fx, name: name}, nil
}

func NewLocalFunctionStat(name Name, fx Function) (LocalFunctionStat, error) {
	return LocalFunctionStat{Function: fx, name: name}, nil
}
