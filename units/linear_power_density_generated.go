// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// LinearPowerDensityUnits defines various units of LinearPowerDensity.
type LinearPowerDensityUnits string

const (
    
        // 
        LinearPowerDensityWattPerMeter LinearPowerDensityUnits = "WattPerMeter"
        // 
        LinearPowerDensityWattPerCentimeter LinearPowerDensityUnits = "WattPerCentimeter"
        // 
        LinearPowerDensityWattPerMillimeter LinearPowerDensityUnits = "WattPerMillimeter"
        // 
        LinearPowerDensityWattPerInch LinearPowerDensityUnits = "WattPerInch"
        // 
        LinearPowerDensityWattPerFoot LinearPowerDensityUnits = "WattPerFoot"
        // 
        LinearPowerDensityMilliwattPerMeter LinearPowerDensityUnits = "MilliwattPerMeter"
        // 
        LinearPowerDensityKilowattPerMeter LinearPowerDensityUnits = "KilowattPerMeter"
        // 
        LinearPowerDensityMegawattPerMeter LinearPowerDensityUnits = "MegawattPerMeter"
        // 
        LinearPowerDensityGigawattPerMeter LinearPowerDensityUnits = "GigawattPerMeter"
        // 
        LinearPowerDensityMilliwattPerCentimeter LinearPowerDensityUnits = "MilliwattPerCentimeter"
        // 
        LinearPowerDensityKilowattPerCentimeter LinearPowerDensityUnits = "KilowattPerCentimeter"
        // 
        LinearPowerDensityMegawattPerCentimeter LinearPowerDensityUnits = "MegawattPerCentimeter"
        // 
        LinearPowerDensityGigawattPerCentimeter LinearPowerDensityUnits = "GigawattPerCentimeter"
        // 
        LinearPowerDensityMilliwattPerMillimeter LinearPowerDensityUnits = "MilliwattPerMillimeter"
        // 
        LinearPowerDensityKilowattPerMillimeter LinearPowerDensityUnits = "KilowattPerMillimeter"
        // 
        LinearPowerDensityMegawattPerMillimeter LinearPowerDensityUnits = "MegawattPerMillimeter"
        // 
        LinearPowerDensityGigawattPerMillimeter LinearPowerDensityUnits = "GigawattPerMillimeter"
        // 
        LinearPowerDensityMilliwattPerInch LinearPowerDensityUnits = "MilliwattPerInch"
        // 
        LinearPowerDensityKilowattPerInch LinearPowerDensityUnits = "KilowattPerInch"
        // 
        LinearPowerDensityMegawattPerInch LinearPowerDensityUnits = "MegawattPerInch"
        // 
        LinearPowerDensityGigawattPerInch LinearPowerDensityUnits = "GigawattPerInch"
        // 
        LinearPowerDensityMilliwattPerFoot LinearPowerDensityUnits = "MilliwattPerFoot"
        // 
        LinearPowerDensityKilowattPerFoot LinearPowerDensityUnits = "KilowattPerFoot"
        // 
        LinearPowerDensityMegawattPerFoot LinearPowerDensityUnits = "MegawattPerFoot"
        // 
        LinearPowerDensityGigawattPerFoot LinearPowerDensityUnits = "GigawattPerFoot"
)

// LinearPowerDensityDto represents a LinearPowerDensity measurement with a numerical value and its corresponding unit.
type LinearPowerDensityDto struct {
    // Value is the numerical representation of the LinearPowerDensity.
	Value float64
    // Unit specifies the unit of measurement for the LinearPowerDensity, as defined in the LinearPowerDensityUnits enumeration.
	Unit  LinearPowerDensityUnits
}

// LinearPowerDensityDtoFactory groups methods for creating and serializing LinearPowerDensityDto objects.
type LinearPowerDensityDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a LinearPowerDensityDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf LinearPowerDensityDtoFactory) FromJSON(data []byte) (*LinearPowerDensityDto, error) {
	a := LinearPowerDensityDto{}

    // Parse JSON into LinearPowerDensityDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a LinearPowerDensityDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a LinearPowerDensityDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}


// LinearPowerDensity represents a measurement in a of LinearPowerDensity.
//
// The Linear Power Density of a substance is its power per unit length.  The term linear density is most often used when describing the characteristics of one-dimensional objects, although linear density can also be used to describe the density of a three-dimensional quantity along one particular dimension.
type LinearPowerDensity struct {
	// value is the base measurement stored internally.
	value       float64
    
    watts_per_meterLazy *float64 
    watts_per_centimeterLazy *float64 
    watts_per_millimeterLazy *float64 
    watts_per_inchLazy *float64 
    watts_per_footLazy *float64 
    milliwatts_per_meterLazy *float64 
    kilowatts_per_meterLazy *float64 
    megawatts_per_meterLazy *float64 
    gigawatts_per_meterLazy *float64 
    milliwatts_per_centimeterLazy *float64 
    kilowatts_per_centimeterLazy *float64 
    megawatts_per_centimeterLazy *float64 
    gigawatts_per_centimeterLazy *float64 
    milliwatts_per_millimeterLazy *float64 
    kilowatts_per_millimeterLazy *float64 
    megawatts_per_millimeterLazy *float64 
    gigawatts_per_millimeterLazy *float64 
    milliwatts_per_inchLazy *float64 
    kilowatts_per_inchLazy *float64 
    megawatts_per_inchLazy *float64 
    gigawatts_per_inchLazy *float64 
    milliwatts_per_footLazy *float64 
    kilowatts_per_footLazy *float64 
    megawatts_per_footLazy *float64 
    gigawatts_per_footLazy *float64 
}

// LinearPowerDensityFactory groups methods for creating LinearPowerDensity instances.
type LinearPowerDensityFactory struct{}

// CreateLinearPowerDensity creates a new LinearPowerDensity instance from the given value and unit.
func (uf LinearPowerDensityFactory) CreateLinearPowerDensity(value float64, unit LinearPowerDensityUnits) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, unit)
}

