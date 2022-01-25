package ns

import (
	"encoding/json"
	"net/http"
	"os"
)

// FreeYachtRequest The structure of the request will be made to the availability endpoint.
type FreeYachtRequest struct {
	Credentials            *Credentials    `json:"credentials"`
	PeriodFrom             *NausysDateTime `json:"periodFrom,omitempty"`
	PeriodTo               *NausysDateTime `json:"periodTo,omitempty"`
	YachtIds               []int64         `json:"yachts,omitempty"`
	PriceFrom              int             `json:"priceFrom,omitempty"`
	PriceTo                int             `json:"priceTo,omitempty"`
	OrderBy                int             `json:"orderby,omitempty"`
	Direction              int             `json:"direction,omitempty"`
	IgnoreAvailability     bool            `json:"ignoreAvailability,omitempty"`
	Periods                []*Period       `json:"periods,omitempty"`
	IncludeExtendedDataSet bool            `json:"includeExtendedDataSet,omitempty"`
}

// AvailabilityService operates over availability requests.
type AvailabilityService service

// GetAvailability returns availability for the specified yachts.
func (as *AvailabilityService) GetAvailability(arq *FreeYachtRequest) (ar *FreeYachtList, err error) {
	arq.Credentials = &Credentials{
		Username: os.Getenv(APIUsernameContainer),
		Password: os.Getenv(APIPasswordContainer),
	}

	req, err := as.client.NewAPIRequest(http.MethodPost, "freeYachts", arq)
	if err != nil {
		return
	}

	res, err := as.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &ar); err != nil {
		return
	}

	return
}
