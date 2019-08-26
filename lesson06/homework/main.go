package main

import (
	"strconv"

	"github.com/kataras/iris"

	"fmt"

	"io"
	"os"
)

type Review struct {
	Reviewer      string
	ReviewContent string
}

type Product struct {
	ID          int
	Name        string
	Score       float64
	NewPrice    int
	OldPrice    int
	Description string
	Reviews     []Review
	Image       string
}

var products = []Product{
	{
		ID:          0,
		Name:        "Kem socola",
		Score:       3.5,
		NewPrice:    120000,
		OldPrice:    150000,
		Description: "Lorem ipsum, dolor sit amet consectetur adipisicing elit. Modi repellendus totam neque perspiciatis pariatur repudiandae dolore ad ratione, optio doloremque, molestiae quia quam, fuga obcaecati tenetur corrupti sequi dolorem repellat!",
		Reviews: []Review{
			{
				Reviewer:      "Anh toi",
				ReviewContent: "Sometimes will mid -.-",
			},
			{
				Reviewer:      "Su gia su that",
				ReviewContent: "Toi da lua anh em bao gio chua ?? :D ??",
			},
		},
		Image: "/resources/image/ic1.png",
	},
	{
		ID:          1,
		Name:        "Kem dau tay",
		Score:       4,
		NewPrice:    98000,
		OldPrice:    0,
		Description: "Lorem ipsum, dolor sit amet consectetur adipisicing elit. Modi repellendus totam neque perspiciatis pariatur repudiandae dolore ad ratione, optio doloremque, molestiae quia quam, fuga obcaecati tenetur corrupti sequi dolorem repellat!",
		Reviews: []Review{
			{
				Reviewer:      "Co tich",
				ReviewContent: "Meo peo ngok nghek, gout boy ^^",
			},
			{
				Reviewer:      "Misa",
				ReviewContent: "Gosu 7k",
			},
		},
		Image: "/resources/image/ic2.png",
	},
	{
		ID:          2,
		Name:        "Welcome The International 2019",
		Score:       5,
		NewPrice:    100000,
		OldPrice:    120000,
		Description: "Lorem ipsum, dolor sit amet consectetur adipisicing elit. Modi repellendus totam neque perspiciatis pariatur repudiandae dolore ad ratione, optio doloremque, molestiae quia quam, fuga obcaecati tenetur corrupti sequi dolorem repellat!",
		Reviews: []Review{
			{
				Reviewer:      "Huan Cao",
				ReviewContent: "Chu nguoi tu tu",
			},
			{
				Reviewer:      "Kuroky",
				ReviewContent: "Quoc truong Lilquit",
			},
		},
		Image: "/resources/image/ic3.png",
	},
	{
		ID:          3,
		Name:        "OG",
		Score:       5,
		NewPrice:    100000,
		OldPrice:    120000,
		Description: "Lorem ipsum, dolor sit amet consectetur adipisicing elit. Modi repellendus totam neque perspiciatis pariatur repudiandae dolore ad ratione, optio doloremque, molestiae quia quam, fuga obcaecati tenetur corrupti sequi dolorem repellat!",
		Reviews: []Review{
			{
				Reviewer:      "Topson",
				ReviewContent: "Hahahaha",
			},
			{
				Reviewer:      "Ana",
				ReviewContent: "The best ember spirit",
			},
		},
		Image: "/resources/image/ic4.png",
	},
}

func main() {
	app := iris.New()

	tmpl := iris.HTML("./view", ".html")
	tmpl.Layout("layout.html")
	tmpl.Reload(true)
	app.RegisterView(tmpl)

	// app.StaticWeb("/resources", "./resources")
	app.HandleDir("/resources", "./resources")
	app.Get("/san-pham", GetShop)
	app.Get("/san-pham/{id}", GetProductById)
	app.Get("/tao-san-pham", GetCreateProductPage)
	app.Post("/tao-san-pham", CreateProduct)
	app.Post("/upload", UploadImage)

	app.Run(iris.Addr(":8080"))
}

func GetShop(ctx iris.Context) {
	ctx.ViewData("products", products)
	ctx.View("shop.html")
}

func GetProductById(ctx iris.Context) {
	id, err := strconv.Atoi(ctx.Params().GetStringDefault("id", "1"))

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
	}

	fmt.Printf("%d \n", id)
	ctx.ViewData("Name", products[id].Name)
	ctx.ViewData("Image", products[id].Image)
	ctx.ViewData("Score", products[id].Score)
	ctx.ViewData("NewPrice", products[id].NewPrice)
	ctx.ViewData("OldPrice", products[id].OldPrice)
	ctx.ViewData("Description", products[id].Description)
	ctx.ViewData("Reviews", products[id].Reviews)

	ctx.View("product.html")
}

func GetCreateProductPage(ctx iris.Context) {
	ctx.ViewLayout(iris.NoLayout)
	ctx.View("create.html")
}

func CreateProduct(ctx iris.Context) {
	product := Product{}
	err := ctx.ReadForm(&product)
	if err != nil && !iris.IsErrPath(err) /* see: https://github.com/kataras/iris/issues/1157 */ {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
	}

	product.ID = len(products)
	products = append(products, product)

	ctx.Redirect("/san-pham")
}

func UploadImage(ctx iris.Context) {
	// Get the file from the request.
	file, info, err := ctx.FormFile("uploadfile")

	fmt.Println(info)

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
		return
	}

	defer file.Close()
	// info.Filename = "ic" + strconv.Itoa(len(products)) + ".png"
	fname := info.Filename

	fmt.Println(fname)

	// Create a file with the same name
	// assuming that you have a folder named 'uploads'
	out, err := os.OpenFile("uploads/"+fname,
		os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
		return
	}
	defer out.Close()

	io.Copy(out, file)
}
