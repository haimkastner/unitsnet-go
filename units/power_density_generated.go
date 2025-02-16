// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// PowerDensityUnits defines various units of PowerDensity.
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

// PowerDensityDto represents a PowerDensity measurement with a numerical value and its corresponding unit.
type PowerDensityDto struct {
    // Value is the numerical representation of the PowerDensity.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the PowerDensity, as defined in the PowerDensityUnits enumeration.
	Unit  PowerDensityUnits `json:"unit"`
}

// PowerDensityDtoFactory groups methods for creating and serializing PowerDensityDto objects.
type PowerDensityDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a PowerDensityDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf PowerDensityDtoFactory) FromJSON(data []byte) (*PowerDensityDto, error) {
	a := PowerDensityDto{}

    // Parse JSON into PowerDensityDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a PowerDensityDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a PowerDensityDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// PowerDensity represents a measurement in a of PowerDensity.
//
// The amount of power in a volume.
type PowerDensity struct {
	// value is the base measurement stored internally.
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

// PowerDensityFactory groups methods for creating PowerDensity instances.
type PowerDensityFactory struct{}

// CreatePowerDensity creates a new PowerDensity instance from the given value and unit.
func (uf PowerDensityFactory) CreatePowerDensity(value float64, unit PowerDensityUnits) (*PowerDensity, error) {
	return newPowerDensity(value, unit)
}

// FromDto converts a PowerDensityDto to a PowerDensity instance.
func (uf PowerDensityFactory) FromDto(dto PowerDensityDto) (*PowerDensity, error) {
	return newPowerDensity(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a PowerDensity instance.
func (uf PowerDensityFactory) FromDtoJSON(data []byte) (*PowerDensity, error) {
	unitDto, err := PowerDensityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse PowerDensityDto from JSON: %w", err)
	}
	return PowerDensityFactory{}.FromDto(*unitDto)
}


// FromWattsPerCubicMeter creates a new PowerDensity instance from a value in WattsPerCubicMeter.
func (uf PowerDensityFactory) FromWattsPerCubicMeter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityWattPerCubicMeter)
}

// FromWattsPerCubicInch creates a new PowerDensity instance from a value in WattsPerCubicInch.
func (uf PowerDensityFactory) FromWattsPerCubicInch(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityWattPerCubicInch)
}

// FromWattsPerCubicFoot creates a new PowerDensity instance from a value in WattsPerCubicFoot.
func (uf PowerDensityFactory) FromWattsPerCubicFoot(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityWattPerCubicFoot)
}

// FromWattsPerLiter creates a new PowerDensity instance from a value in WattsPerLiter.
func (uf PowerDensityFactory) FromWattsPerLiter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityWattPerLiter)
}

// FromPicowattsPerCubicMeter creates a new PowerDensity instance from a value in PicowattsPerCubicMeter.
func (uf PowerDensityFactory) FromPicowattsPerCubicMeter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityPicowattPerCubicMeter)
}

// FromNanowattsPerCubicMeter creates a new PowerDensity instance from a value in NanowattsPerCubicMeter.
func (uf PowerDensityFactory) FromNanowattsPerCubicMeter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityNanowattPerCubicMeter)
}

// FromMicrowattsPerCubicMeter creates a new PowerDensity instance from a value in MicrowattsPerCubicMeter.
func (uf PowerDensityFactory) FromMicrowattsPerCubicMeter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityMicrowattPerCubicMeter)
}

// FromMilliwattsPerCubicMeter creates a new PowerDensity instance from a value in MilliwattsPerCubicMeter.
func (uf PowerDensityFactory) FromMilliwattsPerCubicMeter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityMilliwattPerCubicMeter)
}

// FromDeciwattsPerCubicMeter creates a new PowerDensity instance from a value in DeciwattsPerCubicMeter.
func (uf PowerDensityFactory) FromDeciwattsPerCubicMeter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityDeciwattPerCubicMeter)
}

// FromDecawattsPerCubicMeter creates a new PowerDensity instance from a value in DecawattsPerCubicMeter.
func (uf PowerDensityFactory) FromDecawattsPerCubicMeter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityDecawattPerCubicMeter)
}

// FromKilowattsPerCubicMeter creates a new PowerDensity instance from a value in KilowattsPerCubicMeter.
func (uf PowerDensityFactory) FromKilowattsPerCubicMeter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityKilowattPerCubicMeter)
}

// FromMegawattsPerCubicMeter creates a new PowerDensity instance from a value in MegawattsPerCubicMeter.
func (uf PowerDensityFactory) FromMegawattsPerCubicMeter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityMegawattPerCubicMeter)
}

// FromGigawattsPerCubicMeter creates a new PowerDensity instance from a value in GigawattsPerCubicMeter.
func (uf PowerDensityFactory) FromGigawattsPerCubicMeter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityGigawattPerCubicMeter)
}

