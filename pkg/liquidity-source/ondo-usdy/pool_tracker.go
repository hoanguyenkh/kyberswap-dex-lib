package ondo_usdy

import (
	"context"
	"time"

	"github.com/KyberNetwork/ethrpc"
	"github.com/KyberNetwork/logger"
	"github.com/goccy/go-json"

	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/entity"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/pool"
)

type PoolTracker struct {
	config       *Config
	ethrpcClient *ethrpc.Client
}

func NewPoolTracker(config *Config, ethrpcClient *ethrpc.Client) (*PoolTracker, error) {
	return &PoolTracker{
		config:       config,
		ethrpcClient: ethrpcClient,
	}, nil
}

func (t *PoolTracker) GetNewPoolState(
	ctx context.Context,
	p entity.Pool,
	_ pool.GetNewPoolStateParams,
) (entity.Pool, error) {
	logger.WithFields(logger.Fields{
		"dex_id": p.Exchange,
		"pool":   p.Address,
	}).Info("start fetching pool state")

	oldExtra := PoolExtra{}
	err := json.Unmarshal([]byte(p.Extra), &oldExtra)
	if err != nil {

		return entity.Pool{}, err
	}

	newExtras, blockNumber, err := getExtra(
		ctx, t.config, t.ethrpcClient, []entity.Pool{p}, []string{oldExtra.RWADynamicOracleAddress},
	)
	if err != nil {
		return p, err
	}

	extraBytes, err := json.Marshal(newExtras[0])
	if err != nil {
		return p, err
	}

	p.Extra = string(extraBytes)
	p.BlockNumber = blockNumber
	p.Timestamp = time.Now().Unix()

	return p, nil
}
