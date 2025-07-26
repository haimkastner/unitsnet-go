// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// InformationUnits defines various units of Information.
type InformationUnits string

const (
    
        // 
        InformationByte InformationUnits = "Byte"
        // 
        InformationOctet InformationUnits = "Octet"
        // 
        InformationBit InformationUnits = "Bit"
        // 
        InformationKilobyte InformationUnits = "Kilobyte"
        // 
        InformationMegabyte InformationUnits = "Megabyte"
        // 
        InformationGigabyte InformationUnits = "Gigabyte"
        // 
        InformationTerabyte InformationUnits = "Terabyte"
        // 
        InformationPetabyte InformationUnits = "Petabyte"
        // 
        InformationExabyte InformationUnits = "Exabyte"
        // 
        InformationKibibyte InformationUnits = "Kibibyte"
        // 
        InformationMebibyte InformationUnits = "Mebibyte"
        // 
        InformationGibibyte InformationUnits = "Gibibyte"
        // 
        InformationTebibyte InformationUnits = "Tebibyte"
        // 
        InformationPebibyte InformationUnits = "Pebibyte"
        // 
        InformationExbibyte InformationUnits = "Exbibyte"
        // 
        InformationKilooctet InformationUnits = "Kilooctet"
        // 
        InformationMegaoctet InformationUnits = "Megaoctet"
        // 
        InformationGigaoctet InformationUnits = "Gigaoctet"
        // 
        InformationTeraoctet InformationUnits = "Teraoctet"
        // 
        InformationPetaoctet InformationUnits = "Petaoctet"
        // 
        InformationExaoctet InformationUnits = "Exaoctet"
        // 
        InformationKibioctet InformationUnits = "Kibioctet"
        // 
        InformationMebioctet InformationUnits = "Mebioctet"
        // 
        InformationGibioctet InformationUnits = "Gibioctet"
        // 
        InformationTebioctet InformationUnits = "Tebioctet"
        // 
        InformationPebioctet InformationUnits = "Pebioctet"
        // 
        InformationExbioctet InformationUnits = "Exbioctet"
        // 
        InformationKilobit InformationUnits = "Kilobit"
        // 
        InformationMegabit InformationUnits = "Megabit"
        // 
        InformationGigabit InformationUnits = "Gigabit"
        // 
        InformationTerabit InformationUnits = "Terabit"
        // 
        InformationPetabit InformationUnits = "Petabit"
        // 
        InformationExabit InformationUnits = "Exabit"
        // 
        InformationKibibit InformationUnits = "Kibibit"
        // 
        InformationMebibit InformationUnits = "Mebibit"
        // 
        InformationGibibit InformationUnits = "Gibibit"
        // 
        InformationTebibit InformationUnits = "Tebibit"
        // 
        InformationPebibit InformationUnits = "Pebibit"
        // 
        InformationExbibit InformationUnits = "Exbibit"
)

var internalInformationUnitsMap = map[InformationUnits]bool{
	
	InformationByte: true,
	InformationOctet: true,
	InformationBit: true,
	InformationKilobyte: true,
	InformationMegabyte: true,
	InformationGigabyte: true,
	InformationTerabyte: true,
	InformationPetabyte: true,
	InformationExabyte: true,
	InformationKibibyte: true,
	InformationMebibyte: true,
	InformationGibibyte: true,
	InformationTebibyte: true,
	InformationPebibyte: true,
	InformationExbibyte: true,
	InformationKilooctet: true,
	InformationMegaoctet: true,
	InformationGigaoctet: true,
	InformationTeraoctet: true,
	InformationPetaoctet: true,
	InformationExaoctet: true,
	InformationKibioctet: true,
	InformationMebioctet: true,
	InformationGibioctet: true,
	InformationTebioctet: true,
	InformationPebioctet: true,
	InformationExbioctet: true,
	InformationKilobit: true,
	InformationMegabit: true,
	InformationGigabit: true,
	InformationTerabit: true,
	InformationPetabit: true,
	InformationExabit: true,
	InformationKibibit: true,
	InformationMebibit: true,
	InformationGibibit: true,
	InformationTebibit: true,
	InformationPebibit: true,
	InformationExbibit: true,
}

// InformationDto represents a Information measurement with a numerical value and its corresponding unit.
type InformationDto struct {
    // Value is the numerical representation of the Information.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the Information, as defined in the InformationUnits enumeration.
	Unit  InformationUnits `json:"unit" validate:"required,oneof=Byte Octet Bit Kilobyte Megabyte Gigabyte Terabyte Petabyte Exabyte Kibibyte Mebibyte Gibibyte Tebibyte Pebibyte Exbibyte Kilooctet Megaoctet Gigaoctet Teraoctet Petaoctet Exaoctet Kibioctet Mebioctet Gibioctet Tebioctet Pebioctet Exbioctet Kilobit Megabit Gigabit Terabit Petabit Exabit Kibibit Mebibit Gibibit Tebibit Pebibit Exbibit"`
}

