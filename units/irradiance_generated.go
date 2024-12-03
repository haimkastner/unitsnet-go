// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// IrradianceUnits enumeration
type IrradianceUnits string

const (
    
        // 
        IrradianceWattPerSquareMeter IrradianceUnits = "WattPerSquareMeter"
        // 
        IrradianceWattPerSquareCentimeter IrradianceUnits = "WattPerSquareCentimeter"
        // 
        IrradiancePicowattPerSquareMeter IrradianceUnits = "PicowattPerSquareMeter"
        // 
        IrradianceNanowattPerSquareMeter IrradianceUnits = "NanowattPerSquareMeter"
        // 
        IrradianceMicrowattPerSquareMeter IrradianceUnits = "MicrowattPerSquareMeter"
        // 
        IrradianceMilliwattPerSquareMeter IrradianceUnits = "MilliwattPerSquareMeter"
        // 
        IrradianceKilowattPerSquareMeter IrradianceUnits = "KilowattPerSquareMeter"
        // 
        IrradianceMegawattPerSquareMeter IrradianceUnits = "MegawattPerSquareMeter"
        // 
        IrradiancePicowattPerSquareCentimeter IrradianceUnits = "PicowattPerSquareCentimeter"
        // 
        IrradianceNanowattPerSquareCentimeter IrradianceUnits = "NanowattPerSquareCentimeter"
        // 
        IrradianceMicrowattPerSquareCentimeter IrradianceUnits = "MicrowattPerSquareCentimeter"
        // 
        IrradianceMilliwattPerSquareCentimeter IrradianceUnits = "MilliwattPerSquareCentimeter"
        // 
        IrradianceKilowattPerSquareCentimeter IrradianceUnits = "KilowattPerSquareCentimeter"
        // 
        IrradianceMegawattPerSquareCentimeter IrradianceUnits = "MegawattPerSquareCentimeter"
)

// IrradianceDto represents an Irradiance
type IrradianceDto struct {
	Value float64
	Unit  IrradianceUnits
}

// IrradianceDtoFactory struct to group related functions
type IrradianceDtoFactory struct{}

