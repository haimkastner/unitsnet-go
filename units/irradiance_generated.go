// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// IrradianceUnits defines various units of Irradiance.
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

// IrradianceDto represents a Irradiance measurement with a numerical value and its corresponding unit.
type IrradianceDto struct {
    // Value is the numerical representation of the Irradiance.
	Value float64
    // Unit specifies the unit of measurement for the Irradiance, as defined in the IrradianceUnits enumeration.
	Unit  IrradianceUnits
}

// IrradianceDtoFactory groups methods for creating and serializing IrradianceDto objects.
type IrradianceDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a IrradianceDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf IrradianceDtoFactory) FromJSON(data []byte) (*IrradianceDto, error) {
	a := IrradianceDto{}

    // Parse JSON into IrradianceDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a IrradianceDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a IrradianceDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}


// Irradiance represents a measurement in a of Irradiance.
//
// Irradiance is the intensity of ultraviolet (UV) or visible light incident on a surface.
type Irradiance struct {
	// value is the base measurement stored internally.
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

// IrradianceFactory groups methods for creating Irradiance instances.
type IrradianceFactory struct{}

// CreateIrradiance creates a new Irradiance instance from the given value and unit.
func (uf IrradianceFactory) CreateIrradiance(value float64, unit IrradianceUnits) (*Irradiance, error) {
	return newIrradiance(value, unit)
}

// FromDto converts a IrradianceDto to a Irradiance instance.
func (uf IrradianceFactory) FromDto(dto IrradianceDto) (*Irradiance, error) {
	return newIrradiance(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a Irradiance instance.
func (uf IrradianceFactory) FromDtoJSON(data []byte) (*Irradiance, error) {
	unitDto, err := IrradianceDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse IrradianceDto from JSON: %w", err)
	}
	return IrradianceFactory{}.FromDto(*unitDto)
}


// FromWattsPerSquareMeter creates a new Irradiance instance from a value in WattsPerSquareMeter.
func (uf IrradianceFactory) FromWattsPerSquareMeter(value float64) (*Irradiance, error) {
	return newIrradiance(value, IrradianceWattPerSquareMeter)
}

// FromWattsPerSquareCentimeter creates a new Irradiance instance from a value in WattsPerSquareCentimeter.
func (uf IrradianceFactory) FromWattsPerSquareCentimeter(value float64) (*Irradiance, error) {
	return newIrradiance(value, IrradianceWattPerSquareCentimeter)
}

// FromPicowattsPerSquareMeter creates a new Irradiance instance from a value in PicowattsPerSquareMeter.
func (uf IrradianceFactory) FromPicowattsPerSquareMeter(value float64) (*Irradiance, error) {
	return newIrradiance(value, IrradiancePicowattPerSquareMeter)
}

// FromNanowattsPerSquareMeter creates a new Irradiance instance from a value in NanowattsPerSquareMeter.
func (uf IrradianceFactory) FromNanowattsPerSquareMeter(value float64) (*Irradiance, error) {
	return newIrradiance(value, IrradianceNanowattPerSquareMeter)
}

// FromMicrowattsPerSquareMeter creates a new Irradiance instance from a value in MicrowattsPerSquareMeter.
func (uf IrradianceFactory) FromMicrowattsPerSquareMeter(value float64) (*Irradiance, error) {
	return newIrradiance(value, IrradianceMicrowattPerSquareMeter)
}

// FromMilliwattsPerSquareMeter creates a new Irradiance instance from a value in MilliwattsPerSquareMeter.
func (uf IrradianceFactory) FromMilliwattsPerSquareMeter(value float64) (*Irradiance, error) {
	return newIrradiance(value, IrradianceMilliwattPerSquareMeter)
}

// FromKilowattsPerSquareMeter creates a new Irradiance instance from a value in KilowattsPerSquareMeter.
func (uf IrradianceFactory) FromKilowattsPerSquareMeter(value float64) (*Irradiance, error) {
	return newIrradiance(value, IrradianceKilowattPerSquareMeter)
}

// FromMegawattsPerSquareMeter creates a new Irradiance instance from a value in MegawattsPerSquareMeter.
func (uf IrradianceFactory) FromMegawattsPerSquareMeter(value float64) (*Irradiance, error) {
	return newIrradiance(value, IrradianceMegawattPerSquareMeter)
}

// FromPicowattsPerSquareCentimeter creates a new Irradiance instance from a value in PicowattsPerSquareCentimeter.
func (uf IrradianceFactory) FromPicowattsPerSquareCentimeter(value float64) (*Irradiance, error) {
	return newIrradiance(value, IrradiancePicowattPerSquareCentimeter)
}

// FromNanowattsPerSquareCentimeter creates a new Irradiance instance from a value in NanowattsPerSquareCentimeter.
func (uf IrradianceFactory) FromNanowattsPerSquareCentimeter(value float64) (*Irradiance, error) {
	return newIrradiance(value, IrradianceNanowattPerSquareCentimeter)
}

// FromMicrowattsPerSquareCentimeter creates a new Irradiance instance from a value in MicrowattsPerSquareCentimeter.
func (uf IrradianceFactory) FromMicrowattsPerSquareCentimeter(value float64) (*Irradiance, error) {
	return newIrradiance(value, IrradianceMicrowattPerSquareCentimeter)
}

// FromMilliwattsPerSquareCentimeter creates a new Irradiance instance from a value in MilliwattsPerSquareCentimeter.
func (uf IrradianceFactory) FromMilliwattsPerSquareCentimeter(value float64) (*Irradiance, error) {
	return newIrradiance(value, IrradianceMilliwattPerSquareCentimeter)
}

// FromKilowattsPerSquareCentimeter creates a new Irradiance instance from a value in KilowattsPerSquareCentimeter.
func (uf IrradianceFactory) FromKilowattsPerSquareCentimeter(value float64) (*Irradiance, error) {
	return newIrradiance(value, IrradianceKilowattPerSquareCentimeter)
}

// FromMegawattsPerSquareCentimeter creates a new Irradiance instance from a value in MegawattsPerSquareCentimeter.
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

// BaseValue returns the base value of Irradiance in WattPerSquareMeter unit (the base/default quantity).
func (a *Irradiance) BaseValue() float64 {
	return a.value
}


// WattsPerSquareMeter returns the Irradiance value in WattsPerSquareMeter.
//
// 
func (a *Irradiance) WattsPerSquareMeter() float64 {
	if a.watts_per_square_meterLazy != nil {
		return *a.watts_per_square_meterLazy
	}
	watts_per_square_meter := a.convertFromBase(IrradianceWattPerSquareMeter)
	a.watts_per_square_meterLazy = &watts_per_square_meter
	return watts_per_square_meter
}

// WattsPerSquareCentimeter returns the Irradiance value in WattsPerSquareCentimeter.
//
// 
func (a *Irradiance) WattsPerSquareCentimeter() float64 {
	if a.watts_per_square_centimeterLazy != nil {
		return *a.watts_per_square_centimeterLazy
	}
	watts_per_square_centimeter := a.convertFromBase(IrradianceWattPerSquareCentimeter)
	a.watts_per_square_centimeterLazy = &watts_per_square_centimeter
	return watts_per_square_centimeter
}

// PicowattsPerSquareMeter returns the Irradiance value in PicowattsPerSquareMeter.
//
// 
func (a *Irradiance) PicowattsPerSquareMeter() float64 {
	if a.picowatts_per_square_meterLazy != nil {
		return *a.picowatts_per_square_meterLazy
	}
	picowatts_per_square_meter := a.convertFromBase(IrradiancePicowattPerSquareMeter)
	a.picowatts_per_square_meterLazy = &picowatts_per_square_meter
	return picowatts_per_square_meter
}

// NanowattsPerSquareMeter returns the Irradiance value in NanowattsPerSquareMeter.
//
// 
func (a *Irradiance) NanowattsPerSquareMeter() float64 {
	if a.nanowatts_per_square_meterLazy != nil {
		return *a.nanowatts_per_square_meterLazy
	}
	nanowatts_per_square_meter := a.convertFromBase(IrradianceNanowattPerSquareMeter)
	a.nanowatts_per_square_meterLazy = &nanowatts_per_square_meter
	return nanowatts_per_square_meter
}

// MicrowattsPerSquareMeter returns the Irradiance value in MicrowattsPerSquareMeter.
//
// 
func (a *Irradiance) MicrowattsPerSquareMeter() float64 {
	if a.microwatts_per_square_meterLazy != nil {
		return *a.microwatts_per_square_meterLazy
	}
	microwatts_per_square_meter := a.convertFromBase(IrradianceMicrowattPerSquareMeter)
	a.microwatts_per_square_meterLazy = &microwatts_per_square_meter
	return microwatts_per_square_meter
}

// MilliwattsPerSquareMeter returns the Irradiance value in MilliwattsPerSquareMeter.
//
// 
func (a *Irradiance) MilliwattsPerSquareMeter() float64 {
	if a.milliwatts_per_square_meterLazy != nil {
		return *a.milliwatts_per_square_meterLazy
	}
	milliwatts_per_square_meter := a.convertFromBase(IrradianceMilliwattPerSquareMeter)
	a.milliwatts_per_square_meterLazy = &milliwatts_per_square_meter
	return milliwatts_per_square_meter
}

// KilowattsPerSquareMeter returns the Irradiance value in KilowattsPerSquareMeter.
//
// 
func (a *Irradiance) KilowattsPerSquareMeter() float64 {
	if a.kilowatts_per_square_meterLazy != nil {
		return *a.kilowatts_per_square_meterLazy
	}
	kilowatts_per_square_meter := a.convertFromBase(IrradianceKilowattPerSquareMeter)
	a.kilowatts_per_square_meterLazy = &kilowatts_per_square_meter
	return kilowatts_per_square_meter
}

// MegawattsPerSquareMeter returns the Irradiance value in MegawattsPerSquareMeter.
//
// 
func (a *Irradiance) MegawattsPerSquareMeter() float64 {
	if a.megawatts_per_square_meterLazy != nil {
		return *a.megawatts_per_square_meterLazy
	}
	megawatts_per_square_meter := a.convertFromBase(IrradianceMegawattPerSquareMeter)
	a.megawatts_per_square_meterLazy = &megawatts_per_square_meter
	return megawatts_per_square_meter
}

// PicowattsPerSquareCentimeter returns the Irradiance value in PicowattsPerSquareCentimeter.
//
// 
func (a *Irradiance) PicowattsPerSquareCentimeter() float64 {
	if a.picowatts_per_square_centimeterLazy != nil {
		return *a.picowatts_per_square_centimeterLazy
	}
	picowatts_per_square_centimeter := a.convertFromBase(IrradiancePicowattPerSquareCentimeter)
	a.picowatts_per_square_centimeterLazy = &picowatts_per_square_centimeter
	return picowatts_per_square_centimeter
}

// NanowattsPerSquareCentimeter returns the Irradiance value in NanowattsPerSquareCentimeter.
//
// 
func (a *Irradiance) NanowattsPerSquareCentimeter() float64 {
	if a.nanowatts_per_square_centimeterLazy != nil {
		return *a.nanowatts_per_square_centimeterLazy
	}
	nanowatts_per_square_centimeter := a.convertFromBase(IrradianceNanowattPerSquareCentimeter)
	a.nanowatts_per_square_centimeterLazy = &nanowatts_per_square_centimeter
	return nanowatts_per_square_centimeter
}

// MicrowattsPerSquareCentimeter returns the Irradiance value in MicrowattsPerSquareCentimeter.
//
// 
func (a *Irradiance) MicrowattsPerSquareCentimeter() float64 {
	if a.microwatts_per_square_centimeterLazy != nil {
		return *a.microwatts_per_square_centimeterLazy
	}
	microwatts_per_square_centimeter := a.convertFromBase(IrradianceMicrowattPerSquareCentimeter)
	a.microwatts_per_square_centimeterLazy = &microwatts_per_square_centimeter
	return microwatts_per_square_centimeter
}

// MilliwattsPerSquareCentimeter returns the Irradiance value in MilliwattsPerSquareCentimeter.
//
// 
func (a *Irradiance) MilliwattsPerSquareCentimeter() float64 {
	if a.milliwatts_per_square_centimeterLazy != nil {
		return *a.milliwatts_per_square_centimeterLazy
	}
	milliwatts_per_square_centimeter := a.convertFromBase(IrradianceMilliwattPerSquareCentimeter)
	a.milliwatts_per_square_centimeterLazy = &milliwatts_per_square_centimeter
	return milliwatts_per_square_centimeter
}

// KilowattsPerSquareCentimeter returns the Irradiance value in KilowattsPerSquareCentimeter.
//
// 
func (a *Irradiance) KilowattsPerSquareCentimeter() float64 {
	if a.kilowatts_per_square_centimeterLazy != nil {
		return *a.kilowatts_per_square_centimeterLazy
	}
	kilowatts_per_square_centimeter := a.convertFromBase(IrradianceKilowattPerSquareCentimeter)
	a.kilowatts_per_square_centimeterLazy = &kilowatts_per_square_centimeter
	return kilowatts_per_square_centimeter
}

// MegawattsPerSquareCentimeter returns the Irradiance value in MegawattsPerSquareCentimeter.
//
// 
func (a *Irradiance) MegawattsPerSquareCentimeter() float64 {
	if a.megawatts_per_square_centimeterLazy != nil {
		return *a.megawatts_per_square_centimeterLazy
	}
	megawatts_per_square_centimeter := a.convertFromBase(IrradianceMegawattPerSquareCentimeter)
	a.megawatts_per_square_centimeterLazy = &megawatts_per_square_centimeter
	return megawatts_per_square_centimeter
}


// ToDto creates a IrradianceDto representation from the Irradiance instance.
//
// If the provided holdInUnit is nil, the value will be repesented by WattPerSquareMeter by default.
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

// ToDtoJSON creates a JSON representation of the Irradiance instance.
//
// If the provided holdInUnit is nil, the value will be repesented by WattPerSquareMeter by default.
func (a *Irradiance) ToDtoJSON(holdInUnit *IrradianceUnits) ([]byte, error) {
	// Convert to IrradianceDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a Irradiance to a specific unit value.
// The function uses the provided unit type (IrradianceUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
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
		return math.NaN()
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

// String returns a string representation of the Irradiance in the default unit (WattPerSquareMeter),
// formatted to two decimal places.
func (a Irradiance) String() string {
	return a.ToString(IrradianceWattPerSquareMeter, 2)
}

// ToString formats the Irradiance value as a string with the specified unit and fractional digits.
// It converts the Irradiance to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the Irradiance value will be converted (e.g., WattPerSquareMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the Irradiance with the unit abbreviation.
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

// Equals checks if the given Irradiance is equal to the current Irradiance.
//
// Parameters:
//    other: The Irradiance to compare against.
//
// Returns:
//    bool: Returns true if both Irradiance are equal, false otherwise.
func (a *Irradiance) Equals(other *Irradiance) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current Irradiance with another Irradiance.
// It returns -1 if the current Irradiance is less than the other Irradiance, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The Irradiance to compare against.
//
// Returns:
//    int: -1 if the current Irradiance is less, 1 if greater, and 0 if equal.
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

// Add adds the given Irradiance to the current Irradiance and returns the result.
// The result is a new Irradiance instance with the sum of the values.
//
// Parameters:
//    other: The Irradiance to add to the current Irradiance.
//
// Returns:
//    *Irradiance: A new Irradiance instance representing the sum of both Irradiance.
func (a *Irradiance) Add(other *Irradiance) *Irradiance {
	return &Irradiance{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given Irradiance from the current Irradiance and returns the result.
// The result is a new Irradiance instance with the difference of the values.
//
// Parameters:
//    other: The Irradiance to subtract from the current Irradiance.
//
// Returns:
//    *Irradiance: A new Irradiance instance representing the difference of both Irradiance.
func (a *Irradiance) Subtract(other *Irradiance) *Irradiance {
	return &Irradiance{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current Irradiance by the given Irradiance and returns the result.
// The result is a new Irradiance instance with the product of the values.
//
// Parameters:
//    other: The Irradiance to multiply with the current Irradiance.
//
// Returns:
//    *Irradiance: A new Irradiance instance representing the product of both Irradiance.
func (a *Irradiance) Multiply(other *Irradiance) *Irradiance {
	return &Irradiance{value: a.value * other.BaseValue()}
}

// Divide divides the current Irradiance by the given Irradiance and returns the result.
// The result is a new Irradiance instance with the quotient of the values.
//
// Parameters:
//    other: The Irradiance to divide the current Irradiance by.
//
// Returns:
//    *Irradiance: A new Irradiance instance representing the quotient of both Irradiance.
func (a *Irradiance) Divide(other *Irradiance) *Irradiance {
	return &Irradiance{value: a.value / other.BaseValue()}
}