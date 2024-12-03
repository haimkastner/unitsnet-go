package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// PowerDensityUnits enumeration
type PowerDensityUnits string

const (
    
        // 
        PowerDensityWattPerCubicMeter PowerDensityUnits = "WattPerCubicMeter"
        // 
        PowerDensityWattPerCubicInch PowerDensityUnits = "WattPerCubicInch"
        // 
        PowerDensityWattPerCubicFoot PowerDensityUnits = "WattPerCubicFoot"
        // 
        PowerDensityWattPerLiter PowerDensityUnits = "WattPerLiter"
        // 
        PowerDensityPicowattPerCubicMeter PowerDensityUnits = "PicowattPerCubicMeter"
        // 
        PowerDensityNanowattPerCubicMeter PowerDensityUnits = "NanowattPerCubicMeter"
        // 
        PowerDensityMicrowattPerCubicMeter PowerDensityUnits = "MicrowattPerCubicMeter"
        // 
        PowerDensityMilliwattPerCubicMeter PowerDensityUnits = "MilliwattPerCubicMeter"
        // 
        PowerDensityDeciwattPerCubicMeter PowerDensityUnits = "DeciwattPerCubicMeter"
        // 
        PowerDensityDecawattPerCubicMeter PowerDensityUnits = "DecawattPerCubicMeter"
        // 
        PowerDensityKilowattPerCubicMeter PowerDensityUnits = "KilowattPerCubicMeter"
        // 
        PowerDensityMegawattPerCubicMeter PowerDensityUnits = "MegawattPerCubicMeter"
        // 
        PowerDensityGigawattPerCubicMeter PowerDensityUnits = "GigawattPerCubicMeter"
        // 
        PowerDensityTerawattPerCubicMeter PowerDensityUnits = "TerawattPerCubicMeter"
        // 
        PowerDensityPicowattPerCubicInch PowerDensityUnits = "PicowattPerCubicInch"
        // 
        PowerDensityNanowattPerCubicInch PowerDensityUnits = "NanowattPerCubicInch"
        // 
        PowerDensityMicrowattPerCubicInch PowerDensityUnits = "MicrowattPerCubicInch"
        // 
        PowerDensityMilliwattPerCubicInch PowerDensityUnits = "MilliwattPerCubicInch"
        // 
        PowerDensityDeciwattPerCubicInch PowerDensityUnits = "DeciwattPerCubicInch"
        // 
        PowerDensityDecawattPerCubicInch PowerDensityUnits = "DecawattPerCubicInch"
        // 
        PowerDensityKilowattPerCubicInch PowerDensityUnits = "KilowattPerCubicInch"
        // 
        PowerDensityMegawattPerCubicInch PowerDensityUnits = "MegawattPerCubicInch"
        // 
        PowerDensityGigawattPerCubicInch PowerDensityUnits = "GigawattPerCubicInch"
        // 
        PowerDensityTerawattPerCubicInch PowerDensityUnits = "TerawattPerCubicInch"
        // 
        PowerDensityPicowattPerCubicFoot PowerDensityUnits = "PicowattPerCubicFoot"
        // 
        PowerDensityNanowattPerCubicFoot PowerDensityUnits = "NanowattPerCubicFoot"
        // 
        PowerDensityMicrowattPerCubicFoot PowerDensityUnits = "MicrowattPerCubicFoot"
        // 
        PowerDensityMilliwattPerCubicFoot PowerDensityUnits = "MilliwattPerCubicFoot"
        // 
        PowerDensityDeciwattPerCubicFoot PowerDensityUnits = "DeciwattPerCubicFoot"
        // 
        PowerDensityDecawattPerCubicFoot PowerDensityUnits = "DecawattPerCubicFoot"
        // 
        PowerDensityKilowattPerCubicFoot PowerDensityUnits = "KilowattPerCubicFoot"
        // 
        PowerDensityMegawattPerCubicFoot PowerDensityUnits = "MegawattPerCubicFoot"
        // 
        PowerDensityGigawattPerCubicFoot PowerDensityUnits = "GigawattPerCubicFoot"
        // 
        PowerDensityTerawattPerCubicFoot PowerDensityUnits = "TerawattPerCubicFoot"
        // 
        PowerDensityPicowattPerLiter PowerDensityUnits = "PicowattPerLiter"
        // 
        PowerDensityNanowattPerLiter PowerDensityUnits = "NanowattPerLiter"
        // 
        PowerDensityMicrowattPerLiter PowerDensityUnits = "MicrowattPerLiter"
        // 
        PowerDensityMilliwattPerLiter PowerDensityUnits = "MilliwattPerLiter"
        // 
        PowerDensityDeciwattPerLiter PowerDensityUnits = "DeciwattPerLiter"
        // 
        PowerDensityDecawattPerLiter PowerDensityUnits = "DecawattPerLiter"
        // 
        PowerDensityKilowattPerLiter PowerDensityUnits = "KilowattPerLiter"
        // 
        PowerDensityMegawattPerLiter PowerDensityUnits = "MegawattPerLiter"
        // 
        PowerDensityGigawattPerLiter PowerDensityUnits = "GigawattPerLiter"
        // 
        PowerDensityTerawattPerLiter PowerDensityUnits = "TerawattPerLiter"
)

// PowerDensityDto represents an PowerDensity
type PowerDensityDto struct {
	Value float64
	Unit  PowerDensityUnits
}

// PowerDensityDtoFactory struct to group related functions
type PowerDensityDtoFactory struct{}

func (udf PowerDensityDtoFactory) FromJSON(data []byte) (*PowerDensityDto, error) {
	a := PowerDensityDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a PowerDensityDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// PowerDensity struct
type PowerDensity struct {
	value       float64
    
    watts_per_cubic_meterLazy *float64 
    watts_per_cubic_inchLazy *float64 
    watts_per_cubic_footLazy *float64 
    watts_per_literLazy *float64 
    picowatts_per_cubic_meterLazy *float64 
    nanowatts_per_cubic_meterLazy *float64 
    microwatts_per_cubic_meterLazy *float64 
    milliwatts_per_cubic_meterLazy *float64 
    deciwatts_per_cubic_meterLazy *float64 
    decawatts_per_cubic_meterLazy *float64 
    kilowatts_per_cubic_meterLazy *float64 
    megawatts_per_cubic_meterLazy *float64 
    gigawatts_per_cubic_meterLazy *float64 
    terawatts_per_cubic_meterLazy *float64 
    picowatts_per_cubic_inchLazy *float64 
    nanowatts_per_cubic_inchLazy *float64 
    microwatts_per_cubic_inchLazy *float64 
    milliwatts_per_cubic_inchLazy *float64 
    deciwatts_per_cubic_inchLazy *float64 
    decawatts_per_cubic_inchLazy *float64 
    kilowatts_per_cubic_inchLazy *float64 
    megawatts_per_cubic_inchLazy *float64 
    gigawatts_per_cubic_inchLazy *float64 
    terawatts_per_cubic_inchLazy *float64 
    picowatts_per_cubic_footLazy *float64 
    nanowatts_per_cubic_footLazy *float64 
    microwatts_per_cubic_footLazy *float64 
    milliwatts_per_cubic_footLazy *float64 
    deciwatts_per_cubic_footLazy *float64 
    decawatts_per_cubic_footLazy *float64 
    kilowatts_per_cubic_footLazy *float64 
    megawatts_per_cubic_footLazy *float64 
    gigawatts_per_cubic_footLazy *float64 
    terawatts_per_cubic_footLazy *float64 
    picowatts_per_literLazy *float64 
    nanowatts_per_literLazy *float64 
    microwatts_per_literLazy *float64 
    milliwatts_per_literLazy *float64 
    deciwatts_per_literLazy *float64 
    decawatts_per_literLazy *float64 
    kilowatts_per_literLazy *float64 
    megawatts_per_literLazy *float64 
    gigawatts_per_literLazy *float64 
    terawatts_per_literLazy *float64 
}

// PowerDensityFactory struct to group related functions
type PowerDensityFactory struct{}

func (uf PowerDensityFactory) CreatePowerDensity(value float64, unit PowerDensityUnits) (*PowerDensity, error) {
	return newPowerDensity(value, unit)
}

func (uf PowerDensityFactory) FromDto(dto PowerDensityDto) (*PowerDensity, error) {
	return newPowerDensity(dto.Value, dto.Unit)
}

func (uf PowerDensityFactory) FromDtoJSON(data []byte) (*PowerDensity, error) {
	unitDto, err := PowerDensityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return PowerDensityFactory{}.FromDto(*unitDto)
}


// FromWattPerCubicMeter creates a new PowerDensity instance from WattPerCubicMeter.
func (uf PowerDensityFactory) FromWattsPerCubicMeter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityWattPerCubicMeter)
}

