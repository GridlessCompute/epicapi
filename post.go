package epicapi

import (
	"fmt"
	"strconv"
)

type PostRes struct {
	Result bool   `json:"result"`
	Error  string `json:"error"`
}

// Authenticate authenticates with the device
func (p *Post) Authenticate() error {
	payload := map[string]any{
		"param":    nil,
		"password": p.api.password,
	}

	if err := p.api.post("authenticate", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// BoardEnable enables or disables a board
// BoardEnable(idx []int, enable bool)
func (p *Post) BoardEnable(idx []int, enable bool) error {
	payload := map[string]any{
		"param":    []map[string]any{},
		"password": p.api.password,
	}

	for _, id := range idx {
		b := map[string]any{
			"Data":  enable,
			"Index": id,
		}

		pr := payload["param"].([]map[string]any)
		pr = append(pr, b)

		payload["param"] = pr
	}

	if err := p.api.post("boardenable", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// ClearHashrate clears the hashrate history of the device
func (p *Post) ClearHashrate() error {
	payload := map[string]any{
		"param":    nil,
		"password": p.api.password,
	}

	if err := p.api.post("clearhashrate", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

type StratumConfig struct {
	Login    string
	Password string
	Pool     string
}

// Coin sets the coin to mine
// Coin(confs []StratumConfig, coin string, uniqueId bool)
func (p *Post) Coin(confs []StratumConfig, coin string, uniqueId bool) error {
	payload := map[string]any{
		"param": map[string]any{
			"coin":            coin,
			"stratum_configs": confs,
			"unique_id":       uniqueId,
		},
		"password": p.api.password,
	}

	if err := p.api.post("coin", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// CriticalTemp sets the critical temperature of the device
// CriticalTemp(temp int)
func (p *Post) CriticalTemp(temp int) error {
	payload := map[string]any{
		"param":    temp,
		"password": p.api.password,
	}

	if err := p.api.post("criticaltemp", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// DefaultConfig resets the device to its default configuration
func (p *Post) DefaultConfig() error {
	payload := map[string]any{
		"param":    nil,
		"password": p.api.password,
	}

	if err := p.api.post("defaultconfig", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// FanSpeed sets the fan speed of the device
// FanSpeed(idleSpeed, targetTemp int)
func (p *Post) FanSpeed(idleSpeed, targetTemp int) error {
	payload := map[string]any{
		"param": map[string]any{
			"Auto": map[string]any{
				"Idle Speed":         idleSpeed,
				"Target Temperature": targetTemp,
			},
		},
		"password": p.api.password,
	}

	if err := p.api.post("fanspeed", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// Id sets the device ID
// Id(unique bool)
func (p *Post) Id(unique bool) error {
	payload := map[string]any{
		"param":    unique,
		"password": p.api.password,
	}

	if err := p.api.post("id", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// IdVarient sets the device ID varient
// IdVarient(varient int)
func (p *Post) IdVarient(varient int) error {
	opt := []string{"IpAddress", "MacAddress", "CpuId"}

	if !(varient > 0) && !(varient < 3) {
		return fmt.Errorf("bad varient index")
	}

	payload := map[string]any{
		"param":    opt[varient],
		"password": p.api.password,
	}

	if err := p.api.post("id/varient", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// Identify identifies the device
// Identify(identify bool)
func (p *Post) Identify(identify bool) error {
	payload := map[string]any{
		"param":    identify,
		"password": p.api.password,
	}

	if err := p.api.post("identify", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// IdleOnConnectionLost sets the device to idle when the connection is lost
// IdleOnConnectionLost(idle bool)
func (p *Post) IdleOnConnectionLost(idle bool) error {
	payload := map[string]any{
		"param":    idle,
		"password": p.api.password,
	}

	if err := p.api.post("idleonconnectionlost", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// Miner starts or stops the miner
// Miner(control bool)
func (p *Post) Miner(control bool) error {
	var payload map[string]any

	if control {
		payload = map[string]any{
			"param":    "Autostart",
			"password": p.api.password,
		}
	} else {
		payload = map[string]any{
			"param":    "Stop",
			"password": p.api.password,
		}
	}

	if err := p.api.post("miner", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// Network sets the network settings of the device
// Network(dhcp bool)
func (p *Post) Network(dhcp bool) error {
	payload := map[string]any{
		"param": map[string]any{
			"dhcp": dhcp,
		},
		"password": p.api.password,
	}

	if err := p.api.post("network", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// Overdrive enables or disables overdrive
// Overdrive(od bool)
func (p *Post) Overdrive(od bool) error {
	payload := map[string]any{
		"param":    od,
		"password": p.api.password,
	}

	if err := p.api.post("overdrive", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// Password sets the password of the device
// Password(pass string)
func (p *Post) Password(pass string) error {
	payload := map[string]any{
		"param":    pass,
		"password": p.api.password,
	}

	if err := p.api.post("password", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	p.api.password = pass

	return nil
}

// PerpetualTune enables or disables perpetual tune
// PerpetualTune(pt bool)
func (p *Post) PerpetualTune(pt bool) error {
	payload := map[string]any{
		"param":    pt,
		"password": p.api.password,
	}

	if err := p.api.post("perpetualtune", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// Reboot reboots the device
// Reboot(delay int)
func (p *Post) Reboot(delay int) error {
	payload := map[string]any{
		"param":    delay,
		"password": p.api.password,
	}

	if err := p.api.post("reboot", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// ShutdownTemp sets the shutdown temperature of the device
// ShutdownTemp(temp int)
func (p *Post) ShutdownTemp(temp int) error {
	payload := map[string]any{
		"param":    temp,
		"password": p.api.password,
	}

	if err := p.api.post("shutdowntemp", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// SoftReboot soft reboots the device
func (p *Post) SoftReboot() error {
	payload := map[string]any{
		"param":    nil,
		"password": p.api.password,
	}

	if err := p.api.post("softreboot", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// Tune tunes the device
// Tune(freq, volt float64)
func (p *Post) Tune(freq, volt float64) error {
	volt = volt * 1000
	payload := map[string]any{
		"param": map[string]string{
			"clks":    strconv.FormatFloat(freq, 'f', 2, 64),
			"voltage": strconv.FormatFloat(volt, 'f', 0, 64),
		},
		"password": p.api.password,
	}

	if err := p.api.post("tune", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// TuneClockAll tunes the clock of all boards
// TuneClockAll(freq float64)
func (p *Post) TuneClockAll(freq float64) error {
	payload := map[string]any{
		"param":    freq,
		"password": p.api.password,
	}

	if err := p.api.post("tune/clock/all", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

type BoardTune struct {
	Data  float64
	Index int
}

// TuneClockBoard tunes the clock of a specific board
// TuneClockBoard(settings []BoardTune)
func (p *Post) TuneClockBoard(settings []BoardTune) error {
	payload := map[string]any{
		"param": map[string]any{
			"v": settings,
		},
		"password": p.api.password,
	}

	if err := p.api.post("tune/clock/board", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// TuneVoltage tunes the voltage of the device
// TuneVoltage(volt float64)
func (p *Post) TuneVoltage(volt float64) error {
	payload := map[string]any{
		"param":    int(volt * 1000),
		"password": p.api.password,
	}

	if err := p.api.post("tune/voltage", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
