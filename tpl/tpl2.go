package tpl

// Tpl2 ...
const Tpl2 = `
{
	"name": "core",
	"version": "0.0.0",
	"description": "",
	"main": "build/src/index.js",
	"types": "build/src/index.d.ts",
	"files": [
	  "build/src"
	],
	"license": "LGPL",
	"keywords": [],
	"scripts": {
	  "compile": "babel src -d dist",
	  "test": "echo \"Error: no test specified\" && exit 1",
	  "prepare": "npm run compile",
	  "pretest": "npm run compile",
	  "posttest": "npm run check"
	},
	"devDependencies": {
	  "@babel/cli": "^7.7.7",
	  "@babel/core": "^7.7.7",
	  "@babel/plugin-proposal-class-properties": "^7.7.4",
	  "@babel/plugin-syntax-dynamic-import": "^7.7.4",
	  "@babel/plugin-transform-runtime": "^7.7.6",
	  "@babel/preset-env": "^7.7.7",
	  "babel-loader": "^8.0.6",
	  "prettier": "^1.19.1",
	  "webpack": "^4.41.5",
	  "webpack-cli": "^3.3.10"
	},
	"dependencies": {
	  "arkfbp": "0.0.11",
	  "axios": "^0.19.0",
	  "body-parser": "^1.19.0",
	  "cookie": "^0.4.0",
	  "cookie-parser": "^1.4.4",
	  "debug": "^4.1.1",
	  "express": "^4.17.1",
	  "express-formidable": "^1.2.0",
	  "module-alias": "^2.2.2",
	  "yargs": "^15.0.2"
	}
  }
`
