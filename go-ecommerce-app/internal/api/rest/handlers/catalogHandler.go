package handlers

import (
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/repository"
	"go-ecommerce-app/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CatelogHandler struct {
	svc service.CatelogService
}

func SetupCatelogRoutes(rh *rest.RestHandler) {
	app := rh.App

	// create an instance of user service and pass to handler
	svc := service.CatelogService{
		Repo:   repository.NewCatelogRepository(rh.DB),
		Auth:   rh.Auth,
		Config: rh.Config,
	}

	handler := CatelogHandler{svc: svc}

	//public end points
	// listing product and categories
	app.Get("/products", handler.GetProducts)
	app.Get("/product/:id", handler.GetProduct)
	app.Get("/catagories", handler.GetCategories)
	app.Get("/catagories/:id", handler.GetCategoryById)

	// private
	// manage product and categories
	selRoutes := app.Group("/seller", rh.Auth.AuthorizeSeller)

	//categories
	selRoutes.Post("/catagories", handler.CreateCategories)
	selRoutes.Patch("/catagories/:id", handler.EditCategories)
	selRoutes.Delete("/catagories/:id", handler.DeleteCategories)

	//products
	selRoutes.Post("/products", handler.CreateProduct)
	selRoutes.Get("/products", handler.GetProduct)
	selRoutes.Get("/products/:id", handler.GetProduct)

	selRoutes.Put("/products/:id", handler.EditProduct)
	selRoutes.Patch("/products/:id", handler.UpdateStock) // update stock
	selRoutes.Delete("/products/:id", handler.DeleteProduct)
}

func (h CatelogHandler) GetCategories(ctx *fiber.Ctx) error {

	categories, err := h.svc.GetCategories()

	if err != nil {
		return rest.ErrorMessage(ctx, 404, err)
	}

	return rest.SuccessResponse(ctx, "List of categories", categories)
}

func (h CatelogHandler) GetCategoryById(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	category, err := h.svc.GetCategory(id)
	if err != nil {
		return rest.ErrorMessage(ctx, 404, err)
	}
	return rest.SuccessResponse(ctx, "Category", category)
}

func (h CatelogHandler) CreateCategories(ctx *fiber.Ctx) error {

	req := dto.CreateCategoryRequest{}

	err := ctx.BodyParser(&req)
	if err != nil {
		return rest.BadRequestError(ctx, "create category request is not valid")
	}
	err = h.svc.CreateCategory(req)

	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, " category created successfully", nil)
}

func (h CatelogHandler) EditCategories(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	req := dto.CreateCategoryRequest{}

	err := ctx.BodyParser(&req)
	if err != nil {
		return rest.BadRequestError(ctx, "Edit category request is not valid")
	}
	updatedCat, err := h.svc.EditCategory(id, req)

	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, " category edited successfully", updatedCat)
}

func (h CatelogHandler) DeleteCategories(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	err := h.svc.DeleteCategory(id)

	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, " category deleted successfully ", nil)
}

func (h CatelogHandler) CreateProduct(ctx *fiber.Ctx) error {

	return rest.SuccessResponse(ctx, "creat product endpoint", nil)
}

func (h CatelogHandler) EditProduct(ctx *fiber.Ctx) error {

	return rest.SuccessResponse(ctx, "Edit product endpoint", nil)
}

func (h CatelogHandler) DeleteProduct(ctx *fiber.Ctx) error {

	return rest.SuccessResponse(ctx, "Delete product endpoint", nil)
}

func (h CatelogHandler) GetProduct(ctx *fiber.Ctx) error {

	return rest.SuccessResponse(ctx, "Get product endpoint", nil)
}

func (h CatelogHandler) GetProducts(ctx *fiber.Ctx) error {

	return rest.SuccessResponse(ctx, "Get product endpoint", nil)
}

func (h CatelogHandler) UpdateStock(ctx *fiber.Ctx) error {

	return rest.SuccessResponse(ctx, "update stock endpoint", nil)
}
