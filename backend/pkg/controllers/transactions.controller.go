package controllers

import (
	"encoding/csv"
	"net/http"

	"zterhes/budget/pkg/types"

	"github.com/gin-gonic/gin"
)

func UploadTransactions(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "Could not get uploaded file")
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		c.String(http.StatusInternalServerError, "Could not read CSV file")
		return
	}

	var parsedTransactions []types.TransactionRequest
	for i, row := range records {
		if i == 0 {
			continue
		}
		parsedTransactions = append(parsedTransactions, types.TransactionRequest{
			ID:                row[0],
			Status:            row[1],
			Direction:         row[2],
			CreatedOn:         row[3],
			FinishedOn:        row[4],
			SourceFeeAmount:   row[5],
			SourceFeeCurrency: row[6],
			TargetfeeAmount:   row[7],
			TargetFeeCurrency: row[8],
			SourceName:        row[9],
			SourceAmount:      row[10],
			SourceCurrency:    row[11],
			TargetName:        row[12],
			TargetAmount:      row[13],
			TargetCurrency:    row[14],
			ExchangeRate:      row[15],
			Refference:        row[16],
			Batch:             row[17],
		})
	}

	c.JSON(http.StatusOK, parsedTransactions)
}
