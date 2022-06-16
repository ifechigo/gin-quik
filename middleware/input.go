package middleware

type CreateWalletInput struct {
	Firstname  string `json:"firstname" binding:"required"`
	Lastname string `json:"lastname" binding:"required"`
}

type UpdateWalletInput struct {
	Firstname  string `json:"firstname"`
	Lastname string `json:"lastname"`
}

type CreditWalletInput struct {
	Credit string `json:"amount"`
}

type DebitWalletInput struct {
	Debit string `json:"amount"`
}