// FromWattPerCubicInch creates a new PowerDensity instance from WattPerCubicInch.
func (uf PowerDensityFactory) FromWattsPerCubicInch(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityWattPerCubicInch)
}

// FromWattPerCubicFoot creates a new PowerDensity instance from WattPerCubicFoot.
func (uf PowerDensityFactory) FromWattsPerCubicFoot(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityWattPerCubicFoot)
}

// FromWattPerLiter creates a new PowerDensity instance from WattPerLiter.
func (uf PowerDensityFactory) FromWattsPerLiter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityWattPerLiter)
}

// FromPicowattPerCubicMeter creates a new PowerDensity instance from PicowattPerCubicMeter.
func (uf PowerDensityFactory) FromPicowattsPerCubicMeter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityPicowattPerCubicMeter)
}

// FromNanowattPerCubicMeter creates a new PowerDensity instance from NanowattPerCubicMeter.
func (uf PowerDensityFactory) FromNanowattsPerCubicMeter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityNanowattPerCubicMeter)
}

// FromMicrowattPerCubicMeter creates a new PowerDensity instance from MicrowattPerCubicMeter.
func (uf PowerDensityFactory) FromMicrowattsPerCubicMeter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityMicrowattPerCubicMeter)
}

// FromMilliwattPerCubicMeter creates a new PowerDensity instance from MilliwattPerCubicMeter.
func (uf PowerDensityFactory) FromMilliwattsPerCubicMeter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityMilliwattPerCubicMeter)
}

// FromDeciwattPerCubicMeter creates a new PowerDensity instance from DeciwattPerCubicMeter.
func (uf PowerDensityFactory) FromDeciwattsPerCubicMeter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityDeciwattPerCubicMeter)
}

// FromDecawattPerCubicMeter creates a new PowerDensity instance from DecawattPerCubicMeter.
func (uf PowerDensityFactory) FromDecawattsPerCubicMeter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityDecawattPerCubicMeter)
}

// FromKilowattPerCubicMeter creates a new PowerDensity instance from KilowattPerCubicMeter.
func (uf PowerDensityFactory) FromKilowattsPerCubicMeter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityKilowattPerCubicMeter)
}

// FromMegawattPerCubicMeter creates a new PowerDensity instance from MegawattPerCubicMeter.
func (uf PowerDensityFactory) FromMegawattsPerCubicMeter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityMegawattPerCubicMeter)
}

// FromGigawattPerCubicMeter creates a new PowerDensity instance from GigawattPerCubicMeter.
func (uf PowerDensityFactory) FromGigawattsPerCubicMeter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityGigawattPerCubicMeter)
}

// FromTerawattPerCubicMeter creates a new PowerDensity instance from TerawattPerCubicMeter.
func (uf PowerDensityFactory) FromTerawattsPerCubicMeter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityTerawattPerCubicMeter)
}

// FromPicowattPerCubicInch creates a new PowerDensity instance from PicowattPerCubicInch.
func (uf PowerDensityFactory) FromPicowattsPerCubicInch(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityPicowattPerCubicInch)
}

// FromNanowattPerCubicInch creates a new PowerDensity instance from NanowattPerCubicInch.
func (uf PowerDensityFactory) FromNanowattsPerCubicInch(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityNanowattPerCubicInch)
}

// FromMicrowattPerCubicInch creates a new PowerDensity instance from MicrowattPerCubicInch.
func (uf PowerDensityFactory) FromMicrowattsPerCubicInch(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityMicrowattPerCubicInch)
}

// FromMilliwattPerCubicInch creates a new PowerDensity instance from MilliwattPerCubicInch.
func (uf PowerDensityFactory) FromMilliwattsPerCubicInch(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityMilliwattPerCubicInch)
}

// FromDeciwattPerCubicInch creates a new PowerDensity instance from DeciwattPerCubicInch.
func (uf PowerDensityFactory) FromDeciwattsPerCubicInch(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityDeciwattPerCubicInch)
}

// FromDecawattPerCubicInch creates a new PowerDensity instance from DecawattPerCubicInch.
func (uf PowerDensityFactory) FromDecawattsPerCubicInch(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityDecawattPerCubicInch)
}

// FromKilowattPerCubicInch creates a new PowerDensity instance from KilowattPerCubicInch.
func (uf PowerDensityFactory) FromKilowattsPerCubicInch(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityKilowattPerCubicInch)
}

// FromMegawattPerCubicInch creates a new PowerDensity instance from MegawattPerCubicInch.
func (uf PowerDensityFactory) FromMegawattsPerCubicInch(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityMegawattPerCubicInch)
}

// FromGigawattPerCubicInch creates a new PowerDensity instance from GigawattPerCubicInch.
func (uf PowerDensityFactory) FromGigawattsPerCubicInch(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityGigawattPerCubicInch)
}

// FromTerawattPerCubicInch creates a new PowerDensity instance from TerawattPerCubicInch.
func (uf PowerDensityFactory) FromTerawattsPerCubicInch(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityTerawattPerCubicInch)
}

// FromPicowattPerCubicFoot creates a new PowerDensity instance from PicowattPerCubicFoot.
func (uf PowerDensityFactory) FromPicowattsPerCubicFoot(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityPicowattPerCubicFoot)
}

// FromNanowattPerCubicFoot creates a new PowerDensity instance from NanowattPerCubicFoot.
func (uf PowerDensityFactory) FromNanowattsPerCubicFoot(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityNanowattPerCubicFoot)
}

// FromMicrowattPerCubicFoot creates a new PowerDensity instance from MicrowattPerCubicFoot.
func (uf PowerDensityFactory) FromMicrowattsPerCubicFoot(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityMicrowattPerCubicFoot)
}

// FromMilliwattPerCubicFoot creates a new PowerDensity instance from MilliwattPerCubicFoot.
func (uf PowerDensityFactory) FromMilliwattsPerCubicFoot(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityMilliwattPerCubicFoot)
}

// FromDeciwattPerCubicFoot creates a new PowerDensity instance from DeciwattPerCubicFoot.
func (uf PowerDensityFactory) FromDeciwattsPerCubicFoot(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityDeciwattPerCubicFoot)
}

