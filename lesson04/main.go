package main

import (
	"strconv"

	"fmt"

	"github.com/kataras/iris"
)

// Post struct ...
type Post struct {
	Thumbnail   string
	Title       string
	Description string
	Content     string
	Author      string
	ID          int
}

var posts = []Post{
	{
		ID:          0,
		Thumbnail:   "/static/img/go.jpeg",
		Title:       "Xây dựng web báng hàng bằng Golang",
		Description: "Iris MVC",
		Content:     "The International 2019",
		Author:      "Trịnh Minh Cường",
	},
	{
		ID:          1,
		Thumbnail:   "/static/img/uiux.png",
		Title:       "Thiết kế UI/UX",
		Description: "Don't make me think",
		Content:     "The International 2019",
		Author:      "Trịnh Minh Cường",
	},
	{
		ID:          2,
		Thumbnail:   "/static/img/frontend.jpg",
		Title:       "Web cơ bản HTML5, CSS3, Javascript",
		Description: "Freecodecamp is the best",
		Content:     "The International 2019",
		Author:      "Trịnh Minh Cường",
	},
	{
		ID:          3,
		Thumbnail:   "/static/img/pythonclass.jpg",
		Title:       "Python phân tích dữ liệu",
		Description: "Data Engineering",
		Content:     "The International 2019",
		Author:      "Trịnh Minh Cường",
	},
}

func main() {
	app := iris.New()
	// Đăng kí các file HTML trong thư mục view
	tmpl := iris.HTML("/home/hieuht/go/src/github.com/hieudevx/golang-web/lesson04/view", ".html")
	// Mỗi lần chỉnh sửa file HTML thì chỉ cần refresh lại trang
	tmpl.Reload(true)
	app.RegisterView(tmpl)
	app.Get("/bai-viet", hi)
	app.Get("/bai-viet/{id}", BlogDetailHanlder)
	app.Get("/", func(ctx iris.Context) {
		if err := ctx.View("form.html"); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.WriteString(err.Error())
		}
	})
	app.Post("/form_action", func(ctx iris.Context) {
		post := Post{}
		err := ctx.ReadForm(&post)
		if err != nil && !iris.IsErrPath(err) /* see: https://github.com/kataras/iris/issues/1157 */ {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.WriteString(err.Error())
		}

		post.ID = len(posts)
		posts = append(posts, post)
		// ctx.Writef("Post: %#v", post)
		ctx.Redirect("/bai-viet")
	})
	app.Get("/static/{folder}/{file}", serveFileHandler)
	app.Run(iris.Addr(":8080")) // defaults to that but you can change it.
}

func hi(ctx iris.Context) {
	ctx.ViewData("posts", posts)
	ctx.View("blog.html")
}

func BlogDetailHanlder(ctx iris.Context) {
	id, err := strconv.Atoi(ctx.Params().Get("id"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%d \n", id)
		ctx.ViewData("Title", posts[id].Title)
		ctx.ViewData("Description", posts[id].Description)
		ctx.ViewData("Content", posts[id].Content)
		ctx.ViewData("Author", posts[id].Author)
		ctx.ViewData("Thumbnail", posts[id].Thumbnail)
		ctx.View("blog-detail.html")
	}
}

func serveFileHandler(ctx iris.Context) {
	folder := ctx.Params().Get("folder")
	file := ctx.Params().Get("file")

	fmt.Printf(folder)
	fmt.Println()
	fmt.Printf(file)
	// đường dẫn trỏ đến file trong thư mục assets
	filePath := fmt.Sprintf("/home/hieuht/go/src/github.com/hieudevx/golang-web/lesson04/assets/%s/%s", folder, file)
	ctx.ServeFile(filePath, false)
}
