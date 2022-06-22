package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

var fileName string
var usage string = `To analysis and trace paras of functions.
Usage:
parasTracer <filename>`

func printUsage() {
	fmt.Fprintln(
		os.Stderr,
		usage,
	)
	return
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 { //Error handle
		printUsage()
		return
	}
	fileName = flag.Args()[0]
	doAnalysis()
}
func doAnalysis() {
	//解析最上一层的顶级声明,再将函数信息传入下一层函数分析
	fileSet := token.NewFileSet()
	f, err := parser.ParseFile(fileSet, fileName, "", 0)
	if err != nil {
		fmt.Printf("err = %s", err)
	}
	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if ok { //Got Function
			fmt.Printf("Got function decl, name is %s", fn.Name.Name)
			analysisFunc(fn)
		}

	}
}

//TODO: 注意：循环内的变量分析怎么做？
//TODO: 注意:分支条件结构的变量分析怎么做?
func analysisFunc(fn *ast.FuncDecl) {
	//分析单个函数的传播链
	//获取形参
	//注意形如 s0, s1 string, s2 string这样的声明, s0和s1是在一个field里所以需要再加一个循环解析
	var parasList []string //暂存参数表
	for _, paraField := range fn.Type.Params.List {
		for _, singlePara := range paraField.Names {
			fmt.Printf("Got function para, name is %s", singlePara.Name)
			parasList = append(parasList, singlePara.Name)
		}
	}
	//建立参数传播存储结构
	//每个list第一个为形参(变量传染根),后面的应为不重复的变量名.
	var parasInfo [][]string
	for index, paraName := range parasList {
		parasInfo = append(parasInfo, make([]string, 1))
		parasInfo[index][0] = paraName
	}
	//此时parasInfo应该为[[a],[b],[c],...]

	//准备分析,应该是FuncDecl.Body
	for _, _ = range fn.Body.List {
		//TODO:考虑内部仍然有代码块的情况
		//应该只有stmt类型和BlockStmt类型

	}
}
