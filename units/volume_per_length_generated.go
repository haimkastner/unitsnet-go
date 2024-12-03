// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// VolumePerLengthUnits enumeration
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

// VolumePerLengthDto represents an VolumePerLength
type VolumePerLengthDto struct {
	Value float64
	Unit  VolumePerLengthUnits
}

// VolumePerLengthDtoFactory struct to group related functions
type VolumePerLengthDtoFactory struct{}

func (udf VolumePerLengthDtoFactory) FromJSON(data []byte) (*VolumePerLengthDto, error) {
	a := VolumePerLengthDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a VolumePerLengthDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// VolumePerLength struct
type VolumePerLength struct {
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

// VolumePerLengthFactory struct to group related functions
type VolumePerLengthFactory struct{}

func (uf VolumePerLengthFactory) CreateVolumePerLength(value float64, unit VolumePerLengthUnits) (*VolumePerLength, error) {
	return newVolumePerLength(value, unit)
}

func (uf VolumePerLengthFactory) FromDto(dto VolumePerLengthDto) (*VolumePerLength, error) {
	return newVolumePerLength(dto.Value, dto.Unit)
}

func (uf VolumePerLengthFactory) FromDtoJSON(data []byte) (*VolumePerLength, error) {
	unitDto, err := VolumePerLengthDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return VolumePerLengthFactory{}.FromDto(*unitDto)
}


// FromCubicMeterPerMeter creates a new VolumePerLength instance from CubicMeterPerMeter.
func (uf VolumePerLengthFactory) FromCubicMetersPerMeter(value float64) (*VolumePerLength, error) {
	return newVolumePerLength(value, VolumePerLengthCubicMeterPerMeter)
}

// FromLiterPerMeter creates a new VolumePerLength instance from LiterPerMeter.
func (uf VolumePerLengthFactory) FromLitersPerMeter(value float64) (*VolumePerLength, error) {
	return newVolumePerLength(value, VolumePerLengthLiterPerMeter)
}

// FromLiterPerKilometer creates a new VolumePerLength instance from LiterPerKilometer.
func (uf VolumePerLengthFactory) FromLitersPerKilometer(value float64) (*VolumePerLength, error) {
	return newVolumePerLength(value, VolumePerLengthLiterPerKilometer)
}

// FromLiterPerMillimeter creates a new VolumePerLength instance from LiterPerMillimeter.
func (uf VolumePerLengthFactory) FromLitersPerMillimeter(value float64) (*VolumePerLength, error) {
	return newVolumePerLength(value, VolumePerLengthLiterPerMillimeter)
}

// FromOilBarrelPerFoot creates a new VolumePerLength instance from OilBarrelPerFoot.
func (uf VolumePerLengthFactory) FromOilBarrelsPerFoot(value float64) (*VolumePerLength, error) {
	return newVolumePerLength(value, VolumePerLengthOilBarrelPerFoot)
}

// FromCubicYardPerFoot creates a new VolumePerLength instance from CubicYardPerFoot.
func (uf VolumePerLengthFactory) FromCubicYardsPerFoot(value float64) (*VolumePerLength, error) {
	return newVolumePerLength(value, VolumePerLengthCubicYardPerFoot)
}

// FromCubicYardPerUsSurveyFoot creates a new VolumePerLength instance from CubicYardPerUsSurveyFoot.
func (uf VolumePerLengthFactory) FromCubicYardsPerUsSurveyFoot(value float64) (*VolumePerLength, error) {
	return newVolumePerLength(value, VolumePerLengthCubicYardPerUsSurveyFoot)
}

// FromUsGallonPerMile creates a new VolumePerLength instance from UsGallonPerMile.
func (uf VolumePerLengthFactory) FromUsGallonsPerMile(value float64) (*VolumePerLength, error) {
	return newVolumePerLength(value, VolumePerLengthUsGallonPerMile)
}

// FromImperialGallonPerMile creates a new VolumePerLength instance from ImperialGallonPerMile.
func (uf VolumePerLengthFactory) FromImperialGallonsPerMile(value float64) (*VolumePerLength, error) {
	return newVolumePerLength(value, VolumePerLengthImperialGallonPerMile)
}




// newVolumePerLength creates a new VolumePerLength.
func newVolumePerLength(value float64, fromUnit VolumePerLengthUnits) (*VolumePerLength, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &VolumePerLength{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of VolumePerLength in CubicMeterPerMeter.
func (a *VolumePerLength) BaseValue() float64 {
	return a.value
}


// CubicMeterPerMeter returns the value in CubicMeterPerMeter.
func (a *VolumePerLength) CubicMetersPerMeter() float64 {
	if a.cubic_meters_per_meterLazy != nil {
		return *a.cubic_meters_per_meterLazy
	}
	cubic_meters_per_meter := a.convertFromBase(VolumePerLengthCubicMeterPerMeter)
	a.cubic_meters_per_meterLazy = &cubic_meters_per_meter
	return cubic_meters_per_meter
}

// LiterPerMeter returns the value in LiterPerMeter.
func (a *VolumePerLength) LitersPerMeter() float64 {
	if a.liters_per_meterLazy != nil {
		return *a.liters_per_meterLazy
	}
	liters_per_meter := a.convertFromBase(VolumePerLengthLiterPerMeter)
	a.liters_per_meterLazy = &liters_per_meter
	return liters_per_meter
}

// LiterPerKilometer returns the value in LiterPerKilometer.
func (a *VolumePerLength) LitersPerKilometer() float64 {
	if a.liters_per_kilometerLazy != nil {
		return *a.liters_per_kilometerLazy
	}
	liters_per_kilometer := a.convertFromBase(VolumePerLengthLiterPerKilometer)
	a.liters_per_kilometerLazy = &liters_per_kilometer
	return liters_per_kilometer
}

// LiterPerMillimeter returns the value in LiterPerMillimeter.
func (a *VolumePerLength) LitersPerMillimeter() float64 {
	if a.liters_per_millimeterLazy != nil {
		return *a.liters_per_millimeterLazy
	}
	liters_per_millimeter := a.convertFromBase(VolumePerLengthLiterPerMillimeter)
	a.liters_per_millimeterLazy = &liters_per_millimeter
	return liters_per_millimeter
}

// OilBarrelPerFoot returns the value in OilBarrelPerFoot.
func (a *VolumePerLength) OilBarrelsPerFoot() float64 {
	if a.oil_barrels_per_footLazy != nil {
		return *a.oil_barrels_per_footLazy
	}
	oil_barrels_per_foot := a.convertFromBase(VolumePerLengthOilBarrelPerFoot)
	a.oil_barrels_per_footLazy = &oil_barrels_per_foot
	return oil_barrels_per_foot
}

// CubicYardPerFoot returns the value in CubicYardPerFoot.
func (a *VolumePerLength) CubicYardsPerFoot() float64 {
	if a.cubic_yards_per_footLazy != nil {
		return *a.cubic_yards_per_footLazy
	}
	cubic_yards_per_foot := a.convertFromBase(VolumePerLengthCubicYardPerFoot)
	a.cubic_yards_per_footLazy = &cubic_yards_per_foot
	return cubic_yards_per_foot
}

// CubicYardPerUsSurveyFoot returns the value in CubicYardPerUsSurveyFoot.
func (a *VolumePerLength) CubicYardsPerUsSurveyFoot() float64 {
	if a.cubic_yards_per_us_survey_footLazy != nil {
		return *a.cubic_yards_per_us_survey_footLazy
	}
	cubic_yards_per_us_survey_foot := a.convertFromBase(VolumePerLengthCubicYardPerUsSurveyFoot)
	a.cubic_yards_per_us_survey_footLazy = &cubic_yards_per_us_survey_foot
	return cubic_yards_per_us_survey_foot
}

// UsGallonPerMile returns the value in UsGallonPerMile.
func (a *VolumePerLength) UsGallonsPerMile() float64 {
	if a.us_gallons_per_mileLazy != nil {
		return *a.us_gallons_per_mileLazy
	}
	us_gallons_per_mile := a.convertFromBase(VolumePerLengthUsGallonPerMile)
	a.us_gallons_per_mileLazy = &us_gallons_per_mile
	return us_gallons_per_mile
}

// ImperialGallonPerMile returns the value in ImperialGallonPerMile.
func (a *VolumePerLength) ImperialGallonsPerMile() float64 {
	if a.imperial_gallons_per_mileLazy != nil {
		return *a.imperial_gallons_per_mileLazy
	}
	imperial_gallons_per_mile := a.convertFromBase(VolumePerLengthImperialGallonPerMile)
	a.imperial_gallons_per_mileLazy = &imperial_gallons_per_mile
	return imperial_gallons_per_mile
}


// ToDto creates an VolumePerLengthDto representation.
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

// ToDtoJSON creates an VolumePerLengthDto representation.
func (a *VolumePerLength) ToDtoJSON(holdInUnit *VolumePerLengthUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts VolumePerLength to a specific unit value.
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
		return 0
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
		return (value * 1.91713408) 
	case VolumePerLengthCubicYardPerFoot:
		return (value / 2.50838208) 
	case VolumePerLengthCubicYardPerUsSurveyFoot:
		return (value / 2.50837706323584) 
	case VolumePerLengthUsGallonPerMile:
		return (value * (1000 * 1609.344 / 3.785411784)) 
	case VolumePerLengthImperialGallonPerMile:
		return (value * (1000 * 1609.344 / 4.54609)) 
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
		return (value / 1.91713408) 
	case VolumePerLengthCubicYardPerFoot:
		return (value * 2.50838208) 
	case VolumePerLengthCubicYardPerUsSurveyFoot:
		return (value * 2.50837706323584) 
	case VolumePerLengthUsGallonPerMile:
		return (value / (1000 * 1609.344 / 3.785411784)) 
	case VolumePerLengthImperialGallonPerMile:
		return (value / (1000 * 1609.344 / 4.54609)) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a VolumePerLength) String() string {
	return a.ToString(VolumePerLengthCubicMeterPerMeter, 2)
}

// ToString formats the VolumePerLength to string.
// fractionalDigits -1 for not mention
func (a *VolumePerLength) ToString(unit VolumePerLengthUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *VolumePerLength) getUnitAbbreviation(unit VolumePerLengthUnits) string {
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

// Check if the given VolumePerLength are equals to the current VolumePerLength
func (a *VolumePerLength) Equals(other *VolumePerLength) bool {
	return a.value == other.BaseValue()
}

// Check if the given VolumePerLength are equals to the current VolumePerLength
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

// Add the given VolumePerLength to the current VolumePerLength.
func (a *VolumePerLength) Add(other *VolumePerLength) *VolumePerLength {
	return &VolumePerLength{value: a.value + other.BaseValue()}
}

// Subtract the given VolumePerLength to the current VolumePerLength.
func (a *VolumePerLength) Subtract(other *VolumePerLength) *VolumePerLength {
	return &VolumePerLength{value: a.value - other.BaseValue()}
}

// Multiply the given VolumePerLength to the current VolumePerLength.
func (a *VolumePerLength) Multiply(other *VolumePerLength) *VolumePerLength {
	return &VolumePerLength{value: a.value * other.BaseValue()}
}

// Divide the given VolumePerLength to the current VolumePerLength.
func (a *VolumePerLength) Divide(other *VolumePerLength) *VolumePerLength {
	return &VolumePerLength{value: a.value / other.BaseValue()}
}