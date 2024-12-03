// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// IrradiationUnits enumeration
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

// IrradiationDto represents an Irradiation
type IrradiationDto struct {
	Value float64
	Unit  IrradiationUnits
}

// IrradiationDtoFactory struct to group related functions
type IrradiationDtoFactory struct{}

func (udf IrradiationDtoFactory) FromJSON(data []byte) (*IrradiationDto, error) {
	a := IrradiationDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a IrradiationDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// Irradiation struct
type Irradiation struct {
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

// IrradiationFactory struct to group related functions
type IrradiationFactory struct{}

func (uf IrradiationFactory) CreateIrradiation(value float64, unit IrradiationUnits) (*Irradiation, error) {
	return newIrradiation(value, unit)
}

func (uf IrradiationFactory) FromDto(dto IrradiationDto) (*Irradiation, error) {
	return newIrradiation(dto.Value, dto.Unit)
}

func (uf IrradiationFactory) FromDtoJSON(data []byte) (*Irradiation, error) {
	unitDto, err := IrradiationDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return IrradiationFactory{}.FromDto(*unitDto)
}


// FromJoulePerSquareMeter creates a new Irradiation instance from JoulePerSquareMeter.
func (uf IrradiationFactory) FromJoulesPerSquareMeter(value float64) (*Irradiation, error) {
	return newIrradiation(value, IrradiationJoulePerSquareMeter)
}

// FromJoulePerSquareCentimeter creates a new Irradiation instance from JoulePerSquareCentimeter.
func (uf IrradiationFactory) FromJoulesPerSquareCentimeter(value float64) (*Irradiation, error) {
	return newIrradiation(value, IrradiationJoulePerSquareCentimeter)
}

// FromJoulePerSquareMillimeter creates a new Irradiation instance from JoulePerSquareMillimeter.
func (uf IrradiationFactory) FromJoulesPerSquareMillimeter(value float64) (*Irradiation, error) {
	return newIrradiation(value, IrradiationJoulePerSquareMillimeter)
}

// FromWattHourPerSquareMeter creates a new Irradiation instance from WattHourPerSquareMeter.
func (uf IrradiationFactory) FromWattHoursPerSquareMeter(value float64) (*Irradiation, error) {
	return newIrradiation(value, IrradiationWattHourPerSquareMeter)
}

// FromBtuPerSquareFoot creates a new Irradiation instance from BtuPerSquareFoot.
func (uf IrradiationFactory) FromBtusPerSquareFoot(value float64) (*Irradiation, error) {
	return newIrradiation(value, IrradiationBtuPerSquareFoot)
}

// FromKilojoulePerSquareMeter creates a new Irradiation instance from KilojoulePerSquareMeter.
func (uf IrradiationFactory) FromKilojoulesPerSquareMeter(value float64) (*Irradiation, error) {
	return newIrradiation(value, IrradiationKilojoulePerSquareMeter)
}

// FromMillijoulePerSquareCentimeter creates a new Irradiation instance from MillijoulePerSquareCentimeter.
func (uf IrradiationFactory) FromMillijoulesPerSquareCentimeter(value float64) (*Irradiation, error) {
	return newIrradiation(value, IrradiationMillijoulePerSquareCentimeter)
}

// FromKilowattHourPerSquareMeter creates a new Irradiation instance from KilowattHourPerSquareMeter.
func (uf IrradiationFactory) FromKilowattHoursPerSquareMeter(value float64) (*Irradiation, error) {
	return newIrradiation(value, IrradiationKilowattHourPerSquareMeter)
}

// FromKilobtuPerSquareFoot creates a new Irradiation instance from KilobtuPerSquareFoot.
func (uf IrradiationFactory) FromKilobtusPerSquareFoot(value float64) (*Irradiation, error) {
	return newIrradiation(value, IrradiationKilobtuPerSquareFoot)
}




// newIrradiation creates a new Irradiation.
func newIrradiation(value float64, fromUnit IrradiationUnits) (*Irradiation, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Irradiation{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Irradiation in JoulePerSquareMeter.
func (a *Irradiation) BaseValue() float64 {
	return a.value
}


// JoulePerSquareMeter returns the value in JoulePerSquareMeter.
func (a *Irradiation) JoulesPerSquareMeter() float64 {
	if a.joules_per_square_meterLazy != nil {
		return *a.joules_per_square_meterLazy
	}
	joules_per_square_meter := a.convertFromBase(IrradiationJoulePerSquareMeter)
	a.joules_per_square_meterLazy = &joules_per_square_meter
	return joules_per_square_meter
}

// JoulePerSquareCentimeter returns the value in JoulePerSquareCentimeter.
func (a *Irradiation) JoulesPerSquareCentimeter() float64 {
	if a.joules_per_square_centimeterLazy != nil {
		return *a.joules_per_square_centimeterLazy
	}
	joules_per_square_centimeter := a.convertFromBase(IrradiationJoulePerSquareCentimeter)
	a.joules_per_square_centimeterLazy = &joules_per_square_centimeter
	return joules_per_square_centimeter
}

// JoulePerSquareMillimeter returns the value in JoulePerSquareMillimeter.
func (a *Irradiation) JoulesPerSquareMillimeter() float64 {
	if a.joules_per_square_millimeterLazy != nil {
		return *a.joules_per_square_millimeterLazy
	}
	joules_per_square_millimeter := a.convertFromBase(IrradiationJoulePerSquareMillimeter)
	a.joules_per_square_millimeterLazy = &joules_per_square_millimeter
	return joules_per_square_millimeter
}

// WattHourPerSquareMeter returns the value in WattHourPerSquareMeter.
func (a *Irradiation) WattHoursPerSquareMeter() float64 {
	if a.watt_hours_per_square_meterLazy != nil {
		return *a.watt_hours_per_square_meterLazy
	}
	watt_hours_per_square_meter := a.convertFromBase(IrradiationWattHourPerSquareMeter)
	a.watt_hours_per_square_meterLazy = &watt_hours_per_square_meter
	return watt_hours_per_square_meter
}

// BtuPerSquareFoot returns the value in BtuPerSquareFoot.
func (a *Irradiation) BtusPerSquareFoot() float64 {
	if a.btus_per_square_footLazy != nil {
		return *a.btus_per_square_footLazy
	}
	btus_per_square_foot := a.convertFromBase(IrradiationBtuPerSquareFoot)
	a.btus_per_square_footLazy = &btus_per_square_foot
	return btus_per_square_foot
}

// KilojoulePerSquareMeter returns the value in KilojoulePerSquareMeter.
func (a *Irradiation) KilojoulesPerSquareMeter() float64 {
	if a.kilojoules_per_square_meterLazy != nil {
		return *a.kilojoules_per_square_meterLazy
	}
	kilojoules_per_square_meter := a.convertFromBase(IrradiationKilojoulePerSquareMeter)
	a.kilojoules_per_square_meterLazy = &kilojoules_per_square_meter
	return kilojoules_per_square_meter
}

// MillijoulePerSquareCentimeter returns the value in MillijoulePerSquareCentimeter.
func (a *Irradiation) MillijoulesPerSquareCentimeter() float64 {
	if a.millijoules_per_square_centimeterLazy != nil {
		return *a.millijoules_per_square_centimeterLazy
	}
	millijoules_per_square_centimeter := a.convertFromBase(IrradiationMillijoulePerSquareCentimeter)
	a.millijoules_per_square_centimeterLazy = &millijoules_per_square_centimeter
	return millijoules_per_square_centimeter
}

// KilowattHourPerSquareMeter returns the value in KilowattHourPerSquareMeter.
func (a *Irradiation) KilowattHoursPerSquareMeter() float64 {
	if a.kilowatt_hours_per_square_meterLazy != nil {
		return *a.kilowatt_hours_per_square_meterLazy
	}
	kilowatt_hours_per_square_meter := a.convertFromBase(IrradiationKilowattHourPerSquareMeter)
	a.kilowatt_hours_per_square_meterLazy = &kilowatt_hours_per_square_meter
	return kilowatt_hours_per_square_meter
}

// KilobtuPerSquareFoot returns the value in KilobtuPerSquareFoot.
func (a *Irradiation) KilobtusPerSquareFoot() float64 {
	if a.kilobtus_per_square_footLazy != nil {
		return *a.kilobtus_per_square_footLazy
	}
	kilobtus_per_square_foot := a.convertFromBase(IrradiationKilobtuPerSquareFoot)
	a.kilobtus_per_square_footLazy = &kilobtus_per_square_foot
	return kilobtus_per_square_foot
}


// ToDto creates an IrradiationDto representation.
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

// ToDtoJSON creates an IrradiationDto representation.
func (a *Irradiation) ToDtoJSON(holdInUnit *IrradiationUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts Irradiation to a specific unit value.
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
		return 0
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
		return (value / (52752792631 / 4645152)) 
	case IrradiationKilojoulePerSquareMeter:
		return ((value) / 1000.0) 
	case IrradiationMillijoulePerSquareCentimeter:
		return ((value / 1e4) / 0.001) 
	case IrradiationKilowattHourPerSquareMeter:
		return ((value / 3600) / 1000.0) 
	case IrradiationKilobtuPerSquareFoot:
		return ((value / (52752792631 / 4645152)) / 1000.0) 
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
		return (value * (52752792631 / 4645152)) 
	case IrradiationKilojoulePerSquareMeter:
		return ((value) * 1000.0) 
	case IrradiationMillijoulePerSquareCentimeter:
		return ((value * 1e4) * 0.001) 
	case IrradiationKilowattHourPerSquareMeter:
		return ((value * 3600) * 1000.0) 
	case IrradiationKilobtuPerSquareFoot:
		return ((value * (52752792631 / 4645152)) * 1000.0) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a Irradiation) String() string {
	return a.ToString(IrradiationJoulePerSquareMeter, 2)
}

// ToString formats the Irradiation to string.
// fractionalDigits -1 for not mention
func (a *Irradiation) ToString(unit IrradiationUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Irradiation) getUnitAbbreviation(unit IrradiationUnits) string {
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

// Check if the given Irradiation are equals to the current Irradiation
func (a *Irradiation) Equals(other *Irradiation) bool {
	return a.value == other.BaseValue()
}

// Check if the given Irradiation are equals to the current Irradiation
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

// Add the given Irradiation to the current Irradiation.
func (a *Irradiation) Add(other *Irradiation) *Irradiation {
	return &Irradiation{value: a.value + other.BaseValue()}
}

// Subtract the given Irradiation to the current Irradiation.
func (a *Irradiation) Subtract(other *Irradiation) *Irradiation {
	return &Irradiation{value: a.value - other.BaseValue()}
}

// Multiply the given Irradiation to the current Irradiation.
func (a *Irradiation) Multiply(other *Irradiation) *Irradiation {
	return &Irradiation{value: a.value * other.BaseValue()}
}

// Divide the given Irradiation to the current Irradiation.
func (a *Irradiation) Divide(other *Irradiation) *Irradiation {
	return &Irradiation{value: a.value / other.BaseValue()}
}