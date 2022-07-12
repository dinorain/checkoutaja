package http

import (
	"errors"
	"net/http"
	"strings"

	"github.com/go-playground/validator"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/dinorain/checkoutaja/config"
	"github.com/dinorain/checkoutaja/internal/middlewares"
	"github.com/dinorain/checkoutaja/internal/models"
	"github.com/dinorain/checkoutaja/internal/product"
	"github.com/dinorain/checkoutaja/internal/product/delivery/http/dto"
	"github.com/dinorain/checkoutaja/internal/session"
	"github.com/dinorain/checkoutaja/pkg/constants"
	httpErrors "github.com/dinorain/checkoutaja/pkg/http_errors"
	"github.com/dinorain/checkoutaja/pkg/logger"
	"github.com/dinorain/checkoutaja/pkg/utils"
)

type productHandlersHTTP struct {
	group     *echo.Group
	logger    logger.Logger
	cfg       *config.Config
	mw        middlewares.MiddlewareManager
	v         *validator.Validate
	productUC product.ProductUseCase
	sessUC    session.SessUseCase
}

func NewProductHandlersHTTP(
	group *echo.Group,
	logger logger.Logger,
	cfg *config.Config,
	mw middlewares.MiddlewareManager,
	v *validator.Validate,
	productUC product.ProductUseCase,
	sessUC session.SessUseCase,
) *productHandlersHTTP {
	return &productHandlersHTTP{group: group, logger: logger, cfg: cfg, mw: mw, v: v, productUC: productUC, sessUC: sessUC}
}

