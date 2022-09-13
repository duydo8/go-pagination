package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"pagination/initializers"
	"pagination/models"
	"strconv"
)

type Pagination struct {
	NextPage     int
	PreviousPage int
	CurrentPage  int
	TotalElement int
	TotalPage    int
	User         []models.User
}

func FindByPageable(context *gin.Context) {
	perPage := context.DefaultQuery("size", "10")
	size, _ := strconv.Atoi(perPage)
	pageStr := context.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)

	offset := (page - 1) * size
	// find pageable
	var people []models.User
	initializers.DB.Limit(size).Offset(offset).Find(&people)
	// find count
	var totalRows int64
	initializers.DB.Model(&models.User{}).Count(&totalRows)
	fmt.Println(size)
	fmt.Println(totalRows / int64(size))
	totalPage := math.Ceil(float64(totalRows/int64(size) + 1))
	previousPage := page - 1
	if previousPage < 0 {
		previousPage = 0
	}
	context.JSON(200, gin.H{"data": Pagination{
		NextPage:     page + 1,
		PreviousPage: previousPage,
		CurrentPage:  page,
		TotalElement: int(totalRows),
		TotalPage:    int(math.Ceil(totalPage)),
		User:         people,
	}})
}
