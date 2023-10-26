package main

import (
  "fmt"
  "syscall/js"
)

func concat(_ js.Value, args []js.Value) any {
  var result string
  for _, arg := range args {
    result += arg.String()
  }
  return result
}

func getElementById(_ js.Value, args []js.Value) any {
  id := args[0].String()
  return js.Global().Get("document").Call("getElementById", id)
}

func main() {
	fmt.Println("Go program started...")
  js.Global().Set("concat", js.FuncOf(concat))
  js.Global().Set("getElementById", js.FuncOf(getElementById))
  fmt.Println("Go funcitons setup complete.")
  select {}
}