// FromDecawattPerCubicFoot creates a new PowerDensity instance from DecawattPerCubicFoot.
func (uf PowerDensityFactory) FromDecawattsPerCubicFoot(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityDecawattPerCubicFoot)
}

// FromKilowattPerCubicFoot creates a new PowerDensity instance from KilowattPerCubicFoot.
func (uf PowerDensityFactory) FromKilowattsPerCubicFoot(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityKilowattPerCubicFoot)
}

// FromMegawattPerCubicFoot creates a new PowerDensity instance from MegawattPerCubicFoot.
func (uf PowerDensityFactory) FromMegawattsPerCubicFoot(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityMegawattPerCubicFoot)
}

// FromGigawattPerCubicFoot creates a new PowerDensity instance from GigawattPerCubicFoot.
func (uf PowerDensityFactory) FromGigawattsPerCubicFoot(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityGigawattPerCubicFoot)
}

// FromTerawattPerCubicFoot creates a new PowerDensity instance from TerawattPerCubicFoot.
func (uf PowerDensityFactory) FromTerawattsPerCubicFoot(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityTerawattPerCubicFoot)
}

// FromPicowattPerLiter creates a new PowerDensity instance from PicowattPerLiter.
func (uf PowerDensityFactory) FromPicowattsPerLiter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityPicowattPerLiter)
}

// FromNanowattPerLiter creates a new PowerDensity instance from NanowattPerLiter.
func (uf PowerDensityFactory) FromNanowattsPerLiter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityNanowattPerLiter)
}

// FromMicrowattPerLiter creates a new PowerDensity instance from MicrowattPerLiter.
func (uf PowerDensityFactory) FromMicrowattsPerLiter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityMicrowattPerLiter)
}

// FromMilliwattPerLiter creates a new PowerDensity instance from MilliwattPerLiter.
func (uf PowerDensityFactory) FromMilliwattsPerLiter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityMilliwattPerLiter)
}

// FromDeciwattPerLiter creates a new PowerDensity instance from DeciwattPerLiter.
func (uf PowerDensityFactory) FromDeciwattsPerLiter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityDeciwattPerLiter)
}

// FromDecawattPerLiter creates a new PowerDensity instance from DecawattPerLiter.
func (uf PowerDensityFactory) FromDecawattsPerLiter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityDecawattPerLiter)
}

// FromKilowattPerLiter creates a new PowerDensity instance from KilowattPerLiter.
func (uf PowerDensityFactory) FromKilowattsPerLiter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityKilowattPerLiter)
}

// FromMegawattPerLiter creates a new PowerDensity instance from MegawattPerLiter.
func (uf PowerDensityFactory) FromMegawattsPerLiter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityMegawattPerLiter)
}

// FromGigawattPerLiter creates a new PowerDensity instance from GigawattPerLiter.
func (uf PowerDensityFactory) FromGigawattsPerLiter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityGigawattPerLiter)
}

// FromTerawattPerLiter creates a new PowerDensity instance from TerawattPerLiter.
func (uf PowerDensityFactory) FromTerawattsPerLiter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityTerawattPerLiter)
}




