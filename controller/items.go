package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	model "github.com/huyanhvn/quick-rest-go/model"
	utils "github.com/huyanhvn/quick-rest-go/utils"
)

// @BasePath /api/v1

// GetAllItems returns an array of items in the database.
//
// GET /items
//
// @Summary Get all items.
// @Description Get all items.
// @Tags items
// @Produce json
// @Success      200  {array}   model.Item
// @Router /items [get]
func (c *Controller) GetAllItems(ctx *gin.Context) {
	items, err := model.GetAllItems()
	if err != nil {
		utils.NewHTTPError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, items)
}

// GetItemByID gets an item by primary key.
//
// GET /items/:id
//
// @Summary Get an item by ID.
// @Description Get an item by ID.
// @Tags items
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} model.Item
// @Router /items/{id} [get]
func (c *Controller) GetItemByID(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		utils.NewHTTPError(ctx, http.StatusBadRequest, err)
		return
	}

	item, err := model.GetItemByID(aid)
	if err != nil {
		utils.NewHTTPError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.IndentedJSON(http.StatusFound, item)
}

// CreateItem creates a new item in the database.
//
// POST /items
//
// @Summary Create an item.
// @Description Create an item.
// @Tags items
// @Accept json
// @Produce json
// @Param item body model.CreateItem true "Create an item"
// @Success 201
// @Router /items [post]
func (c *Controller) CreateItem(ctx *gin.Context) {
	var item model.CreateItem
	if err := ctx.ShouldBindJSON(&item); err != nil {
		utils.NewHTTPError(ctx, http.StatusBadRequest, err)
		return
	}

	if err := model.Create(item); err != nil {
		utils.NewHTTPError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.IndentedJSON(http.StatusCreated, "OK.")
}

// UpdateItem updates an existing item in the database.
//
// PATCH /items/:id
//
// @Summary Update an item.
// @Description Update an item.
// @Tags items
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Param item body model.UpdateItem true "Update an item"
// @Success 200
// @Router /items/{id} [patch]
func (c *Controller) UpdateItem(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		utils.NewHTTPError(ctx, http.StatusBadRequest, err)
		return
	}

	var item model.UpdateItem
	if err := ctx.ShouldBindJSON(&item); err != nil {
		utils.NewHTTPError(ctx, http.StatusBadRequest, err)
		return
	}

	if err := model.Update(aid, item); err != nil {
		utils.NewHTTPError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.IndentedJSON(http.StatusOK, "Updated.")
}

// DeleteItem deletes an item by its primary key.
//
// DELETE /items/:id
//
// @Summary Delete an item.
// @Description Delete an item.
// @Tags items
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Success 200
// @Router /items/{id} [delete]
func (c *Controller) DeleteItem(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		utils.NewHTTPError(ctx, http.StatusBadRequest, err)
		return
	}

	if err := model.Delete(aid); err != nil {
		utils.NewHTTPError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"data": true})
}
