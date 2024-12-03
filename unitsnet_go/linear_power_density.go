package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// LinearPowerDensityUnits enumeration
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

// LinearPowerDensityDto represents an LinearPowerDensity
type LinearPowerDensityDto struct {
	Value float64
	Unit  LinearPowerDensityUnits
}

// LinearPowerDensityDtoFactory struct to group related functions
type LinearPowerDensityDtoFactory struct{}

func (udf LinearPowerDensityDtoFactory) FromJSON(data []byte) (*LinearPowerDensityDto, error) {
	a := LinearPowerDensityDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a LinearPowerDensityDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// LinearPowerDensity struct
type LinearPowerDensity struct {
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

// LinearPowerDensityFactory struct to group related functions
type LinearPowerDensityFactory struct{}

func (uf LinearPowerDensityFactory) CreateLinearPowerDensity(value float64, unit LinearPowerDensityUnits) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, unit)
}

func (uf LinearPowerDensityFactory) FromDto(dto LinearPowerDensityDto) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(dto.Value, dto.Unit)
}

func (uf LinearPowerDensityFactory) FromDtoJSON(data []byte) (*LinearPowerDensity, error) {
	unitDto, err := LinearPowerDensityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return LinearPowerDensityFactory{}.FromDto(*unitDto)
}


// FromWattPerMeter creates a new LinearPowerDensity instance from WattPerMeter.
func (uf LinearPowerDensityFactory) FromWattsPerMeter(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityWattPerMeter)
}

// FromWattPerCentimeter creates a new LinearPowerDensity instance from WattPerCentimeter.
func (uf LinearPowerDensityFactory) FromWattsPerCentimeter(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityWattPerCentimeter)
}

// FromWattPerMillimeter creates a new LinearPowerDensity instance from WattPerMillimeter.
func (uf LinearPowerDensityFactory) FromWattsPerMillimeter(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityWattPerMillimeter)
}

// FromWattPerInch creates a new LinearPowerDensity instance from WattPerInch.
func (uf LinearPowerDensityFactory) FromWattsPerInch(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityWattPerInch)
}

// FromWattPerFoot creates a new LinearPowerDensity instance from WattPerFoot.
func (uf LinearPowerDensityFactory) FromWattsPerFoot(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityWattPerFoot)
}

// FromMilliwattPerMeter creates a new LinearPowerDensity instance from MilliwattPerMeter.
func (uf LinearPowerDensityFactory) FromMilliwattsPerMeter(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityMilliwattPerMeter)
}

// FromKilowattPerMeter creates a new LinearPowerDensity instance from KilowattPerMeter.
func (uf LinearPowerDensityFactory) FromKilowattsPerMeter(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityKilowattPerMeter)
}

// FromMegawattPerMeter creates a new LinearPowerDensity instance from MegawattPerMeter.
func (uf LinearPowerDensityFactory) FromMegawattsPerMeter(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityMegawattPerMeter)
}

// FromGigawattPerMeter creates a new LinearPowerDensity instance from GigawattPerMeter.
func (uf LinearPowerDensityFactory) FromGigawattsPerMeter(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityGigawattPerMeter)
}

// FromMilliwattPerCentimeter creates a new LinearPowerDensity instance from MilliwattPerCentimeter.
func (uf LinearPowerDensityFactory) FromMilliwattsPerCentimeter(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityMilliwattPerCentimeter)
}

// FromKilowattPerCentimeter creates a new LinearPowerDensity instance from KilowattPerCentimeter.
func (uf LinearPowerDensityFactory) FromKilowattsPerCentimeter(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityKilowattPerCentimeter)
}

// FromMegawattPerCentimeter creates a new LinearPowerDensity instance from MegawattPerCentimeter.
func (uf LinearPowerDensityFactory) FromMegawattsPerCentimeter(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityMegawattPerCentimeter)
}

// FromGigawattPerCentimeter creates a new LinearPowerDensity instance from GigawattPerCentimeter.
func (uf LinearPowerDensityFactory) FromGigawattsPerCentimeter(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityGigawattPerCentimeter)
}

