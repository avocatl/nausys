package ns

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// OccupancyService operates over occupancy requests.
type OccupancyService service

// All Provides all reservations for specified company in specified
// year regardless who made them.
func (ocs *OccupancyService) All(companyID int64, year uint) (olr *OccupancyListResponse, err error) {
	c := &Credentials{
		Username: os.Getenv(APIUsernameContainer),
		Password: os.Getenv(APIPasswordContainer),
	}

	target := fmt.Sprintf("%s/occupancy/%d/%d", ReservationURL, companyID, year)

	req, err := ocs.client.NewAPIRequest(http.MethodPost, target, c)
	if err != nil {
		return
	}

	res, err := ocs.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &olr); err != nil {
		return
	}

	return
}
