package controller

import (
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/loveyandex/go-mongo-celestia-scan/model"
	"github.com/loveyandex/go-mongo-celestia-scan/service"
	"github.com/loveyandex/go-mongo-celestia-scan/util"

	jwtware "github.com/gofiber/jwt/v3"

	"github.com/golang-jwt/jwt/v4"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(app *fiber.App) *UserController {
	ccccc := &UserController{userService: service.NewUserService()}
	ccccc.initRouter(app)
	ccccc.secureInitRouter(app)
	return ccccc
}

func (a *UserController) initRouter(app *fiber.App) {
	r := app.Group("/user")

	r.Post("/login", a.login)

}
func (a *UserController) secureInitRouter(app *fiber.App) {
	r := app.Group("/user")

	r.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

}

func (uc *UserController) xxxxx(c *fiber.Ctx) error {

	return nil
}

func (uc *UserController) login(c *fiber.Ctx) error {

	username := c.FormValue("user")
	pass := c.FormValue("pass")

	//verify is username and pass is in db and corrected

	um, err := uc.userService.GetUserByPhone(username)

	if err != nil {
		//log.Fatal(err)
		c.Status(fiber.StatusNotFound)
		return c.JSON(map[string]interface{}{"status": 404,
			"code": "user not found"})
	}
	fmt.Printf("um: %v\n", um)

	b := util.ComparePasswords(um.PassWord, pass)

	if !b {
		return errors.New("password is inccorect")
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"name": "John Doe",
		"user": util.ClaimBody{
			UserId: um.ID.Hex(),
			Roles:  []string{string(model.ADMIN)},
			Phone:  um.Phone,
		},
		"exp": time.Now().Add(time.Hour * 24 * 30 * 12).Unix(),
	}

	// claims["roles"] = [1]int64{1}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})

}
