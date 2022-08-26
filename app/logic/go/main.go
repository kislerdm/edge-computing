//go:build js && !unittest
// +build js,!unittest

package main

import (
	"syscall/js"

	"edgecomputing/logic"
)

func main() {
	js.Global().Set(
		"start", js.FuncOf(
			func(this js.Value, args []js.Value) interface{} {
				if len(args) < 3 {
					return map[string]interface{}{"error": "no r, g, b input provided"}
				}

				o, err := logic.Start(
					args[0].Float(),
					args[1].Float(),
					args[2].Float(),
				)

				if err != nil {
					return map[string]interface{}{"error": err.Error()}
				}

				return map[string]interface{}{
					"name":    o.Name,
					"is_warm": o.IsWarm,
				}
			},
		),
	)
	select {}
}