// FromTerawattsPerCubicMeter creates a new PowerDensity instance from a value in TerawattsPerCubicMeter.
func (uf PowerDensityFactory) FromTerawattsPerCubicMeter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityTerawattPerCubicMeter)
}

// FromPicowattsPerCubicInch creates a new PowerDensity instance from a value in PicowattsPerCubicInch.
func (uf PowerDensityFactory) FromPicowattsPerCubicInch(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityPicowattPerCubicInch)
}

// FromNanowattsPerCubicInch creates a new PowerDensity instance from a value in NanowattsPerCubicInch.
func (uf PowerDensityFactory) FromNanowattsPerCubicInch(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityNanowattPerCubicInch)
}

// FromMicrowattsPerCubicInch creates a new PowerDensity instance from a value in MicrowattsPerCubicInch.
func (uf PowerDensityFactory) FromMicrowattsPerCubicInch(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityMicrowattPerCubicInch)
}

// FromMilliwattsPerCubicInch creates a new PowerDensity instance from a value in MilliwattsPerCubicInch.
func (uf PowerDensityFactory) FromMilliwattsPerCubicInch(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityMilliwattPerCubicInch)
}

// FromDeciwattsPerCubicInch creates a new PowerDensity instance from a value in DeciwattsPerCubicInch.
func (uf PowerDensityFactory) FromDeciwattsPerCubicInch(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityDeciwattPerCubicInch)
}

// FromDecawattsPerCubicInch creates a new PowerDensity instance from a value in DecawattsPerCubicInch.
func (uf PowerDensityFactory) FromDecawattsPerCubicInch(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityDecawattPerCubicInch)
}

// FromKilowattsPerCubicInch creates a new PowerDensity instance from a value in KilowattsPerCubicInch.
func (uf PowerDensityFactory) FromKilowattsPerCubicInch(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityKilowattPerCubicInch)
}

// FromMegawattsPerCubicInch creates a new PowerDensity instance from a value in MegawattsPerCubicInch.
func (uf PowerDensityFactory) FromMegawattsPerCubicInch(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityMegawattPerCubicInch)
}

// FromGigawattsPerCubicInch creates a new PowerDensity instance from a value in GigawattsPerCubicInch.
func (uf PowerDensityFactory) FromGigawattsPerCubicInch(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityGigawattPerCubicInch)
}

// FromTerawattsPerCubicInch creates a new PowerDensity instance from a value in TerawattsPerCubicInch.
func (uf PowerDensityFactory) FromTerawattsPerCubicInch(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityTerawattPerCubicInch)
}

// FromPicowattsPerCubicFoot creates a new PowerDensity instance from a value in PicowattsPerCubicFoot.
func (uf PowerDensityFactory) FromPicowattsPerCubicFoot(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityPicowattPerCubicFoot)
}

// FromNanowattsPerCubicFoot creates a new PowerDensity instance from a value in NanowattsPerCubicFoot.
func (uf PowerDensityFactory) FromNanowattsPerCubicFoot(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityNanowattPerCubicFoot)
}

// FromMicrowattsPerCubicFoot creates a new PowerDensity instance from a value in MicrowattsPerCubicFoot.
func (uf PowerDensityFactory) FromMicrowattsPerCubicFoot(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityMicrowattPerCubicFoot)
}

// FromMilliwattsPerCubicFoot creates a new PowerDensity instance from a value in MilliwattsPerCubicFoot.
func (uf PowerDensityFactory) FromMilliwattsPerCubicFoot(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityMilliwattPerCubicFoot)
}

// FromDeciwattsPerCubicFoot creates a new PowerDensity instance from a value in DeciwattsPerCubicFoot.
func (uf PowerDensityFactory) FromDeciwattsPerCubicFoot(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityDeciwattPerCubicFoot)
}

// FromDecawattsPerCubicFoot creates a new PowerDensity instance from a value in DecawattsPerCubicFoot.
func (uf PowerDensityFactory) FromDecawattsPerCubicFoot(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityDecawattPerCubicFoot)
}

// FromKilowattsPerCubicFoot creates a new PowerDensity instance from a value in KilowattsPerCubicFoot.
func (uf PowerDensityFactory) FromKilowattsPerCubicFoot(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityKilowattPerCubicFoot)
}

// FromMegawattsPerCubicFoot creates a new PowerDensity instance from a value in MegawattsPerCubicFoot.
func (uf PowerDensityFactory) FromMegawattsPerCubicFoot(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityMegawattPerCubicFoot)
}

// FromGigawattsPerCubicFoot creates a new PowerDensity instance from a value in GigawattsPerCubicFoot.
func (uf PowerDensityFactory) FromGigawattsPerCubicFoot(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityGigawattPerCubicFoot)
}

// FromTerawattsPerCubicFoot creates a new PowerDensity instance from a value in TerawattsPerCubicFoot.
func (uf PowerDensityFactory) FromTerawattsPerCubicFoot(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityTerawattPerCubicFoot)
}

