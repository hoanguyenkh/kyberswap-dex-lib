package composablestable

import (
	"math/big"

	"github.com/goccy/go-json"
	"github.com/holiman/uint256"

	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/entity"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/balancer-v2/math"
	poolpkg "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/pool"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/util/bignumber"
)

type PoolSimulator struct {
	poolpkg.Pool

	Paused                 bool
	CanNotUpdateTokenRates bool

	RegularSimulator *RegularSimulator
	BptSimulator     *BptSimulator

	Vault       string
	PoolID      string
	PoolTypeVer int
}

func NewPoolSimulator(entityPool entity.Pool) (*PoolSimulator, error) {
	var (
		extra       Extra
		staticExtra StaticExtra

		tokens   = make([]string, len(entityPool.Tokens))
		reserves = make([]*big.Int, len(entityPool.Tokens))
	)

	if err := json.Unmarshal([]byte(entityPool.Extra), &extra); err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(entityPool.StaticExtra), &staticExtra); err != nil {
		return nil, err
	}

	for idx := 0; idx < len(entityPool.Tokens); idx++ {
		tokens[idx] = entityPool.Tokens[idx].Address
		reserves[idx] = bignumber.NewBig10(entityPool.Reserves[idx])
	}

	pool := poolpkg.Pool{
		Info: poolpkg.PoolInfo{
			Address:     entityPool.Address,
			Exchange:    entityPool.Exchange,
			Type:        entityPool.Type,
			Tokens:      tokens,
			Reserves:    reserves,
			Checked:     true,
			BlockNumber: entityPool.BlockNumber,
		},
	}

	regularSimulator := RegularSimulator{
		Pool:              pool,
		BptIndex:          staticExtra.BptIndex,
		ScalingFactors:    extra.ScalingFactors,
		Amp:               extra.Amp,
		SwapFeePercentage: extra.SwapFeePercentage,
	}

	bptSimulator := BptSimulator{
		Pool:                            pool,
		BptIndex:                        staticExtra.BptIndex,
		BptTotalSupply:                  extra.BptTotalSupply,
		Amp:                             extra.Amp,
		ScalingFactors:                  extra.ScalingFactors,
		LastJoinExit:                    extra.LastJoinExit,
		RateProviders:                   extra.RateProviders,
		TokenRateCaches:                 extra.TokenRateCaches,
		SwapFeePercentage:               extra.SwapFeePercentage,
		ProtocolFeePercentageCache:      extra.ProtocolFeePercentageCache,
		TokenExemptFromYieldProtocolFee: extra.IsTokenExemptFromYieldProtocolFee,
		ExemptFromYieldProtocolFee:      extra.IsExemptFromYieldProtocolFee,
		InRecoveryMode:                  extra.InRecoveryMode,

		PoolTypeVer: staticExtra.PoolTypeVer,
	}

	return &PoolSimulator{
		Pool:                   pool,
		Paused:                 extra.Paused,
		CanNotUpdateTokenRates: extra.CanNotUpdateTokenRates,
		RegularSimulator:       &regularSimulator,
		BptSimulator:           &bptSimulator,
		Vault:                  staticExtra.Vault,
		PoolID:                 staticExtra.PoolID,
		PoolTypeVer:            staticExtra.PoolTypeVer,
	}, nil
}

func (s *PoolSimulator) CalcAmountOut(params poolpkg.CalcAmountOutParams) (*poolpkg.CalcAmountOutResult, error) {
	if s.Paused {
		return nil, ErrPoolPaused
	}

	if s.CanNotUpdateTokenRates {
		return nil, ErrBeforeSwapJoinExit
	}

	tokenAmountIn := params.TokenAmountIn
	tokenOut := params.TokenOut

	indexIn := s.GetTokenIndex(tokenAmountIn.Token)
	indexOut := s.GetTokenIndex(tokenOut)
	if indexIn == unknownInt || indexOut == unknownInt {
		return nil, ErrUnknownToken
	}

	amountIn, overflow := uint256.FromBig(tokenAmountIn.Amount)
	if overflow {
		return nil, ErrOverflow
	}

	balances := make([]*uint256.Int, len(s.Info.Reserves))
	for i, reserve := range s.Info.Reserves {
		r, overflow := uint256.FromBig(reserve)
		if overflow {
			return nil, ErrOverflow
		}
		balances[i] = r
	}

	var (
		amountOut *uint256.Int
		fee       *poolpkg.TokenAmount
		swapInfo  *SwapInfo
		err       error
	)
	if tokenAmountIn.Token == s.Info.Address || tokenOut == s.Info.Address {
		amountOut, fee, swapInfo, err = s.BptSimulator.swap(amountIn, balances, indexIn, indexOut)
	} else {
		amountOut, fee, swapInfo, err = s.RegularSimulator.swap(amountIn, balances, indexIn, indexOut)
	}
	if err != nil {
		return nil, err
	}

	return &poolpkg.CalcAmountOutResult{
		TokenAmountOut: &poolpkg.TokenAmount{
			Token:  tokenOut,
			Amount: amountOut.ToBig(),
		},
		Fee:      fee,
		Gas:      DefaultGas.Swap,
		SwapInfo: swapInfo,
	}, nil
}

func (s *PoolSimulator) GetMetaInfo(tokenIn string, tokenOut string) interface{} {
	return PoolMetaInfo{
		Vault:         s.Vault,
		PoolID:        s.PoolID,
		TokenOutIndex: s.GetTokenIndex(tokenOut),
		BlockNumber:   s.Info.BlockNumber,
	}
}

func (s *PoolSimulator) UpdateBalance(params poolpkg.UpdateBalanceParams) {
	if params.TokenAmountIn.Token == s.Info.Address || params.TokenAmountOut.Token == s.Info.Address {
		s.BptSimulator.updateBalance(params)
		return
	}

	s.RegularSimulator.updateBalance(params)
}

func _downscaleDown(amount *uint256.Int, scalingFactor *uint256.Int) (*uint256.Int, error) {
	return math.FixedPoint.DivDown(amount, scalingFactor)
}

func _upscaleArray(balances []*uint256.Int, scalingFactors []*uint256.Int) ([]*uint256.Int, error) {
	upscaled := make([]*uint256.Int, len(balances))
	for i, balance := range balances {
		upscaledI, err := _upscale(balance, scalingFactors[i])
		if err != nil {
			return nil, err
		}
		upscaled[i] = upscaledI
	}
	return upscaled, nil
}

func _upscale(amount *uint256.Int, scalingFactor *uint256.Int) (*uint256.Int, error) {
	return math.FixedPoint.MulDown(amount, scalingFactor)
}

func _dropBptItem(amounts []*uint256.Int, bptIndex int) []*uint256.Int {
	amountsWithoutBpt := make([]*uint256.Int, len(amounts)-1)

	for i := 0; i < len(amountsWithoutBpt); i++ {
		if i < bptIndex {
			amountsWithoutBpt[i] = amounts[i]
			continue
		}
		amountsWithoutBpt[i] = amounts[i+1]
	}

	return amountsWithoutBpt
}

func _skipBptIndex(index int, bptIndex int) int {
	if index < bptIndex {
		return index
	}
	return index - 1
}
