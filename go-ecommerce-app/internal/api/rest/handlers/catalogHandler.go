package handlers

import (
	"fmt"
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/domain"
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
	app.Get("/categories", handler.GetCategories)
	app.Get("/categories/:id", handler.GetCategoryById)

	// private
	// manage product and categories
	selRoutes := app.Group("/seller", rh.Auth.AuthorizeSeller)

	//categories
	selRoutes.Post("/categories", handler.CreateCategories)
	selRoutes.Patch("/categories/:id", handler.EditCategories)
	selRoutes.Delete("/categories/:id", handler.DeleteCategories)

	//products
	selRoutes.Post("/products", handler.CreateProduct)
	selRoutes.Get("/products", handler.GetProducts)
	selRoutes.Get("/products/:id", handler.GetProduct)

	selRoutes.Put("/products/:id", handler.EditProduct)
	selRoutes.Patch("/products/:id", handler.UpdateStock) // update stock
	selRoutes.Delete("/products/:id", handler.DeleteProduct)
}

func (h CatelogHandler) GetCategories(ctx *fiber.Ctx) error {

	fmt.Println(" CATE : Request received")

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

	req := dto.CreateProductRequest{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return rest.BadRequestError(ctx, "Invalid product request")
	}

	user := h.svc.Auth.GetCurrentUser(ctx)
	err = h.svc.CreateProduct(req, user)
	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "Product created successfully", nil)
}

func (h CatelogHandler) EditProduct(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))
	req := dto.CreateProductRequest{}

	err := ctx.BodyParser(&req)

	if err != nil {
		return rest.BadRequestError(ctx, "edit product request is not valid")
	}
	user := h.svc.Auth.GetCurrentUser(ctx)
	product, err := h.svc.EditProduct(id, req, user)

	if err != nil {
		return rest.InternalError(ctx, err)
	}
	return rest.SuccessResponse(ctx, "Edit product ", product)
}

func (h CatelogHandler) DeleteProduct(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	user := h.svc.Auth.GetCurrentUser(ctx)

	err := h.svc.DeleteProduct(id, user)

	return rest.SuccessResponse(ctx, "Delete product endpoint", err)
}

func (h CatelogHandler) GetProduct(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))
	fmt.Println("INput ", id)
	product, err := h.svc.GetProductById(id)
	if err != nil {
		return rest.BadRequestError(ctx, "product not found")
	}

	return rest.SuccessResponse(ctx, "Get product endpoint", product)
}

func (h CatelogHandler) GetProducts(ctx *fiber.Ctx) error {

	products, err := h.svc.GetProducts()
	if err != nil {
		return rest.ErrorMessage(ctx, 404, err)
	}

	return rest.SuccessResponse(ctx, "List of products", products)
}

func (h CatelogHandler) UpdateStock(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	req := dto.UpdateStockRequest{}
	err := ctx.BodyParser(&req)

	if err != nil {
		return rest.BadRequestError(ctx, "Invalid update stock input")
	}
	user := h.svc.Auth.GetCurrentUser(ctx)
	product := domain.Product{
		ID:     uint(id),
		Stock:  uint(req.Stock),
		UserId: int(user.ID),
	}

	updateProduct, err := h.svc.UpdateProductStock(product)
	if err != nil {
		return rest.ErrorMessage(ctx, 404, err)
	}

	return rest.SuccessResponse(ctx, "update stock endpoint", updateProduct)
}
