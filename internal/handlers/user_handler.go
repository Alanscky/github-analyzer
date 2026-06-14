package handlers

import (
	"net/http"

	"github-analyzer/internal/models"
	"github-analyzer/internal/services"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	id, err := services.CreateUser(user)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuario creado",
		"id":      id,
	})
}

func GetAllUsers(c *gin.Context) {

	users, err := services.GetAllUsers()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {

	id := c.Param("id")

	user, err := services.GetUserByID(id)

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "Usuario no encontrado",
		})

		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {

	id := c.Param("id")

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	err := services.UpdateUser(id, user)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Usuario actualizado",
	})
}

func DeleteUser(c *gin.Context) {

	id := c.Param("id")

	err := services.DeleteUser(id)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Usuario eliminado",
	})
}

func Login(c *gin.Context) {

	var request models.LoginRequest

	if err :=
		c.ShouldBindJSON(&request); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	ok := services.Login(
		request.Username,
		request.Password,
	)

	if !ok {

		c.JSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "Credenciales incorrectas",
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "Login exitoso",
		},
	)
}
