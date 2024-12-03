// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// MassMomentOfInertiaUnits enumeration
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

// MassMomentOfInertiaDto represents an MassMomentOfInertia
type MassMomentOfInertiaDto struct {
	Value float64
	Unit  MassMomentOfInertiaUnits
}

// MassMomentOfInertiaDtoFactory struct to group related functions
type MassMomentOfInertiaDtoFactory struct{}

func (udf MassMomentOfInertiaDtoFactory) FromJSON(data []byte) (*MassMomentOfInertiaDto, error) {
	a := MassMomentOfInertiaDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a MassMomentOfInertiaDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// MassMomentOfInertia struct
type MassMomentOfInertia struct {
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

// MassMomentOfInertiaFactory struct to group related functions
type MassMomentOfInertiaFactory struct{}

func (uf MassMomentOfInertiaFactory) CreateMassMomentOfInertia(value float64, unit MassMomentOfInertiaUnits) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, unit)
}

func (uf MassMomentOfInertiaFactory) FromDto(dto MassMomentOfInertiaDto) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(dto.Value, dto.Unit)
}

func (uf MassMomentOfInertiaFactory) FromDtoJSON(data []byte) (*MassMomentOfInertia, error) {
	unitDto, err := MassMomentOfInertiaDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return MassMomentOfInertiaFactory{}.FromDto(*unitDto)
}


// FromGramSquareMeter creates a new MassMomentOfInertia instance from GramSquareMeter.
func (uf MassMomentOfInertiaFactory) FromGramSquareMeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaGramSquareMeter)
}

// FromGramSquareDecimeter creates a new MassMomentOfInertia instance from GramSquareDecimeter.
func (uf MassMomentOfInertiaFactory) FromGramSquareDecimeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaGramSquareDecimeter)
}

// FromGramSquareCentimeter creates a new MassMomentOfInertia instance from GramSquareCentimeter.
func (uf MassMomentOfInertiaFactory) FromGramSquareCentimeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaGramSquareCentimeter)
}

// FromGramSquareMillimeter creates a new MassMomentOfInertia instance from GramSquareMillimeter.
func (uf MassMomentOfInertiaFactory) FromGramSquareMillimeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaGramSquareMillimeter)
}

// FromTonneSquareMeter creates a new MassMomentOfInertia instance from TonneSquareMeter.
func (uf MassMomentOfInertiaFactory) FromTonneSquareMeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaTonneSquareMeter)
}

// FromTonneSquareDecimeter creates a new MassMomentOfInertia instance from TonneSquareDecimeter.
func (uf MassMomentOfInertiaFactory) FromTonneSquareDecimeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaTonneSquareDecimeter)
}

// FromTonneSquareCentimeter creates a new MassMomentOfInertia instance from TonneSquareCentimeter.
func (uf MassMomentOfInertiaFactory) FromTonneSquareCentimeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaTonneSquareCentimeter)
}

// FromTonneSquareMilimeter creates a new MassMomentOfInertia instance from TonneSquareMilimeter.
func (uf MassMomentOfInertiaFactory) FromTonneSquareMilimeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaTonneSquareMilimeter)
}

// FromPoundSquareFoot creates a new MassMomentOfInertia instance from PoundSquareFoot.
func (uf MassMomentOfInertiaFactory) FromPoundSquareFeet(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaPoundSquareFoot)
}

// FromPoundSquareInch creates a new MassMomentOfInertia instance from PoundSquareInch.
func (uf MassMomentOfInertiaFactory) FromPoundSquareInches(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaPoundSquareInch)
}

// FromSlugSquareFoot creates a new MassMomentOfInertia instance from SlugSquareFoot.
func (uf MassMomentOfInertiaFactory) FromSlugSquareFeet(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaSlugSquareFoot)
}

// FromSlugSquareInch creates a new MassMomentOfInertia instance from SlugSquareInch.
func (uf MassMomentOfInertiaFactory) FromSlugSquareInches(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaSlugSquareInch)
}

