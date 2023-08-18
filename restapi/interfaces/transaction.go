package interfaces

import (
	"rest-api/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type Itransaction interface{
 CreateTransaction(*models.Transaction)(*mongo.InsertOneResult,error)
 GetTransactionsbyCustomerId(TransactionId int )
}