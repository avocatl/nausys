package ns

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// Response represents an API response.
//
// This wraps the standard http.Response returned from MMK
// and provides convenient access to things the decoded body.
type Response struct {
	*http.Response
	content []byte
}

// Error reports details on a failed API request.
type Error struct {
	Code     int            `json:"code"`
	Message  string         `json:"message"`
	Content  string         `json:"content,omitempty"`
	Response *http.Response `json:"response"` // the full response that produced the error
}

// Error function complies with the error interface.
func (e *Error) Error() string {
	return fmt.Sprintf("%v:\n%v", e.Message, e.Content)
}

// Credentials is a struct used for authentication.
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Period is a struct that defines a time interval.
type Period struct {
	PeriodFrom NausysDateTime `json:"periodFrom,omitempty"`
	PeriodTo   NausysDateTime `json:"periodTo,omitempty"`
}

// FreeYachtList is a list of all free yachts available from Nausys
type FreeYachtList struct {
	Status       string                    `json:"status,omitempty"`
	ErrorCode    int                       `json:"errorCode,omitempty"`
	PeriodFrom   *NausysDateTime           `json:"periodFrom,omitempty"`
	PeriodTo     *NausysDateTime           `json:"periodTo,omitempty"`
	FreeYachts   []FreeYacht               `json:"freeYachts,omitempty"`
	Price        YachtReservationPriceInfo `json:"price,omitempty"`
	PaymentPlans PaymentPlan               `json:"paymentPlans,omitempty"`
}

// FreeYacht is a free yacht object with timeframe of which it is free and locations.
type FreeYacht struct {
	YachtId        int64           `json:"yachtId,omitempty"`
	PeriodFrom     *NausysDateTime `json:"periodFrom,omitempty"`
	PeriodTo       *NausysDateTime `json:"periodTo,omitempty"`
	LocationFromId int64           `json:"locationFromId,omitempty"`
	LocationToId   int64           `json:"locationToId,omitempty"`
}

// YachtReservationPriceInfo contains the information on the reservation price of a yacht including discounts
// and currency.
type YachtReservationPriceInfo struct {
	PriceListPrice string      `json:"priceListPrice,omitempty"`
	ClientPrice    string      `json:"clientPrice,omitempty"`
	Currency       string      `json:"currency,omitempty"`
	Discounts      []*Discount `json:"discount,omitempty"`
}

// Discount describes the discount and type applied to an item.
type Discount struct {
	DiscountedItemId int64  `json:"discountedItemId,omitempty"`
	Amount           string `json:"amount,omitempty"`
	Type             string `json:"type,omitempty"`
}

// PaymentPlan describes a payment plan that can be used for payment of a yacht reservation.
type PaymentPlan struct {
	Date       *NausysDateTime `json:"date,omitempty"`
	Percentage int             `json:"percentage,omitempty"`
}

// NausysDateTime allows to perform (un)marshal operations with JSON
// on MMK's date time formatted response objects.
type NausysDateTime struct {
	time.Time
}

// MarshalJSON overrides the default marshal action
// for the Time struct. Returns date as YYYY-MM-DD HH:ii:ss formatted string.
func (d *NausysDateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Time.Format("2006-01-02"))
}

// UnmarshalJSON overrides the default unmarshal action
// for the Time struct.
func (d *NausysDateTime) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = strings.Trim(s, "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}
