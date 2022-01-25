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

// FreeYachtListResponse is a list of all free yachts available from Nausys
type FreeYachtListResponse struct {
	Status       string          `json:"status,omitempty"`
	ErrorCode    int             `json:"errorCode,omitempty"`
	PeriodFrom   *NausysDateTime `json:"periodFrom,omitempty"`
	PeriodTo     *NausysDateTime `json:"periodTo,omitempty"`
	FreeYachts   []FreeYacht     `json:"freeYachts,omitempty"`
	PaymentPlans PaymentPlan     `json:"paymentPlans,omitempty"`
}

// FreeYacht is a free yacht object with timeframe of which it is free and locations.
type FreeYacht struct {
	YachtId        int64                     `json:"yachtId,omitempty"`
	PeriodFrom     *NausysDateTime           `json:"periodFrom,omitempty"`
	PeriodTo       *NausysDateTime           `json:"periodTo,omitempty"`
	Price          YachtReservationPriceInfo `json:"price,omitempty"`
	LocationFromId int64                     `json:"locationFromId,omitempty"`
	LocationToId   int64                     `json:"locationToId,omitempty"`
}

// YachtReservationPriceInfo contains the information on the reservation price of a yacht including discounts
// and currency.
type YachtReservationPriceInfo struct {
	PriceListPrice           string      `json:"priceListPrice,omitempty"`
	ClientPrice              string      `json:"clientPrice,omitempty"`
	Currency                 string      `json:"currency,omitempty"`
	DepositAmount            string      `json:"depositAmount,omitempty"`
	DepositWhenInsuredAmount string      `json:"depositWhenInsuredAmount,omitempty"`
	Discounts                []*Discount `json:"discounts,omitempty"`
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

// CompanyListResponse is a list of all companies from Nausys.
type CompanyListResponse struct {
	Status    string    `json:"status,omitempty"`
	ErrorCode int       `json:"errorCode,omitempty"`
	Company   []Company `json:"companies,omitempty"`
}

// Company is a company object with full charter company information.
type Company struct {
	CountryID    int64         `json:"countryId,omitempty"`
	Name         string        `json:"name,omitempty"`
	Address      string        `json:"address,omitempty"`
	City         string        `json:"city,omitempty"`
	Zip          string        `json:"zip,omitempty"`
	Phone        string        `json:"phone,omitempty"`
	Fax          string        `json:"fax,omitempty"`
	Mobile       string        `json:"mobile,omitempty"`
	Vatcode      string        `json:"vatcode,omitempty"`
	Web          string        `json:"web,omitempty"`
	Email        string        `json:"email,omitempty"`
	Pac          bool          `json:"pac,omitempty"`
	BankAccounts []BankAccount `json:"bankAccounts,omitempty"`
}

// BankAccount is a company bank account information.
type BankAccount struct {
	BankName      string `json:"bankName,omitempty"`
	BankAddress   string `json:"bankAddress,omitempty"`
	AccountNumber string `json:"accountNumber,omitempty"`
	Swift         string `json:"swift,omitempty"`
	Iban          string `json:"iban,omitempty"`
}

// OccupancyListResponse is a list of all a company occupancy from Nausys.
type OccupancyListResponse struct {
	CompanyId    int64         `json:"companyId,omitempty"`
	Year         uint          `json:"year,omitempty"`
	Reservations []Reservation `json:"reservations,omitempty"`
}

// Reservation is a reservation object used in occupancy.
type Reservation struct {
	ID              int64           `json:"id,omitempty"`
	YachtID         int64           `json:"yachtId,omitempty"`
	LocationFromID  int64           `json:"locationFromId,omitempty"`
	LocationToID    int64           `json:"locationToId,omitempty"`
	ReservationType string          `json:"reservationType,omitempty"`
	PeriodFrom      *NausysDateTime `json:"periodFrom,omitempty"`
	CheckInTime     *NausysDateTime `json:"checkInTime,omitempty"`
	PeriodTo        *NausysDateTime `json:"periodTo,omitempty"`
	CheckOutTime    *NausysDateTime `json:"checkOutTime,omitempty"`
}

// NausysDateTime allows to perform (un)marshal operations with JSON
// on MMK's date time formatted response objects.
type NausysDateTime struct {
	time.Time
}

// MarshalJSON overrides the default marshal action
// for the Time struct. Returns date as YYYY-MM-DD HH:ii:ss formatted string.
func (d *NausysDateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Time.Format("02.01.2006"))
}

// UnmarshalJSON overrides the default unmarshal action
// for the Time struct.
func (d *NausysDateTime) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = strings.Trim(s, "\"")
	t, err := time.Parse("02.01.2006", s)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}
