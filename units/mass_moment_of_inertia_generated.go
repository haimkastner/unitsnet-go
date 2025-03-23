// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// MassMomentOfInertiaUnits defines various units of MassMomentOfInertia.
type MassMomentOfInertiaUnits string

const (
    
        // 
        MassMomentOfInertiaGramSquareMeter MassMomentOfInertiaUnits = "GramSquareMeter"
        // 
        MassMomentOfInertiaGramSquareDecimeter MassMomentOfInertiaUnits = "GramSquareDecimeter"
        // 
        MassMomentOfInertiaGramSquareCentimeter MassMomentOfInertiaUnits = "GramSquareCentimeter"
        // 
        MassMomentOfInertiaGramSquareMillimeter MassMomentOfInertiaUnits = "GramSquareMillimeter"
        // 
        MassMomentOfInertiaTonneSquareMeter MassMomentOfInertiaUnits = "TonneSquareMeter"
        // 
        MassMomentOfInertiaTonneSquareDecimeter MassMomentOfInertiaUnits = "TonneSquareDecimeter"
        // 
        MassMomentOfInertiaTonneSquareCentimeter MassMomentOfInertiaUnits = "TonneSquareCentimeter"
        // 
        MassMomentOfInertiaTonneSquareMilimeter MassMomentOfInertiaUnits = "TonneSquareMilimeter"
        // 
        MassMomentOfInertiaPoundSquareFoot MassMomentOfInertiaUnits = "PoundSquareFoot"
        // 
        MassMomentOfInertiaPoundSquareInch MassMomentOfInertiaUnits = "PoundSquareInch"
        // 
        MassMomentOfInertiaSlugSquareFoot MassMomentOfInertiaUnits = "SlugSquareFoot"
        // 
        MassMomentOfInertiaSlugSquareInch MassMomentOfInertiaUnits = "SlugSquareInch"
        // 
        MassMomentOfInertiaMilligramSquareMeter MassMomentOfInertiaUnits = "MilligramSquareMeter"
        // 
        MassMomentOfInertiaKilogramSquareMeter MassMomentOfInertiaUnits = "KilogramSquareMeter"
        // 
        MassMomentOfInertiaMilligramSquareDecimeter MassMomentOfInertiaUnits = "MilligramSquareDecimeter"
        // 
        MassMomentOfInertiaKilogramSquareDecimeter MassMomentOfInertiaUnits = "KilogramSquareDecimeter"
        // 
        MassMomentOfInertiaMilligramSquareCentimeter MassMomentOfInertiaUnits = "MilligramSquareCentimeter"
        // 
        MassMomentOfInertiaKilogramSquareCentimeter MassMomentOfInertiaUnits = "KilogramSquareCentimeter"
        // 
        MassMomentOfInertiaMilligramSquareMillimeter MassMomentOfInertiaUnits = "MilligramSquareMillimeter"
        // 
        MassMomentOfInertiaKilogramSquareMillimeter MassMomentOfInertiaUnits = "KilogramSquareMillimeter"
        // 
        MassMomentOfInertiaKilotonneSquareMeter MassMomentOfInertiaUnits = "KilotonneSquareMeter"
        // 
        MassMomentOfInertiaMegatonneSquareMeter MassMomentOfInertiaUnits = "MegatonneSquareMeter"
        // 
        MassMomentOfInertiaKilotonneSquareDecimeter MassMomentOfInertiaUnits = "KilotonneSquareDecimeter"
        // 
        MassMomentOfInertiaMegatonneSquareDecimeter MassMomentOfInertiaUnits = "MegatonneSquareDecimeter"
        // 
        MassMomentOfInertiaKilotonneSquareCentimeter MassMomentOfInertiaUnits = "KilotonneSquareCentimeter"
        // 
        MassMomentOfInertiaMegatonneSquareCentimeter MassMomentOfInertiaUnits = "MegatonneSquareCentimeter"
        // 
        MassMomentOfInertiaKilotonneSquareMilimeter MassMomentOfInertiaUnits = "KilotonneSquareMilimeter"
        // 
        MassMomentOfInertiaMegatonneSquareMilimeter MassMomentOfInertiaUnits = "MegatonneSquareMilimeter"
)

var internalMassMomentOfInertiaUnitsMap = map[MassMomentOfInertiaUnits]bool{
	
	MassMomentOfInertiaGramSquareMeter: true,
	MassMomentOfInertiaGramSquareDecimeter: true,
	MassMomentOfInertiaGramSquareCentimeter: true,
	MassMomentOfInertiaGramSquareMillimeter: true,
	MassMomentOfInertiaTonneSquareMeter: true,
	MassMomentOfInertiaTonneSquareDecimeter: true,
	MassMomentOfInertiaTonneSquareCentimeter: true,
	MassMomentOfInertiaTonneSquareMilimeter: true,
	MassMomentOfInertiaPoundSquareFoot: true,
	MassMomentOfInertiaPoundSquareInch: true,
	MassMomentOfInertiaSlugSquareFoot: true,
	MassMomentOfInertiaSlugSquareInch: true,
	MassMomentOfInertiaMilligramSquareMeter: true,
	MassMomentOfInertiaKilogramSquareMeter: true,
	MassMomentOfInertiaMilligramSquareDecimeter: true,
	MassMomentOfInertiaKilogramSquareDecimeter: true,
	MassMomentOfInertiaMilligramSquareCentimeter: true,
	MassMomentOfInertiaKilogramSquareCentimeter: true,
	MassMomentOfInertiaMilligramSquareMillimeter: true,
	MassMomentOfInertiaKilogramSquareMillimeter: true,
	MassMomentOfInertiaKilotonneSquareMeter: true,
	MassMomentOfInertiaMegatonneSquareMeter: true,
	MassMomentOfInertiaKilotonneSquareDecimeter: true,
	MassMomentOfInertiaMegatonneSquareDecimeter: true,
	MassMomentOfInertiaKilotonneSquareCentimeter: true,
	MassMomentOfInertiaMegatonneSquareCentimeter: true,
	MassMomentOfInertiaKilotonneSquareMilimeter: true,
	MassMomentOfInertiaMegatonneSquareMilimeter: true,
}

