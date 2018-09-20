package model

type Mnemonic struct {
	Value   string
}

type AddrJsonStructByInsight struct {
	AddrStr                 string  `json:"addrStr"`
	Balance                 float64 `json:"balance"`
	BalanceSat              int     `json:"balanceSat"`
	TotalReceived           float64 `json:"totalReceived"`
	TotalReceivedSat        int     `json:"totalReceivedSat"`
	TotalSent               int     `json:"totalSent"`
	TotalSentSat            int     `json:"totalSentSat"`
	UnconfirmedBalance      int     `json:"unconfirmedBalance"`
	UnconfirmedBalanceSat   int     `json:"unconfirmedBalanceSat"`
	UnconfirmedTxApperances int     `json:"unconfirmedTxApperances"`
	TxApperances            int     `json:"txApperances"`
}

type AutoGeneratedByBlockcypher struct {
	Address            string `json:"address"`
	TotalReceived      int    `json:"total_received"`
	TotalSent          int    `json:"total_sent"`
	Balance            int    `json:"balance"`
	UnconfirmedBalance int    `json:"unconfirmed_balance"`
	FinalBalance       int    `json:"final_balance"`
	NTx                int    `json:"n_tx"`
	UnconfirmedNTx     int    `json:"unconfirmed_n_tx"`
	FinalNTx           int    `json:"final_n_tx"`
}

type AddrJsonStruct struct {
	Hash160       string `json:"hash160"`
	Address       string `json:"address"`
	NTx           int    `json:"n_tx"`
	TotalReceived int    `json:"total_received"`
	TotalSent     int    `json:"total_sent"`
	FinalBalance  int    `json:"final_balance"`
	Txs           []struct {
		Ver    int `json:"ver"`
		Inputs []struct {
			Sequence int64  `json:"sequence"`
			Witness  string `json:"witness"`
			PrevOut  struct {
				Spent   bool   `json:"spent"`
				TxIndex int    `json:"tx_index"`
				Type    int    `json:"type"`
				Addr    string `json:"addr"`
				Value   int    `json:"value"`
				N       int    `json:"n"`
				Script  string `json:"script"`
			} `json:"prev_out"`
			Script string `json:"script"`
		} `json:"inputs"`
		Weight      int    `json:"weight"`
		BlockHeight int    `json:"block_height"`
		RelayedBy   string `json:"relayed_by"`
		Out         []struct {
			Spent   bool   `json:"spent"`
			TxIndex int    `json:"tx_index"`
			Type    int    `json:"type"`
			Addr    string `json:"addr"`
			Value   int    `json:"value"`
			N       int    `json:"n"`
			Script  string `json:"script"`
		} `json:"out"`
		LockTime int    `json:"lock_time"`
		Result   int    `json:"result"`
		Size     int    `json:"size"`
		Time     int    `json:"time"`
		TxIndex  int    `json:"tx_index"`
		VinSz    int    `json:"vin_sz"`
		Hash     string `json:"hash"`
		VoutSz   int    `json:"vout_sz"`
	} `json:"txs"`
}