// FromDto converts a LinearPowerDensityDto to a LinearPowerDensity instance.
func (uf LinearPowerDensityFactory) FromDto(dto LinearPowerDensityDto) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a LinearPowerDensity instance.
func (uf LinearPowerDensityFactory) FromDtoJSON(data []byte) (*LinearPowerDensity, error) {
	unitDto, err := LinearPowerDensityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse LinearPowerDensityDto from JSON: %w", err)
	}
	return LinearPowerDensityFactory{}.FromDto(*unitDto)
}


// FromWattsPerMeter creates a new LinearPowerDensity instance from a value in WattsPerMeter.
func (uf LinearPowerDensityFactory) FromWattsPerMeter(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityWattPerMeter)
}

// FromWattsPerCentimeter creates a new LinearPowerDensity instance from a value in WattsPerCentimeter.
func (uf LinearPowerDensityFactory) FromWattsPerCentimeter(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityWattPerCentimeter)
}

// FromWattsPerMillimeter creates a new LinearPowerDensity instance from a value in WattsPerMillimeter.
func (uf LinearPowerDensityFactory) FromWattsPerMillimeter(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityWattPerMillimeter)
}

// FromWattsPerInch creates a new LinearPowerDensity instance from a value in WattsPerInch.
func (uf LinearPowerDensityFactory) FromWattsPerInch(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityWattPerInch)
}

// FromWattsPerFoot creates a new LinearPowerDensity instance from a value in WattsPerFoot.
func (uf LinearPowerDensityFactory) FromWattsPerFoot(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityWattPerFoot)
}

// FromMilliwattsPerMeter creates a new LinearPowerDensity instance from a value in MilliwattsPerMeter.
func (uf LinearPowerDensityFactory) FromMilliwattsPerMeter(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityMilliwattPerMeter)
}

// FromKilowattsPerMeter creates a new LinearPowerDensity instance from a value in KilowattsPerMeter.
func (uf LinearPowerDensityFactory) FromKilowattsPerMeter(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityKilowattPerMeter)
}

// FromMegawattsPerMeter creates a new LinearPowerDensity instance from a value in MegawattsPerMeter.
func (uf LinearPowerDensityFactory) FromMegawattsPerMeter(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityMegawattPerMeter)
}

// FromGigawattsPerMeter creates a new LinearPowerDensity instance from a value in GigawattsPerMeter.
func (uf LinearPowerDensityFactory) FromGigawattsPerMeter(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityGigawattPerMeter)
}

// FromMilliwattsPerCentimeter creates a new LinearPowerDensity instance from a value in MilliwattsPerCentimeter.
func (uf LinearPowerDensityFactory) FromMilliwattsPerCentimeter(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityMilliwattPerCentimeter)
}

// FromKilowattsPerCentimeter creates a new LinearPowerDensity instance from a value in KilowattsPerCentimeter.
func (uf LinearPowerDensityFactory) FromKilowattsPerCentimeter(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityKilowattPerCentimeter)
}

// FromMegawattsPerCentimeter creates a new LinearPowerDensity instance from a value in MegawattsPerCentimeter.
func (uf LinearPowerDensityFactory) FromMegawattsPerCentimeter(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityMegawattPerCentimeter)
}

// FromGigawattsPerCentimeter creates a new LinearPowerDensity instance from a value in GigawattsPerCentimeter.
func (uf LinearPowerDensityFactory) FromGigawattsPerCentimeter(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityGigawattPerCentimeter)
}

// FromMilliwattsPerMillimeter creates a new LinearPowerDensity instance from a value in MilliwattsPerMillimeter.
func (uf LinearPowerDensityFactory) FromMilliwattsPerMillimeter(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityMilliwattPerMillimeter)
}

// FromKilowattsPerMillimeter creates a new LinearPowerDensity instance from a value in KilowattsPerMillimeter.
func (uf LinearPowerDensityFactory) FromKilowattsPerMillimeter(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityKilowattPerMillimeter)
}

