package epicapi

type CapabilitiesResult struct {
	BoardSerialNumbers   []string `json:"Board Serial Numbers"`
	ChipType             string   `json:"Chip Type"`
	ChipsPerBank         int      `json:"Chips Per Bank"`
	Coins                []string `json:"Coins"`
	DefaultClock         int      `json:"Default Clock"`
	DefaultHashrate      int      `json:"Default Hashrate"`
	DefaultVoltage       int      `json:"Default Voltage"`
	Display              []string `json:"Display"`
	MaxClockConfigs      int      `json:"Max Clock Configs"`
	MaxHBs               int      `json:"Max HBs"`
	Model                string   `json:"Model"`
	ModelSubtype         string   `json:"Model Subtype"`
	PerformanceEstimator struct {
		ChipCount              int `json:"Chip Count"`
		HashesPerSecondPerChip int `json:"Hashes Per Second Per Chip"`
	} `json:"Performance Estimator"`
	PerpetualTune []struct {
		Algorithm   string `json:"algorithm"`
		Description string `json:"description"`
		Max         int    `json:"max"`
		Min         int    `json:"min"`
		Name        string `json:"name"`
	} `json:"PerpetualTune"`
	PsuInfo struct {
		FW      int `json:"FW"`
		MaxVout int `json:"Max Vout"`
		MinVout int `json:"Min Vout"`
		PSUType int `json:"PSU Type"`
	} `json:"Psu Info"`
	TempSensorInfo struct {
		Count int    `json:"Count"`
		Type  string `json:"Type"`
	} `json:"Temp Sensor Info"`
	TunePresets []struct {
		Clk      int `json:"clk"`
		Hashrate int `json:"hashrate"`
		Power    int `json:"power"`
		Voltage  int `json:"voltage"`
	} `json:"Tune Presets"`
	VoltageEstimator struct {
		Slope      int `json:"Slope"`
		YIntercept int `json:"Y Intercept"`
	} `json:"Voltage Estimator"`
}

type ClocksResult []struct {
	Index int       `json:"Index"` // Index of the board.
	Data  []float64 `json:"Data"`  // Clock data for the chips on the board.
	Total float64   `json:"Total"` // Total number of clock entries.
}

type ClocksConfigResult []struct {
	Data  [][]any `json:"Data"`
	Index int     `json:"Index"`
	Total []any   `json:"Total"`
}

type HashrateResult []struct {
	Index int         `json:"Index"` // Index of the board.
	Data  [][]float64 `json:"Data"`  // Hashrate data for the chips on the board.
	Total []float64   `json:"Total"` // Total hashrate values.
}

type HashrateHistoryContinuousResult []struct {
	Hashrate  int `json:"hashrate"`
	Timestamp int `json:"timestamp"`
}

type HashrateHistoryDiscreteResult []struct {
	Hashrate  int `json:"hashrate"`
	Timestamp int `json:"timestamp"`
}

type HistoryResult struct {
	History []string `json:"History"`
}

type NetworkResult struct {
	Dhcp struct {
		Address    string `json:"address"`
		DNS        any    `json:"dns"`
		Gateway    string `json:"gateway"`
		MacAddress string `json:"mac_address"`
		Netmask    string `json:"netmask"`
	} `json:"dhcp"`
}

type PerpetualTuneResult struct {
	Algorithms struct {
		AdditionalProp1 struct {
			MinThrottleTarget int    `json:"Min Throttle Target"`
			Optimized         bool   `json:"Optimized"`
			Target            int    `json:"Target"`
			ThrottleStep      int    `json:"Throttle Step"`
			ThrottleTarget    int    `json:"Throttle Target"`
			Unit              string `json:"Unit"`
		} `json:"additionalProp1"`
		AdditionalProp2 struct {
			MinThrottleTarget int    `json:"Min Throttle Target"`
			Optimized         bool   `json:"Optimized"`
			Target            int    `json:"Target"`
			ThrottleStep      int    `json:"Throttle Step"`
			ThrottleTarget    int    `json:"Throttle Target"`
			Unit              string `json:"Unit"`
		} `json:"additionalProp2"`
		AdditionalProp3 struct {
			MinThrottleTarget int    `json:"Min Throttle Target"`
			Optimized         bool   `json:"Optimized"`
			Target            int    `json:"Target"`
			ThrottleStep      int    `json:"Throttle Step"`
			ThrottleTarget    int    `json:"Throttle Target"`
			Unit              string `json:"Unit"`
		} `json:"additionalProp3"`
	} `json:"Algorithms"`
	CurrentAlgorithm string `json:"Current Algorithm"`
	Enable           bool   `json:"Enable"`
	Running          bool   `json:"Running"`
}

