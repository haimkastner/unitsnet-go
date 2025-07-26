// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// VolumePerLengthUnits defines various units of VolumePerLength.
type VolumePerLengthUnits string

const (
    
        // 
        VolumePerLengthCubicMeterPerMeter VolumePerLengthUnits = "CubicMeterPerMeter"
        // 
        VolumePerLengthLiterPerMeter VolumePerLengthUnits = "LiterPerMeter"
        // 
        VolumePerLengthLiterPerKilometer VolumePerLengthUnits = "LiterPerKilometer"
        // 
        VolumePerLengthLiterPerMillimeter VolumePerLengthUnits = "LiterPerMillimeter"
        // 
        VolumePerLengthOilBarrelPerFoot VolumePerLengthUnits = "OilBarrelPerFoot"
        // 
        VolumePerLengthCubicYardPerFoot VolumePerLengthUnits = "CubicYardPerFoot"
        // 
        VolumePerLengthCubicYardPerUsSurveyFoot VolumePerLengthUnits = "CubicYardPerUsSurveyFoot"
        // 
        VolumePerLengthUsGallonPerMile VolumePerLengthUnits = "UsGallonPerMile"
        // 
        VolumePerLengthImperialGallonPerMile VolumePerLengthUnits = "ImperialGallonPerMile"
)

var internalVolumePerLengthUnitsMap = map[VolumePerLengthUnits]bool{
	
	VolumePerLengthCubicMeterPerMeter: true,
	VolumePerLengthLiterPerMeter: true,
	VolumePerLengthLiterPerKilometer: true,
	VolumePerLengthLiterPerMillimeter: true,
	VolumePerLengthOilBarrelPerFoot: true,
	VolumePerLengthCubicYardPerFoot: true,
	VolumePerLengthCubicYardPerUsSurveyFoot: true,
	VolumePerLengthUsGallonPerMile: true,
	VolumePerLengthImperialGallonPerMile: true,
}

// VolumePerLengthDto represents a VolumePerLength measurement with a numerical value and its corresponding unit.
type VolumePerLengthDto struct {
    // Value is the numerical representation of the VolumePerLength.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the VolumePerLength, as defined in the VolumePerLengthUnits enumeration.
	Unit  VolumePerLengthUnits `json:"unit" validate:"required,oneof=CubicMeterPerMeter LiterPerMeter LiterPerKilometer LiterPerMillimeter OilBarrelPerFoot CubicYardPerFoot CubicYardPerUsSurveyFoot UsGallonPerMile ImperialGallonPerMile"`
}

