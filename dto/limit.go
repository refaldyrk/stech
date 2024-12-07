package dto

type LimitResponse struct {
	ID            string  `json:"id"`
	KonsumenID    string  `json:"konsumen_id"`
	Tenor         int     `json:"tenor"`
	LimitPinjaman float64 `json:"limit_pinjaman"`

	Konsumen UserResponse `json:"konsumen,omitempty"`
}
