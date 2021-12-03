package models

// Block data
type Block struct {
	BlockNumber       int64         `json:"block_number"`
	Timestamp         uint64        `json:"timestamp"`
	Difficulty        uint64        `json:"difficulty"`
	Hash              string        `json:"hash"`
	TransactionsCount int           `json:"transactions_count"`
	Transactions      []Transaction `json:"transactions"`
}

type Transaction struct {
	Hash     string `json:"hash"`
	Value    string `json:"value"`
	Gas      uint64 `json:"gas"`
	GasPrice uint64 `json:"gas_price"`
	Nonce    uint64 `json:"nonce"`
	To       string `json:"to"`
	Pending  bool   `json:"pending"`
}

// TransferRequest
type TransferRequest struct {
	PrivKey string `json:"priv_key"`
	To      string `json:"to"`
	Amount  int64  `json:"amount"`
}

// HashResponse
type HashResponse struct {
	Hash string `json:"hash"`
}

// BlanceResponse
type BlanceResponse struct {
	Address string `json:"address"`
	Blance  string `json:"blance"`
	Symbol  string `json:"symbol"`
	Units   string `json:"units"`
}

// Error
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
