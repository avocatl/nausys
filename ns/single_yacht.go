package ns

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// YachtsService operates over company requests.
type YachtsService service

// Find retrieves a yacht with the yacht ID.
func (sys *YachtsService) Find(y int) (r YachtListResponse, err error) {
	cred := &Credentials{
		Username: os.Getenv(APIUsernameContainer),
		Password: os.Getenv(APIPasswordContainer),
	}

	target := fmt.Sprintf("%s/yacht/%d", CatalogueURL, y)

	req, err := sys.client.NewAPIRequest(http.MethodPost, target, cred)
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
