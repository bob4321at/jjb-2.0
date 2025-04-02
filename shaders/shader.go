package shaders

var Base_Shader = `//kage:unit pixels
			package main

			func Fragment(targetCoords vec4, srcPos vec2, _ vec4) vec4 {
				col := imageSrc0At(srcPos.xy)
				return vec4(col.x, col.y, col.z, col.w)
			}
`

var Snake_Ball_Shader = `//kage:unit pixels
			package main

			func Fragment(targetCoords vec4, srcPos vec2, _ vec4) vec4 {
				col := imageSrc0At(srcPos.xy)
				return vec4(col.x, col.y, col.z, col.w)
			}
`

var Snake_Head_Shader = `//kage:unit pixels
			package main

			func Fragment(targetCoords vec4, srcPos vec2, _ vec4) vec4 {
				col := imageSrc0At(srcPos.xy)
				return vec4(col.x, col.y, col.z, col.w)
			}
`

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
