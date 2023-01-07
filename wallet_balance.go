package iiko

type WalletType int

const (
	WalletTypeDepositOrCorporateNutrition WalletType = 0
	WalletTypeBonusProgram                WalletType = 1
	WalletTypeProductsProgram             WalletType = 2
	WalletTypeDiscountProgram             WalletType = 3
	WalletTypeCertificateProgram          WalletType = 4
)

type WalletBalance struct {
	Id      string     `json:"id"`
	Name    string     `json:"name"`
	Type    WalletType `json:"type"`
	Balance float64    `json:"balance"`
}