// FromPicowattsPerLiter creates a new PowerDensity instance from a value in PicowattsPerLiter.
func (uf PowerDensityFactory) FromPicowattsPerLiter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityPicowattPerLiter)
}

// FromNanowattsPerLiter creates a new PowerDensity instance from a value in NanowattsPerLiter.
func (uf PowerDensityFactory) FromNanowattsPerLiter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityNanowattPerLiter)
}

// FromMicrowattsPerLiter creates a new PowerDensity instance from a value in MicrowattsPerLiter.
func (uf PowerDensityFactory) FromMicrowattsPerLiter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityMicrowattPerLiter)
}

// FromMilliwattsPerLiter creates a new PowerDensity instance from a value in MilliwattsPerLiter.
func (uf PowerDensityFactory) FromMilliwattsPerLiter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityMilliwattPerLiter)
}

// FromDeciwattsPerLiter creates a new PowerDensity instance from a value in DeciwattsPerLiter.
func (uf PowerDensityFactory) FromDeciwattsPerLiter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityDeciwattPerLiter)
}

// FromDecawattsPerLiter creates a new PowerDensity instance from a value in DecawattsPerLiter.
func (uf PowerDensityFactory) FromDecawattsPerLiter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityDecawattPerLiter)
}

// FromKilowattsPerLiter creates a new PowerDensity instance from a value in KilowattsPerLiter.
func (uf PowerDensityFactory) FromKilowattsPerLiter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityKilowattPerLiter)
}

// FromMegawattsPerLiter creates a new PowerDensity instance from a value in MegawattsPerLiter.
func (uf PowerDensityFactory) FromMegawattsPerLiter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityMegawattPerLiter)
}

// FromGigawattsPerLiter creates a new PowerDensity instance from a value in GigawattsPerLiter.
func (uf PowerDensityFactory) FromGigawattsPerLiter(value float64) (*PowerDensity, error) {
	return newPowerDensity(value, PowerDensityGigawattPerLiter)
}

// FromTerawattsPerLiter creates a new PowerDensity instance from a value in TerawattsPerLiter.
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

// BaseValue returns the base value of PowerDensity in WattPerCubicMeter unit (the base/default quantity).
func (a *PowerDensity) BaseValue() float64 {
	return a.value
}


// WattsPerCubicMeter returns the PowerDensity value in WattsPerCubicMeter.
//
// 
func (a *PowerDensity) WattsPerCubicMeter() float64 {
	if a.watts_per_cubic_meterLazy != nil {
		return *a.watts_per_cubic_meterLazy
	}
	watts_per_cubic_meter := a.convertFromBase(PowerDensityWattPerCubicMeter)
	a.watts_per_cubic_meterLazy = &watts_per_cubic_meter
	return watts_per_cubic_meter
}

// WattsPerCubicInch returns the PowerDensity value in WattsPerCubicInch.
//
// 
func (a *PowerDensity) WattsPerCubicInch() float64 {
	if a.watts_per_cubic_inchLazy != nil {
		return *a.watts_per_cubic_inchLazy
	}
	watts_per_cubic_inch := a.convertFromBase(PowerDensityWattPerCubicInch)
	a.watts_per_cubic_inchLazy = &watts_per_cubic_inch
	return watts_per_cubic_inch
}

// WattsPerCubicFoot returns the PowerDensity value in WattsPerCubicFoot.
//
// 
func (a *PowerDensity) WattsPerCubicFoot() float64 {
	if a.watts_per_cubic_footLazy != nil {
		return *a.watts_per_cubic_footLazy
	}
	watts_per_cubic_foot := a.convertFromBase(PowerDensityWattPerCubicFoot)
	a.watts_per_cubic_footLazy = &watts_per_cubic_foot
	return watts_per_cubic_foot
}

// WattsPerLiter returns the PowerDensity value in WattsPerLiter.
//
// 
func (a *PowerDensity) WattsPerLiter() float64 {
	if a.watts_per_literLazy != nil {
		return *a.watts_per_literLazy
	}
	watts_per_liter := a.convertFromBase(PowerDensityWattPerLiter)
	a.watts_per_literLazy = &watts_per_liter
	return watts_per_liter
}

// PicowattsPerCubicMeter returns the PowerDensity value in PicowattsPerCubicMeter.
//
// 
func (a *PowerDensity) PicowattsPerCubicMeter() float64 {
	if a.picowatts_per_cubic_meterLazy != nil {
		return *a.picowatts_per_cubic_meterLazy
	}
	picowatts_per_cubic_meter := a.convertFromBase(PowerDensityPicowattPerCubicMeter)
	a.picowatts_per_cubic_meterLazy = &picowatts_per_cubic_meter
	return picowatts_per_cubic_meter
}