// FromMilligramSquareMeter creates a new MassMomentOfInertia instance from MilligramSquareMeter.
func (uf MassMomentOfInertiaFactory) FromMilligramSquareMeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaMilligramSquareMeter)
}

// FromKilogramSquareMeter creates a new MassMomentOfInertia instance from KilogramSquareMeter.
func (uf MassMomentOfInertiaFactory) FromKilogramSquareMeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaKilogramSquareMeter)
}

// FromMilligramSquareDecimeter creates a new MassMomentOfInertia instance from MilligramSquareDecimeter.
func (uf MassMomentOfInertiaFactory) FromMilligramSquareDecimeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaMilligramSquareDecimeter)
}

// FromKilogramSquareDecimeter creates a new MassMomentOfInertia instance from KilogramSquareDecimeter.
func (uf MassMomentOfInertiaFactory) FromKilogramSquareDecimeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaKilogramSquareDecimeter)
}

// FromMilligramSquareCentimeter creates a new MassMomentOfInertia instance from MilligramSquareCentimeter.
func (uf MassMomentOfInertiaFactory) FromMilligramSquareCentimeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaMilligramSquareCentimeter)
}

// FromKilogramSquareCentimeter creates a new MassMomentOfInertia instance from KilogramSquareCentimeter.
func (uf MassMomentOfInertiaFactory) FromKilogramSquareCentimeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaKilogramSquareCentimeter)
}

// FromMilligramSquareMillimeter creates a new MassMomentOfInertia instance from MilligramSquareMillimeter.
func (uf MassMomentOfInertiaFactory) FromMilligramSquareMillimeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaMilligramSquareMillimeter)
}

// FromKilogramSquareMillimeter creates a new MassMomentOfInertia instance from KilogramSquareMillimeter.
func (uf MassMomentOfInertiaFactory) FromKilogramSquareMillimeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaKilogramSquareMillimeter)
}

// FromKilotonneSquareMeter creates a new MassMomentOfInertia instance from KilotonneSquareMeter.
func (uf MassMomentOfInertiaFactory) FromKilotonneSquareMeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaKilotonneSquareMeter)
}

// FromMegatonneSquareMeter creates a new MassMomentOfInertia instance from MegatonneSquareMeter.
func (uf MassMomentOfInertiaFactory) FromMegatonneSquareMeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaMegatonneSquareMeter)
}

// FromKilotonneSquareDecimeter creates a new MassMomentOfInertia instance from KilotonneSquareDecimeter.
func (uf MassMomentOfInertiaFactory) FromKilotonneSquareDecimeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaKilotonneSquareDecimeter)
}

// FromMegatonneSquareDecimeter creates a new MassMomentOfInertia instance from MegatonneSquareDecimeter.
func (uf MassMomentOfInertiaFactory) FromMegatonneSquareDecimeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaMegatonneSquareDecimeter)
}

// FromKilotonneSquareCentimeter creates a new MassMomentOfInertia instance from KilotonneSquareCentimeter.
func (uf MassMomentOfInertiaFactory) FromKilotonneSquareCentimeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaKilotonneSquareCentimeter)
}

// FromMegatonneSquareCentimeter creates a new MassMomentOfInertia instance from MegatonneSquareCentimeter.
func (uf MassMomentOfInertiaFactory) FromMegatonneSquareCentimeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaMegatonneSquareCentimeter)
}

// FromKilotonneSquareMilimeter creates a new MassMomentOfInertia instance from KilotonneSquareMilimeter.
func (uf MassMomentOfInertiaFactory) FromKilotonneSquareMilimeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaKilotonneSquareMilimeter)
}

// FromMegatonneSquareMilimeter creates a new MassMomentOfInertia instance from MegatonneSquareMilimeter.
func (uf MassMomentOfInertiaFactory) FromMegatonneSquareMilimeters(value float64) (*MassMomentOfInertia, error) {
	return newMassMomentOfInertia(value, MassMomentOfInertiaMegatonneSquareMilimeter)
}




