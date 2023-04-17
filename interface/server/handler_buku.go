package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imron16/go-sewaBuku/entity"
	buku "github.com/imron16/go-sewaBuku/usecase/buku/response"
)

type bukuHandler struct {
	service buku.Service
}

func newBukuHanlder(service buku.Service) *bukuHandler {
	return &bukuHandler{service: service}
}

func (h *bukuHandler) GetAllBuku() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		result, err := h.service.GetAllBuku()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprint(err)})
		} else {
			c.JSON(http.StatusOK, gin.H{"Data": result})
		}

	}
}

func (h *bukuHandler) GetBukuById() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var requestBody entity.Buku
	}
}
