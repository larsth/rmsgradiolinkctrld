package config

import (
	"errors"
	"fmt"
	"sync"

	"github.com/larsth/rmsgradiolinkctrld/logging"
)

var (
	ErrLocationConfigHasBenInited          error = errors.New("This LocationConfig type has already been initialized.")
	ErrLocationConfigNoHwDevicesConfigured error = errors.New("This LocationConfig type has no hardware devices.")
)

type HwDeviceConfig struct {
	Name      string `"json:name"`
	ShortName string `"json:short_name"`
	Type      string `"json:type"`     //Both IPv4 and IPv6 are supported
	IPAddress string `json:ip_address` //Port 80 is assumed
}

type Config struct {
	Name              string                     `"json:name"`
	ShortName         string                     `"json:short_name"`
	Host              string                     `"json:host"`
	HwDeviceConfigs   []*HwDeviceConfig          `"json:hardware_devices"`
	hwDeviceLUT       map[string]*HwDeviceConfig `"json:-"`
	mutex_hwDeviceLUT sync.RWMutex               `"json:-"`
	InitError         error                      `"json:-"`
}

func (lc *Config) Init() error {
	lc.mutex_hwDeviceLUT.Lock()
	defer lc.mutex_hwDeviceLUT.Unlock()

	if lc.hwDeviceLUT != nil {
		if lc.InitError == nil {
			lc.InitError = ErrLocationConfigHasBenInited
		}
		return ErrLocationConfigHasBenInited
	}

	if len(lc.HwDeviceConfigs) == 0 {
		lc.InitError = ErrLocationConfigNoHwDevicesConfigured
		logging.Fatalf("%s", ErrLocationConfigNoHwDevicesConfigured.Error())
		return ErrLocationConfigNoHwDevicesConfigured //make the compiler happy
	}

	lc.hwDeviceLUT = make(map[string]*HwDeviceConfig, len(lc.HwDeviceConfigs))
	for _, hdc := range lc.HwDeviceConfigs {
		lc.hwDeviceLUT[hdc.ShortName] = hdc
	}

	return nil //no error
}

func (lc *Config) IndexOf(shortName string) (*HwDeviceConfig, error) {
	lc.mutex_hwDeviceLUT.RLock()
	defer lc.mutex_hwDeviceLUT.RUnlock()

	const formatstr string = "No such shortname. Shortname: \"%s\""
	var (
		hwc *HwDeviceConfig
		ok  bool
	)

	if hwc, ok = lc.hwDeviceLUT[shortName]; ok == false {
		return nil, fmt.Errorf(formatstr, shortName)
	}
	return hwc, nil
}
