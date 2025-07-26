// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// IrradiationUnits defines various units of Irradiation.
type IrradiationUnits string

const (
    
        // 
        IrradiationJoulePerSquareMeter IrradiationUnits = "JoulePerSquareMeter"
        // 
        IrradiationJoulePerSquareCentimeter IrradiationUnits = "JoulePerSquareCentimeter"
        // 
        IrradiationJoulePerSquareMillimeter IrradiationUnits = "JoulePerSquareMillimeter"
        // 
        IrradiationWattHourPerSquareMeter IrradiationUnits = "WattHourPerSquareMeter"
        // 
        IrradiationBtuPerSquareFoot IrradiationUnits = "BtuPerSquareFoot"
        // 
        IrradiationKilojoulePerSquareMeter IrradiationUnits = "KilojoulePerSquareMeter"
        // 
        IrradiationMillijoulePerSquareCentimeter IrradiationUnits = "MillijoulePerSquareCentimeter"
        // 
        IrradiationKilowattHourPerSquareMeter IrradiationUnits = "KilowattHourPerSquareMeter"
        // 
        IrradiationKilobtuPerSquareFoot IrradiationUnits = "KilobtuPerSquareFoot"
)

var internalIrradiationUnitsMap = map[IrradiationUnits]bool{
	
	IrradiationJoulePerSquareMeter: true,
	IrradiationJoulePerSquareCentimeter: true,
	IrradiationJoulePerSquareMillimeter: true,
	IrradiationWattHourPerSquareMeter: true,
	IrradiationBtuPerSquareFoot: true,
	IrradiationKilojoulePerSquareMeter: true,
	IrradiationMillijoulePerSquareCentimeter: true,
	IrradiationKilowattHourPerSquareMeter: true,
	IrradiationKilobtuPerSquareFoot: true,
}

// IrradiationDto represents a Irradiation measurement with a numerical value and its corresponding unit.
type IrradiationDto struct {
    // Value is the numerical representation of the Irradiation.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the Irradiation, as defined in the IrradiationUnits enumeration.
	Unit  IrradiationUnits `json:"unit" validate:"required,oneof=JoulePerSquareMeter JoulePerSquareCentimeter JoulePerSquareMillimeter WattHourPerSquareMeter BtuPerSquareFoot KilojoulePerSquareMeter MillijoulePerSquareCentimeter KilowattHourPerSquareMeter KilobtuPerSquareFoot"`
}

