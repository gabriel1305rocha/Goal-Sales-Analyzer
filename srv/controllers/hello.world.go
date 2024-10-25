package controllers

import (
	"github.com/astaxie/beego"
)

type HelloController struct {
	beego.Controller
}

func (c *HelloController) Get() {
	c.Ctx.Output.ContentType("text/html")
	c.Data["content"] = `
		<!DOCTYPE html>
		<html lang="pt-BR">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Hello World</title>
		</head>
		<body>
			<h1>Hello World</h1>
			<p>Welcome to Goal Sales Analyzer with Beego</p>
		</body>
		</html>
	`
	c.Ctx.WriteString(c.Data["content"].(string))
}
