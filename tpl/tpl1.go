package tpl

// Tpl1 ...
const Tpl1 = `
{
	"presets": [
		[
			"@babel/preset-env", {
				"targets": {
					"node": "current"
				}
			}
		]
	],
	"plugins": [
		"@babel/plugin-proposal-object-rest-spread",
		"@babel/plugin-syntax-dynamic-import",
		"@babel/plugin-proposal-class-properties"
	]
}
`
