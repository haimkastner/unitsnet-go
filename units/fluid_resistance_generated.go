// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// FluidResistanceUnits defines various units of FluidResistance.
type FluidResistanceUnits string

const (
    
        // 
        FluidResistancePascalSecondPerLiter FluidResistanceUnits = "PascalSecondPerLiter"
        // 
        FluidResistancePascalMinutePerLiter FluidResistanceUnits = "PascalMinutePerLiter"
        // 
        FluidResistancePascalSecondPerMilliliter FluidResistanceUnits = "PascalSecondPerMilliliter"
        // 
        FluidResistancePascalMinutePerMilliliter FluidResistanceUnits = "PascalMinutePerMilliliter"
        // 
        FluidResistancePascalSecondPerCubicMeter FluidResistanceUnits = "PascalSecondPerCubicMeter"
        // 
        FluidResistancePascalMinutePerCubicMeter FluidResistanceUnits = "PascalMinutePerCubicMeter"
        // 
        FluidResistancePascalSecondPerCubicCentimeter FluidResistanceUnits = "PascalSecondPerCubicCentimeter"
        // 
        FluidResistancePascalMinutePerCubicCentimeter FluidResistanceUnits = "PascalMinutePerCubicCentimeter"
        // 
        FluidResistanceDyneSecondPerCentimeterToTheFifth FluidResistanceUnits = "DyneSecondPerCentimeterToTheFifth"
        // 
        FluidResistanceMillimeterMercurySecondPerLiter FluidResistanceUnits = "MillimeterMercurySecondPerLiter"
        // 
        FluidResistanceMillimeterMercuryMinutePerLiter FluidResistanceUnits = "MillimeterMercuryMinutePerLiter"
        // 
        FluidResistanceMillimeterMercurySecondPerMilliliter FluidResistanceUnits = "MillimeterMercurySecondPerMilliliter"
        // 
        FluidResistanceMillimeterMercuryMinutePerMilliliter FluidResistanceUnits = "MillimeterMercuryMinutePerMilliliter"
        // 
        FluidResistanceMillimeterMercurySecondPerCubicCentimeter FluidResistanceUnits = "MillimeterMercurySecondPerCubicCentimeter"
        // 
        FluidResistanceMillimeterMercuryMinutePerCubicCentimeter FluidResistanceUnits = "MillimeterMercuryMinutePerCubicCentimeter"
        // 
        FluidResistanceMillimeterMercurySecondPerCubicMeter FluidResistanceUnits = "MillimeterMercurySecondPerCubicMeter"
        // 
        FluidResistanceMillimeterMercuryMinutePerCubicMeter FluidResistanceUnits = "MillimeterMercuryMinutePerCubicMeter"
        // 
        FluidResistanceWoodUnit FluidResistanceUnits = "WoodUnit"
        // 
        FluidResistanceMegapascalSecondPerCubicMeter FluidResistanceUnits = "MegapascalSecondPerCubicMeter"
)

var internalFluidResistanceUnitsMap = map[FluidResistanceUnits]bool{
	
	FluidResistancePascalSecondPerLiter: true,
	FluidResistancePascalMinutePerLiter: true,
	FluidResistancePascalSecondPerMilliliter: true,
	FluidResistancePascalMinutePerMilliliter: true,
	FluidResistancePascalSecondPerCubicMeter: true,
	FluidResistancePascalMinutePerCubicMeter: true,
	FluidResistancePascalSecondPerCubicCentimeter: true,
	FluidResistancePascalMinutePerCubicCentimeter: true,
	FluidResistanceDyneSecondPerCentimeterToTheFifth: true,
	FluidResistanceMillimeterMercurySecondPerLiter: true,
	FluidResistanceMillimeterMercuryMinutePerLiter: true,
	FluidResistanceMillimeterMercurySecondPerMilliliter: true,
	FluidResistanceMillimeterMercuryMinutePerMilliliter: true,
	FluidResistanceMillimeterMercurySecondPerCubicCentimeter: true,
	FluidResistanceMillimeterMercuryMinutePerCubicCentimeter: true,
	FluidResistanceMillimeterMercurySecondPerCubicMeter: true,
	FluidResistanceMillimeterMercuryMinutePerCubicMeter: true,
	FluidResistanceWoodUnit: true,
	FluidResistanceMegapascalSecondPerCubicMeter: true,
}

// FluidResistanceDto represents a FluidResistance measurement with a numerical value and its corresponding unit.
type FluidResistanceDto struct {
    // Value is the numerical representation of the FluidResistance.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the FluidResistance, as defined in the FluidResistanceUnits enumeration.
	Unit  FluidResistanceUnits `json:"unit" validate:"required,oneof=PascalSecondPerLiter PascalMinutePerLiter PascalSecondPerMilliliter PascalMinutePerMilliliter PascalSecondPerCubicMeter PascalMinutePerCubicMeter PascalSecondPerCubicCentimeter PascalMinutePerCubicCentimeter DyneSecondPerCentimeterToTheFifth MillimeterMercurySecondPerLiter MillimeterMercuryMinutePerLiter MillimeterMercurySecondPerMilliliter MillimeterMercuryMinutePerMilliliter MillimeterMercurySecondPerCubicCentimeter MillimeterMercuryMinutePerCubicCentimeter MillimeterMercurySecondPerCubicMeter MillimeterMercuryMinutePerCubicMeter WoodUnit MegapascalSecondPerCubicMeter"`
}

