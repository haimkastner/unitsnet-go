// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// DoseAreaProductUnits defines various units of DoseAreaProduct.
type DoseAreaProductUnits string

const (
    
        // 
        DoseAreaProductGraySquareMeter DoseAreaProductUnits = "GraySquareMeter"
        // 
        DoseAreaProductGraySquareDecimeter DoseAreaProductUnits = "GraySquareDecimeter"
        // 
        DoseAreaProductGraySquareCentimeter DoseAreaProductUnits = "GraySquareCentimeter"
        // 
        DoseAreaProductGraySquareMillimeter DoseAreaProductUnits = "GraySquareMillimeter"
        // 
        DoseAreaProductMicrograySquareMeter DoseAreaProductUnits = "MicrograySquareMeter"
        // 
        DoseAreaProductMilligraySquareMeter DoseAreaProductUnits = "MilligraySquareMeter"
        // 
        DoseAreaProductCentigraySquareMeter DoseAreaProductUnits = "CentigraySquareMeter"
        // 
        DoseAreaProductDecigraySquareMeter DoseAreaProductUnits = "DecigraySquareMeter"
        // 
        DoseAreaProductMicrograySquareDecimeter DoseAreaProductUnits = "MicrograySquareDecimeter"
        // 
        DoseAreaProductMilligraySquareDecimeter DoseAreaProductUnits = "MilligraySquareDecimeter"
        // 
        DoseAreaProductCentigraySquareDecimeter DoseAreaProductUnits = "CentigraySquareDecimeter"
        // 
        DoseAreaProductDecigraySquareDecimeter DoseAreaProductUnits = "DecigraySquareDecimeter"
        // 
        DoseAreaProductMicrograySquareCentimeter DoseAreaProductUnits = "MicrograySquareCentimeter"
        // 
        DoseAreaProductMilligraySquareCentimeter DoseAreaProductUnits = "MilligraySquareCentimeter"
        // 
        DoseAreaProductCentigraySquareCentimeter DoseAreaProductUnits = "CentigraySquareCentimeter"
        // 
        DoseAreaProductDecigraySquareCentimeter DoseAreaProductUnits = "DecigraySquareCentimeter"
        // 
        DoseAreaProductMicrograySquareMillimeter DoseAreaProductUnits = "MicrograySquareMillimeter"
        // 
        DoseAreaProductMilligraySquareMillimeter DoseAreaProductUnits = "MilligraySquareMillimeter"
        // 
        DoseAreaProductCentigraySquareMillimeter DoseAreaProductUnits = "CentigraySquareMillimeter"
        // 
        DoseAreaProductDecigraySquareMillimeter DoseAreaProductUnits = "DecigraySquareMillimeter"
)

var internalDoseAreaProductUnitsMap = map[DoseAreaProductUnits]bool{
	
	DoseAreaProductGraySquareMeter: true,
	DoseAreaProductGraySquareDecimeter: true,
	DoseAreaProductGraySquareCentimeter: true,
	DoseAreaProductGraySquareMillimeter: true,
	DoseAreaProductMicrograySquareMeter: true,
	DoseAreaProductMilligraySquareMeter: true,
	DoseAreaProductCentigraySquareMeter: true,
	DoseAreaProductDecigraySquareMeter: true,
	DoseAreaProductMicrograySquareDecimeter: true,
	DoseAreaProductMilligraySquareDecimeter: true,
	DoseAreaProductCentigraySquareDecimeter: true,
	DoseAreaProductDecigraySquareDecimeter: true,
	DoseAreaProductMicrograySquareCentimeter: true,
	DoseAreaProductMilligraySquareCentimeter: true,
	DoseAreaProductCentigraySquareCentimeter: true,
	DoseAreaProductDecigraySquareCentimeter: true,
	DoseAreaProductMicrograySquareMillimeter: true,
	DoseAreaProductMilligraySquareMillimeter: true,
	DoseAreaProductCentigraySquareMillimeter: true,
	DoseAreaProductDecigraySquareMillimeter: true,
}

// DoseAreaProductDto represents a DoseAreaProduct measurement with a numerical value and its corresponding unit.
type DoseAreaProductDto struct {
    // Value is the numerical representation of the DoseAreaProduct.
	Value float64 `json:"value" validate:"required"`
    // Unit specifies the unit of measurement for the DoseAreaProduct, as defined in the DoseAreaProductUnits enumeration.
	Unit  DoseAreaProductUnits `json:"unit" validate:"required,oneof=GraySquareMeter,GraySquareDecimeter,GraySquareCentimeter,GraySquareMillimeter,MicrograySquareMeter,MilligraySquareMeter,CentigraySquareMeter,DecigraySquareMeter,MicrograySquareDecimeter,MilligraySquareDecimeter,CentigraySquareDecimeter,DecigraySquareDecimeter,MicrograySquareCentimeter,MilligraySquareCentimeter,CentigraySquareCentimeter,DecigraySquareCentimeter,MicrograySquareMillimeter,MilligraySquareMillimeter,CentigraySquareMillimeter,DecigraySquareMillimeter"`
}