// FromMilliwattPerMillimeter creates a new LinearPowerDensity instance from MilliwattPerMillimeter.
func (uf LinearPowerDensityFactory) FromMilliwattsPerMillimeter(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityMilliwattPerMillimeter)
}

// FromKilowattPerMillimeter creates a new LinearPowerDensity instance from KilowattPerMillimeter.
func (uf LinearPowerDensityFactory) FromKilowattsPerMillimeter(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityKilowattPerMillimeter)
}

// FromMegawattPerMillimeter creates a new LinearPowerDensity instance from MegawattPerMillimeter.
func (uf LinearPowerDensityFactory) FromMegawattsPerMillimeter(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityMegawattPerMillimeter)
}

// FromGigawattPerMillimeter creates a new LinearPowerDensity instance from GigawattPerMillimeter.
func (uf LinearPowerDensityFactory) FromGigawattsPerMillimeter(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityGigawattPerMillimeter)
}

// FromMilliwattPerInch creates a new LinearPowerDensity instance from MilliwattPerInch.
func (uf LinearPowerDensityFactory) FromMilliwattsPerInch(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityMilliwattPerInch)
}

// FromKilowattPerInch creates a new LinearPowerDensity instance from KilowattPerInch.
func (uf LinearPowerDensityFactory) FromKilowattsPerInch(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityKilowattPerInch)
}

// FromMegawattPerInch creates a new LinearPowerDensity instance from MegawattPerInch.
func (uf LinearPowerDensityFactory) FromMegawattsPerInch(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityMegawattPerInch)
}

// FromGigawattPerInch creates a new LinearPowerDensity instance from GigawattPerInch.
func (uf LinearPowerDensityFactory) FromGigawattsPerInch(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityGigawattPerInch)
}

// FromMilliwattPerFoot creates a new LinearPowerDensity instance from MilliwattPerFoot.
func (uf LinearPowerDensityFactory) FromMilliwattsPerFoot(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityMilliwattPerFoot)
}

// FromKilowattPerFoot creates a new LinearPowerDensity instance from KilowattPerFoot.
func (uf LinearPowerDensityFactory) FromKilowattsPerFoot(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityKilowattPerFoot)
}

// FromMegawattPerFoot creates a new LinearPowerDensity instance from MegawattPerFoot.
func (uf LinearPowerDensityFactory) FromMegawattsPerFoot(value float64) (*LinearPowerDensity, error) {
	return newLinearPowerDensity(value, LinearPowerDensityMegawattPerFoot)
}

// FromGigawattPerFoot creates a new LinearPowerDensity instance from GigawattPerFoot.
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

// BaseValue returns the base value of LinearPowerDensity in WattPerMeter.
func (a *LinearPowerDensity) BaseValue() float64 {
	return a.value
}


// WattPerMeter returns the value in WattPerMeter.
func (a *LinearPowerDensity) WattsPerMeter() float64 {
	if a.watts_per_meterLazy != nil {
		return *a.watts_per_meterLazy
	}
	watts_per_meter := a.convertFromBase(LinearPowerDensityWattPerMeter)
	a.watts_per_meterLazy = &watts_per_meter
	return watts_per_meter
}

// WattPerCentimeter returns the value in WattPerCentimeter.
func (a *LinearPowerDensity) WattsPerCentimeter() float64 {
	if a.watts_per_centimeterLazy != nil {
		return *a.watts_per_centimeterLazy
	}
	watts_per_centimeter := a.convertFromBase(LinearPowerDensityWattPerCentimeter)
	a.watts_per_centimeterLazy = &watts_per_centimeter
	return watts_per_centimeter
}

// WattPerMillimeter returns the value in WattPerMillimeter.
func (a *LinearPowerDensity) WattsPerMillimeter() float64 {
	if a.watts_per_millimeterLazy != nil {
		return *a.watts_per_millimeterLazy
	}
	watts_per_millimeter := a.convertFromBase(LinearPowerDensityWattPerMillimeter)
	a.watts_per_millimeterLazy = &watts_per_millimeter
	return watts_per_millimeter
}

