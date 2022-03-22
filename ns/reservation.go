package ns

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// InfoRequest describes a request create an info reservations.
type InfoRequest struct {
	Credentials                    *Credentials `json:"credentials,omitempty"`
	ClientInfo                     *ClientInfo  `json:"client,omitempty"`
	YachtId                        int64        `json:"yachtId,omitempty"`
	PeriodFrom                     *NausysDate  `json:"periodFrom,omitempty"`
	PeriodTo                       *NausysDate  `json:"periodTo,omitempty"`
	Services                       []int64      `json:"services,omitempty"`
	Equipment                      []int64      `json:"equipment,omitempty"`
	OnlinePayment                  string       `json:"onlinePayment,omitempty"`
	PromoCode                      string       `json:"promoCode,omitempty"`
	NumberOfPayments               int          `json:"numberOfPayments,omitempty"`
	PaymentCurrency                string       `json:"paymentCurrency,omitempty"`
	UseDepositPayment              string       `json:"useDepositPayment,omitempty"`
	AgencyClientDiscountAmount     string       `json:"agencyClientDiscountAmount,omitempty"`
	AgencyClientDiscountAmountType string       `json:"agencyClientDiscountAmountType,omitempty"`
}

// ReservationsRequest describes a request to get reservations
type ReservationsRequest struct {
	Credentials           *Credentials `json:"credentials,omitempty"`
	PeriodFrom            *NausysDate  `json:"periodFrom,omitempty"`
	PeriodTo              *NausysDate  `json:"periodTo,omitempty"`
	IncludeWaitingOptions bool         `json:"includeWaitingOptions,omitempty"`
	Reservations          []int64      `json:"reservations,omitempty"`
}

// OptionBookingRequest describes a request to create an option or a booking
type OptionBookingRequest struct {
	Credentials         *Credentials `json:"credentials,omitempty"`
	ID                  int64        `json:"ID,omitempty"`
	Uuid                string       `json:"uuid,omitempty"`
	CreateWaitingOption bool         `json:"createWaitingOption,omitempty"`
}

// ReservationService operates over reservation requests.
type ReservationService service

// GetReservation gets a reservation using the reservation id.
func (rsrv *ReservationService) GetReservation(rr *ReservationsRequest) (r *Reservation, err error) {
	rr.Credentials = &Credentials{
		Username: os.Getenv(APIUsernameContainer),
		Password: os.Getenv(APIPasswordContainer),
	}

	target := fmt.Sprintf("%s/reservations", ReservationURL)

	req, err := rsrv.client.NewAPIRequest(http.MethodPost, target, rr)
	if err != nil {
		return
	}

	res, err := rsrv.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &r); err != nil {
		return
	}

	return
}

// CreateInfo sends a request to create an info reservation.
func (rsrv *ReservationService) CreateInfo(ir *InfoRequest) (r *Reservation, err error) {
	ir.Credentials = &Credentials{
		Username: os.Getenv(APIUsernameContainer),
		Password: os.Getenv(APIPasswordContainer),
	}

	target := fmt.Sprintf("%s/createInfo", BookingURL)

	req, err := rsrv.client.NewAPIRequest(http.MethodPost, target, ir)
	if err != nil {
		return
	}

	res, err := rsrv.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &r); err != nil {
		return
	}

	return
}

// CreateOption sends a request to create an option reservation.
func (rsrv *ReservationService) CreateOption(obr *OptionBookingRequest) (r *Reservation, err error) {
	obr.Credentials = &Credentials{
		Username: os.Getenv(APIUsernameContainer),
		Password: os.Getenv(APIPasswordContainer),
	}

	target := fmt.Sprintf("%s/createOption", BookingURL)

	req, err := rsrv.client.NewAPIRequest(http.MethodPost, target, obr)
	if err != nil {
		return
	}

	res, err := rsrv.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &r); err != nil {
		return
	}

	return
}

// CreateBooking sends a post request to create a booking reservation.
func (rsrv *ReservationService) CreateBooking(obr *OptionBookingRequest) (r *Reservation, err error) {
	obr.Credentials = &Credentials{
		Username: os.Getenv(APIUsernameContainer),
		Password: os.Getenv(APIPasswordContainer),
	}

	target := fmt.Sprintf("%s/createBooking", BookingURL)

	req, err := rsrv.client.NewAPIRequest(http.MethodPost, target, obr)
	if err != nil {
		return
	}

	res, err := rsrv.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &r); err != nil {
		return
	}

	return
}
