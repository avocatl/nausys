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
	PeriodFrom *NausysDate `json:"periodFrom,omitempty"`
	PeriodTo   *NausysDate `json:"periodTo,omitempty"`
}

// FreeYachtListResponse is a list of all free yachts available from Nausys
type FreeYachtListResponse struct {
	Status       string      `json:"status,omitempty"`
	ErrorCode    int         `json:"errorCode,omitempty"`
	PeriodFrom   *NausysDate `json:"periodFrom,omitempty"`
	PeriodTo     *NausysDate `json:"periodTo,omitempty"`
	FreeYachts   []FreeYacht `json:"freeYachts,omitempty"`
	PaymentPlans PaymentPlan `json:"paymentPlans,omitempty"`
}

// FreeYacht is a free yacht object with timeframe of which it is free and locations.
type FreeYacht struct {
	YachtId        int64                     `json:"yachtId,omitempty"`
	PeriodFrom     *NausysDate               `json:"periodFrom,omitempty"`
	PeriodTo       *NausysDate               `json:"periodTo,omitempty"`
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
	DiscountedItemId int64   `json:"discountedItemId,omitempty"`
	Amount           float32 `json:"amount,omitempty"`
	Type             string  `json:"type,omitempty"`
}

// PaymentPlan describes a payment plan that can be used for payment of a yacht reservation.
type PaymentPlan struct {
	Date       *NausysDate `json:"date,omitempty"`
	Percentage int         `json:"percentage,omitempty"`
}

// BookingPaymentPlan describes a payment plan that was created as part of a bookingd
type BookingPaymentPlan struct {
	ID                      int64           `json:"id,omitempty"`
	Date                    *NausysDate     `json:"date,omitempty"`
	Amount                  string          `json:"amount,omitempty"`
	AmountPaymentCurrency   string          `json:"amountPaymentCurrency,omitempty"`
	Paid                    bool            `json:"paid,omitempty"`
	OnlinePaymentLink       string          `json:"onlinePaymentLink,omitempty"`
	OnlinePaymentValidUTill *NausysDateTime `json:"onlinePaymentValidUTill,omitempty"`
}

// CompanyListResponse is a list of all companies from Nausys.
type CompanyListResponse struct {
	Status    string    `json:"status,omitempty"`
	ErrorCode int       `json:"errorCode,omitempty"`
	Company   []Company `json:"companies,omitempty"`
}

// Company is a company object with full charter company information.
type Company struct {
	ID           int64         `json:"id,omitempty"`
	CountryID    int64         `json:"countryId,omitempty"`
	Name         string        `json:"name,omitempty"`
	CompanyName  string        `json:"companyName,omitempty"`
	Address      string        `json:"address,omitempty"`
	City         string        `json:"city,omitempty"`
	Zip          string        `json:"zip,omitempty"`
	Phone        string        `json:"phone,omitempty"`
	Mobile       string        `json:"mobile,omitempty"`
	Vatcode      string        `json:"vatcode,omitempty"`
	Web          string        `json:"web,omitempty"`
	Email        string        `json:"email,omitempty"`
	Pac          int           `json:"pac,omitempty"`
	BankAccounts []BankAccount `json:"bankAccounts,omitempty"`
}

// BankAccount is a company bank account information.
type BankAccount struct {
	BankName    string `json:"bankName,omitempty"`
	BankAddress string `json:"bankAddress,omitempty"`
	Swift       string `json:"swift,omitempty"`
	Iban        string `json:"iban,omitempty"`
}

// OccupancyListResponse is a list of all a company occupancy from Nausys.
type OccupancyListResponse struct {
	Status       string        `json:"status,omitempty"`
	CompanyId    int64         `json:"companyId,omitempty"`
	Year         uint          `json:"year,omitempty"`
	Reservations []Reservation `json:"reservations,omitempty"`
}

// Reservation is a reservation object used in occupancy.
type Reservation struct {
	ID              int64       `json:"id,omitempty"`
	YachtID         int64       `json:"yachtId,omitempty"`
	LocationFromID  int64       `json:"locationFromId,omitempty"`
	LocationToID    int64       `json:"locationToId,omitempty"`
	ReservationType string      `json:"reservationType,omitempty"`
	PeriodFrom      *NausysDate `json:"periodFrom,omitempty"`
	CheckInTime     *NausysTime `json:"checkInTime,omitempty"`
	PeriodTo        *NausysDate `json:"periodTo,omitempty"`
	CheckOutTime    *NausysTime `json:"checkOutTime,omitempty"`
}

