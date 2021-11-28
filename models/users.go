package models

type User struct {
	Id       int      `json:"id"`
	UserName string   `json:"username"`
	Password string   `json:"password,omitempty"`
	Wallet   *Wallet  `gorm:"foreignKey:UserId;references:Id" json:"wallet"`
	Bitcoin  *Bitcoin `gorm:"foreignKey:UserId;references:Id" json:"bitcoin"`
}

type Wallet struct {
	WalletId int     `json:"walletId,omitempty"`
	UserId   int     `json:"userId"`
	Value    float64 `json:"value"`
}

type Bitcoin struct {
	BitcoinId int     `json:"bitcoinId,omitempty"`
	UserId    int     `json:"userId"`
	Amount    float64 `json:"amount"`
	Value     float64 `json:"value"`
}

type Signup struct {
	UserName        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type LoginCredentials struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}

type ValidateToken struct {
	Token string `json:"token"`
}
