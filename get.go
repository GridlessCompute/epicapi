package epicapi

import (
	"encoding/json"
	"fmt"
)

// Capabilities returns a list of the device's capabilities
func (g *get) Capabilities() (CapabilitiesResult, error) {
	var epres CapabilitiesResult

	res, err := g.api.GET("capabilities")
	if err != nil {
		return CapabilitiesResult{}, fmt.Errorf("%w", err)
	}

	if err := json.Unmarshal(res, &epres); err != nil {
		return CapabilitiesResult{}, fmt.Errorf("error Unmarshaling res body for capabilities (%w)", err)
	}

	return epres, nil
}

// Clocks returns the clock speeds of the device
func (g *get) Clocks() (ClocksResult, error) {
	var clcks ClocksResult

	res, err := g.api.GET("clocks")
	if err != nil {
		return ClocksResult{}, fmt.Errorf("%w", err)
	}

	if err := json.Unmarshal(res, &clcks); err != nil {
		return ClocksResult{}, fmt.Errorf("error Unmarshaling res body for clocks (%w)", err)
	}

	return clcks, nil
}

// ClocksConfig returns the clock configuration of the device
func (g *get) ClocksConfig() (ClocksConfigResult, error) {
	var clcksCnfg ClocksConfigResult

	res, err := g.api.GET("clocks/config")
	if err != nil {
		return ClocksConfigResult{}, fmt.Errorf("%w", err)
	}

	if err := json.Unmarshal(res, &clcksCnfg); err != nil {
		return ClocksConfigResult{}, fmt.Errorf("error Unmarshaling res body for clocks/config (%w)", err)
	}

	return clcksCnfg, nil
}

// Hashrate returns the hashrate of the device
func (g *get) Hashrate() (HashrateResult, error) {
	var hsrt HashrateResult

	res, err := g.api.GET("hashrate")
	if err != nil {
		return HashrateResult{}, fmt.Errorf("%w", err)
	}

	if err := json.Unmarshal(res, &hsrt); err != nil {
		return HashrateResult{}, fmt.Errorf("error Unmarshaling res body for hashrate (%w)", err)
	}

	return hsrt, nil
}

// HashrateHistoryContinuous returns the continuous hashrate history of the device
func (g *get) HashrateHistoryContinuous() (HashrateHistoryContinuousResult, error) {
	var hsrtHstryCnts HashrateHistoryContinuousResult

	res, err := g.api.GET("hashrate/history/continuous")
	if err != nil {
		return HashrateHistoryContinuousResult{}, fmt.Errorf("%w", err)
	}

	if err := json.Unmarshal(res, &hsrtHstryCnts); err != nil {
		return HashrateHistoryContinuousResult{}, fmt.Errorf("error Unmarshaling res body for hashrate/history/continuous (%w)", err)
	}

	return hsrtHstryCnts, nil
}

// HashrateHistoryDiscrete returns the discrete hashrate history of the device
func (g *get) HashrateHistoryDiscrete() (HashrateHistoryDiscreteResult, error) {
	var hsrtHstryDscrt HashrateHistoryDiscreteResult

	res, err := g.api.GET("hashrate/history/discrete")
	if err != nil {
		return HashrateHistoryDiscreteResult{}, fmt.Errorf("%w", err)
	}

	if err := json.Unmarshal(res, &hsrtHstryDscrt); err != nil {
		return HashrateHistoryDiscreteResult{}, fmt.Errorf("error Unmarshaling res body for hashrate/history/discrete (%w)", err)
	}

	return hsrtHstryDscrt, nil
}

// History returns the history of the device
func (g *get) History() (HistoryResult, error) {
	var hstry HistoryResult

	res, err := g.api.GET("history")
	if err != nil {
		return HistoryResult{}, fmt.Errorf("%w", err)
	}

	if err := json.Unmarshal(res, &hstry); err != nil {
		return HistoryResult{}, fmt.Errorf("error Unmarshaling res body for history (%w)", err)
	}

	return hstry, nil
}

