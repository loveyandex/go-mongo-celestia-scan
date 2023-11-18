package controller

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/loveyandex/go-mongo-celestia-scan/service"
	"github.com/loveyandex/go-mongo-celestia-scan/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CmnCtl[T any] struct {
	Path   string
	CmnSrv *service.CmnSrv[T]
}

func (a *CmnCtl[T]) Setup(app *fiber.App) {
	a.initRouter(app)
	a.secureInitRouter(app)
}
func (a *CmnCtl[T]) initRouter(app *fiber.App) {
	r := app.Group(a.Path)
	r.Get("/", a.getAll)

}
func (a *CmnCtl[T]) secureInitRouter(app *fiber.App) {
	r := app.Group(a.Path)

	r.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	r.Post("/", a.make)
	// r.Get("/:id", a.Get)
}

func (cmnCtl *CmnCtl[T]) getAll(c *fiber.Ctx) error {
	var xxxxxxx T
	err := c.BodyParser(&xxxxxxx)
	if err != nil {
		return err
	}

	i, err := cmnCtl.CmnSrv.Create(&xxxxxxx)
	if err != nil {
		return err
	}
	id := i.(primitive.ObjectID).Hex()
	rrrrrrrrrr, err := cmnCtl.CmnSrv.Get(id)

	return util.Jackson(c, rrrrrrrrrr, err)
}

func (cmnCtl *CmnCtl[T]) make(c *fiber.Ctx) error {
	var xxxxxxx T
	err := c.BodyParser(&xxxxxxx)
	if err != nil {
		return err
	}

	i, err := cmnCtl.CmnSrv.Create(&xxxxxxx)
	if err != nil {
		return err
	}
	id := i.(primitive.ObjectID).Hex()
	rrrrrrrrrr, err := cmnCtl.CmnSrv.Get(id)

	return util.Jackson(c, rrrrrrrrrr, err)
}
