package ns

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// OccupancyRequest The structure of the request will be made to the occupancy endpoint.
type OccupancyRequest struct {
	Credentials *Credentials `json:"credentials"`
}

// OccupancyService operates over occupancy requests.
type OccupancyService service

// Occupancy Provides all reservations for specified company in specified
// year regardless who made them.
func (ocs *OccupancyService) Occupancy(arq *OccupancyRequest, companyID int64, year uint) (ol []*OccupancyList, err error) {
	arq.Credentials = &Credentials{
		Username: os.Getenv(APIUsernameContainer),
		Password: os.Getenv(APIPasswordContainer),
	}

	req, err := ocs.client.NewAPIRequest(http.MethodPost, fmt.Sprintf("occupancy/%d/%d", companyID, year), arq)
	if err != nil {
		return
	}

	res, err := ocs.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &ol); err != nil {
		return
	}

	return
}