type SummaryResult struct {
	Status struct {
		OperatingState    string `json:"Operating State"`     // Current operating state of the miner.
		LastCommand       any    `json:"Last Command"`        // Last command executed.
		LastCommandResult any    `json:"Last Command Result"` // Result of the last command.
		LastError         any    `json:"Last Error"`          // Last error encountered.
	} `json:"Status"`
	Hostname   string `json:"Hostname"` // Hostname of the miner.
	PresetInfo struct {
		TargetPower int `json:"Target Power"` // Target power setting.
	} `json:"PresetInfo"`
	Software string `json:"Software"` // Software version.
	Mining   struct {
		Coin      string `json:"Coin"`      // Coin being mined.
		Algorithm string `json:"Algorithm"` // Mining algorithm.
	} `json:"Mining"`
	Stratum struct {
		ConfigID              int     `json:"Config Id"`                // Stratum configuration ID.
		CurrentPool           string  `json:"Current Pool"`             // Current mining pool.
		CurrentUser           string  `json:"Current User"`             // Current mining user.
		IsPoolConnected       bool    `json:"IsPoolConnected"`          // Whether the pool is connected.
		AverageLatency        float64 `json:"Average Latency"`          // Average latency to the pool.
		WorkerUniqueID        bool    `json:"Worker Unique Id"`         // Whether worker unique ID is used.
		WorkerUniqueIDVariant string  `json:"Worker Unique Id Variant"` // Variant of worker unique ID.
	} `json:"Stratum"`
	Session struct {
		StartupTimestamp           float64 `json:"Startup Timestamp"`             // Timestamp of miner startup.
		StartupString              string  `json:"Startup String"`                // Startup string.
		Uptime                     float64 `json:"Uptime"`                        // Miner uptime in seconds.
		LastWorkTimestamp          float64 `json:"Last Work Timestamp"`           // Timestamp of last work.
		WorkReceived               float64 `json:"WorkReceived"`                  // Number of work received.
		ActiveHBs                  float64 `json:"Active HBs"`                    // Number of active hash boards.
		AverageMHs                 float64 `json:"Average MHs"`                   // Average hashrate in MH/s.
		LastAverageMHs             float64 `json:"LastAverageMHs"`                // Last average hashrate.
		Accepted                   float64 `json:"Accepted"`                      // Number of accepted shares.
		Rejected                   float64 `json:"Rejected"`                      // Number of rejected shares.
		Submitted                  float64 `json:"Submitted"`                     // Number of submitted shares.
		LastAcceptedShareTimestamp float64 `json:"Last Accepted Share Timestamp"` // Timestamp of last accepted share.
		Difficulty                 float64 `json:"Difficulty"`                    // Current mining difficulty.
	} `json:"Session"`
	HBs []struct {
		Index         int       `json:"Index"`          // Index of the hash board.
		InputVoltage  float64   `json:"Input Voltage"`  // Input voltage of the hash board.
		OutputVoltage float64   `json:"Output Voltage"` // Output voltage of the hash board.
		InputCurrent  float64   `json:"Input Current"`  // Input current of the hash board.
		OutputCurrent float64   `json:"Output Current"` // Output current of the hash board.
		InputPower    float64   `json:"Input Power"`    // Input power of the hash board.
		OutputPower   float64   `json:"Output Power"`   // Output power of the hash board.
		Temperature   float64   `json:"Temperature"`    // Temperature of the hash board.
		CoreClock     []any     `json:"Core Clock"`     // Core clock data.
		Hashrate      []float64 `json:"Hashrate"`       // Hashrate data.
		CoreClockAvg  float64   `json:"Core Clock Avg"` // Average core clock.
	} `json:"HBs"`
	HBStatus []struct {
		Index    int  `json:"Index"`    // Index of the hash board.
		Enabled  bool `json:"Enabled"`  // Whether the hash board is enabled.
		Detected bool `json:"Detected"` // Whether the hash board is detected.
	} `json:"HBStatus"`
	Fans struct {
		FansSpeed int `json:"Fans Speed"` // Current fan speed.
		FanMode   struct {
			Auto struct {
				TargetTemperature int `json:"Target Temperature"` // Target temperature for auto fan mode.
				IdleSpeed         int `json:"Idle Speed"`         // Idle speed for auto fan mode.
			} `json:"Auto"`
		} `json:"Fan Mode"`
		MinimumWorkingFans int `json:"Minimum Working Fans"` // Minimum number of working fans.
	} `json:"Fans"`
	FansRpm struct {
		FansSpeed1 int `json:"Fans Speed 1"` // Fan speed for fan 1.
		FansSpeed2 int `json:"Fans Speed 2"` // Fan speed for fan 2.
		FansSpeed3 int `json:"Fans Speed 3"` // Fan speed for fan 3.
		FansSpeed4 int `json:"Fans Speed 4"` // Fan speed for fan 4.
	} `json:"Fans Rpm"`
	Misc struct {
		LocateMinerState bool    `json:"Locate Miner State"` // State of locate miner function.
		ShutdownTemp     float64 `json:"Shutdown Temp"`      // Shutdown temperature.
		CriticalTemp     float64 `json:"Critical Temp"`      // Critical temperature.
	} `json:"Misc"`
	StratumConfigs []struct {
		Pool     string `json:"pool"`     // Stratum pool URL.
		Login    string `json:"login"`    // Stratum login.
		Password string `json:"password"` // Stratum password.
	} `json:"StratumConfigs"`
	PerpetualTune struct {
		Running   bool `json:"Running"` // Whether perpetual tune is running.
		Algorithm struct {
			ChipTune struct {
				Optimized         bool   `json:"Optimized"`           // Whether chip tune is optimized.
				Target            int    `json:"Target"`              // Chip tune target.
				ThrottleTarget    any    `json:"Throttle Target"`     // Throttle target.
				MinThrottleTarget int    `json:"Min Throttle Target"` // Minimum throttle target.
				ThrottleStep      int    `json:"Throttle Step"`       // Throttle step.
				Unit              string `json:"Unit"`                // Unit for chip tune.
			} `json:"ChipTune"`
		} `json:"Algorithm"`
	} `json:"PerpetualTune"`
	PowerSupplyStats struct {
		InputVoltage  float64 `json:"Input Voltage"`  // Input voltage of power supply.
		OutputVoltage float64 `json:"Output Voltage"` // Output voltage of power supply.
		InputCurrent  float64 `json:"Input Current"`  // Input current of power supply.
		OutputCurrent float64 `json:"Output Current"` // Output current of power supply.
		InputPower    float64 `json:"Input Power"`    // Input power of power supply.
		OutputPower   float64 `json:"Output Power"`   // Output power of power supply.
		TargetVoltage float64 `json:"Target Voltage"` // Target voltage of power supply.
	} `json:"Power Supply Stats"`
	HwConfig struct {
		BoardsTargetClock []struct {
			Index int `json:"Index"` // Index of the board.
			Data  int `json:"Data"`  // Target clock data for the board.
		} `json:"HwConfig"`
	} `json:"HwConfig"`
	IdleOnConnectionLost  bool `json:"IdleOnConnectionLost"`     // Whether to idle on connection lost.
	Overdrive             bool `json:"Overdrive"`                // Whether overdrive is enabled.
	DisableBoardOnFailure bool `json:"Disable Board On Failure"` // Whether to disable board on failure.
}

type TempsResult []struct {
	Data  []int   `json:"Data"`
	Index float64 `json:"Index"`
	Total float64 `json:"Total"`
}

type TempsBoardResult []struct {
	Data  []int   `json:"Data"`
	Index float64 `json:"Index"`
	Total float64 `json:"Total"`
}

type TempsChipResult []struct {
	Index int       `json:"Index"` // Index of the board.
	Data  []float64 `json:"Data"`  // Temperature data for the chips on the board.
	Total float64   `json:"Total"` // Total number of temperature entries.
}

type VoltageResult []struct {
	Index int       `json:"Index"` // Index of the board.
	Data  []float64 `json:"Data"`  // Voltage data for the chips on the board.
	Total float64   `json:"Total"` // Total number of voltage entries.
}
