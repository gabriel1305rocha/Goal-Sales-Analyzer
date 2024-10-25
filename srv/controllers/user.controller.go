package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gabriel1305rocha/Goal-Sales-Analyzer/models"
	"strconv"
)

type UserController struct {
	beego.Controller
}

// POST: Cria um novo usuário
func (c *UserController) CreateUser() {
	if c.Ctx.Request.Method == "POST" {
		name := c.GetString("name")
		email := c.GetString("email")
		password := c.GetString("password")
		ageStr := c.GetString("age")

		age, err := strconv.Atoi(ageStr)
		if err != nil {
			c.Data["error"] = "Invalid age"
			c.Ctx.Output.Status = 400 // Bad Request
			c.Data["json"] = map[string]string{"error": "Invalid age"}
			c.ServeJSON()
			return
		}

		user := models.User{Name: name, Email: email, Password: password, Age: age}
		if err := models.CreateUser(models.Db, &user); err != nil {
			c.Data["error"] = "Error on create user"
			c.Ctx.Output.Status = 500 // Internal Server Error
			c.Data["json"] = map[string]string{"error": "Error on create user"}
			c.ServeJSON()
			return
		}

		c.Data["json"] = map[string]string{"status": "User created successfully"}
		c.ServeJSON()
	} else {
		c.Ctx.Output.Status = 405 // Method Not Allowed
		c.Data["json"] = map[string]string{"error": "Method not allowed"}
		c.ServeJSON()
	}
}

// GET: Lista todos os usuários
func (c *UserController) ListUsers() {
	users, err := models.GetAllUsers(models.Db)
	if err != nil {
		c.Data["error"] = "Error on Get All users"
		c.Ctx.Output.Status = 500 // Internal Server Error
		c.Data["json"] = map[string]string{"error": "Error on Get All users"}
		c.ServeJSON()
		return
	}

	c.Data["json"] = users
	c.ServeJSON()
}
