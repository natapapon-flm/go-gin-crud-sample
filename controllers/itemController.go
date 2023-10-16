package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/natapapon-flm/go-gin/database"
	"github.com/natapapon-flm/go-gin/models"
)

// GET '/'
func GetAllItems(c *gin.Context) {
	items := []models.Item{}

	database.DB.Model(&models.Item{}).Find(&items)

	c.JSON(http.StatusOK, gin.H{ "data": items, "result": true, "status": 200 })
}

// GET '/:Id'
func GetItemsById(c *gin.Context) {
	id := c.Param("Id")

	var item models.Item
  if err := database.DB.Model(&models.Item{}).Where("id = ?", id).First(&item).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!", "result": false, "status": 404})
    return
  }

	if item.Id != "" {
		c.JSON((http.StatusOK), gin.H{"data": item, "status": 200, "result": true})
	}
}

// POST '/'
func Create(c *gin.Context) {
	var itemInput models.ItemRequst

	if err := c.ShouldBindJSON(&itemInput); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

	item := models.Item{Id: uuid.NewString(), ItemName: itemInput.ItemName, ItemPrice: itemInput.ItemPrice, Amount: itemInput.Amount}

	result := database.DB.Create(&item);

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	if result.RowsAffected > 0 {
		var newItem models.Item
		
		database.DB.Model(&models.Item{}).First(&newItem, "id = ?", item.Id)

		c.JSON(http.StatusAccepted, gin.H{"data": newItem, "result": true, "status": 201})
	}
}

// PATCH '/:Id'
func Update(c *gin.Context) {
	var id = c.Param("Id");
	
	var itemRequest models.ItemRequst

	var item models.Item
  if err := database.DB.Model(&models.Item{}).Where("id = ?", id).First(&item).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!", "result": false, "status": 404})
    return
  }

	if err := c.ShouldBindJSON(&itemRequest); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

	result := database.DB.Model(&item).Updates(itemRequest)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result, "result": false, "status": 500})
	}

	var itemUpdated models.Item
		database.DB.Model(&models.Item{}).Where("id = ?", id).First(&itemUpdated)
	
	if result.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{"data": itemUpdated, "result": true, "status": 200})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": itemUpdated, "result": true, "status": 200, "messsage": "No data updated"})
	}
}

// DELETE '/:Id'
func Delete(c *gin.Context) {
	var id = c.Param("Id");
	var item models.Item

	if err := database.DB.Model(&models.Item{}).Where("id = ?", id).First(&item).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!", "result": false, "status": 404})
    return
  }

  result := database.DB.Delete(&item)

	if result.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{"message": `Success fully remove item ID :: ${id}`, "result": true, "status": 200})
	}
}