// NanowattsPerCubicMeter returns the PowerDensity value in NanowattsPerCubicMeter.
//
// 
func (a *PowerDensity) NanowattsPerCubicMeter() float64 {
	if a.nanowatts_per_cubic_meterLazy != nil {
		return *a.nanowatts_per_cubic_meterLazy
	}
	nanowatts_per_cubic_meter := a.convertFromBase(PowerDensityNanowattPerCubicMeter)
	a.nanowatts_per_cubic_meterLazy = &nanowatts_per_cubic_meter
	return nanowatts_per_cubic_meter
}

// MicrowattsPerCubicMeter returns the PowerDensity value in MicrowattsPerCubicMeter.
//
// 
func (a *PowerDensity) MicrowattsPerCubicMeter() float64 {
	if a.microwatts_per_cubic_meterLazy != nil {
		return *a.microwatts_per_cubic_meterLazy
	}
	microwatts_per_cubic_meter := a.convertFromBase(PowerDensityMicrowattPerCubicMeter)
	a.microwatts_per_cubic_meterLazy = &microwatts_per_cubic_meter
	return microwatts_per_cubic_meter
}

// MilliwattsPerCubicMeter returns the PowerDensity value in MilliwattsPerCubicMeter.
//
// 
func (a *PowerDensity) MilliwattsPerCubicMeter() float64 {
	if a.milliwatts_per_cubic_meterLazy != nil {
		return *a.milliwatts_per_cubic_meterLazy
	}
	milliwatts_per_cubic_meter := a.convertFromBase(PowerDensityMilliwattPerCubicMeter)
	a.milliwatts_per_cubic_meterLazy = &milliwatts_per_cubic_meter
	return milliwatts_per_cubic_meter
}

// DeciwattsPerCubicMeter returns the PowerDensity value in DeciwattsPerCubicMeter.
//
// 
func (a *PowerDensity) DeciwattsPerCubicMeter() float64 {
	if a.deciwatts_per_cubic_meterLazy != nil {
		return *a.deciwatts_per_cubic_meterLazy
	}
	deciwatts_per_cubic_meter := a.convertFromBase(PowerDensityDeciwattPerCubicMeter)
	a.deciwatts_per_cubic_meterLazy = &deciwatts_per_cubic_meter
	return deciwatts_per_cubic_meter
}

// DecawattsPerCubicMeter returns the PowerDensity value in DecawattsPerCubicMeter.
//
// 
func (a *PowerDensity) DecawattsPerCubicMeter() float64 {
	if a.decawatts_per_cubic_meterLazy != nil {
		return *a.decawatts_per_cubic_meterLazy
	}
	decawatts_per_cubic_meter := a.convertFromBase(PowerDensityDecawattPerCubicMeter)
	a.decawatts_per_cubic_meterLazy = &decawatts_per_cubic_meter
	return decawatts_per_cubic_meter
}

// KilowattsPerCubicMeter returns the PowerDensity value in KilowattsPerCubicMeter.
//
// 
func (a *PowerDensity) KilowattsPerCubicMeter() float64 {
	if a.kilowatts_per_cubic_meterLazy != nil {
		return *a.kilowatts_per_cubic_meterLazy
	}
	kilowatts_per_cubic_meter := a.convertFromBase(PowerDensityKilowattPerCubicMeter)
	a.kilowatts_per_cubic_meterLazy = &kilowatts_per_cubic_meter
	return kilowatts_per_cubic_meter
}

// MegawattsPerCubicMeter returns the PowerDensity value in MegawattsPerCubicMeter.
//
// 
func (a *PowerDensity) MegawattsPerCubicMeter() float64 {
	if a.megawatts_per_cubic_meterLazy != nil {
		return *a.megawatts_per_cubic_meterLazy
	}
	megawatts_per_cubic_meter := a.convertFromBase(PowerDensityMegawattPerCubicMeter)
	a.megawatts_per_cubic_meterLazy = &megawatts_per_cubic_meter
	return megawatts_per_cubic_meter
}

// GigawattsPerCubicMeter returns the PowerDensity value in GigawattsPerCubicMeter.
//
// 
func (a *PowerDensity) GigawattsPerCubicMeter() float64 {
	if a.gigawatts_per_cubic_meterLazy != nil {
		return *a.gigawatts_per_cubic_meterLazy
	}
	gigawatts_per_cubic_meter := a.convertFromBase(PowerDensityGigawattPerCubicMeter)
	a.gigawatts_per_cubic_meterLazy = &gigawatts_per_cubic_meter
	return gigawatts_per_cubic_meter
}

// TerawattsPerCubicMeter returns the PowerDensity value in TerawattsPerCubicMeter.
//
// 
func (a *PowerDensity) TerawattsPerCubicMeter() float64 {
	if a.terawatts_per_cubic_meterLazy != nil {
		return *a.terawatts_per_cubic_meterLazy
	}
	terawatts_per_cubic_meter := a.convertFromBase(PowerDensityTerawattPerCubicMeter)
	a.terawatts_per_cubic_meterLazy = &terawatts_per_cubic_meter
	return terawatts_per_cubic_meter
}