// WattPerInch returns the value in WattPerInch.
func (a *LinearPowerDensity) WattsPerInch() float64 {
	if a.watts_per_inchLazy != nil {
		return *a.watts_per_inchLazy
	}
	watts_per_inch := a.convertFromBase(LinearPowerDensityWattPerInch)
	a.watts_per_inchLazy = &watts_per_inch
	return watts_per_inch
}

// WattPerFoot returns the value in WattPerFoot.
func (a *LinearPowerDensity) WattsPerFoot() float64 {
	if a.watts_per_footLazy != nil {
		return *a.watts_per_footLazy
	}
	watts_per_foot := a.convertFromBase(LinearPowerDensityWattPerFoot)
	a.watts_per_footLazy = &watts_per_foot
	return watts_per_foot
}

// MilliwattPerMeter returns the value in MilliwattPerMeter.
func (a *LinearPowerDensity) MilliwattsPerMeter() float64 {
	if a.milliwatts_per_meterLazy != nil {
		return *a.milliwatts_per_meterLazy
	}
	milliwatts_per_meter := a.convertFromBase(LinearPowerDensityMilliwattPerMeter)
	a.milliwatts_per_meterLazy = &milliwatts_per_meter
	return milliwatts_per_meter
}

// KilowattPerMeter returns the value in KilowattPerMeter.
func (a *LinearPowerDensity) KilowattsPerMeter() float64 {
	if a.kilowatts_per_meterLazy != nil {
		return *a.kilowatts_per_meterLazy
	}
	kilowatts_per_meter := a.convertFromBase(LinearPowerDensityKilowattPerMeter)
	a.kilowatts_per_meterLazy = &kilowatts_per_meter
	return kilowatts_per_meter
}

// MegawattPerMeter returns the value in MegawattPerMeter.
func (a *LinearPowerDensity) MegawattsPerMeter() float64 {
	if a.megawatts_per_meterLazy != nil {
		return *a.megawatts_per_meterLazy
	}
	megawatts_per_meter := a.convertFromBase(LinearPowerDensityMegawattPerMeter)
	a.megawatts_per_meterLazy = &megawatts_per_meter
	return megawatts_per_meter
}

// GigawattPerMeter returns the value in GigawattPerMeter.
func (a *LinearPowerDensity) GigawattsPerMeter() float64 {
	if a.gigawatts_per_meterLazy != nil {
		return *a.gigawatts_per_meterLazy
	}
	gigawatts_per_meter := a.convertFromBase(LinearPowerDensityGigawattPerMeter)
	a.gigawatts_per_meterLazy = &gigawatts_per_meter
	return gigawatts_per_meter
}

// MilliwattPerCentimeter returns the value in MilliwattPerCentimeter.
func (a *LinearPowerDensity) MilliwattsPerCentimeter() float64 {
	if a.milliwatts_per_centimeterLazy != nil {
		return *a.milliwatts_per_centimeterLazy
	}
	milliwatts_per_centimeter := a.convertFromBase(LinearPowerDensityMilliwattPerCentimeter)
	a.milliwatts_per_centimeterLazy = &milliwatts_per_centimeter
	return milliwatts_per_centimeter
}

// KilowattPerCentimeter returns the value in KilowattPerCentimeter.
func (a *LinearPowerDensity) KilowattsPerCentimeter() float64 {
	if a.kilowatts_per_centimeterLazy != nil {
		return *a.kilowatts_per_centimeterLazy
	}
	kilowatts_per_centimeter := a.convertFromBase(LinearPowerDensityKilowattPerCentimeter)
	a.kilowatts_per_centimeterLazy = &kilowatts_per_centimeter
	return kilowatts_per_centimeter
}

// MegawattPerCentimeter returns the value in MegawattPerCentimeter.
func (a *LinearPowerDensity) MegawattsPerCentimeter() float64 {
	if a.megawatts_per_centimeterLazy != nil {
		return *a.megawatts_per_centimeterLazy
	}
	megawatts_per_centimeter := a.convertFromBase(LinearPowerDensityMegawattPerCentimeter)
	a.megawatts_per_centimeterLazy = &megawatts_per_centimeter
	return megawatts_per_centimeter
}

