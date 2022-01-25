package ns

import (
	"encoding/json"
	"net/http"
	"os"
)

// CompanyRequest The structure of the request will be made to the company endpoint.
type CompanyRequest struct {
	Credentials *Credentials `json:"credentials"`
}

// CompanyService operates over company requests.
type CompanyService service

// All returns all companies.
func (cs *CompanyService) All(crq *CompanyRequest) (clr *CompanyListResponse, err error) {
	crq.Credentials = &Credentials{
		Username: os.Getenv(APIUsernameContainer),
		Password: os.Getenv(APIPasswordContainer),
	}

	req, err := cs.client.NewAPIRequest(http.MethodPost, "charterCompanies", crq)
	if err != nil {
		return
	}

	res, err := cs.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &clr); err != nil {
		return
	}

	return
}
