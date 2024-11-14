package sfrxeth

import (
	"errors"
	"math/big"

	"github.com/KyberNetwork/blockchain-toolkit/number"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/entity"
	poolpkg "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/pool"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/util/bignumber"
	"github.com/goccy/go-json"
	"github.com/holiman/uint256"
)

type (
	PoolSimulator struct {
		poolpkg.Pool

		submitPaused bool
		totalSupply  *uint256.Int
		totalAssets  *uint256.Int

		gas Gas
	}
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrSubmitPaused = errors.New("submit is paused")
)

func NewPoolSimulator(entityPool entity.Pool) (*PoolSimulator, error) {
	var extra PoolExtra
	if err := json.Unmarshal([]byte(entityPool.Extra), &extra); err != nil {
		return nil, err
	}

	var (
		tokens   = make([]string, len(entityPool.Tokens))
		reserves = make([]*big.Int, len(entityPool.Tokens))
	)

	for i := 0; i < len(entityPool.Tokens); i++ {
		tokens[i] = entityPool.Tokens[i].Address
		reserves[i] = bignumber.NewBig10(entityPool.Reserves[i])
	}

	return &PoolSimulator{
		Pool: poolpkg.Pool{Info: poolpkg.PoolInfo{
			Address:     entityPool.Address,
			ReserveUsd:  entityPool.ReserveUsd,
			Exchange:    entityPool.Exchange,
			Type:        entityPool.Type,
			Tokens:      tokens,
			Reserves:    reserves,
			BlockNumber: entityPool.BlockNumber,
		}},
		submitPaused: extra.SubmitPaused,
		totalAssets:  uint256.MustFromDecimal(entityPool.Reserves[0]),
		totalSupply:  uint256.MustFromDecimal(entityPool.Reserves[1]),
		gas:          defaultGas,
	}, nil
}

func (s *PoolSimulator) CalcAmountOut(params poolpkg.CalcAmountOutParams) (*poolpkg.CalcAmountOutResult, error) {
	if s.submitPaused {
		return nil, ErrSubmitPaused
	}

	if params.TokenAmountIn.Token != s.Pool.Info.Tokens[0] || params.TokenOut != s.Pool.Info.Tokens[1] {
		return nil, ErrInvalidToken
	}

	amountOut, err := s.submitAndDeposit(uint256.MustFromBig(params.TokenAmountIn.Amount))
	if err != nil {
		return nil, err
	}

	return &poolpkg.CalcAmountOutResult{
		TokenAmountOut: &poolpkg.TokenAmount{Token: params.TokenOut, Amount: amountOut.ToBig()},
		Fee:            &poolpkg.TokenAmount{Token: params.TokenOut, Amount: bignumber.ZeroBI},
		Gas:            s.gas.SubmitAndDeposit,
	}, nil
}

func (s *PoolSimulator) UpdateBalance(params poolpkg.UpdateBalanceParams) {
	supply, overflow := uint256.FromBig(params.TokenAmountOut.Amount)
	if overflow {
		return
	}
	s.totalSupply.Add(s.totalSupply, supply)

	assets, overflow := uint256.FromBig(params.TokenAmountIn.Amount)
	if overflow {
		return
	}
	s.totalAssets.Add(s.totalAssets, assets)
}

func (s *PoolSimulator) GetMetaInfo(_ string, _ string) interface{} {
	return PoolMeta{
		BlockNumber: s.Pool.Info.BlockNumber,
	}
}

func (s *PoolSimulator) CanSwapTo(token string) []string {
	if token == s.Pool.Info.Tokens[1] {
		return []string{s.Pool.Info.Tokens[0]}
	}
	return []string{}
}

func (s *PoolSimulator) CanSwapFrom(token string) []string {
	if token == s.Pool.Info.Tokens[0] {
		return []string{s.Pool.Info.Tokens[1]}
	}
	return []string{}
}

func (s *PoolSimulator) submitAndDeposit(amountIn *uint256.Int) (*uint256.Int, error) {
	return s.previewDeposit(amountIn)
}

func (s *PoolSimulator) previewDeposit(assets *uint256.Int) (*uint256.Int, error) {
	if s.totalSupply.IsZero() {
		return assets.Clone(), nil
	}

	shares, overflow := new(uint256.Int).MulDivOverflow(assets, s.totalSupply, s.totalAssets)
	if overflow {
		return nil, number.ErrOverflow
	}
	return shares, nil
}