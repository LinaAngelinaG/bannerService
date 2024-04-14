package schedule

import "time"

const (
	CallbackTime        = time.Second * 20
	MempoolCheckTime    = time.Minute
	FinalizeTime        = time.Minute
	SendTransactionTime = 5 * time.Second
	//SendTransactionTime = 8*time.Minute
)