// FluidResistanceDtoFactory groups methods for creating and serializing FluidResistanceDto objects.
type FluidResistanceDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a FluidResistanceDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf FluidResistanceDtoFactory) FromJSON(data []byte) (*FluidResistanceDto, error) {
	a := FluidResistanceDto{}

    // Parse JSON into FluidResistanceDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a FluidResistanceDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a FluidResistanceDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// FluidResistance represents a measurement in a of FluidResistance.
//
// Fluid Resistance is a force acting opposite to the relative motion of any object moving with respect to a surrounding fluid. Fluid Resistance is sometimes referred to as drag or fluid friction.
type FluidResistance struct {
	// value is the base measurement stored internally.
	value       float64
    
    pascal_seconds_per_literLazy *float64 
    pascal_minutes_per_literLazy *float64 
    pascal_seconds_per_milliliterLazy *float64 
    pascal_minutes_per_milliliterLazy *float64 
    pascal_seconds_per_cubic_meterLazy *float64 
    pascal_minutes_per_cubic_meterLazy *float64 
    pascal_seconds_per_cubic_centimeterLazy *float64 
    pascal_minutes_per_cubic_centimeterLazy *float64 
    dyne_seconds_per_centimeter_to_the_fifthLazy *float64 
    millimeter_mercury_seconds_per_literLazy *float64 
    millimeter_mercury_minutes_per_literLazy *float64 
    millimeter_mercury_seconds_per_milliliterLazy *float64 
    millimeter_mercury_minutes_per_milliliterLazy *float64 
    millimeter_mercury_seconds_per_cubic_centimeterLazy *float64 
    millimeter_mercury_minutes_per_cubic_centimeterLazy *float64 
    millimeter_mercury_seconds_per_cubic_meterLazy *float64 
    millimeter_mercury_minutes_per_cubic_meterLazy *float64 
    wood_unitsLazy *float64 
    megapascal_seconds_per_cubic_meterLazy *float64 
}

// FluidResistanceFactory groups methods for creating FluidResistance instances.
type FluidResistanceFactory struct{}

// CreateFluidResistance creates a new FluidResistance instance from the given value and unit.
func (uf FluidResistanceFactory) CreateFluidResistance(value float64, unit FluidResistanceUnits) (*FluidResistance, error) {
	return newFluidResistance(value, unit)
}

// FromDto converts a FluidResistanceDto to a FluidResistance instance.
func (uf FluidResistanceFactory) FromDto(dto FluidResistanceDto) (*FluidResistance, error) {
	return newFluidResistance(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a FluidResistance instance.
func (uf FluidResistanceFactory) FromDtoJSON(data []byte) (*FluidResistance, error) {
	unitDto, err := FluidResistanceDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse FluidResistanceDto from JSON: %w", err)
	}
	return FluidResistanceFactory{}.FromDto(*unitDto)
}


// FromPascalSecondsPerLiter creates a new FluidResistance instance from a value in PascalSecondsPerLiter.
func (uf FluidResistanceFactory) FromPascalSecondsPerLiter(value float64) (*FluidResistance, error) {
	return newFluidResistance(value, FluidResistancePascalSecondPerLiter)
}

// FromPascalMinutesPerLiter creates a new FluidResistance instance from a value in PascalMinutesPerLiter.
func (uf FluidResistanceFactory) FromPascalMinutesPerLiter(value float64) (*FluidResistance, error) {
	return newFluidResistance(value, FluidResistancePascalMinutePerLiter)
}

// FromPascalSecondsPerMilliliter creates a new FluidResistance instance from a value in PascalSecondsPerMilliliter.
func (uf FluidResistanceFactory) FromPascalSecondsPerMilliliter(value float64) (*FluidResistance, error) {
	return newFluidResistance(value, FluidResistancePascalSecondPerMilliliter)
}

// FromPascalMinutesPerMilliliter creates a new FluidResistance instance from a value in PascalMinutesPerMilliliter.
func (uf FluidResistanceFactory) FromPascalMinutesPerMilliliter(value float64) (*FluidResistance, error) {
	return newFluidResistance(value, FluidResistancePascalMinutePerMilliliter)
}

// FromPascalSecondsPerCubicMeter creates a new FluidResistance instance from a value in PascalSecondsPerCubicMeter.
func (uf FluidResistanceFactory) FromPascalSecondsPerCubicMeter(value float64) (*FluidResistance, error) {
	return newFluidResistance(value, FluidResistancePascalSecondPerCubicMeter)
}

// FromPascalMinutesPerCubicMeter creates a new FluidResistance instance from a value in PascalMinutesPerCubicMeter.
func (uf FluidResistanceFactory) FromPascalMinutesPerCubicMeter(value float64) (*FluidResistance, error) {
	return newFluidResistance(value, FluidResistancePascalMinutePerCubicMeter)
}

// FromPascalSecondsPerCubicCentimeter creates a new FluidResistance instance from a value in PascalSecondsPerCubicCentimeter.
func (uf FluidResistanceFactory) FromPascalSecondsPerCubicCentimeter(value float64) (*FluidResistance, error) {
	return newFluidResistance(value, FluidResistancePascalSecondPerCubicCentimeter)
}

// FromPascalMinutesPerCubicCentimeter creates a new FluidResistance instance from a value in PascalMinutesPerCubicCentimeter.
func (uf FluidResistanceFactory) FromPascalMinutesPerCubicCentimeter(value float64) (*FluidResistance, error) {
	return newFluidResistance(value, FluidResistancePascalMinutePerCubicCentimeter)
}

// FromDyneSecondsPerCentimeterToTheFifth creates a new FluidResistance instance from a value in DyneSecondsPerCentimeterToTheFifth.
func (uf FluidResistanceFactory) FromDyneSecondsPerCentimeterToTheFifth(value float64) (*FluidResistance, error) {
	return newFluidResistance(value, FluidResistanceDyneSecondPerCentimeterToTheFifth)
}

// FromMillimeterMercurySecondsPerLiter creates a new FluidResistance instance from a value in MillimeterMercurySecondsPerLiter.
func (uf FluidResistanceFactory) FromMillimeterMercurySecondsPerLiter(value float64) (*FluidResistance, error) {
	return newFluidResistance(value, FluidResistanceMillimeterMercurySecondPerLiter)
}

// FromMillimeterMercuryMinutesPerLiter creates a new FluidResistance instance from a value in MillimeterMercuryMinutesPerLiter.
func (uf FluidResistanceFactory) FromMillimeterMercuryMinutesPerLiter(value float64) (*FluidResistance, error) {
	return newFluidResistance(value, FluidResistanceMillimeterMercuryMinutePerLiter)
}

// FromMillimeterMercurySecondsPerMilliliter creates a new FluidResistance instance from a value in MillimeterMercurySecondsPerMilliliter.
func (uf FluidResistanceFactory) FromMillimeterMercurySecondsPerMilliliter(value float64) (*FluidResistance, error) {
	return newFluidResistance(value, FluidResistanceMillimeterMercurySecondPerMilliliter)
}

// FromMillimeterMercuryMinutesPerMilliliter creates a new FluidResistance instance from a value in MillimeterMercuryMinutesPerMilliliter.
func (uf FluidResistanceFactory) FromMillimeterMercuryMinutesPerMilliliter(value float64) (*FluidResistance, error) {
	return newFluidResistance(value, FluidResistanceMillimeterMercuryMinutePerMilliliter)
}

// FromMillimeterMercurySecondsPerCubicCentimeter creates a new FluidResistance instance from a value in MillimeterMercurySecondsPerCubicCentimeter.
func (uf FluidResistanceFactory) FromMillimeterMercurySecondsPerCubicCentimeter(value float64) (*FluidResistance, error) {
	return newFluidResistance(value, FluidResistanceMillimeterMercurySecondPerCubicCentimeter)
}

// FromMillimeterMercuryMinutesPerCubicCentimeter creates a new FluidResistance instance from a value in MillimeterMercuryMinutesPerCubicCentimeter.
func (uf FluidResistanceFactory) FromMillimeterMercuryMinutesPerCubicCentimeter(value float64) (*FluidResistance, error) {
	return newFluidResistance(value, FluidResistanceMillimeterMercuryMinutePerCubicCentimeter)
}

// FromMillimeterMercurySecondsPerCubicMeter creates a new FluidResistance instance from a value in MillimeterMercurySecondsPerCubicMeter.
func (uf FluidResistanceFactory) FromMillimeterMercurySecondsPerCubicMeter(value float64) (*FluidResistance, error) {
	return newFluidResistance(value, FluidResistanceMillimeterMercurySecondPerCubicMeter)
}

// FromMillimeterMercuryMinutesPerCubicMeter creates a new FluidResistance instance from a value in MillimeterMercuryMinutesPerCubicMeter.
func (uf FluidResistanceFactory) FromMillimeterMercuryMinutesPerCubicMeter(value float64) (*FluidResistance, error) {
	return newFluidResistance(value, FluidResistanceMillimeterMercuryMinutePerCubicMeter)
}

// FromWoodUnits creates a new FluidResistance instance from a value in WoodUnits.
func (uf FluidResistanceFactory) FromWoodUnits(value float64) (*FluidResistance, error) {
	return newFluidResistance(value, FluidResistanceWoodUnit)
}

// FromMegapascalSecondsPerCubicMeter creates a new FluidResistance instance from a value in MegapascalSecondsPerCubicMeter.
func (uf FluidResistanceFactory) FromMegapascalSecondsPerCubicMeter(value float64) (*FluidResistance, error) {
	return newFluidResistance(value, FluidResistanceMegapascalSecondPerCubicMeter)
}


// newFluidResistance creates a new FluidResistance.
func newFluidResistance(value float64, fromUnit FluidResistanceUnits) (*FluidResistance, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalFluidResistanceUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in FluidResistanceUnits", fromUnit)
	}
	a := &FluidResistance{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of FluidResistance in PascalSecondPerCubicMeter unit (the base/default quantity).
func (a *FluidResistance) BaseValue() float64 {
	return a.value
}


// PascalSecondsPerLiter returns the FluidResistance value in PascalSecondsPerLiter.
//
// 
func (a *FluidResistance) PascalSecondsPerLiter() float64 {
	if a.pascal_seconds_per_literLazy != nil {
		return *a.pascal_seconds_per_literLazy
	}
	pascal_seconds_per_liter := a.convertFromBase(FluidResistancePascalSecondPerLiter)
	a.pascal_seconds_per_literLazy = &pascal_seconds_per_liter
	return pascal_seconds_per_liter
}

// PascalMinutesPerLiter returns the FluidResistance value in PascalMinutesPerLiter.
//
// 
func (a *FluidResistance) PascalMinutesPerLiter() float64 {
	if a.pascal_minutes_per_literLazy != nil {
		return *a.pascal_minutes_per_literLazy
	}
	pascal_minutes_per_liter := a.convertFromBase(FluidResistancePascalMinutePerLiter)
	a.pascal_minutes_per_literLazy = &pascal_minutes_per_liter
	return pascal_minutes_per_liter
}

// PascalSecondsPerMilliliter returns the FluidResistance value in PascalSecondsPerMilliliter.
//
// 
func (a *FluidResistance) PascalSecondsPerMilliliter() float64 {
	if a.pascal_seconds_per_milliliterLazy != nil {
		return *a.pascal_seconds_per_milliliterLazy
	}
	pascal_seconds_per_milliliter := a.convertFromBase(FluidResistancePascalSecondPerMilliliter)
	a.pascal_seconds_per_milliliterLazy = &pascal_seconds_per_milliliter
	return pascal_seconds_per_milliliter
}

// PascalMinutesPerMilliliter returns the FluidResistance value in PascalMinutesPerMilliliter.
//
// 
func (a *FluidResistance) PascalMinutesPerMilliliter() float64 {
	if a.pascal_minutes_per_milliliterLazy != nil {
		return *a.pascal_minutes_per_milliliterLazy
	}
	pascal_minutes_per_milliliter := a.convertFromBase(FluidResistancePascalMinutePerMilliliter)
	a.pascal_minutes_per_milliliterLazy = &pascal_minutes_per_milliliter
	return pascal_minutes_per_milliliter
}

// PascalSecondsPerCubicMeter returns the FluidResistance value in PascalSecondsPerCubicMeter.
//
// 
func (a *FluidResistance) PascalSecondsPerCubicMeter() float64 {
	if a.pascal_seconds_per_cubic_meterLazy != nil {
		return *a.pascal_seconds_per_cubic_meterLazy
	}
	pascal_seconds_per_cubic_meter := a.convertFromBase(FluidResistancePascalSecondPerCubicMeter)
	a.pascal_seconds_per_cubic_meterLazy = &pascal_seconds_per_cubic_meter
	return pascal_seconds_per_cubic_meter
}

// PascalMinutesPerCubicMeter returns the FluidResistance value in PascalMinutesPerCubicMeter.
//
// 
func (a *FluidResistance) PascalMinutesPerCubicMeter() float64 {
	if a.pascal_minutes_per_cubic_meterLazy != nil {
		return *a.pascal_minutes_per_cubic_meterLazy
	}
	pascal_minutes_per_cubic_meter := a.convertFromBase(FluidResistancePascalMinutePerCubicMeter)
	a.pascal_minutes_per_cubic_meterLazy = &pascal_minutes_per_cubic_meter
	return pascal_minutes_per_cubic_meter
}

// PascalSecondsPerCubicCentimeter returns the FluidResistance value in PascalSecondsPerCubicCentimeter.
//
// 
func (a *FluidResistance) PascalSecondsPerCubicCentimeter() float64 {
	if a.pascal_seconds_per_cubic_centimeterLazy != nil {
		return *a.pascal_seconds_per_cubic_centimeterLazy
	}
	pascal_seconds_per_cubic_centimeter := a.convertFromBase(FluidResistancePascalSecondPerCubicCentimeter)
	a.pascal_seconds_per_cubic_centimeterLazy = &pascal_seconds_per_cubic_centimeter
	return pascal_seconds_per_cubic_centimeter
}

// PascalMinutesPerCubicCentimeter returns the FluidResistance value in PascalMinutesPerCubicCentimeter.
//
// 
func (a *FluidResistance) PascalMinutesPerCubicCentimeter() float64 {
	if a.pascal_minutes_per_cubic_centimeterLazy != nil {
		return *a.pascal_minutes_per_cubic_centimeterLazy
	}
	pascal_minutes_per_cubic_centimeter := a.convertFromBase(FluidResistancePascalMinutePerCubicCentimeter)
	a.pascal_minutes_per_cubic_centimeterLazy = &pascal_minutes_per_cubic_centimeter
	return pascal_minutes_per_cubic_centimeter
}

// DyneSecondsPerCentimeterToTheFifth returns the FluidResistance value in DyneSecondsPerCentimeterToTheFifth.
//
// 
func (a *FluidResistance) DyneSecondsPerCentimeterToTheFifth() float64 {
	if a.dyne_seconds_per_centimeter_to_the_fifthLazy != nil {
		return *a.dyne_seconds_per_centimeter_to_the_fifthLazy
	}
	dyne_seconds_per_centimeter_to_the_fifth := a.convertFromBase(FluidResistanceDyneSecondPerCentimeterToTheFifth)
	a.dyne_seconds_per_centimeter_to_the_fifthLazy = &dyne_seconds_per_centimeter_to_the_fifth
	return dyne_seconds_per_centimeter_to_the_fifth
}

// MillimeterMercurySecondsPerLiter returns the FluidResistance value in MillimeterMercurySecondsPerLiter.
//
// 
func (a *FluidResistance) MillimeterMercurySecondsPerLiter() float64 {
	if a.millimeter_mercury_seconds_per_literLazy != nil {
		return *a.millimeter_mercury_seconds_per_literLazy
	}
	millimeter_mercury_seconds_per_liter := a.convertFromBase(FluidResistanceMillimeterMercurySecondPerLiter)
	a.millimeter_mercury_seconds_per_literLazy = &millimeter_mercury_seconds_per_liter
	return millimeter_mercury_seconds_per_liter
}

// MillimeterMercuryMinutesPerLiter returns the FluidResistance value in MillimeterMercuryMinutesPerLiter.
//
// 
func (a *FluidResistance) MillimeterMercuryMinutesPerLiter() float64 {
	if a.millimeter_mercury_minutes_per_literLazy != nil {
		return *a.millimeter_mercury_minutes_per_literLazy
	}
	millimeter_mercury_minutes_per_liter := a.convertFromBase(FluidResistanceMillimeterMercuryMinutePerLiter)
	a.millimeter_mercury_minutes_per_literLazy = &millimeter_mercury_minutes_per_liter
	return millimeter_mercury_minutes_per_liter
}

// MillimeterMercurySecondsPerMilliliter returns the FluidResistance value in MillimeterMercurySecondsPerMilliliter.
//
// 
func (a *FluidResistance) MillimeterMercurySecondsPerMilliliter() float64 {
	if a.millimeter_mercury_seconds_per_milliliterLazy != nil {
		return *a.millimeter_mercury_seconds_per_milliliterLazy
	}
	millimeter_mercury_seconds_per_milliliter := a.convertFromBase(FluidResistanceMillimeterMercurySecondPerMilliliter)
	a.millimeter_mercury_seconds_per_milliliterLazy = &millimeter_mercury_seconds_per_milliliter
	return millimeter_mercury_seconds_per_milliliter
}

// MillimeterMercuryMinutesPerMilliliter returns the FluidResistance value in MillimeterMercuryMinutesPerMilliliter.
//
// 
func (a *FluidResistance) MillimeterMercuryMinutesPerMilliliter() float64 {
	if a.millimeter_mercury_minutes_per_milliliterLazy != nil {
		return *a.millimeter_mercury_minutes_per_milliliterLazy
	}
	millimeter_mercury_minutes_per_milliliter := a.convertFromBase(FluidResistanceMillimeterMercuryMinutePerMilliliter)
	a.millimeter_mercury_minutes_per_milliliterLazy = &millimeter_mercury_minutes_per_milliliter
	return millimeter_mercury_minutes_per_milliliter
}

// MillimeterMercurySecondsPerCubicCentimeter returns the FluidResistance value in MillimeterMercurySecondsPerCubicCentimeter.
//
// 
func (a *FluidResistance) MillimeterMercurySecondsPerCubicCentimeter() float64 {
	if a.millimeter_mercury_seconds_per_cubic_centimeterLazy != nil {
		return *a.millimeter_mercury_seconds_per_cubic_centimeterLazy
	}
	millimeter_mercury_seconds_per_cubic_centimeter := a.convertFromBase(FluidResistanceMillimeterMercurySecondPerCubicCentimeter)
	a.millimeter_mercury_seconds_per_cubic_centimeterLazy = &millimeter_mercury_seconds_per_cubic_centimeter
	return millimeter_mercury_seconds_per_cubic_centimeter
}

// MillimeterMercuryMinutesPerCubicCentimeter returns the FluidResistance value in MillimeterMercuryMinutesPerCubicCentimeter.
//
// 
func (a *FluidResistance) MillimeterMercuryMinutesPerCubicCentimeter() float64 {
	if a.millimeter_mercury_minutes_per_cubic_centimeterLazy != nil {
		return *a.millimeter_mercury_minutes_per_cubic_centimeterLazy
	}
	millimeter_mercury_minutes_per_cubic_centimeter := a.convertFromBase(FluidResistanceMillimeterMercuryMinutePerCubicCentimeter)
	a.millimeter_mercury_minutes_per_cubic_centimeterLazy = &millimeter_mercury_minutes_per_cubic_centimeter
	return millimeter_mercury_minutes_per_cubic_centimeter
}

// MillimeterMercurySecondsPerCubicMeter returns the FluidResistance value in MillimeterMercurySecondsPerCubicMeter.
//
// 
func (a *FluidResistance) MillimeterMercurySecondsPerCubicMeter() float64 {
	if a.millimeter_mercury_seconds_per_cubic_meterLazy != nil {
		return *a.millimeter_mercury_seconds_per_cubic_meterLazy
	}
	millimeter_mercury_seconds_per_cubic_meter := a.convertFromBase(FluidResistanceMillimeterMercurySecondPerCubicMeter)
	a.millimeter_mercury_seconds_per_cubic_meterLazy = &millimeter_mercury_seconds_per_cubic_meter
	return millimeter_mercury_seconds_per_cubic_meter
}

// MillimeterMercuryMinutesPerCubicMeter returns the FluidResistance value in MillimeterMercuryMinutesPerCubicMeter.
//
// 
func (a *FluidResistance) MillimeterMercuryMinutesPerCubicMeter() float64 {
	if a.millimeter_mercury_minutes_per_cubic_meterLazy != nil {
		return *a.millimeter_mercury_minutes_per_cubic_meterLazy
	}
	millimeter_mercury_minutes_per_cubic_meter := a.convertFromBase(FluidResistanceMillimeterMercuryMinutePerCubicMeter)
	a.millimeter_mercury_minutes_per_cubic_meterLazy = &millimeter_mercury_minutes_per_cubic_meter
	return millimeter_mercury_minutes_per_cubic_meter
}

// WoodUnits returns the FluidResistance value in WoodUnits.
//
// 
func (a *FluidResistance) WoodUnits() float64 {
	if a.wood_unitsLazy != nil {
		return *a.wood_unitsLazy
	}
	wood_units := a.convertFromBase(FluidResistanceWoodUnit)
	a.wood_unitsLazy = &wood_units
	return wood_units
}

// MegapascalSecondsPerCubicMeter returns the FluidResistance value in MegapascalSecondsPerCubicMeter.
//
// 
func (a *FluidResistance) MegapascalSecondsPerCubicMeter() float64 {
	if a.megapascal_seconds_per_cubic_meterLazy != nil {
		return *a.megapascal_seconds_per_cubic_meterLazy
	}
	megapascal_seconds_per_cubic_meter := a.convertFromBase(FluidResistanceMegapascalSecondPerCubicMeter)
	a.megapascal_seconds_per_cubic_meterLazy = &megapascal_seconds_per_cubic_meter
	return megapascal_seconds_per_cubic_meter
}


// ToDto creates a FluidResistanceDto representation from the FluidResistance instance.
//
// If the provided holdInUnit is nil, the value will be repesented by PascalSecondPerCubicMeter by default.
func (a *FluidResistance) ToDto(holdInUnit *FluidResistanceUnits) FluidResistanceDto {
	if holdInUnit == nil {
		defaultUnit := FluidResistancePascalSecondPerCubicMeter // Default value
		holdInUnit = &defaultUnit
	}

	return FluidResistanceDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the FluidResistance instance.
//
// If the provided holdInUnit is nil, the value will be repesented by PascalSecondPerCubicMeter by default.
func (a *FluidResistance) ToDtoJSON(holdInUnit *FluidResistanceUnits) ([]byte, error) {
	// Convert to FluidResistanceDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a FluidResistance to a specific unit value.
// The function uses the provided unit type (FluidResistanceUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *FluidResistance) Convert(toUnit FluidResistanceUnits) float64 {
	switch toUnit { 
    case FluidResistancePascalSecondPerLiter:
		return a.PascalSecondsPerLiter()
    case FluidResistancePascalMinutePerLiter:
		return a.PascalMinutesPerLiter()
    case FluidResistancePascalSecondPerMilliliter:
		return a.PascalSecondsPerMilliliter()
    case FluidResistancePascalMinutePerMilliliter:
		return a.PascalMinutesPerMilliliter()
    case FluidResistancePascalSecondPerCubicMeter:
		return a.PascalSecondsPerCubicMeter()
    case FluidResistancePascalMinutePerCubicMeter:
		return a.PascalMinutesPerCubicMeter()
    case FluidResistancePascalSecondPerCubicCentimeter:
		return a.PascalSecondsPerCubicCentimeter()
    case FluidResistancePascalMinutePerCubicCentimeter:
		return a.PascalMinutesPerCubicCentimeter()
    case FluidResistanceDyneSecondPerCentimeterToTheFifth:
		return a.DyneSecondsPerCentimeterToTheFifth()
    case FluidResistanceMillimeterMercurySecondPerLiter:
		return a.MillimeterMercurySecondsPerLiter()
    case FluidResistanceMillimeterMercuryMinutePerLiter:
		return a.MillimeterMercuryMinutesPerLiter()
    case FluidResistanceMillimeterMercurySecondPerMilliliter:
		return a.MillimeterMercurySecondsPerMilliliter()
    case FluidResistanceMillimeterMercuryMinutePerMilliliter:
		return a.MillimeterMercuryMinutesPerMilliliter()
    case FluidResistanceMillimeterMercurySecondPerCubicCentimeter:
		return a.MillimeterMercurySecondsPerCubicCentimeter()
    case FluidResistanceMillimeterMercuryMinutePerCubicCentimeter:
		return a.MillimeterMercuryMinutesPerCubicCentimeter()
    case FluidResistanceMillimeterMercurySecondPerCubicMeter:
		return a.MillimeterMercurySecondsPerCubicMeter()
    case FluidResistanceMillimeterMercuryMinutePerCubicMeter:
		return a.MillimeterMercuryMinutesPerCubicMeter()
    case FluidResistanceWoodUnit:
		return a.WoodUnits()
    case FluidResistanceMegapascalSecondPerCubicMeter:
		return a.MegapascalSecondsPerCubicMeter()
	default:
		return math.NaN()
	}
}

func (a *FluidResistance) convertFromBase(toUnit FluidResistanceUnits) float64 {
    value := a.value
	switch toUnit { 
	case FluidResistancePascalSecondPerLiter:
		return (value / 1e3) 
	case FluidResistancePascalMinutePerLiter:
		return (value / 6e4) 
	case FluidResistancePascalSecondPerMilliliter:
		return (value / 1e6) 
	case FluidResistancePascalMinutePerMilliliter:
		return (value / 6e7) 
	case FluidResistancePascalSecondPerCubicMeter:
		return (value) 
	case FluidResistancePascalMinutePerCubicMeter:
		return (value / 60) 
	case FluidResistancePascalSecondPerCubicCentimeter:
		return (value / 1e6) 
	case FluidResistancePascalMinutePerCubicCentimeter:
		return (value / 6e7) 
	case FluidResistanceDyneSecondPerCentimeterToTheFifth:
		return (value / 1e5) 
	case FluidResistanceMillimeterMercurySecondPerLiter:
		return (value / 1.33322368e5) 
	case FluidResistanceMillimeterMercuryMinutePerLiter:
		return (value / 7.99934208e6) 
	case FluidResistanceMillimeterMercurySecondPerMilliliter:
		return (value / 1.33322368e8) 
	case FluidResistanceMillimeterMercuryMinutePerMilliliter:
		return (value / 7.99934208e9) 
	case FluidResistanceMillimeterMercurySecondPerCubicCentimeter:
		return (value / 1.33322368e8) 
	case FluidResistanceMillimeterMercuryMinutePerCubicCentimeter:
		return (value / 7.99934208e9) 
	case FluidResistanceMillimeterMercurySecondPerCubicMeter:
		return (value / 133.322368) 
	case FluidResistanceMillimeterMercuryMinutePerCubicMeter:
		return (value / 7.99934208e3) 
	case FluidResistanceWoodUnit:
		return (value / 7.99934208e6) 
	case FluidResistanceMegapascalSecondPerCubicMeter:
		return ((value) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *FluidResistance) convertToBase(value float64, fromUnit FluidResistanceUnits) float64 {
	switch fromUnit { 
	case FluidResistancePascalSecondPerLiter:
		return (value * 1e3) 
	case FluidResistancePascalMinutePerLiter:
		return (value * 6e4) 
	case FluidResistancePascalSecondPerMilliliter:
		return (value * 1e6) 
	case FluidResistancePascalMinutePerMilliliter:
		return (value * 6e7) 
	case FluidResistancePascalSecondPerCubicMeter:
		return (value) 
	case FluidResistancePascalMinutePerCubicMeter:
		return (value * 60) 
	case FluidResistancePascalSecondPerCubicCentimeter:
		return (value * 1e6) 
	case FluidResistancePascalMinutePerCubicCentimeter:
		return (value * 6e7) 
	case FluidResistanceDyneSecondPerCentimeterToTheFifth:
		return (value * 1e5) 
	case FluidResistanceMillimeterMercurySecondPerLiter:
		return (value * 1.33322368e5) 
	case FluidResistanceMillimeterMercuryMinutePerLiter:
		return (value * 7.99934208e6) 
	case FluidResistanceMillimeterMercurySecondPerMilliliter:
		return (value * 1.33322368e8) 
	case FluidResistanceMillimeterMercuryMinutePerMilliliter:
		return (value * 7.99934208e9) 
	case FluidResistanceMillimeterMercurySecondPerCubicCentimeter:
		return (value * 1.33322368e8) 
	case FluidResistanceMillimeterMercuryMinutePerCubicCentimeter:
		return (value * 7.99934208e9) 
	case FluidResistanceMillimeterMercurySecondPerCubicMeter:
		return (value * 133.322368) 
	case FluidResistanceMillimeterMercuryMinutePerCubicMeter:
		return (value * 7.99934208e3) 
	case FluidResistanceWoodUnit:
		return (value * 7.99934208e6) 
	case FluidResistanceMegapascalSecondPerCubicMeter:
		return ((value) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the FluidResistance in the default unit (PascalSecondPerCubicMeter),
// formatted to two decimal places.
func (a FluidResistance) String() string {
	return a.ToString(FluidResistancePascalSecondPerCubicMeter, 2)
}

// ToString formats the FluidResistance value as a string with the specified unit and fractional digits.
// It converts the FluidResistance to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the FluidResistance value will be converted (e.g., PascalSecondPerCubicMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the FluidResistance with the unit abbreviation.
func (a *FluidResistance) ToString(unit FluidResistanceUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetFluidResistanceAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetFluidResistanceAbbreviation(unit))
}

// Equals checks if the given FluidResistance is equal to the current FluidResistance.
//
// Parameters:
//    other: The FluidResistance to compare against.
//
// Returns:
//    bool: Returns true if both FluidResistance are equal, false otherwise.
func (a *FluidResistance) Equals(other *FluidResistance) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current FluidResistance with another FluidResistance.
// It returns -1 if the current FluidResistance is less than the other FluidResistance, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The FluidResistance to compare against.
//
// Returns:
//    int: -1 if the current FluidResistance is less, 1 if greater, and 0 if equal.
func (a *FluidResistance) CompareTo(other *FluidResistance) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given FluidResistance to the current FluidResistance and returns the result.
// The result is a new FluidResistance instance with the sum of the values.
//
// Parameters:
//    other: The FluidResistance to add to the current FluidResistance.
//
// Returns:
//    *FluidResistance: A new FluidResistance instance representing the sum of both FluidResistance.
func (a *FluidResistance) Add(other *FluidResistance) *FluidResistance {
	return &FluidResistance{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given FluidResistance from the current FluidResistance and returns the result.
// The result is a new FluidResistance instance with the difference of the values.
//
// Parameters:
//    other: The FluidResistance to subtract from the current FluidResistance.
//
// Returns:
//    *FluidResistance: A new FluidResistance instance representing the difference of both FluidResistance.
func (a *FluidResistance) Subtract(other *FluidResistance) *FluidResistance {
	return &FluidResistance{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current FluidResistance by the given FluidResistance and returns the result.
// The result is a new FluidResistance instance with the product of the values.
//
// Parameters:
//    other: The FluidResistance to multiply with the current FluidResistance.
//
// Returns:
//    *FluidResistance: A new FluidResistance instance representing the product of both FluidResistance.
func (a *FluidResistance) Multiply(other *FluidResistance) *FluidResistance {
	return &FluidResistance{value: a.value * other.BaseValue()}
}

// Divide divides the current FluidResistance by the given FluidResistance and returns the result.
// The result is a new FluidResistance instance with the quotient of the values.
//
// Parameters:
//    other: The FluidResistance to divide the current FluidResistance by.
//
// Returns:
//    *FluidResistance: A new FluidResistance instance representing the quotient of both FluidResistance.
func (a *FluidResistance) Divide(other *FluidResistance) *FluidResistance {
	return &FluidResistance{value: a.value / other.BaseValue()}
}

// GetFluidResistanceAbbreviation gets the unit abbreviation.
func GetFluidResistanceAbbreviation(unit FluidResistanceUnits) string {
	switch unit { 
	case FluidResistancePascalSecondPerLiter:
		return "Pa·s/l" 
	case FluidResistancePascalMinutePerLiter:
		return "Pa·min/l" 
	case FluidResistancePascalSecondPerMilliliter:
		return "Pa·s/ml" 
	case FluidResistancePascalMinutePerMilliliter:
		return "Pa·min/ml" 
	case FluidResistancePascalSecondPerCubicMeter:
		return "Pa·s/m³" 
	case FluidResistancePascalMinutePerCubicMeter:
		return "Pa·min/m³" 
	case FluidResistancePascalSecondPerCubicCentimeter:
		return "Pa·s/cm³" 
	case FluidResistancePascalMinutePerCubicCentimeter:
		return "Pa·min/cm³" 
	case FluidResistanceDyneSecondPerCentimeterToTheFifth:
		return "dyn·s/cm⁵" 
	case FluidResistanceMillimeterMercurySecondPerLiter:
		return "mmHg·s/l" 
	case FluidResistanceMillimeterMercuryMinutePerLiter:
		return "mmHg·min/l" 
	case FluidResistanceMillimeterMercurySecondPerMilliliter:
		return "mmHg·s/ml" 
	case FluidResistanceMillimeterMercuryMinutePerMilliliter:
		return "mmHg·min/ml" 
	case FluidResistanceMillimeterMercurySecondPerCubicCentimeter:
		return "mmHg·s/cm³" 
	case FluidResistanceMillimeterMercuryMinutePerCubicCentimeter:
		return "mmHg·min/cm³" 
	case FluidResistanceMillimeterMercurySecondPerCubicMeter:
		return "mmHg·s/m³" 
	case FluidResistanceMillimeterMercuryMinutePerCubicMeter:
		return "mmHg·min/m³" 
	case FluidResistanceWoodUnit:
		return "WU" 
	case FluidResistanceMegapascalSecondPerCubicMeter:
		return "MPa·s/m³" 
	default:
		return ""
	}
}