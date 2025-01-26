package konsumen

import (
	"fmt"
	"xyz/entity"
	"xyz/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	repo := NewKonsumenRepository(db)
	service := NewKonsumenService(repo)
	handler := NewKonsumenController(service)

	routes := app.Group("/konsumen")
	routes.Get("/master-data/all", handler.MasterDataAll)
	routes.Get("/detail/:nik", handler.DetailKonsumen)
	routes.Post("/create-konsumen", handler.CreateKonsumen)
	routes.Post("/update-konsumen", handler.UpdateKonsumen)
	routes.Post("/upload-ktp",middleware.DeserializeUser, handler.UploadFotoKtp)
	routes.Post("/upload-selfie",middleware.DeserializeUser, handler.UploadSelfie)
}

type KonsumenController interface {
	CreateKonsumen(ctx *fiber.Ctx) error
	UpdateKonsumen(ctx *fiber.Ctx) error
	MasterDataAll(ctx *fiber.Ctx) error
	DetailKonsumen(ctx *fiber.Ctx) error
	UploadFotoKtp(ctx *fiber.Ctx) error
	UploadSelfie(ctx *fiber.Ctx) error
}

type konsumenController struct {
	konsumenService KonsumenService
}

func NewKonsumenController(aS KonsumenService) KonsumenController {
	return &konsumenController{
		konsumenService: aS,
	}
}

func (c *konsumenController) UploadFotoKtp(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		fmt.Println("error file ", err)
		return err
	}
	destination := fmt.Sprintf("./uploads/%s", file.Filename)
	if err := ctx.SaveFile(file, destination); err != nil {
		return err
	}
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	c.konsumenService.updateKtp(entity.Konsumen{Nik: details.Nik, FotoKtp: destination})
	return ctx.JSON(map[string]interface{}{"data": "/uploads/" + file.Filename, "message": "Berhasil"})
}

func (c *konsumenController) UploadSelfie(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		fmt.Println("error file ", err)
		return err
	}
	destination := fmt.Sprintf("./uploads/%s", file.Filename)
	if err := ctx.SaveFile(file, destination); err != nil {
		return err
	}
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	c.konsumenService.updateSelfie(entity.Konsumen{Nik: details.Nik, FotoSelfie: destination})
	return ctx.JSON(map[string]interface{}{"data": "/uploads/" + file.Filename, "message": "Berhasil"})
}

func (c *konsumenController) DetailKonsumen(ctx *fiber.Ctx) error {
	nik := ctx.Params("nik")
	return ctx.JSON(c.konsumenService.DetailKonsumen(nik))
}

func (c *konsumenController) MasterDataAll(ctx *fiber.Ctx) error {
	return ctx.JSON(map[string]interface{}{"data": c.konsumenService.MasterDataAll(), "status": "success"})
}

func (c *konsumenController) UpdateKonsumen(ctx *fiber.Ctx) error {
	var body entity.Konsumen
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	body.UpdatedBy = details.Name
	err = c.konsumenService.Update(body)
	if err != nil {
		return ctx.JSON(map[string]interface{}{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil Update"})
}
func (c *konsumenController) CreateKonsumen(ctx *fiber.Ctx) error {
	var body entity.Konsumen
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	body.CreatedBy = details.Name
	fmt.Println("ini body yaa ", body)
	err = c.konsumenService.CreateKonsumen(body)
	if err != nil {
		return ctx.JSON(map[string]string{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil create"})
}