// MassMomentOfInertiaDto represents a MassMomentOfInertia measurement with a numerical value and its corresponding unit.
type MassMomentOfInertiaDto struct {
    // Value is the numerical representation of the MassMomentOfInertia.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the MassMomentOfInertia, as defined in the MassMomentOfInertiaUnits enumeration.
	Unit  MassMomentOfInertiaUnits `json:"unit" validate:"required,oneof=GramSquareMeter GramSquareDecimeter GramSquareCentimeter GramSquareMillimeter TonneSquareMeter TonneSquareDecimeter TonneSquareCentimeter TonneSquareMilimeter PoundSquareFoot PoundSquareInch SlugSquareFoot SlugSquareInch MilligramSquareMeter KilogramSquareMeter MilligramSquareDecimeter KilogramSquareDecimeter MilligramSquareCentimeter KilogramSquareCentimeter MilligramSquareMillimeter KilogramSquareMillimeter KilotonneSquareMeter MegatonneSquareMeter KilotonneSquareDecimeter MegatonneSquareDecimeter KilotonneSquareCentimeter MegatonneSquareCentimeter KilotonneSquareMilimeter MegatonneSquareMilimeter"`
}

// MassMomentOfInertiaDtoFactory groups methods for creating and serializing MassMomentOfInertiaDto objects.
type MassMomentOfInertiaDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a MassMomentOfInertiaDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf MassMomentOfInertiaDtoFactory) FromJSON(data []byte) (*MassMomentOfInertiaDto, error) {
	a := MassMomentOfInertiaDto{}

    // Parse JSON into MassMomentOfInertiaDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a MassMomentOfInertiaDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a MassMomentOfInertiaDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// MassMomentOfInertia represents a measurement in a of MassMomentOfInertia.
//
// A property of body reflects how its mass is distributed with regard to an axis.
type MassMomentOfInertia struct {
	// value is the base measurement stored internally.
	value       float64
    
    gram_square_metersLazy *float64 
    gram_square_decimetersLazy *float64 
    gram_square_centimetersLazy *float64 
    gram_square_millimetersLazy *float64 
    tonne_square_metersLazy *float64 
    tonne_square_decimetersLazy *float64 
    tonne_square_centimetersLazy *float64 
    tonne_square_milimetersLazy *float64 
    pound_square_feetLazy *float64 
    pound_square_inchesLazy *float64 
    slug_square_feetLazy *float64 
    slug_square_inchesLazy *float64 
    milligram_square_metersLazy *float64 
    kilogram_square_metersLazy *float64 
    milligram_square_decimetersLazy *float64 
    kilogram_square_decimetersLazy *float64 
    milligram_square_centimetersLazy *float64 
    kilogram_square_centimetersLazy *float64 
    milligram_square_millimetersLazy *float64 
    kilogram_square_millimetersLazy *float64 
    kilotonne_square_metersLazy *float64 
    megatonne_square_metersLazy *float64 
    kilotonne_square_decimetersLazy *float64 
    megatonne_square_decimetersLazy *float64 
    kilotonne_square_centimetersLazy *float64 
    megatonne_square_centimetersLazy *float64 
    kilotonne_square_milimetersLazy *float64 
    megatonne_square_milimetersLazy *float64 
}

// MassMomentOfInertiaFactory groups methods for creating MassMomentOfInertia instances.
type MassMomentOfInertiaFactory struct{}

// CreateMassMomentOfInertia creates a new MassMomentOfInertia instance from the given value and unit.
func (uf MassMomentOfInertiaFactory) CreateMassMomentOfInertia(value float64, unit MassMomentOfInertiaUnits) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, unit)
}

// FromDto converts a MassMomentOfInertiaDto to a MassMomentOfInertia instance.
func (uf MassMomentOfInertiaFactory) FromDto(dto MassMomentOfInertiaDto) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a MassMomentOfInertia instance.
func (uf MassMomentOfInertiaFactory) FromDtoJSON(data []byte) (*MassMomentOfInertia, error) {
	unitDto, err := MassMomentOfInertiaDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse MassMomentOfInertiaDto from JSON: %w", err)
	}
	return MassMomentOfInertiaFactory{}.FromDto(*unitDto)
}


// FromGramSquareMeters creates a new MassMomentOfInertia instance from a value in GramSquareMeters.
func (uf MassMomentOfInertiaFactory) FromGramSquareMeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaGramSquareMeter)
}

// FromGramSquareDecimeters creates a new MassMomentOfInertia instance from a value in GramSquareDecimeters.
func (uf MassMomentOfInertiaFactory) FromGramSquareDecimeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaGramSquareDecimeter)
}

// FromGramSquareCentimeters creates a new MassMomentOfInertia instance from a value in GramSquareCentimeters.
func (uf MassMomentOfInertiaFactory) FromGramSquareCentimeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaGramSquareCentimeter)
}

// FromGramSquareMillimeters creates a new MassMomentOfInertia instance from a value in GramSquareMillimeters.
func (uf MassMomentOfInertiaFactory) FromGramSquareMillimeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaGramSquareMillimeter)
}

// FromTonneSquareMeters creates a new MassMomentOfInertia instance from a value in TonneSquareMeters.
func (uf MassMomentOfInertiaFactory) FromTonneSquareMeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaTonneSquareMeter)
}

// FromTonneSquareDecimeters creates a new MassMomentOfInertia instance from a value in TonneSquareDecimeters.
func (uf MassMomentOfInertiaFactory) FromTonneSquareDecimeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaTonneSquareDecimeter)
}

// FromTonneSquareCentimeters creates a new MassMomentOfInertia instance from a value in TonneSquareCentimeters.
func (uf MassMomentOfInertiaFactory) FromTonneSquareCentimeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaTonneSquareCentimeter)
}

// FromTonneSquareMilimeters creates a new MassMomentOfInertia instance from a value in TonneSquareMilimeters.
func (uf MassMomentOfInertiaFactory) FromTonneSquareMilimeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaTonneSquareMilimeter)
}