// IrradiationDtoFactory groups methods for creating and serializing IrradiationDto objects.
type IrradiationDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a IrradiationDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf IrradiationDtoFactory) FromJSON(data []byte) (*IrradiationDto, error) {
	a := IrradiationDto{}

    // Parse JSON into IrradiationDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a IrradiationDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a IrradiationDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// Irradiation represents a measurement in a of Irradiation.
//
// Irradiation is the process by which an object is exposed to radiation. The exposure can originate from various sources, including natural sources.
type Irradiation struct {
	// value is the base measurement stored internally.
	value       float64
    
    joules_per_square_meterLazy *float64 
    joules_per_square_centimeterLazy *float64 
    joules_per_square_millimeterLazy *float64 
    watt_hours_per_square_meterLazy *float64 
    btus_per_square_footLazy *float64 
    kilojoules_per_square_meterLazy *float64 
    millijoules_per_square_centimeterLazy *float64 
    kilowatt_hours_per_square_meterLazy *float64 
    kilobtus_per_square_footLazy *float64 
}

// IrradiationFactory groups methods for creating Irradiation instances.
type IrradiationFactory struct{}

// CreateIrradiation creates a new Irradiation instance from the given value and unit.
func (uf IrradiationFactory) CreateIrradiation(value float64, unit IrradiationUnits) (*Irradiation, error) {
	return newIrradiation(value, unit)
}

// FromDto converts a IrradiationDto to a Irradiation instance.
func (uf IrradiationFactory) FromDto(dto IrradiationDto) (*Irradiation, error) {
	return newIrradiation(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a Irradiation instance.
func (uf IrradiationFactory) FromDtoJSON(data []byte) (*Irradiation, error) {
	unitDto, err := IrradiationDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse IrradiationDto from JSON: %w", err)
	}
	return IrradiationFactory{}.FromDto(*unitDto)
}


// FromJoulesPerSquareMeter creates a new Irradiation instance from a value in JoulesPerSquareMeter.
func (uf IrradiationFactory) FromJoulesPerSquareMeter(value float64) (*Irradiation, error) {
	return newIrradiation(value, IrradiationJoulePerSquareMeter)
}

// FromJoulesPerSquareCentimeter creates a new Irradiation instance from a value in JoulesPerSquareCentimeter.
func (uf IrradiationFactory) FromJoulesPerSquareCentimeter(value float64) (*Irradiation, error) {
	return newIrradiation(value, IrradiationJoulePerSquareCentimeter)
}

// FromJoulesPerSquareMillimeter creates a new Irradiation instance from a value in JoulesPerSquareMillimeter.
func (uf IrradiationFactory) FromJoulesPerSquareMillimeter(value float64) (*Irradiation, error) {
	return newIrradiation(value, IrradiationJoulePerSquareMillimeter)
}

// FromWattHoursPerSquareMeter creates a new Irradiation instance from a value in WattHoursPerSquareMeter.
func (uf IrradiationFactory) FromWattHoursPerSquareMeter(value float64) (*Irradiation, error) {
	return newIrradiation(value, IrradiationWattHourPerSquareMeter)
}

// FromBtusPerSquareFoot creates a new Irradiation instance from a value in BtusPerSquareFoot.
func (uf IrradiationFactory) FromBtusPerSquareFoot(value float64) (*Irradiation, error) {
	return newIrradiation(value, IrradiationBtuPerSquareFoot)
}

// FromKilojoulesPerSquareMeter creates a new Irradiation instance from a value in KilojoulesPerSquareMeter.
func (uf IrradiationFactory) FromKilojoulesPerSquareMeter(value float64) (*Irradiation, error) {
	return newIrradiation(value, IrradiationKilojoulePerSquareMeter)
}

// FromMillijoulesPerSquareCentimeter creates a new Irradiation instance from a value in MillijoulesPerSquareCentimeter.
func (uf IrradiationFactory) FromMillijoulesPerSquareCentimeter(value float64) (*Irradiation, error) {
	return newIrradiation(value, IrradiationMillijoulePerSquareCentimeter)
}

// FromKilowattHoursPerSquareMeter creates a new Irradiation instance from a value in KilowattHoursPerSquareMeter.
func (uf IrradiationFactory) FromKilowattHoursPerSquareMeter(value float64) (*Irradiation, error) {
	return newIrradiation(value, IrradiationKilowattHourPerSquareMeter)
}

// FromKilobtusPerSquareFoot creates a new Irradiation instance from a value in KilobtusPerSquareFoot.
func (uf IrradiationFactory) FromKilobtusPerSquareFoot(value float64) (*Irradiation, error) {
	return newIrradiation(value, IrradiationKilobtuPerSquareFoot)
}


// newIrradiation creates a new Irradiation.
func newIrradiation(value float64, fromUnit IrradiationUnits) (*Irradiation, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalIrradiationUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in IrradiationUnits", fromUnit)
	}
	a := &Irradiation{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Irradiation in JoulePerSquareMeter unit (the base/default quantity).
func (a *Irradiation) BaseValue() float64 {
	return a.value
}


// JoulesPerSquareMeter returns the Irradiation value in JoulesPerSquareMeter.
//
// 
func (a *Irradiation) JoulesPerSquareMeter() float64 {
	if a.joules_per_square_meterLazy != nil {
		return *a.joules_per_square_meterLazy
	}
	joules_per_square_meter := a.convertFromBase(IrradiationJoulePerSquareMeter)
	a.joules_per_square_meterLazy = &joules_per_square_meter
	return joules_per_square_meter
}

// JoulesPerSquareCentimeter returns the Irradiation value in JoulesPerSquareCentimeter.
//
// 
func (a *Irradiation) JoulesPerSquareCentimeter() float64 {
	if a.joules_per_square_centimeterLazy != nil {
		return *a.joules_per_square_centimeterLazy
	}
	joules_per_square_centimeter := a.convertFromBase(IrradiationJoulePerSquareCentimeter)
	a.joules_per_square_centimeterLazy = &joules_per_square_centimeter
	return joules_per_square_centimeter
}

// JoulesPerSquareMillimeter returns the Irradiation value in JoulesPerSquareMillimeter.
//
// 
func (a *Irradiation) JoulesPerSquareMillimeter() float64 {
	if a.joules_per_square_millimeterLazy != nil {
		return *a.joules_per_square_millimeterLazy
	}
	joules_per_square_millimeter := a.convertFromBase(IrradiationJoulePerSquareMillimeter)
	a.joules_per_square_millimeterLazy = &joules_per_square_millimeter
	return joules_per_square_millimeter
}

// WattHoursPerSquareMeter returns the Irradiation value in WattHoursPerSquareMeter.
//
// 
func (a *Irradiation) WattHoursPerSquareMeter() float64 {
	if a.watt_hours_per_square_meterLazy != nil {
		return *a.watt_hours_per_square_meterLazy
	}
	watt_hours_per_square_meter := a.convertFromBase(IrradiationWattHourPerSquareMeter)
	a.watt_hours_per_square_meterLazy = &watt_hours_per_square_meter
	return watt_hours_per_square_meter
}

// BtusPerSquareFoot returns the Irradiation value in BtusPerSquareFoot.
//
// 
func (a *Irradiation) BtusPerSquareFoot() float64 {
	if a.btus_per_square_footLazy != nil {
		return *a.btus_per_square_footLazy
	}
	btus_per_square_foot := a.convertFromBase(IrradiationBtuPerSquareFoot)
	a.btus_per_square_footLazy = &btus_per_square_foot
	return btus_per_square_foot
}

// KilojoulesPerSquareMeter returns the Irradiation value in KilojoulesPerSquareMeter.
//
// 
func (a *Irradiation) KilojoulesPerSquareMeter() float64 {
	if a.kilojoules_per_square_meterLazy != nil {
		return *a.kilojoules_per_square_meterLazy
	}
	kilojoules_per_square_meter := a.convertFromBase(IrradiationKilojoulePerSquareMeter)
	a.kilojoules_per_square_meterLazy = &kilojoules_per_square_meter
	return kilojoules_per_square_meter
}

// MillijoulesPerSquareCentimeter returns the Irradiation value in MillijoulesPerSquareCentimeter.
//
// 
func (a *Irradiation) MillijoulesPerSquareCentimeter() float64 {
	if a.millijoules_per_square_centimeterLazy != nil {
		return *a.millijoules_per_square_centimeterLazy
	}
	millijoules_per_square_centimeter := a.convertFromBase(IrradiationMillijoulePerSquareCentimeter)
	a.millijoules_per_square_centimeterLazy = &millijoules_per_square_centimeter
	return millijoules_per_square_centimeter
}

// KilowattHoursPerSquareMeter returns the Irradiation value in KilowattHoursPerSquareMeter.
//
// 
func (a *Irradiation) KilowattHoursPerSquareMeter() float64 {
	if a.kilowatt_hours_per_square_meterLazy != nil {
		return *a.kilowatt_hours_per_square_meterLazy
	}
	kilowatt_hours_per_square_meter := a.convertFromBase(IrradiationKilowattHourPerSquareMeter)
	a.kilowatt_hours_per_square_meterLazy = &kilowatt_hours_per_square_meter
	return kilowatt_hours_per_square_meter
}

// KilobtusPerSquareFoot returns the Irradiation value in KilobtusPerSquareFoot.
//
// 
func (a *Irradiation) KilobtusPerSquareFoot() float64 {
	if a.kilobtus_per_square_footLazy != nil {
		return *a.kilobtus_per_square_footLazy
	}
	kilobtus_per_square_foot := a.convertFromBase(IrradiationKilobtuPerSquareFoot)
	a.kilobtus_per_square_footLazy = &kilobtus_per_square_foot
	return kilobtus_per_square_foot
}


// ToDto creates a IrradiationDto representation from the Irradiation instance.
//
// If the provided holdInUnit is nil, the value will be repesented by JoulePerSquareMeter by default.
func (a *Irradiation) ToDto(holdInUnit *IrradiationUnits) IrradiationDto {
	if holdInUnit == nil {
		defaultUnit := IrradiationJoulePerSquareMeter // Default value
		holdInUnit = &defaultUnit
	}

	return IrradiationDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the Irradiation instance.
//
// If the provided holdInUnit is nil, the value will be repesented by JoulePerSquareMeter by default.
func (a *Irradiation) ToDtoJSON(holdInUnit *IrradiationUnits) ([]byte, error) {
	// Convert to IrradiationDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a Irradiation to a specific unit value.
// The function uses the provided unit type (IrradiationUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *Irradiation) Convert(toUnit IrradiationUnits) float64 {
	switch toUnit { 
    case IrradiationJoulePerSquareMeter:
		return a.JoulesPerSquareMeter()
    case IrradiationJoulePerSquareCentimeter:
		return a.JoulesPerSquareCentimeter()
    case IrradiationJoulePerSquareMillimeter:
		return a.JoulesPerSquareMillimeter()
    case IrradiationWattHourPerSquareMeter:
		return a.WattHoursPerSquareMeter()
    case IrradiationBtuPerSquareFoot:
		return a.BtusPerSquareFoot()
    case IrradiationKilojoulePerSquareMeter:
		return a.KilojoulesPerSquareMeter()
    case IrradiationMillijoulePerSquareCentimeter:
		return a.MillijoulesPerSquareCentimeter()
    case IrradiationKilowattHourPerSquareMeter:
		return a.KilowattHoursPerSquareMeter()
    case IrradiationKilobtuPerSquareFoot:
		return a.KilobtusPerSquareFoot()
	default:
		return math.NaN()
	}
}

func (a *Irradiation) convertFromBase(toUnit IrradiationUnits) float64 {
    value := a.value
	switch toUnit { 
	case IrradiationJoulePerSquareMeter:
		return (value) 
	case IrradiationJoulePerSquareCentimeter:
		return (value / 1e4) 
	case IrradiationJoulePerSquareMillimeter:
		return (value / 1e6) 
	case IrradiationWattHourPerSquareMeter:
		return (value / 3600) 
	case IrradiationBtuPerSquareFoot:
		return (value * 9.290304e-2 / 1055.05585262) 
	case IrradiationKilojoulePerSquareMeter:
		return ((value) / 1000.0) 
	case IrradiationMillijoulePerSquareCentimeter:
		return ((value / 1e4) / 0.001) 
	case IrradiationKilowattHourPerSquareMeter:
		return ((value / 3600) / 1000.0) 
	case IrradiationKilobtuPerSquareFoot:
		return ((value * 9.290304e-2 / 1055.05585262) / 1000.0) 
	default:
		return math.NaN()
	}
}

func (a *Irradiation) convertToBase(value float64, fromUnit IrradiationUnits) float64 {
	switch fromUnit { 
	case IrradiationJoulePerSquareMeter:
		return (value) 
	case IrradiationJoulePerSquareCentimeter:
		return (value * 1e4) 
	case IrradiationJoulePerSquareMillimeter:
		return (value * 1e6) 
	case IrradiationWattHourPerSquareMeter:
		return (value * 3600) 
	case IrradiationBtuPerSquareFoot:
		return (value * 1055.05585262 / 9.290304e-2) 
	case IrradiationKilojoulePerSquareMeter:
		return ((value) * 1000.0) 
	case IrradiationMillijoulePerSquareCentimeter:
		return ((value * 1e4) * 0.001) 
	case IrradiationKilowattHourPerSquareMeter:
		return ((value * 3600) * 1000.0) 
	case IrradiationKilobtuPerSquareFoot:
		return ((value * 1055.05585262 / 9.290304e-2) * 1000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the Irradiation in the default unit (JoulePerSquareMeter),
// formatted to two decimal places.
func (a Irradiation) String() string {
	return a.ToString(IrradiationJoulePerSquareMeter, 2)
}

// ToString formats the Irradiation value as a string with the specified unit and fractional digits.
// It converts the Irradiation to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the Irradiation value will be converted (e.g., JoulePerSquareMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the Irradiation with the unit abbreviation.
func (a *Irradiation) ToString(unit IrradiationUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetIrradiationAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetIrradiationAbbreviation(unit))
}

// Equals checks if the given Irradiation is equal to the current Irradiation.
//
// Parameters:
//    other: The Irradiation to compare against.
//
// Returns:
//    bool: Returns true if both Irradiation are equal, false otherwise.
func (a *Irradiation) Equals(other *Irradiation) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current Irradiation with another Irradiation.
// It returns -1 if the current Irradiation is less than the other Irradiation, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The Irradiation to compare against.
//
// Returns:
//    int: -1 if the current Irradiation is less, 1 if greater, and 0 if equal.
func (a *Irradiation) CompareTo(other *Irradiation) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given Irradiation to the current Irradiation and returns the result.
// The result is a new Irradiation instance with the sum of the values.
//
// Parameters:
//    other: The Irradiation to add to the current Irradiation.
//
// Returns:
//    *Irradiation: A new Irradiation instance representing the sum of both Irradiation.
func (a *Irradiation) Add(other *Irradiation) *Irradiation {
	return &Irradiation{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given Irradiation from the current Irradiation and returns the result.
// The result is a new Irradiation instance with the difference of the values.
//
// Parameters:
//    other: The Irradiation to subtract from the current Irradiation.
//
// Returns:
//    *Irradiation: A new Irradiation instance representing the difference of both Irradiation.
func (a *Irradiation) Subtract(other *Irradiation) *Irradiation {
	return &Irradiation{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current Irradiation by the given Irradiation and returns the result.
// The result is a new Irradiation instance with the product of the values.
//
// Parameters:
//    other: The Irradiation to multiply with the current Irradiation.
//
// Returns:
//    *Irradiation: A new Irradiation instance representing the product of both Irradiation.
func (a *Irradiation) Multiply(other *Irradiation) *Irradiation {
	return &Irradiation{value: a.value * other.BaseValue()}
}

// Divide divides the current Irradiation by the given Irradiation and returns the result.
// The result is a new Irradiation instance with the quotient of the values.
//
// Parameters:
//    other: The Irradiation to divide the current Irradiation by.
//
// Returns:
//    *Irradiation: A new Irradiation instance representing the quotient of both Irradiation.
func (a *Irradiation) Divide(other *Irradiation) *Irradiation {
	return &Irradiation{value: a.value / other.BaseValue()}
}

// GetIrradiationAbbreviation gets the unit abbreviation.
func GetIrradiationAbbreviation(unit IrradiationUnits) string {
	switch unit { 
	case IrradiationJoulePerSquareMeter:
		return "J/m²" 
	case IrradiationJoulePerSquareCentimeter:
		return "J/cm²" 
	case IrradiationJoulePerSquareMillimeter:
		return "J/mm²" 
	case IrradiationWattHourPerSquareMeter:
		return "Wh/m²" 
	case IrradiationBtuPerSquareFoot:
		return "Btu/ft²" 
	case IrradiationKilojoulePerSquareMeter:
		return "kJ/m²" 
	case IrradiationMillijoulePerSquareCentimeter:
		return "mJ/cm²" 
	case IrradiationKilowattHourPerSquareMeter:
		return "kWh/m²" 
	case IrradiationKilobtuPerSquareFoot:
		return "kBtu/ft²" 
	default:
		return ""
	}
}