// newPowerDensity creates a new PowerDensity.
func newPowerDensity(value float64, fromUnit PowerDensityUnits) (*PowerDensity, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &PowerDensity{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of PowerDensity in WattPerCubicMeter.
func (a *PowerDensity) BaseValue() float64 {
	return a.value
}


// WattPerCubicMeter returns the value in WattPerCubicMeter.
func (a *PowerDensity) WattsPerCubicMeter() float64 {
	if a.watts_per_cubic_meterLazy != nil {
		return *a.watts_per_cubic_meterLazy
	}
	watts_per_cubic_meter := a.convertFromBase(PowerDensityWattPerCubicMeter)
	a.watts_per_cubic_meterLazy = &watts_per_cubic_meter
	return watts_per_cubic_meter
}

// WattPerCubicInch returns the value in WattPerCubicInch.
func (a *PowerDensity) WattsPerCubicInch() float64 {
	if a.watts_per_cubic_inchLazy != nil {
		return *a.watts_per_cubic_inchLazy
	}
	watts_per_cubic_inch := a.convertFromBase(PowerDensityWattPerCubicInch)
	a.watts_per_cubic_inchLazy = &watts_per_cubic_inch
	return watts_per_cubic_inch
}

// WattPerCubicFoot returns the value in WattPerCubicFoot.
func (a *PowerDensity) WattsPerCubicFoot() float64 {
	if a.watts_per_cubic_footLazy != nil {
		return *a.watts_per_cubic_footLazy
	}
	watts_per_cubic_foot := a.convertFromBase(PowerDensityWattPerCubicFoot)
	a.watts_per_cubic_footLazy = &watts_per_cubic_foot
	return watts_per_cubic_foot
}

// WattPerLiter returns the value in WattPerLiter.
func (a *PowerDensity) WattsPerLiter() float64 {
	if a.watts_per_literLazy != nil {
		return *a.watts_per_literLazy
	}
	watts_per_liter := a.convertFromBase(PowerDensityWattPerLiter)
	a.watts_per_literLazy = &watts_per_liter
	return watts_per_liter
}

// PicowattPerCubicMeter returns the value in PicowattPerCubicMeter.
func (a *PowerDensity) PicowattsPerCubicMeter() float64 {
	if a.picowatts_per_cubic_meterLazy != nil {
		return *a.picowatts_per_cubic_meterLazy
	}
	picowatts_per_cubic_meter := a.convertFromBase(PowerDensityPicowattPerCubicMeter)
	a.picowatts_per_cubic_meterLazy = &picowatts_per_cubic_meter
	return picowatts_per_cubic_meter
}

// NanowattPerCubicMeter returns the value in NanowattPerCubicMeter.
func (a *PowerDensity) NanowattsPerCubicMeter() float64 {
	if a.nanowatts_per_cubic_meterLazy != nil {
		return *a.nanowatts_per_cubic_meterLazy
	}
	nanowatts_per_cubic_meter := a.convertFromBase(PowerDensityNanowattPerCubicMeter)
	a.nanowatts_per_cubic_meterLazy = &nanowatts_per_cubic_meter
	return nanowatts_per_cubic_meter
}

// MicrowattPerCubicMeter returns the value in MicrowattPerCubicMeter.
func (a *PowerDensity) MicrowattsPerCubicMeter() float64 {
	if a.microwatts_per_cubic_meterLazy != nil {
		return *a.microwatts_per_cubic_meterLazy
	}
	microwatts_per_cubic_meter := a.convertFromBase(PowerDensityMicrowattPerCubicMeter)
	a.microwatts_per_cubic_meterLazy = &microwatts_per_cubic_meter
	return microwatts_per_cubic_meter
}

// MilliwattPerCubicMeter returns the value in MilliwattPerCubicMeter.
func (a *PowerDensity) MilliwattsPerCubicMeter() float64 {
	if a.milliwatts_per_cubic_meterLazy != nil {
		return *a.milliwatts_per_cubic_meterLazy
	}
	milliwatts_per_cubic_meter := a.convertFromBase(PowerDensityMilliwattPerCubicMeter)
	a.milliwatts_per_cubic_meterLazy = &milliwatts_per_cubic_meter
	return milliwatts_per_cubic_meter
}

// DeciwattPerCubicMeter returns the value in DeciwattPerCubicMeter.
func (a *PowerDensity) DeciwattsPerCubicMeter() float64 {
	if a.deciwatts_per_cubic_meterLazy != nil {
		return *a.deciwatts_per_cubic_meterLazy
	}
	deciwatts_per_cubic_meter := a.convertFromBase(PowerDensityDeciwattPerCubicMeter)
	a.deciwatts_per_cubic_meterLazy = &deciwatts_per_cubic_meter
	return deciwatts_per_cubic_meter
}

// DecawattPerCubicMeter returns the value in DecawattPerCubicMeter.
func (a *PowerDensity) DecawattsPerCubicMeter() float64 {
	if a.decawatts_per_cubic_meterLazy != nil {
		return *a.decawatts_per_cubic_meterLazy
	}
	decawatts_per_cubic_meter := a.convertFromBase(PowerDensityDecawattPerCubicMeter)
	a.decawatts_per_cubic_meterLazy = &decawatts_per_cubic_meter
	return decawatts_per_cubic_meter
}

// KilowattPerCubicMeter returns the value in KilowattPerCubicMeter.
func (a *PowerDensity) KilowattsPerCubicMeter() float64 {
	if a.kilowatts_per_cubic_meterLazy != nil {
		return *a.kilowatts_per_cubic_meterLazy
	}
	kilowatts_per_cubic_meter := a.convertFromBase(PowerDensityKilowattPerCubicMeter)
	a.kilowatts_per_cubic_meterLazy = &kilowatts_per_cubic_meter
	return kilowatts_per_cubic_meter
}

// MegawattPerCubicMeter returns the value in MegawattPerCubicMeter.
func (a *PowerDensity) MegawattsPerCubicMeter() float64 {
	if a.megawatts_per_cubic_meterLazy != nil {
		return *a.megawatts_per_cubic_meterLazy
	}
	megawatts_per_cubic_meter := a.convertFromBase(PowerDensityMegawattPerCubicMeter)
	a.megawatts_per_cubic_meterLazy = &megawatts_per_cubic_meter
	return megawatts_per_cubic_meter
}

// GigawattPerCubicMeter returns the value in GigawattPerCubicMeter.
func (a *PowerDensity) GigawattsPerCubicMeter() float64 {
	if a.gigawatts_per_cubic_meterLazy != nil {
		return *a.gigawatts_per_cubic_meterLazy
	}
	gigawatts_per_cubic_meter := a.convertFromBase(PowerDensityGigawattPerCubicMeter)
	a.gigawatts_per_cubic_meterLazy = &gigawatts_per_cubic_meter
	return gigawatts_per_cubic_meter
}

// TerawattPerCubicMeter returns the value in TerawattPerCubicMeter.
func (a *PowerDensity) TerawattsPerCubicMeter() float64 {
	if a.terawatts_per_cubic_meterLazy != nil {
		return *a.terawatts_per_cubic_meterLazy
	}
	terawatts_per_cubic_meter := a.convertFromBase(PowerDensityTerawattPerCubicMeter)
	a.terawatts_per_cubic_meterLazy = &terawatts_per_cubic_meter
	return terawatts_per_cubic_meter
}

// PicowattPerCubicInch returns the value in PicowattPerCubicInch.
func (a *PowerDensity) PicowattsPerCubicInch() float64 {
	if a.picowatts_per_cubic_inchLazy != nil {
		return *a.picowatts_per_cubic_inchLazy
	}
	picowatts_per_cubic_inch := a.convertFromBase(PowerDensityPicowattPerCubicInch)
	a.picowatts_per_cubic_inchLazy = &picowatts_per_cubic_inch
	return picowatts_per_cubic_inch
}

// NanowattPerCubicInch returns the value in NanowattPerCubicInch.
func (a *PowerDensity) NanowattsPerCubicInch() float64 {
	if a.nanowatts_per_cubic_inchLazy != nil {
		return *a.nanowatts_per_cubic_inchLazy
	}
	nanowatts_per_cubic_inch := a.convertFromBase(PowerDensityNanowattPerCubicInch)
	a.nanowatts_per_cubic_inchLazy = &nanowatts_per_cubic_inch
	return nanowatts_per_cubic_inch
}

// MicrowattPerCubicInch returns the value in MicrowattPerCubicInch.
func (a *PowerDensity) MicrowattsPerCubicInch() float64 {
	if a.microwatts_per_cubic_inchLazy != nil {
		return *a.microwatts_per_cubic_inchLazy
	}
	microwatts_per_cubic_inch := a.convertFromBase(PowerDensityMicrowattPerCubicInch)
	a.microwatts_per_cubic_inchLazy = &microwatts_per_cubic_inch
	return microwatts_per_cubic_inch
}

// MilliwattPerCubicInch returns the value in MilliwattPerCubicInch.
func (a *PowerDensity) MilliwattsPerCubicInch() float64 {
	if a.milliwatts_per_cubic_inchLazy != nil {
		return *a.milliwatts_per_cubic_inchLazy
	}
	milliwatts_per_cubic_inch := a.convertFromBase(PowerDensityMilliwattPerCubicInch)
	a.milliwatts_per_cubic_inchLazy = &milliwatts_per_cubic_inch
	return milliwatts_per_cubic_inch
}

// DeciwattPerCubicInch returns the value in DeciwattPerCubicInch.
func (a *PowerDensity) DeciwattsPerCubicInch() float64 {
	if a.deciwatts_per_cubic_inchLazy != nil {
		return *a.deciwatts_per_cubic_inchLazy
	}
	deciwatts_per_cubic_inch := a.convertFromBase(PowerDensityDeciwattPerCubicInch)
	a.deciwatts_per_cubic_inchLazy = &deciwatts_per_cubic_inch
	return deciwatts_per_cubic_inch
}

// DecawattPerCubicInch returns the value in DecawattPerCubicInch.
func (a *PowerDensity) DecawattsPerCubicInch() float64 {
	if a.decawatts_per_cubic_inchLazy != nil {
		return *a.decawatts_per_cubic_inchLazy
	}
	decawatts_per_cubic_inch := a.convertFromBase(PowerDensityDecawattPerCubicInch)
	a.decawatts_per_cubic_inchLazy = &decawatts_per_cubic_inch
	return decawatts_per_cubic_inch
}

// KilowattPerCubicInch returns the value in KilowattPerCubicInch.
func (a *PowerDensity) KilowattsPerCubicInch() float64 {
	if a.kilowatts_per_cubic_inchLazy != nil {
		return *a.kilowatts_per_cubic_inchLazy
	}
	kilowatts_per_cubic_inch := a.convertFromBase(PowerDensityKilowattPerCubicInch)
	a.kilowatts_per_cubic_inchLazy = &kilowatts_per_cubic_inch
	return kilowatts_per_cubic_inch
}

// MegawattPerCubicInch returns the value in MegawattPerCubicInch.
func (a *PowerDensity) MegawattsPerCubicInch() float64 {
	if a.megawatts_per_cubic_inchLazy != nil {
		return *a.megawatts_per_cubic_inchLazy
	}
	megawatts_per_cubic_inch := a.convertFromBase(PowerDensityMegawattPerCubicInch)
	a.megawatts_per_cubic_inchLazy = &megawatts_per_cubic_inch
	return megawatts_per_cubic_inch
}

// GigawattPerCubicInch returns the value in GigawattPerCubicInch.
func (a *PowerDensity) GigawattsPerCubicInch() float64 {
	if a.gigawatts_per_cubic_inchLazy != nil {
		return *a.gigawatts_per_cubic_inchLazy
	}
	gigawatts_per_cubic_inch := a.convertFromBase(PowerDensityGigawattPerCubicInch)
	a.gigawatts_per_cubic_inchLazy = &gigawatts_per_cubic_inch
	return gigawatts_per_cubic_inch
}

// TerawattPerCubicInch returns the value in TerawattPerCubicInch.
func (a *PowerDensity) TerawattsPerCubicInch() float64 {
	if a.terawatts_per_cubic_inchLazy != nil {
		return *a.terawatts_per_cubic_inchLazy
	}
	terawatts_per_cubic_inch := a.convertFromBase(PowerDensityTerawattPerCubicInch)
	a.terawatts_per_cubic_inchLazy = &terawatts_per_cubic_inch
	return terawatts_per_cubic_inch
}

// PicowattPerCubicFoot returns the value in PicowattPerCubicFoot.
func (a *PowerDensity) PicowattsPerCubicFoot() float64 {
	if a.picowatts_per_cubic_footLazy != nil {
		return *a.picowatts_per_cubic_footLazy
	}
	picowatts_per_cubic_foot := a.convertFromBase(PowerDensityPicowattPerCubicFoot)
	a.picowatts_per_cubic_footLazy = &picowatts_per_cubic_foot
	return picowatts_per_cubic_foot
}

// NanowattPerCubicFoot returns the value in NanowattPerCubicFoot.
func (a *PowerDensity) NanowattsPerCubicFoot() float64 {
	if a.nanowatts_per_cubic_footLazy != nil {
		return *a.nanowatts_per_cubic_footLazy
	}
	nanowatts_per_cubic_foot := a.convertFromBase(PowerDensityNanowattPerCubicFoot)
	a.nanowatts_per_cubic_footLazy = &nanowatts_per_cubic_foot
	return nanowatts_per_cubic_foot
}

// MicrowattPerCubicFoot returns the value in MicrowattPerCubicFoot.
func (a *PowerDensity) MicrowattsPerCubicFoot() float64 {
	if a.microwatts_per_cubic_footLazy != nil {
		return *a.microwatts_per_cubic_footLazy
	}
	microwatts_per_cubic_foot := a.convertFromBase(PowerDensityMicrowattPerCubicFoot)
	a.microwatts_per_cubic_footLazy = &microwatts_per_cubic_foot
	return microwatts_per_cubic_foot
}

// MilliwattPerCubicFoot returns the value in MilliwattPerCubicFoot.
func (a *PowerDensity) MilliwattsPerCubicFoot() float64 {
	if a.milliwatts_per_cubic_footLazy != nil {
		return *a.milliwatts_per_cubic_footLazy
	}
	milliwatts_per_cubic_foot := a.convertFromBase(PowerDensityMilliwattPerCubicFoot)
	a.milliwatts_per_cubic_footLazy = &milliwatts_per_cubic_foot
	return milliwatts_per_cubic_foot
}

// DeciwattPerCubicFoot returns the value in DeciwattPerCubicFoot.
func (a *PowerDensity) DeciwattsPerCubicFoot() float64 {
	if a.deciwatts_per_cubic_footLazy != nil {
		return *a.deciwatts_per_cubic_footLazy
	}
	deciwatts_per_cubic_foot := a.convertFromBase(PowerDensityDeciwattPerCubicFoot)
	a.deciwatts_per_cubic_footLazy = &deciwatts_per_cubic_foot
	return deciwatts_per_cubic_foot
}

// DecawattPerCubicFoot returns the value in DecawattPerCubicFoot.
func (a *PowerDensity) DecawattsPerCubicFoot() float64 {
	if a.decawatts_per_cubic_footLazy != nil {
		return *a.decawatts_per_cubic_footLazy
	}
	decawatts_per_cubic_foot := a.convertFromBase(PowerDensityDecawattPerCubicFoot)
	a.decawatts_per_cubic_footLazy = &decawatts_per_cubic_foot
	return decawatts_per_cubic_foot
}

// KilowattPerCubicFoot returns the value in KilowattPerCubicFoot.
func (a *PowerDensity) KilowattsPerCubicFoot() float64 {
	if a.kilowatts_per_cubic_footLazy != nil {
		return *a.kilowatts_per_cubic_footLazy
	}
	kilowatts_per_cubic_foot := a.convertFromBase(PowerDensityKilowattPerCubicFoot)
	a.kilowatts_per_cubic_footLazy = &kilowatts_per_cubic_foot
	return kilowatts_per_cubic_foot
}

// MegawattPerCubicFoot returns the value in MegawattPerCubicFoot.
func (a *PowerDensity) MegawattsPerCubicFoot() float64 {
	if a.megawatts_per_cubic_footLazy != nil {
		return *a.megawatts_per_cubic_footLazy
	}
	megawatts_per_cubic_foot := a.convertFromBase(PowerDensityMegawattPerCubicFoot)
	a.megawatts_per_cubic_footLazy = &megawatts_per_cubic_foot
	return megawatts_per_cubic_foot
}

// GigawattPerCubicFoot returns the value in GigawattPerCubicFoot.
func (a *PowerDensity) GigawattsPerCubicFoot() float64 {
	if a.gigawatts_per_cubic_footLazy != nil {
		return *a.gigawatts_per_cubic_footLazy
	}
	gigawatts_per_cubic_foot := a.convertFromBase(PowerDensityGigawattPerCubicFoot)
	a.gigawatts_per_cubic_footLazy = &gigawatts_per_cubic_foot
	return gigawatts_per_cubic_foot
}

// TerawattPerCubicFoot returns the value in TerawattPerCubicFoot.
func (a *PowerDensity) TerawattsPerCubicFoot() float64 {
	if a.terawatts_per_cubic_footLazy != nil {
		return *a.terawatts_per_cubic_footLazy
	}
	terawatts_per_cubic_foot := a.convertFromBase(PowerDensityTerawattPerCubicFoot)
	a.terawatts_per_cubic_footLazy = &terawatts_per_cubic_foot
	return terawatts_per_cubic_foot
}

// PicowattPerLiter returns the value in PicowattPerLiter.
func (a *PowerDensity) PicowattsPerLiter() float64 {
	if a.picowatts_per_literLazy != nil {
		return *a.picowatts_per_literLazy
	}
	picowatts_per_liter := a.convertFromBase(PowerDensityPicowattPerLiter)
	a.picowatts_per_literLazy = &picowatts_per_liter
	return picowatts_per_liter
}

// NanowattPerLiter returns the value in NanowattPerLiter.
func (a *PowerDensity) NanowattsPerLiter() float64 {
	if a.nanowatts_per_literLazy != nil {
		return *a.nanowatts_per_literLazy
	}
	nanowatts_per_liter := a.convertFromBase(PowerDensityNanowattPerLiter)
	a.nanowatts_per_literLazy = &nanowatts_per_liter
	return nanowatts_per_liter
}

// MicrowattPerLiter returns the value in MicrowattPerLiter.
func (a *PowerDensity) MicrowattsPerLiter() float64 {
	if a.microwatts_per_literLazy != nil {
		return *a.microwatts_per_literLazy
	}
	microwatts_per_liter := a.convertFromBase(PowerDensityMicrowattPerLiter)
	a.microwatts_per_literLazy = &microwatts_per_liter
	return microwatts_per_liter
}

// MilliwattPerLiter returns the value in MilliwattPerLiter.
func (a *PowerDensity) MilliwattsPerLiter() float64 {
	if a.milliwatts_per_literLazy != nil {
		return *a.milliwatts_per_literLazy
	}
	milliwatts_per_liter := a.convertFromBase(PowerDensityMilliwattPerLiter)
	a.milliwatts_per_literLazy = &milliwatts_per_liter
	return milliwatts_per_liter
}

// DeciwattPerLiter returns the value in DeciwattPerLiter.
func (a *PowerDensity) DeciwattsPerLiter() float64 {
	if a.deciwatts_per_literLazy != nil {
		return *a.deciwatts_per_literLazy
	}
	deciwatts_per_liter := a.convertFromBase(PowerDensityDeciwattPerLiter)
	a.deciwatts_per_literLazy = &deciwatts_per_liter
	return deciwatts_per_liter
}

// DecawattPerLiter returns the value in DecawattPerLiter.
func (a *PowerDensity) DecawattsPerLiter() float64 {
	if a.decawatts_per_literLazy != nil {
		return *a.decawatts_per_literLazy
	}
	decawatts_per_liter := a.convertFromBase(PowerDensityDecawattPerLiter)
	a.decawatts_per_literLazy = &decawatts_per_liter
	return decawatts_per_liter
}

// KilowattPerLiter returns the value in KilowattPerLiter.
func (a *PowerDensity) KilowattsPerLiter() float64 {
	if a.kilowatts_per_literLazy != nil {
		return *a.kilowatts_per_literLazy
	}
	kilowatts_per_liter := a.convertFromBase(PowerDensityKilowattPerLiter)
	a.kilowatts_per_literLazy = &kilowatts_per_liter
	return kilowatts_per_liter
}

// MegawattPerLiter returns the value in MegawattPerLiter.
func (a *PowerDensity) MegawattsPerLiter() float64 {
	if a.megawatts_per_literLazy != nil {
		return *a.megawatts_per_literLazy
	}
	megawatts_per_liter := a.convertFromBase(PowerDensityMegawattPerLiter)
	a.megawatts_per_literLazy = &megawatts_per_liter
	return megawatts_per_liter
}

// GigawattPerLiter returns the value in GigawattPerLiter.
func (a *PowerDensity) GigawattsPerLiter() float64 {
	if a.gigawatts_per_literLazy != nil {
		return *a.gigawatts_per_literLazy
	}
	gigawatts_per_liter := a.convertFromBase(PowerDensityGigawattPerLiter)
	a.gigawatts_per_literLazy = &gigawatts_per_liter
	return gigawatts_per_liter
}

// TerawattPerLiter returns the value in TerawattPerLiter.
func (a *PowerDensity) TerawattsPerLiter() float64 {
	if a.terawatts_per_literLazy != nil {
		return *a.terawatts_per_literLazy
	}
	terawatts_per_liter := a.convertFromBase(PowerDensityTerawattPerLiter)
	a.terawatts_per_literLazy = &terawatts_per_liter
	return terawatts_per_liter
}


// ToDto creates an PowerDensityDto representation.
func (a *PowerDensity) ToDto(holdInUnit *PowerDensityUnits) PowerDensityDto {
	if holdInUnit == nil {
		defaultUnit := PowerDensityWattPerCubicMeter // Default value
		holdInUnit = &defaultUnit
	}

	return PowerDensityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an PowerDensityDto representation.
func (a *PowerDensity) ToDtoJSON(holdInUnit *PowerDensityUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts PowerDensity to a specific unit value.
func (a *PowerDensity) Convert(toUnit PowerDensityUnits) float64 {
	switch toUnit { 
    case PowerDensityWattPerCubicMeter:
		return a.WattsPerCubicMeter()
    case PowerDensityWattPerCubicInch:
		return a.WattsPerCubicInch()
    case PowerDensityWattPerCubicFoot:
		return a.WattsPerCubicFoot()
    case PowerDensityWattPerLiter:
		return a.WattsPerLiter()
    case PowerDensityPicowattPerCubicMeter:
		return a.PicowattsPerCubicMeter()
    case PowerDensityNanowattPerCubicMeter:
		return a.NanowattsPerCubicMeter()
    case PowerDensityMicrowattPerCubicMeter:
		return a.MicrowattsPerCubicMeter()
    case PowerDensityMilliwattPerCubicMeter:
		return a.MilliwattsPerCubicMeter()
    case PowerDensityDeciwattPerCubicMeter:
		return a.DeciwattsPerCubicMeter()
    case PowerDensityDecawattPerCubicMeter:
		return a.DecawattsPerCubicMeter()
    case PowerDensityKilowattPerCubicMeter:
		return a.KilowattsPerCubicMeter()
    case PowerDensityMegawattPerCubicMeter:
		return a.MegawattsPerCubicMeter()
    case PowerDensityGigawattPerCubicMeter:
		return a.GigawattsPerCubicMeter()
    case PowerDensityTerawattPerCubicMeter:
		return a.TerawattsPerCubicMeter()
    case PowerDensityPicowattPerCubicInch:
		return a.PicowattsPerCubicInch()
    case PowerDensityNanowattPerCubicInch:
		return a.NanowattsPerCubicInch()
    case PowerDensityMicrowattPerCubicInch:
		return a.MicrowattsPerCubicInch()
    case PowerDensityMilliwattPerCubicInch:
		return a.MilliwattsPerCubicInch()
    case PowerDensityDeciwattPerCubicInch:
		return a.DeciwattsPerCubicInch()
    case PowerDensityDecawattPerCubicInch:
		return a.DecawattsPerCubicInch()
    case PowerDensityKilowattPerCubicInch:
		return a.KilowattsPerCubicInch()
    case PowerDensityMegawattPerCubicInch:
		return a.MegawattsPerCubicInch()
    case PowerDensityGigawattPerCubicInch:
		return a.GigawattsPerCubicInch()
    case PowerDensityTerawattPerCubicInch:
		return a.TerawattsPerCubicInch()
    case PowerDensityPicowattPerCubicFoot:
		return a.PicowattsPerCubicFoot()
    case PowerDensityNanowattPerCubicFoot:
		return a.NanowattsPerCubicFoot()
    case PowerDensityMicrowattPerCubicFoot:
		return a.MicrowattsPerCubicFoot()
    case PowerDensityMilliwattPerCubicFoot:
		return a.MilliwattsPerCubicFoot()
    case PowerDensityDeciwattPerCubicFoot:
		return a.DeciwattsPerCubicFoot()
    case PowerDensityDecawattPerCubicFoot:
		return a.DecawattsPerCubicFoot()
    case PowerDensityKilowattPerCubicFoot:
		return a.KilowattsPerCubicFoot()
    case PowerDensityMegawattPerCubicFoot:
		return a.MegawattsPerCubicFoot()
    case PowerDensityGigawattPerCubicFoot:
		return a.GigawattsPerCubicFoot()
    case PowerDensityTerawattPerCubicFoot:
		return a.TerawattsPerCubicFoot()
    case PowerDensityPicowattPerLiter:
		return a.PicowattsPerLiter()
    case PowerDensityNanowattPerLiter:
		return a.NanowattsPerLiter()
    case PowerDensityMicrowattPerLiter:
		return a.MicrowattsPerLiter()
    case PowerDensityMilliwattPerLiter:
		return a.MilliwattsPerLiter()
    case PowerDensityDeciwattPerLiter:
		return a.DeciwattsPerLiter()
    case PowerDensityDecawattPerLiter:
		return a.DecawattsPerLiter()
    case PowerDensityKilowattPerLiter:
		return a.KilowattsPerLiter()
    case PowerDensityMegawattPerLiter:
		return a.MegawattsPerLiter()
    case PowerDensityGigawattPerLiter:
		return a.GigawattsPerLiter()
    case PowerDensityTerawattPerLiter:
		return a.TerawattsPerLiter()
	default:
		return 0
	}
}

func (a *PowerDensity) convertFromBase(toUnit PowerDensityUnits) float64 {
    value := a.value
	switch toUnit { 
	case PowerDensityWattPerCubicMeter:
		return (value) 
	case PowerDensityWattPerCubicInch:
		return (value / 6.102374409473228e4) 
	case PowerDensityWattPerCubicFoot:
		return (value / 3.531466672148859e1) 
	case PowerDensityWattPerLiter:
		return (value / 1.0e3) 
	case PowerDensityPicowattPerCubicMeter:
		return ((value) / 1e-12) 
	case PowerDensityNanowattPerCubicMeter:
		return ((value) / 1e-09) 
	case PowerDensityMicrowattPerCubicMeter:
		return ((value) / 1e-06) 
	case PowerDensityMilliwattPerCubicMeter:
		return ((value) / 0.001) 
	case PowerDensityDeciwattPerCubicMeter:
		return ((value) / 0.1) 
	case PowerDensityDecawattPerCubicMeter:
		return ((value) / 10.0) 
	case PowerDensityKilowattPerCubicMeter:
		return ((value) / 1000.0) 
	case PowerDensityMegawattPerCubicMeter:
		return ((value) / 1000000.0) 
	case PowerDensityGigawattPerCubicMeter:
		return ((value) / 1000000000.0) 
	case PowerDensityTerawattPerCubicMeter:
		return ((value) / 1000000000000.0) 
	case PowerDensityPicowattPerCubicInch:
		return ((value / 6.102374409473228e4) / 1e-12) 
	case PowerDensityNanowattPerCubicInch:
		return ((value / 6.102374409473228e4) / 1e-09) 
	case PowerDensityMicrowattPerCubicInch:
		return ((value / 6.102374409473228e4) / 1e-06) 
	case PowerDensityMilliwattPerCubicInch:
		return ((value / 6.102374409473228e4) / 0.001) 
	case PowerDensityDeciwattPerCubicInch:
		return ((value / 6.102374409473228e4) / 0.1) 
	case PowerDensityDecawattPerCubicInch:
		return ((value / 6.102374409473228e4) / 10.0) 
	case PowerDensityKilowattPerCubicInch:
		return ((value / 6.102374409473228e4) / 1000.0) 
	case PowerDensityMegawattPerCubicInch:
		return ((value / 6.102374409473228e4) / 1000000.0) 
	case PowerDensityGigawattPerCubicInch:
		return ((value / 6.102374409473228e4) / 1000000000.0) 
	case PowerDensityTerawattPerCubicInch:
		return ((value / 6.102374409473228e4) / 1000000000000.0) 
	case PowerDensityPicowattPerCubicFoot:
		return ((value / 3.531466672148859e1) / 1e-12) 
	case PowerDensityNanowattPerCubicFoot:
		return ((value / 3.531466672148859e1) / 1e-09) 
	case PowerDensityMicrowattPerCubicFoot:
		return ((value / 3.531466672148859e1) / 1e-06) 
	case PowerDensityMilliwattPerCubicFoot:
		return ((value / 3.531466672148859e1) / 0.001) 
	case PowerDensityDeciwattPerCubicFoot:
		return ((value / 3.531466672148859e1) / 0.1) 
	case PowerDensityDecawattPerCubicFoot:
		return ((value / 3.531466672148859e1) / 10.0) 
	case PowerDensityKilowattPerCubicFoot:
		return ((value / 3.531466672148859e1) / 1000.0) 
	case PowerDensityMegawattPerCubicFoot:
		return ((value / 3.531466672148859e1) / 1000000.0) 
	case PowerDensityGigawattPerCubicFoot:
		return ((value / 3.531466672148859e1) / 1000000000.0) 
	case PowerDensityTerawattPerCubicFoot:
		return ((value / 3.531466672148859e1) / 1000000000000.0) 
	case PowerDensityPicowattPerLiter:
		return ((value / 1.0e3) / 1e-12) 
	case PowerDensityNanowattPerLiter:
		return ((value / 1.0e3) / 1e-09) 
	case PowerDensityMicrowattPerLiter:
		return ((value / 1.0e3) / 1e-06) 
	case PowerDensityMilliwattPerLiter:
		return ((value / 1.0e3) / 0.001) 
	case PowerDensityDeciwattPerLiter:
		return ((value / 1.0e3) / 0.1) 
	case PowerDensityDecawattPerLiter:
		return ((value / 1.0e3) / 10.0) 
	case PowerDensityKilowattPerLiter:
		return ((value / 1.0e3) / 1000.0) 
	case PowerDensityMegawattPerLiter:
		return ((value / 1.0e3) / 1000000.0) 
	case PowerDensityGigawattPerLiter:
		return ((value / 1.0e3) / 1000000000.0) 
	case PowerDensityTerawattPerLiter:
		return ((value / 1.0e3) / 1000000000000.0) 
	default:
		return math.NaN()
	}
}

func (a *PowerDensity) convertToBase(value float64, fromUnit PowerDensityUnits) float64 {
	switch fromUnit { 
	case PowerDensityWattPerCubicMeter:
		return (value) 
	case PowerDensityWattPerCubicInch:
		return (value * 6.102374409473228e4) 
	case PowerDensityWattPerCubicFoot:
		return (value * 3.531466672148859e1) 
	case PowerDensityWattPerLiter:
		return (value * 1.0e3) 
	case PowerDensityPicowattPerCubicMeter:
		return ((value) * 1e-12) 
	case PowerDensityNanowattPerCubicMeter:
		return ((value) * 1e-09) 
	case PowerDensityMicrowattPerCubicMeter:
		return ((value) * 1e-06) 
	case PowerDensityMilliwattPerCubicMeter:
		return ((value) * 0.001) 
	case PowerDensityDeciwattPerCubicMeter:
		return ((value) * 0.1) 
	case PowerDensityDecawattPerCubicMeter:
		return ((value) * 10.0) 
	case PowerDensityKilowattPerCubicMeter:
		return ((value) * 1000.0) 
	case PowerDensityMegawattPerCubicMeter:
		return ((value) * 1000000.0) 
	case PowerDensityGigawattPerCubicMeter:
		return ((value) * 1000000000.0) 
	case PowerDensityTerawattPerCubicMeter:
		return ((value) * 1000000000000.0) 
	case PowerDensityPicowattPerCubicInch:
		return ((value * 6.102374409473228e4) * 1e-12) 
	case PowerDensityNanowattPerCubicInch:
		return ((value * 6.102374409473228e4) * 1e-09) 
	case PowerDensityMicrowattPerCubicInch:
		return ((value * 6.102374409473228e4) * 1e-06) 
	case PowerDensityMilliwattPerCubicInch:
		return ((value * 6.102374409473228e4) * 0.001) 
	case PowerDensityDeciwattPerCubicInch:
		return ((value * 6.102374409473228e4) * 0.1) 
	case PowerDensityDecawattPerCubicInch:
		return ((value * 6.102374409473228e4) * 10.0) 
	case PowerDensityKilowattPerCubicInch:
		return ((value * 6.102374409473228e4) * 1000.0) 
	case PowerDensityMegawattPerCubicInch:
		return ((value * 6.102374409473228e4) * 1000000.0) 
	case PowerDensityGigawattPerCubicInch:
		return ((value * 6.102374409473228e4) * 1000000000.0) 
	case PowerDensityTerawattPerCubicInch:
		return ((value * 6.102374409473228e4) * 1000000000000.0) 
	case PowerDensityPicowattPerCubicFoot:
		return ((value * 3.531466672148859e1) * 1e-12) 
	case PowerDensityNanowattPerCubicFoot:
		return ((value * 3.531466672148859e1) * 1e-09) 
	case PowerDensityMicrowattPerCubicFoot:
		return ((value * 3.531466672148859e1) * 1e-06) 
	case PowerDensityMilliwattPerCubicFoot:
		return ((value * 3.531466672148859e1) * 0.001) 
	case PowerDensityDeciwattPerCubicFoot:
		return ((value * 3.531466672148859e1) * 0.1) 
	case PowerDensityDecawattPerCubicFoot:
		return ((value * 3.531466672148859e1) * 10.0) 
	case PowerDensityKilowattPerCubicFoot:
		return ((value * 3.531466672148859e1) * 1000.0) 
	case PowerDensityMegawattPerCubicFoot:
		return ((value * 3.531466672148859e1) * 1000000.0) 
	case PowerDensityGigawattPerCubicFoot:
		return ((value * 3.531466672148859e1) * 1000000000.0) 
	case PowerDensityTerawattPerCubicFoot:
		return ((value * 3.531466672148859e1) * 1000000000000.0) 
	case PowerDensityPicowattPerLiter:
		return ((value * 1.0e3) * 1e-12) 
	case PowerDensityNanowattPerLiter:
		return ((value * 1.0e3) * 1e-09) 
	case PowerDensityMicrowattPerLiter:
		return ((value * 1.0e3) * 1e-06) 
	case PowerDensityMilliwattPerLiter:
		return ((value * 1.0e3) * 0.001) 
	case PowerDensityDeciwattPerLiter:
		return ((value * 1.0e3) * 0.1) 
	case PowerDensityDecawattPerLiter:
		return ((value * 1.0e3) * 10.0) 
	case PowerDensityKilowattPerLiter:
		return ((value * 1.0e3) * 1000.0) 
	case PowerDensityMegawattPerLiter:
		return ((value * 1.0e3) * 1000000.0) 
	case PowerDensityGigawattPerLiter:
		return ((value * 1.0e3) * 1000000000.0) 
	case PowerDensityTerawattPerLiter:
		return ((value * 1.0e3) * 1000000000000.0) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a PowerDensity) String() string {
	return a.ToString(PowerDensityWattPerCubicMeter, 2)
}

// ToString formats the PowerDensity to string.
// fractionalDigits -1 for not mention
func (a *PowerDensity) ToString(unit PowerDensityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *PowerDensity) getUnitAbbreviation(unit PowerDensityUnits) string {
	switch unit { 
	case PowerDensityWattPerCubicMeter:
		return "W/m³" 
	case PowerDensityWattPerCubicInch:
		return "W/in³" 
	case PowerDensityWattPerCubicFoot:
		return "W/ft³" 
	case PowerDensityWattPerLiter:
		return "W/l" 
	case PowerDensityPicowattPerCubicMeter:
		return "pW/m³" 
	case PowerDensityNanowattPerCubicMeter:
		return "nW/m³" 
	case PowerDensityMicrowattPerCubicMeter:
		return "μW/m³" 
	case PowerDensityMilliwattPerCubicMeter:
		return "mW/m³" 
	case PowerDensityDeciwattPerCubicMeter:
		return "dW/m³" 
	case PowerDensityDecawattPerCubicMeter:
		return "daW/m³" 
	case PowerDensityKilowattPerCubicMeter:
		return "kW/m³" 
	case PowerDensityMegawattPerCubicMeter:
		return "MW/m³" 
	case PowerDensityGigawattPerCubicMeter:
		return "GW/m³" 
	case PowerDensityTerawattPerCubicMeter:
		return "TW/m³" 
	case PowerDensityPicowattPerCubicInch:
		return "pW/in³" 
	case PowerDensityNanowattPerCubicInch:
		return "nW/in³" 
	case PowerDensityMicrowattPerCubicInch:
		return "μW/in³" 
	case PowerDensityMilliwattPerCubicInch:
		return "mW/in³" 
	case PowerDensityDeciwattPerCubicInch:
		return "dW/in³" 
	case PowerDensityDecawattPerCubicInch:
		return "daW/in³" 
	case PowerDensityKilowattPerCubicInch:
		return "kW/in³" 
	case PowerDensityMegawattPerCubicInch:
		return "MW/in³" 
	case PowerDensityGigawattPerCubicInch:
		return "GW/in³" 
	case PowerDensityTerawattPerCubicInch:
		return "TW/in³" 
	case PowerDensityPicowattPerCubicFoot:
		return "pW/ft³" 
	case PowerDensityNanowattPerCubicFoot:
		return "nW/ft³" 
	case PowerDensityMicrowattPerCubicFoot:
		return "μW/ft³" 
	case PowerDensityMilliwattPerCubicFoot:
		return "mW/ft³" 
	case PowerDensityDeciwattPerCubicFoot:
		return "dW/ft³" 
	case PowerDensityDecawattPerCubicFoot:
		return "daW/ft³" 
	case PowerDensityKilowattPerCubicFoot:
		return "kW/ft³" 
	case PowerDensityMegawattPerCubicFoot:
		return "MW/ft³" 
	case PowerDensityGigawattPerCubicFoot:
		return "GW/ft³" 
	case PowerDensityTerawattPerCubicFoot:
		return "TW/ft³" 
	case PowerDensityPicowattPerLiter:
		return "pW/l" 
	case PowerDensityNanowattPerLiter:
		return "nW/l" 
	case PowerDensityMicrowattPerLiter:
		return "μW/l" 
	case PowerDensityMilliwattPerLiter:
		return "mW/l" 
	case PowerDensityDeciwattPerLiter:
		return "dW/l" 
	case PowerDensityDecawattPerLiter:
		return "daW/l" 
	case PowerDensityKilowattPerLiter:
		return "kW/l" 
	case PowerDensityMegawattPerLiter:
		return "MW/l" 
	case PowerDensityGigawattPerLiter:
		return "GW/l" 
	case PowerDensityTerawattPerLiter:
		return "TW/l" 
	default:
		return ""
	}
}

// Check if the given PowerDensity are equals to the current PowerDensity
func (a *PowerDensity) Equals(other *PowerDensity) bool {
	return a.value == other.BaseValue()
}

// Check if the given PowerDensity are equals to the current PowerDensity
func (a *PowerDensity) CompareTo(other *PowerDensity) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given PowerDensity to the current PowerDensity.
func (a *PowerDensity) Add(other *PowerDensity) *PowerDensity {
	return &PowerDensity{value: a.value + other.BaseValue()}
}

// Subtract the given PowerDensity to the current PowerDensity.
func (a *PowerDensity) Subtract(other *PowerDensity) *PowerDensity {
	return &PowerDensity{value: a.value - other.BaseValue()}
}

// Multiply the given PowerDensity to the current PowerDensity.
func (a *PowerDensity) Multiply(other *PowerDensity) *PowerDensity {
	return &PowerDensity{value: a.value * other.BaseValue()}
}

// Divide the given PowerDensity to the current PowerDensity.
func (a *PowerDensity) Divide(other *PowerDensity) *PowerDensity {
	return &PowerDensity{value: a.value / other.BaseValue()}
}