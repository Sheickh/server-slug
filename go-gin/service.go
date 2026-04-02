package gogin

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"

	"net/http"
	"server-slug/service"
)

func GetAll(ctx *gin.Context) {
	all, err := service.GetAlllinks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, all)
}

func GetAllSlug(ctx *gin.Context) {
	slug := ctx.Param("slug")
	all, err := service.GetLinkBySlug(slug)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, all)
}

func Post(ctx *gin.Context) {
	var req service.DataCreate

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if _, err := service.ValidateAndNormalizeURL(req.URL); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "неккоректный url",
		})
		return
	}

	if req.Slug == "" || req.URL == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("проверьте входные данные"),
		})
		return

	}
	if service.SlugAlert(req.Slug) == true {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Slug с таким именем уже существует",
		})
	}

	link, err := service.PostLink(req.URL, req.Slug)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, link)
}

func Patch(ctx *gin.Context) {
	id := ctx.Param("id")
	id1, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var req service.DataUpdate
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var data service.Data
	if _, err := service.ValidateAndNormalizeURL(*req.URL); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "неккоректный url",
		})
		return
	}

	if req.Slug != nil{
		req.Slug = &data.Slug
	}

	if req.URL != nil{
		req.Slug = &data.URL
	}

	link, err := service.PatchLinkById(id1, *req.Slug, *req.URL)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, link)
}

func Delete(ctx *gin.Context) {

    id := ctx.Query("id")
    slug := ctx.Query("slug")
    
    if id1, err := strconv.Atoi(id); err != nil{
        service.DeleteLinkByID(id1)
    } else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "неверный id",
		})
		return
	}
	if slug != "" {
        service.DeleteLinkBySlug(slug)
    } else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "неверный slug",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "данные удалены",
	})
	// Использование: DELETE /links?id=123 или DELETE /links?slug=my-slug
}

func Redirect(ctx *gin.Context) {
	slug := ctx.Param("slug")
	err, r := service.Redirect(slug)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "неверный запрос",
		})
		return
	}

	ctx.JSON(http.StatusFound, gin.H{
		"location": r, 
	})
}

/*
func DeleteID(ctx *gin.Context) {
	id := ctx.Param("id")
	id1, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = service.DeleteLinkByID(id1)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "deleted",
	})
}

func DeleteSlug(ctx *gin.Context) {
	slug := ctx.Param("slug")

	err := service.DeleteLinkBySlug(slug)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "deleted",
	})
	
}
*/