// YachtListResponse is a response that contains a list of yacht objects
type YachtListResponse struct {
	Status    string  `json:"status,omitempty"`
	ErrorCode int     `json:"errorCode,omitempty"`
	Yachts    []Yacht `json:"yachts,omitempty"`
	YachtIDs  []int64 `json:"yachtIds,omitempty"`
}

// Yacht describes a single yacht
type Yacht struct {
	ID                          int64            `json:"id,omitempty"`
	Name                        string           `json:"name,omitempty"`
	CompanyID                   int64            `json:"companyId,omitempty"`
	BaseID                      int64            `json:"baseId,omitempty"`
	LocationID                  int64            `json:"locationId,omitempty"`
	YachtModelID                int64            `json:"yachtModelId,omitempty"`
	Draft                       float64          `json:"draft,omitempty"`
	Cabins                      int              `json:"cabins,omitempty"`
	CabinCrew                   int              `json:"cabinCrew,omitempty"`
	BerthsCabin                 int              `json:"berthsCabin,omitempty"`
	BerthsSalon                 int              `json:"berthsSalon,omitempty"`
	BerthsCrew                  int              `json:"berthsCrew,omitempty"`
	BerthsTotal                 int              `json:"berthsTotal,omitempty"`
	Wc                          int              `json:"wc,omitempty"`
	WcCrew                      int              `json:"wcCrew,omitempty"`
	Engines                     int              `json:"engines,omitempty"`
	EnginePower                 float64          `json:"enginePower,omitempty"`
	SteeringTypeID              int64            `json:"steeringTypeId,omitempty"`
	SailTypeID                  int64            `json:"sailTypeId,omitempty"`
	SailRenewed                 int              `json:"sailRenewed,omitempty"`
	GenoaTypeID                 int64            `json:"genoaTypeId,omitempty"`
	GenoaRenewed                int              `json:"genoaRenewed,omitempty"`
	StandardYachtEquipment      []YachtEquipment `json:"standardYachtEquipment,omitempty"`
	Euminia                     Euminia          `json:"euminia,omitempty"`
	MainPictureUrl              string           `json:"mainPictureUrl,omitempty"`
	PicturesUrl                 []string         `json:"picturesUrl,omitempty"`
	Commission                  float64          `json:"commission,omitempty"`
	Deposit                     float64          `json:"deposit,omitempty"`
	DepositCurrency             string           `json:"depositCurrency,omitempty"`
	MaxDiscount                 float64          `json:"maxDiscount,omitempty"`
	SeasonSpecificData          []YachtSeason    `json:"seasonSpecificData,omitempty"`
	NeedsOptionApproval         bool             `json:"needsOptionApproval,omitempty"`
	CanMakeBookingFixed         bool             `json:"canMakeBookingFixed,omitempty"`
	FlagsID                     []int64          `json:"flagsId,omitempty"`
	CharterType                 string           `json:"charterType,omitempty"`
	FuelTank                    int              `json:"fuelTank,omitempty"`
	WaterTank                   int              `json:"waterTank,omitempty"`
	MastLength                  float64          `json:"mastLength,omitempty"`
	PropulsionType              string           `json:"propulsionType,omitempty"`
	OneWayPeriods               []OneWayPeriod   `json:"oneWayPeriods,omitempty"`
	NumberOfRudderBlades        int              `json:"numberOfRudderBlades,omitempty"`
	EngineBuilderID             int64            `json:"engineBuilderId,omitempty"`
	HullColor                   string           `json:"hullColor,omitempty"`
	ThirdPartyInsuranceAmount   float64          `json:"thirdPartyInsuranceAmount,omitempty"`
	ThirdPartyInsuranceCurrency string           `json:"thirdPartyInsuranceCurrency,omitempty"`
	CheckInPeriods              []CheckInPeriod  `json:"checkInPeriods,omitempty"`
}

