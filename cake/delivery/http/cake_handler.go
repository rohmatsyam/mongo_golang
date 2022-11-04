package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rohmatsyam/mongo_golang/domain"
	"net/http"
)

type CakeHandler struct {
	cakeUseCase domain.CakeUseCase
}

func NewCakeHandler(r *gin.Engine, cakeUC domain.CakeUseCase) {
	handler := &CakeHandler{
		cakeUseCase: cakeUC,
	}
	router := r.Group("/cakes")
	router.GET("/", handler.GetCakes)
	router.GET("/:id", handler.GetCakeWithId)
	router.POST("/", handler.CreateCake)
	router.PATCH("/:id", handler.UpdateCake)
	router.DELETE("/:id", handler.DeleteCake)
}

func (h CakeHandler) GetCakes(c *gin.Context) {
	res, err := h.cakeUseCase.GetCakesUC()
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"code":   http.StatusBadGateway,
			"status": err.Error(),
		})
		return
	}

	if len(res) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusOK,
			"status": "data masih kosong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"status": fmt.Sprintf("data found with length : %d", len(res)),
		"data":   res,
	})
}

func (h CakeHandler) GetCakeWithId(c *gin.Context) {
	res, err := h.cakeUseCase.GetCakeIdUC(c)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"code":   http.StatusBadGateway,
			"status": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusFound,
		"status": "data found",
		"data":   res,
	})
}

func (h CakeHandler) CreateCake(c *gin.Context) {
	res, err := h.cakeUseCase.CreateCakeUC(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"status": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusCreated,
		"status": "Successfully created a cake",
		"data":   res,
	})
}

func (h CakeHandler) UpdateCake(c *gin.Context) {
	res, err := h.cakeUseCase.UpdateCakeUC(c)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"code":   http.StatusBadGateway,
			"status": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusFound,
		"status": "data found",
		"data":   res,
	})
}
func (h CakeHandler) DeleteCake(c *gin.Context) {
	err := h.cakeUseCase.DeleteCakeUC(c)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"code":   http.StatusBadGateway,
			"status": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusFound,
		"status": "succesfully deleted data",
	})
}
