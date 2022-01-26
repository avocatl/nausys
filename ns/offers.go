package ns

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type OffersService service

func (osrv *OffersService) GetOffers(orq *FreeYachtRequest) (or *FreeYachtListResponse, err error) {
	orq.Credentials = &Credentials{
		Username: os.Getenv(APIUsernameContainer),
		Password: os.Getenv(APIPasswordContainer),
	}

	target := fmt.Sprintf("%s/freeYachts", ReservationURL)

	req, err := osrv.client.NewAPIRequest(http.MethodPost, target, orq)
	if err != nil {
		return
	}

	res, err := osrv.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &or); err != nil {
		return
	}

	return
}
