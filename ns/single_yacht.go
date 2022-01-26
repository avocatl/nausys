package ns

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type SingleYachtService service

func (sys *SingleYachtService) GetYacht(y int) (r YachtListResponse, err error) {
	cred := &Credentials{
		Username: os.Getenv(APIUsernameContainer),
		Password: os.Getenv(APIPasswordContainer),
	}

	target := fmt.Sprintf("%s/yacht/%d", CatalogueURL, y)

	body, err := json.Marshal(cred)
	if err != nil {
		return
	}

	req, err := sys.client.NewAPIRequest(http.MethodPost, target, body)
	if err != nil {
		return
	}

	res, err := sys.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &r); err != nil {
		return
	}

	return
}
