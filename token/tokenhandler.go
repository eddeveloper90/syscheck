package token

const randomLength = 5

type GtToken struct {
	ClientId       int64
	ClientName     int64
	RegexCheck     bool
	IpCheck        bool
	TokenId        int64
	ExpirationUnix int64
	RandomStr      string
	Uuid           string
	Ip             string
	CheckSum       string
	ExpireDateTime string
}

func (gt *GtToken) loadDbTokenTokenId (dbToken string, tokenId int64) {
	gt.TokenId = tokenId

}

func (gt *GtToken) loadToken(token string) {

}

func (gt *GtToken) loadUnixIp(expirationUnix int64, ip string) {

}
