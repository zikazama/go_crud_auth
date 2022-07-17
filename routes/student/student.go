package routes

import (
	"strconv"

	"go_crud_auth/config"
	"go_crud_auth/models"

	"github.com/gin-gonic/gin"
)

func ReadStudent(c *gin.Context) {
	id := c.Param("studentId")
	studentId, _ := strconv.ParseUint(string(id), 10, 64)
	data := []models.Student{}

	if id != "" {
		data := models.Student{Student_id: studentId}
		if config.DB.Find(&data).RecordNotFound() {
			c.JSON(404, gin.H{
				"message": "Student tidak ditemukan",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Student ditemukan",
			"data":    data,
		})
		return
	} else {
		config.DB.Find(&data)
		c.JSON(200, gin.H{
			"message": "Student ditemukan",
			"data":    data,
		})
		return
	}
}

func StoreStudent(c *gin.Context) {

	var student = models.Student{}

	c.Bind(&student)

	config.DB.Create(&student)

	c.JSON(201, gin.H{
		"message": "Berhasil menambahkan student",
		"data":    student,
	})
}

func UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	studentId, _ := strconv.ParseUint(string(id), 10, 64)
	var data = models.Student{Student_id: studentId}
	if config.DB.Find(&data).RecordNotFound() {
		c.JSON(404, gin.H{
			"message": "Student tidak ditemukan",
		})
		return
	}

	var student = models.Student{}

	c.Bind(&student)

	student.Student_id = studentId

	config.DB.Model(&data).Updates(&student)

	c.JSON(200, gin.H{
		"message": "Berhasil memperbarui student",
		"data":    student,
	})
}

func DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	studentId, _ := strconv.ParseUint(string(id), 10, 64)
	var data = models.Student{Student_id: studentId}
	if config.DB.Where("id = ?", id).Find(&data).RecordNotFound() {
		c.JSON(404, gin.H{
			"message": "Student tidak ditemukan",
		})
		return
	}

	config.DB.Delete(&data)

	c.JSON(200, gin.H{
		"message": "Berhasil menghapus student",
	})
}
