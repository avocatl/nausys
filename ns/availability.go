package ns

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// AvailabilityRequest The structure of the request will be made to the availability endpoint.
type AvailabilityRequest struct {
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
func (as *AvailabilityService) GetAvailability(arq AvailabilityRequest) (ar []*FreeYachtList, err error) {
	arq.Credentials = &Credentials{
		Username: os.Getenv(APIUsernameContainer),
		Password: os.Getenv(APIPasswordContainer),
	}

	var target = fmt.Sprintf("freeYachts")

	body, err := json.Marshal(arq)
	if err != nil {
		return
	}

	req, err := as.client.NewAPIRequest(http.MethodPost, target, body)
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