// FromMegawattsPerMillimeter creates a new LinearPowerDensity instance from a value in MegawattsPerMillimeter.
func (uf LinearPowerDensityFactory) FromMegawattsPerMillimeter(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityMegawattPerMillimeter)
}

// FromGigawattsPerMillimeter creates a new LinearPowerDensity instance from a value in GigawattsPerMillimeter.
func (uf LinearPowerDensityFactory) FromGigawattsPerMillimeter(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityGigawattPerMillimeter)
}

// FromMilliwattsPerInch creates a new LinearPowerDensity instance from a value in MilliwattsPerInch.
func (uf LinearPowerDensityFactory) FromMilliwattsPerInch(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityMilliwattPerInch)
}

// FromKilowattsPerInch creates a new LinearPowerDensity instance from a value in KilowattsPerInch.
func (uf LinearPowerDensityFactory) FromKilowattsPerInch(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityKilowattPerInch)
}

// FromMegawattsPerInch creates a new LinearPowerDensity instance from a value in MegawattsPerInch.
func (uf LinearPowerDensityFactory) FromMegawattsPerInch(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityMegawattPerInch)
}

// FromGigawattsPerInch creates a new LinearPowerDensity instance from a value in GigawattsPerInch.
func (uf LinearPowerDensityFactory) FromGigawattsPerInch(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityGigawattPerInch)
}

// FromMilliwattsPerFoot creates a new LinearPowerDensity instance from a value in MilliwattsPerFoot.
func (uf LinearPowerDensityFactory) FromMilliwattsPerFoot(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityMilliwattPerFoot)
}

// FromKilowattsPerFoot creates a new LinearPowerDensity instance from a value in KilowattsPerFoot.
func (uf LinearPowerDensityFactory) FromKilowattsPerFoot(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityKilowattPerFoot)
}

// FromMegawattsPerFoot creates a new LinearPowerDensity instance from a value in MegawattsPerFoot.
func (uf LinearPowerDensityFactory) FromMegawattsPerFoot(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityMegawattPerFoot)
}

// FromGigawattsPerFoot creates a new LinearPowerDensity instance from a value in GigawattsPerFoot.
func (uf LinearPowerDensityFactory) FromGigawattsPerFoot(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityGigawattPerFoot)
}


