package controllers

import (
	"assignment2/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetOrders(c *gin.Context) {
	var (
		orders []models.Order
		result gin.H
	)

	idb.DB.Preload("Items").Find(&orders)
	if len(orders) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": orders,
			"count":  len(orders),
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) CreateOrder(c *gin.Context) {
	var (
		order  models.Order
		items  []models.Item
		result gin.H
	)

	err := c.BindJSON(&order)
	if err != nil {
		result = gin.H{
			"error": err.Error(),
			"code":  "400",
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, result)
		return
	}
	// check if item is empty
	items = order.Items
	if len(items) == 0 {
		result = gin.H{
			"error": "Item can't be empty",
			"code":  "400",
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, result)
		return
	}

	idb.DB.Create(&order)
	result = gin.H{
		"result": order,
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdateOrder(c *gin.Context) {
	var (
		order    models.Order
		newOrder models.Order
		result   gin.H
	)

	id := c.Param("orderId")
	// check if data exist
	err := idb.DB.Preload("Items").First(&order, "id=?", id).Error
	if err != nil {
		result = gin.H{
			"error": "Order Id not found",
			"code":  "400",
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, result)
		return
	}

	err = c.BindJSON(&newOrder)
	if err != nil {
		result = gin.H{
			"error": err.Error(),
			"code":  "400",
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, result)
		return
	}

	err = idb.DB.Model(&order).Update(newOrder).Error
	if err != nil {
		result = gin.H{
			"error": err.Error(),
			"code":  "400",
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, result)
		return
	}

	c.JSON(http.StatusOK, newOrder)
}

func (idb *InDB) DeleteOrder(c *gin.Context) {
	var (
		order  models.Order
		result gin.H
	)
	id := c.Param("orderId")
	err := idb.DB.Preload("Items").First(&order, "id=?", id).Error
	if err != nil {
		result = gin.H{
			"error": "Order Id not found",
			"code":  "400",
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, result)
		return
	}

	err = idb.DB.Delete(&order).Error
	// delete the items associated with the order
	for _, item := range order.Items {
		err = idb.DB.Delete(&item).Error
	}
	if err != nil {
		result = gin.H{
			"error": err.Error(),
			"code":  "400",
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, result)
		return
	} else {
		result = gin.H{
			"result": "Data deleted successfully",
		}
	}

	c.JSON(http.StatusOK, result)
}