// FromPoundSquareFeet creates a new MassMomentOfInertia instance from a value in PoundSquareFeet.
func (uf MassMomentOfInertiaFactory) FromPoundSquareFeet(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaPoundSquareFoot)
}

// FromPoundSquareInches creates a new MassMomentOfInertia instance from a value in PoundSquareInches.
func (uf MassMomentOfInertiaFactory) FromPoundSquareInches(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaPoundSquareInch)
}

// FromSlugSquareFeet creates a new MassMomentOfInertia instance from a value in SlugSquareFeet.
func (uf MassMomentOfInertiaFactory) FromSlugSquareFeet(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaSlugSquareFoot)
}

// FromSlugSquareInches creates a new MassMomentOfInertia instance from a value in SlugSquareInches.
func (uf MassMomentOfInertiaFactory) FromSlugSquareInches(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaSlugSquareInch)
}

// FromMilligramSquareMeters creates a new MassMomentOfInertia instance from a value in MilligramSquareMeters.
func (uf MassMomentOfInertiaFactory) FromMilligramSquareMeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaMilligramSquareMeter)
}

// FromKilogramSquareMeters creates a new MassMomentOfInertia instance from a value in KilogramSquareMeters.
func (uf MassMomentOfInertiaFactory) FromKilogramSquareMeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaKilogramSquareMeter)
}

// FromMilligramSquareDecimeters creates a new MassMomentOfInertia instance from a value in MilligramSquareDecimeters.
func (uf MassMomentOfInertiaFactory) FromMilligramSquareDecimeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaMilligramSquareDecimeter)
}

// FromKilogramSquareDecimeters creates a new MassMomentOfInertia instance from a value in KilogramSquareDecimeters.
func (uf MassMomentOfInertiaFactory) FromKilogramSquareDecimeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaKilogramSquareDecimeter)
}

// FromMilligramSquareCentimeters creates a new MassMomentOfInertia instance from a value in MilligramSquareCentimeters.
func (uf MassMomentOfInertiaFactory) FromMilligramSquareCentimeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaMilligramSquareCentimeter)
}

// FromKilogramSquareCentimeters creates a new MassMomentOfInertia instance from a value in KilogramSquareCentimeters.
func (uf MassMomentOfInertiaFactory) FromKilogramSquareCentimeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaKilogramSquareCentimeter)
}

// FromMilligramSquareMillimeters creates a new MassMomentOfInertia instance from a value in MilligramSquareMillimeters.
func (uf MassMomentOfInertiaFactory) FromMilligramSquareMillimeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaMilligramSquareMillimeter)
}

// FromKilogramSquareMillimeters creates a new MassMomentOfInertia instance from a value in KilogramSquareMillimeters.
func (uf MassMomentOfInertiaFactory) FromKilogramSquareMillimeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaKilogramSquareMillimeter)
}

// FromKilotonneSquareMeters creates a new MassMomentOfInertia instance from a value in KilotonneSquareMeters.
func (uf MassMomentOfInertiaFactory) FromKilotonneSquareMeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaKilotonneSquareMeter)
}

// FromMegatonneSquareMeters creates a new MassMomentOfInertia instance from a value in MegatonneSquareMeters.
func (uf MassMomentOfInertiaFactory) FromMegatonneSquareMeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaMegatonneSquareMeter)
}

// FromKilotonneSquareDecimeters creates a new MassMomentOfInertia instance from a value in KilotonneSquareDecimeters.
func (uf MassMomentOfInertiaFactory) FromKilotonneSquareDecimeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaKilotonneSquareDecimeter)
}

// FromMegatonneSquareDecimeters creates a new MassMomentOfInertia instance from a value in MegatonneSquareDecimeters.
func (uf MassMomentOfInertiaFactory) FromMegatonneSquareDecimeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaMegatonneSquareDecimeter)
}

// FromKilotonneSquareCentimeters creates a new MassMomentOfInertia instance from a value in KilotonneSquareCentimeters.
func (uf MassMomentOfInertiaFactory) FromKilotonneSquareCentimeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaKilotonneSquareCentimeter)
}

// FromMegatonneSquareCentimeters creates a new MassMomentOfInertia instance from a value in MegatonneSquareCentimeters.
func (uf MassMomentOfInertiaFactory) FromMegatonneSquareCentimeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaMegatonneSquareCentimeter)
}

// FromKilotonneSquareMilimeters creates a new MassMomentOfInertia instance from a value in KilotonneSquareMilimeters.
func (uf MassMomentOfInertiaFactory) FromKilotonneSquareMilimeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaKilotonneSquareMilimeter)
}

// FromMegatonneSquareMilimeters creates a new MassMomentOfInertia instance from a value in MegatonneSquareMilimeters.
func (uf MassMomentOfInertiaFactory) FromMegatonneSquareMilimeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaMegatonneSquareMilimeter)
}


