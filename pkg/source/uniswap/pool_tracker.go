package uniswap

import (
	"context"
	"time"

	"github.com/KyberNetwork/ethrpc"
	"github.com/KyberNetwork/logger"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/entity"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/pool"
)

type PoolTracker struct {
	ethrpcClient *ethrpc.Client
}

func NewPoolTracker(
	ethrpcClient *ethrpc.Client,
) (*PoolTracker, error) {
	return &PoolTracker{
		ethrpcClient: ethrpcClient,
	}, nil
}

func (d *PoolTracker) GetNewPoolState(
	ctx context.Context,
	p entity.Pool,
	params pool.GetNewPoolStateParams,
) (entity.Pool, error) {
	logger.Infof("[Uniswap V2] Start getting new state of pool: %v", p.Address)

	var (
		err      error
		reserves Reserves
	)

	latestSyncEvent := d.findLatestSyncEvent(params.Logs)
	if latestSyncEvent == nil {
		reserves, err = d.fetchReservesFromNode(ctx, p.Address)
		if err != nil {
			logger.WithFields(logger.Fields{
				"poolAddress": p.Address,
				"error":       err,
			}).Error("Fail to fetch reserves from node")
		}
	} else {
		reserves, err = decodeSyncEvent(*latestSyncEvent)
		if err != nil {
			logger.WithFields(logger.Fields{
				"poolAddress": p.Address,
				"event":       latestSyncEvent,
				"error":       err,
			}).Error("Fail to decode sync event")
		}
	}

	if err != nil {
		return entity.Pool{}, err
	}

	p.Timestamp = time.Now().Unix()
	p.Reserves = entity.PoolReserves{
		reserves.Reserve0.String(),
		reserves.Reserve1.String(),
	}

	logger.Infof("[Uniswap V2] Finish getting new state of pool: %v", p.Address)

	return p, nil
}

func (d *PoolTracker) fetchReservesFromNode(ctx context.Context, poolAddress string) (Reserves, error) {
	var reserves Reserves

	rpcRequest := d.ethrpcClient.NewRequest()
	rpcRequest.SetContext(ctx)

	rpcRequest.AddCall(&ethrpc.Call{
		ABI:    uniswapV2PairABI,
		Target: poolAddress,
		Method: pairMethodGetReserves,
		Params: nil,
	}, []interface{}{&reserves})

	_, err := rpcRequest.Call()
	if err != nil {
		logger.Errorf("failed to process tryAggregate for pool: %v, err: %v", poolAddress, err)
		return Reserves{}, err
	}

	return reserves, nil
}

func (d *PoolTracker) findLatestSyncEvent(logs []types.Log) *types.Log {
	var latestSyncEvent *types.Log
	for _, log := range logs {
		if log.Removed || !isSyncEvent(log) {
			continue
		}

		if latestSyncEvent == nil ||
			latestSyncEvent.BlockNumber < log.BlockNumber ||
			(latestSyncEvent.BlockNumber == log.BlockNumber && latestSyncEvent.Index < log.Index) {
			latestSyncEvent = &log
		}
	}

	return latestSyncEvent
}

func isSyncEvent(log types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}

	return log.Topics[0] == uniswapV2PairABI.Events["Sync"].ID
}

func decodeSyncEvent(log types.Log) (Reserves, error) {
	filterer, err := NewUniswapFilterer(log.Address, nil)
	if err != nil {
		return Reserves{}, err
	}

	syncEvent, err := filterer.ParseSync(log)
	if err != nil {
		return Reserves{}, err
	}

	return Reserves{
		Reserve0: syncEvent.Reserve0,
		Reserve1: syncEvent.Reserve1,
	}, nil
}
