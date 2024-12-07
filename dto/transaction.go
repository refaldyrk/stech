package dto

type TransaksiResponse struct {
	ID            string  `json:"id"`
	KonsumenID    string  `json:"konsumen_id"`
	NomorKontrak  string  `json:"nomor_kontrak"`
	NamaAsset     string  `json:"nama_asset"`
	OTR           float64 `json:"otr"`
	AdminFee      float64 `json:"admin_fee"`
	JumlahCicilan float64 `json:"jumlah_cicilan"`
	JumlahBunga   float64 `json:"jumlah_bunga"`
	CreatedAt     int64   `json:"created_at"`
}

type TransaksiRequestDTO struct {
	NomorKontrak  string  `json:"nomor_kontrak" binding:"required"`
	NamaAsset     string  `json:"nama_asset" binding:"required"`
	OTR           float64 `json:"otr" binding:"required"`
	AdminFee      float64 `json:"admin_fee" binding:"required"`
	JumlahCicilan float64 `json:"jumlah_cicilan" binding:"required"`
	JumlahBunga   float64 `json:"jumlah_bunga" binding:"required"`
}

type TransaksiResponseDTO struct {
	TransactionID  string `json:"transaction_id"`
	Tenor          int    `json:"tenor"`
	LimitAvailable any    `json:"limit_available"`
}
