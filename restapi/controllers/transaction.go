package controllers

import (
	"net/http"
	"rest-api/interfaces"
	"rest-api/models"

	"github.com/gin-gonic/gin"
)

type TransactionController struct{
	 TransactionService  interfaces.Itransaction
}

// func InitialiseTransactionController( transaction interfaces.Itransaction)(TransactionController){
// 	return TransactionController{transactionService} 
// }

func InitTransController(transactionService interfaces.Itransaction) TransactionController {
	return TransactionController{transactionService}
}

func (t *TransactionController)CreateTransaction(ctx *gin.Context){
	var trans *models.Transaction
	if err := ctx.ShouldBindJSON(&trans); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	newtrans, err := t.TransactionService.CreateTransaction(trans)
	if(err!=nil){
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newtrans})

}