// GigawattPerCentimeter returns the value in GigawattPerCentimeter.
func (a *LinearPowerDensity) GigawattsPerCentimeter() float64 {
	if a.gigawatts_per_centimeterLazy != nil {
		return *a.gigawatts_per_centimeterLazy
	}
	gigawatts_per_centimeter := a.convertFromBase(LinearPowerDensityGigawattPerCentimeter)
	a.gigawatts_per_centimeterLazy = &gigawatts_per_centimeter
	return gigawatts_per_centimeter
}

// MilliwattPerMillimeter returns the value in MilliwattPerMillimeter.
func (a *LinearPowerDensity) MilliwattsPerMillimeter() float64 {
	if a.milliwatts_per_millimeterLazy != nil {
		return *a.milliwatts_per_millimeterLazy
	}
	milliwatts_per_millimeter := a.convertFromBase(LinearPowerDensityMilliwattPerMillimeter)
	a.milliwatts_per_millimeterLazy = &milliwatts_per_millimeter
	return milliwatts_per_millimeter
}

// KilowattPerMillimeter returns the value in KilowattPerMillimeter.
func (a *LinearPowerDensity) KilowattsPerMillimeter() float64 {
	if a.kilowatts_per_millimeterLazy != nil {
		return *a.kilowatts_per_millimeterLazy
	}
	kilowatts_per_millimeter := a.convertFromBase(LinearPowerDensityKilowattPerMillimeter)
	a.kilowatts_per_millimeterLazy = &kilowatts_per_millimeter
	return kilowatts_per_millimeter
}

// MegawattPerMillimeter returns the value in MegawattPerMillimeter.
func (a *LinearPowerDensity) MegawattsPerMillimeter() float64 {
	if a.megawatts_per_millimeterLazy != nil {
		return *a.megawatts_per_millimeterLazy
	}
	megawatts_per_millimeter := a.convertFromBase(LinearPowerDensityMegawattPerMillimeter)
	a.megawatts_per_millimeterLazy = &megawatts_per_millimeter
	return megawatts_per_millimeter
}

// GigawattPerMillimeter returns the value in GigawattPerMillimeter.
func (a *LinearPowerDensity) GigawattsPerMillimeter() float64 {
	if a.gigawatts_per_millimeterLazy != nil {
		return *a.gigawatts_per_millimeterLazy
	}
	gigawatts_per_millimeter := a.convertFromBase(LinearPowerDensityGigawattPerMillimeter)
	a.gigawatts_per_millimeterLazy = &gigawatts_per_millimeter
	return gigawatts_per_millimeter
}

// MilliwattPerInch returns the value in MilliwattPerInch.
func (a *LinearPowerDensity) MilliwattsPerInch() float64 {
	if a.milliwatts_per_inchLazy != nil {
		return *a.milliwatts_per_inchLazy
	}
	milliwatts_per_inch := a.convertFromBase(LinearPowerDensityMilliwattPerInch)
	a.milliwatts_per_inchLazy = &milliwatts_per_inch
	return milliwatts_per_inch
}

// KilowattPerInch returns the value in KilowattPerInch.
func (a *LinearPowerDensity) KilowattsPerInch() float64 {
	if a.kilowatts_per_inchLazy != nil {
		return *a.kilowatts_per_inchLazy
	}
	kilowatts_per_inch := a.convertFromBase(LinearPowerDensityKilowattPerInch)
	a.kilowatts_per_inchLazy = &kilowatts_per_inch
	return kilowatts_per_inch
}

// MegawattPerInch returns the value in MegawattPerInch.
func (a *LinearPowerDensity) MegawattsPerInch() float64 {
	if a.megawatts_per_inchLazy != nil {
		return *a.megawatts_per_inchLazy
	}
	megawatts_per_inch := a.convertFromBase(LinearPowerDensityMegawattPerInch)
	a.megawatts_per_inchLazy = &megawatts_per_inch
	return megawatts_per_inch
}