// PicowattsPerCubicInch returns the PowerDensity value in PicowattsPerCubicInch.
//
// 
func (a *PowerDensity) PicowattsPerCubicInch() float64 {
	if a.picowatts_per_cubic_inchLazy != nil {
		return *a.picowatts_per_cubic_inchLazy
	}
	picowatts_per_cubic_inch := a.convertFromBase(PowerDensityPicowattPerCubicInch)
	a.picowatts_per_cubic_inchLazy = &picowatts_per_cubic_inch
	return picowatts_per_cubic_inch
}

// NanowattsPerCubicInch returns the PowerDensity value in NanowattsPerCubicInch.
//
// 
func (a *PowerDensity) NanowattsPerCubicInch() float64 {
	if a.nanowatts_per_cubic_inchLazy != nil {
		return *a.nanowatts_per_cubic_inchLazy
	}
	nanowatts_per_cubic_inch := a.convertFromBase(PowerDensityNanowattPerCubicInch)
	a.nanowatts_per_cubic_inchLazy = &nanowatts_per_cubic_inch
	return nanowatts_per_cubic_inch
}

// MicrowattsPerCubicInch returns the PowerDensity value in MicrowattsPerCubicInch.
//
// 
func (a *PowerDensity) MicrowattsPerCubicInch() float64 {
	if a.microwatts_per_cubic_inchLazy != nil {
		return *a.microwatts_per_cubic_inchLazy
	}
	microwatts_per_cubic_inch := a.convertFromBase(PowerDensityMicrowattPerCubicInch)
	a.microwatts_per_cubic_inchLazy = &microwatts_per_cubic_inch
	return microwatts_per_cubic_inch
}

// MilliwattsPerCubicInch returns the PowerDensity value in MilliwattsPerCubicInch.
//
// 
func (a *PowerDensity) MilliwattsPerCubicInch() float64 {
	if a.milliwatts_per_cubic_inchLazy != nil {
		return *a.milliwatts_per_cubic_inchLazy
	}
	milliwatts_per_cubic_inch := a.convertFromBase(PowerDensityMilliwattPerCubicInch)
	a.milliwatts_per_cubic_inchLazy = &milliwatts_per_cubic_inch
	return milliwatts_per_cubic_inch
}

// DeciwattsPerCubicInch returns the PowerDensity value in DeciwattsPerCubicInch.
//
// 
func (a *PowerDensity) DeciwattsPerCubicInch() float64 {
	if a.deciwatts_per_cubic_inchLazy != nil {
		return *a.deciwatts_per_cubic_inchLazy
	}
	deciwatts_per_cubic_inch := a.convertFromBase(PowerDensityDeciwattPerCubicInch)
	a.deciwatts_per_cubic_inchLazy = &deciwatts_per_cubic_inch
	return deciwatts_per_cubic_inch
}

// DecawattsPerCubicInch returns the PowerDensity value in DecawattsPerCubicInch.
//
// 
func (a *PowerDensity) DecawattsPerCubicInch() float64 {
	if a.decawatts_per_cubic_inchLazy != nil {
		return *a.decawatts_per_cubic_inchLazy
	}
	decawatts_per_cubic_inch := a.convertFromBase(PowerDensityDecawattPerCubicInch)
	a.decawatts_per_cubic_inchLazy = &decawatts_per_cubic_inch
	return decawatts_per_cubic_inch
}

// KilowattsPerCubicInch returns the PowerDensity value in KilowattsPerCubicInch.
//
// 
func (a *PowerDensity) KilowattsPerCubicInch() float64 {
	if a.kilowatts_per_cubic_inchLazy != nil {
		return *a.kilowatts_per_cubic_inchLazy
	}
	kilowatts_per_cubic_inch := a.convertFromBase(PowerDensityKilowattPerCubicInch)
	a.kilowatts_per_cubic_inchLazy = &kilowatts_per_cubic_inch
	return kilowatts_per_cubic_inch
}

// MegawattsPerCubicInch returns the PowerDensity value in MegawattsPerCubicInch.
//
// 
func (a *PowerDensity) MegawattsPerCubicInch() float64 {
	if a.megawatts_per_cubic_inchLazy != nil {
		return *a.megawatts_per_cubic_inchLazy
	}
	megawatts_per_cubic_inch := a.convertFromBase(PowerDensityMegawattPerCubicInch)
	a.megawatts_per_cubic_inchLazy = &megawatts_per_cubic_inch
	return megawatts_per_cubic_inch
}

