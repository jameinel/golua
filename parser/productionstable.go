// Code generated by gocc; DO NOT EDIT.

package parser

import "github.com/arnodel/golua/ast"
import "github.com/arnodel/golua/token"
import "github.com/arnodel/golua/ops"

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : Chunk	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Chunk : StatList "return" ExpList ";"	<< ast.NewBlockStat(X[0].([]ast.Stat), X[2].([]ast.ExpNode)) >>`,
		Id:         "Chunk",
		NTType:     1,
		Index:      1,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBlockStat(X[0].([]ast.Stat), X[2].([]ast.ExpNode))
		},
	},
	ProdTabEntry{
		String: `Chunk : StatList "return" ExpList	<< ast.NewBlockStat(X[0].([]ast.Stat), X[2].([]ast.ExpNode)) >>`,
		Id:         "Chunk",
		NTType:     1,
		Index:      2,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBlockStat(X[0].([]ast.Stat), X[2].([]ast.ExpNode))
		},
	},
	ProdTabEntry{
		String: `Chunk : StatList "return" ";"	<< ast.NewBlockStat(X[0].([]ast.Stat), []ast.ExpNode{}) >>`,
		Id:         "Chunk",
		NTType:     1,
		Index:      3,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBlockStat(X[0].([]ast.Stat), []ast.ExpNode{})
		},
	},
	ProdTabEntry{
		String: `Chunk : StatList "return"	<< ast.NewBlockStat(X[0].([]ast.Stat), []ast.ExpNode{}) >>`,
		Id:         "Chunk",
		NTType:     1,
		Index:      4,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBlockStat(X[0].([]ast.Stat), []ast.ExpNode{})
		},
	},
	ProdTabEntry{
		String: `Chunk : StatList	<< ast.NewBlockStat(X[0].([]ast.Stat), nil) >>`,
		Id:         "Chunk",
		NTType:     1,
		Index:      5,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBlockStat(X[0].([]ast.Stat), nil)
		},
	},
	ProdTabEntry{
		String: `StatList : empty	<< []ast.Stat{}, nil >>`,
		Id:         "StatList",
		NTType:     2,
		Index:      6,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return []ast.Stat{}, nil
		},
	},
	ProdTabEntry{
		String: `StatList : StatList Stat	<< append(X[0].([]ast.Stat), X[1].(ast.Stat)), nil >>`,
		Id:         "StatList",
		NTType:     2,
		Index:      7,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return append(X[0].([]ast.Stat), X[1].(ast.Stat)), nil
		},
	},
	ProdTabEntry{
		String: `Stat : ";"	<< ast.NewEmptyStat() >>`,
		Id:         "Stat",
		NTType:     3,
		Index:      8,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewEmptyStat()
		},
	},
	ProdTabEntry{
		String: `Stat : AssignStat	<<  >>`,
		Id:         "Stat",
		NTType:     3,
		Index:      9,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Stat : FunctionCall ";"	<<  >>`,
		Id:         "Stat",
		NTType:     3,
		Index:      10,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Stat : Label	<<  >>`,
		Id:         "Stat",
		NTType:     3,
		Index:      11,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Stat : BreakStat	<<  >>`,
		Id:         "Stat",
		NTType:     3,
		Index:      12,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Stat : GotoStat	<<  >>`,
		Id:         "Stat",
		NTType:     3,
		Index:      13,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Stat : BlockStat	<<  >>`,
		Id:         "Stat",
		NTType:     3,
		Index:      14,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Stat : WhileStat	<<  >>`,
		Id:         "Stat",
		NTType:     3,
		Index:      15,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Stat : RepeatStat	<<  >>`,
		Id:         "Stat",
		NTType:     3,
		Index:      16,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Stat : IfStat	<<  >>`,
		Id:         "Stat",
		NTType:     3,
		Index:      17,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Stat : ForStat	<<  >>`,
		Id:         "Stat",
		NTType:     3,
		Index:      18,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Stat : ForInStat	<<  >>`,
		Id:         "Stat",
		NTType:     3,
		Index:      19,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Stat : FunctionStat	<<  >>`,
		Id:         "Stat",
		NTType:     3,
		Index:      20,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Stat : LocalFunctionStat	<<  >>`,
		Id:         "Stat",
		NTType:     3,
		Index:      21,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Stat : LocalStat	<<  >>`,
		Id:         "Stat",
		NTType:     3,
		Index:      22,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `AssignStat : VarList "=" ExpList	<< ast.NewAssignStat(X[0].([]ast.Var), X[2].([]ast.ExpNode)) >>`,
		Id:         "AssignStat",
		NTType:     4,
		Index:      23,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewAssignStat(X[0].([]ast.Var), X[2].([]ast.ExpNode))
		},
	},
	ProdTabEntry{
		String: `BreakStat : "break"	<< ast.NewBreakStat() >>`,
		Id:         "BreakStat",
		NTType:     5,
		Index:      24,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBreakStat()
		},
	},
	ProdTabEntry{
		String: `GotoStat : "goto" Name	<< ast.NewGotoStat(X[1].(ast.Name)) >>`,
		Id:         "GotoStat",
		NTType:     6,
		Index:      25,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewGotoStat(X[1].(ast.Name))
		},
	},
	ProdTabEntry{
		String: `BlockStat : "do" Chunk "end"	<< X[1], nil >>`,
		Id:         "BlockStat",
		NTType:     7,
		Index:      26,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `WhileStat : "while" Exp "do" Chunk "end"	<< ast.NewWhileStat(X[1].(ast.ExpNode), X[3].(ast.BlockStat)) >>`,
		Id:         "WhileStat",
		NTType:     8,
		Index:      27,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewWhileStat(X[1].(ast.ExpNode), X[3].(ast.BlockStat))
		},
	},
	ProdTabEntry{
		String: `RepeatStat : "repeat" Chunk "until" Exp	<< ast.NewRepeatStat(X[1].(ast.BlockStat), X[3].(ast.ExpNode)) >>`,
		Id:         "RepeatStat",
		NTType:     9,
		Index:      28,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewRepeatStat(X[1].(ast.BlockStat), X[3].(ast.ExpNode))
		},
	},
	ProdTabEntry{
		String: `ElseIf : Exp "then" Chunk	<< ast.NewIfStat().AddElseIf(X[0].(ast.ExpNode), X[2].(ast.BlockStat)) >>`,
		Id:         "ElseIf",
		NTType:     10,
		Index:      29,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIfStat().AddElseIf(X[0].(ast.ExpNode), X[2].(ast.BlockStat))
		},
	},
	ProdTabEntry{
		String: `ElseIf : ElseIf "elseif" Exp "then" Chunk	<< X[0].(ast.IfStat).AddElseIf(X[2].(ast.ExpNode), X[4].(ast.BlockStat)) >>`,
		Id:         "ElseIf",
		NTType:     10,
		Index:      30,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0].(ast.IfStat).AddElseIf(X[2].(ast.ExpNode), X[4].(ast.BlockStat))
		},
	},
	ProdTabEntry{
		String: `Else : "elseif" ElseIf "else" Chunk "end"	<< X[1].(ast.IfStat).AddElse(X[3].(ast.BlockStat)) >>`,
		Id:         "Else",
		NTType:     11,
		Index:      31,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1].(ast.IfStat).AddElse(X[3].(ast.BlockStat))
		},
	},
	ProdTabEntry{
		String: `Else : "elseif" ElseIf "end"	<< X[1], nil >>`,
		Id:         "Else",
		NTType:     11,
		Index:      32,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `Else : "else" Chunk "end"	<< ast.NewIfStat().AddElse(X[1].(ast.BlockStat)) >>`,
		Id:         "Else",
		NTType:     11,
		Index:      33,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIfStat().AddElse(X[1].(ast.BlockStat))
		},
	},
	ProdTabEntry{
		String: `Else : "end"	<< ast.NewIfStat(), nil >>`,
		Id:         "Else",
		NTType:     11,
		Index:      34,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIfStat(), nil
		},
	},
	ProdTabEntry{
		String: `IfStat : "if" Exp "then" Chunk Else	<< X[4].(ast.IfStat).AddIf(X[1].(ast.ExpNode), X[3].(ast.BlockStat)) >>`,
		Id:         "IfStat",
		NTType:     12,
		Index:      35,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[4].(ast.IfStat).AddIf(X[1].(ast.ExpNode), X[3].(ast.BlockStat))
		},
	},
	ProdTabEntry{
		String: `ForList : Exp "," Exp	<< []ast.ExpNode{X[0].(ast.ExpNode), X[2].(ast.ExpNode), ast.Int(1)}, nil >>`,
		Id:         "ForList",
		NTType:     13,
		Index:      36,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return []ast.ExpNode{X[0].(ast.ExpNode), X[2].(ast.ExpNode), ast.Int(1)}, nil
		},
	},
	ProdTabEntry{
		String: `ForList : Exp "," Exp "," Exp	<< []ast.ExpNode{X[0].(ast.ExpNode), X[2].(ast.ExpNode), X[4].(ast.ExpNode)}, nil >>`,
		Id:         "ForList",
		NTType:     13,
		Index:      37,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return []ast.ExpNode{X[0].(ast.ExpNode), X[2].(ast.ExpNode), X[4].(ast.ExpNode)}, nil
		},
	},
	ProdTabEntry{
		String: `ForStat : "for" Name "=" ForList "do" Chunk "end"	<< ast.NewForStat(X[1].(ast.Name), X[3].([]ast.ExpNode), X[5].(ast.BlockStat)) >>`,
		Id:         "ForStat",
		NTType:     14,
		Index:      38,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewForStat(X[1].(ast.Name), X[3].([]ast.ExpNode), X[5].(ast.BlockStat))
		},
	},
	ProdTabEntry{
		String: `ForInStat : "for" NameList "in" ExpList "do" Chunk "end"	<< ast.NewForInStat(X[1].([]ast.Name), X[3].([]ast.ExpNode), X[5].(ast.BlockStat)) >>`,
		Id:         "ForInStat",
		NTType:     15,
		Index:      39,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewForInStat(X[1].([]ast.Name), X[3].([]ast.ExpNode), X[5].(ast.BlockStat))
		},
	},
	ProdTabEntry{
		String: `FunctionStat : "function" FuncName FuncBody	<< ast.NewFunctionStat(X[1].(ast.FunctionName), X[2].(ast.Function)) >>`,
		Id:         "FunctionStat",
		NTType:     16,
		Index:      40,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFunctionStat(X[1].(ast.FunctionName), X[2].(ast.Function))
		},
	},
	ProdTabEntry{
		String: `LocalFunctionStat : "local" "function" Name FuncBody	<< ast.NewLocalFunctionStat(X[2].(ast.Name), X[3].(ast.Function)) >>`,
		Id:         "LocalFunctionStat",
		NTType:     17,
		Index:      41,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLocalFunctionStat(X[2].(ast.Name), X[3].(ast.Function))
		},
	},
	ProdTabEntry{
		String: `LocalStat : "local" NameList	<< ast.NewLocalStat(X[1].([]ast.Name), nil) >>`,
		Id:         "LocalStat",
		NTType:     18,
		Index:      42,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLocalStat(X[1].([]ast.Name), nil)
		},
	},
	ProdTabEntry{
		String: `LocalStat : "local" NameList "=" ExpList	<< ast.NewLocalStat(X[1].([]ast.Name), X[3].([]ast.ExpNode)) >>`,
		Id:         "LocalStat",
		NTType:     18,
		Index:      43,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLocalStat(X[1].([]ast.Name), X[3].([]ast.ExpNode))
		},
	},
	ProdTabEntry{
		String: `Label : "::" Name "::"	<< ast.NewLabelStat(X[0].(ast.Name)) >>`,
		Id:         "Label",
		NTType:     19,
		Index:      44,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLabelStat(X[0].(ast.Name))
		},
	},
	ProdTabEntry{
		String: `DottedName : Name	<< X[0].(ast.Name), nil >>`,
		Id:         "DottedName",
		NTType:     20,
		Index:      45,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0].(ast.Name), nil
		},
	},
	ProdTabEntry{
		String: `DottedName : DottedName "." Name	<< ast.NewIndexExp(X[0].(ast.ExpNode), ast.String(X[2].(ast.Name))) >>`,
		Id:         "DottedName",
		NTType:     20,
		Index:      46,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIndexExp(X[0].(ast.ExpNode), ast.String(X[2].(ast.Name)))
		},
	},
	ProdTabEntry{
		String: `FuncName : DottedName	<< ast.NewFunctionName(X[0].(ast.Var), ast.Name("")) >>`,
		Id:         "FuncName",
		NTType:     21,
		Index:      47,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFunctionName(X[0].(ast.Var), ast.Name(""))
		},
	},
	ProdTabEntry{
		String: `FuncName : DottedName ":" Name	<< ast.NewFunctionName(X[0].(ast.Var), X[2].(ast.Name)) >>`,
		Id:         "FuncName",
		NTType:     21,
		Index:      48,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFunctionName(X[0].(ast.Var), X[2].(ast.Name))
		},
	},
	ProdTabEntry{
		String: `VarList : Var	<< []ast.Var{X[0].(ast.Var)}, nil >>`,
		Id:         "VarList",
		NTType:     22,
		Index:      49,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return []ast.Var{X[0].(ast.Var)}, nil
		},
	},
	ProdTabEntry{
		String: `VarList : VarList "," Var	<< append(X[0].([]ast.Var), X[2].(ast.Var)), nil >>`,
		Id:         "VarList",
		NTType:     22,
		Index:      50,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return append(X[0].([]ast.Var), X[2].(ast.Var)), nil
		},
	},
	ProdTabEntry{
		String: `NameList : Name	<< []ast.Name{X[0].(ast.Name)}, nil >>`,
		Id:         "NameList",
		NTType:     23,
		Index:      51,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return []ast.Name{X[0].(ast.Name)}, nil
		},
	},
	ProdTabEntry{
		String: `NameList : NameList "," Name	<< append(X[0].([]ast.Name), X[2].(ast.Name)), nil >>`,
		Id:         "NameList",
		NTType:     23,
		Index:      52,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return append(X[0].([]ast.Name), X[2].(ast.Name)), nil
		},
	},
	ProdTabEntry{
		String: `ExpList : Exp	<< []ast.ExpNode{X[0].(ast.ExpNode)}, nil >>`,
		Id:         "ExpList",
		NTType:     24,
		Index:      53,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return []ast.ExpNode{X[0].(ast.ExpNode)}, nil
		},
	},
	ProdTabEntry{
		String: `ExpList : ExpList "," Exp	<< append(X[0].([]ast.ExpNode), X[2].(ast.ExpNode)), nil >>`,
		Id:         "ExpList",
		NTType:     24,
		Index:      54,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return append(X[0].([]ast.ExpNode), X[2].(ast.ExpNode)), nil
		},
	},
	ProdTabEntry{
		String: `Exp : AndExp	<<  >>`,
		Id:         "Exp",
		NTType:     25,
		Index:      55,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Exp : Exp "or" AndExp	<< ast.NewBinOp(X[0].(ast.ExpNode), ops.OpOr, X[2].(ast.ExpNode)) >>`,
		Id:         "Exp",
		NTType:     25,
		Index:      56,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBinOp(X[0].(ast.ExpNode), ops.OpOr, X[2].(ast.ExpNode))
		},
	},
	ProdTabEntry{
		String: `AndExp : CompExp	<<  >>`,
		Id:         "AndExp",
		NTType:     26,
		Index:      57,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `AndExp : AndExp "and" CompExp	<< ast.NewBinOp(X[0].(ast.ExpNode), ops.OpAnd, X[2].(ast.ExpNode)) >>`,
		Id:         "AndExp",
		NTType:     26,
		Index:      58,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBinOp(X[0].(ast.ExpNode), ops.OpAnd, X[2].(ast.ExpNode))
		},
	},
	ProdTabEntry{
		String: `CompOp : "<"	<< ops.OpLt, nil >>`,
		Id:         "CompOp",
		NTType:     27,
		Index:      59,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ops.OpLt, nil
		},
	},
	ProdTabEntry{
		String: `CompOp : "<="	<< ops.OpLeq, nil >>`,
		Id:         "CompOp",
		NTType:     27,
		Index:      60,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ops.OpLeq, nil
		},
	},
	ProdTabEntry{
		String: `CompOp : ">"	<< ops.OpGt, nil >>`,
		Id:         "CompOp",
		NTType:     27,
		Index:      61,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ops.OpGt, nil
		},
	},
	ProdTabEntry{
		String: `CompOp : ">="	<< ops.OpGeq, nil >>`,
		Id:         "CompOp",
		NTType:     27,
		Index:      62,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ops.OpGeq, nil
		},
	},
	ProdTabEntry{
		String: `CompOp : "=="	<< ops.OpEq, nil >>`,
		Id:         "CompOp",
		NTType:     27,
		Index:      63,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ops.OpEq, nil
		},
	},
	ProdTabEntry{
		String: `CompOp : "~="	<< ops.OpNeq, nil >>`,
		Id:         "CompOp",
		NTType:     27,
		Index:      64,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ops.OpNeq, nil
		},
	},
	ProdTabEntry{
		String: `CompExp : BitOrExp	<<  >>`,
		Id:         "CompExp",
		NTType:     28,
		Index:      65,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `CompExp : CompExp CompOp BitOrExp	<< ast.NewBinOp(X[0].(ast.ExpNode), X[1].(ops.Op), X[2].(ast.ExpNode)) >>`,
		Id:         "CompExp",
		NTType:     28,
		Index:      66,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBinOp(X[0].(ast.ExpNode), X[1].(ops.Op), X[2].(ast.ExpNode))
		},
	},
	ProdTabEntry{
		String: `BitOrExp : BitXorExp	<<  >>`,
		Id:         "BitOrExp",
		NTType:     29,
		Index:      67,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `BitOrExp : BitOrExp "|" BitXorExp	<< ast.NewBinOp(X[0].(ast.ExpNode), ops.OpBitOr, X[2].(ast.ExpNode)) >>`,
		Id:         "BitOrExp",
		NTType:     29,
		Index:      68,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBinOp(X[0].(ast.ExpNode), ops.OpBitOr, X[2].(ast.ExpNode))
		},
	},
	ProdTabEntry{
		String: `BitXorExp : BitAndExp	<<  >>`,
		Id:         "BitXorExp",
		NTType:     30,
		Index:      69,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `BitXorExp : BitXorExp "~" BitAndExp	<< ast.NewBinOp(X[0].(ast.ExpNode), ops.OpBitXor, X[2].(ast.ExpNode)) >>`,
		Id:         "BitXorExp",
		NTType:     30,
		Index:      70,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBinOp(X[0].(ast.ExpNode), ops.OpBitXor, X[2].(ast.ExpNode))
		},
	},
	ProdTabEntry{
		String: `BitAndExp : ConcatExp	<<  >>`,
		Id:         "BitAndExp",
		NTType:     31,
		Index:      71,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `BitAndExp : BitAndExp "&" ConcatExp	<< ast.NewBinOp(X[0].(ast.ExpNode), ops.OpBitAnd, X[2].(ast.ExpNode)) >>`,
		Id:         "BitAndExp",
		NTType:     31,
		Index:      72,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBinOp(X[0].(ast.ExpNode), ops.OpBitAnd, X[2].(ast.ExpNode))
		},
	},
	ProdTabEntry{
		String: `ConcatExp : ShiftExp	<<  >>`,
		Id:         "ConcatExp",
		NTType:     32,
		Index:      73,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ConcatExp : ConcatExp ".." ShiftExp	<< ast.NewBinOp(X[0].(ast.ExpNode), ops.OpConcat, X[2].(ast.ExpNode)) >>`,
		Id:         "ConcatExp",
		NTType:     32,
		Index:      74,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBinOp(X[0].(ast.ExpNode), ops.OpConcat, X[2].(ast.ExpNode))
		},
	},
	ProdTabEntry{
		String: `ShiftOp : "<<"	<< ops.OpShiftL, nil >>`,
		Id:         "ShiftOp",
		NTType:     33,
		Index:      75,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ops.OpShiftL, nil
		},
	},
	ProdTabEntry{
		String: `ShiftOp : ">>"	<< ops.OpShiftR, nil >>`,
		Id:         "ShiftOp",
		NTType:     33,
		Index:      76,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ops.OpShiftR, nil
		},
	},
	ProdTabEntry{
		String: `ShiftExp : Sum	<<  >>`,
		Id:         "ShiftExp",
		NTType:     34,
		Index:      77,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ShiftExp : ShiftExp ShiftOp Sum	<< ast.NewBinOp(X[0].(ast.ExpNode), X[1].(ops.Op), X[2].(ast.ExpNode)) >>`,
		Id:         "ShiftExp",
		NTType:     34,
		Index:      78,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBinOp(X[0].(ast.ExpNode), X[1].(ops.Op), X[2].(ast.ExpNode))
		},
	},
	ProdTabEntry{
		String: `SumOp : "+"	<< ops.OpAdd, nil >>`,
		Id:         "SumOp",
		NTType:     35,
		Index:      79,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ops.OpAdd, nil
		},
	},
	ProdTabEntry{
		String: `SumOp : "-"	<< ops.OpSub, nil >>`,
		Id:         "SumOp",
		NTType:     35,
		Index:      80,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ops.OpSub, nil
		},
	},
	ProdTabEntry{
		String: `Sum : Term	<<  >>`,
		Id:         "Sum",
		NTType:     36,
		Index:      81,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Sum : Sum SumOp Term	<< ast.NewBinOp(X[0].(ast.ExpNode), X[1].(ops.Op), X[2].(ast.ExpNode)) >>`,
		Id:         "Sum",
		NTType:     36,
		Index:      82,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBinOp(X[0].(ast.ExpNode), X[1].(ops.Op), X[2].(ast.ExpNode))
		},
	},
	ProdTabEntry{
		String: `TermOp : "*"	<< ops.OpMul, nil >>`,
		Id:         "TermOp",
		NTType:     37,
		Index:      83,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ops.OpMul, nil
		},
	},
	ProdTabEntry{
		String: `TermOp : "/"	<< ops.OpDiv, nil >>`,
		Id:         "TermOp",
		NTType:     37,
		Index:      84,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ops.OpDiv, nil
		},
	},
	ProdTabEntry{
		String: `TermOp : "%!"(MISSING)	<< ops.OpMod, nil >>`,
		Id:         "TermOp",
		NTType:     37,
		Index:      85,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ops.OpMod, nil
		},
	},
	ProdTabEntry{
		String: `TermOp : "//"	<< ops.OpFloorDiv, nil >>`,
		Id:         "TermOp",
		NTType:     37,
		Index:      86,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ops.OpFloorDiv, nil
		},
	},
	ProdTabEntry{
		String: `Term : Factor	<<  >>`,
		Id:         "Term",
		NTType:     38,
		Index:      87,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Term : Term TermOp Factor	<< ast.NewBinOp(X[0].(ast.ExpNode), X[1].(ops.Op), X[2].(ast.ExpNode)) >>`,
		Id:         "Term",
		NTType:     38,
		Index:      88,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBinOp(X[0].(ast.ExpNode), X[1].(ops.Op), X[2].(ast.ExpNode))
		},
	},
	ProdTabEntry{
		String: `UnOp : "-"	<< ops.OpNeg, nil >>`,
		Id:         "UnOp",
		NTType:     39,
		Index:      89,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ops.OpNeg, nil
		},
	},
	ProdTabEntry{
		String: `UnOp : "not"	<< ops.OpNot, nil >>`,
		Id:         "UnOp",
		NTType:     39,
		Index:      90,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ops.OpNot, nil
		},
	},
	ProdTabEntry{
		String: `UnOp : "#"	<< ops.OpLen, nil >>`,
		Id:         "UnOp",
		NTType:     39,
		Index:      91,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ops.OpLen, nil
		},
	},
	ProdTabEntry{
		String: `UnOp : "~"	<< ops.OpBitNot, nil >>`,
		Id:         "UnOp",
		NTType:     39,
		Index:      92,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ops.OpBitNot, nil
		},
	},
	ProdTabEntry{
		String: `Factor : UnOp Factor	<< ast.NewUnOp(X[0].(ops.Op), X[1].(ast.ExpNode)) >>`,
		Id:         "Factor",
		NTType:     40,
		Index:      93,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewUnOp(X[0].(ops.Op), X[1].(ast.ExpNode))
		},
	},
	ProdTabEntry{
		String: `Factor : Power	<<  >>`,
		Id:         "Factor",
		NTType:     40,
		Index:      94,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Power : Atom	<<  >>`,
		Id:         "Power",
		NTType:     41,
		Index:      95,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Power : Atom "^" Factor	<< ast.NewBinOp(X[0].(ast.ExpNode), ops.OpPow, X[2].(ast.ExpNode)) >>`,
		Id:         "Power",
		NTType:     41,
		Index:      96,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBinOp(X[0].(ast.ExpNode), ops.OpPow, X[2].(ast.ExpNode))
		},
	},
	ProdTabEntry{
		String: `Var : Name	<<  >>`,
		Id:         "Var",
		NTType:     42,
		Index:      97,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Var : PrefixExp "[" Exp "]"	<< ast.NewIndexExp(X[0].(ast.ExpNode), X[2].(ast.ExpNode)) >>`,
		Id:         "Var",
		NTType:     42,
		Index:      98,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIndexExp(X[0].(ast.ExpNode), X[2].(ast.ExpNode))
		},
	},
	ProdTabEntry{
		String: `Var : PrefixExp "." Name	<< ast.NewIndexExp(X[0].(ast.ExpNode), X[2].(ast.ExpNode)) >>`,
		Id:         "Var",
		NTType:     42,
		Index:      99,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIndexExp(X[0].(ast.ExpNode), X[2].(ast.ExpNode))
		},
	},
	ProdTabEntry{
		String: `PrefixExp : Var	<<  >>`,
		Id:         "PrefixExp",
		NTType:     43,
		Index:      100,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `PrefixExp : FunctionCall	<<  >>`,
		Id:         "PrefixExp",
		NTType:     43,
		Index:      101,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `PrefixExp : "(" Exp ")"	<< X[1], nil >>`,
		Id:         "PrefixExp",
		NTType:     43,
		Index:      102,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `FunctionCall : PrefixExp Args	<< ast.NewFunctionCall(X[0].(ast.ExpNode), ast.Name(""), X[1].([]ast.ExpNode)) >>`,
		Id:         "FunctionCall",
		NTType:     44,
		Index:      103,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFunctionCall(X[0].(ast.ExpNode), ast.Name(""), X[1].([]ast.ExpNode))
		},
	},
	ProdTabEntry{
		String: `FunctionCall : PrefixExp ":" Name Args	<< ast.NewFunctionCall(X[0].(ast.ExpNode), X[2].(ast.Name), X[3].([]ast.ExpNode)) >>`,
		Id:         "FunctionCall",
		NTType:     44,
		Index:      104,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFunctionCall(X[0].(ast.ExpNode), X[2].(ast.Name), X[3].([]ast.ExpNode))
		},
	},
	ProdTabEntry{
		String: `Atom : "nil"	<< ast.Nil, nil >>`,
		Id:         "Atom",
		NTType:     45,
		Index:      105,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.Nil, nil
		},
	},
	ProdTabEntry{
		String: `Atom : "true"	<< ast.True, nil >>`,
		Id:         "Atom",
		NTType:     45,
		Index:      106,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.True, nil
		},
	},
	ProdTabEntry{
		String: `Atom : "false"	<< ast.False, nil >>`,
		Id:         "Atom",
		NTType:     45,
		Index:      107,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.False, nil
		},
	},
	ProdTabEntry{
		String: `Atom : "..."	<< ast.Etc, nil >>`,
		Id:         "Atom",
		NTType:     45,
		Index:      108,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.Etc, nil
		},
	},
	ProdTabEntry{
		String: `Atom : numdec	<< ast.NewNumber(X[0].(*token.Token)) >>`,
		Id:         "Atom",
		NTType:     45,
		Index:      109,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewNumber(X[0].(*token.Token))
		},
	},
	ProdTabEntry{
		String: `Atom : numhex	<< ast.NewNumber(X[0].(*token.Token)) >>`,
		Id:         "Atom",
		NTType:     45,
		Index:      110,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewNumber(X[0].(*token.Token))
		},
	},
	ProdTabEntry{
		String: `Atom : string	<< ast.NewString(X[0].(*token.Token)), nil >>`,
		Id:         "Atom",
		NTType:     45,
		Index:      111,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewString(X[0].(*token.Token)), nil
		},
	},
	ProdTabEntry{
		String: `Atom : FunctionDef	<<  >>`,
		Id:         "Atom",
		NTType:     45,
		Index:      112,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Atom : TableConstructor	<<  >>`,
		Id:         "Atom",
		NTType:     45,
		Index:      113,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Atom : PrefixExp	<<  >>`,
		Id:         "Atom",
		NTType:     45,
		Index:      114,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Args : "(" ")"	<< []ast.ExpNode{}, nil >>`,
		Id:         "Args",
		NTType:     46,
		Index:      115,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return []ast.ExpNode{}, nil
		},
	},
	ProdTabEntry{
		String: `Args : "(" ExpList ")"	<< X[1].([]ast.ExpNode), nil >>`,
		Id:         "Args",
		NTType:     46,
		Index:      116,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1].([]ast.ExpNode), nil
		},
	},
	ProdTabEntry{
		String: `Args : TableConstructor	<< []ast.ExpNode{X[0].(ast.ExpNode)}, nil >>`,
		Id:         "Args",
		NTType:     46,
		Index:      117,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return []ast.ExpNode{X[0].(ast.ExpNode)}, nil
		},
	},
	ProdTabEntry{
		String: `Args : string	<< []ast.ExpNode{ast.NewString(X[0].(*token.Token))}, nil >>`,
		Id:         "Args",
		NTType:     46,
		Index:      118,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return []ast.ExpNode{ast.NewString(X[0].(*token.Token))}, nil
		},
	},
	ProdTabEntry{
		String: `FunctionDef : "function" FuncBody	<< X[1], nil >>`,
		Id:         "FunctionDef",
		NTType:     47,
		Index:      119,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `ParList : NameList	<< ast.NewParList(X[0].([]ast.Name), false) >>`,
		Id:         "ParList",
		NTType:     48,
		Index:      120,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewParList(X[0].([]ast.Name), false)
		},
	},
	ProdTabEntry{
		String: `ParList : NameList "," "..."	<< ast.NewParList(X[0].([]ast.Name), true) >>`,
		Id:         "ParList",
		NTType:     48,
		Index:      121,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewParList(X[0].([]ast.Name), true)
		},
	},
	ProdTabEntry{
		String: `ParList : "..."	<< ast.NewParList(nil, true) >>`,
		Id:         "ParList",
		NTType:     48,
		Index:      122,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewParList(nil, true)
		},
	},
	ProdTabEntry{
		String: `ParList : empty	<< ast.NewParList(nil, false) >>`,
		Id:         "ParList",
		NTType:     48,
		Index:      123,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewParList(nil, false)
		},
	},
	ProdTabEntry{
		String: `FuncBody : "(" ParList ")" Chunk "end"	<< ast.NewFunction(X[1].(ast.ParList), X[3].(ast.BlockStat)) >>`,
		Id:         "FuncBody",
		NTType:     49,
		Index:      124,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFunction(X[1].(ast.ParList), X[3].(ast.BlockStat))
		},
	},
	ProdTabEntry{
		String: `TableConstructor : "{" FieldList "}"	<< ast.NewTableConstructor(X[1].([]ast.TableField)) >>`,
		Id:         "TableConstructor",
		NTType:     50,
		Index:      125,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewTableConstructor(X[1].([]ast.TableField))
		},
	},
	ProdTabEntry{
		String: `TableConstructor : "{" FieldList FieldSep "}"	<< ast.NewTableConstructor(X[1].([]ast.TableField)) >>`,
		Id:         "TableConstructor",
		NTType:     50,
		Index:      126,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewTableConstructor(X[1].([]ast.TableField))
		},
	},
	ProdTabEntry{
		String: `FieldSep : ","	<<  >>`,
		Id:         "FieldSep",
		NTType:     51,
		Index:      127,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `FieldSep : ";"	<<  >>`,
		Id:         "FieldSep",
		NTType:     51,
		Index:      128,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `FieldList : Field	<< []ast.TableField{X[0].(ast.TableField)}, nil >>`,
		Id:         "FieldList",
		NTType:     52,
		Index:      129,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return []ast.TableField{X[0].(ast.TableField)}, nil
		},
	},
	ProdTabEntry{
		String: `FieldList : FieldList FieldSep Field	<< append(X[0].([]ast.TableField), X[2].(ast.TableField)), nil >>`,
		Id:         "FieldList",
		NTType:     52,
		Index:      130,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return append(X[0].([]ast.TableField), X[2].(ast.TableField)), nil
		},
	},
	ProdTabEntry{
		String: `Field : "[" Exp "]" "=" Exp	<< ast.NewTableField(X[1].(ast.ExpNode), X[4].(ast.ExpNode)) >>`,
		Id:         "Field",
		NTType:     53,
		Index:      131,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewTableField(X[1].(ast.ExpNode), X[4].(ast.ExpNode))
		},
	},
	ProdTabEntry{
		String: `Field : Name "=" Exp	<< ast.NewTableField(X[0].(ast.ExpNode), X[2].(ast.ExpNode)) >>`,
		Id:         "Field",
		NTType:     53,
		Index:      132,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewTableField(X[0].(ast.ExpNode), X[2].(ast.ExpNode))
		},
	},
	ProdTabEntry{
		String: `Field : Exp	<< ast.NewTableField(ast.NoTableKey{}, X[0].(ast.ExpNode)) >>`,
		Id:         "Field",
		NTType:     53,
		Index:      133,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewTableField(ast.NoTableKey{}, X[0].(ast.ExpNode))
		},
	},
	ProdTabEntry{
		String: `Name : ident	<< ast.NewName(X[0].(*token.Token)) >>`,
		Id:         "Name",
		NTType:     54,
		Index:      134,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewName(X[0].(*token.Token))
		},
	},
}
