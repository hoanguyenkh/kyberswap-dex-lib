package dmm

import (
	"testing"

	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/entity"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/util/testutil"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func TestMsgpackMarshalUnmarshal(t *testing.T) {
	poolEntities := []entity.Pool{
		{
			Exchange:  "kyberswap",
			Type:      "dmm",
			Timestamp: 1685615099,
			Reserves:  entity.PoolReserves{"2766560101102", "1840989218168603319854"},
			Tokens: entity.PoolTokens{
				{
					Address:   "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
					Weight:    50,
					Swappable: true,
				},
				{
					Address:   "0xdd974d5c2e2928dea5f71b9825b8b646686bd200",
					Weight:    50,
					Swappable: true,
				},
			},
			Extra:        "{\"vReserves\":[\"867857435362478004\",\"2348002479022720085946\"],\"feeInPrecision\":\"1503833623506882\"}",
			ReserveUsd:   100000,
			AmplifiedTvl: 100000,
		},
	}
	var err error
	pools := make([]*PoolSimulator, len(poolEntities))
	for i, poolEntity := range poolEntities {
		pools[i], err = NewPoolSimulator(poolEntity)
		require.NoError(t, err)
	}
	for _, pool := range pools {
		b, err := pool.MarshalMsg(nil)
		require.NoError(t, err)
		actual := new(PoolSimulator)
		_, err = actual.UnmarshalMsg(b)
		require.NoError(t, err)
		require.Empty(t, cmp.Diff(pool, actual, testutil.CmpOpts(PoolSimulator{})...))
	}
}
