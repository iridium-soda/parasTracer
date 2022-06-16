# parasTracer
一个用于[Massive Code Runner](https://github.com/iridium-soda/massive-coderunner)的原理验证小项目.用于追踪函数参数的传播.
# Description
检查函数中每个参数的传播路径,返回和每个函数形参*相关*的变量.
> 相关:被形参和形参相关的变量赋值或为相关运算的返回结果,或作为形参传入函数的返回值.

Example:

```go
package main

func funcA(parA int, parB int, parC int)int {
	var a = parA + parB
	var b=a+parC
	var d=funcB(a,b)
	return d
	
}
func funcB(a int,b int)int{
	return a+b
}
func main() {
	funcA(1,2,3)
}
```
则分析结果应为:
- 和`parA`相关的变量有:`a`,`b`(由`a`经由`+`传导),d(由`a`或`b`经由`funcB`传导)
- 和`parB`相关的变量有:`a`,`b`,`d`
- 和`parC`相关的变量有:`b`,`d`
# Usage
```shell
parasTracer <filename>
```
Output:
```
funcName: <Name> paraName:<Name> releated vars:[<Name>,...]
```