// GigawattsPerCubicInch returns the PowerDensity value in GigawattsPerCubicInch.
//
// 
func (a *PowerDensity) GigawattsPerCubicInch() float64 {
	if a.gigawatts_per_cubic_inchLazy != nil {
		return *a.gigawatts_per_cubic_inchLazy
	}
	gigawatts_per_cubic_inch := a.convertFromBase(PowerDensityGigawattPerCubicInch)
	a.gigawatts_per_cubic_inchLazy = &gigawatts_per_cubic_inch
	return gigawatts_per_cubic_inch
}

// TerawattsPerCubicInch returns the PowerDensity value in TerawattsPerCubicInch.
//
// 
func (a *PowerDensity) TerawattsPerCubicInch() float64 {
	if a.terawatts_per_cubic_inchLazy != nil {
		return *a.terawatts_per_cubic_inchLazy
	}
	terawatts_per_cubic_inch := a.convertFromBase(PowerDensityTerawattPerCubicInch)
	a.terawatts_per_cubic_inchLazy = &terawatts_per_cubic_inch
	return terawatts_per_cubic_inch
}

// PicowattsPerCubicFoot returns the PowerDensity value in PicowattsPerCubicFoot.
//
// 
func (a *PowerDensity) PicowattsPerCubicFoot() float64 {
	if a.picowatts_per_cubic_footLazy != nil {
		return *a.picowatts_per_cubic_footLazy
	}
	picowatts_per_cubic_foot := a.convertFromBase(PowerDensityPicowattPerCubicFoot)
	a.picowatts_per_cubic_footLazy = &picowatts_per_cubic_foot
	return picowatts_per_cubic_foot
}

// NanowattsPerCubicFoot returns the PowerDensity value in NanowattsPerCubicFoot.
//
// 
func (a *PowerDensity) NanowattsPerCubicFoot() float64 {
	if a.nanowatts_per_cubic_footLazy != nil {
		return *a.nanowatts_per_cubic_footLazy
	}
	nanowatts_per_cubic_foot := a.convertFromBase(PowerDensityNanowattPerCubicFoot)
	a.nanowatts_per_cubic_footLazy = &nanowatts_per_cubic_foot
	return nanowatts_per_cubic_foot
}

// MicrowattsPerCubicFoot returns the PowerDensity value in MicrowattsPerCubicFoot.
//
// 
func (a *PowerDensity) MicrowattsPerCubicFoot() float64 {
	if a.microwatts_per_cubic_footLazy != nil {
		return *a.microwatts_per_cubic_footLazy
	}
	microwatts_per_cubic_foot := a.convertFromBase(PowerDensityMicrowattPerCubicFoot)
	a.microwatts_per_cubic_footLazy = &microwatts_per_cubic_foot
	return microwatts_per_cubic_foot
}

// MilliwattsPerCubicFoot returns the PowerDensity value in MilliwattsPerCubicFoot.
//
// 
func (a *PowerDensity) MilliwattsPerCubicFoot() float64 {
	if a.milliwatts_per_cubic_footLazy != nil {
		return *a.milliwatts_per_cubic_footLazy
	}
	milliwatts_per_cubic_foot := a.convertFromBase(PowerDensityMilliwattPerCubicFoot)
	a.milliwatts_per_cubic_footLazy = &milliwatts_per_cubic_foot
	return milliwatts_per_cubic_foot
}

// DeciwattsPerCubicFoot returns the PowerDensity value in DeciwattsPerCubicFoot.
//
// 
func (a *PowerDensity) DeciwattsPerCubicFoot() float64 {
	if a.deciwatts_per_cubic_footLazy != nil {
		return *a.deciwatts_per_cubic_footLazy
	}
	deciwatts_per_cubic_foot := a.convertFromBase(PowerDensityDeciwattPerCubicFoot)
	a.deciwatts_per_cubic_footLazy = &deciwatts_per_cubic_foot
	return deciwatts_per_cubic_foot
}

// DecawattsPerCubicFoot returns the PowerDensity value in DecawattsPerCubicFoot.
//
// 
func (a *PowerDensity) DecawattsPerCubicFoot() float64 {
	if a.decawatts_per_cubic_footLazy != nil {
		return *a.decawatts_per_cubic_footLazy
	}
	decawatts_per_cubic_foot := a.convertFromBase(PowerDensityDecawattPerCubicFoot)
	a.decawatts_per_cubic_footLazy = &decawatts_per_cubic_foot
	return decawatts_per_cubic_foot
}

// KilowattsPerCubicFoot returns the PowerDensity value in KilowattsPerCubicFoot.
//
// 
func (a *PowerDensity) KilowattsPerCubicFoot() float64 {
	if a.kilowatts_per_cubic_footLazy != nil {
		return *a.kilowatts_per_cubic_footLazy
	}
	kilowatts_per_cubic_foot := a.convertFromBase(PowerDensityKilowattPerCubicFoot)
	a.kilowatts_per_cubic_footLazy = &kilowatts_per_cubic_foot
	return kilowatts_per_cubic_foot
}