// VolumePerLengthDtoFactory groups methods for creating and serializing VolumePerLengthDto objects.
type VolumePerLengthDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a VolumePerLengthDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf VolumePerLengthDtoFactory) FromJSON(data []byte) (*VolumePerLengthDto, error) {
	a := VolumePerLengthDto{}

    // Parse JSON into VolumePerLengthDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a VolumePerLengthDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a VolumePerLengthDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// VolumePerLength represents a measurement in a of VolumePerLength.
//
// Volume, typically of fluid, that a container can hold within a unit of length.
type VolumePerLength struct {
	// value is the base measurement stored internally.
	value       float64
    
    cubic_meters_per_meterLazy *float64 
    liters_per_meterLazy *float64 
    liters_per_kilometerLazy *float64 
    liters_per_millimeterLazy *float64 
    oil_barrels_per_footLazy *float64 
    cubic_yards_per_footLazy *float64 
    cubic_yards_per_us_survey_footLazy *float64 
    us_gallons_per_mileLazy *float64 
    imperial_gallons_per_mileLazy *float64 
}

// VolumePerLengthFactory groups methods for creating VolumePerLength instances.
type VolumePerLengthFactory struct{}

// CreateVolumePerLength creates a new VolumePerLength instance from the given value and unit.
func (uf VolumePerLengthFactory) CreateVolumePerLength(value float64, unit VolumePerLengthUnits) (*VolumePerLength, error) {
	return newVolumePerLength(value, unit)
}

// FromDto converts a VolumePerLengthDto to a VolumePerLength instance.
func (uf VolumePerLengthFactory) FromDto(dto VolumePerLengthDto) (*VolumePerLength, error) {
	return newVolumePerLength(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a VolumePerLength instance.
func (uf VolumePerLengthFactory) FromDtoJSON(data []byte) (*VolumePerLength, error) {
	unitDto, err := VolumePerLengthDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse VolumePerLengthDto from JSON: %w", err)
	}
	return VolumePerLengthFactory{}.FromDto(*unitDto)
}


// FromCubicMetersPerMeter creates a new VolumePerLength instance from a value in CubicMetersPerMeter.
func (uf VolumePerLengthFactory) FromCubicMetersPerMeter(value float64) (*VolumePerLength, error) {
	return newVolumePerLength(value, VolumePerLengthCubicMeterPerMeter)
}

// FromLitersPerMeter creates a new VolumePerLength instance from a value in LitersPerMeter.
func (uf VolumePerLengthFactory) FromLitersPerMeter(value float64) (*VolumePerLength, error) {
	return newVolumePerLength(value, VolumePerLengthLiterPerMeter)
}

// FromLitersPerKilometer creates a new VolumePerLength instance from a value in LitersPerKilometer.
func (uf VolumePerLengthFactory) FromLitersPerKilometer(value float64) (*VolumePerLength, error) {
	return newVolumePerLength(value, VolumePerLengthLiterPerKilometer)
}

// FromLitersPerMillimeter creates a new VolumePerLength instance from a value in LitersPerMillimeter.
func (uf VolumePerLengthFactory) FromLitersPerMillimeter(value float64) (*VolumePerLength, error) {
	return newVolumePerLength(value, VolumePerLengthLiterPerMillimeter)
}

// FromOilBarrelsPerFoot creates a new VolumePerLength instance from a value in OilBarrelsPerFoot.
func (uf VolumePerLengthFactory) FromOilBarrelsPerFoot(value float64) (*VolumePerLength, error) {
	return newVolumePerLength(value, VolumePerLengthOilBarrelPerFoot)
}

// FromCubicYardsPerFoot creates a new VolumePerLength instance from a value in CubicYardsPerFoot.
func (uf VolumePerLengthFactory) FromCubicYardsPerFoot(value float64) (*VolumePerLength, error) {
	return newVolumePerLength(value, VolumePerLengthCubicYardPerFoot)
}

// FromCubicYardsPerUsSurveyFoot creates a new VolumePerLength instance from a value in CubicYardsPerUsSurveyFoot.
func (uf VolumePerLengthFactory) FromCubicYardsPerUsSurveyFoot(value float64) (*VolumePerLength, error) {
	return newVolumePerLength(value, VolumePerLengthCubicYardPerUsSurveyFoot)
}

// FromUsGallonsPerMile creates a new VolumePerLength instance from a value in UsGallonsPerMile.
func (uf VolumePerLengthFactory) FromUsGallonsPerMile(value float64) (*VolumePerLength, error) {
	return newVolumePerLength(value, VolumePerLengthUsGallonPerMile)
}

// FromImperialGallonsPerMile creates a new VolumePerLength instance from a value in ImperialGallonsPerMile.
func (uf VolumePerLengthFactory) FromImperialGallonsPerMile(value float64) (*VolumePerLength, error) {
	return newVolumePerLength(value, VolumePerLengthImperialGallonPerMile)
}


// newVolumePerLength creates a new VolumePerLength.
func newVolumePerLength(value float64, fromUnit VolumePerLengthUnits) (*VolumePerLength, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalVolumePerLengthUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in VolumePerLengthUnits", fromUnit)
	}
	a := &VolumePerLength{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of VolumePerLength in CubicMeterPerMeter unit (the base/default quantity).
func (a *VolumePerLength) BaseValue() float64 {
	return a.value
}


// CubicMetersPerMeter returns the VolumePerLength value in CubicMetersPerMeter.
//
// 
func (a *VolumePerLength) CubicMetersPerMeter() float64 {
	if a.cubic_meters_per_meterLazy != nil {
		return *a.cubic_meters_per_meterLazy
	}
	cubic_meters_per_meter := a.convertFromBase(VolumePerLengthCubicMeterPerMeter)
	a.cubic_meters_per_meterLazy = &cubic_meters_per_meter
	return cubic_meters_per_meter
}

// LitersPerMeter returns the VolumePerLength value in LitersPerMeter.
//
// 
func (a *VolumePerLength) LitersPerMeter() float64 {
	if a.liters_per_meterLazy != nil {
		return *a.liters_per_meterLazy
	}
	liters_per_meter := a.convertFromBase(VolumePerLengthLiterPerMeter)
	a.liters_per_meterLazy = &liters_per_meter
	return liters_per_meter
}

// LitersPerKilometer returns the VolumePerLength value in LitersPerKilometer.
//
// 
func (a *VolumePerLength) LitersPerKilometer() float64 {
	if a.liters_per_kilometerLazy != nil {
		return *a.liters_per_kilometerLazy
	}
	liters_per_kilometer := a.convertFromBase(VolumePerLengthLiterPerKilometer)
	a.liters_per_kilometerLazy = &liters_per_kilometer
	return liters_per_kilometer
}

// LitersPerMillimeter returns the VolumePerLength value in LitersPerMillimeter.
//
// 
func (a *VolumePerLength) LitersPerMillimeter() float64 {
	if a.liters_per_millimeterLazy != nil {
		return *a.liters_per_millimeterLazy
	}
	liters_per_millimeter := a.convertFromBase(VolumePerLengthLiterPerMillimeter)
	a.liters_per_millimeterLazy = &liters_per_millimeter
	return liters_per_millimeter
}

// OilBarrelsPerFoot returns the VolumePerLength value in OilBarrelsPerFoot.
//
// 
func (a *VolumePerLength) OilBarrelsPerFoot() float64 {
	if a.oil_barrels_per_footLazy != nil {
		return *a.oil_barrels_per_footLazy
	}
	oil_barrels_per_foot := a.convertFromBase(VolumePerLengthOilBarrelPerFoot)
	a.oil_barrels_per_footLazy = &oil_barrels_per_foot
	return oil_barrels_per_foot
}

// CubicYardsPerFoot returns the VolumePerLength value in CubicYardsPerFoot.
//
// 
func (a *VolumePerLength) CubicYardsPerFoot() float64 {
	if a.cubic_yards_per_footLazy != nil {
		return *a.cubic_yards_per_footLazy
	}
	cubic_yards_per_foot := a.convertFromBase(VolumePerLengthCubicYardPerFoot)
	a.cubic_yards_per_footLazy = &cubic_yards_per_foot
	return cubic_yards_per_foot
}

// CubicYardsPerUsSurveyFoot returns the VolumePerLength value in CubicYardsPerUsSurveyFoot.
//
// 
func (a *VolumePerLength) CubicYardsPerUsSurveyFoot() float64 {
	if a.cubic_yards_per_us_survey_footLazy != nil {
		return *a.cubic_yards_per_us_survey_footLazy
	}
	cubic_yards_per_us_survey_foot := a.convertFromBase(VolumePerLengthCubicYardPerUsSurveyFoot)
	a.cubic_yards_per_us_survey_footLazy = &cubic_yards_per_us_survey_foot
	return cubic_yards_per_us_survey_foot
}

// UsGallonsPerMile returns the VolumePerLength value in UsGallonsPerMile.
//
// 
func (a *VolumePerLength) UsGallonsPerMile() float64 {
	if a.us_gallons_per_mileLazy != nil {
		return *a.us_gallons_per_mileLazy
	}
	us_gallons_per_mile := a.convertFromBase(VolumePerLengthUsGallonPerMile)
	a.us_gallons_per_mileLazy = &us_gallons_per_mile
	return us_gallons_per_mile
}

// ImperialGallonsPerMile returns the VolumePerLength value in ImperialGallonsPerMile.
//
// 
func (a *VolumePerLength) ImperialGallonsPerMile() float64 {
	if a.imperial_gallons_per_mileLazy != nil {
		return *a.imperial_gallons_per_mileLazy
	}
	imperial_gallons_per_mile := a.convertFromBase(VolumePerLengthImperialGallonPerMile)
	a.imperial_gallons_per_mileLazy = &imperial_gallons_per_mile
	return imperial_gallons_per_mile
}


// ToDto creates a VolumePerLengthDto representation from the VolumePerLength instance.
//
// If the provided holdInUnit is nil, the value will be repesented by CubicMeterPerMeter by default.
func (a *VolumePerLength) ToDto(holdInUnit *VolumePerLengthUnits) VolumePerLengthDto {
	if holdInUnit == nil {
		defaultUnit := VolumePerLengthCubicMeterPerMeter // Default value
		holdInUnit = &defaultUnit
	}

	return VolumePerLengthDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the VolumePerLength instance.
//
// If the provided holdInUnit is nil, the value will be repesented by CubicMeterPerMeter by default.
func (a *VolumePerLength) ToDtoJSON(holdInUnit *VolumePerLengthUnits) ([]byte, error) {
	// Convert to VolumePerLengthDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a VolumePerLength to a specific unit value.
// The function uses the provided unit type (VolumePerLengthUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *VolumePerLength) Convert(toUnit VolumePerLengthUnits) float64 {
	switch toUnit { 
    case VolumePerLengthCubicMeterPerMeter:
		return a.CubicMetersPerMeter()
    case VolumePerLengthLiterPerMeter:
		return a.LitersPerMeter()
    case VolumePerLengthLiterPerKilometer:
		return a.LitersPerKilometer()
    case VolumePerLengthLiterPerMillimeter:
		return a.LitersPerMillimeter()
    case VolumePerLengthOilBarrelPerFoot:
		return a.OilBarrelsPerFoot()
    case VolumePerLengthCubicYardPerFoot:
		return a.CubicYardsPerFoot()
    case VolumePerLengthCubicYardPerUsSurveyFoot:
		return a.CubicYardsPerUsSurveyFoot()
    case VolumePerLengthUsGallonPerMile:
		return a.UsGallonsPerMile()
    case VolumePerLengthImperialGallonPerMile:
		return a.ImperialGallonsPerMile()
	default:
		return math.NaN()
	}
}

func (a *VolumePerLength) convertFromBase(toUnit VolumePerLengthUnits) float64 {
    value := a.value
	switch toUnit { 
	case VolumePerLengthCubicMeterPerMeter:
		return (value) 
	case VolumePerLengthLiterPerMeter:
		return (value * 1000) 
	case VolumePerLengthLiterPerKilometer:
		return (value * 1e6) 
	case VolumePerLengthLiterPerMillimeter:
		return (value) 
	case VolumePerLengthOilBarrelPerFoot:
		return (value * 0.3048 / 0.158987294928) 
	case VolumePerLengthCubicYardPerFoot:
		return (value * 0.3048 / 0.764554857984) 
	case VolumePerLengthCubicYardPerUsSurveyFoot:
		return (value * 1200 / (0.764554857984 * 3937)) 
	case VolumePerLengthUsGallonPerMile:
		return (value * 1609.344 / 0.003785411784) 
	case VolumePerLengthImperialGallonPerMile:
		return (value * 1609.344 / 0.00454609) 
	default:
		return math.NaN()
	}
}

func (a *VolumePerLength) convertToBase(value float64, fromUnit VolumePerLengthUnits) float64 {
	switch fromUnit { 
	case VolumePerLengthCubicMeterPerMeter:
		return (value) 
	case VolumePerLengthLiterPerMeter:
		return (value / 1000) 
	case VolumePerLengthLiterPerKilometer:
		return (value / 1e6) 
	case VolumePerLengthLiterPerMillimeter:
		return (value) 
	case VolumePerLengthOilBarrelPerFoot:
		return (value * 0.158987294928 / 0.3048) 
	case VolumePerLengthCubicYardPerFoot:
		return (value * 0.764554857984 / 0.3048) 
	case VolumePerLengthCubicYardPerUsSurveyFoot:
		return (value * 0.764554857984 * 3937 / 1200) 
	case VolumePerLengthUsGallonPerMile:
		return (value * 0.003785411784 / 1609.344) 
	case VolumePerLengthImperialGallonPerMile:
		return (value * 0.00454609 / 1609.344) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the VolumePerLength in the default unit (CubicMeterPerMeter),
// formatted to two decimal places.
func (a VolumePerLength) String() string {
	return a.ToString(VolumePerLengthCubicMeterPerMeter, 2)
}

// ToString formats the VolumePerLength value as a string with the specified unit and fractional digits.
// It converts the VolumePerLength to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the VolumePerLength value will be converted (e.g., CubicMeterPerMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the VolumePerLength with the unit abbreviation.
func (a *VolumePerLength) ToString(unit VolumePerLengthUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetVolumePerLengthAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetVolumePerLengthAbbreviation(unit))
}

// Equals checks if the given VolumePerLength is equal to the current VolumePerLength.
//
// Parameters:
//    other: The VolumePerLength to compare against.
//
// Returns:
//    bool: Returns true if both VolumePerLength are equal, false otherwise.
func (a *VolumePerLength) Equals(other *VolumePerLength) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current VolumePerLength with another VolumePerLength.
// It returns -1 if the current VolumePerLength is less than the other VolumePerLength, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The VolumePerLength to compare against.
//
// Returns:
//    int: -1 if the current VolumePerLength is less, 1 if greater, and 0 if equal.
func (a *VolumePerLength) CompareTo(other *VolumePerLength) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given VolumePerLength to the current VolumePerLength and returns the result.
// The result is a new VolumePerLength instance with the sum of the values.
//
// Parameters:
//    other: The VolumePerLength to add to the current VolumePerLength.
//
// Returns:
//    *VolumePerLength: A new VolumePerLength instance representing the sum of both VolumePerLength.
func (a *VolumePerLength) Add(other *VolumePerLength) *VolumePerLength {
	return &VolumePerLength{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given VolumePerLength from the current VolumePerLength and returns the result.
// The result is a new VolumePerLength instance with the difference of the values.
//
// Parameters:
//    other: The VolumePerLength to subtract from the current VolumePerLength.
//
// Returns:
//    *VolumePerLength: A new VolumePerLength instance representing the difference of both VolumePerLength.
func (a *VolumePerLength) Subtract(other *VolumePerLength) *VolumePerLength {
	return &VolumePerLength{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current VolumePerLength by the given VolumePerLength and returns the result.
// The result is a new VolumePerLength instance with the product of the values.
//
// Parameters:
//    other: The VolumePerLength to multiply with the current VolumePerLength.
//
// Returns:
//    *VolumePerLength: A new VolumePerLength instance representing the product of both VolumePerLength.
func (a *VolumePerLength) Multiply(other *VolumePerLength) *VolumePerLength {
	return &VolumePerLength{value: a.value * other.BaseValue()}
}

// Divide divides the current VolumePerLength by the given VolumePerLength and returns the result.
// The result is a new VolumePerLength instance with the quotient of the values.
//
// Parameters:
//    other: The VolumePerLength to divide the current VolumePerLength by.
//
// Returns:
//    *VolumePerLength: A new VolumePerLength instance representing the quotient of both VolumePerLength.
func (a *VolumePerLength) Divide(other *VolumePerLength) *VolumePerLength {
	return &VolumePerLength{value: a.value / other.BaseValue()}
}

// GetVolumePerLengthAbbreviation gets the unit abbreviation.
func GetVolumePerLengthAbbreviation(unit VolumePerLengthUnits) string {
	switch unit { 
	case VolumePerLengthCubicMeterPerMeter:
		return "m³/m" 
	case VolumePerLengthLiterPerMeter:
		return "l/m" 
	case VolumePerLengthLiterPerKilometer:
		return "l/km" 
	case VolumePerLengthLiterPerMillimeter:
		return "l/mm" 
	case VolumePerLengthOilBarrelPerFoot:
		return "bbl/ft" 
	case VolumePerLengthCubicYardPerFoot:
		return "yd³/ft" 
	case VolumePerLengthCubicYardPerUsSurveyFoot:
		return "yd³/ftUS" 
	case VolumePerLengthUsGallonPerMile:
		return "gal (U.S.)/mi" 
	case VolumePerLengthImperialGallonPerMile:
		return "gal (imp.)/mi" 
	default:
		return ""
	}
}