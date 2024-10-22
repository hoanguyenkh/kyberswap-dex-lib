package client

import (
	"context"
	"encoding/json"
	"errors"

	dexalot "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/dexalot"
	"github.com/KyberNetwork/logger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-resty/resty/v2"
)

const (
	pathQuote    = "api/rfq/firm"
	headerApiKey = "x-apikey"
)

var (
	ErrRFQFailed = errors.New("rfq failed")
)

type HTTPClient struct {
	config *dexalot.HTTPClientConfig
	client *resty.Client
}

func NewHTTPClient(config *dexalot.HTTPClientConfig) *HTTPClient {
	client := resty.New().
		SetBaseURL(config.BaseURL).
		SetTimeout(config.Timeout.Duration).
		SetRetryCount(config.RetryCount).
		SetHeader(headerApiKey, config.APIKey)

	return &HTTPClient{
		config: config,
		client: client,
	}
}

func (c *HTTPClient) Quote(ctx context.Context, params dexalot.FirmQuoteParams) (dexalot.FirmQuoteResult, error) {
	// token address case-sensitive
	req := c.client.R().
		SetContext(ctx).
		// the SellTokens address must follow the HEX format
		SetBody(map[string]interface{}{
			dexalot.ParamsChainID:     params.ChainID,
			dexalot.ParamsTakerAsset:  common.HexToAddress(params.TakerAsset).Hex(),
			dexalot.ParamsMakerAsset:  common.HexToAddress(params.MakerAsset).Hex(),
			dexalot.ParamsTakerAmount: params.TakerAmount,
			dexalot.ParamsUserAddress: params.UserAddress,
			dexalot.ParamsExecutor:    params.Executor,
		})
	var result dexalot.FirmQuoteResult
	var fail dexalot.FirmQuoteFail
	resp, err := req.SetResult(&result).SetError(&fail).Post(pathQuote)
	if err != nil {
		return dexalot.FirmQuoteResult{}, err
	}

	respBytes := resp.Body()
	_ = json.Unmarshal(respBytes, &result)
	_ = json.Unmarshal(respBytes, &fail)

	if !resp.IsSuccess() || fail.Failed() {
		return dexalot.FirmQuoteResult{}, ErrRFQFailed
	}

	return result, nil
}

func parseRFQError(errorCode int, message string) error {
	switch errorCode {
	default:
		logger.
			WithFields(logger.Fields{"client": "dexalot", "errorCode": errorCode, "message": message}).
			Error("rfq failed")
		return ErrRFQFailed
	}
}