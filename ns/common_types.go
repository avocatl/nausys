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
	ID           int64         `json:"id,omitempty"`
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

// YachtListResponse is a response that contains a list of yacht objects
type YachtListResponse struct {
	Status    string  `json:"status,omitempty"`
	ErrorCode int     `json:"errorCode,omitempty"`
	Yachts    []Yacht `json:"yachts,omitempty"`
	YachtIds  []int64 `json:"yachtIds,omitempty"`
}

// Yacht describes a single yacht
type Yacht struct {
	ID                          int64            `json:"id,omitempty"`
	Name                        string           `json:"name,omitempty"`
	CompanyId                   int64            `json:"companyId,omitempty"`
	BaseId                      int64            `json:"baseId,omitempty"`
	LocationId                  int64            `json:"locationId,omitempty"`
	YachtModelId                int64            `json:"yachtModelId,omitempty"`
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
	SteeringTypeId              int64            `json:"steeringTypeId,omitempty"`
	SailTypeId                  int64            `json:"sailTypeId,omitempty"`
	SailRenewed                 int              `json:"sailRenewed,omitempty"`
	GenoaTypeId                 int64            `json:"genoaTypeId,omitempty"`
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
	FlagsId                     []int64          `json:"flagsId,omitempty"`
	CharterType                 string           `json:"charterType,omitempty"`
	FuelTank                    int              `json:"fuelTank,omitempty"`
	WaterTank                   int              `json:"waterTank,omitempty"`
	MastLength                  float64          `json:"mastLength,omitempty"`
	PropulsionType              string           `json:"propulsionType,omitempty"`
	OneWayPeriods               []OneWayPeriod   `json:"oneWayPeriods,omitempty"`
	NumberOfRudderBlades        int              `json:"numberOfRudderBlades,omitempty"`
	EngineBuilderId             int64            `json:"engineBuilderId,omitempty"`
	HullColor                   string           `json:"hullColor,omitempty"`
	ThirdPartyInsuranceAmount   float64          `json:"thirdPartyInsuranceAmount,omitempty"`
	ThirdPartyInsuranceCurrency string           `json:"thirdPartyInsuranceCurrency,omitempty"`
	CheckInPeriods              []CheckInPeriod  `json:"checkInPeriods,omitempty"`
}

// YachtEquipment is amounts of equipment found on a yacht with their descriptions
type YachtEquipment struct {
	ID          int64             `json:"id,omitempty"`
	Quantity    int               `json:"quantity,omitempty"`
	EquipmentId int64             `json:"equipmentId,omitempty"`
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
	SeasonId                 int64                      `json:"seasonId,omitempty"`
	BaseId                   int64                      `json:"baseId,omitempty"`
	LocationId               int64                      `json:"locationId,omitempty"`
	AdditionalYachtEquipment []AdditionalYachtEquipment `json:"additionalYachtEquipment,omitempty"`
	Services                 []YachtService             `json:"services,omitempty"`
	Prices                   []YachtPrice               `json:"prices,omitempty"`
	RegularDiscounts         []Discount                 `json:"regularDiscounts,omitempty"`
}

// OneWayPeriod describes the period it would take to travel one way
type OneWayPeriod struct {
	Id         int64          `json:"id,omitempty"`
	PeriodFrom NausysDateTime `json:"periodFrom,omitempty"`
	PeriodTo   NausysDateTime `json:"periodTo,omitempty"`
	BaseId     int64          `json:"baseId,omitempty"`
	LocationId int64          `json:"locationId,omitempty"`
}

// CheckInPeriod describes the minimum amount of days that a yacht can be booked and check in days
type CheckInPeriod struct {
	DateFrom                   NausysDateTime `json:"dateFrom,omitempty"`
	DateTo                     NausysDateTime `json:"dateTo,omitempty"`
	MinimalReservationDuration int            `json:"minimalReservationDuration,omitempty"`
	CheckInMonday              bool           `json:"checkInMonday,omitempty"`
	CheckInTuesday             bool           `json:"checkInTuesday,omitempty"`
	CheckInWednesday           bool           `json:"checkInWednesday,omitempty"`
	CheckInThursday            bool           `json:"checkInThursday,omitempty"`
	CheckInFriday              bool           `json:"checkInFriday,omitempty"`
	CheckInSaturday            bool           `json:"checkInSaturday,omitempty"`
	CheckInSunday              bool           `json:"checkInSunday,omitempty"`
	CheckOutMonday             bool           `json:"checkOutMonday,omitempty"`
	CheckOutTuesday            bool           `json:"checkOutTuesday,omitempty"`
	CheckOutWednesday          bool           `json:"checkOutWednesday,omitempty"`
	CheckOutThursday           bool           `json:"checkOutThursday,omitempty"`
	CheckOutFriday             bool           `json:"checkOutFriday,omitempty"`
	CheckOutSaturday           bool           `json:"checkOutSaturday,omitempty"`
	CheckOutSunday             bool           `json:"checkOutSunday,omitempty"`
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
	Id                        int64             `json:"id,omitempty"`
	Quantity                  int               `json:"quantity,omitempty"`
	Price                     string            `json:"price,omitempty"`
	Currency                  string            `json:"currency,omitempty"`
	EquipmentId               int64             `json:"equipmentId,omitempty"`
	Comment                   InternationalText `json:"comment,omitempty"`
	PriceMeasureId            int64             `json:"priceMeasureId,omitempty"`
	CalculationType           string            `json:"calculationType,omitempty"`
	Condition                 InternationalText `json:"condition,omitempty"`
	Amount                    string            `json:"amount,omitempty"`
	AmountIsPercentage        bool              `json:"amountIsPercentage,omitempty"`
	PercentageCalculationType string            `json:"percentageCalculationType,omitempty"`
	ValidForBases             []int64           `json:"validForBases,omitempty"`
	MinimumPrice              string            `json:"minimumPrice,omitempty"`
}

// YachtService describes a service used with a yacht
type YachtService struct {
	Id                        int64             `json:"id,omitempty"`
	ServiceId                 int64             `json:"serviceId,omitempty"`
	Price                     string            `json:"price,omitempty"`
	Currency                  string            `json:"currency,omitempty"`
	PriceMeasureId            int64             `json:"priceMeasureId,omitempty"`
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

// YachtPrice describes a yacht price
type YachtPrice struct {
	Id         int64   `json:"id,omitempty"`
	DateFrom   string  `json:"dateFrom,omitempty"`
	DateTo     string  `json:"dateTo,omitempty"`
	Price      string  `json:"price,omitempty"`
	Currency   string  `json:"currency,omitempty"`
	Type       string  `json:"type,omitempty"`
	LocationId []int64 `json:"locationId,omitempty"`
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