// YachtEquipment is amounts of equipment found on a yacht with their descriptions
type YachtEquipment struct {
	ID          int64             `json:"id,omitempty"`
	Quantity    int               `json:"quantity,omitempty"`
	EquipmentID int64             `json:"equipmentId,omitempty"`
	Highlight   bool              `json:"highlight,omitempty"`
	Comment     InternationalText `json:"comment,omitempty"`
}

// Euminia is an overall rating given to a yacht
type Euminia struct {
	Cleanliness      string `json:"cleanliness,omitempty"`
	Equipment        string `json:"equipment,omitempty"`
	PersonalService  string `json:"personalService,omitempty"`
	PricePerformance string `json:"pricePerformance,omitempty"`
	Recommendation   string `json:"recommendation,omitempty"`
	Total            string `json:"total,omitempty"`
	Reviews          string `json:"reviews,omitempty"`
}

// YachtSeason describes a season tied to a yacht with available equipment, services, prices and discounts for that season
type YachtSeason struct {
	SeasonID                 int64                      `json:"seasonId,omitempty"`
	BaseID                   int64                      `json:"baseId,omitempty"`
	LocationID               int64                      `json:"locationId,omitempty"`
	AdditionalYachtEquipment []AdditionalYachtEquipment `json:"additionalYachtEquipment,omitempty"`
	Services                 []YachtService             `json:"services,omitempty"`
	Prices                   []YachtPrice               `json:"prices,omitempty"`
	RegularDiscounts         []Discount                 `json:"regularDiscounts,omitempty"`
}

// OneWayPeriod describes the period it would take to travel one way
type OneWayPeriod struct {
	ID         int64       `json:"id,omitempty"`
	PeriodFrom *NausysDate `json:"periodFrom,omitempty"`
	PeriodTo   *NausysDate `json:"periodTo,omitempty"`
	BaseID     int64       `json:"baseId,omitempty"`
	LocationID int64       `json:"locationId,omitempty"`
}

// CheckInPeriod describes the minimum amount of days that a yacht can be booked and check in days
type CheckInPeriod struct {
	DateFrom                   *NausysDate `json:"dateFrom,omitempty"`
	DateTo                     *NausysDate `json:"dateTo,omitempty"`
	MinimalReservationDuration int         `json:"minimalReservationDuration,omitempty"`
	CheckInMonday              bool        `json:"checkInMonday,omitempty"`
	CheckInTuesday             bool        `json:"checkInTuesday,omitempty"`
	CheckInWednesday           bool        `json:"checkInWednesday,omitempty"`
	CheckInThursday            bool        `json:"checkInThursday,omitempty"`
	CheckInFriday              bool        `json:"checkInFriday,omitempty"`
	CheckInSaturday            bool        `json:"checkInSaturday,omitempty"`
	CheckInSunday              bool        `json:"checkInSunday,omitempty"`
	CheckOutMonday             bool        `json:"checkOutMonday,omitempty"`
	CheckOutTuesday            bool        `json:"checkOutTuesday,omitempty"`
	CheckOutWednesday          bool        `json:"checkOutWednesday,omitempty"`
	CheckOutThursday           bool        `json:"checkOutThursday,omitempty"`
	CheckOutFriday             bool        `json:"checkOutFriday,omitempty"`
	CheckOutSaturday           bool        `json:"checkOutSaturday,omitempty"`
	CheckOutSunday             bool        `json:"checkOutSunday,omitempty"`
}

// InternationalText is a list of translations
type InternationalText struct {
	TextDE string `json:"textDE,omitempty"`
	TextEN string `json:"textEN,omitempty"`
	TextHR string `json:"textHR,omitempty"`
	TextIT string `json:"textIT,omitempty"`
	TextSI string `json:"textSI,omitempty"`
	TextRU string `json:"textRU,omitempty"`
	TextCZ string `json:"textCZ,omitempty"`
	TextFR string `json:"textFR,omitempty"`
	TextPL string `json:"textPL,omitempty"`
	TextSK string `json:"textSK,omitempty"`
	TextNL string `json:"textNL,omitempty"`
	TextES string `json:"textES,omitempty"`
}

