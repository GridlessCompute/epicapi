# epicapi
ePIC API wrapper for Go

## Usage
```go
import "github.com/GridlessCompute/epicapi"

func main() {
	// Creates a new API client
	// New(ip, port, password string)
	api := epicapi.New("192.168.1.1", "80", "password")
}
```

## Functions

### GET
```go
// Capabilities returns a list of the device's capabilities
func (g *Get) Capabilities() (CapabilitiesResult, error)
```
```go
// Clocks returns the clock speeds of the device
func (g *Get) Clocks() (ClocksResult, error)
```
```go
// ClocksConfig returns the clock configuration of the device
func (g *Get) ClocksConfig() (ClocksConfigResult, error)
```
```go
// Hashrate returns the hashrate of the device
func (g *Get) Hashrate() (HashrateResult, error)
```
```go
// HashrateHistoryContinuous returns the continuous hashrate history of the device
func (g *Get) HashrateHistoryContinuous() (HashrateHistoryContinuousResult, error)
```
```go
// HashrateHistoryDiscrete returns the discrete hashrate history of the device
func (g. *Get) HashrateHistoryDiscrete() (HashrateHistoryDiscreteResult, error)
```
```go
// History returns the history of the device
func (g *Get) History() (HistoryResult, error)
```
```go
// Network returns the network settings of the device
func (g *Get) Network() (NetworkResult, error)
```
```go
// PerpetualTune returns the perpetual tune settings of the device
func (g *Get) PerpetualTune() (PerpetualTuneResult, error)
```
```go
// Summary returns a summary of the device
func (g *Get) Summary() (SummaryResult, error)
```
```go
// Temps returns the temperatures of the device
func (g *Get) Temps() (TempsResult, error)
```
```go
// TempsBoard returns the board temperatures of the device
func (g *Get) TempsBoard() (TempsBoardResult, error)
```
```go
// TempsChip returns the chip temperatures of the device
func (g *Get) TempsChip() (TempsChipResult, error)
```
```go
// Voltages returns the voltages of the device
func (g *Get) Voltages() (VoltageResult, error)
```

### POST
```go
// Authenticate authenticates with the device
func (p *Post) Authenticate() error
```
```go
// BoardEnable enables or disables a board
// BoardEnable(idx []int, enable bool)
func (p *Post) BoardEnable(idx []int, enable bool) error
```
```go
// ClearHashrate clears the hashrate history of the device
func (p *Post) ClearHashrate() error
```
```go
// Coin sets the coin to mine
// Coin(confs []StratumConfig, coin string, uniqueId bool)
func (p *Post) Coin(confs []StratumConfig, coin string, uniqueId bool) error
```
```go
// CriticalTemp sets the critical temperature of the device
// CriticalTemp(temp int)
func (p *Post) CriticalTemp(temp int) error
```
```go
// DefaultConfig resets the device to its default configuration
func (p *Post) DefaultConfig() error
```
```go
// FanSpeed sets the fan speed of the device
// FanSpeed(idleSpeed, targetTemp int)
func (p *Post) FanSpeed(idleSpeed, targetTemp int) error
```
```go
// Id sets the device ID
// Id(unique bool)
func (p *Post) Id(unique bool) error
```
```go
// IdVarient sets the device ID varient
// IdVarient(varient int)
func (p *Post) IdVarient(varient int) error
```
```go
// Identify identifies the device
// Identify(identify bool)
func (p *Post) Identify(identify bool) error
```
```go
// IdleOnConnectionLost sets the device to idle when the connection is lost
// IdleOnConnectionLost(idle bool)
func (p *Post) IdleOnConnectionLost(idle bool) error
```
```go
// Miner starts or stops the miner
// Miner(control bool)
func (p *Post) Miner(control bool) error
```
```go
// Network sets the network settings of the device
// Network(dhcp bool)
func (p *Post) Network(dhcp bool) error
```
```go
// Overdrive enables or disables overdrive
// Overdrive(od bool)
func (p *Post) Overdrive(od bool) error
```
```go
// Password sets the password of the device
// Password(pass string)
func (p *Post) Password(pass string) error
```
```go
// PerpetualTune enables or disables perpetual tune
// PerpetualTune(pt bool)
func (p *Post) PerpetualTune(pt bool) error
```
```go
// Reboot reboots the device
// Reboot(delay int)
func (p *Post) Reboot(delay int) error
```
```go
// ShutdownTemp sets the shutdown temperature of the device
// ShutdownTemp(temp int)
func (p *Post) ShutdownTemp(temp int) error
```
```go
// SoftReboot soft reboots the device
func (p *Post) SoftReboot() error
```
```go
// Tune tunes the device
// Tune(freq, volt float64)
func (p *Post) Tune(freq, volt float64) error
```
```go
// TuneClockAll tunes the clock of all boards
// TuneClockAll(freq float64)
func (p *Post) TuneClockAll(freq float64) error
```
```go
// TuneClockBoard tunes the clock of a specific board
// TuneClockBoard(settings []BoardTune)
func (p *Post) TuneClockBoard(settings []BoardTune) error
```
```go
// TuneVoltage tunes the voltage of the device
// TuneVoltage(volt float64)
func (p *Post) TuneVoltage(volt float64) error
```