// newLinearPowerDensity creates a new LinearPowerDensity.
func newLinearPowerDensity(value float64, fromUnit LinearPowerDensityUnits) (*LinearPowerDensity, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &LinearPowerDensity{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of LinearPowerDensity in WattPerMeter unit (the base/default quantity).
func (a *LinearPowerDensity) BaseValue() float64 {
	return a.value
}


// WattsPerMeter returns the LinearPowerDensity value in WattsPerMeter.
//
// 
func (a *LinearPowerDensity) WattsPerMeter() float64 {
	if a.watts_per_meterLazy != nil {
		return *a.watts_per_meterLazy
	}
	watts_per_meter := a.convertFromBase(LinearPowerDensityWattPerMeter)
	a.watts_per_meterLazy = &watts_per_meter
	return watts_per_meter
}

// WattsPerCentimeter returns the LinearPowerDensity value in WattsPerCentimeter.
//
// 
func (a *LinearPowerDensity) WattsPerCentimeter() float64 {
	if a.watts_per_centimeterLazy != nil {
		return *a.watts_per_centimeterLazy
	}
	watts_per_centimeter := a.convertFromBase(LinearPowerDensityWattPerCentimeter)
	a.watts_per_centimeterLazy = &watts_per_centimeter
	return watts_per_centimeter
}

// WattsPerMillimeter returns the LinearPowerDensity value in WattsPerMillimeter.
//
// 
func (a *LinearPowerDensity) WattsPerMillimeter() float64 {
	if a.watts_per_millimeterLazy != nil {
		return *a.watts_per_millimeterLazy
	}
	watts_per_millimeter := a.convertFromBase(LinearPowerDensityWattPerMillimeter)
	a.watts_per_millimeterLazy = &watts_per_millimeter
	return watts_per_millimeter
}

// WattsPerInch returns the LinearPowerDensity value in WattsPerInch.
//
// 
func (a *LinearPowerDensity) WattsPerInch() float64 {
	if a.watts_per_inchLazy != nil {
		return *a.watts_per_inchLazy
	}
	watts_per_inch := a.convertFromBase(LinearPowerDensityWattPerInch)
	a.watts_per_inchLazy = &watts_per_inch
	return watts_per_inch
}

// WattsPerFoot returns the LinearPowerDensity value in WattsPerFoot.
//
// 
func (a *LinearPowerDensity) WattsPerFoot() float64 {
	if a.watts_per_footLazy != nil {
		return *a.watts_per_footLazy
	}
	watts_per_foot := a.convertFromBase(LinearPowerDensityWattPerFoot)
	a.watts_per_footLazy = &watts_per_foot
	return watts_per_foot
}

// MilliwattsPerMeter returns the LinearPowerDensity value in MilliwattsPerMeter.
//
// 
func (a *LinearPowerDensity) MilliwattsPerMeter() float64 {
	if a.milliwatts_per_meterLazy != nil {
		return *a.milliwatts_per_meterLazy
	}
	milliwatts_per_meter := a.convertFromBase(LinearPowerDensityMilliwattPerMeter)
	a.milliwatts_per_meterLazy = &milliwatts_per_meter
	return milliwatts_per_meter
}

// KilowattsPerMeter returns the LinearPowerDensity value in KilowattsPerMeter.
//
// 
func (a *LinearPowerDensity) KilowattsPerMeter() float64 {
	if a.kilowatts_per_meterLazy != nil {
		return *a.kilowatts_per_meterLazy
	}
	kilowatts_per_meter := a.convertFromBase(LinearPowerDensityKilowattPerMeter)
	a.kilowatts_per_meterLazy = &kilowatts_per_meter
	return kilowatts_per_meter
}

// MegawattsPerMeter returns the LinearPowerDensity value in MegawattsPerMeter.
//
// 
func (a *LinearPowerDensity) MegawattsPerMeter() float64 {
	if a.megawatts_per_meterLazy != nil {
		return *a.megawatts_per_meterLazy
	}
	megawatts_per_meter := a.convertFromBase(LinearPowerDensityMegawattPerMeter)
	a.megawatts_per_meterLazy = &megawatts_per_meter
	return megawatts_per_meter
}

// GigawattsPerMeter returns the LinearPowerDensity value in GigawattsPerMeter.
//
// 
func (a *LinearPowerDensity) GigawattsPerMeter() float64 {
	if a.gigawatts_per_meterLazy != nil {
		return *a.gigawatts_per_meterLazy
	}
	gigawatts_per_meter := a.convertFromBase(LinearPowerDensityGigawattPerMeter)
	a.gigawatts_per_meterLazy = &gigawatts_per_meter
	return gigawatts_per_meter
}

// MilliwattsPerCentimeter returns the LinearPowerDensity value in MilliwattsPerCentimeter.
//
// 
func (a *LinearPowerDensity) MilliwattsPerCentimeter() float64 {
	if a.milliwatts_per_centimeterLazy != nil {
		return *a.milliwatts_per_centimeterLazy
	}
	milliwatts_per_centimeter := a.convertFromBase(LinearPowerDensityMilliwattPerCentimeter)
	a.milliwatts_per_centimeterLazy = &milliwatts_per_centimeter
	return milliwatts_per_centimeter
}

// KilowattsPerCentimeter returns the LinearPowerDensity value in KilowattsPerCentimeter.
//
// 
func (a *LinearPowerDensity) KilowattsPerCentimeter() float64 {
	if a.kilowatts_per_centimeterLazy != nil {
		return *a.kilowatts_per_centimeterLazy
	}
	kilowatts_per_centimeter := a.convertFromBase(LinearPowerDensityKilowattPerCentimeter)
	a.kilowatts_per_centimeterLazy = &kilowatts_per_centimeter
	return kilowatts_per_centimeter
}

// MegawattsPerCentimeter returns the LinearPowerDensity value in MegawattsPerCentimeter.
//
// 
func (a *LinearPowerDensity) MegawattsPerCentimeter() float64 {
	if a.megawatts_per_centimeterLazy != nil {
		return *a.megawatts_per_centimeterLazy
	}
	megawatts_per_centimeter := a.convertFromBase(LinearPowerDensityMegawattPerCentimeter)
	a.megawatts_per_centimeterLazy = &megawatts_per_centimeter
	return megawatts_per_centimeter
}

// GigawattsPerCentimeter returns the LinearPowerDensity value in GigawattsPerCentimeter.
//
// 
func (a *LinearPowerDensity) GigawattsPerCentimeter() float64 {
	if a.gigawatts_per_centimeterLazy != nil {
		return *a.gigawatts_per_centimeterLazy
	}
	gigawatts_per_centimeter := a.convertFromBase(LinearPowerDensityGigawattPerCentimeter)
	a.gigawatts_per_centimeterLazy = &gigawatts_per_centimeter
	return gigawatts_per_centimeter
}

// MilliwattsPerMillimeter returns the LinearPowerDensity value in MilliwattsPerMillimeter.
//
// 
func (a *LinearPowerDensity) MilliwattsPerMillimeter() float64 {
	if a.milliwatts_per_millimeterLazy != nil {
		return *a.milliwatts_per_millimeterLazy
	}
	milliwatts_per_millimeter := a.convertFromBase(LinearPowerDensityMilliwattPerMillimeter)
	a.milliwatts_per_millimeterLazy = &milliwatts_per_millimeter
	return milliwatts_per_millimeter
}

// KilowattsPerMillimeter returns the LinearPowerDensity value in KilowattsPerMillimeter.
//
// 
func (a *LinearPowerDensity) KilowattsPerMillimeter() float64 {
	if a.kilowatts_per_millimeterLazy != nil {
		return *a.kilowatts_per_millimeterLazy
	}
	kilowatts_per_millimeter := a.convertFromBase(LinearPowerDensityKilowattPerMillimeter)
	a.kilowatts_per_millimeterLazy = &kilowatts_per_millimeter
	return kilowatts_per_millimeter
}

// MegawattsPerMillimeter returns the LinearPowerDensity value in MegawattsPerMillimeter.
//
// 
func (a *LinearPowerDensity) MegawattsPerMillimeter() float64 {
	if a.megawatts_per_millimeterLazy != nil {
		return *a.megawatts_per_millimeterLazy
	}
	megawatts_per_millimeter := a.convertFromBase(LinearPowerDensityMegawattPerMillimeter)
	a.megawatts_per_millimeterLazy = &megawatts_per_millimeter
	return megawatts_per_millimeter
}

// GigawattsPerMillimeter returns the LinearPowerDensity value in GigawattsPerMillimeter.
//
// 
func (a *LinearPowerDensity) GigawattsPerMillimeter() float64 {
	if a.gigawatts_per_millimeterLazy != nil {
		return *a.gigawatts_per_millimeterLazy
	}
	gigawatts_per_millimeter := a.convertFromBase(LinearPowerDensityGigawattPerMillimeter)
	a.gigawatts_per_millimeterLazy = &gigawatts_per_millimeter
	return gigawatts_per_millimeter
}

// MilliwattsPerInch returns the LinearPowerDensity value in MilliwattsPerInch.
//
// 
func (a *LinearPowerDensity) MilliwattsPerInch() float64 {
	if a.milliwatts_per_inchLazy != nil {
		return *a.milliwatts_per_inchLazy
	}
	milliwatts_per_inch := a.convertFromBase(LinearPowerDensityMilliwattPerInch)
	a.milliwatts_per_inchLazy = &milliwatts_per_inch
	return milliwatts_per_inch
}

// KilowattsPerInch returns the LinearPowerDensity value in KilowattsPerInch.
//
// 
func (a *LinearPowerDensity) KilowattsPerInch() float64 {
	if a.kilowatts_per_inchLazy != nil {
		return *a.kilowatts_per_inchLazy
	}
	kilowatts_per_inch := a.convertFromBase(LinearPowerDensityKilowattPerInch)
	a.kilowatts_per_inchLazy = &kilowatts_per_inch
	return kilowatts_per_inch
}

// MegawattsPerInch returns the LinearPowerDensity value in MegawattsPerInch.
//
// 
func (a *LinearPowerDensity) MegawattsPerInch() float64 {
	if a.megawatts_per_inchLazy != nil {
		return *a.megawatts_per_inchLazy
	}
	megawatts_per_inch := a.convertFromBase(LinearPowerDensityMegawattPerInch)
	a.megawatts_per_inchLazy = &megawatts_per_inch
	return megawatts_per_inch
}

// GigawattsPerInch returns the LinearPowerDensity value in GigawattsPerInch.
//
// 
func (a *LinearPowerDensity) GigawattsPerInch() float64 {
	if a.gigawatts_per_inchLazy != nil {
		return *a.gigawatts_per_inchLazy
	}
	gigawatts_per_inch := a.convertFromBase(LinearPowerDensityGigawattPerInch)
	a.gigawatts_per_inchLazy = &gigawatts_per_inch
	return gigawatts_per_inch
}

// MilliwattsPerFoot returns the LinearPowerDensity value in MilliwattsPerFoot.
//
// 
func (a *LinearPowerDensity) MilliwattsPerFoot() float64 {
	if a.milliwatts_per_footLazy != nil {
		return *a.milliwatts_per_footLazy
	}
	milliwatts_per_foot := a.convertFromBase(LinearPowerDensityMilliwattPerFoot)
	a.milliwatts_per_footLazy = &milliwatts_per_foot
	return milliwatts_per_foot
}

// KilowattsPerFoot returns the LinearPowerDensity value in KilowattsPerFoot.
//
// 
func (a *LinearPowerDensity) KilowattsPerFoot() float64 {
	if a.kilowatts_per_footLazy != nil {
		return *a.kilowatts_per_footLazy
	}
	kilowatts_per_foot := a.convertFromBase(LinearPowerDensityKilowattPerFoot)
	a.kilowatts_per_footLazy = &kilowatts_per_foot
	return kilowatts_per_foot
}

// MegawattsPerFoot returns the LinearPowerDensity value in MegawattsPerFoot.
//
// 
func (a *LinearPowerDensity) MegawattsPerFoot() float64 {
	if a.megawatts_per_footLazy != nil {
		return *a.megawatts_per_footLazy
	}
	megawatts_per_foot := a.convertFromBase(LinearPowerDensityMegawattPerFoot)
	a.megawatts_per_footLazy = &megawatts_per_foot
	return megawatts_per_foot
}

// GigawattsPerFoot returns the LinearPowerDensity value in GigawattsPerFoot.
//
// 
func (a *LinearPowerDensity) GigawattsPerFoot() float64 {
	if a.gigawatts_per_footLazy != nil {
		return *a.gigawatts_per_footLazy
	}
	gigawatts_per_foot := a.convertFromBase(LinearPowerDensityGigawattPerFoot)
	a.gigawatts_per_footLazy = &gigawatts_per_foot
	return gigawatts_per_foot
}


// ToDto creates a LinearPowerDensityDto representation from the LinearPowerDensity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by WattPerMeter by default.
func (a *LinearPowerDensity) ToDto(holdInUnit *LinearPowerDensityUnits) LinearPowerDensityDto {
	if holdInUnit == nil {
		defaultUnit := LinearPowerDensityWattPerMeter // Default value
		holdInUnit = &defaultUnit
	}

	return LinearPowerDensityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the LinearPowerDensity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by WattPerMeter by default.
func (a *LinearPowerDensity) ToDtoJSON(holdInUnit *LinearPowerDensityUnits) ([]byte, error) {
	// Convert to LinearPowerDensityDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a LinearPowerDensity to a specific unit value.
// The function uses the provided unit type (LinearPowerDensityUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *LinearPowerDensity) Convert(toUnit LinearPowerDensityUnits) float64 {
	switch toUnit { 
    case LinearPowerDensityWattPerMeter:
		return a.WattsPerMeter()
    case LinearPowerDensityWattPerCentimeter:
		return a.WattsPerCentimeter()
    case LinearPowerDensityWattPerMillimeter:
		return a.WattsPerMillimeter()
    case LinearPowerDensityWattPerInch:
		return a.WattsPerInch()
    case LinearPowerDensityWattPerFoot:
		return a.WattsPerFoot()
    case LinearPowerDensityMilliwattPerMeter:
		return a.MilliwattsPerMeter()
    case LinearPowerDensityKilowattPerMeter:
		return a.KilowattsPerMeter()
    case LinearPowerDensityMegawattPerMeter:
		return a.MegawattsPerMeter()
    case LinearPowerDensityGigawattPerMeter:
		return a.GigawattsPerMeter()
    case LinearPowerDensityMilliwattPerCentimeter:
		return a.MilliwattsPerCentimeter()
    case LinearPowerDensityKilowattPerCentimeter:
		return a.KilowattsPerCentimeter()
    case LinearPowerDensityMegawattPerCentimeter:
		return a.MegawattsPerCentimeter()
    case LinearPowerDensityGigawattPerCentimeter:
		return a.GigawattsPerCentimeter()
    case LinearPowerDensityMilliwattPerMillimeter:
		return a.MilliwattsPerMillimeter()
    case LinearPowerDensityKilowattPerMillimeter:
		return a.KilowattsPerMillimeter()
    case LinearPowerDensityMegawattPerMillimeter:
		return a.MegawattsPerMillimeter()
    case LinearPowerDensityGigawattPerMillimeter:
		return a.GigawattsPerMillimeter()
    case LinearPowerDensityMilliwattPerInch:
		return a.MilliwattsPerInch()
    case LinearPowerDensityKilowattPerInch:
		return a.KilowattsPerInch()
    case LinearPowerDensityMegawattPerInch:
		return a.MegawattsPerInch()
    case LinearPowerDensityGigawattPerInch:
		return a.GigawattsPerInch()
    case LinearPowerDensityMilliwattPerFoot:
		return a.MilliwattsPerFoot()
    case LinearPowerDensityKilowattPerFoot:
		return a.KilowattsPerFoot()
    case LinearPowerDensityMegawattPerFoot:
		return a.MegawattsPerFoot()
    case LinearPowerDensityGigawattPerFoot:
		return a.GigawattsPerFoot()
	default:
		return math.NaN()
	}
}

func (a *LinearPowerDensity) convertFromBase(toUnit LinearPowerDensityUnits) float64 {
    value := a.value
	switch toUnit { 
	case LinearPowerDensityWattPerMeter:
		return (value) 
	case LinearPowerDensityWattPerCentimeter:
		return (value / 1e2) 
	case LinearPowerDensityWattPerMillimeter:
		return (value / 1e3) 
	case LinearPowerDensityWattPerInch:
		return (value / 39.37007874) 
	case LinearPowerDensityWattPerFoot:
		return (value / 3.280839895) 
	case LinearPowerDensityMilliwattPerMeter:
		return ((value) / 0.001) 
	case LinearPowerDensityKilowattPerMeter:
		return ((value) / 1000.0) 
	case LinearPowerDensityMegawattPerMeter:
		return ((value) / 1000000.0) 
	case LinearPowerDensityGigawattPerMeter:
		return ((value) / 1000000000.0) 
	case LinearPowerDensityMilliwattPerCentimeter:
		return ((value / 1e2) / 0.001) 
	case LinearPowerDensityKilowattPerCentimeter:
		return ((value / 1e2) / 1000.0) 
	case LinearPowerDensityMegawattPerCentimeter:
		return ((value / 1e2) / 1000000.0) 
	case LinearPowerDensityGigawattPerCentimeter:
		return ((value / 1e2) / 1000000000.0) 
	case LinearPowerDensityMilliwattPerMillimeter:
		return ((value / 1e3) / 0.001) 
	case LinearPowerDensityKilowattPerMillimeter:
		return ((value / 1e3) / 1000.0) 
	case LinearPowerDensityMegawattPerMillimeter:
		return ((value / 1e3) / 1000000.0) 
	case LinearPowerDensityGigawattPerMillimeter:
		return ((value / 1e3) / 1000000000.0) 
	case LinearPowerDensityMilliwattPerInch:
		return ((value / 39.37007874) / 0.001) 
	case LinearPowerDensityKilowattPerInch:
		return ((value / 39.37007874) / 1000.0) 
	case LinearPowerDensityMegawattPerInch:
		return ((value / 39.37007874) / 1000000.0) 
	case LinearPowerDensityGigawattPerInch:
		return ((value / 39.37007874) / 1000000000.0) 
	case LinearPowerDensityMilliwattPerFoot:
		return ((value / 3.280839895) / 0.001) 
	case LinearPowerDensityKilowattPerFoot:
		return ((value / 3.280839895) / 1000.0) 
	case LinearPowerDensityMegawattPerFoot:
		return ((value / 3.280839895) / 1000000.0) 
	case LinearPowerDensityGigawattPerFoot:
		return ((value / 3.280839895) / 1000000000.0) 
	default:
		return math.NaN()
	}
}

func (a *LinearPowerDensity) convertToBase(value float64, fromUnit LinearPowerDensityUnits) float64 {
	switch fromUnit { 
	case LinearPowerDensityWattPerMeter:
		return (value) 
	case LinearPowerDensityWattPerCentimeter:
		return (value * 1e2) 
	case LinearPowerDensityWattPerMillimeter:
		return (value * 1e3) 
	case LinearPowerDensityWattPerInch:
		return (value * 39.37007874) 
	case LinearPowerDensityWattPerFoot:
		return (value * 3.280839895) 
	case LinearPowerDensityMilliwattPerMeter:
		return ((value) * 0.001) 
	case LinearPowerDensityKilowattPerMeter:
		return ((value) * 1000.0) 
	case LinearPowerDensityMegawattPerMeter:
		return ((value) * 1000000.0) 
	case LinearPowerDensityGigawattPerMeter:
		return ((value) * 1000000000.0) 
	case LinearPowerDensityMilliwattPerCentimeter:
		return ((value * 1e2) * 0.001) 
	case LinearPowerDensityKilowattPerCentimeter:
		return ((value * 1e2) * 1000.0) 
	case LinearPowerDensityMegawattPerCentimeter:
		return ((value * 1e2) * 1000000.0) 
	case LinearPowerDensityGigawattPerCentimeter:
		return ((value * 1e2) * 1000000000.0) 
	case LinearPowerDensityMilliwattPerMillimeter:
		return ((value * 1e3) * 0.001) 
	case LinearPowerDensityKilowattPerMillimeter:
		return ((value * 1e3) * 1000.0) 
	case LinearPowerDensityMegawattPerMillimeter:
		return ((value * 1e3) * 1000000.0) 
	case LinearPowerDensityGigawattPerMillimeter:
		return ((value * 1e3) * 1000000000.0) 
	case LinearPowerDensityMilliwattPerInch:
		return ((value * 39.37007874) * 0.001) 
	case LinearPowerDensityKilowattPerInch:
		return ((value * 39.37007874) * 1000.0) 
	case LinearPowerDensityMegawattPerInch:
		return ((value * 39.37007874) * 1000000.0) 
	case LinearPowerDensityGigawattPerInch:
		return ((value * 39.37007874) * 1000000000.0) 
	case LinearPowerDensityMilliwattPerFoot:
		return ((value * 3.280839895) * 0.001) 
	case LinearPowerDensityKilowattPerFoot:
		return ((value * 3.280839895) * 1000.0) 
	case LinearPowerDensityMegawattPerFoot:
		return ((value * 3.280839895) * 1000000.0) 
	case LinearPowerDensityGigawattPerFoot:
		return ((value * 3.280839895) * 1000000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the LinearPowerDensity in the default unit (WattPerMeter),
// formatted to two decimal places.
func (a LinearPowerDensity) String() string {
	return a.ToString(LinearPowerDensityWattPerMeter, 2)
}

// ToString formats the LinearPowerDensity value as a string with the specified unit and fractional digits.
// It converts the LinearPowerDensity to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the LinearPowerDensity value will be converted (e.g., WattPerMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the LinearPowerDensity with the unit abbreviation.
func (a *LinearPowerDensity) ToString(unit LinearPowerDensityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *LinearPowerDensity) getUnitAbbreviation(unit LinearPowerDensityUnits) string {
	switch unit { 
	case LinearPowerDensityWattPerMeter:
		return "W/m" 
	case LinearPowerDensityWattPerCentimeter:
		return "W/cm" 
	case LinearPowerDensityWattPerMillimeter:
		return "W/mm" 
	case LinearPowerDensityWattPerInch:
		return "W/in" 
	case LinearPowerDensityWattPerFoot:
		return "W/ft" 
	case LinearPowerDensityMilliwattPerMeter:
		return "mW/m" 
	case LinearPowerDensityKilowattPerMeter:
		return "kW/m" 
	case LinearPowerDensityMegawattPerMeter:
		return "MW/m" 
	case LinearPowerDensityGigawattPerMeter:
		return "GW/m" 
	case LinearPowerDensityMilliwattPerCentimeter:
		return "mW/cm" 
	case LinearPowerDensityKilowattPerCentimeter:
		return "kW/cm" 
	case LinearPowerDensityMegawattPerCentimeter:
		return "MW/cm" 
	case LinearPowerDensityGigawattPerCentimeter:
		return "GW/cm" 
	case LinearPowerDensityMilliwattPerMillimeter:
		return "mW/mm" 
	case LinearPowerDensityKilowattPerMillimeter:
		return "kW/mm" 
	case LinearPowerDensityMegawattPerMillimeter:
		return "MW/mm" 
	case LinearPowerDensityGigawattPerMillimeter:
		return "GW/mm" 
	case LinearPowerDensityMilliwattPerInch:
		return "mW/in" 
	case LinearPowerDensityKilowattPerInch:
		return "kW/in" 
	case LinearPowerDensityMegawattPerInch:
		return "MW/in" 
	case LinearPowerDensityGigawattPerInch:
		return "GW/in" 
	case LinearPowerDensityMilliwattPerFoot:
		return "mW/ft" 
	case LinearPowerDensityKilowattPerFoot:
		return "kW/ft" 
	case LinearPowerDensityMegawattPerFoot:
		return "MW/ft" 
	case LinearPowerDensityGigawattPerFoot:
		return "GW/ft" 
	default:
		return ""
	}
}

// Equals checks if the given LinearPowerDensity is equal to the current LinearPowerDensity.
//
// Parameters:
//    other: The LinearPowerDensity to compare against.
//
// Returns:
//    bool: Returns true if both LinearPowerDensity are equal, false otherwise.
func (a *LinearPowerDensity) Equals(other *LinearPowerDensity) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current LinearPowerDensity with another LinearPowerDensity.
// It returns -1 if the current LinearPowerDensity is less than the other LinearPowerDensity, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The LinearPowerDensity to compare against.
//
// Returns:
//    int: -1 if the current LinearPowerDensity is less, 1 if greater, and 0 if equal.
func (a *LinearPowerDensity) CompareTo(other *LinearPowerDensity) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given LinearPowerDensity to the current LinearPowerDensity and returns the result.
// The result is a new LinearPowerDensity instance with the sum of the values.
//
// Parameters:
//    other: The LinearPowerDensity to add to the current LinearPowerDensity.
//
// Returns:
//    *LinearPowerDensity: A new LinearPowerDensity instance representing the sum of both LinearPowerDensity.
func (a *LinearPowerDensity) Add(other *LinearPowerDensity) *LinearPowerDensity {
	return &LinearPowerDensity{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given LinearPowerDensity from the current LinearPowerDensity and returns the result.
// The result is a new LinearPowerDensity instance with the difference of the values.
//
// Parameters:
//    other: The LinearPowerDensity to subtract from the current LinearPowerDensity.
//
// Returns:
//    *LinearPowerDensity: A new LinearPowerDensity instance representing the difference of both LinearPowerDensity.
func (a *LinearPowerDensity) Subtract(other *LinearPowerDensity) *LinearPowerDensity {
	return &LinearPowerDensity{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current LinearPowerDensity by the given LinearPowerDensity and returns the result.
// The result is a new LinearPowerDensity instance with the product of the values.
//
// Parameters:
//    other: The LinearPowerDensity to multiply with the current LinearPowerDensity.
//
// Returns:
//    *LinearPowerDensity: A new LinearPowerDensity instance representing the product of both LinearPowerDensity.
func (a *LinearPowerDensity) Multiply(other *LinearPowerDensity) *LinearPowerDensity {
	return &LinearPowerDensity{value: a.value * other.BaseValue()}
}

// Divide divides the current LinearPowerDensity by the given LinearPowerDensity and returns the result.
// The result is a new LinearPowerDensity instance with the quotient of the values.
//
// Parameters:
//    other: The LinearPowerDensity to divide the current LinearPowerDensity by.
//
// Returns:
//    *LinearPowerDensity: A new LinearPowerDensity instance representing the quotient of both LinearPowerDensity.
func (a *LinearPowerDensity) Divide(other *LinearPowerDensity) *LinearPowerDensity {
	return &LinearPowerDensity{value: a.value / other.BaseValue()}
}