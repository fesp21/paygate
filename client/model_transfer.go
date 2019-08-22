/*
 * Paygate API
 *
 * Paygate is a RESTful API enabling Automated Clearing House ([ACH](https://en.wikipedia.org/wiki/Automated_Clearing_House)) transactions to be submitted and received without a deep understanding of a full NACHA file specification.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

type Transfer struct {
	// Optional ID to uniquely identify this transfer. If omitted, one will be generated
	ID string `json:"ID,omitempty"`
	// Type of transaction being actioned against the receiving institution. Expected values are pull (debits) or push (credits). Only one period used to signify decimal value will be included.
	TransferType string `json:"transferType,omitempty"`
	// Amount of money. USD - United States.
	Amount string `json:"amount"`
	// ID of the Originator account initiating the transfer.
	Originator string `json:"originator"`
	// ID of the Originator Depository to be be used to override the default depository.
	OriginatorDepository string `json:"originatorDepository,omitempty"`
	// ID of the Receiver account the transfer was sent to.
	Receiver string `json:"receiver"`
	// ID of the Receiver Depository to be used to override the default depository
	ReceiverDepository string `json:"receiverDepository,omitempty"`
	// Brief description of the transaction, that may appear on the receiving entity’s financial statement
	Description string `json:"description"`
	// Standard Entry Class code will be generated based on Receiver type for CCD and PPD
	StandardEntryClassCode string `json:"standardEntryClassCode,omitempty"`
	// Defines the state of the Transfer
	Status string `json:"status,omitempty"`
	// When set to true this indicates the transfer should be processed the same day if possible.
	SameDay   bool      `json:"sameDay,omitempty"`
	Created   time.Time `json:"created,omitempty"`
	CCDDetail CcdDetail `json:"CCDDetail,omitempty"`
	IATDetail IatDetail `json:"IATDetail,omitempty"`
	TELDetail TelDetail `json:"TELDetail,omitempty"`
	WEBDetail WebDetail `json:"WEBDetail,omitempty"`
}