/* func (g *Get) Log() (LogResult, error) {
	var lg LogResult

	res, err := g.api.GET("log")
	if err != nil {
		return LogResult{}, fmt.Errorf("%w", err)
	}

	if err := json.Unmarshal(res, &lg); err != nil {
		return LogResult{}, fmt.Errorf("error Unmarshaling res body for log (%w)", err)
	}

	return lg, nil
} */

// Network returns the network settings of the device
func (g *get) Network() (NetworkResult, error) {
	var ntwrk NetworkResult

	res, err := g.api.GET("network")
	if err != nil {
		return NetworkResult{}, fmt.Errorf("%w", err)
	}

	if err := json.Unmarshal(res, &ntwrk); err != nil {
		return NetworkResult{}, fmt.Errorf("error Unmarshaling res body for network (%w)", err)
	}

	return ntwrk, nil
}

// PerpetualTune returns the perpetual tune settings of the device
func (g *get) PerpetualTune() (PerpetualTuneResult, error) {
	var prptlTn PerpetualTuneResult

	res, err := g.api.GET("perpetual-tune")
	if err != nil {
		return PerpetualTuneResult{}, fmt.Errorf("%w", err)
	}

	if err := json.Unmarshal(res, &prptlTn); err != nil {
		return PerpetualTuneResult{}, fmt.Errorf("error Unmarshaling res body for perpetual-tune (%w)", err)
	}

	return prptlTn, nil
}

// Summary returns a summary of the device
func (g *get) Summary() (SummaryResult, error) {
	var smry SummaryResult

	res, err := g.api.GET("summary")
	if err != nil {
		return SummaryResult{}, fmt.Errorf("%w", err)
	}

	if err := json.Unmarshal(res, &smry); err != nil {
		return SummaryResult{}, fmt.Errorf("error Unmarshaling res body for summary (%w)", err)
	}

	return smry, nil
}

// Temps returns the temperatures of the device
func (g *get) Temps() (TempsResult, error) {
	var tmps TempsResult

	res, err := g.api.GET("temps")
	if err != nil {
		return TempsResult{}, fmt.Errorf("%w", err)
	}

	if err := json.Unmarshal(res, &tmps); err != nil {
		return TempsResult{}, fmt.Errorf("error Unmarshaling res body for temps (%w)", err)
	}

	return tmps, nil
}

// TempsBoard returns the board temperatures of the device
func (g *get) TempsBoard() (TempsBoardResult, error) {
	var tmpsBrd TempsBoardResult

	res, err := g.api.GET("temps/board")
	if err != nil {
		return TempsBoardResult{}, fmt.Errorf("%w", err)
	}

	if err := json.Unmarshal(res, &tmpsBrd); err != nil {
		return TempsBoardResult{}, fmt.Errorf("error Unmarshaling res body for temps/board (%w)", err)
	}

	return tmpsBrd, nil
}

// TempsChip returns the chip temperatures of the device
func (g *get) TempsChip() (TempsChipResult, error) {
	var tmpsChp TempsChipResult

	res, err := g.api.GET("temps/chip")
	if err != nil {
		return TempsChipResult{}, fmt.Errorf("%w", err)
	}

	if err := json.Unmarshal(res, &tmpsChp); err != nil {
		return TempsChipResult{}, fmt.Errorf("error Unmarshaling res body for temps/chip (%w)", err)
	}

	return tmpsChp, nil
}

// Voltages returns the voltages of the device
func (g *get) Voltages() (VoltageResult, error) {
	var vltgs VoltageResult

	res, err := g.api.GET("voltages")
	if err != nil {
		return VoltageResult{}, fmt.Errorf("%w", err)
	}

	if err := json.Unmarshal(res, &vltgs); err != nil {
		return VoltageResult{}, fmt.Errorf("error Unmarshaling res body for voltages (%w)", err)
	}

	return vltgs, nil
}
