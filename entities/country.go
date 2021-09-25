package entities

type Country struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Currency    Currency `json:"currency"`
	CallingCode string   `json:"callingCode"`
}
