package shaders

var Enemy_Shader = `//kage:unit pixels
			package main

			var I_Frames float

			func Fragment(targetCoords vec4, srcPos vec2, _ vec4) vec4 {
				col := imageSrc0At(srcPos.xy)
				if col.w != 0 {
					return vec4(col.x + I_Frames, col.y + I_Frames, col.z + I_Frames, col.w)
				} else {
					return col
				}
			}
`

var Player_Shader = `//kage:unit pixels
			package main

			var I_Frames float

			func Fragment(targetCoords vec4, srcPos vec2, _ vec4) vec4 {
				col := imageSrc0At(srcPos.xy)
				if col.w != 0 {
					return vec4(col.x + I_Frames, col.y + I_Frames, col.z + I_Frames, col.w)
				} else {
					return col
				}
			}
`

var Test_Guy_Ball_Shader = `//kage:unit pixels
	package main

	var X float
	var Y float
	var Z float

	func Fragment(targetCoords vec4, srcPos vec2, _ vec4) vec4 {
		col := imageSrc0At(srcPos.xy)
		if col.w != 0 {
			return vec4((col.x + X)/2, (col.y + Y)/2, (col.z + Z)/2, col.w)
		} else {
			return col
		}
	}`