// MegawattsPerCubicFoot returns the PowerDensity value in MegawattsPerCubicFoot.
//
// 
func (a *PowerDensity) MegawattsPerCubicFoot() float64 {
	if a.megawatts_per_cubic_footLazy != nil {
		return *a.megawatts_per_cubic_footLazy
	}
	megawatts_per_cubic_foot := a.convertFromBase(PowerDensityMegawattPerCubicFoot)
	a.megawatts_per_cubic_footLazy = &megawatts_per_cubic_foot
	return megawatts_per_cubic_foot
}

// GigawattsPerCubicFoot returns the PowerDensity value in GigawattsPerCubicFoot.
//
// 
func (a *PowerDensity) GigawattsPerCubicFoot() float64 {
	if a.gigawatts_per_cubic_footLazy != nil {
		return *a.gigawatts_per_cubic_footLazy
	}
	gigawatts_per_cubic_foot := a.convertFromBase(PowerDensityGigawattPerCubicFoot)
	a.gigawatts_per_cubic_footLazy = &gigawatts_per_cubic_foot
	return gigawatts_per_cubic_foot
}

// TerawattsPerCubicFoot returns the PowerDensity value in TerawattsPerCubicFoot.
//
// 
func (a *PowerDensity) TerawattsPerCubicFoot() float64 {
	if a.terawatts_per_cubic_footLazy != nil {
		return *a.terawatts_per_cubic_footLazy
	}
	terawatts_per_cubic_foot := a.convertFromBase(PowerDensityTerawattPerCubicFoot)
	a.terawatts_per_cubic_footLazy = &terawatts_per_cubic_foot
	return terawatts_per_cubic_foot
}

// PicowattsPerLiter returns the PowerDensity value in PicowattsPerLiter.
//
// 
func (a *PowerDensity) PicowattsPerLiter() float64 {
	if a.picowatts_per_literLazy != nil {
		return *a.picowatts_per_literLazy
	}
	picowatts_per_liter := a.convertFromBase(PowerDensityPicowattPerLiter)
	a.picowatts_per_literLazy = &picowatts_per_liter
	return picowatts_per_liter
}

// NanowattsPerLiter returns the PowerDensity value in NanowattsPerLiter.
//
// 
func (a *PowerDensity) NanowattsPerLiter() float64 {
	if a.nanowatts_per_literLazy != nil {
		return *a.nanowatts_per_literLazy
	}
	nanowatts_per_liter := a.convertFromBase(PowerDensityNanowattPerLiter)
	a.nanowatts_per_literLazy = &nanowatts_per_liter
	return nanowatts_per_liter
}

// MicrowattsPerLiter returns the PowerDensity value in MicrowattsPerLiter.
//
// 
func (a *PowerDensity) MicrowattsPerLiter() float64 {
	if a.microwatts_per_literLazy != nil {
		return *a.microwatts_per_literLazy
	}
	microwatts_per_liter := a.convertFromBase(PowerDensityMicrowattPerLiter)
	a.microwatts_per_literLazy = &microwatts_per_liter
	return microwatts_per_liter
}

// MilliwattsPerLiter returns the PowerDensity value in MilliwattsPerLiter.
//
// 
func (a *PowerDensity) MilliwattsPerLiter() float64 {
	if a.milliwatts_per_literLazy != nil {
		return *a.milliwatts_per_literLazy
	}
	milliwatts_per_liter := a.convertFromBase(PowerDensityMilliwattPerLiter)
	a.milliwatts_per_literLazy = &milliwatts_per_liter
	return milliwatts_per_liter
}

// DeciwattsPerLiter returns the PowerDensity value in DeciwattsPerLiter.
//
// 
func (a *PowerDensity) DeciwattsPerLiter() float64 {
	if a.deciwatts_per_literLazy != nil {
		return *a.deciwatts_per_literLazy
	}
	deciwatts_per_liter := a.convertFromBase(PowerDensityDeciwattPerLiter)
	a.deciwatts_per_literLazy = &deciwatts_per_liter
	return deciwatts_per_liter
}

// DecawattsPerLiter returns the PowerDensity value in DecawattsPerLiter.
//
// 
func (a *PowerDensity) DecawattsPerLiter() float64 {
	if a.decawatts_per_literLazy != nil {
		return *a.decawatts_per_literLazy
	}
	decawatts_per_liter := a.convertFromBase(PowerDensityDecawattPerLiter)
	a.decawatts_per_literLazy = &decawatts_per_liter
	return decawatts_per_liter
}

// KilowattsPerLiter returns the PowerDensity value in KilowattsPerLiter.
//
// 
func (a *PowerDensity) KilowattsPerLiter() float64 {
	if a.kilowatts_per_literLazy != nil {
		return *a.kilowatts_per_literLazy
	}
	kilowatts_per_liter := a.convertFromBase(PowerDensityKilowattPerLiter)
	a.kilowatts_per_literLazy = &kilowatts_per_liter
	return kilowatts_per_liter
}