// AdditionalYachtEquipment describes equipment that can be booked with a yacht
type AdditionalYachtEquipment struct {
	ID                        int64             `json:"id,omitempty"`
	Quantity                  int               `json:"quantity,omitempty"`
	Price                     string            `json:"price,omitempty"`
	Currency                  string            `json:"currency,omitempty"`
	EquipmentID               int64             `json:"equipmentId,omitempty"`
	Comment                   InternationalText `json:"comment,omitempty"`
	PriceMeasureID            int64             `json:"priceMeasureId,omitempty"`
	CalculationType           string            `json:"calculationType,omitempty"`
	Condition                 InternationalText `json:"condition,omitempty"`
	Amount                    string            `json:"amount,omitempty"`
	AmountIsPercentage        bool              `json:"amountIsPercentage,omitempty"`
	PercentageCalculationType string            `json:"percentageCalculationType,omitempty"`
	ValidForBases             []int64           `json:"validForBases,omitempty"`
	MinimumPrice              string            `json:"minimumPrice,omitempty"`
}

// Payment describes a payment object
type Payment struct {
	ID                    int64  `json:"id,omitempty"`
	Date                  string `json:"date,omitempty"`
	Amount                string `json:"amount,omitempty"`
	AmountPaymentCurrency string `json:"amountPaymentCurrency,omitempty"`
	PaymentCurrency       string `json:"paymentCurrency,omitempty"`
}

// YachtService describes a service used with a yacht
type YachtService struct {
	ID                        int64             `json:"id,omitempty"`
	ServiceID                 int64             `json:"serviceId,omitempty"`
	Price                     string            `json:"price,omitempty"`
	Currency                  string            `json:"currency,omitempty"`
	PriceMeasureID            int64             `json:"priceMeasureId,omitempty"`
	CalculationType           string            `json:"calculationType,omitempty"`
	Description               InternationalText `json:"description,omitempty"`
	Obligatory                bool              `json:"obligatory,omitempty"`
	Amount                    string            `json:"amount,omitempty"`
	AmountIsPercentage        bool              `json:"amountIsPercentage,omitempty"`
	PercentageCalculationType string            `json:"percentageCalculationType,omitempty"`
	ValidPeriodFrom           string            `json:"validPeriodFrom,omitempty"`
	ValidPeriodTo             string            `json:"validPeriodTo,omitempty"`
	ValidMinPax               int               `json:"validMinPax,omitempty"`
	ValidMaxPax               int               `json:"validMaxPax,omitempty"`
	ValidForBases             []int64           `json:"validForBases,omitempty"`
	MinimumPrice              string            `json:"minimumPrice,omitempty"`
}

// YachtPrice describes a yacht price.
type YachtPrice struct {
	ID         int64   `json:"id,omitempty"`
	DateFrom   string  `json:"dateFrom,omitempty"`
	DateTo     string  `json:"dateTo,omitempty"`
	Price      float32 `json:"price,omitempty"`
	Currency   string  `json:"currency,omitempty"`
	Type       string  `json:"type,omitempty"`
	LocationID []int64 `json:"locationId,omitempty"`
}

