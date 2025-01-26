package transaksi

import (
	"fmt"
	"xyz/entity"
	"xyz/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	repo := NewTransaksiRepository(db)
	service := NewTransaksiService(repo)
	handler := NewTransaksiController(service)

	routes := app.Group("/transaksi")
	routes.Get("/master-data/all", handler.MasterDataAll)
	routes.Get("/detail/:id", handler.DetailTransaksi)
	routes.Get("/detail/free/:id", handler.DetailTransaksi)
	routes.Post("/create-transaksi", middleware.DeserializeUser, handler.CreateTransaksi)
	routes.Post("/update-transaksi", handler.UpdateTransaksi)
}

type TransaksiController interface {
	CreateTransaksi(ctx *fiber.Ctx) error
	UpdateTransaksi(ctx *fiber.Ctx) error
	MasterDataAll(ctx *fiber.Ctx) error
	DetailTransaksi(ctx *fiber.Ctx) error
}

type transaksiController struct {
	transaksiService TransaksiService
}

func NewTransaksiController(aS TransaksiService) TransaksiController {
	return &transaksiController{
		transaksiService: aS,
	}
}

func (tr *transaksiController) DetailTransaksi(ctx *fiber.Ctx) error {
	nik := ctx.Params("nik")
	return ctx.JSON(tr.transaksiService.DetailTransaksi(nik))
}

func (tr *transaksiController) MasterDataAll(ctx *fiber.Ctx) error {
	return ctx.JSON(map[string]interface{}{"data": tr.transaksiService.MasterDataAll(), "status": "success"})
}

func (tr *transaksiController) UpdateTransaksi(ctx *fiber.Ctx) error {
	var body entity.Transaksi
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	body.UpdatedBy = details.Name
	err = tr.transaksiService.Update(body)
	if err != nil {
		return ctx.JSON(map[string]interface{}{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil Update"})
}
func (tr *transaksiController) CreateTransaksi(ctx *fiber.Ctx) error {
	var body TransaksiRequest
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	body.Nik = details.Nik
	err = tr.transaksiService.CreateTransaksi(body)
	if err != nil {
		return ctx.JSON(map[string]string{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil create"})
}
