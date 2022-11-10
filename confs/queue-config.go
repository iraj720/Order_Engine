package confs

import "os"

// const (
// 	xxAcountingReq  = "xxAcountingReq"
// 	xxAcountingResp = "xxAcountingResp"
// 	consumerGroup   = "xxAccountingGrp0"
// 	consumerGroupId = "0"
// )

func GetXXRequestQueue() string {
	return os.Getenv("XX_REQUEST_QUEUE")
}

func GetXXResponseQueue() string {
	return os.Getenv("XX_RESPONSE_QUEUE")
}

func GetCryptoRequestQueue() string {
	return os.Getenv("CRYPTO_REQUEST_QUEUE")
}

func GetCryptoResponseQueue() string {
	return os.Getenv("CRYPTO_RESPONSE_QUEUE")
}
