package controllers

import (
	"errors"
	"net/http"

	"github.com/alpakih/simpel-api/helper"
	"github.com/alpakih/simpel-api/models"
	"github.com/alpakih/simpel-api/request"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrderHandler struct {
	db *gorm.DB
}

func OrderController(db *gorm.DB) OrderHandler {
	return OrderHandler{db: db}
}

func (r *OrderHandler) GetAll(c *gin.Context) {
	var data []models.Order
	if err := r.db.Preload("OrderItem").Find(&data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": http.StatusText(http.StatusInternalServerError),
			"error":   err.Error(),
		})

		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": http.StatusText(http.StatusOK),
		"data":    data,
	})

}

func (r *OrderHandler) GetByID(c *gin.Context) {
	var data models.Order

	id := c.Param("id")
	if err := r.db.Preload("OrderItem").First(&data, "id =?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": http.StatusText(http.StatusNotFound),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": http.StatusText(http.StatusInternalServerError),
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": http.StatusText(http.StatusOK),
		"data":    data,
	})

}

func (r *OrderHandler) Store(c *gin.Context) {
	var req request.CreateOrder

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": http.StatusText(http.StatusInternalServerError),
			"error":   err.Error(),
		})
		return
	}

	dateParse, err := helper.ParseDateString(req.OrderAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": http.StatusText(http.StatusInternalServerError),
			"error":   err.Error(),
		})
		return
	}

	items := []models.OrderItem{}

	for _, v := range req.Items {
		var item models.OrderItem
		item.ItemCode = v.ItemCode
		item.Description = v.Description
		item.Quantity = v.Quantity
		items = append(items, item)
	}

	entity := models.Order{
		CustomerName: req.CustomerName,
		OrderAt:      dateParse,
		OrderItem:    items,
	}

	if err := r.db.Debug().Create(&entity).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": http.StatusText(http.StatusInternalServerError),
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": http.StatusText(http.StatusOK),
		"data":    entity,
	})

}

func (r *OrderHandler) Update(c *gin.Context) {
	var req request.UpdateOrder
	var data models.Order

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": http.StatusText(http.StatusInternalServerError),
			"error":   err.Error(),
		})
		return
	}

	db := r.db.Debug()

	if err := db.First(&data, "id =?", req.OrderID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": http.StatusText(http.StatusNotFound),
				"error":   err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": http.StatusText(http.StatusInternalServerError),
			"error":   err.Error(),
		})
		return
	}

	dateParse, err := helper.ParseDateString(req.OrderAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": http.StatusText(http.StatusInternalServerError),
			"error":   err.Error(),
		})
		return
	}

	data.CustomerName = req.CustomerName
	data.OrderAt = dateParse

	if err := db.Omit("OrderItem").Save(&data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": http.StatusText(http.StatusInternalServerError),
			"error":   err.Error(),
		})
		return
	}

	// Update columns to new value on `id` conflict
	items := []models.OrderItem{}

	for _, v := range req.Items {
		item := models.OrderItem{}
		item.ID = v.ID
		item.ItemCode = v.ItemCode
		item.Description = v.Description
		item.Quantity = v.Quantity
		item.OrderRef = data.ID
		items = append(items, item)
	}

	if err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"item_code", "description", "quantity"}),
	}).Create(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": http.StatusText(http.StatusInternalServerError),
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": http.StatusText(http.StatusOK),
		"data":    data,
	})

}

func (r *OrderHandler) Destroy(c *gin.Context) {
	var data models.Order

	id := c.Param("id")
	db := r.db.Debug()
	if err := db.Preload("OrderItem").First(&data, "id =?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": http.StatusText(http.StatusNotFound),
				"error":   err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": http.StatusText(http.StatusInternalServerError),
			"error":   err.Error(),
		})
		return
	}
	if err := db.Delete(&data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": http.StatusText(http.StatusInternalServerError),
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "delete data success",
	})

}
