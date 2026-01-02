package handlers

import (
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/helper"
	"go-ecommerce-app/internal/repository"
	"go-ecommerce-app/internal/service"
	"go-ecommerce-app/pkg/payment"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TransactionHandler struct {
	svc           service.TransactionService
	paymentClient payment.PaymentClient
}

func initializeTransactionService(db *gorm.DB, auth helper.Auth) service.TransactionService {
	return service.TransactionService{
		Repo: repository.NewTansactionRepository(db),
		Auth: auth,
	}
}

func SetupTransactionRoutes(as *rest.RestHandler) {
	app := as.App
	svc := initializeTransactionService(as.DB, as.Auth)

	handler := TransactionHandler{
		svc:           svc,
		paymentClient: as.Pc,
	}
	secRoute := app.Group("/", as.Auth.Authorize)
	secRoute.Get("/payment", handler.MakePayment)

	sellerRoute := app.Group("/seller", as.Auth.AuthorizeSeller)
	sellerRoute.Get("/orders", handler.GetOrders)
	sellerRoute.Get("/orders/:id", handler.GetOrderDetails)
}

func (h *TransactionHandler) MakePayment(ctx *fiber.Ctx) error {

	sessionUrl, err := h.paymentClient.CreatePayment(100, 123, "456")
	if err != nil {
		return ctx.Status(400).JSON(err)
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message":     "login",
		"result":      sessionUrl,
		"payment_url": "success", // senthil need to check
	})
}

func (h *TransactionHandler) GetOrders(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON("success")
}

func (h *TransactionHandler) GetOrderDetails(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON("success")
}
