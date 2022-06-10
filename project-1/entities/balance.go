package entities

type TopUp struct {
	ID          int
	User_id     int
	Telp        string
	Saldo_topup int
	Created_at  []uint8
}

type Transfer struct {
	ID            int
	Nama          string
	UserID        int
	TelpSender    string
	TelpReceiver  string
	SaldoTransfer int
	StatusID      int
	CreateAt      []uint8
}