// ClientInfo describes a client object.
type ClientInfo struct {
	Company   bool   `json:"company,omitempty"`
	VatNr     string `json:"vatNr,omitempty"`
	Name      string `json:"name,omitempty"`
	Surname   string `json:"surname,omitempty"`
	Address   string `json:"address,omitempty"`
	Zip       string `json:"zip,omitempty"`
	City      string `json:"city,omitempty"`
	CountryId int64  `json:"countryId,omitempty"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Mobile    string `json:"mobile,omitempty"`
	Skype     string `json:"skype,omitempty"`
}

// ReservationsList is the response struct.
type ReservationsList struct {
	Status       string             `json:"status,omitempty"`
	ErrorCode    int                `json:"errorCode,omitempty"`
	Reservations []*ReservationInfo `json:"reservations,omitempty"`
}

// ReservationInfo defines all the information and types that a reservation will contain
type ReservationInfo struct {
	ID                             int64                       `json:"id,omitempty"`
	Uuid                           string                      `json:"uuid,omitempty"`
	ReservationStatus              string                      `json:"reservationStatus,omitempty"`
	WaitingForOption               bool                        `json:"waitingForOption,omitempty"`
	YachtID                        int64                       `json:"yachtID,omitempty"`
	BaseFromId                     int64                       `json:"baseFromId,omitempty"`
	BaseToId                       int64                       `json:"baseToId,omitempty"`
	LocationFromId                 int64                       `json:"locationFromId,omitempty"`
	LocationToId                   int64                       `json:"locationToId,omitempty"`
	PeriodFrom                     *NausysDateTime             `json:"periodFrom,omitempty"`
	PeriodTo                       *NausysDateTime             `json:"periodTo,omitempty"`
	OptionTill                     string                      `json:"optionTill,omitempty"`
	Agency                         string                      `json:"agency,omitempty"`
	AgencyVatID                    string                      `json:"agencyVATID,omitempty"`
	Client                         *ClientInfo                 `json:"client,omitempty"`
	Discounts                      []*Discount                 `json:"discounts,omitempty"`
	AdditionalEquipment            []*AdditionalYachtEquipment `json:"additionalEquipment,omitempty"`
	Services                       []*YachtsService            `json:"services,omitempty"`
	PriceListPrice                 string                      `json:"priceListPrice,omitempty"`
	AgencyPrice                    string                      `json:"agencyPrice,omitempty"`
	ClientPrice                    string                      `json:"clientPrice,omitempty"`
	Currency                       string                      `json:"currency,omitempty"`
	PaymentCurrency                string                      `json:"paymentCurrency,omitempty"`
	LocalizedFinalPrice            string                      `json:"localizedFinalPrice,omitempty"`
	OnlinePaymentAmount            string                      `json:"onlinePaymentAmount,omitempty"`
	Approved                       bool                        `json:"approved,omitempty"`
	CrewListLink                   string                      `json:"crewlistlink,omitempty"`
	CreatedDate                    string                      `json:"createdDate,omitempty"`
	PaymentPlan                    []*PaymentPlan              `json:"paymentPlan,omitempty"`
	Payments                       []*Payment                  `json:"payments,omitempty"`
	UseDepositPayment              bool                        `json:"useDepositPayment,omitempty"`
	NumberOfPayments               int                         `json:"numberOfPayments,omitempty"`
	OwnerBooking                   bool                        `json:"ownerBooking,omitempty"`
	AgencyAdditionalDiscountAmount string                      `json:"agencyAdditionalDiscountAmount,omitempty"`
	AgencyClientFinalPrice         string                      `json:"agencyClientFinalPrice,omitempty"`
}

// NausysDate allows to perform (un)marshal operations with JSON
// on Nausys's date formatted response objects.
type NausysDate struct {
	time.Time
}

// NausysDateTime allows to perform (un)marshal operations with JSON
// on Nausys's date time formatted response objects.
type NausysDateTime struct {
	time.Time
}

// NausysTime allows to (un)marshalling operations with JSON on
// Nausys's time formatted response objects.
type NausysTime struct {
	time.Time
}

// MarshalJSON overrides the default marshal action
// for the Time struct. Returns date as YYYY-MM-DD formatted string.
func (d *NausysDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Time.Format("02.01.2006"))
}

// UnmarshalJSON overrides the default unmarshal action
// for the Time struct.
func (d *NausysDate) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = strings.Trim(s, "\"")
	t, err := time.Parse("02.01.2006", s)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

// MarshalJSON overrides the default marshal action
// for the Time struct. Returns date as YYYY-MM-DD HH:ii:ss formatted string.
func (dt *NausysDateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(dt.Time.Format("02.01.2006 15:04"))
}

// UnmarshalJSON overrides the default unmarshal action
// for the Time struct.
func (dt *NausysDateTime) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = strings.Trim(s, "\"")
	t, err := time.Parse("02.01.2006 15:04", s)
	if err != nil {
		return err
	}
	dt.Time = t
	return nil
}

// MarshalJSON overrides the default marshal action
// for the Time struct. Returns date as YYYY-MM-DD HH:ii:ss formatted string.
func (t *NausysTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Time.Format("15:04:05"))
}

// UnmarshalJSON overrides the default unmarshal action
// for the Time struct.
func (t *NausysTime) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = strings.Trim(s, "\"")
	ti, err := time.Parse("15:04:05", s)
	if err != nil {
		return err
	}
	t.Time = ti
	return nil
}
