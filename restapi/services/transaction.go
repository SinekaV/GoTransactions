package services

import (
	"context"
	"time"

	"rest-api/interfaces"
	"rest-api/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionService struct {
	mongoclient *mongo.Collection
	cxt         context.Context
}

// // GetTransactionsbyCustomerId implements interfaces.Itransaction.
func (*TransactionService) GetTransactionsbyCustomerId(TransactionId int) {
	panic("unimplemented")}


func InitialiseTransactionService(collection *mongo.Collection, cxt context.Context) interfaces.Itransaction {
	return &TransactionService{collection, cxt}
}

func (t *TransactionService) CreateTransaction(transaction *models.Transaction) (*mongo.InsertOneResult,error){
	transaction.Id=primitive.NewObjectID()
	transaction.TransactionDate=time.Now()
	res, err := t.mongoclient.InsertOne(t.cxt, &transaction)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetTransactionsbyCustomerId(t *models.Transaction) []models.Transaction {
	panic("unimplemented")
}