// InformationDtoFactory groups methods for creating and serializing InformationDto objects.
type InformationDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a InformationDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf InformationDtoFactory) FromJSON(data []byte) (*InformationDto, error) {
	a := InformationDto{}

    // Parse JSON into InformationDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a InformationDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a InformationDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// Information represents a measurement in a of Information.
//
// In computing and telecommunications, a unit of information is the capacity of some standard data storage system or communication channel, used to measure the capacities of other systems and channels. In information theory, units of information are also used to measure the information contents or entropy of random variables.
type Information struct {
	// value is the base measurement stored internally.
	value       float64
    
    bytesLazy *float64 
    octetsLazy *float64 
    bitsLazy *float64 
    kilobytesLazy *float64 
    megabytesLazy *float64 
    gigabytesLazy *float64 
    terabytesLazy *float64 
    petabytesLazy *float64 
    exabytesLazy *float64 
    kibibytesLazy *float64 
    mebibytesLazy *float64 
    gibibytesLazy *float64 
    tebibytesLazy *float64 
    pebibytesLazy *float64 
    exbibytesLazy *float64 
    kilooctetsLazy *float64 
    megaoctetsLazy *float64 
    gigaoctetsLazy *float64 
    teraoctetsLazy *float64 
    petaoctetsLazy *float64 
    exaoctetsLazy *float64 
    kibioctetsLazy *float64 
    mebioctetsLazy *float64 
    gibioctetsLazy *float64 
    tebioctetsLazy *float64 
    pebioctetsLazy *float64 
    exbioctetsLazy *float64 
    kilobitsLazy *float64 
    megabitsLazy *float64 
    gigabitsLazy *float64 
    terabitsLazy *float64 
    petabitsLazy *float64 
    exabitsLazy *float64 
    kibibitsLazy *float64 
    mebibitsLazy *float64 
    gibibitsLazy *float64 
    tebibitsLazy *float64 
    pebibitsLazy *float64 
    exbibitsLazy *float64 
}

// InformationFactory groups methods for creating Information instances.
type InformationFactory struct{}

// CreateInformation creates a new Information instance from the given value and unit.
func (uf InformationFactory) CreateInformation(value float64, unit InformationUnits) (*Information, error) {
	return newInformation(value, unit)
}

// FromDto converts a InformationDto to a Information instance.
func (uf InformationFactory) FromDto(dto InformationDto) (*Information, error) {
	return newInformation(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a Information instance.
func (uf InformationFactory) FromDtoJSON(data []byte) (*Information, error) {
	unitDto, err := InformationDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse InformationDto from JSON: %w", err)
	}
	return InformationFactory{}.FromDto(*unitDto)
}


// FromBytes creates a new Information instance from a value in Bytes.
func (uf InformationFactory) FromBytes(value float64) (*Information, error) {
	return newInformation(value, InformationByte)
}

// FromOctets creates a new Information instance from a value in Octets.
func (uf InformationFactory) FromOctets(value float64) (*Information, error) {
	return newInformation(value, InformationOctet)
}

// FromBits creates a new Information instance from a value in Bits.
func (uf InformationFactory) FromBits(value float64) (*Information, error) {
	return newInformation(value, InformationBit)
}

// FromKilobytes creates a new Information instance from a value in Kilobytes.
func (uf InformationFactory) FromKilobytes(value float64) (*Information, error) {
	return newInformation(value, InformationKilobyte)
}

// FromMegabytes creates a new Information instance from a value in Megabytes.
func (uf InformationFactory) FromMegabytes(value float64) (*Information, error) {
	return newInformation(value, InformationMegabyte)
}

// FromGigabytes creates a new Information instance from a value in Gigabytes.
func (uf InformationFactory) FromGigabytes(value float64) (*Information, error) {
	return newInformation(value, InformationGigabyte)
}

// FromTerabytes creates a new Information instance from a value in Terabytes.
func (uf InformationFactory) FromTerabytes(value float64) (*Information, error) {
	return newInformation(value, InformationTerabyte)
}

// FromPetabytes creates a new Information instance from a value in Petabytes.
func (uf InformationFactory) FromPetabytes(value float64) (*Information, error) {
	return newInformation(value, InformationPetabyte)
}

// FromExabytes creates a new Information instance from a value in Exabytes.
func (uf InformationFactory) FromExabytes(value float64) (*Information, error) {
	return newInformation(value, InformationExabyte)
}

// FromKibibytes creates a new Information instance from a value in Kibibytes.
func (uf InformationFactory) FromKibibytes(value float64) (*Information, error) {
	return newInformation(value, InformationKibibyte)
}

// FromMebibytes creates a new Information instance from a value in Mebibytes.
func (uf InformationFactory) FromMebibytes(value float64) (*Information, error) {
	return newInformation(value, InformationMebibyte)
}

// FromGibibytes creates a new Information instance from a value in Gibibytes.
func (uf InformationFactory) FromGibibytes(value float64) (*Information, error) {
	return newInformation(value, InformationGibibyte)
}

// FromTebibytes creates a new Information instance from a value in Tebibytes.
func (uf InformationFactory) FromTebibytes(value float64) (*Information, error) {
	return newInformation(value, InformationTebibyte)
}

// FromPebibytes creates a new Information instance from a value in Pebibytes.
func (uf InformationFactory) FromPebibytes(value float64) (*Information, error) {
	return newInformation(value, InformationPebibyte)
}

// FromExbibytes creates a new Information instance from a value in Exbibytes.
func (uf InformationFactory) FromExbibytes(value float64) (*Information, error) {
	return newInformation(value, InformationExbibyte)
}

// FromKilooctets creates a new Information instance from a value in Kilooctets.
func (uf InformationFactory) FromKilooctets(value float64) (*Information, error) {
	return newInformation(value, InformationKilooctet)
}

// FromMegaoctets creates a new Information instance from a value in Megaoctets.
func (uf InformationFactory) FromMegaoctets(value float64) (*Information, error) {
	return newInformation(value, InformationMegaoctet)
}

// FromGigaoctets creates a new Information instance from a value in Gigaoctets.
func (uf InformationFactory) FromGigaoctets(value float64) (*Information, error) {
	return newInformation(value, InformationGigaoctet)
}

// FromTeraoctets creates a new Information instance from a value in Teraoctets.
func (uf InformationFactory) FromTeraoctets(value float64) (*Information, error) {
	return newInformation(value, InformationTeraoctet)
}

// FromPetaoctets creates a new Information instance from a value in Petaoctets.
func (uf InformationFactory) FromPetaoctets(value float64) (*Information, error) {
	return newInformation(value, InformationPetaoctet)
}

// FromExaoctets creates a new Information instance from a value in Exaoctets.
func (uf InformationFactory) FromExaoctets(value float64) (*Information, error) {
	return newInformation(value, InformationExaoctet)
}

// FromKibioctets creates a new Information instance from a value in Kibioctets.
func (uf InformationFactory) FromKibioctets(value float64) (*Information, error) {
	return newInformation(value, InformationKibioctet)
}

// FromMebioctets creates a new Information instance from a value in Mebioctets.
func (uf InformationFactory) FromMebioctets(value float64) (*Information, error) {
	return newInformation(value, InformationMebioctet)
}

// FromGibioctets creates a new Information instance from a value in Gibioctets.
func (uf InformationFactory) FromGibioctets(value float64) (*Information, error) {
	return newInformation(value, InformationGibioctet)
}

// FromTebioctets creates a new Information instance from a value in Tebioctets.
func (uf InformationFactory) FromTebioctets(value float64) (*Information, error) {
	return newInformation(value, InformationTebioctet)
}

// FromPebioctets creates a new Information instance from a value in Pebioctets.
func (uf InformationFactory) FromPebioctets(value float64) (*Information, error) {
	return newInformation(value, InformationPebioctet)
}

// FromExbioctets creates a new Information instance from a value in Exbioctets.
func (uf InformationFactory) FromExbioctets(value float64) (*Information, error) {
	return newInformation(value, InformationExbioctet)
}

// FromKilobits creates a new Information instance from a value in Kilobits.
func (uf InformationFactory) FromKilobits(value float64) (*Information, error) {
	return newInformation(value, InformationKilobit)
}

// FromMegabits creates a new Information instance from a value in Megabits.
func (uf InformationFactory) FromMegabits(value float64) (*Information, error) {
	return newInformation(value, InformationMegabit)
}

// FromGigabits creates a new Information instance from a value in Gigabits.
func (uf InformationFactory) FromGigabits(value float64) (*Information, error) {
	return newInformation(value, InformationGigabit)
}

// FromTerabits creates a new Information instance from a value in Terabits.
func (uf InformationFactory) FromTerabits(value float64) (*Information, error) {
	return newInformation(value, InformationTerabit)
}

// FromPetabits creates a new Information instance from a value in Petabits.
func (uf InformationFactory) FromPetabits(value float64) (*Information, error) {
	return newInformation(value, InformationPetabit)
}

// FromExabits creates a new Information instance from a value in Exabits.
func (uf InformationFactory) FromExabits(value float64) (*Information, error) {
	return newInformation(value, InformationExabit)
}

// FromKibibits creates a new Information instance from a value in Kibibits.
func (uf InformationFactory) FromKibibits(value float64) (*Information, error) {
	return newInformation(value, InformationKibibit)
}

// FromMebibits creates a new Information instance from a value in Mebibits.
func (uf InformationFactory) FromMebibits(value float64) (*Information, error) {
	return newInformation(value, InformationMebibit)
}

// FromGibibits creates a new Information instance from a value in Gibibits.
func (uf InformationFactory) FromGibibits(value float64) (*Information, error) {
	return newInformation(value, InformationGibibit)
}

// FromTebibits creates a new Information instance from a value in Tebibits.
func (uf InformationFactory) FromTebibits(value float64) (*Information, error) {
	return newInformation(value, InformationTebibit)
}

// FromPebibits creates a new Information instance from a value in Pebibits.
func (uf InformationFactory) FromPebibits(value float64) (*Information, error) {
	return newInformation(value, InformationPebibit)
}

// FromExbibits creates a new Information instance from a value in Exbibits.
func (uf InformationFactory) FromExbibits(value float64) (*Information, error) {
	return newInformation(value, InformationExbibit)
}


// newInformation creates a new Information.
func newInformation(value float64, fromUnit InformationUnits) (*Information, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalInformationUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in InformationUnits", fromUnit)
	}
	a := &Information{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Information in Bit unit (the base/default quantity).
func (a *Information) BaseValue() float64 {
	return a.value
}


// Bytes returns the Information value in Bytes.
//
// 
func (a *Information) Bytes() float64 {
	if a.bytesLazy != nil {
		return *a.bytesLazy
	}
	bytes := a.convertFromBase(InformationByte)
	a.bytesLazy = &bytes
	return bytes
}

// Octets returns the Information value in Octets.
//
// 
func (a *Information) Octets() float64 {
	if a.octetsLazy != nil {
		return *a.octetsLazy
	}
	octets := a.convertFromBase(InformationOctet)
	a.octetsLazy = &octets
	return octets
}

// Bits returns the Information value in Bits.
//
// 
func (a *Information) Bits() float64 {
	if a.bitsLazy != nil {
		return *a.bitsLazy
	}
	bits := a.convertFromBase(InformationBit)
	a.bitsLazy = &bits
	return bits
}

// Kilobytes returns the Information value in Kilobytes.
//
// 
func (a *Information) Kilobytes() float64 {
	if a.kilobytesLazy != nil {
		return *a.kilobytesLazy
	}
	kilobytes := a.convertFromBase(InformationKilobyte)
	a.kilobytesLazy = &kilobytes
	return kilobytes
}

// Megabytes returns the Information value in Megabytes.
//
// 
func (a *Information) Megabytes() float64 {
	if a.megabytesLazy != nil {
		return *a.megabytesLazy
	}
	megabytes := a.convertFromBase(InformationMegabyte)
	a.megabytesLazy = &megabytes
	return megabytes
}

// Gigabytes returns the Information value in Gigabytes.
//
// 
func (a *Information) Gigabytes() float64 {
	if a.gigabytesLazy != nil {
		return *a.gigabytesLazy
	}
	gigabytes := a.convertFromBase(InformationGigabyte)
	a.gigabytesLazy = &gigabytes
	return gigabytes
}

// Terabytes returns the Information value in Terabytes.
//
// 
func (a *Information) Terabytes() float64 {
	if a.terabytesLazy != nil {
		return *a.terabytesLazy
	}
	terabytes := a.convertFromBase(InformationTerabyte)
	a.terabytesLazy = &terabytes
	return terabytes
}

// Petabytes returns the Information value in Petabytes.
//
// 
func (a *Information) Petabytes() float64 {
	if a.petabytesLazy != nil {
		return *a.petabytesLazy
	}
	petabytes := a.convertFromBase(InformationPetabyte)
	a.petabytesLazy = &petabytes
	return petabytes
}

// Exabytes returns the Information value in Exabytes.
//
// 
func (a *Information) Exabytes() float64 {
	if a.exabytesLazy != nil {
		return *a.exabytesLazy
	}
	exabytes := a.convertFromBase(InformationExabyte)
	a.exabytesLazy = &exabytes
	return exabytes
}

// Kibibytes returns the Information value in Kibibytes.
//
// 
func (a *Information) Kibibytes() float64 {
	if a.kibibytesLazy != nil {
		return *a.kibibytesLazy
	}
	kibibytes := a.convertFromBase(InformationKibibyte)
	a.kibibytesLazy = &kibibytes
	return kibibytes
}

// Mebibytes returns the Information value in Mebibytes.
//
// 
func (a *Information) Mebibytes() float64 {
	if a.mebibytesLazy != nil {
		return *a.mebibytesLazy
	}
	mebibytes := a.convertFromBase(InformationMebibyte)
	a.mebibytesLazy = &mebibytes
	return mebibytes
}

// Gibibytes returns the Information value in Gibibytes.
//
// 
func (a *Information) Gibibytes() float64 {
	if a.gibibytesLazy != nil {
		return *a.gibibytesLazy
	}
	gibibytes := a.convertFromBase(InformationGibibyte)
	a.gibibytesLazy = &gibibytes
	return gibibytes
}

// Tebibytes returns the Information value in Tebibytes.
//
// 
func (a *Information) Tebibytes() float64 {
	if a.tebibytesLazy != nil {
		return *a.tebibytesLazy
	}
	tebibytes := a.convertFromBase(InformationTebibyte)
	a.tebibytesLazy = &tebibytes
	return tebibytes
}

// Pebibytes returns the Information value in Pebibytes.
//
// 
func (a *Information) Pebibytes() float64 {
	if a.pebibytesLazy != nil {
		return *a.pebibytesLazy
	}
	pebibytes := a.convertFromBase(InformationPebibyte)
	a.pebibytesLazy = &pebibytes
	return pebibytes
}

// Exbibytes returns the Information value in Exbibytes.
//
// 
func (a *Information) Exbibytes() float64 {
	if a.exbibytesLazy != nil {
		return *a.exbibytesLazy
	}
	exbibytes := a.convertFromBase(InformationExbibyte)
	a.exbibytesLazy = &exbibytes
	return exbibytes
}

// Kilooctets returns the Information value in Kilooctets.
//
// 
func (a *Information) Kilooctets() float64 {
	if a.kilooctetsLazy != nil {
		return *a.kilooctetsLazy
	}
	kilooctets := a.convertFromBase(InformationKilooctet)
	a.kilooctetsLazy = &kilooctets
	return kilooctets
}

// Megaoctets returns the Information value in Megaoctets.
//
// 
func (a *Information) Megaoctets() float64 {
	if a.megaoctetsLazy != nil {
		return *a.megaoctetsLazy
	}
	megaoctets := a.convertFromBase(InformationMegaoctet)
	a.megaoctetsLazy = &megaoctets
	return megaoctets
}

// Gigaoctets returns the Information value in Gigaoctets.
//
// 
func (a *Information) Gigaoctets() float64 {
	if a.gigaoctetsLazy != nil {
		return *a.gigaoctetsLazy
	}
	gigaoctets := a.convertFromBase(InformationGigaoctet)
	a.gigaoctetsLazy = &gigaoctets
	return gigaoctets
}

// Teraoctets returns the Information value in Teraoctets.
//
// 
func (a *Information) Teraoctets() float64 {
	if a.teraoctetsLazy != nil {
		return *a.teraoctetsLazy
	}
	teraoctets := a.convertFromBase(InformationTeraoctet)
	a.teraoctetsLazy = &teraoctets
	return teraoctets
}

// Petaoctets returns the Information value in Petaoctets.
//
// 
func (a *Information) Petaoctets() float64 {
	if a.petaoctetsLazy != nil {
		return *a.petaoctetsLazy
	}
	petaoctets := a.convertFromBase(InformationPetaoctet)
	a.petaoctetsLazy = &petaoctets
	return petaoctets
}

// Exaoctets returns the Information value in Exaoctets.
//
// 
func (a *Information) Exaoctets() float64 {
	if a.exaoctetsLazy != nil {
		return *a.exaoctetsLazy
	}
	exaoctets := a.convertFromBase(InformationExaoctet)
	a.exaoctetsLazy = &exaoctets
	return exaoctets
}

// Kibioctets returns the Information value in Kibioctets.
//
// 
func (a *Information) Kibioctets() float64 {
	if a.kibioctetsLazy != nil {
		return *a.kibioctetsLazy
	}
	kibioctets := a.convertFromBase(InformationKibioctet)
	a.kibioctetsLazy = &kibioctets
	return kibioctets
}

// Mebioctets returns the Information value in Mebioctets.
//
// 
func (a *Information) Mebioctets() float64 {
	if a.mebioctetsLazy != nil {
		return *a.mebioctetsLazy
	}
	mebioctets := a.convertFromBase(InformationMebioctet)
	a.mebioctetsLazy = &mebioctets
	return mebioctets
}

// Gibioctets returns the Information value in Gibioctets.
//
// 
func (a *Information) Gibioctets() float64 {
	if a.gibioctetsLazy != nil {
		return *a.gibioctetsLazy
	}
	gibioctets := a.convertFromBase(InformationGibioctet)
	a.gibioctetsLazy = &gibioctets
	return gibioctets
}

// Tebioctets returns the Information value in Tebioctets.
//
// 
func (a *Information) Tebioctets() float64 {
	if a.tebioctetsLazy != nil {
		return *a.tebioctetsLazy
	}
	tebioctets := a.convertFromBase(InformationTebioctet)
	a.tebioctetsLazy = &tebioctets
	return tebioctets
}

// Pebioctets returns the Information value in Pebioctets.
//
// 
func (a *Information) Pebioctets() float64 {
	if a.pebioctetsLazy != nil {
		return *a.pebioctetsLazy
	}
	pebioctets := a.convertFromBase(InformationPebioctet)
	a.pebioctetsLazy = &pebioctets
	return pebioctets
}

// Exbioctets returns the Information value in Exbioctets.
//
// 
func (a *Information) Exbioctets() float64 {
	if a.exbioctetsLazy != nil {
		return *a.exbioctetsLazy
	}
	exbioctets := a.convertFromBase(InformationExbioctet)
	a.exbioctetsLazy = &exbioctets
	return exbioctets
}

// Kilobits returns the Information value in Kilobits.
//
// 
func (a *Information) Kilobits() float64 {
	if a.kilobitsLazy != nil {
		return *a.kilobitsLazy
	}
	kilobits := a.convertFromBase(InformationKilobit)
	a.kilobitsLazy = &kilobits
	return kilobits
}

// Megabits returns the Information value in Megabits.
//
// 
func (a *Information) Megabits() float64 {
	if a.megabitsLazy != nil {
		return *a.megabitsLazy
	}
	megabits := a.convertFromBase(InformationMegabit)
	a.megabitsLazy = &megabits
	return megabits
}

// Gigabits returns the Information value in Gigabits.
//
// 
func (a *Information) Gigabits() float64 {
	if a.gigabitsLazy != nil {
		return *a.gigabitsLazy
	}
	gigabits := a.convertFromBase(InformationGigabit)
	a.gigabitsLazy = &gigabits
	return gigabits
}

// Terabits returns the Information value in Terabits.
//
// 
func (a *Information) Terabits() float64 {
	if a.terabitsLazy != nil {
		return *a.terabitsLazy
	}
	terabits := a.convertFromBase(InformationTerabit)
	a.terabitsLazy = &terabits
	return terabits
}

// Petabits returns the Information value in Petabits.
//
// 
func (a *Information) Petabits() float64 {
	if a.petabitsLazy != nil {
		return *a.petabitsLazy
	}
	petabits := a.convertFromBase(InformationPetabit)
	a.petabitsLazy = &petabits
	return petabits
}

// Exabits returns the Information value in Exabits.
//
// 
func (a *Information) Exabits() float64 {
	if a.exabitsLazy != nil {
		return *a.exabitsLazy
	}
	exabits := a.convertFromBase(InformationExabit)
	a.exabitsLazy = &exabits
	return exabits
}

// Kibibits returns the Information value in Kibibits.
//
// 
func (a *Information) Kibibits() float64 {
	if a.kibibitsLazy != nil {
		return *a.kibibitsLazy
	}
	kibibits := a.convertFromBase(InformationKibibit)
	a.kibibitsLazy = &kibibits
	return kibibits
}

// Mebibits returns the Information value in Mebibits.
//
// 
func (a *Information) Mebibits() float64 {
	if a.mebibitsLazy != nil {
		return *a.mebibitsLazy
	}
	mebibits := a.convertFromBase(InformationMebibit)
	a.mebibitsLazy = &mebibits
	return mebibits
}

// Gibibits returns the Information value in Gibibits.
//
// 
func (a *Information) Gibibits() float64 {
	if a.gibibitsLazy != nil {
		return *a.gibibitsLazy
	}
	gibibits := a.convertFromBase(InformationGibibit)
	a.gibibitsLazy = &gibibits
	return gibibits
}

// Tebibits returns the Information value in Tebibits.
//
// 
func (a *Information) Tebibits() float64 {
	if a.tebibitsLazy != nil {
		return *a.tebibitsLazy
	}
	tebibits := a.convertFromBase(InformationTebibit)
	a.tebibitsLazy = &tebibits
	return tebibits
}

// Pebibits returns the Information value in Pebibits.
//
// 
func (a *Information) Pebibits() float64 {
	if a.pebibitsLazy != nil {
		return *a.pebibitsLazy
	}
	pebibits := a.convertFromBase(InformationPebibit)
	a.pebibitsLazy = &pebibits
	return pebibits
}

// Exbibits returns the Information value in Exbibits.
//
// 
func (a *Information) Exbibits() float64 {
	if a.exbibitsLazy != nil {
		return *a.exbibitsLazy
	}
	exbibits := a.convertFromBase(InformationExbibit)
	a.exbibitsLazy = &exbibits
	return exbibits
}


// ToDto creates a InformationDto representation from the Information instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Bit by default.
func (a *Information) ToDto(holdInUnit *InformationUnits) InformationDto {
	if holdInUnit == nil {
		defaultUnit := InformationBit // Default value
		holdInUnit = &defaultUnit
	}

	return InformationDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the Information instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Bit by default.
func (a *Information) ToDtoJSON(holdInUnit *InformationUnits) ([]byte, error) {
	// Convert to InformationDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a Information to a specific unit value.
// The function uses the provided unit type (InformationUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *Information) Convert(toUnit InformationUnits) float64 {
	switch toUnit { 
    case InformationByte:
		return a.Bytes()
    case InformationOctet:
		return a.Octets()
    case InformationBit:
		return a.Bits()
    case InformationKilobyte:
		return a.Kilobytes()
    case InformationMegabyte:
		return a.Megabytes()
    case InformationGigabyte:
		return a.Gigabytes()
    case InformationTerabyte:
		return a.Terabytes()
    case InformationPetabyte:
		return a.Petabytes()
    case InformationExabyte:
		return a.Exabytes()
    case InformationKibibyte:
		return a.Kibibytes()
    case InformationMebibyte:
		return a.Mebibytes()
    case InformationGibibyte:
		return a.Gibibytes()
    case InformationTebibyte:
		return a.Tebibytes()
    case InformationPebibyte:
		return a.Pebibytes()
    case InformationExbibyte:
		return a.Exbibytes()
    case InformationKilooctet:
		return a.Kilooctets()
    case InformationMegaoctet:
		return a.Megaoctets()
    case InformationGigaoctet:
		return a.Gigaoctets()
    case InformationTeraoctet:
		return a.Teraoctets()
    case InformationPetaoctet:
		return a.Petaoctets()
    case InformationExaoctet:
		return a.Exaoctets()
    case InformationKibioctet:
		return a.Kibioctets()
    case InformationMebioctet:
		return a.Mebioctets()
    case InformationGibioctet:
		return a.Gibioctets()
    case InformationTebioctet:
		return a.Tebioctets()
    case InformationPebioctet:
		return a.Pebioctets()
    case InformationExbioctet:
		return a.Exbioctets()
    case InformationKilobit:
		return a.Kilobits()
    case InformationMegabit:
		return a.Megabits()
    case InformationGigabit:
		return a.Gigabits()
    case InformationTerabit:
		return a.Terabits()
    case InformationPetabit:
		return a.Petabits()
    case InformationExabit:
		return a.Exabits()
    case InformationKibibit:
		return a.Kibibits()
    case InformationMebibit:
		return a.Mebibits()
    case InformationGibibit:
		return a.Gibibits()
    case InformationTebibit:
		return a.Tebibits()
    case InformationPebibit:
		return a.Pebibits()
    case InformationExbibit:
		return a.Exbibits()
	default:
		return math.NaN()
	}
}

func (a *Information) convertFromBase(toUnit InformationUnits) float64 {
    value := a.value
	switch toUnit { 
	case InformationByte:
		return (value / 8) 
	case InformationOctet:
		return (value / 8) 
	case InformationBit:
		return (value) 
	case InformationKilobyte:
		return ((value / 8) / 1000.0) 
	case InformationMegabyte:
		return ((value / 8) / 1000000.0) 
	case InformationGigabyte:
		return ((value / 8) / 1000000000.0) 
	case InformationTerabyte:
		return ((value / 8) / 1000000000000.0) 
	case InformationPetabyte:
		return ((value / 8) / 1000000000000000.0) 
	case InformationExabyte:
		return ((value / 8) / 1e+18) 
	case InformationKibibyte:
		return ((value / 8) / 1024) 
	case InformationMebibyte:
		return ((value / 8) / 1048576) 
	case InformationGibibyte:
		return ((value / 8) / 1073741824) 
	case InformationTebibyte:
		return ((value / 8) / 1099511627776) 
	case InformationPebibyte:
		return ((value / 8) / 1125899906842624) 
	case InformationExbibyte:
		return ((value / 8) / 1152921504606846976) 
	case InformationKilooctet:
		return ((value / 8) / 1000.0) 
	case InformationMegaoctet:
		return ((value / 8) / 1000000.0) 
	case InformationGigaoctet:
		return ((value / 8) / 1000000000.0) 
	case InformationTeraoctet:
		return ((value / 8) / 1000000000000.0) 
	case InformationPetaoctet:
		return ((value / 8) / 1000000000000000.0) 
	case InformationExaoctet:
		return ((value / 8) / 1e+18) 
	case InformationKibioctet:
		return ((value / 8) / 1024) 
	case InformationMebioctet:
		return ((value / 8) / 1048576) 
	case InformationGibioctet:
		return ((value / 8) / 1073741824) 
	case InformationTebioctet:
		return ((value / 8) / 1099511627776) 
	case InformationPebioctet:
		return ((value / 8) / 1125899906842624) 
	case InformationExbioctet:
		return ((value / 8) / 1152921504606846976) 
	case InformationKilobit:
		return ((value) / 1000.0) 
	case InformationMegabit:
		return ((value) / 1000000.0) 
	case InformationGigabit:
		return ((value) / 1000000000.0) 
	case InformationTerabit:
		return ((value) / 1000000000000.0) 
	case InformationPetabit:
		return ((value) / 1000000000000000.0) 
	case InformationExabit:
		return ((value) / 1e+18) 
	case InformationKibibit:
		return ((value) / 1024) 
	case InformationMebibit:
		return ((value) / 1048576) 
	case InformationGibibit:
		return ((value) / 1073741824) 
	case InformationTebibit:
		return ((value) / 1099511627776) 
	case InformationPebibit:
		return ((value) / 1125899906842624) 
	case InformationExbibit:
		return ((value) / 1152921504606846976) 
	default:
		return math.NaN()
	}
}

func (a *Information) convertToBase(value float64, fromUnit InformationUnits) float64 {
	switch fromUnit { 
	case InformationByte:
		return (value * 8) 
	case InformationOctet:
		return (value * 8) 
	case InformationBit:
		return (value) 
	case InformationKilobyte:
		return ((value * 8) * 1000.0) 
	case InformationMegabyte:
		return ((value * 8) * 1000000.0) 
	case InformationGigabyte:
		return ((value * 8) * 1000000000.0) 
	case InformationTerabyte:
		return ((value * 8) * 1000000000000.0) 
	case InformationPetabyte:
		return ((value * 8) * 1000000000000000.0) 
	case InformationExabyte:
		return ((value * 8) * 1e+18) 
	case InformationKibibyte:
		return ((value * 8) * 1024) 
	case InformationMebibyte:
		return ((value * 8) * 1048576) 
	case InformationGibibyte:
		return ((value * 8) * 1073741824) 
	case InformationTebibyte:
		return ((value * 8) * 1099511627776) 
	case InformationPebibyte:
		return ((value * 8) * 1125899906842624) 
	case InformationExbibyte:
		return ((value * 8) * 1152921504606846976) 
	case InformationKilooctet:
		return ((value * 8) * 1000.0) 
	case InformationMegaoctet:
		return ((value * 8) * 1000000.0) 
	case InformationGigaoctet:
		return ((value * 8) * 1000000000.0) 
	case InformationTeraoctet:
		return ((value * 8) * 1000000000000.0) 
	case InformationPetaoctet:
		return ((value * 8) * 1000000000000000.0) 
	case InformationExaoctet:
		return ((value * 8) * 1e+18) 
	case InformationKibioctet:
		return ((value * 8) * 1024) 
	case InformationMebioctet:
		return ((value * 8) * 1048576) 
	case InformationGibioctet:
		return ((value * 8) * 1073741824) 
	case InformationTebioctet:
		return ((value * 8) * 1099511627776) 
	case InformationPebioctet:
		return ((value * 8) * 1125899906842624) 
	case InformationExbioctet:
		return ((value * 8) * 1152921504606846976) 
	case InformationKilobit:
		return ((value) * 1000.0) 
	case InformationMegabit:
		return ((value) * 1000000.0) 
	case InformationGigabit:
		return ((value) * 1000000000.0) 
	case InformationTerabit:
		return ((value) * 1000000000000.0) 
	case InformationPetabit:
		return ((value) * 1000000000000000.0) 
	case InformationExabit:
		return ((value) * 1e+18) 
	case InformationKibibit:
		return ((value) * 1024) 
	case InformationMebibit:
		return ((value) * 1048576) 
	case InformationGibibit:
		return ((value) * 1073741824) 
	case InformationTebibit:
		return ((value) * 1099511627776) 
	case InformationPebibit:
		return ((value) * 1125899906842624) 
	case InformationExbibit:
		return ((value) * 1152921504606846976) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the Information in the default unit (Bit),
// formatted to two decimal places.
func (a Information) String() string {
	return a.ToString(InformationBit, 2)
}

// ToString formats the Information value as a string with the specified unit and fractional digits.
// It converts the Information to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the Information value will be converted (e.g., Bit))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the Information with the unit abbreviation.
func (a *Information) ToString(unit InformationUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetInformationAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetInformationAbbreviation(unit))
}

// Equals checks if the given Information is equal to the current Information.
//
// Parameters:
//    other: The Information to compare against.
//
// Returns:
//    bool: Returns true if both Information are equal, false otherwise.
func (a *Information) Equals(other *Information) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current Information with another Information.
// It returns -1 if the current Information is less than the other Information, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The Information to compare against.
//
// Returns:
//    int: -1 if the current Information is less, 1 if greater, and 0 if equal.
func (a *Information) CompareTo(other *Information) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given Information to the current Information and returns the result.
// The result is a new Information instance with the sum of the values.
//
// Parameters:
//    other: The Information to add to the current Information.
//
// Returns:
//    *Information: A new Information instance representing the sum of both Information.
func (a *Information) Add(other *Information) *Information {
	return &Information{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given Information from the current Information and returns the result.
// The result is a new Information instance with the difference of the values.
//
// Parameters:
//    other: The Information to subtract from the current Information.
//
// Returns:
//    *Information: A new Information instance representing the difference of both Information.
func (a *Information) Subtract(other *Information) *Information {
	return &Information{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current Information by the given Information and returns the result.
// The result is a new Information instance with the product of the values.
//
// Parameters:
//    other: The Information to multiply with the current Information.
//
// Returns:
//    *Information: A new Information instance representing the product of both Information.
func (a *Information) Multiply(other *Information) *Information {
	return &Information{value: a.value * other.BaseValue()}
}

// Divide divides the current Information by the given Information and returns the result.
// The result is a new Information instance with the quotient of the values.
//
// Parameters:
//    other: The Information to divide the current Information by.
//
// Returns:
//    *Information: A new Information instance representing the quotient of both Information.
func (a *Information) Divide(other *Information) *Information {
	return &Information{value: a.value / other.BaseValue()}
}

// GetInformationAbbreviation gets the unit abbreviation.
func GetInformationAbbreviation(unit InformationUnits) string {
	switch unit { 
	case InformationByte:
		return "B" 
	case InformationOctet:
		return "o" 
	case InformationBit:
		return "b" 
	case InformationKilobyte:
		return "kB" 
	case InformationMegabyte:
		return "MB" 
	case InformationGigabyte:
		return "GB" 
	case InformationTerabyte:
		return "TB" 
	case InformationPetabyte:
		return "PB" 
	case InformationExabyte:
		return "EB" 
	case InformationKibibyte:
		return "KiBB" 
	case InformationMebibyte:
		return "MiBB" 
	case InformationGibibyte:
		return "GiBB" 
	case InformationTebibyte:
		return "TiBB" 
	case InformationPebibyte:
		return "PiBB" 
	case InformationExbibyte:
		return "EiBB" 
	case InformationKilooctet:
		return "ko" 
	case InformationMegaoctet:
		return "Mo" 
	case InformationGigaoctet:
		return "Go" 
	case InformationTeraoctet:
		return "To" 
	case InformationPetaoctet:
		return "Po" 
	case InformationExaoctet:
		return "Eo" 
	case InformationKibioctet:
		return "KiBo" 
	case InformationMebioctet:
		return "MiBo" 
	case InformationGibioctet:
		return "GiBo" 
	case InformationTebioctet:
		return "TiBo" 
	case InformationPebioctet:
		return "PiBo" 
	case InformationExbioctet:
		return "EiBo" 
	case InformationKilobit:
		return "kb" 
	case InformationMegabit:
		return "Mb" 
	case InformationGigabit:
		return "Gb" 
	case InformationTerabit:
		return "Tb" 
	case InformationPetabit:
		return "Pb" 
	case InformationExabit:
		return "Eb" 
	case InformationKibibit:
		return "KiBb" 
	case InformationMebibit:
		return "MiBb" 
	case InformationGibibit:
		return "GiBb" 
	case InformationTebibit:
		return "TiBb" 
	case InformationPebibit:
		return "PiBb" 
	case InformationExbibit:
		return "EiBb" 
	default:
		return ""
	}
}