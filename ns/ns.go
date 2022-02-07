package ns

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"runtime"
	"strings"
)

// Nausys API global constants.
const (
	BaseURL              = "http://ws.nausys.com/CBMS-external/rest/"
	CatalogueURL         = "catalogue/v6"
	ReservationURL       = "yachtReservation/v6"
	RequestContentType   = "application/json"
	APIUsernameContainer = "NAUSYS_API_USERNAME"
	APIPasswordContainer = "NAUSYS_API_PASSWORD"
)

// Nausys API global errors.
var (
	errBadBaseURL = errors.New("malformed base url, it must contain a trailing slash")
)

// Client manages communication with Nausys API.
type Client struct {
	BaseURL      *url.URL
	userAgent    string
	client       *http.Client
	common       service // Reuse a single struct instead of allocating one for each service on the heap.
	Availability *AvailabilityService
	Offers       *OffersService
	Occupancy    *OccupancyService
	Company      *CompanyService
	Yacht        *YachtsService
}

// NewClient returns a new Nausys HTTP API client.
// You can pass a previously built http client, if none is provided then
// http.DefaultClient will be used.
func NewClient(baseClient *http.Client) (nausys *Client, err error) {
	if baseClient == nil {
		baseClient = http.DefaultClient
	}

	u, _ := url.Parse(BaseURL)

	nausys = &Client{
		BaseURL: u,
		client:  baseClient,
	}

	nausys.common.client = nausys

	// golang base user agent binding
	nausys.userAgent = strings.Join([]string{
		runtime.GOOS,
		runtime.GOARCH,
		runtime.Version(),
	}, ";")
	nausys.Availability = (*AvailabilityService)(&nausys.common)
	nausys.Offers = (*OffersService)(&nausys.common)
	nausys.Occupancy = (*OccupancyService)(&nausys.common)
	nausys.Company = (*CompanyService)(&nausys.common)
	nausys.Yacht = (*YachtsService)(&nausys.common)
	return
}

type service struct {
	client *Client
}

// NewAPIRequest is a wrapper around the http.NewRequest function.
func (c *Client) NewAPIRequest(method string, uri string, body interface{}) (req *http.Request, err error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, errBadBaseURL
	}

	u, err := c.BaseURL.Parse(uri)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err = http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", RequestContentType)
	req.Header.Set("Accept", RequestContentType)
	req.Header.Set("User-Agent", c.userAgent)

	return
}

// Do sends an API request and returns the API response or returned as an
// error if an API error has occurred.
func (c *Client) Do(req *http.Request) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	response, _ := newResponse(resp)
	err = CheckResponse(resp)
	if err != nil {
		return response, err
	}

	return response, nil
}

func newResponse(r *http.Response) (*Response, error) {
	var res Response
	c, err := ioutil.ReadAll(r.Body)
	if err == nil {
		res.content = c
	}
	err = json.NewDecoder(r.Body).Decode(&res)
	r.Body = io.NopCloser(bytes.NewBuffer(c))
	res.Response = r
	return &res, err
}

// CheckResponse checks the API response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range.
// API error responses are expected to have either no response
// body, or a JSON response body.
func CheckResponse(r *http.Response) error {
	if r.StatusCode >= http.StatusMultipleChoices {
		return newError(r)
	}
	return nil
}

/*
Constructor for Error
*/
func newError(r *http.Response) *Error {
	var e Error
	e.Response = r
	e.Code = r.StatusCode
	e.Message = r.Status
	c, err := ioutil.ReadAll(r.Body)
	if err == nil {
		e.Content = string(c)
	}
	r.Body = io.NopCloser(bytes.NewBuffer(c))
	return &e
}
