package main

import (
	"github.com/kataras/iris"
)

// User struct ...
type User struct {
	ID       int
	Name     string
	Password string
}

var users = []User{
	{
		ID:       0,
		Name:     "hieuht",
		Password: "12345678",
	},
	{
		ID:       1,
		Name:     "cotich",
		Password: "meobeongocnghech",
	},
	{
		ID:       2,
		Name:     "bachrau",
		Password: "tatcachilaculua",
	},
	{
		ID:       3,
		Name:     "dendi",
		Password: "anhtoi",
	},
}

func main() {
	app := iris.New()
	// Đăng kí các file HTML trong thư mục view
	tmpl := iris.HTML("./views", ".html")
	// Mỗi lần chỉnh sửa file HTML thì chỉ cần refresh lại trang
	tmpl.Reload(true)
	app.RegisterView(tmpl)
	app.Get("/home", HomeView)
	app.Get("/register", RegisterView)
	app.Get("/sign-in", SigninView)
	app.Post("/register", RegisterHandler)
	app.Post("sign-in", SigninHandler)

	app.Run(iris.Addr(":8080")) // defaults to that but you can change it.

}

func HomeView(ctx iris.Context) {
	ctx.ViewData("users", users)
	ctx.View("home.html")
}

func RegisterView(ctx iris.Context) {
	ctx.View("register.html")
}

func SigninView(ctx iris.Context) {
	ctx.View("sign-in.html")
}

func RegisterHandler(ctx iris.Context) {
	user := User{}
	err := ctx.ReadForm(&user)
	if err != nil && !iris.IsErrPath(err) /* see: https://github.com/kataras/iris/issues/1157 */ {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
	}
	user.ID = len(users)
	users = append(users, user)

	// ctx.Writef("User: %#v", user)
	ctx.Redirect("/home")
}

func SigninHandler(ctx iris.Context) {
	user := User{}
	err := ctx.ReadForm(&user)
	if err != nil && !iris.IsErrPath(err) /* see: https://github.com/kataras/iris/issues/1157 */ {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
	}

	isRight := false
	var userID int
	for _, v := range users {
		if user.Name == v.Name && user.Password == v.Password {
			isRight = true
			userID = v.ID
			break
		}
	}
	if isRight {
		ctx.Writef("Sign in successfully \nUser id: %d", userID)
	} else {
		ctx.Writef("Failed to sign in")
	}
}