// newMassMomentOfInertia creates a new MassMomentOfInertia.
func newMassMomentOfInertia(value float64, fromUnit MassMomentOfInertiaUnits) (*MassMomentOfInertia, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalMassMomentOfInertiaUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in MassMomentOfInertiaUnits", fromUnit)
	}
	a := &MassMomentOfInertia{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of MassMomentOfInertia in KilogramSquareMeter unit (the base/default quantity).
func (a *MassMomentOfInertia) BaseValue() float64 {
	return a.value
}


// GramSquareMeters returns the MassMomentOfInertia value in GramSquareMeters.
//
// 
func (a *MassMomentOfInertia) GramSquareMeters() float64 {
	if a.gram_square_metersLazy != nil {
		return *a.gram_square_metersLazy
	}
	gram_square_meters := a.convertFromBase(MassMomentOfInertiaGramSquareMeter)
	a.gram_square_metersLazy = &gram_square_meters
	return gram_square_meters
}

// GramSquareDecimeters returns the MassMomentOfInertia value in GramSquareDecimeters.
//
// 
func (a *MassMomentOfInertia) GramSquareDecimeters() float64 {
	if a.gram_square_decimetersLazy != nil {
		return *a.gram_square_decimetersLazy
	}
	gram_square_decimeters := a.convertFromBase(MassMomentOfInertiaGramSquareDecimeter)
	a.gram_square_decimetersLazy = &gram_square_decimeters
	return gram_square_decimeters
}

// GramSquareCentimeters returns the MassMomentOfInertia value in GramSquareCentimeters.
//
// 
func (a *MassMomentOfInertia) GramSquareCentimeters() float64 {
	if a.gram_square_centimetersLazy != nil {
		return *a.gram_square_centimetersLazy
	}
	gram_square_centimeters := a.convertFromBase(MassMomentOfInertiaGramSquareCentimeter)
	a.gram_square_centimetersLazy = &gram_square_centimeters
	return gram_square_centimeters
}

// GramSquareMillimeters returns the MassMomentOfInertia value in GramSquareMillimeters.
//
// 
func (a *MassMomentOfInertia) GramSquareMillimeters() float64 {
	if a.gram_square_millimetersLazy != nil {
		return *a.gram_square_millimetersLazy
	}
	gram_square_millimeters := a.convertFromBase(MassMomentOfInertiaGramSquareMillimeter)
	a.gram_square_millimetersLazy = &gram_square_millimeters
	return gram_square_millimeters
}

// TonneSquareMeters returns the MassMomentOfInertia value in TonneSquareMeters.
//
// 
func (a *MassMomentOfInertia) TonneSquareMeters() float64 {
	if a.tonne_square_metersLazy != nil {
		return *a.tonne_square_metersLazy
	}
	tonne_square_meters := a.convertFromBase(MassMomentOfInertiaTonneSquareMeter)
	a.tonne_square_metersLazy = &tonne_square_meters
	return tonne_square_meters
}

// TonneSquareDecimeters returns the MassMomentOfInertia value in TonneSquareDecimeters.
//
// 
func (a *MassMomentOfInertia) TonneSquareDecimeters() float64 {
	if a.tonne_square_decimetersLazy != nil {
		return *a.tonne_square_decimetersLazy
	}
	tonne_square_decimeters := a.convertFromBase(MassMomentOfInertiaTonneSquareDecimeter)
	a.tonne_square_decimetersLazy = &tonne_square_decimeters
	return tonne_square_decimeters
}

// TonneSquareCentimeters returns the MassMomentOfInertia value in TonneSquareCentimeters.
//
// 
func (a *MassMomentOfInertia) TonneSquareCentimeters() float64 {
	if a.tonne_square_centimetersLazy != nil {
		return *a.tonne_square_centimetersLazy
	}
	tonne_square_centimeters := a.convertFromBase(MassMomentOfInertiaTonneSquareCentimeter)
	a.tonne_square_centimetersLazy = &tonne_square_centimeters
	return tonne_square_centimeters
}

// TonneSquareMilimeters returns the MassMomentOfInertia value in TonneSquareMilimeters.
//
// 
func (a *MassMomentOfInertia) TonneSquareMilimeters() float64 {
	if a.tonne_square_milimetersLazy != nil {
		return *a.tonne_square_milimetersLazy
	}
	tonne_square_milimeters := a.convertFromBase(MassMomentOfInertiaTonneSquareMilimeter)
	a.tonne_square_milimetersLazy = &tonne_square_milimeters
	return tonne_square_milimeters
}

// PoundSquareFeet returns the MassMomentOfInertia value in PoundSquareFeet.
//
// 
func (a *MassMomentOfInertia) PoundSquareFeet() float64 {
	if a.pound_square_feetLazy != nil {
		return *a.pound_square_feetLazy
	}
	pound_square_feet := a.convertFromBase(MassMomentOfInertiaPoundSquareFoot)
	a.pound_square_feetLazy = &pound_square_feet
	return pound_square_feet
}

// PoundSquareInches returns the MassMomentOfInertia value in PoundSquareInches.
//
// 
func (a *MassMomentOfInertia) PoundSquareInches() float64 {
	if a.pound_square_inchesLazy != nil {
		return *a.pound_square_inchesLazy
	}
	pound_square_inches := a.convertFromBase(MassMomentOfInertiaPoundSquareInch)
	a.pound_square_inchesLazy = &pound_square_inches
	return pound_square_inches
}

// SlugSquareFeet returns the MassMomentOfInertia value in SlugSquareFeet.
//
// 
func (a *MassMomentOfInertia) SlugSquareFeet() float64 {
	if a.slug_square_feetLazy != nil {
		return *a.slug_square_feetLazy
	}
	slug_square_feet := a.convertFromBase(MassMomentOfInertiaSlugSquareFoot)
	a.slug_square_feetLazy = &slug_square_feet
	return slug_square_feet
}

// SlugSquareInches returns the MassMomentOfInertia value in SlugSquareInches.
//
// 
func (a *MassMomentOfInertia) SlugSquareInches() float64 {
	if a.slug_square_inchesLazy != nil {
		return *a.slug_square_inchesLazy
	}
	slug_square_inches := a.convertFromBase(MassMomentOfInertiaSlugSquareInch)
	a.slug_square_inchesLazy = &slug_square_inches
	return slug_square_inches
}

// MilligramSquareMeters returns the MassMomentOfInertia value in MilligramSquareMeters.
//
// 
func (a *MassMomentOfInertia) MilligramSquareMeters() float64 {
	if a.milligram_square_metersLazy != nil {
		return *a.milligram_square_metersLazy
	}
	milligram_square_meters := a.convertFromBase(MassMomentOfInertiaMilligramSquareMeter)
	a.milligram_square_metersLazy = &milligram_square_meters
	return milligram_square_meters
}

// KilogramSquareMeters returns the MassMomentOfInertia value in KilogramSquareMeters.
//
// 
func (a *MassMomentOfInertia) KilogramSquareMeters() float64 {
	if a.kilogram_square_metersLazy != nil {
		return *a.kilogram_square_metersLazy
	}
	kilogram_square_meters := a.convertFromBase(MassMomentOfInertiaKilogramSquareMeter)
	a.kilogram_square_metersLazy = &kilogram_square_meters
	return kilogram_square_meters
}

// MilligramSquareDecimeters returns the MassMomentOfInertia value in MilligramSquareDecimeters.
//
// 
func (a *MassMomentOfInertia) MilligramSquareDecimeters() float64 {
	if a.milligram_square_decimetersLazy != nil {
		return *a.milligram_square_decimetersLazy
	}
	milligram_square_decimeters := a.convertFromBase(MassMomentOfInertiaMilligramSquareDecimeter)
	a.milligram_square_decimetersLazy = &milligram_square_decimeters
	return milligram_square_decimeters
}

// KilogramSquareDecimeters returns the MassMomentOfInertia value in KilogramSquareDecimeters.
//
// 
func (a *MassMomentOfInertia) KilogramSquareDecimeters() float64 {
	if a.kilogram_square_decimetersLazy != nil {
		return *a.kilogram_square_decimetersLazy
	}
	kilogram_square_decimeters := a.convertFromBase(MassMomentOfInertiaKilogramSquareDecimeter)
	a.kilogram_square_decimetersLazy = &kilogram_square_decimeters
	return kilogram_square_decimeters
}

// MilligramSquareCentimeters returns the MassMomentOfInertia value in MilligramSquareCentimeters.
//
// 
func (a *MassMomentOfInertia) MilligramSquareCentimeters() float64 {
	if a.milligram_square_centimetersLazy != nil {
		return *a.milligram_square_centimetersLazy
	}
	milligram_square_centimeters := a.convertFromBase(MassMomentOfInertiaMilligramSquareCentimeter)
	a.milligram_square_centimetersLazy = &milligram_square_centimeters
	return milligram_square_centimeters
}

// KilogramSquareCentimeters returns the MassMomentOfInertia value in KilogramSquareCentimeters.
//
// 
func (a *MassMomentOfInertia) KilogramSquareCentimeters() float64 {
	if a.kilogram_square_centimetersLazy != nil {
		return *a.kilogram_square_centimetersLazy
	}
	kilogram_square_centimeters := a.convertFromBase(MassMomentOfInertiaKilogramSquareCentimeter)
	a.kilogram_square_centimetersLazy = &kilogram_square_centimeters
	return kilogram_square_centimeters
}

// MilligramSquareMillimeters returns the MassMomentOfInertia value in MilligramSquareMillimeters.
//
// 
func (a *MassMomentOfInertia) MilligramSquareMillimeters() float64 {
	if a.milligram_square_millimetersLazy != nil {
		return *a.milligram_square_millimetersLazy
	}
	milligram_square_millimeters := a.convertFromBase(MassMomentOfInertiaMilligramSquareMillimeter)
	a.milligram_square_millimetersLazy = &milligram_square_millimeters
	return milligram_square_millimeters
}

// KilogramSquareMillimeters returns the MassMomentOfInertia value in KilogramSquareMillimeters.
//
// 
func (a *MassMomentOfInertia) KilogramSquareMillimeters() float64 {
	if a.kilogram_square_millimetersLazy != nil {
		return *a.kilogram_square_millimetersLazy
	}
	kilogram_square_millimeters := a.convertFromBase(MassMomentOfInertiaKilogramSquareMillimeter)
	a.kilogram_square_millimetersLazy = &kilogram_square_millimeters
	return kilogram_square_millimeters
}

// KilotonneSquareMeters returns the MassMomentOfInertia value in KilotonneSquareMeters.
//
// 
func (a *MassMomentOfInertia) KilotonneSquareMeters() float64 {
	if a.kilotonne_square_metersLazy != nil {
		return *a.kilotonne_square_metersLazy
	}
	kilotonne_square_meters := a.convertFromBase(MassMomentOfInertiaKilotonneSquareMeter)
	a.kilotonne_square_metersLazy = &kilotonne_square_meters
	return kilotonne_square_meters
}

// MegatonneSquareMeters returns the MassMomentOfInertia value in MegatonneSquareMeters.
//
// 
func (a *MassMomentOfInertia) MegatonneSquareMeters() float64 {
	if a.megatonne_square_metersLazy != nil {
		return *a.megatonne_square_metersLazy
	}
	megatonne_square_meters := a.convertFromBase(MassMomentOfInertiaMegatonneSquareMeter)
	a.megatonne_square_metersLazy = &megatonne_square_meters
	return megatonne_square_meters
}

// KilotonneSquareDecimeters returns the MassMomentOfInertia value in KilotonneSquareDecimeters.
//
// 
func (a *MassMomentOfInertia) KilotonneSquareDecimeters() float64 {
	if a.kilotonne_square_decimetersLazy != nil {
		return *a.kilotonne_square_decimetersLazy
	}
	kilotonne_square_decimeters := a.convertFromBase(MassMomentOfInertiaKilotonneSquareDecimeter)
	a.kilotonne_square_decimetersLazy = &kilotonne_square_decimeters
	return kilotonne_square_decimeters
}

// MegatonneSquareDecimeters returns the MassMomentOfInertia value in MegatonneSquareDecimeters.
//
// 
func (a *MassMomentOfInertia) MegatonneSquareDecimeters() float64 {
	if a.megatonne_square_decimetersLazy != nil {
		return *a.megatonne_square_decimetersLazy
	}
	megatonne_square_decimeters := a.convertFromBase(MassMomentOfInertiaMegatonneSquareDecimeter)
	a.megatonne_square_decimetersLazy = &megatonne_square_decimeters
	return megatonne_square_decimeters
}

// KilotonneSquareCentimeters returns the MassMomentOfInertia value in KilotonneSquareCentimeters.
//
// 
func (a *MassMomentOfInertia) KilotonneSquareCentimeters() float64 {
	if a.kilotonne_square_centimetersLazy != nil {
		return *a.kilotonne_square_centimetersLazy
	}
	kilotonne_square_centimeters := a.convertFromBase(MassMomentOfInertiaKilotonneSquareCentimeter)
	a.kilotonne_square_centimetersLazy = &kilotonne_square_centimeters
	return kilotonne_square_centimeters
}

// MegatonneSquareCentimeters returns the MassMomentOfInertia value in MegatonneSquareCentimeters.
//
// 
func (a *MassMomentOfInertia) MegatonneSquareCentimeters() float64 {
	if a.megatonne_square_centimetersLazy != nil {
		return *a.megatonne_square_centimetersLazy
	}
	megatonne_square_centimeters := a.convertFromBase(MassMomentOfInertiaMegatonneSquareCentimeter)
	a.megatonne_square_centimetersLazy = &megatonne_square_centimeters
	return megatonne_square_centimeters
}

// KilotonneSquareMilimeters returns the MassMomentOfInertia value in KilotonneSquareMilimeters.
//
// 
func (a *MassMomentOfInertia) KilotonneSquareMilimeters() float64 {
	if a.kilotonne_square_milimetersLazy != nil {
		return *a.kilotonne_square_milimetersLazy
	}
	kilotonne_square_milimeters := a.convertFromBase(MassMomentOfInertiaKilotonneSquareMilimeter)
	a.kilotonne_square_milimetersLazy = &kilotonne_square_milimeters
	return kilotonne_square_milimeters
}

// MegatonneSquareMilimeters returns the MassMomentOfInertia value in MegatonneSquareMilimeters.
//
// 
func (a *MassMomentOfInertia) MegatonneSquareMilimeters() float64 {
	if a.megatonne_square_milimetersLazy != nil {
		return *a.megatonne_square_milimetersLazy
	}
	megatonne_square_milimeters := a.convertFromBase(MassMomentOfInertiaMegatonneSquareMilimeter)
	a.megatonne_square_milimetersLazy = &megatonne_square_milimeters
	return megatonne_square_milimeters
}


// ToDto creates a MassMomentOfInertiaDto representation from the MassMomentOfInertia instance.
//
// If the provided holdInUnit is nil, the value will be repesented by KilogramSquareMeter by default.
func (a *MassMomentOfInertia) ToDto(holdInUnit *MassMomentOfInertiaUnits) MassMomentOfInertiaDto {
	if holdInUnit == nil {
		defaultUnit := MassMomentOfInertiaKilogramSquareMeter // Default value
		holdInUnit = &defaultUnit
	}

	return MassMomentOfInertiaDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the MassMomentOfInertia instance.
//
// If the provided holdInUnit is nil, the value will be repesented by KilogramSquareMeter by default.
func (a *MassMomentOfInertia) ToDtoJSON(holdInUnit *MassMomentOfInertiaUnits) ([]byte, error) {
	// Convert to MassMomentOfInertiaDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a MassMomentOfInertia to a specific unit value.
// The function uses the provided unit type (MassMomentOfInertiaUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *MassMomentOfInertia) Convert(toUnit MassMomentOfInertiaUnits) float64 {
	switch toUnit { 
    case MassMomentOfInertiaGramSquareMeter:
		return a.GramSquareMeters()
    case MassMomentOfInertiaGramSquareDecimeter:
		return a.GramSquareDecimeters()
    case MassMomentOfInertiaGramSquareCentimeter:
		return a.GramSquareCentimeters()
    case MassMomentOfInertiaGramSquareMillimeter:
		return a.GramSquareMillimeters()
    case MassMomentOfInertiaTonneSquareMeter:
		return a.TonneSquareMeters()
    case MassMomentOfInertiaTonneSquareDecimeter:
		return a.TonneSquareDecimeters()
    case MassMomentOfInertiaTonneSquareCentimeter:
		return a.TonneSquareCentimeters()
    case MassMomentOfInertiaTonneSquareMilimeter:
		return a.TonneSquareMilimeters()
    case MassMomentOfInertiaPoundSquareFoot:
		return a.PoundSquareFeet()
    case MassMomentOfInertiaPoundSquareInch:
		return a.PoundSquareInches()
    case MassMomentOfInertiaSlugSquareFoot:
		return a.SlugSquareFeet()
    case MassMomentOfInertiaSlugSquareInch:
		return a.SlugSquareInches()
    case MassMomentOfInertiaMilligramSquareMeter:
		return a.MilligramSquareMeters()
    case MassMomentOfInertiaKilogramSquareMeter:
		return a.KilogramSquareMeters()
    case MassMomentOfInertiaMilligramSquareDecimeter:
		return a.MilligramSquareDecimeters()
    case MassMomentOfInertiaKilogramSquareDecimeter:
		return a.KilogramSquareDecimeters()
    case MassMomentOfInertiaMilligramSquareCentimeter:
		return a.MilligramSquareCentimeters()
    case MassMomentOfInertiaKilogramSquareCentimeter:
		return a.KilogramSquareCentimeters()
    case MassMomentOfInertiaMilligramSquareMillimeter:
		return a.MilligramSquareMillimeters()
    case MassMomentOfInertiaKilogramSquareMillimeter:
		return a.KilogramSquareMillimeters()
    case MassMomentOfInertiaKilotonneSquareMeter:
		return a.KilotonneSquareMeters()
    case MassMomentOfInertiaMegatonneSquareMeter:
		return a.MegatonneSquareMeters()
    case MassMomentOfInertiaKilotonneSquareDecimeter:
		return a.KilotonneSquareDecimeters()
    case MassMomentOfInertiaMegatonneSquareDecimeter:
		return a.MegatonneSquareDecimeters()
    case MassMomentOfInertiaKilotonneSquareCentimeter:
		return a.KilotonneSquareCentimeters()
    case MassMomentOfInertiaMegatonneSquareCentimeter:
		return a.MegatonneSquareCentimeters()
    case MassMomentOfInertiaKilotonneSquareMilimeter:
		return a.KilotonneSquareMilimeters()
    case MassMomentOfInertiaMegatonneSquareMilimeter:
		return a.MegatonneSquareMilimeters()
	default:
		return math.NaN()
	}
}

func (a *MassMomentOfInertia) convertFromBase(toUnit MassMomentOfInertiaUnits) float64 {
    value := a.value
	switch toUnit { 
	case MassMomentOfInertiaGramSquareMeter:
		return (value * 1e3) 
	case MassMomentOfInertiaGramSquareDecimeter:
		return (value * 1e5) 
	case MassMomentOfInertiaGramSquareCentimeter:
		return (value * 1e7) 
	case MassMomentOfInertiaGramSquareMillimeter:
		return (value * 1e9) 
	case MassMomentOfInertiaTonneSquareMeter:
		return (value * 1e-3) 
	case MassMomentOfInertiaTonneSquareDecimeter:
		return (value * 1e-1) 
	case MassMomentOfInertiaTonneSquareCentimeter:
		return (value * 1e1) 
	case MassMomentOfInertiaTonneSquareMilimeter:
		return (value * 1e3) 
	case MassMomentOfInertiaPoundSquareFoot:
		return (value / 4.21401101e-2) 
	case MassMomentOfInertiaPoundSquareInch:
		return (value / 2.9263965e-4) 
	case MassMomentOfInertiaSlugSquareFoot:
		return (value / 1.3558179619) 
	case MassMomentOfInertiaSlugSquareInch:
		return (value / 9.41540242e-3) 
	case MassMomentOfInertiaMilligramSquareMeter:
		return ((value * 1e3) / 0.001) 
	case MassMomentOfInertiaKilogramSquareMeter:
		return ((value * 1e3) / 1000.0) 
	case MassMomentOfInertiaMilligramSquareDecimeter:
		return ((value * 1e5) / 0.001) 
	case MassMomentOfInertiaKilogramSquareDecimeter:
		return ((value * 1e5) / 1000.0) 
	case MassMomentOfInertiaMilligramSquareCentimeter:
		return ((value * 1e7) / 0.001) 
	case MassMomentOfInertiaKilogramSquareCentimeter:
		return ((value * 1e7) / 1000.0) 
	case MassMomentOfInertiaMilligramSquareMillimeter:
		return ((value * 1e9) / 0.001) 
	case MassMomentOfInertiaKilogramSquareMillimeter:
		return ((value * 1e9) / 1000.0) 
	case MassMomentOfInertiaKilotonneSquareMeter:
		return ((value * 1e-3) / 1000.0) 
	case MassMomentOfInertiaMegatonneSquareMeter:
		return ((value * 1e-3) / 1000000.0) 
	case MassMomentOfInertiaKilotonneSquareDecimeter:
		return ((value * 1e-1) / 1000.0) 
	case MassMomentOfInertiaMegatonneSquareDecimeter:
		return ((value * 1e-1) / 1000000.0) 
	case MassMomentOfInertiaKilotonneSquareCentimeter:
		return ((value * 1e1) / 1000.0) 
	case MassMomentOfInertiaMegatonneSquareCentimeter:
		return ((value * 1e1) / 1000000.0) 
	case MassMomentOfInertiaKilotonneSquareMilimeter:
		return ((value * 1e3) / 1000.0) 
	case MassMomentOfInertiaMegatonneSquareMilimeter:
		return ((value * 1e3) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *MassMomentOfInertia) convertToBase(value float64, fromUnit MassMomentOfInertiaUnits) float64 {
	switch fromUnit { 
	case MassMomentOfInertiaGramSquareMeter:
		return (value / 1e3) 
	case MassMomentOfInertiaGramSquareDecimeter:
		return (value / 1e5) 
	case MassMomentOfInertiaGramSquareCentimeter:
		return (value / 1e7) 
	case MassMomentOfInertiaGramSquareMillimeter:
		return (value / 1e9) 
	case MassMomentOfInertiaTonneSquareMeter:
		return (value / 1e-3) 
	case MassMomentOfInertiaTonneSquareDecimeter:
		return (value / 1e-1) 
	case MassMomentOfInertiaTonneSquareCentimeter:
		return (value / 1e1) 
	case MassMomentOfInertiaTonneSquareMilimeter:
		return (value / 1e3) 
	case MassMomentOfInertiaPoundSquareFoot:
		return (value * 4.21401101e-2) 
	case MassMomentOfInertiaPoundSquareInch:
		return (value * 2.9263965e-4) 
	case MassMomentOfInertiaSlugSquareFoot:
		return (value * 1.3558179619) 
	case MassMomentOfInertiaSlugSquareInch:
		return (value * 9.41540242e-3) 
	case MassMomentOfInertiaMilligramSquareMeter:
		return ((value / 1e3) * 0.001) 
	case MassMomentOfInertiaKilogramSquareMeter:
		return ((value / 1e3) * 1000.0) 
	case MassMomentOfInertiaMilligramSquareDecimeter:
		return ((value / 1e5) * 0.001) 
	case MassMomentOfInertiaKilogramSquareDecimeter:
		return ((value / 1e5) * 1000.0) 
	case MassMomentOfInertiaMilligramSquareCentimeter:
		return ((value / 1e7) * 0.001) 
	case MassMomentOfInertiaKilogramSquareCentimeter:
		return ((value / 1e7) * 1000.0) 
	case MassMomentOfInertiaMilligramSquareMillimeter:
		return ((value / 1e9) * 0.001) 
	case MassMomentOfInertiaKilogramSquareMillimeter:
		return ((value / 1e9) * 1000.0) 
	case MassMomentOfInertiaKilotonneSquareMeter:
		return ((value / 1e-3) * 1000.0) 
	case MassMomentOfInertiaMegatonneSquareMeter:
		return ((value / 1e-3) * 1000000.0) 
	case MassMomentOfInertiaKilotonneSquareDecimeter:
		return ((value / 1e-1) * 1000.0) 
	case MassMomentOfInertiaMegatonneSquareDecimeter:
		return ((value / 1e-1) * 1000000.0) 
	case MassMomentOfInertiaKilotonneSquareCentimeter:
		return ((value / 1e1) * 1000.0) 
	case MassMomentOfInertiaMegatonneSquareCentimeter:
		return ((value / 1e1) * 1000000.0) 
	case MassMomentOfInertiaKilotonneSquareMilimeter:
		return ((value / 1e3) * 1000.0) 
	case MassMomentOfInertiaMegatonneSquareMilimeter:
		return ((value / 1e3) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the MassMomentOfInertia in the default unit (KilogramSquareMeter),
// formatted to two decimal places.
func (a MassMomentOfInertia) String() string {
	return a.ToString(MassMomentOfInertiaKilogramSquareMeter, 2)
}

// ToString formats the MassMomentOfInertia value as a string with the specified unit and fractional digits.
// It converts the MassMomentOfInertia to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the MassMomentOfInertia value will be converted (e.g., KilogramSquareMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the MassMomentOfInertia with the unit abbreviation.
func (a *MassMomentOfInertia) ToString(unit MassMomentOfInertiaUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetMassMomentOfInertiaAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetMassMomentOfInertiaAbbreviation(unit))
}

// Equals checks if the given MassMomentOfInertia is equal to the current MassMomentOfInertia.
//
// Parameters:
//    other: The MassMomentOfInertia to compare against.
//
// Returns:
//    bool: Returns true if both MassMomentOfInertia are equal, false otherwise.
func (a *MassMomentOfInertia) Equals(other *MassMomentOfInertia) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current MassMomentOfInertia with another MassMomentOfInertia.
// It returns -1 if the current MassMomentOfInertia is less than the other MassMomentOfInertia, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The MassMomentOfInertia to compare against.
//
// Returns:
//    int: -1 if the current MassMomentOfInertia is less, 1 if greater, and 0 if equal.
func (a *MassMomentOfInertia) CompareTo(other *MassMomentOfInertia) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given MassMomentOfInertia to the current MassMomentOfInertia and returns the result.
// The result is a new MassMomentOfInertia instance with the sum of the values.
//
// Parameters:
//    other: The MassMomentOfInertia to add to the current MassMomentOfInertia.
//
// Returns:
//    *MassMomentOfInertia: A new MassMomentOfInertia instance representing the sum of both MassMomentOfInertia.
func (a *MassMomentOfInertia) Add(other *MassMomentOfInertia) *MassMomentOfInertia {
	return &MassMomentOfInertia{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given MassMomentOfInertia from the current MassMomentOfInertia and returns the result.
// The result is a new MassMomentOfInertia instance with the difference of the values.
//
// Parameters:
//    other: The MassMomentOfInertia to subtract from the current MassMomentOfInertia.
//
// Returns:
//    *MassMomentOfInertia: A new MassMomentOfInertia instance representing the difference of both MassMomentOfInertia.
func (a *MassMomentOfInertia) Subtract(other *MassMomentOfInertia) *MassMomentOfInertia {
	return &MassMomentOfInertia{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current MassMomentOfInertia by the given MassMomentOfInertia and returns the result.
// The result is a new MassMomentOfInertia instance with the product of the values.
//
// Parameters:
//    other: The MassMomentOfInertia to multiply with the current MassMomentOfInertia.
//
// Returns:
//    *MassMomentOfInertia: A new MassMomentOfInertia instance representing the product of both MassMomentOfInertia.
func (a *MassMomentOfInertia) Multiply(other *MassMomentOfInertia) *MassMomentOfInertia {
	return &MassMomentOfInertia{value: a.value * other.BaseValue()}
}

// Divide divides the current MassMomentOfInertia by the given MassMomentOfInertia and returns the result.
// The result is a new MassMomentOfInertia instance with the quotient of the values.
//
// Parameters:
//    other: The MassMomentOfInertia to divide the current MassMomentOfInertia by.
//
// Returns:
//    *MassMomentOfInertia: A new MassMomentOfInertia instance representing the quotient of both MassMomentOfInertia.
func (a *MassMomentOfInertia) Divide(other *MassMomentOfInertia) *MassMomentOfInertia {
	return &MassMomentOfInertia{value: a.value / other.BaseValue()}
}

// GetMassMomentOfInertiaAbbreviation gets the unit abbreviation.
func GetMassMomentOfInertiaAbbreviation(unit MassMomentOfInertiaUnits) string {
	switch unit { 
	case MassMomentOfInertiaGramSquareMeter:
		return "g·m²" 
	case MassMomentOfInertiaGramSquareDecimeter:
		return "g·dm²" 
	case MassMomentOfInertiaGramSquareCentimeter:
		return "g·cm²" 
	case MassMomentOfInertiaGramSquareMillimeter:
		return "g·mm²" 
	case MassMomentOfInertiaTonneSquareMeter:
		return "t·m²" 
	case MassMomentOfInertiaTonneSquareDecimeter:
		return "t·dm²" 
	case MassMomentOfInertiaTonneSquareCentimeter:
		return "t·cm²" 
	case MassMomentOfInertiaTonneSquareMilimeter:
		return "t·mm²" 
	case MassMomentOfInertiaPoundSquareFoot:
		return "lb·ft²" 
	case MassMomentOfInertiaPoundSquareInch:
		return "lb·in²" 
	case MassMomentOfInertiaSlugSquareFoot:
		return "slug·ft²" 
	case MassMomentOfInertiaSlugSquareInch:
		return "slug·in²" 
	case MassMomentOfInertiaMilligramSquareMeter:
		return "mg·m²" 
	case MassMomentOfInertiaKilogramSquareMeter:
		return "kg·m²" 
	case MassMomentOfInertiaMilligramSquareDecimeter:
		return "mg·dm²" 
	case MassMomentOfInertiaKilogramSquareDecimeter:
		return "kg·dm²" 
	case MassMomentOfInertiaMilligramSquareCentimeter:
		return "mg·cm²" 
	case MassMomentOfInertiaKilogramSquareCentimeter:
		return "kg·cm²" 
	case MassMomentOfInertiaMilligramSquareMillimeter:
		return "mg·mm²" 
	case MassMomentOfInertiaKilogramSquareMillimeter:
		return "kg·mm²" 
	case MassMomentOfInertiaKilotonneSquareMeter:
		return "kt·m²" 
	case MassMomentOfInertiaMegatonneSquareMeter:
		return "Mt·m²" 
	case MassMomentOfInertiaKilotonneSquareDecimeter:
		return "kt·dm²" 
	case MassMomentOfInertiaMegatonneSquareDecimeter:
		return "Mt·dm²" 
	case MassMomentOfInertiaKilotonneSquareCentimeter:
		return "kt·cm²" 
	case MassMomentOfInertiaMegatonneSquareCentimeter:
		return "Mt·cm²" 
	case MassMomentOfInertiaKilotonneSquareMilimeter:
		return "kt·mm²" 
	case MassMomentOfInertiaMegatonneSquareMilimeter:
		return "Mt·mm²" 
	default:
		return ""
	}
}