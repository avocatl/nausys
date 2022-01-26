package ns

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// CompanyService operates over company requests.
type CompanyService service

// All returns all companies.
func (cs *CompanyService) All() (clr *CompanyListResponse, err error) {
	c := &Credentials{
		Username: os.Getenv(APIUsernameContainer),
		Password: os.Getenv(APIPasswordContainer),
	}

	target := fmt.Sprintf("%s/charterCompanies", CatalogueURL)

	req, err := cs.client.NewAPIRequest(http.MethodPost, target, c)
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
