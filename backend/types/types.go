package types

import "time"

type Transaction struct {
	ID                string
	Status            string
	Direction         string
	CreatedOn         time.Time
	FinishedOn        time.Time
	SourceFeeAmount   float64
	SourceFeeCurrency string
	TargetfeeAmount   float64
	TargetFeeCurrency string
	SourceName        string
	SourceAmount      float64
	SourceCurrency    string
	TargetName        string
	TargetAmount      float64
	TargetCurrency    string
	ExchangeRate      float64
	Refference        string
	Batch             string
}

type TransactionRequest struct {
	ID                string
	Status            string
	Direction         string
	CreatedOn         string
	FinishedOn        string
	SourceFeeAmount   string
	SourceFeeCurrency string
	TargetfeeAmount   string
	TargetFeeCurrency string
	SourceName        string
	SourceAmount      string
	SourceCurrency    string
	TargetName        string
	TargetAmount      string
	TargetCurrency    string
	ExchangeRate      string
	Refference        string
	Batch             string
}