// MegawattsPerLiter returns the PowerDensity value in MegawattsPerLiter.
//
// 
func (a *PowerDensity) MegawattsPerLiter() float64 {
	if a.megawatts_per_literLazy != nil {
		return *a.megawatts_per_literLazy
	}
	megawatts_per_liter := a.convertFromBase(PowerDensityMegawattPerLiter)
	a.megawatts_per_literLazy = &megawatts_per_liter
	return megawatts_per_liter
}

// GigawattsPerLiter returns the PowerDensity value in GigawattsPerLiter.
//
// 
func (a *PowerDensity) GigawattsPerLiter() float64 {
	if a.gigawatts_per_literLazy != nil {
		return *a.gigawatts_per_literLazy
	}
	gigawatts_per_liter := a.convertFromBase(PowerDensityGigawattPerLiter)
	a.gigawatts_per_literLazy = &gigawatts_per_liter
	return gigawatts_per_liter
}

// TerawattsPerLiter returns the PowerDensity value in TerawattsPerLiter.
//
// 
func (a *PowerDensity) TerawattsPerLiter() float64 {
	if a.terawatts_per_literLazy != nil {
		return *a.terawatts_per_literLazy
	}
	terawatts_per_liter := a.convertFromBase(PowerDensityTerawattPerLiter)
	a.terawatts_per_literLazy = &terawatts_per_liter
	return terawatts_per_liter
}


// ToDto creates a PowerDensityDto representation from the PowerDensity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by WattPerCubicMeter by default.
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

// ToDtoJSON creates a JSON representation of the PowerDensity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by WattPerCubicMeter by default.
func (a *PowerDensity) ToDtoJSON(holdInUnit *PowerDensityUnits) ([]byte, error) {
	// Convert to PowerDensityDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a PowerDensity to a specific unit value.
// The function uses the provided unit type (PowerDensityUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
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
		return math.NaN()
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

// String returns a string representation of the PowerDensity in the default unit (WattPerCubicMeter),
// formatted to two decimal places.
func (a PowerDensity) String() string {
	return a.ToString(PowerDensityWattPerCubicMeter, 2)
}

// ToString formats the PowerDensity value as a string with the specified unit and fractional digits.
// It converts the PowerDensity to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the PowerDensity value will be converted (e.g., WattPerCubicMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the PowerDensity with the unit abbreviation.
func (a *PowerDensity) ToString(unit PowerDensityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetPowerDensityAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetPowerDensityAbbreviation(unit))
}

// Equals checks if the given PowerDensity is equal to the current PowerDensity.
//
// Parameters:
//    other: The PowerDensity to compare against.
//
// Returns:
//    bool: Returns true if both PowerDensity are equal, false otherwise.
func (a *PowerDensity) Equals(other *PowerDensity) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current PowerDensity with another PowerDensity.
// It returns -1 if the current PowerDensity is less than the other PowerDensity, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The PowerDensity to compare against.
//
// Returns:
//    int: -1 if the current PowerDensity is less, 1 if greater, and 0 if equal.
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

// Add adds the given PowerDensity to the current PowerDensity and returns the result.
// The result is a new PowerDensity instance with the sum of the values.
//
// Parameters:
//    other: The PowerDensity to add to the current PowerDensity.
//
// Returns:
//    *PowerDensity: A new PowerDensity instance representing the sum of both PowerDensity.
func (a *PowerDensity) Add(other *PowerDensity) *PowerDensity {
	return &PowerDensity{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given PowerDensity from the current PowerDensity and returns the result.
// The result is a new PowerDensity instance with the difference of the values.
//
// Parameters:
//    other: The PowerDensity to subtract from the current PowerDensity.
//
// Returns:
//    *PowerDensity: A new PowerDensity instance representing the difference of both PowerDensity.
func (a *PowerDensity) Subtract(other *PowerDensity) *PowerDensity {
	return &PowerDensity{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current PowerDensity by the given PowerDensity and returns the result.
// The result is a new PowerDensity instance with the product of the values.
//
// Parameters:
//    other: The PowerDensity to multiply with the current PowerDensity.
//
// Returns:
//    *PowerDensity: A new PowerDensity instance representing the product of both PowerDensity.
func (a *PowerDensity) Multiply(other *PowerDensity) *PowerDensity {
	return &PowerDensity{value: a.value * other.BaseValue()}
}

// Divide divides the current PowerDensity by the given PowerDensity and returns the result.
// The result is a new PowerDensity instance with the quotient of the values.
//
// Parameters:
//    other: The PowerDensity to divide the current PowerDensity by.
//
// Returns:
//    *PowerDensity: A new PowerDensity instance representing the quotient of both PowerDensity.
func (a *PowerDensity) Divide(other *PowerDensity) *PowerDensity {
	return &PowerDensity{value: a.value / other.BaseValue()}
}

// GetPowerDensityAbbreviation gets the unit abbreviation.
func GetPowerDensityAbbreviation(unit PowerDensityUnits) string {
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