package epicapi

import (
	"fmt"
	"strconv"
)

type PostRes struct {
	Result bool   `json:"result"`
	Error  string `json:"error"`
}

func (p *Post) Authenticate() error {
	payload := map[string]any{
		"param":    nil,
		"password": p.api.Pass,
	}

	if err := p.api.POST("authenticate", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (p *Post) BoardEnable(idx []int, enable bool) error {
	payload := map[string]any{
		"param":    []map[string]any{},
		"password": p.api.Pass,
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

	if err := p.api.POST("boardenable", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (p *Post) ClearHashrate() error {
	payload := map[string]any{
		"param":    nil,
		"password": p.api.Pass,
	}

	if err := p.api.POST("clearhashrate", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

type StratumConfig struct {
	Login    string
	Password string
	Pool     string
}

func (p *Post) Coin(confs []StratumConfig, coin string, uniqueId bool) error {
	payload := map[string]any{
		"param": map[string]any{
			"coin":            coin,
			"stratum_configs": confs,
			"unique_id":       uniqueId,
		},
		"password": p.api.Pass,
	}

	if err := p.api.POST("coin", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (p *Post) CriticalTemp(temp int) error {
	payload := map[string]any{
		"param":    temp,
		"password": p.api.Pass,
	}

	if err := p.api.POST("criticaltemp", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (p *Post) DefaultConfig() error {
	payload := map[string]any{
		"param":    nil,
		"password": p.api.Pass,
	}

	if err := p.api.POST("defaultconfig", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (p *Post) FanSpeed(idleSpeed, targetTemp int) error {
	payload := map[string]any{
		"param": map[string]any{
			"Auto": map[string]any{
				"Idle Speed":         idleSpeed,
				"Target Temperature": targetTemp,
			},
		},
		"password": p.api.Pass,
	}

	if err := p.api.POST("fanspeed", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (p *Post) Id(unique bool) error {
	payload := map[string]any{
		"param":    unique,
		"password": p.api.Pass,
	}

	if err := p.api.POST("id", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (p *Post) IdVarient(varient int) error {
	opt := []string{"IpAddress", "MacAddress", "CpuId"}

	if !(varient > 0) && !(varient < 3) {
		return fmt.Errorf("Bad varient index")
	}

	payload := map[string]any{
		"param":    opt[varient],
		"password": p.api.Pass,
	}

	if err := p.api.POST("id/varient", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (p *Post) Identify(identify bool) error {
	payload := map[string]any{
		"param":    identify,
		"password": p.api.Pass,
	}

	if err := p.api.POST("identify", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (p *Post) IdleOnConnectionLost(idle bool) error {
	payload := map[string]any{
		"param":    idle,
		"password": p.api.Pass,
	}

	if err := p.api.POST("idleonconnectionlost", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (p *Post) Miner(control bool) error {
	var payload map[string]any

	if control {
		payload = map[string]any{
			"param":    "Autostart",
			"password": p.api.Pass,
		}
	} else {
		payload = map[string]any{
			"param":    "Stop",
			"password": p.api.Pass,
		}
	}

	if err := p.api.POST("miner", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (p *Post) Network(dhcp bool) error {
	payload := map[string]any{
		"param": map[string]any{
			"dhcp": dhcp,
		},
		"password": p.api.Pass,
	}

	if err := p.api.POST("network", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (p *Post) Overdrive(od bool) error {
	payload := map[string]any{
		"param":    od,
		"password": p.api.Pass,
	}

	if err := p.api.POST("overdrive", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (p *Post) Password(pass string) error {
	payload := map[string]any{
		"param":    pass,
		"password": p.api.Pass,
	}

	if err := p.api.POST("password", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	p.api.Pass = pass

	return nil
}

func (p *Post) PerpetualTune(pt bool) error {
	payload := map[string]any{
		"param":    pt,
		"password": p.api.Pass,
	}

	if err := p.api.POST("perpetualtune", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (p *Post) Reboot(delay int) error {
	payload := map[string]any{
		"param":    delay,
		"password": p.api.Pass,
	}

	if err := p.api.POST("reboot", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (p *Post) ShutdownTemp(temp int) error {
	payload := map[string]any{
		"param":    temp,
		"password": p.api.Pass,
	}

	if err := p.api.POST("shutdowntemp", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (p *Post) SoftReboot() error {
	payload := map[string]any{
		"param":    nil,
		"password": p.api.Pass,
	}

	if err := p.api.POST("softreboot", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (p *Post) Tune(freq, volt float64) error {
	volt = volt * 1000
	payload := map[string]any{
		"param": map[string]string{
			"clks":    strconv.FormatFloat(freq, 'f', 2, 64),
			"voltage": strconv.FormatFloat(volt, 'f', 0, 64),
		},
		"password": p.api.Pass,
	}

	if err := p.api.POST("tune", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (p *Post) TuneClockAll(freq float64) error {
	payload := map[string]any{
		"param":    freq,
		"password": p.api.Pass,
	}

	if err := p.api.POST("tune/clock/all", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

type BoardTune struct {
	Data  float64
	Index int
}

func (p *Post) TuneClockBoard(settings []BoardTune) error {
	payload := map[string]any{
		"param": map[string]any{
			"v": settings,
		},
		"password": p.api.Pass,
	}

	if err := p.api.POST("tune/clock/board", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (p *Post) TuneVoltage(volt float64) error {
	payload := map[string]any{
		"param":    int(volt * 1000),
		"password": p.api.Pass,
	}

	if err := p.api.POST("tune/voltage", payload); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
