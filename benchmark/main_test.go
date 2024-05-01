package benchmark

import (
	"flag"
	"fmt"
	"monkey/compiler"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"monkey/vm"
	"testing"
	"time"
)

var input = `
let fibonacci = fn(x) {
  if (x == 0) {
    0
  } else {
    if (x == 1) {
      return 1;
    } else {
      fibonacci(x - 1) + fibonacci(x - 2);
    }
  }
};
fibonacci(35);
`

func Test(t *testing.T) {
	flag.Parse()

	var (
		duration1 time.Duration
		duration2 time.Duration
		result1   object.Object
		result2   object.Object
	)

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	comp := compiler.New()
	err := comp.Compile(program)
	if err != nil {
		fmt.Printf("compiler error: %s", err)
		return
	}

	machine := vm.New(comp.Bytecode())

	start := time.Now()

	err = machine.Run()
	if err != nil {
		fmt.Printf("vm error: %s", err)
		return
	}
	duration1 = time.Since(start)
	result1 = machine.LastPoppedStackElem()

	env := object.NewEnvironment()
	start = time.Now()
	result2 = evaluator.Eval(program, env)
	duration2 = time.Since(start)

	fmt.Printf(
		"engine=vm, result=%s, duration=%s\n",
		result1.Inspect(),
		duration1)

	fmt.Printf(
		"engine=eval, result=%s, duration=%s\n",
		result2.Inspect(),
		duration2)
}