func (udf IrradianceDtoFactory) FromJSON(data []byte) (*IrradianceDto, error) {
	a := IrradianceDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a IrradianceDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// Irradiance struct
type Irradiance struct {
	value       float64
    
    watts_per_square_meterLazy *float64 
    watts_per_square_centimeterLazy *float64 
    picowatts_per_square_meterLazy *float64 
    nanowatts_per_square_meterLazy *float64 
    microwatts_per_square_meterLazy *float64 
    milliwatts_per_square_meterLazy *float64 
    kilowatts_per_square_meterLazy *float64 
    megawatts_per_square_meterLazy *float64 
    picowatts_per_square_centimeterLazy *float64 
    nanowatts_per_square_centimeterLazy *float64 
    microwatts_per_square_centimeterLazy *float64 
    milliwatts_per_square_centimeterLazy *float64 
    kilowatts_per_square_centimeterLazy *float64 
    megawatts_per_square_centimeterLazy *float64 
}

// IrradianceFactory struct to group related functions
type IrradianceFactory struct{}

func (uf IrradianceFactory) CreateIrradiance(value float64, unit IrradianceUnits) (*Irradiance, error) {
	return newIrradiance(value, unit)
}

func (uf IrradianceFactory) FromDto(dto IrradianceDto) (*Irradiance, error) {
	return newIrradiance(dto.Value, dto.Unit)
}

func (uf IrradianceFactory) FromDtoJSON(data []byte) (*Irradiance, error) {
	unitDto, err := IrradianceDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return IrradianceFactory{}.FromDto(*unitDto)
}


// FromWattPerSquareMeter creates a new Irradiance instance from WattPerSquareMeter.
func (uf IrradianceFactory) FromWattsPerSquareMeter(value float64) (*Irradiance, error) {
	return newIrradiance(value, IrradianceWattPerSquareMeter)
}

// FromWattPerSquareCentimeter creates a new Irradiance instance from WattPerSquareCentimeter.
func (uf IrradianceFactory) FromWattsPerSquareCentimeter(value float64) (*Irradiance, error) {
	return newIrradiance(value, IrradianceWattPerSquareCentimeter)
}

// FromPicowattPerSquareMeter creates a new Irradiance instance from PicowattPerSquareMeter.
func (uf IrradianceFactory) FromPicowattsPerSquareMeter(value float64) (*Irradiance, error) {
	return newIrradiance(value, IrradiancePicowattPerSquareMeter)
}

// FromNanowattPerSquareMeter creates a new Irradiance instance from NanowattPerSquareMeter.
func (uf IrradianceFactory) FromNanowattsPerSquareMeter(value float64) (*Irradiance, error) {
	return newIrradiance(value, IrradianceNanowattPerSquareMeter)
}

// FromMicrowattPerSquareMeter creates a new Irradiance instance from MicrowattPerSquareMeter.
func (uf IrradianceFactory) FromMicrowattsPerSquareMeter(value float64) (*Irradiance, error) {
	return newIrradiance(value, IrradianceMicrowattPerSquareMeter)
}

// FromMilliwattPerSquareMeter creates a new Irradiance instance from MilliwattPerSquareMeter.
func (uf IrradianceFactory) FromMilliwattsPerSquareMeter(value float64) (*Irradiance, error) {
	return newIrradiance(value, IrradianceMilliwattPerSquareMeter)
}

// FromKilowattPerSquareMeter creates a new Irradiance instance from KilowattPerSquareMeter.
func (uf IrradianceFactory) FromKilowattsPerSquareMeter(value float64) (*Irradiance, error) {
	return newIrradiance(value, IrradianceKilowattPerSquareMeter)
}

// FromMegawattPerSquareMeter creates a new Irradiance instance from MegawattPerSquareMeter.
func (uf IrradianceFactory) FromMegawattsPerSquareMeter(value float64) (*Irradiance, error) {
	return newIrradiance(value, IrradianceMegawattPerSquareMeter)
}

// FromPicowattPerSquareCentimeter creates a new Irradiance instance from PicowattPerSquareCentimeter.
func (uf IrradianceFactory) FromPicowattsPerSquareCentimeter(value float64) (*Irradiance, error) {
	return newIrradiance(value, IrradiancePicowattPerSquareCentimeter)
}

// FromNanowattPerSquareCentimeter creates a new Irradiance instance from NanowattPerSquareCentimeter.
func (uf IrradianceFactory) FromNanowattsPerSquareCentimeter(value float64) (*Irradiance, error) {
	return newIrradiance(value, IrradianceNanowattPerSquareCentimeter)
}

// FromMicrowattPerSquareCentimeter creates a new Irradiance instance from MicrowattPerSquareCentimeter.
func (uf IrradianceFactory) FromMicrowattsPerSquareCentimeter(value float64) (*Irradiance, error) {
	return newIrradiance(value, IrradianceMicrowattPerSquareCentimeter)
}

// FromMilliwattPerSquareCentimeter creates a new Irradiance instance from MilliwattPerSquareCentimeter.
func (uf IrradianceFactory) FromMilliwattsPerSquareCentimeter(value float64) (*Irradiance, error) {
	return newIrradiance(value, IrradianceMilliwattPerSquareCentimeter)
}

// FromKilowattPerSquareCentimeter creates a new Irradiance instance from KilowattPerSquareCentimeter.
func (uf IrradianceFactory) FromKilowattsPerSquareCentimeter(value float64) (*Irradiance, error) {
	return newIrradiance(value, IrradianceKilowattPerSquareCentimeter)
}

// FromMegawattPerSquareCentimeter creates a new Irradiance instance from MegawattPerSquareCentimeter.
func (uf IrradianceFactory) FromMegawattsPerSquareCentimeter(value float64) (*Irradiance, error) {
	return newIrradiance(value, IrradianceMegawattPerSquareCentimeter)
}




// newIrradiance creates a new Irradiance.
func newIrradiance(value float64, fromUnit IrradianceUnits) (*Irradiance, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Irradiance{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Irradiance in WattPerSquareMeter.
func (a *Irradiance) BaseValue() float64 {
	return a.value
}


// WattPerSquareMeter returns the value in WattPerSquareMeter.
func (a *Irradiance) WattsPerSquareMeter() float64 {
	if a.watts_per_square_meterLazy != nil {
		return *a.watts_per_square_meterLazy
	}
	watts_per_square_meter := a.convertFromBase(IrradianceWattPerSquareMeter)
	a.watts_per_square_meterLazy = &watts_per_square_meter
	return watts_per_square_meter
}

// WattPerSquareCentimeter returns the value in WattPerSquareCentimeter.
func (a *Irradiance) WattsPerSquareCentimeter() float64 {
	if a.watts_per_square_centimeterLazy != nil {
		return *a.watts_per_square_centimeterLazy
	}
	watts_per_square_centimeter := a.convertFromBase(IrradianceWattPerSquareCentimeter)
	a.watts_per_square_centimeterLazy = &watts_per_square_centimeter
	return watts_per_square_centimeter
}

// PicowattPerSquareMeter returns the value in PicowattPerSquareMeter.
func (a *Irradiance) PicowattsPerSquareMeter() float64 {
	if a.picowatts_per_square_meterLazy != nil {
		return *a.picowatts_per_square_meterLazy
	}
	picowatts_per_square_meter := a.convertFromBase(IrradiancePicowattPerSquareMeter)
	a.picowatts_per_square_meterLazy = &picowatts_per_square_meter
	return picowatts_per_square_meter
}

// NanowattPerSquareMeter returns the value in NanowattPerSquareMeter.
func (a *Irradiance) NanowattsPerSquareMeter() float64 {
	if a.nanowatts_per_square_meterLazy != nil {
		return *a.nanowatts_per_square_meterLazy
	}
	nanowatts_per_square_meter := a.convertFromBase(IrradianceNanowattPerSquareMeter)
	a.nanowatts_per_square_meterLazy = &nanowatts_per_square_meter
	return nanowatts_per_square_meter
}

// MicrowattPerSquareMeter returns the value in MicrowattPerSquareMeter.
func (a *Irradiance) MicrowattsPerSquareMeter() float64 {
	if a.microwatts_per_square_meterLazy != nil {
		return *a.microwatts_per_square_meterLazy
	}
	microwatts_per_square_meter := a.convertFromBase(IrradianceMicrowattPerSquareMeter)
	a.microwatts_per_square_meterLazy = &microwatts_per_square_meter
	return microwatts_per_square_meter
}

// MilliwattPerSquareMeter returns the value in MilliwattPerSquareMeter.
func (a *Irradiance) MilliwattsPerSquareMeter() float64 {
	if a.milliwatts_per_square_meterLazy != nil {
		return *a.milliwatts_per_square_meterLazy
	}
	milliwatts_per_square_meter := a.convertFromBase(IrradianceMilliwattPerSquareMeter)
	a.milliwatts_per_square_meterLazy = &milliwatts_per_square_meter
	return milliwatts_per_square_meter
}

// KilowattPerSquareMeter returns the value in KilowattPerSquareMeter.
func (a *Irradiance) KilowattsPerSquareMeter() float64 {
	if a.kilowatts_per_square_meterLazy != nil {
		return *a.kilowatts_per_square_meterLazy
	}
	kilowatts_per_square_meter := a.convertFromBase(IrradianceKilowattPerSquareMeter)
	a.kilowatts_per_square_meterLazy = &kilowatts_per_square_meter
	return kilowatts_per_square_meter
}

// MegawattPerSquareMeter returns the value in MegawattPerSquareMeter.
func (a *Irradiance) MegawattsPerSquareMeter() float64 {
	if a.megawatts_per_square_meterLazy != nil {
		return *a.megawatts_per_square_meterLazy
	}
	megawatts_per_square_meter := a.convertFromBase(IrradianceMegawattPerSquareMeter)
	a.megawatts_per_square_meterLazy = &megawatts_per_square_meter
	return megawatts_per_square_meter
}

// PicowattPerSquareCentimeter returns the value in PicowattPerSquareCentimeter.
func (a *Irradiance) PicowattsPerSquareCentimeter() float64 {
	if a.picowatts_per_square_centimeterLazy != nil {
		return *a.picowatts_per_square_centimeterLazy
	}
	picowatts_per_square_centimeter := a.convertFromBase(IrradiancePicowattPerSquareCentimeter)
	a.picowatts_per_square_centimeterLazy = &picowatts_per_square_centimeter
	return picowatts_per_square_centimeter
}

// NanowattPerSquareCentimeter returns the value in NanowattPerSquareCentimeter.
func (a *Irradiance) NanowattsPerSquareCentimeter() float64 {
	if a.nanowatts_per_square_centimeterLazy != nil {
		return *a.nanowatts_per_square_centimeterLazy
	}
	nanowatts_per_square_centimeter := a.convertFromBase(IrradianceNanowattPerSquareCentimeter)
	a.nanowatts_per_square_centimeterLazy = &nanowatts_per_square_centimeter
	return nanowatts_per_square_centimeter
}

// MicrowattPerSquareCentimeter returns the value in MicrowattPerSquareCentimeter.
func (a *Irradiance) MicrowattsPerSquareCentimeter() float64 {
	if a.microwatts_per_square_centimeterLazy != nil {
		return *a.microwatts_per_square_centimeterLazy
	}
	microwatts_per_square_centimeter := a.convertFromBase(IrradianceMicrowattPerSquareCentimeter)
	a.microwatts_per_square_centimeterLazy = &microwatts_per_square_centimeter
	return microwatts_per_square_centimeter
}

// MilliwattPerSquareCentimeter returns the value in MilliwattPerSquareCentimeter.
func (a *Irradiance) MilliwattsPerSquareCentimeter() float64 {
	if a.milliwatts_per_square_centimeterLazy != nil {
		return *a.milliwatts_per_square_centimeterLazy
	}
	milliwatts_per_square_centimeter := a.convertFromBase(IrradianceMilliwattPerSquareCentimeter)
	a.milliwatts_per_square_centimeterLazy = &milliwatts_per_square_centimeter
	return milliwatts_per_square_centimeter
}

// KilowattPerSquareCentimeter returns the value in KilowattPerSquareCentimeter.
func (a *Irradiance) KilowattsPerSquareCentimeter() float64 {
	if a.kilowatts_per_square_centimeterLazy != nil {
		return *a.kilowatts_per_square_centimeterLazy
	}
	kilowatts_per_square_centimeter := a.convertFromBase(IrradianceKilowattPerSquareCentimeter)
	a.kilowatts_per_square_centimeterLazy = &kilowatts_per_square_centimeter
	return kilowatts_per_square_centimeter
}

// MegawattPerSquareCentimeter returns the value in MegawattPerSquareCentimeter.
func (a *Irradiance) MegawattsPerSquareCentimeter() float64 {
	if a.megawatts_per_square_centimeterLazy != nil {
		return *a.megawatts_per_square_centimeterLazy
	}
	megawatts_per_square_centimeter := a.convertFromBase(IrradianceMegawattPerSquareCentimeter)
	a.megawatts_per_square_centimeterLazy = &megawatts_per_square_centimeter
	return megawatts_per_square_centimeter
}


// ToDto creates an IrradianceDto representation.
func (a *Irradiance) ToDto(holdInUnit *IrradianceUnits) IrradianceDto {
	if holdInUnit == nil {
		defaultUnit := IrradianceWattPerSquareMeter // Default value
		holdInUnit = &defaultUnit
	}

	return IrradianceDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an IrradianceDto representation.
func (a *Irradiance) ToDtoJSON(holdInUnit *IrradianceUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts Irradiance to a specific unit value.
func (a *Irradiance) Convert(toUnit IrradianceUnits) float64 {
	switch toUnit { 
    case IrradianceWattPerSquareMeter:
		return a.WattsPerSquareMeter()
    case IrradianceWattPerSquareCentimeter:
		return a.WattsPerSquareCentimeter()
    case IrradiancePicowattPerSquareMeter:
		return a.PicowattsPerSquareMeter()
    case IrradianceNanowattPerSquareMeter:
		return a.NanowattsPerSquareMeter()
    case IrradianceMicrowattPerSquareMeter:
		return a.MicrowattsPerSquareMeter()
    case IrradianceMilliwattPerSquareMeter:
		return a.MilliwattsPerSquareMeter()
    case IrradianceKilowattPerSquareMeter:
		return a.KilowattsPerSquareMeter()
    case IrradianceMegawattPerSquareMeter:
		return a.MegawattsPerSquareMeter()
    case IrradiancePicowattPerSquareCentimeter:
		return a.PicowattsPerSquareCentimeter()
    case IrradianceNanowattPerSquareCentimeter:
		return a.NanowattsPerSquareCentimeter()
    case IrradianceMicrowattPerSquareCentimeter:
		return a.MicrowattsPerSquareCentimeter()
    case IrradianceMilliwattPerSquareCentimeter:
		return a.MilliwattsPerSquareCentimeter()
    case IrradianceKilowattPerSquareCentimeter:
		return a.KilowattsPerSquareCentimeter()
    case IrradianceMegawattPerSquareCentimeter:
		return a.MegawattsPerSquareCentimeter()
	default:
		return 0
	}
}

func (a *Irradiance) convertFromBase(toUnit IrradianceUnits) float64 {
    value := a.value
	switch toUnit { 
	case IrradianceWattPerSquareMeter:
		return (value) 
	case IrradianceWattPerSquareCentimeter:
		return (value * 0.0001) 
	case IrradiancePicowattPerSquareMeter:
		return ((value) / 1e-12) 
	case IrradianceNanowattPerSquareMeter:
		return ((value) / 1e-09) 
	case IrradianceMicrowattPerSquareMeter:
		return ((value) / 1e-06) 
	case IrradianceMilliwattPerSquareMeter:
		return ((value) / 0.001) 
	case IrradianceKilowattPerSquareMeter:
		return ((value) / 1000.0) 
	case IrradianceMegawattPerSquareMeter:
		return ((value) / 1000000.0) 
	case IrradiancePicowattPerSquareCentimeter:
		return ((value * 0.0001) / 1e-12) 
	case IrradianceNanowattPerSquareCentimeter:
		return ((value * 0.0001) / 1e-09) 
	case IrradianceMicrowattPerSquareCentimeter:
		return ((value * 0.0001) / 1e-06) 
	case IrradianceMilliwattPerSquareCentimeter:
		return ((value * 0.0001) / 0.001) 
	case IrradianceKilowattPerSquareCentimeter:
		return ((value * 0.0001) / 1000.0) 
	case IrradianceMegawattPerSquareCentimeter:
		return ((value * 0.0001) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *Irradiance) convertToBase(value float64, fromUnit IrradianceUnits) float64 {
	switch fromUnit { 
	case IrradianceWattPerSquareMeter:
		return (value) 
	case IrradianceWattPerSquareCentimeter:
		return (value * 10000) 
	case IrradiancePicowattPerSquareMeter:
		return ((value) * 1e-12) 
	case IrradianceNanowattPerSquareMeter:
		return ((value) * 1e-09) 
	case IrradianceMicrowattPerSquareMeter:
		return ((value) * 1e-06) 
	case IrradianceMilliwattPerSquareMeter:
		return ((value) * 0.001) 
	case IrradianceKilowattPerSquareMeter:
		return ((value) * 1000.0) 
	case IrradianceMegawattPerSquareMeter:
		return ((value) * 1000000.0) 
	case IrradiancePicowattPerSquareCentimeter:
		return ((value * 10000) * 1e-12) 
	case IrradianceNanowattPerSquareCentimeter:
		return ((value * 10000) * 1e-09) 
	case IrradianceMicrowattPerSquareCentimeter:
		return ((value * 10000) * 1e-06) 
	case IrradianceMilliwattPerSquareCentimeter:
		return ((value * 10000) * 0.001) 
	case IrradianceKilowattPerSquareCentimeter:
		return ((value * 10000) * 1000.0) 
	case IrradianceMegawattPerSquareCentimeter:
		return ((value * 10000) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a Irradiance) String() string {
	return a.ToString(IrradianceWattPerSquareMeter, 2)
}

// ToString formats the Irradiance to string.
// fractionalDigits -1 for not mention
func (a *Irradiance) ToString(unit IrradianceUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Irradiance) getUnitAbbreviation(unit IrradianceUnits) string {
	switch unit { 
	case IrradianceWattPerSquareMeter:
		return "W/m²" 
	case IrradianceWattPerSquareCentimeter:
		return "W/cm²" 
	case IrradiancePicowattPerSquareMeter:
		return "pW/m²" 
	case IrradianceNanowattPerSquareMeter:
		return "nW/m²" 
	case IrradianceMicrowattPerSquareMeter:
		return "μW/m²" 
	case IrradianceMilliwattPerSquareMeter:
		return "mW/m²" 
	case IrradianceKilowattPerSquareMeter:
		return "kW/m²" 
	case IrradianceMegawattPerSquareMeter:
		return "MW/m²" 
	case IrradiancePicowattPerSquareCentimeter:
		return "pW/cm²" 
	case IrradianceNanowattPerSquareCentimeter:
		return "nW/cm²" 
	case IrradianceMicrowattPerSquareCentimeter:
		return "μW/cm²" 
	case IrradianceMilliwattPerSquareCentimeter:
		return "mW/cm²" 
	case IrradianceKilowattPerSquareCentimeter:
		return "kW/cm²" 
	case IrradianceMegawattPerSquareCentimeter:
		return "MW/cm²" 
	default:
		return ""
	}
}

// Check if the given Irradiance are equals to the current Irradiance
func (a *Irradiance) Equals(other *Irradiance) bool {
	return a.value == other.BaseValue()
}

// Check if the given Irradiance are equals to the current Irradiance
func (a *Irradiance) CompareTo(other *Irradiance) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given Irradiance to the current Irradiance.
func (a *Irradiance) Add(other *Irradiance) *Irradiance {
	return &Irradiance{value: a.value + other.BaseValue()}
}

// Subtract the given Irradiance to the current Irradiance.
func (a *Irradiance) Subtract(other *Irradiance) *Irradiance {
	return &Irradiance{value: a.value - other.BaseValue()}
}

// Multiply the given Irradiance to the current Irradiance.
func (a *Irradiance) Multiply(other *Irradiance) *Irradiance {
	return &Irradiance{value: a.value * other.BaseValue()}
}

// Divide the given Irradiance to the current Irradiance.
func (a *Irradiance) Divide(other *Irradiance) *Irradiance {
	return &Irradiance{value: a.value / other.BaseValue()}
}