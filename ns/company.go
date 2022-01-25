package ns

import (
	"encoding/json"
	"net/http"
	"os"
)

// CompanyRequest The structure of the request will be made to the company endpoint.
type CompanyRequest struct {
	Credentials            *Credentials    `json:"credentials"`
}

// CompanyService operates over company requests.
type CompanyService service

// Companies returns all companies.
func (cs *CompanyService) Companies(arq *CompanyRequest) (cl []*CompanyList, err error) {
	arq.Credentials = &Credentials{
		Username: os.Getenv(APIUsernameContainer),
		Password: os.Getenv(APIPasswordContainer),
	}

	req, err := cs.client.NewAPIRequest(http.MethodPost, "charterCompanies", arq)
	if err != nil {
		return
	}

	res, err := cs.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &cl); err != nil {
		return
	}

	return
}
