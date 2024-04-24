//go:generate go run github.com/tinylib/msgp -unexported -tests=false -v
//msgp:tuple PriceFeed
//msgp:ignore RoundData
//msgp:shim *big.Int as:[]byte using:msgpencode.EncodeInt/msgpencode.DecodeInt

package gmx

import (
	"math/big"

	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/util/bignumber"
)

type PriceFeed struct {
	RoundID *big.Int            `json:"roundId"`
	Answer  *big.Int            `json:"answer"`
	Answers map[string]*big.Int `json:"answers"`
}

type RoundData struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}

func NewPriceFeed() *PriceFeed {
	return &PriceFeed{
		Answers: make(map[string]*big.Int),
	}
}

const (
	priceFeedMethodLatestRoundData = "latestRoundData"
	priceFeedMethodGetRoundData    = "getRoundData"
	minRoundCount                  = 2
)

func (pf *PriceFeed) LatestRound() *big.Int {
	return pf.RoundID
}

func (pf *PriceFeed) LatestAnswer() *big.Int {
	return pf.Answer
}

// GetRoundData returns roundId, answer, startedAt, updatedAt, answeredInRound
func (pf *PriceFeed) GetRoundData(roundID *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int) {
	return roundID, pf.Answers[roundID.String()], bignumber.ZeroBI, bignumber.ZeroBI, bignumber.ZeroBI
}
