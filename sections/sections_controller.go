package sections

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vanntrong/asana-clone-be/configs"
	"github.com/vanntrong/asana-clone-be/middleware"
	"github.com/vanntrong/asana-clone-be/utils"
)

type SectionsController struct {
	SectionsService ISectionsService
}

func registerRoutes(router *gin.RouterGroup, ctrl *SectionsController) {
	v1 := router.Group("/sections")
	v1.GET("/", middleware.AuthMiddleware, ctrl.GetList)
	v1.GET("/:id", middleware.AuthMiddleware, ctrl.GetById)
	v1.POST("/", middleware.AuthMiddleware, ctrl.CreateSection)
	v1.PUT("/:id", middleware.AuthMiddleware, ctrl.UpdateSection)
}

func NewSectionsController(app *gin.RouterGroup, sectionsService ISectionsService) {
	ctrl := &SectionsController{
		sectionsService,
	}
	registerRoutes(app, ctrl)
}

func (ctrl *SectionsController) GetList(c *gin.Context) {
	userId := c.Request.Header.Get(configs.HeaderUserId)

	var query GetListSectionValidation

	isValid := utils.ValidationQuery(c, &query)

	if !isValid {
		return
	}

	sections, err := ctrl.SectionsService.GetList(userId, query.ProjectId)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	utils.GenerateResponse(c, sections, http.StatusOK)
}

func (ctrl *SectionsController) CreateSection(c *gin.Context) {
	userId := c.Request.Header.Get(configs.HeaderUserId)

	var body CreateSectionValidation

	isValid := utils.Validation(c, &body)

	if !isValid {
		return
	}

	section, err := ctrl.SectionsService.CreateSection(userId, body)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	utils.GenerateResponse(c, section, http.StatusOK)
}

func (ctrl *SectionsController) UpdateSection(c *gin.Context) {
	userId := c.Request.Header.Get(configs.HeaderUserId)

	var body UpdateSectionValidation

	isValid := utils.Validation(c, &body)

	if !isValid {
		return
	}

	sectionId := c.Param("id")

	section, err := ctrl.SectionsService.UpdateSection(userId, sectionId, body)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	utils.GenerateResponse(c, section, http.StatusOK)
}

func (ctrl *SectionsController) GetById(c *gin.Context) {
	userId := c.Request.Header.Get(configs.HeaderUserId)

	sectionId := c.Param("id")

	section, err := ctrl.SectionsService.GetById(userId, sectionId)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	utils.GenerateResponse(c, section, http.StatusOK)
}