// newMassMomentOfInertia creates a new MassMomentOfInertia.
func newMassMomentOfInertia(value float64, fromUnit MassMomentOfInertiaUnits) (*MassMomentOfInertia, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &MassMomentOfInertia{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of MassMomentOfInertia in KilogramSquareMeter.
func (a *MassMomentOfInertia) BaseValue() float64 {
	return a.value
}


// GramSquareMeter returns the value in GramSquareMeter.
func (a *MassMomentOfInertia) GramSquareMeters() float64 {
	if a.gram_square_metersLazy != nil {
		return *a.gram_square_metersLazy
	}
	gram_square_meters := a.convertFromBase(MassMomentOfInertiaGramSquareMeter)
	a.gram_square_metersLazy = &gram_square_meters
	return gram_square_meters
}

// GramSquareDecimeter returns the value in GramSquareDecimeter.
func (a *MassMomentOfInertia) GramSquareDecimeters() float64 {
	if a.gram_square_decimetersLazy != nil {
		return *a.gram_square_decimetersLazy
	}
	gram_square_decimeters := a.convertFromBase(MassMomentOfInertiaGramSquareDecimeter)
	a.gram_square_decimetersLazy = &gram_square_decimeters
	return gram_square_decimeters
}

// GramSquareCentimeter returns the value in GramSquareCentimeter.
func (a *MassMomentOfInertia) GramSquareCentimeters() float64 {
	if a.gram_square_centimetersLazy != nil {
		return *a.gram_square_centimetersLazy
	}
	gram_square_centimeters := a.convertFromBase(MassMomentOfInertiaGramSquareCentimeter)
	a.gram_square_centimetersLazy = &gram_square_centimeters
	return gram_square_centimeters
}

// GramSquareMillimeter returns the value in GramSquareMillimeter.
func (a *MassMomentOfInertia) GramSquareMillimeters() float64 {
	if a.gram_square_millimetersLazy != nil {
		return *a.gram_square_millimetersLazy
	}
	gram_square_millimeters := a.convertFromBase(MassMomentOfInertiaGramSquareMillimeter)
	a.gram_square_millimetersLazy = &gram_square_millimeters
	return gram_square_millimeters
}

// TonneSquareMeter returns the value in TonneSquareMeter.
func (a *MassMomentOfInertia) TonneSquareMeters() float64 {
	if a.tonne_square_metersLazy != nil {
		return *a.tonne_square_metersLazy
	}
	tonne_square_meters := a.convertFromBase(MassMomentOfInertiaTonneSquareMeter)
	a.tonne_square_metersLazy = &tonne_square_meters
	return tonne_square_meters
}

// TonneSquareDecimeter returns the value in TonneSquareDecimeter.
func (a *MassMomentOfInertia) TonneSquareDecimeters() float64 {
	if a.tonne_square_decimetersLazy != nil {
		return *a.tonne_square_decimetersLazy
	}
	tonne_square_decimeters := a.convertFromBase(MassMomentOfInertiaTonneSquareDecimeter)
	a.tonne_square_decimetersLazy = &tonne_square_decimeters
	return tonne_square_decimeters
}

// TonneSquareCentimeter returns the value in TonneSquareCentimeter.
func (a *MassMomentOfInertia) TonneSquareCentimeters() float64 {
	if a.tonne_square_centimetersLazy != nil {
		return *a.tonne_square_centimetersLazy
	}
	tonne_square_centimeters := a.convertFromBase(MassMomentOfInertiaTonneSquareCentimeter)
	a.tonne_square_centimetersLazy = &tonne_square_centimeters
	return tonne_square_centimeters
}

// TonneSquareMilimeter returns the value in TonneSquareMilimeter.
func (a *MassMomentOfInertia) TonneSquareMilimeters() float64 {
	if a.tonne_square_milimetersLazy != nil {
		return *a.tonne_square_milimetersLazy
	}
	tonne_square_milimeters := a.convertFromBase(MassMomentOfInertiaTonneSquareMilimeter)
	a.tonne_square_milimetersLazy = &tonne_square_milimeters
	return tonne_square_milimeters
}

// PoundSquareFoot returns the value in PoundSquareFoot.
func (a *MassMomentOfInertia) PoundSquareFeet() float64 {
	if a.pound_square_feetLazy != nil {
		return *a.pound_square_feetLazy
	}
	pound_square_feet := a.convertFromBase(MassMomentOfInertiaPoundSquareFoot)
	a.pound_square_feetLazy = &pound_square_feet
	return pound_square_feet
}

// PoundSquareInch returns the value in PoundSquareInch.
func (a *MassMomentOfInertia) PoundSquareInches() float64 {
	if a.pound_square_inchesLazy != nil {
		return *a.pound_square_inchesLazy
	}
	pound_square_inches := a.convertFromBase(MassMomentOfInertiaPoundSquareInch)
	a.pound_square_inchesLazy = &pound_square_inches
	return pound_square_inches
}

// SlugSquareFoot returns the value in SlugSquareFoot.
func (a *MassMomentOfInertia) SlugSquareFeet() float64 {
	if a.slug_square_feetLazy != nil {
		return *a.slug_square_feetLazy
	}
	slug_square_feet := a.convertFromBase(MassMomentOfInertiaSlugSquareFoot)
	a.slug_square_feetLazy = &slug_square_feet
	return slug_square_feet
}

// SlugSquareInch returns the value in SlugSquareInch.
func (a *MassMomentOfInertia) SlugSquareInches() float64 {
	if a.slug_square_inchesLazy != nil {
		return *a.slug_square_inchesLazy
	}
	slug_square_inches := a.convertFromBase(MassMomentOfInertiaSlugSquareInch)
	a.slug_square_inchesLazy = &slug_square_inches
	return slug_square_inches
}

// MilligramSquareMeter returns the value in MilligramSquareMeter.
func (a *MassMomentOfInertia) MilligramSquareMeters() float64 {
	if a.milligram_square_metersLazy != nil {
		return *a.milligram_square_metersLazy
	}
	milligram_square_meters := a.convertFromBase(MassMomentOfInertiaMilligramSquareMeter)
	a.milligram_square_metersLazy = &milligram_square_meters
	return milligram_square_meters
}

// KilogramSquareMeter returns the value in KilogramSquareMeter.
func (a *MassMomentOfInertia) KilogramSquareMeters() float64 {
	if a.kilogram_square_metersLazy != nil {
		return *a.kilogram_square_metersLazy
	}
	kilogram_square_meters := a.convertFromBase(MassMomentOfInertiaKilogramSquareMeter)
	a.kilogram_square_metersLazy = &kilogram_square_meters
	return kilogram_square_meters
}

// MilligramSquareDecimeter returns the value in MilligramSquareDecimeter.
func (a *MassMomentOfInertia) MilligramSquareDecimeters() float64 {
	if a.milligram_square_decimetersLazy != nil {
		return *a.milligram_square_decimetersLazy
	}
	milligram_square_decimeters := a.convertFromBase(MassMomentOfInertiaMilligramSquareDecimeter)
	a.milligram_square_decimetersLazy = &milligram_square_decimeters
	return milligram_square_decimeters
}

// KilogramSquareDecimeter returns the value in KilogramSquareDecimeter.
func (a *MassMomentOfInertia) KilogramSquareDecimeters() float64 {
	if a.kilogram_square_decimetersLazy != nil {
		return *a.kilogram_square_decimetersLazy
	}
	kilogram_square_decimeters := a.convertFromBase(MassMomentOfInertiaKilogramSquareDecimeter)
	a.kilogram_square_decimetersLazy = &kilogram_square_decimeters
	return kilogram_square_decimeters
}

// MilligramSquareCentimeter returns the value in MilligramSquareCentimeter.
func (a *MassMomentOfInertia) MilligramSquareCentimeters() float64 {
	if a.milligram_square_centimetersLazy != nil {
		return *a.milligram_square_centimetersLazy
	}
	milligram_square_centimeters := a.convertFromBase(MassMomentOfInertiaMilligramSquareCentimeter)
	a.milligram_square_centimetersLazy = &milligram_square_centimeters
	return milligram_square_centimeters
}

// KilogramSquareCentimeter returns the value in KilogramSquareCentimeter.
func (a *MassMomentOfInertia) KilogramSquareCentimeters() float64 {
	if a.kilogram_square_centimetersLazy != nil {
		return *a.kilogram_square_centimetersLazy
	}
	kilogram_square_centimeters := a.convertFromBase(MassMomentOfInertiaKilogramSquareCentimeter)
	a.kilogram_square_centimetersLazy = &kilogram_square_centimeters
	return kilogram_square_centimeters
}

// MilligramSquareMillimeter returns the value in MilligramSquareMillimeter.
func (a *MassMomentOfInertia) MilligramSquareMillimeters() float64 {
	if a.milligram_square_millimetersLazy != nil {
		return *a.milligram_square_millimetersLazy
	}
	milligram_square_millimeters := a.convertFromBase(MassMomentOfInertiaMilligramSquareMillimeter)
	a.milligram_square_millimetersLazy = &milligram_square_millimeters
	return milligram_square_millimeters
}

// KilogramSquareMillimeter returns the value in KilogramSquareMillimeter.
func (a *MassMomentOfInertia) KilogramSquareMillimeters() float64 {
	if a.kilogram_square_millimetersLazy != nil {
		return *a.kilogram_square_millimetersLazy
	}
	kilogram_square_millimeters := a.convertFromBase(MassMomentOfInertiaKilogramSquareMillimeter)
	a.kilogram_square_millimetersLazy = &kilogram_square_millimeters
	return kilogram_square_millimeters
}

// KilotonneSquareMeter returns the value in KilotonneSquareMeter.
func (a *MassMomentOfInertia) KilotonneSquareMeters() float64 {
	if a.kilotonne_square_metersLazy != nil {
		return *a.kilotonne_square_metersLazy
	}
	kilotonne_square_meters := a.convertFromBase(MassMomentOfInertiaKilotonneSquareMeter)
	a.kilotonne_square_metersLazy = &kilotonne_square_meters
	return kilotonne_square_meters
}

// MegatonneSquareMeter returns the value in MegatonneSquareMeter.
func (a *MassMomentOfInertia) MegatonneSquareMeters() float64 {
	if a.megatonne_square_metersLazy != nil {
		return *a.megatonne_square_metersLazy
	}
	megatonne_square_meters := a.convertFromBase(MassMomentOfInertiaMegatonneSquareMeter)
	a.megatonne_square_metersLazy = &megatonne_square_meters
	return megatonne_square_meters
}

// KilotonneSquareDecimeter returns the value in KilotonneSquareDecimeter.
func (a *MassMomentOfInertia) KilotonneSquareDecimeters() float64 {
	if a.kilotonne_square_decimetersLazy != nil {
		return *a.kilotonne_square_decimetersLazy
	}
	kilotonne_square_decimeters := a.convertFromBase(MassMomentOfInertiaKilotonneSquareDecimeter)
	a.kilotonne_square_decimetersLazy = &kilotonne_square_decimeters
	return kilotonne_square_decimeters
}

// MegatonneSquareDecimeter returns the value in MegatonneSquareDecimeter.
func (a *MassMomentOfInertia) MegatonneSquareDecimeters() float64 {
	if a.megatonne_square_decimetersLazy != nil {
		return *a.megatonne_square_decimetersLazy
	}
	megatonne_square_decimeters := a.convertFromBase(MassMomentOfInertiaMegatonneSquareDecimeter)
	a.megatonne_square_decimetersLazy = &megatonne_square_decimeters
	return megatonne_square_decimeters
}

// KilotonneSquareCentimeter returns the value in KilotonneSquareCentimeter.
func (a *MassMomentOfInertia) KilotonneSquareCentimeters() float64 {
	if a.kilotonne_square_centimetersLazy != nil {
		return *a.kilotonne_square_centimetersLazy
	}
	kilotonne_square_centimeters := a.convertFromBase(MassMomentOfInertiaKilotonneSquareCentimeter)
	a.kilotonne_square_centimetersLazy = &kilotonne_square_centimeters
	return kilotonne_square_centimeters
}

// MegatonneSquareCentimeter returns the value in MegatonneSquareCentimeter.
func (a *MassMomentOfInertia) MegatonneSquareCentimeters() float64 {
	if a.megatonne_square_centimetersLazy != nil {
		return *a.megatonne_square_centimetersLazy
	}
	megatonne_square_centimeters := a.convertFromBase(MassMomentOfInertiaMegatonneSquareCentimeter)
	a.megatonne_square_centimetersLazy = &megatonne_square_centimeters
	return megatonne_square_centimeters
}

// KilotonneSquareMilimeter returns the value in KilotonneSquareMilimeter.
func (a *MassMomentOfInertia) KilotonneSquareMilimeters() float64 {
	if a.kilotonne_square_milimetersLazy != nil {
		return *a.kilotonne_square_milimetersLazy
	}
	kilotonne_square_milimeters := a.convertFromBase(MassMomentOfInertiaKilotonneSquareMilimeter)
	a.kilotonne_square_milimetersLazy = &kilotonne_square_milimeters
	return kilotonne_square_milimeters
}

// MegatonneSquareMilimeter returns the value in MegatonneSquareMilimeter.
func (a *MassMomentOfInertia) MegatonneSquareMilimeters() float64 {
	if a.megatonne_square_milimetersLazy != nil {
		return *a.megatonne_square_milimetersLazy
	}
	megatonne_square_milimeters := a.convertFromBase(MassMomentOfInertiaMegatonneSquareMilimeter)
	a.megatonne_square_milimetersLazy = &megatonne_square_milimeters
	return megatonne_square_milimeters
}


// ToDto creates an MassMomentOfInertiaDto representation.
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

// ToDtoJSON creates an MassMomentOfInertiaDto representation.
func (a *MassMomentOfInertia) ToDtoJSON(holdInUnit *MassMomentOfInertiaUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts MassMomentOfInertia to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a MassMomentOfInertia) String() string {
	return a.ToString(MassMomentOfInertiaKilogramSquareMeter, 2)
}

// ToString formats the MassMomentOfInertia to string.
// fractionalDigits -1 for not mention
func (a *MassMomentOfInertia) ToString(unit MassMomentOfInertiaUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *MassMomentOfInertia) getUnitAbbreviation(unit MassMomentOfInertiaUnits) string {
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

// Check if the given MassMomentOfInertia are equals to the current MassMomentOfInertia
func (a *MassMomentOfInertia) Equals(other *MassMomentOfInertia) bool {
	return a.value == other.BaseValue()
}

// Check if the given MassMomentOfInertia are equals to the current MassMomentOfInertia
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

// Add the given MassMomentOfInertia to the current MassMomentOfInertia.
func (a *MassMomentOfInertia) Add(other *MassMomentOfInertia) *MassMomentOfInertia {
	return &MassMomentOfInertia{value: a.value + other.BaseValue()}
}

// Subtract the given MassMomentOfInertia to the current MassMomentOfInertia.
func (a *MassMomentOfInertia) Subtract(other *MassMomentOfInertia) *MassMomentOfInertia {
	return &MassMomentOfInertia{value: a.value - other.BaseValue()}
}

// Multiply the given MassMomentOfInertia to the current MassMomentOfInertia.
func (a *MassMomentOfInertia) Multiply(other *MassMomentOfInertia) *MassMomentOfInertia {
	return &MassMomentOfInertia{value: a.value * other.BaseValue()}
}

// Divide the given MassMomentOfInertia to the current MassMomentOfInertia.
func (a *MassMomentOfInertia) Divide(other *MassMomentOfInertia) *MassMomentOfInertia {
	return &MassMomentOfInertia{value: a.value / other.BaseValue()}
}