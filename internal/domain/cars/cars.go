package cars

type Car struct {
	CID       string
	Brand     string
	Model     string
	Color     string
	Year      int
	Number    string
	Price     float64
	Available bool
}

type Rent struct {
	RID   string
	CID   string
	UID   string
	Hours int
}

type AddCarRequest struct {
	Brand  string  `json:"brand" validate:"required"`
	Model  string  `json:"model" validate:"required"`
	Color  string  `json:"color" validate:"required"`
	Year   int     `json:"year" validate:"required"`
	Number string  `json:"number" validate:"required"`
	Price  float64 `json:"price" validate:"required"`
}

type StartRentRequest struct {
	CID   string `json:"cid"`
	Hours int    `json:"hours"`
}

type EndRentRequest struct {
	CID        string `json:"cid"`
	TotalHours int    `json:"totalHours"`
}