// GigawattPerInch returns the value in GigawattPerInch.
func (a *LinearPowerDensity) GigawattsPerInch() float64 {
	if a.gigawatts_per_inchLazy != nil {
		return *a.gigawatts_per_inchLazy
	}
	gigawatts_per_inch := a.convertFromBase(LinearPowerDensityGigawattPerInch)
	a.gigawatts_per_inchLazy = &gigawatts_per_inch
	return gigawatts_per_inch
}

// MilliwattPerFoot returns the value in MilliwattPerFoot.
func (a *LinearPowerDensity) MilliwattsPerFoot() float64 {
	if a.milliwatts_per_footLazy != nil {
		return *a.milliwatts_per_footLazy
	}
	milliwatts_per_foot := a.convertFromBase(LinearPowerDensityMilliwattPerFoot)
	a.milliwatts_per_footLazy = &milliwatts_per_foot
	return milliwatts_per_foot
}

// KilowattPerFoot returns the value in KilowattPerFoot.
func (a *LinearPowerDensity) KilowattsPerFoot() float64 {
	if a.kilowatts_per_footLazy != nil {
		return *a.kilowatts_per_footLazy
	}
	kilowatts_per_foot := a.convertFromBase(LinearPowerDensityKilowattPerFoot)
	a.kilowatts_per_footLazy = &kilowatts_per_foot
	return kilowatts_per_foot
}

// MegawattPerFoot returns the value in MegawattPerFoot.
func (a *LinearPowerDensity) MegawattsPerFoot() float64 {
	if a.megawatts_per_footLazy != nil {
		return *a.megawatts_per_footLazy
	}
	megawatts_per_foot := a.convertFromBase(LinearPowerDensityMegawattPerFoot)
	a.megawatts_per_footLazy = &megawatts_per_foot
	return megawatts_per_foot
}

// GigawattPerFoot returns the value in GigawattPerFoot.
func (a *LinearPowerDensity) GigawattsPerFoot() float64 {
	if a.gigawatts_per_footLazy != nil {
		return *a.gigawatts_per_footLazy
	}
	gigawatts_per_foot := a.convertFromBase(LinearPowerDensityGigawattPerFoot)
	a.gigawatts_per_footLazy = &gigawatts_per_foot
	return gigawatts_per_foot
}


// ToDto creates an LinearPowerDensityDto representation.
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

// ToDtoJSON creates an LinearPowerDensityDto representation.
func (a *LinearPowerDensity) ToDtoJSON(holdInUnit *LinearPowerDensityUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts LinearPowerDensity to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a LinearPowerDensity) String() string {
	return a.ToString(LinearPowerDensityWattPerMeter, 2)
}

// ToString formats the LinearPowerDensity to string.
// fractionalDigits -1 for not mention
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

// Check if the given LinearPowerDensity are equals to the current LinearPowerDensity
func (a *LinearPowerDensity) Equals(other *LinearPowerDensity) bool {
	return a.value == other.BaseValue()
}

// Check if the given LinearPowerDensity are equals to the current LinearPowerDensity
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

// Add the given LinearPowerDensity to the current LinearPowerDensity.
func (a *LinearPowerDensity) Add(other *LinearPowerDensity) *LinearPowerDensity {
	return &LinearPowerDensity{value: a.value + other.BaseValue()}
}

// Subtract the given LinearPowerDensity to the current LinearPowerDensity.
func (a *LinearPowerDensity) Subtract(other *LinearPowerDensity) *LinearPowerDensity {
	return &LinearPowerDensity{value: a.value - other.BaseValue()}
}

// Multiply the given LinearPowerDensity to the current LinearPowerDensity.
func (a *LinearPowerDensity) Multiply(other *LinearPowerDensity) *LinearPowerDensity {
	return &LinearPowerDensity{value: a.value * other.BaseValue()}
}

// Divide the given LinearPowerDensity to the current LinearPowerDensity.
func (a *LinearPowerDensity) Divide(other *LinearPowerDensity) *LinearPowerDensity {
	return &LinearPowerDensity{value: a.value / other.BaseValue()}
}