// DoseAreaProductDtoFactory groups methods for creating and serializing DoseAreaProductDto objects.
type DoseAreaProductDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a DoseAreaProductDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf DoseAreaProductDtoFactory) FromJSON(data []byte) (*DoseAreaProductDto, error) {
	a := DoseAreaProductDto{}

    // Parse JSON into DoseAreaProductDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a DoseAreaProductDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a DoseAreaProductDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// DoseAreaProduct represents a measurement in a of DoseAreaProduct.
//
// It is defined as the absorbed dose multiplied by the area irradiated.
type DoseAreaProduct struct {
	// value is the base measurement stored internally.
	value       float64
    
    gray_square_metersLazy *float64 
    gray_square_decimetersLazy *float64 
    gray_square_centimetersLazy *float64 
    gray_square_millimetersLazy *float64 
    microgray_square_metersLazy *float64 
    milligray_square_metersLazy *float64 
    centigray_square_metersLazy *float64 
    decigray_square_metersLazy *float64 
    microgray_square_decimetersLazy *float64 
    milligray_square_decimetersLazy *float64 
    centigray_square_decimetersLazy *float64 
    decigray_square_decimetersLazy *float64 
    microgray_square_centimetersLazy *float64 
    milligray_square_centimetersLazy *float64 
    centigray_square_centimetersLazy *float64 
    decigray_square_centimetersLazy *float64 
    microgray_square_millimetersLazy *float64 
    milligray_square_millimetersLazy *float64 
    centigray_square_millimetersLazy *float64 
    decigray_square_millimetersLazy *float64 
}

// DoseAreaProductFactory groups methods for creating DoseAreaProduct instances.
type DoseAreaProductFactory struct{}

// CreateDoseAreaProduct creates a new DoseAreaProduct instance from the given value and unit.
func (uf DoseAreaProductFactory) CreateDoseAreaProduct(value float64, unit DoseAreaProductUnits) (*DoseAreaProduct, error) {
	return newDoseAreaProduct(value, unit)
}

// FromDto converts a DoseAreaProductDto to a DoseAreaProduct instance.
func (uf DoseAreaProductFactory) FromDto(dto DoseAreaProductDto) (*DoseAreaProduct, error) {
	return newDoseAreaProduct(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a DoseAreaProduct instance.
func (uf DoseAreaProductFactory) FromDtoJSON(data []byte) (*DoseAreaProduct, error) {
	unitDto, err := DoseAreaProductDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse DoseAreaProductDto from JSON: %w", err)
	}
	return DoseAreaProductFactory{}.FromDto(*unitDto)
}


// FromGraySquareMeters creates a new DoseAreaProduct instance from a value in GraySquareMeters.
func (uf DoseAreaProductFactory) FromGraySquareMeters(value float64) (*DoseAreaProduct, error) {
	return newDoseAreaProduct(value, DoseAreaProductGraySquareMeter)
}

// FromGraySquareDecimeters creates a new DoseAreaProduct instance from a value in GraySquareDecimeters.
func (uf DoseAreaProductFactory) FromGraySquareDecimeters(value float64) (*DoseAreaProduct, error) {
	return newDoseAreaProduct(value, DoseAreaProductGraySquareDecimeter)
}

// FromGraySquareCentimeters creates a new DoseAreaProduct instance from a value in GraySquareCentimeters.
func (uf DoseAreaProductFactory) FromGraySquareCentimeters(value float64) (*DoseAreaProduct, error) {
	return newDoseAreaProduct(value, DoseAreaProductGraySquareCentimeter)
}

// FromGraySquareMillimeters creates a new DoseAreaProduct instance from a value in GraySquareMillimeters.
func (uf DoseAreaProductFactory) FromGraySquareMillimeters(value float64) (*DoseAreaProduct, error) {
	return newDoseAreaProduct(value, DoseAreaProductGraySquareMillimeter)
}

// FromMicrograySquareMeters creates a new DoseAreaProduct instance from a value in MicrograySquareMeters.
func (uf DoseAreaProductFactory) FromMicrograySquareMeters(value float64) (*DoseAreaProduct, error) {
	return newDoseAreaProduct(value, DoseAreaProductMicrograySquareMeter)
}

// FromMilligraySquareMeters creates a new DoseAreaProduct instance from a value in MilligraySquareMeters.
func (uf DoseAreaProductFactory) FromMilligraySquareMeters(value float64) (*DoseAreaProduct, error) {
	return newDoseAreaProduct(value, DoseAreaProductMilligraySquareMeter)
}

// FromCentigraySquareMeters creates a new DoseAreaProduct instance from a value in CentigraySquareMeters.
func (uf DoseAreaProductFactory) FromCentigraySquareMeters(value float64) (*DoseAreaProduct, error) {
	return newDoseAreaProduct(value, DoseAreaProductCentigraySquareMeter)
}

// FromDecigraySquareMeters creates a new DoseAreaProduct instance from a value in DecigraySquareMeters.
func (uf DoseAreaProductFactory) FromDecigraySquareMeters(value float64) (*DoseAreaProduct, error) {
	return newDoseAreaProduct(value, DoseAreaProductDecigraySquareMeter)
}

// FromMicrograySquareDecimeters creates a new DoseAreaProduct instance from a value in MicrograySquareDecimeters.
func (uf DoseAreaProductFactory) FromMicrograySquareDecimeters(value float64) (*DoseAreaProduct, error) {
	return newDoseAreaProduct(value, DoseAreaProductMicrograySquareDecimeter)
}

// FromMilligraySquareDecimeters creates a new DoseAreaProduct instance from a value in MilligraySquareDecimeters.
func (uf DoseAreaProductFactory) FromMilligraySquareDecimeters(value float64) (*DoseAreaProduct, error) {
	return newDoseAreaProduct(value, DoseAreaProductMilligraySquareDecimeter)
}

// FromCentigraySquareDecimeters creates a new DoseAreaProduct instance from a value in CentigraySquareDecimeters.
func (uf DoseAreaProductFactory) FromCentigraySquareDecimeters(value float64) (*DoseAreaProduct, error) {
	return newDoseAreaProduct(value, DoseAreaProductCentigraySquareDecimeter)
}

// FromDecigraySquareDecimeters creates a new DoseAreaProduct instance from a value in DecigraySquareDecimeters.
func (uf DoseAreaProductFactory) FromDecigraySquareDecimeters(value float64) (*DoseAreaProduct, error) {
	return newDoseAreaProduct(value, DoseAreaProductDecigraySquareDecimeter)
}

// FromMicrograySquareCentimeters creates a new DoseAreaProduct instance from a value in MicrograySquareCentimeters.
func (uf DoseAreaProductFactory) FromMicrograySquareCentimeters(value float64) (*DoseAreaProduct, error) {
	return newDoseAreaProduct(value, DoseAreaProductMicrograySquareCentimeter)
}

// FromMilligraySquareCentimeters creates a new DoseAreaProduct instance from a value in MilligraySquareCentimeters.
func (uf DoseAreaProductFactory) FromMilligraySquareCentimeters(value float64) (*DoseAreaProduct, error) {
	return newDoseAreaProduct(value, DoseAreaProductMilligraySquareCentimeter)
}

// FromCentigraySquareCentimeters creates a new DoseAreaProduct instance from a value in CentigraySquareCentimeters.
func (uf DoseAreaProductFactory) FromCentigraySquareCentimeters(value float64) (*DoseAreaProduct, error) {
	return newDoseAreaProduct(value, DoseAreaProductCentigraySquareCentimeter)
}

// FromDecigraySquareCentimeters creates a new DoseAreaProduct instance from a value in DecigraySquareCentimeters.
func (uf DoseAreaProductFactory) FromDecigraySquareCentimeters(value float64) (*DoseAreaProduct, error) {
	return newDoseAreaProduct(value, DoseAreaProductDecigraySquareCentimeter)
}

// FromMicrograySquareMillimeters creates a new DoseAreaProduct instance from a value in MicrograySquareMillimeters.
func (uf DoseAreaProductFactory) FromMicrograySquareMillimeters(value float64) (*DoseAreaProduct, error) {
	return newDoseAreaProduct(value, DoseAreaProductMicrograySquareMillimeter)
}

// FromMilligraySquareMillimeters creates a new DoseAreaProduct instance from a value in MilligraySquareMillimeters.
func (uf DoseAreaProductFactory) FromMilligraySquareMillimeters(value float64) (*DoseAreaProduct, error) {
	return newDoseAreaProduct(value, DoseAreaProductMilligraySquareMillimeter)
}

// FromCentigraySquareMillimeters creates a new DoseAreaProduct instance from a value in CentigraySquareMillimeters.
func (uf DoseAreaProductFactory) FromCentigraySquareMillimeters(value float64) (*DoseAreaProduct, error) {
	return newDoseAreaProduct(value, DoseAreaProductCentigraySquareMillimeter)
}

// FromDecigraySquareMillimeters creates a new DoseAreaProduct instance from a value in DecigraySquareMillimeters.
func (uf DoseAreaProductFactory) FromDecigraySquareMillimeters(value float64) (*DoseAreaProduct, error) {
	return newDoseAreaProduct(value, DoseAreaProductDecigraySquareMillimeter)
}


// newDoseAreaProduct creates a new DoseAreaProduct.
func newDoseAreaProduct(value float64, fromUnit DoseAreaProductUnits) (*DoseAreaProduct, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalDoseAreaProductUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in DoseAreaProductUnits", fromUnit)
	}
	a := &DoseAreaProduct{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of DoseAreaProduct in GraySquareMeter unit (the base/default quantity).
func (a *DoseAreaProduct) BaseValue() float64 {
	return a.value
}


// GraySquareMeters returns the DoseAreaProduct value in GraySquareMeters.
//
// 
func (a *DoseAreaProduct) GraySquareMeters() float64 {
	if a.gray_square_metersLazy != nil {
		return *a.gray_square_metersLazy
	}
	gray_square_meters := a.convertFromBase(DoseAreaProductGraySquareMeter)
	a.gray_square_metersLazy = &gray_square_meters
	return gray_square_meters
}

// GraySquareDecimeters returns the DoseAreaProduct value in GraySquareDecimeters.
//
// 
func (a *DoseAreaProduct) GraySquareDecimeters() float64 {
	if a.gray_square_decimetersLazy != nil {
		return *a.gray_square_decimetersLazy
	}
	gray_square_decimeters := a.convertFromBase(DoseAreaProductGraySquareDecimeter)
	a.gray_square_decimetersLazy = &gray_square_decimeters
	return gray_square_decimeters
}

// GraySquareCentimeters returns the DoseAreaProduct value in GraySquareCentimeters.
//
// 
func (a *DoseAreaProduct) GraySquareCentimeters() float64 {
	if a.gray_square_centimetersLazy != nil {
		return *a.gray_square_centimetersLazy
	}
	gray_square_centimeters := a.convertFromBase(DoseAreaProductGraySquareCentimeter)
	a.gray_square_centimetersLazy = &gray_square_centimeters
	return gray_square_centimeters
}

// GraySquareMillimeters returns the DoseAreaProduct value in GraySquareMillimeters.
//
// 
func (a *DoseAreaProduct) GraySquareMillimeters() float64 {
	if a.gray_square_millimetersLazy != nil {
		return *a.gray_square_millimetersLazy
	}
	gray_square_millimeters := a.convertFromBase(DoseAreaProductGraySquareMillimeter)
	a.gray_square_millimetersLazy = &gray_square_millimeters
	return gray_square_millimeters
}

// MicrograySquareMeters returns the DoseAreaProduct value in MicrograySquareMeters.
//
// 
func (a *DoseAreaProduct) MicrograySquareMeters() float64 {
	if a.microgray_square_metersLazy != nil {
		return *a.microgray_square_metersLazy
	}
	microgray_square_meters := a.convertFromBase(DoseAreaProductMicrograySquareMeter)
	a.microgray_square_metersLazy = &microgray_square_meters
	return microgray_square_meters
}

// MilligraySquareMeters returns the DoseAreaProduct value in MilligraySquareMeters.
//
// 
func (a *DoseAreaProduct) MilligraySquareMeters() float64 {
	if a.milligray_square_metersLazy != nil {
		return *a.milligray_square_metersLazy
	}
	milligray_square_meters := a.convertFromBase(DoseAreaProductMilligraySquareMeter)
	a.milligray_square_metersLazy = &milligray_square_meters
	return milligray_square_meters
}

// CentigraySquareMeters returns the DoseAreaProduct value in CentigraySquareMeters.
//
// 
func (a *DoseAreaProduct) CentigraySquareMeters() float64 {
	if a.centigray_square_metersLazy != nil {
		return *a.centigray_square_metersLazy
	}
	centigray_square_meters := a.convertFromBase(DoseAreaProductCentigraySquareMeter)
	a.centigray_square_metersLazy = &centigray_square_meters
	return centigray_square_meters
}

// DecigraySquareMeters returns the DoseAreaProduct value in DecigraySquareMeters.
//
// 
func (a *DoseAreaProduct) DecigraySquareMeters() float64 {
	if a.decigray_square_metersLazy != nil {
		return *a.decigray_square_metersLazy
	}
	decigray_square_meters := a.convertFromBase(DoseAreaProductDecigraySquareMeter)
	a.decigray_square_metersLazy = &decigray_square_meters
	return decigray_square_meters
}

// MicrograySquareDecimeters returns the DoseAreaProduct value in MicrograySquareDecimeters.
//
// 
func (a *DoseAreaProduct) MicrograySquareDecimeters() float64 {
	if a.microgray_square_decimetersLazy != nil {
		return *a.microgray_square_decimetersLazy
	}
	microgray_square_decimeters := a.convertFromBase(DoseAreaProductMicrograySquareDecimeter)
	a.microgray_square_decimetersLazy = &microgray_square_decimeters
	return microgray_square_decimeters
}

// MilligraySquareDecimeters returns the DoseAreaProduct value in MilligraySquareDecimeters.
//
// 
func (a *DoseAreaProduct) MilligraySquareDecimeters() float64 {
	if a.milligray_square_decimetersLazy != nil {
		return *a.milligray_square_decimetersLazy
	}
	milligray_square_decimeters := a.convertFromBase(DoseAreaProductMilligraySquareDecimeter)
	a.milligray_square_decimetersLazy = &milligray_square_decimeters
	return milligray_square_decimeters
}

// CentigraySquareDecimeters returns the DoseAreaProduct value in CentigraySquareDecimeters.
//
// 
func (a *DoseAreaProduct) CentigraySquareDecimeters() float64 {
	if a.centigray_square_decimetersLazy != nil {
		return *a.centigray_square_decimetersLazy
	}
	centigray_square_decimeters := a.convertFromBase(DoseAreaProductCentigraySquareDecimeter)
	a.centigray_square_decimetersLazy = &centigray_square_decimeters
	return centigray_square_decimeters
}

// DecigraySquareDecimeters returns the DoseAreaProduct value in DecigraySquareDecimeters.
//
// 
func (a *DoseAreaProduct) DecigraySquareDecimeters() float64 {
	if a.decigray_square_decimetersLazy != nil {
		return *a.decigray_square_decimetersLazy
	}
	decigray_square_decimeters := a.convertFromBase(DoseAreaProductDecigraySquareDecimeter)
	a.decigray_square_decimetersLazy = &decigray_square_decimeters
	return decigray_square_decimeters
}

// MicrograySquareCentimeters returns the DoseAreaProduct value in MicrograySquareCentimeters.
//
// 
func (a *DoseAreaProduct) MicrograySquareCentimeters() float64 {
	if a.microgray_square_centimetersLazy != nil {
		return *a.microgray_square_centimetersLazy
	}
	microgray_square_centimeters := a.convertFromBase(DoseAreaProductMicrograySquareCentimeter)
	a.microgray_square_centimetersLazy = &microgray_square_centimeters
	return microgray_square_centimeters
}

// MilligraySquareCentimeters returns the DoseAreaProduct value in MilligraySquareCentimeters.
//
// 
func (a *DoseAreaProduct) MilligraySquareCentimeters() float64 {
	if a.milligray_square_centimetersLazy != nil {
		return *a.milligray_square_centimetersLazy
	}
	milligray_square_centimeters := a.convertFromBase(DoseAreaProductMilligraySquareCentimeter)
	a.milligray_square_centimetersLazy = &milligray_square_centimeters
	return milligray_square_centimeters
}

// CentigraySquareCentimeters returns the DoseAreaProduct value in CentigraySquareCentimeters.
//
// 
func (a *DoseAreaProduct) CentigraySquareCentimeters() float64 {
	if a.centigray_square_centimetersLazy != nil {
		return *a.centigray_square_centimetersLazy
	}
	centigray_square_centimeters := a.convertFromBase(DoseAreaProductCentigraySquareCentimeter)
	a.centigray_square_centimetersLazy = &centigray_square_centimeters
	return centigray_square_centimeters
}

// DecigraySquareCentimeters returns the DoseAreaProduct value in DecigraySquareCentimeters.
//
// 
func (a *DoseAreaProduct) DecigraySquareCentimeters() float64 {
	if a.decigray_square_centimetersLazy != nil {
		return *a.decigray_square_centimetersLazy
	}
	decigray_square_centimeters := a.convertFromBase(DoseAreaProductDecigraySquareCentimeter)
	a.decigray_square_centimetersLazy = &decigray_square_centimeters
	return decigray_square_centimeters
}

// MicrograySquareMillimeters returns the DoseAreaProduct value in MicrograySquareMillimeters.
//
// 
func (a *DoseAreaProduct) MicrograySquareMillimeters() float64 {
	if a.microgray_square_millimetersLazy != nil {
		return *a.microgray_square_millimetersLazy
	}
	microgray_square_millimeters := a.convertFromBase(DoseAreaProductMicrograySquareMillimeter)
	a.microgray_square_millimetersLazy = &microgray_square_millimeters
	return microgray_square_millimeters
}

// MilligraySquareMillimeters returns the DoseAreaProduct value in MilligraySquareMillimeters.
//
// 
func (a *DoseAreaProduct) MilligraySquareMillimeters() float64 {
	if a.milligray_square_millimetersLazy != nil {
		return *a.milligray_square_millimetersLazy
	}
	milligray_square_millimeters := a.convertFromBase(DoseAreaProductMilligraySquareMillimeter)
	a.milligray_square_millimetersLazy = &milligray_square_millimeters
	return milligray_square_millimeters
}

// CentigraySquareMillimeters returns the DoseAreaProduct value in CentigraySquareMillimeters.
//
// 
func (a *DoseAreaProduct) CentigraySquareMillimeters() float64 {
	if a.centigray_square_millimetersLazy != nil {
		return *a.centigray_square_millimetersLazy
	}
	centigray_square_millimeters := a.convertFromBase(DoseAreaProductCentigraySquareMillimeter)
	a.centigray_square_millimetersLazy = &centigray_square_millimeters
	return centigray_square_millimeters
}

// DecigraySquareMillimeters returns the DoseAreaProduct value in DecigraySquareMillimeters.
//
// 
func (a *DoseAreaProduct) DecigraySquareMillimeters() float64 {
	if a.decigray_square_millimetersLazy != nil {
		return *a.decigray_square_millimetersLazy
	}
	decigray_square_millimeters := a.convertFromBase(DoseAreaProductDecigraySquareMillimeter)
	a.decigray_square_millimetersLazy = &decigray_square_millimeters
	return decigray_square_millimeters
}


// ToDto creates a DoseAreaProductDto representation from the DoseAreaProduct instance.
//
// If the provided holdInUnit is nil, the value will be repesented by GraySquareMeter by default.
func (a *DoseAreaProduct) ToDto(holdInUnit *DoseAreaProductUnits) DoseAreaProductDto {
	if holdInUnit == nil {
		defaultUnit := DoseAreaProductGraySquareMeter // Default value
		holdInUnit = &defaultUnit
	}

	return DoseAreaProductDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the DoseAreaProduct instance.
//
// If the provided holdInUnit is nil, the value will be repesented by GraySquareMeter by default.
func (a *DoseAreaProduct) ToDtoJSON(holdInUnit *DoseAreaProductUnits) ([]byte, error) {
	// Convert to DoseAreaProductDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a DoseAreaProduct to a specific unit value.
// The function uses the provided unit type (DoseAreaProductUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *DoseAreaProduct) Convert(toUnit DoseAreaProductUnits) float64 {
	switch toUnit { 
    case DoseAreaProductGraySquareMeter:
		return a.GraySquareMeters()
    case DoseAreaProductGraySquareDecimeter:
		return a.GraySquareDecimeters()
    case DoseAreaProductGraySquareCentimeter:
		return a.GraySquareCentimeters()
    case DoseAreaProductGraySquareMillimeter:
		return a.GraySquareMillimeters()
    case DoseAreaProductMicrograySquareMeter:
		return a.MicrograySquareMeters()
    case DoseAreaProductMilligraySquareMeter:
		return a.MilligraySquareMeters()
    case DoseAreaProductCentigraySquareMeter:
		return a.CentigraySquareMeters()
    case DoseAreaProductDecigraySquareMeter:
		return a.DecigraySquareMeters()
    case DoseAreaProductMicrograySquareDecimeter:
		return a.MicrograySquareDecimeters()
    case DoseAreaProductMilligraySquareDecimeter:
		return a.MilligraySquareDecimeters()
    case DoseAreaProductCentigraySquareDecimeter:
		return a.CentigraySquareDecimeters()
    case DoseAreaProductDecigraySquareDecimeter:
		return a.DecigraySquareDecimeters()
    case DoseAreaProductMicrograySquareCentimeter:
		return a.MicrograySquareCentimeters()
    case DoseAreaProductMilligraySquareCentimeter:
		return a.MilligraySquareCentimeters()
    case DoseAreaProductCentigraySquareCentimeter:
		return a.CentigraySquareCentimeters()
    case DoseAreaProductDecigraySquareCentimeter:
		return a.DecigraySquareCentimeters()
    case DoseAreaProductMicrograySquareMillimeter:
		return a.MicrograySquareMillimeters()
    case DoseAreaProductMilligraySquareMillimeter:
		return a.MilligraySquareMillimeters()
    case DoseAreaProductCentigraySquareMillimeter:
		return a.CentigraySquareMillimeters()
    case DoseAreaProductDecigraySquareMillimeter:
		return a.DecigraySquareMillimeters()
	default:
		return math.NaN()
	}
}

func (a *DoseAreaProduct) convertFromBase(toUnit DoseAreaProductUnits) float64 {
    value := a.value
	switch toUnit { 
	case DoseAreaProductGraySquareMeter:
		return (value) 
	case DoseAreaProductGraySquareDecimeter:
		return (value * 100) 
	case DoseAreaProductGraySquareCentimeter:
		return (value * 10000) 
	case DoseAreaProductGraySquareMillimeter:
		return (value * 1000000) 
	case DoseAreaProductMicrograySquareMeter:
		return ((value) / 1e-06) 
	case DoseAreaProductMilligraySquareMeter:
		return ((value) / 0.001) 
	case DoseAreaProductCentigraySquareMeter:
		return ((value) / 0.01) 
	case DoseAreaProductDecigraySquareMeter:
		return ((value) / 0.1) 
	case DoseAreaProductMicrograySquareDecimeter:
		return ((value * 100) / 1e-06) 
	case DoseAreaProductMilligraySquareDecimeter:
		return ((value * 100) / 0.001) 
	case DoseAreaProductCentigraySquareDecimeter:
		return ((value * 100) / 0.01) 
	case DoseAreaProductDecigraySquareDecimeter:
		return ((value * 100) / 0.1) 
	case DoseAreaProductMicrograySquareCentimeter:
		return ((value * 10000) / 1e-06) 
	case DoseAreaProductMilligraySquareCentimeter:
		return ((value * 10000) / 0.001) 
	case DoseAreaProductCentigraySquareCentimeter:
		return ((value * 10000) / 0.01) 
	case DoseAreaProductDecigraySquareCentimeter:
		return ((value * 10000) / 0.1) 
	case DoseAreaProductMicrograySquareMillimeter:
		return ((value * 1000000) / 1e-06) 
	case DoseAreaProductMilligraySquareMillimeter:
		return ((value * 1000000) / 0.001) 
	case DoseAreaProductCentigraySquareMillimeter:
		return ((value * 1000000) / 0.01) 
	case DoseAreaProductDecigraySquareMillimeter:
		return ((value * 1000000) / 0.1) 
	default:
		return math.NaN()
	}
}

func (a *DoseAreaProduct) convertToBase(value float64, fromUnit DoseAreaProductUnits) float64 {
	switch fromUnit { 
	case DoseAreaProductGraySquareMeter:
		return (value) 
	case DoseAreaProductGraySquareDecimeter:
		return (value / 100) 
	case DoseAreaProductGraySquareCentimeter:
		return (value / 10000) 
	case DoseAreaProductGraySquareMillimeter:
		return (value / 1000000) 
	case DoseAreaProductMicrograySquareMeter:
		return ((value) * 1e-06) 
	case DoseAreaProductMilligraySquareMeter:
		return ((value) * 0.001) 
	case DoseAreaProductCentigraySquareMeter:
		return ((value) * 0.01) 
	case DoseAreaProductDecigraySquareMeter:
		return ((value) * 0.1) 
	case DoseAreaProductMicrograySquareDecimeter:
		return ((value / 100) * 1e-06) 
	case DoseAreaProductMilligraySquareDecimeter:
		return ((value / 100) * 0.001) 
	case DoseAreaProductCentigraySquareDecimeter:
		return ((value / 100) * 0.01) 
	case DoseAreaProductDecigraySquareDecimeter:
		return ((value / 100) * 0.1) 
	case DoseAreaProductMicrograySquareCentimeter:
		return ((value / 10000) * 1e-06) 
	case DoseAreaProductMilligraySquareCentimeter:
		return ((value / 10000) * 0.001) 
	case DoseAreaProductCentigraySquareCentimeter:
		return ((value / 10000) * 0.01) 
	case DoseAreaProductDecigraySquareCentimeter:
		return ((value / 10000) * 0.1) 
	case DoseAreaProductMicrograySquareMillimeter:
		return ((value / 1000000) * 1e-06) 
	case DoseAreaProductMilligraySquareMillimeter:
		return ((value / 1000000) * 0.001) 
	case DoseAreaProductCentigraySquareMillimeter:
		return ((value / 1000000) * 0.01) 
	case DoseAreaProductDecigraySquareMillimeter:
		return ((value / 1000000) * 0.1) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the DoseAreaProduct in the default unit (GraySquareMeter),
// formatted to two decimal places.
func (a DoseAreaProduct) String() string {
	return a.ToString(DoseAreaProductGraySquareMeter, 2)
}

// ToString formats the DoseAreaProduct value as a string with the specified unit and fractional digits.
// It converts the DoseAreaProduct to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the DoseAreaProduct value will be converted (e.g., GraySquareMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the DoseAreaProduct with the unit abbreviation.
func (a *DoseAreaProduct) ToString(unit DoseAreaProductUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetDoseAreaProductAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetDoseAreaProductAbbreviation(unit))
}

// Equals checks if the given DoseAreaProduct is equal to the current DoseAreaProduct.
//
// Parameters:
//    other: The DoseAreaProduct to compare against.
//
// Returns:
//    bool: Returns true if both DoseAreaProduct are equal, false otherwise.
func (a *DoseAreaProduct) Equals(other *DoseAreaProduct) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current DoseAreaProduct with another DoseAreaProduct.
// It returns -1 if the current DoseAreaProduct is less than the other DoseAreaProduct, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The DoseAreaProduct to compare against.
//
// Returns:
//    int: -1 if the current DoseAreaProduct is less, 1 if greater, and 0 if equal.
func (a *DoseAreaProduct) CompareTo(other *DoseAreaProduct) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given DoseAreaProduct to the current DoseAreaProduct and returns the result.
// The result is a new DoseAreaProduct instance with the sum of the values.
//
// Parameters:
//    other: The DoseAreaProduct to add to the current DoseAreaProduct.
//
// Returns:
//    *DoseAreaProduct: A new DoseAreaProduct instance representing the sum of both DoseAreaProduct.
func (a *DoseAreaProduct) Add(other *DoseAreaProduct) *DoseAreaProduct {
	return &DoseAreaProduct{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given DoseAreaProduct from the current DoseAreaProduct and returns the result.
// The result is a new DoseAreaProduct instance with the difference of the values.
//
// Parameters:
//    other: The DoseAreaProduct to subtract from the current DoseAreaProduct.
//
// Returns:
//    *DoseAreaProduct: A new DoseAreaProduct instance representing the difference of both DoseAreaProduct.
func (a *DoseAreaProduct) Subtract(other *DoseAreaProduct) *DoseAreaProduct {
	return &DoseAreaProduct{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current DoseAreaProduct by the given DoseAreaProduct and returns the result.
// The result is a new DoseAreaProduct instance with the product of the values.
//
// Parameters:
//    other: The DoseAreaProduct to multiply with the current DoseAreaProduct.
//
// Returns:
//    *DoseAreaProduct: A new DoseAreaProduct instance representing the product of both DoseAreaProduct.
func (a *DoseAreaProduct) Multiply(other *DoseAreaProduct) *DoseAreaProduct {
	return &DoseAreaProduct{value: a.value * other.BaseValue()}
}

// Divide divides the current DoseAreaProduct by the given DoseAreaProduct and returns the result.
// The result is a new DoseAreaProduct instance with the quotient of the values.
//
// Parameters:
//    other: The DoseAreaProduct to divide the current DoseAreaProduct by.
//
// Returns:
//    *DoseAreaProduct: A new DoseAreaProduct instance representing the quotient of both DoseAreaProduct.
func (a *DoseAreaProduct) Divide(other *DoseAreaProduct) *DoseAreaProduct {
	return &DoseAreaProduct{value: a.value / other.BaseValue()}
}

// GetDoseAreaProductAbbreviation gets the unit abbreviation.
func GetDoseAreaProductAbbreviation(unit DoseAreaProductUnits) string {
	switch unit { 
	case DoseAreaProductGraySquareMeter:
		return "Gy·m²" 
	case DoseAreaProductGraySquareDecimeter:
		return "Gy·dm²" 
	case DoseAreaProductGraySquareCentimeter:
		return "Gy·cm²" 
	case DoseAreaProductGraySquareMillimeter:
		return "Gy·mm²" 
	case DoseAreaProductMicrograySquareMeter:
		return "μGy·m²" 
	case DoseAreaProductMilligraySquareMeter:
		return "mGy·m²" 
	case DoseAreaProductCentigraySquareMeter:
		return "cGy·m²" 
	case DoseAreaProductDecigraySquareMeter:
		return "dGy·m²" 
	case DoseAreaProductMicrograySquareDecimeter:
		return "μGy·dm²" 
	case DoseAreaProductMilligraySquareDecimeter:
		return "mGy·dm²" 
	case DoseAreaProductCentigraySquareDecimeter:
		return "cGy·dm²" 
	case DoseAreaProductDecigraySquareDecimeter:
		return "dGy·dm²" 
	case DoseAreaProductMicrograySquareCentimeter:
		return "μGy·cm²" 
	case DoseAreaProductMilligraySquareCentimeter:
		return "mGy·cm²" 
	case DoseAreaProductCentigraySquareCentimeter:
		return "cGy·cm²" 
	case DoseAreaProductDecigraySquareCentimeter:
		return "dGy·cm²" 
	case DoseAreaProductMicrograySquareMillimeter:
		return "μGy·mm²" 
	case DoseAreaProductMilligraySquareMillimeter:
		return "mGy·mm²" 
	case DoseAreaProductCentigraySquareMillimeter:
		return "cGy·mm²" 
	case DoseAreaProductDecigraySquareMillimeter:
		return "dGy·mm²" 
	default:
		return ""
	}
}