// Create
// @Tags Products
// @Summary To create product
// @Description Seller create product
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param payload body dto.ProductCreateRequestDto true "Payload"
// @Success 200 {object} dto.ProductCreateResponseDto
// @Router /product [post]
func (h *productHandlersHTTP) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		createDto := &dto.ProductCreateRequestDto{}
		if err := c.Bind(createDto); err != nil {
			h.logger.WarnMsg("bind", err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		if err := h.v.StructCtx(ctx, createDto); err != nil {
			h.logger.WarnMsg("validate", err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		sessID, err := h.getSessionIDFromCtx(c)
		if err != nil {
			h.logger.Errorf("getSessionIDFromCtx: %v", err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		session, err := h.sessUC.GetSessionByID(ctx, sessID)
		if err != nil {
			h.logger.Errorf("sessUC.GetSessionByID: %v", err)
			if errors.Is(err, redis.Nil) {
				httpErrors.NewUnauthorizedError(c, err, h.cfg.Http.DebugErrorsResponse)
			}
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		product, err := h.registerReqToProductModel(createDto, session.UserID)
		if err != nil {
			h.logger.Errorf("registerReqToProductModel: %v", err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		createdProduct, err := h.productUC.Create(ctx, product)
		if err != nil {
			h.logger.Errorf("productUC.Create: %v", err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		return c.JSON(http.StatusCreated, dto.ProductCreateResponseDto{ProductID: createdProduct.ProductID})
	}
}

// FindAll
// @Tags Products
// @Summary Find all products
// @Description Find all products
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param size query string false "pagination size"
// @Param page query string false "pagination page"
// @Param seller_id query string false "seller id"
// @Success 200 {object} dto.ProductFindResponseDto
// @Router /product [get]
func (h *productHandlersHTTP) FindAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		pq := utils.NewPaginationFromQueryParams(c.QueryParam(constants.Size), c.QueryParam(constants.Page))

		var products []models.Product
		if c.QueryParam("seller_id") != "" {
			if res, err := h.productUC.FindAllBySellerId(ctx, c.QueryParam("seller_id"), pq); err != nil {
				h.logger.Errorf("productUC.FindAllBySellerId: %v", err)
				return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
			} else {
				products = res
			}
		} else {
			if res, err := h.productUC.FindAll(ctx, pq); err != nil {
				h.logger.Errorf("productUC.FindAll: %v", err)
				return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
			} else {
				products = res
			}
		}

		return c.JSON(http.StatusOK, dto.ProductFindResponseDto{
			Data: products,
			Meta: utils.PaginationMetaDto{
				Limit:  pq.GetLimit(),
				Offset: pq.GetOffset(),
				Page:   pq.GetPage(),
			},
		})
	}
}

// FindByID
// @Tags Products
// @Summary Find product
// @Description Find existing product by id
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} dto.ProductResponseDto
// @Router /product/{id} [get]
func (h *productHandlersHTTP) FindByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		productUUID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			h.logger.WarnMsg("uuid.FromString", err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		product, err := h.productUC.CachedFindById(ctx, productUUID)
		if err != nil {
			h.logger.Errorf("productUC.CachedFindById: %v", err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		return c.JSON(http.StatusOK, dto.ProductResponseFromModel(product))
	}
}

// UpdateByID
// @Tags Products
// @Summary Update product
// @Description Update existing product
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Product ID"
// @Success 200 {object} dto.ProductResponseDto
// @Router /product/{id} [put]
func (h *productHandlersHTTP) UpdateByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		productUUID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			h.logger.WarnMsg("uuid.FromString", err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		updateDto := &dto.ProductUpdateRequestDto{}
		if err := c.Bind(updateDto); err != nil {
			h.logger.WarnMsg("bind", err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		if err := h.v.StructCtx(ctx, updateDto); err != nil {
			h.logger.WarnMsg("validate", err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		product, err := h.productUC.FindById(ctx, productUUID)
		if err != nil {
			h.logger.Errorf("productUC.FindById: %v", err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		product, err = h.updateReqToProductModel(product, updateDto)
		if err != nil {
			h.logger.Errorf("updateReqToProductModel: %v", err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		product, err = h.productUC.UpdateById(ctx, product)
		if err != nil {
			h.logger.Errorf("productUC.UpdateById: %v", err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		return c.JSON(http.StatusOK, dto.ProductResponseFromModel(product))
	}
}

// DeleteByID
// @Tags Products
// @Summary Delete product
// @Description Delete existing product
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} nil
// @Param id path string true "Product ID"
// @Router /product/{id} [delete]
func (h *productHandlersHTTP) DeleteByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		productUUID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			h.logger.WarnMsg("uuid.FromString", err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		if err := h.productUC.DeleteById(ctx, productUUID); err != nil {
			h.logger.Errorf("productUC.DeleteById: %v", err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		return c.JSON(http.StatusOK, nil)
	}
}

func (h *productHandlersHTTP) getSessionIDFromCtx(c echo.Context) (string, error) {
	user, ok := c.Get("user").(*jwt.Token)
	if !ok {
		h.logger.Warnf("jwt.Token: %+v", c.Get("user"))
		return "", errors.New("invalid token header")
	}

	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		h.logger.Warnf("jwt.MapClaims: %+v", c.Get("user"))
		return "", errors.New("invalid token header")
	}

	sessionID, ok := claims["session_id"].(string)
	if !ok {
		h.logger.Warnf("session_id: %+v", claims)
		return "", errors.New("invalid token header")
	}

	return sessionID, nil
}

func (h *productHandlersHTTP) registerReqToProductModel(r *dto.ProductCreateRequestDto, sellerID uuid.UUID) (*models.Product, error) {
	productCandidate := &models.Product{
		Name:        r.Name,
		Description: r.Description,
		Price:       r.Price,
		SellerID:    sellerID,
	}

	if err := productCandidate.PrepareCreate(); err != nil {
		return nil, err
	}

	return productCandidate, nil
}

func (h *productHandlersHTTP) updateReqToProductModel(updateCandidate *models.Product, r *dto.ProductUpdateRequestDto) (*models.Product, error) {
	if r.Name != nil {
		updateCandidate.Name = strings.TrimSpace(*r.Name)
	}
	if r.Description != nil {
		updateCandidate.Description = strings.TrimSpace(*r.Description)
	}

	return updateCandidate, nil
}
