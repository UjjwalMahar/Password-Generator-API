package controllers

import (
	"net/http"

	"github.com/UjjwalMahar/Password-Gen/logic"
	"github.com/gin-gonic/gin"
)

//Func for Root endpoint
func GetRoot(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "Password Generator Created with ❤️ by Ujjwal Mahar",
	})
}

func PostPasswordGen(c *gin.Context){

	//Taking input by user
	var PassInfo struct{
		Length int
		ReqSpecialChar bool
		ReqDigit bool
	}

	//bind
	if c.Bind(&PassInfo) != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to read the body",
		})

		return
	}

	//getting generated password from my logic

	password := logic.PasswordGen(PassInfo.Length, PassInfo.ReqSpecialChar , PassInfo.ReqDigit)
	if password == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate a password"})
		return
	}

	c.JSON(200, gin.H{
		"GeneratedPassword": password,
	})


}