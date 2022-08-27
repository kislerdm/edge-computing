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

				var r, g, b int = args[0].Int(), args[1].Int(), args[2].Int()

				if r < 0 || r > 255 || g < 0 || g > 255 || b < 0 || b > 255 {
					return map[string]interface{}{"error": "wrong RGB input"}
				}

				o, err := logic.Start(uint8(r), uint8(g), uint8(b))
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
