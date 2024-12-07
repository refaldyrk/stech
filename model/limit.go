package model

import "kreditplus-test/dto"

type Limit struct {
	ID            string  `db:"id"`
	KonsumenID    string  `db:"konsumen_id"`
	Tenor         int     `db:"tenor"`
	LimitPinjaman float64 `db:"limit_pinjaman"`

	Konsumen Customer `db:"konsumen,omitempty"`
}

func (l Limit) ToDTO() dto.LimitResponse {
	return dto.LimitResponse{
		ID:            l.ID,
		KonsumenID:    l.KonsumenID,
		Tenor:         l.Tenor,
		LimitPinjaman: l.LimitPinjaman,
		Konsumen:      l.Konsumen.ToDTO(),
	}
}
