// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// BitRateUnits defines various units of BitRate.
type BitRateUnits string

const (
    
        // 
        BitRateBitPerSecond BitRateUnits = "BitPerSecond"
        // 
        BitRateBytePerSecond BitRateUnits = "BytePerSecond"
        // 
        BitRateKilobitPerSecond BitRateUnits = "KilobitPerSecond"
        // 
        BitRateMegabitPerSecond BitRateUnits = "MegabitPerSecond"
        // 
        BitRateGigabitPerSecond BitRateUnits = "GigabitPerSecond"
        // 
        BitRateTerabitPerSecond BitRateUnits = "TerabitPerSecond"
        // 
        BitRatePetabitPerSecond BitRateUnits = "PetabitPerSecond"
        // 
        BitRateExabitPerSecond BitRateUnits = "ExabitPerSecond"
        // 
        BitRateKibibitPerSecond BitRateUnits = "KibibitPerSecond"
        // 
        BitRateMebibitPerSecond BitRateUnits = "MebibitPerSecond"
        // 
        BitRateGibibitPerSecond BitRateUnits = "GibibitPerSecond"
        // 
        BitRateTebibitPerSecond BitRateUnits = "TebibitPerSecond"
        // 
        BitRatePebibitPerSecond BitRateUnits = "PebibitPerSecond"
        // 
        BitRateExbibitPerSecond BitRateUnits = "ExbibitPerSecond"
        // 
        BitRateKilobytePerSecond BitRateUnits = "KilobytePerSecond"
        // 
        BitRateMegabytePerSecond BitRateUnits = "MegabytePerSecond"
        // 
        BitRateGigabytePerSecond BitRateUnits = "GigabytePerSecond"
        // 
        BitRateTerabytePerSecond BitRateUnits = "TerabytePerSecond"
        // 
        BitRatePetabytePerSecond BitRateUnits = "PetabytePerSecond"
        // 
        BitRateExabytePerSecond BitRateUnits = "ExabytePerSecond"
        // 
        BitRateKibibytePerSecond BitRateUnits = "KibibytePerSecond"
        // 
        BitRateMebibytePerSecond BitRateUnits = "MebibytePerSecond"
        // 
        BitRateGibibytePerSecond BitRateUnits = "GibibytePerSecond"
        // 
        BitRateTebibytePerSecond BitRateUnits = "TebibytePerSecond"
        // 
        BitRatePebibytePerSecond BitRateUnits = "PebibytePerSecond"
        // 
        BitRateExbibytePerSecond BitRateUnits = "ExbibytePerSecond"
)

// BitRateDto represents a BitRate measurement with a numerical value and its corresponding unit.
type BitRateDto struct {
    // Value is the numerical representation of the BitRate.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the BitRate, as defined in the BitRateUnits enumeration.
	Unit  BitRateUnits `json:"unit"`
}

// BitRateDtoFactory groups methods for creating and serializing BitRateDto objects.
type BitRateDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a BitRateDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf BitRateDtoFactory) FromJSON(data []byte) (*BitRateDto, error) {
	a := BitRateDto{}

    // Parse JSON into BitRateDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a BitRateDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a BitRateDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// BitRate represents a measurement in a of BitRate.
//
// In telecommunications and computing, bit rate is the number of bits that are conveyed or processed per unit of time.
type BitRate struct {
	// value is the base measurement stored internally.
	value       float64
    
    bits_per_secondLazy *float64 
    bytes_per_secondLazy *float64 
    kilobits_per_secondLazy *float64 
    megabits_per_secondLazy *float64 
    gigabits_per_secondLazy *float64 
    terabits_per_secondLazy *float64 
    petabits_per_secondLazy *float64 
    exabits_per_secondLazy *float64 
    kibibits_per_secondLazy *float64 
    mebibits_per_secondLazy *float64 
    gibibits_per_secondLazy *float64 
    tebibits_per_secondLazy *float64 
    pebibits_per_secondLazy *float64 
    exbibits_per_secondLazy *float64 
    kilobytes_per_secondLazy *float64 
    megabytes_per_secondLazy *float64 
    gigabytes_per_secondLazy *float64 
    terabytes_per_secondLazy *float64 
    petabytes_per_secondLazy *float64 
    exabytes_per_secondLazy *float64 
    kibibytes_per_secondLazy *float64 
    mebibytes_per_secondLazy *float64 
    gibibytes_per_secondLazy *float64 
    tebibytes_per_secondLazy *float64 
    pebibytes_per_secondLazy *float64 
    exbibytes_per_secondLazy *float64 
}

// BitRateFactory groups methods for creating BitRate instances.
type BitRateFactory struct{}

// CreateBitRate creates a new BitRate instance from the given value and unit.
func (uf BitRateFactory) CreateBitRate(value float64, unit BitRateUnits) (*BitRate, error) {
	return newBitRate(value, unit)
}

// FromDto converts a BitRateDto to a BitRate instance.
func (uf BitRateFactory) FromDto(dto BitRateDto) (*BitRate, error) {
	return newBitRate(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a BitRate instance.
func (uf BitRateFactory) FromDtoJSON(data []byte) (*BitRate, error) {
	unitDto, err := BitRateDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse BitRateDto from JSON: %w", err)
	}
	return BitRateFactory{}.FromDto(*unitDto)
}


// FromBitsPerSecond creates a new BitRate instance from a value in BitsPerSecond.
func (uf BitRateFactory) FromBitsPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateBitPerSecond)
}

// FromBytesPerSecond creates a new BitRate instance from a value in BytesPerSecond.
func (uf BitRateFactory) FromBytesPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateBytePerSecond)
}

// FromKilobitsPerSecond creates a new BitRate instance from a value in KilobitsPerSecond.
func (uf BitRateFactory) FromKilobitsPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateKilobitPerSecond)
}

// FromMegabitsPerSecond creates a new BitRate instance from a value in MegabitsPerSecond.
func (uf BitRateFactory) FromMegabitsPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateMegabitPerSecond)
}

// FromGigabitsPerSecond creates a new BitRate instance from a value in GigabitsPerSecond.
func (uf BitRateFactory) FromGigabitsPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateGigabitPerSecond)
}

// FromTerabitsPerSecond creates a new BitRate instance from a value in TerabitsPerSecond.
func (uf BitRateFactory) FromTerabitsPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateTerabitPerSecond)
}

// FromPetabitsPerSecond creates a new BitRate instance from a value in PetabitsPerSecond.
func (uf BitRateFactory) FromPetabitsPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRatePetabitPerSecond)
}

// FromExabitsPerSecond creates a new BitRate instance from a value in ExabitsPerSecond.
func (uf BitRateFactory) FromExabitsPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateExabitPerSecond)
}

// FromKibibitsPerSecond creates a new BitRate instance from a value in KibibitsPerSecond.
func (uf BitRateFactory) FromKibibitsPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateKibibitPerSecond)
}

// FromMebibitsPerSecond creates a new BitRate instance from a value in MebibitsPerSecond.
func (uf BitRateFactory) FromMebibitsPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateMebibitPerSecond)
}

// FromGibibitsPerSecond creates a new BitRate instance from a value in GibibitsPerSecond.
func (uf BitRateFactory) FromGibibitsPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateGibibitPerSecond)
}

// FromTebibitsPerSecond creates a new BitRate instance from a value in TebibitsPerSecond.
func (uf BitRateFactory) FromTebibitsPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateTebibitPerSecond)
}

// FromPebibitsPerSecond creates a new BitRate instance from a value in PebibitsPerSecond.
func (uf BitRateFactory) FromPebibitsPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRatePebibitPerSecond)
}

// FromExbibitsPerSecond creates a new BitRate instance from a value in ExbibitsPerSecond.
func (uf BitRateFactory) FromExbibitsPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateExbibitPerSecond)
}

// FromKilobytesPerSecond creates a new BitRate instance from a value in KilobytesPerSecond.
func (uf BitRateFactory) FromKilobytesPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateKilobytePerSecond)
}

// FromMegabytesPerSecond creates a new BitRate instance from a value in MegabytesPerSecond.
func (uf BitRateFactory) FromMegabytesPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateMegabytePerSecond)
}

// FromGigabytesPerSecond creates a new BitRate instance from a value in GigabytesPerSecond.
func (uf BitRateFactory) FromGigabytesPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateGigabytePerSecond)
}

// FromTerabytesPerSecond creates a new BitRate instance from a value in TerabytesPerSecond.
func (uf BitRateFactory) FromTerabytesPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateTerabytePerSecond)
}

// FromPetabytesPerSecond creates a new BitRate instance from a value in PetabytesPerSecond.
func (uf BitRateFactory) FromPetabytesPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRatePetabytePerSecond)
}

// FromExabytesPerSecond creates a new BitRate instance from a value in ExabytesPerSecond.
func (uf BitRateFactory) FromExabytesPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateExabytePerSecond)
}

// FromKibibytesPerSecond creates a new BitRate instance from a value in KibibytesPerSecond.
func (uf BitRateFactory) FromKibibytesPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateKibibytePerSecond)
}

// FromMebibytesPerSecond creates a new BitRate instance from a value in MebibytesPerSecond.
func (uf BitRateFactory) FromMebibytesPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateMebibytePerSecond)
}

// FromGibibytesPerSecond creates a new BitRate instance from a value in GibibytesPerSecond.
func (uf BitRateFactory) FromGibibytesPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateGibibytePerSecond)
}

// FromTebibytesPerSecond creates a new BitRate instance from a value in TebibytesPerSecond.
func (uf BitRateFactory) FromTebibytesPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateTebibytePerSecond)
}

// FromPebibytesPerSecond creates a new BitRate instance from a value in PebibytesPerSecond.
func (uf BitRateFactory) FromPebibytesPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRatePebibytePerSecond)
}

// FromExbibytesPerSecond creates a new BitRate instance from a value in ExbibytesPerSecond.
func (uf BitRateFactory) FromExbibytesPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateExbibytePerSecond)
}


// newBitRate creates a new BitRate.
func newBitRate(value float64, fromUnit BitRateUnits) (*BitRate, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &BitRate{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of BitRate in BitPerSecond unit (the base/default quantity).
func (a *BitRate) BaseValue() float64 {
	return a.value
}


// BitsPerSecond returns the BitRate value in BitsPerSecond.
//
// 
func (a *BitRate) BitsPerSecond() float64 {
	if a.bits_per_secondLazy != nil {
		return *a.bits_per_secondLazy
	}
	bits_per_second := a.convertFromBase(BitRateBitPerSecond)
	a.bits_per_secondLazy = &bits_per_second
	return bits_per_second
}

// BytesPerSecond returns the BitRate value in BytesPerSecond.
//
// 
func (a *BitRate) BytesPerSecond() float64 {
	if a.bytes_per_secondLazy != nil {
		return *a.bytes_per_secondLazy
	}
	bytes_per_second := a.convertFromBase(BitRateBytePerSecond)
	a.bytes_per_secondLazy = &bytes_per_second
	return bytes_per_second
}

// KilobitsPerSecond returns the BitRate value in KilobitsPerSecond.
//
// 
func (a *BitRate) KilobitsPerSecond() float64 {
	if a.kilobits_per_secondLazy != nil {
		return *a.kilobits_per_secondLazy
	}
	kilobits_per_second := a.convertFromBase(BitRateKilobitPerSecond)
	a.kilobits_per_secondLazy = &kilobits_per_second
	return kilobits_per_second
}

// MegabitsPerSecond returns the BitRate value in MegabitsPerSecond.
//
// 
func (a *BitRate) MegabitsPerSecond() float64 {
	if a.megabits_per_secondLazy != nil {
		return *a.megabits_per_secondLazy
	}
	megabits_per_second := a.convertFromBase(BitRateMegabitPerSecond)
	a.megabits_per_secondLazy = &megabits_per_second
	return megabits_per_second
}

// GigabitsPerSecond returns the BitRate value in GigabitsPerSecond.
//
// 
func (a *BitRate) GigabitsPerSecond() float64 {
	if a.gigabits_per_secondLazy != nil {
		return *a.gigabits_per_secondLazy
	}
	gigabits_per_second := a.convertFromBase(BitRateGigabitPerSecond)
	a.gigabits_per_secondLazy = &gigabits_per_second
	return gigabits_per_second
}

// TerabitsPerSecond returns the BitRate value in TerabitsPerSecond.
//
// 
func (a *BitRate) TerabitsPerSecond() float64 {
	if a.terabits_per_secondLazy != nil {
		return *a.terabits_per_secondLazy
	}
	terabits_per_second := a.convertFromBase(BitRateTerabitPerSecond)
	a.terabits_per_secondLazy = &terabits_per_second
	return terabits_per_second
}

// PetabitsPerSecond returns the BitRate value in PetabitsPerSecond.
//
// 
func (a *BitRate) PetabitsPerSecond() float64 {
	if a.petabits_per_secondLazy != nil {
		return *a.petabits_per_secondLazy
	}
	petabits_per_second := a.convertFromBase(BitRatePetabitPerSecond)
	a.petabits_per_secondLazy = &petabits_per_second
	return petabits_per_second
}

// ExabitsPerSecond returns the BitRate value in ExabitsPerSecond.
//
// 
func (a *BitRate) ExabitsPerSecond() float64 {
	if a.exabits_per_secondLazy != nil {
		return *a.exabits_per_secondLazy
	}
	exabits_per_second := a.convertFromBase(BitRateExabitPerSecond)
	a.exabits_per_secondLazy = &exabits_per_second
	return exabits_per_second
}

// KibibitsPerSecond returns the BitRate value in KibibitsPerSecond.
//
// 
func (a *BitRate) KibibitsPerSecond() float64 {
	if a.kibibits_per_secondLazy != nil {
		return *a.kibibits_per_secondLazy
	}
	kibibits_per_second := a.convertFromBase(BitRateKibibitPerSecond)
	a.kibibits_per_secondLazy = &kibibits_per_second
	return kibibits_per_second
}

// MebibitsPerSecond returns the BitRate value in MebibitsPerSecond.
//
// 
func (a *BitRate) MebibitsPerSecond() float64 {
	if a.mebibits_per_secondLazy != nil {
		return *a.mebibits_per_secondLazy
	}
	mebibits_per_second := a.convertFromBase(BitRateMebibitPerSecond)
	a.mebibits_per_secondLazy = &mebibits_per_second
	return mebibits_per_second
}

// GibibitsPerSecond returns the BitRate value in GibibitsPerSecond.
//
// 
func (a *BitRate) GibibitsPerSecond() float64 {
	if a.gibibits_per_secondLazy != nil {
		return *a.gibibits_per_secondLazy
	}
	gibibits_per_second := a.convertFromBase(BitRateGibibitPerSecond)
	a.gibibits_per_secondLazy = &gibibits_per_second
	return gibibits_per_second
}

// TebibitsPerSecond returns the BitRate value in TebibitsPerSecond.
//
// 
func (a *BitRate) TebibitsPerSecond() float64 {
	if a.tebibits_per_secondLazy != nil {
		return *a.tebibits_per_secondLazy
	}
	tebibits_per_second := a.convertFromBase(BitRateTebibitPerSecond)
	a.tebibits_per_secondLazy = &tebibits_per_second
	return tebibits_per_second
}

// PebibitsPerSecond returns the BitRate value in PebibitsPerSecond.
//
// 
func (a *BitRate) PebibitsPerSecond() float64 {
	if a.pebibits_per_secondLazy != nil {
		return *a.pebibits_per_secondLazy
	}
	pebibits_per_second := a.convertFromBase(BitRatePebibitPerSecond)
	a.pebibits_per_secondLazy = &pebibits_per_second
	return pebibits_per_second
}

// ExbibitsPerSecond returns the BitRate value in ExbibitsPerSecond.
//
// 
func (a *BitRate) ExbibitsPerSecond() float64 {
	if a.exbibits_per_secondLazy != nil {
		return *a.exbibits_per_secondLazy
	}
	exbibits_per_second := a.convertFromBase(BitRateExbibitPerSecond)
	a.exbibits_per_secondLazy = &exbibits_per_second
	return exbibits_per_second
}

// KilobytesPerSecond returns the BitRate value in KilobytesPerSecond.
//
// 
func (a *BitRate) KilobytesPerSecond() float64 {
	if a.kilobytes_per_secondLazy != nil {
		return *a.kilobytes_per_secondLazy
	}
	kilobytes_per_second := a.convertFromBase(BitRateKilobytePerSecond)
	a.kilobytes_per_secondLazy = &kilobytes_per_second
	return kilobytes_per_second
}

// MegabytesPerSecond returns the BitRate value in MegabytesPerSecond.
//
// 
func (a *BitRate) MegabytesPerSecond() float64 {
	if a.megabytes_per_secondLazy != nil {
		return *a.megabytes_per_secondLazy
	}
	megabytes_per_second := a.convertFromBase(BitRateMegabytePerSecond)
	a.megabytes_per_secondLazy = &megabytes_per_second
	return megabytes_per_second
}

// GigabytesPerSecond returns the BitRate value in GigabytesPerSecond.
//
// 
func (a *BitRate) GigabytesPerSecond() float64 {
	if a.gigabytes_per_secondLazy != nil {
		return *a.gigabytes_per_secondLazy
	}
	gigabytes_per_second := a.convertFromBase(BitRateGigabytePerSecond)
	a.gigabytes_per_secondLazy = &gigabytes_per_second
	return gigabytes_per_second
}

// TerabytesPerSecond returns the BitRate value in TerabytesPerSecond.
//
// 
func (a *BitRate) TerabytesPerSecond() float64 {
	if a.terabytes_per_secondLazy != nil {
		return *a.terabytes_per_secondLazy
	}
	terabytes_per_second := a.convertFromBase(BitRateTerabytePerSecond)
	a.terabytes_per_secondLazy = &terabytes_per_second
	return terabytes_per_second
}

// PetabytesPerSecond returns the BitRate value in PetabytesPerSecond.
//
// 
func (a *BitRate) PetabytesPerSecond() float64 {
	if a.petabytes_per_secondLazy != nil {
		return *a.petabytes_per_secondLazy
	}
	petabytes_per_second := a.convertFromBase(BitRatePetabytePerSecond)
	a.petabytes_per_secondLazy = &petabytes_per_second
	return petabytes_per_second
}

// ExabytesPerSecond returns the BitRate value in ExabytesPerSecond.
//
// 
func (a *BitRate) ExabytesPerSecond() float64 {
	if a.exabytes_per_secondLazy != nil {
		return *a.exabytes_per_secondLazy
	}
	exabytes_per_second := a.convertFromBase(BitRateExabytePerSecond)
	a.exabytes_per_secondLazy = &exabytes_per_second
	return exabytes_per_second
}

// KibibytesPerSecond returns the BitRate value in KibibytesPerSecond.
//
// 
func (a *BitRate) KibibytesPerSecond() float64 {
	if a.kibibytes_per_secondLazy != nil {
		return *a.kibibytes_per_secondLazy
	}
	kibibytes_per_second := a.convertFromBase(BitRateKibibytePerSecond)
	a.kibibytes_per_secondLazy = &kibibytes_per_second
	return kibibytes_per_second
}

// MebibytesPerSecond returns the BitRate value in MebibytesPerSecond.
//
// 
func (a *BitRate) MebibytesPerSecond() float64 {
	if a.mebibytes_per_secondLazy != nil {
		return *a.mebibytes_per_secondLazy
	}
	mebibytes_per_second := a.convertFromBase(BitRateMebibytePerSecond)
	a.mebibytes_per_secondLazy = &mebibytes_per_second
	return mebibytes_per_second
}

// GibibytesPerSecond returns the BitRate value in GibibytesPerSecond.
//
// 
func (a *BitRate) GibibytesPerSecond() float64 {
	if a.gibibytes_per_secondLazy != nil {
		return *a.gibibytes_per_secondLazy
	}
	gibibytes_per_second := a.convertFromBase(BitRateGibibytePerSecond)
	a.gibibytes_per_secondLazy = &gibibytes_per_second
	return gibibytes_per_second
}

// TebibytesPerSecond returns the BitRate value in TebibytesPerSecond.
//
// 
func (a *BitRate) TebibytesPerSecond() float64 {
	if a.tebibytes_per_secondLazy != nil {
		return *a.tebibytes_per_secondLazy
	}
	tebibytes_per_second := a.convertFromBase(BitRateTebibytePerSecond)
	a.tebibytes_per_secondLazy = &tebibytes_per_second
	return tebibytes_per_second
}

// PebibytesPerSecond returns the BitRate value in PebibytesPerSecond.
//
// 
func (a *BitRate) PebibytesPerSecond() float64 {
	if a.pebibytes_per_secondLazy != nil {
		return *a.pebibytes_per_secondLazy
	}
	pebibytes_per_second := a.convertFromBase(BitRatePebibytePerSecond)
	a.pebibytes_per_secondLazy = &pebibytes_per_second
	return pebibytes_per_second
}

// ExbibytesPerSecond returns the BitRate value in ExbibytesPerSecond.
//
// 
func (a *BitRate) ExbibytesPerSecond() float64 {
	if a.exbibytes_per_secondLazy != nil {
		return *a.exbibytes_per_secondLazy
	}
	exbibytes_per_second := a.convertFromBase(BitRateExbibytePerSecond)
	a.exbibytes_per_secondLazy = &exbibytes_per_second
	return exbibytes_per_second
}


// ToDto creates a BitRateDto representation from the BitRate instance.
//
// If the provided holdInUnit is nil, the value will be repesented by BitPerSecond by default.
func (a *BitRate) ToDto(holdInUnit *BitRateUnits) BitRateDto {
	if holdInUnit == nil {
		defaultUnit := BitRateBitPerSecond // Default value
		holdInUnit = &defaultUnit
	}

	return BitRateDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the BitRate instance.
//
// If the provided holdInUnit is nil, the value will be repesented by BitPerSecond by default.
func (a *BitRate) ToDtoJSON(holdInUnit *BitRateUnits) ([]byte, error) {
	// Convert to BitRateDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a BitRate to a specific unit value.
// The function uses the provided unit type (BitRateUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *BitRate) Convert(toUnit BitRateUnits) float64 {
	switch toUnit { 
    case BitRateBitPerSecond:
		return a.BitsPerSecond()
    case BitRateBytePerSecond:
		return a.BytesPerSecond()
    case BitRateKilobitPerSecond:
		return a.KilobitsPerSecond()
    case BitRateMegabitPerSecond:
		return a.MegabitsPerSecond()
    case BitRateGigabitPerSecond:
		return a.GigabitsPerSecond()
    case BitRateTerabitPerSecond:
		return a.TerabitsPerSecond()
    case BitRatePetabitPerSecond:
		return a.PetabitsPerSecond()
    case BitRateExabitPerSecond:
		return a.ExabitsPerSecond()
    case BitRateKibibitPerSecond:
		return a.KibibitsPerSecond()
    case BitRateMebibitPerSecond:
		return a.MebibitsPerSecond()
    case BitRateGibibitPerSecond:
		return a.GibibitsPerSecond()
    case BitRateTebibitPerSecond:
		return a.TebibitsPerSecond()
    case BitRatePebibitPerSecond:
		return a.PebibitsPerSecond()
    case BitRateExbibitPerSecond:
		return a.ExbibitsPerSecond()
    case BitRateKilobytePerSecond:
		return a.KilobytesPerSecond()
    case BitRateMegabytePerSecond:
		return a.MegabytesPerSecond()
    case BitRateGigabytePerSecond:
		return a.GigabytesPerSecond()
    case BitRateTerabytePerSecond:
		return a.TerabytesPerSecond()
    case BitRatePetabytePerSecond:
		return a.PetabytesPerSecond()
    case BitRateExabytePerSecond:
		return a.ExabytesPerSecond()
    case BitRateKibibytePerSecond:
		return a.KibibytesPerSecond()
    case BitRateMebibytePerSecond:
		return a.MebibytesPerSecond()
    case BitRateGibibytePerSecond:
		return a.GibibytesPerSecond()
    case BitRateTebibytePerSecond:
		return a.TebibytesPerSecond()
    case BitRatePebibytePerSecond:
		return a.PebibytesPerSecond()
    case BitRateExbibytePerSecond:
		return a.ExbibytesPerSecond()
	default:
		return math.NaN()
	}
}

func (a *BitRate) convertFromBase(toUnit BitRateUnits) float64 {
    value := a.value
	switch toUnit { 
	case BitRateBitPerSecond:
		return (value) 
	case BitRateBytePerSecond:
		return (value / 8) 
	case BitRateKilobitPerSecond:
		return ((value) / 1000.0) 
	case BitRateMegabitPerSecond:
		return ((value) / 1000000.0) 
	case BitRateGigabitPerSecond:
		return ((value) / 1000000000.0) 
	case BitRateTerabitPerSecond:
		return ((value) / 1000000000000.0) 
	case BitRatePetabitPerSecond:
		return ((value) / 1000000000000000.0) 
	case BitRateExabitPerSecond:
		return ((value) / 1e+18) 
	case BitRateKibibitPerSecond:
		return ((value) / 1024) 
	case BitRateMebibitPerSecond:
		return ((value) / 1048576) 
	case BitRateGibibitPerSecond:
		return ((value) / 1073741824) 
	case BitRateTebibitPerSecond:
		return ((value) / 1099511627776) 
	case BitRatePebibitPerSecond:
		return ((value) / 1125899906842624) 
	case BitRateExbibitPerSecond:
		return ((value) / 1152921504606846976) 
	case BitRateKilobytePerSecond:
		return ((value / 8) / 1000.0) 
	case BitRateMegabytePerSecond:
		return ((value / 8) / 1000000.0) 
	case BitRateGigabytePerSecond:
		return ((value / 8) / 1000000000.0) 
	case BitRateTerabytePerSecond:
		return ((value / 8) / 1000000000000.0) 
	case BitRatePetabytePerSecond:
		return ((value / 8) / 1000000000000000.0) 
	case BitRateExabytePerSecond:
		return ((value / 8) / 1e+18) 
	case BitRateKibibytePerSecond:
		return ((value / 8) / 1024) 
	case BitRateMebibytePerSecond:
		return ((value / 8) / 1048576) 
	case BitRateGibibytePerSecond:
		return ((value / 8) / 1073741824) 
	case BitRateTebibytePerSecond:
		return ((value / 8) / 1099511627776) 
	case BitRatePebibytePerSecond:
		return ((value / 8) / 1125899906842624) 
	case BitRateExbibytePerSecond:
		return ((value / 8) / 1152921504606846976) 
	default:
		return math.NaN()
	}
}

func (a *BitRate) convertToBase(value float64, fromUnit BitRateUnits) float64 {
	switch fromUnit { 
	case BitRateBitPerSecond:
		return (value) 
	case BitRateBytePerSecond:
		return (value * 8) 
	case BitRateKilobitPerSecond:
		return ((value) * 1000.0) 
	case BitRateMegabitPerSecond:
		return ((value) * 1000000.0) 
	case BitRateGigabitPerSecond:
		return ((value) * 1000000000.0) 
	case BitRateTerabitPerSecond:
		return ((value) * 1000000000000.0) 
	case BitRatePetabitPerSecond:
		return ((value) * 1000000000000000.0) 
	case BitRateExabitPerSecond:
		return ((value) * 1e+18) 
	case BitRateKibibitPerSecond:
		return ((value) * 1024) 
	case BitRateMebibitPerSecond:
		return ((value) * 1048576) 
	case BitRateGibibitPerSecond:
		return ((value) * 1073741824) 
	case BitRateTebibitPerSecond:
		return ((value) * 1099511627776) 
	case BitRatePebibitPerSecond:
		return ((value) * 1125899906842624) 
	case BitRateExbibitPerSecond:
		return ((value) * 1152921504606846976) 
	case BitRateKilobytePerSecond:
		return ((value * 8) * 1000.0) 
	case BitRateMegabytePerSecond:
		return ((value * 8) * 1000000.0) 
	case BitRateGigabytePerSecond:
		return ((value * 8) * 1000000000.0) 
	case BitRateTerabytePerSecond:
		return ((value * 8) * 1000000000000.0) 
	case BitRatePetabytePerSecond:
		return ((value * 8) * 1000000000000000.0) 
	case BitRateExabytePerSecond:
		return ((value * 8) * 1e+18) 
	case BitRateKibibytePerSecond:
		return ((value * 8) * 1024) 
	case BitRateMebibytePerSecond:
		return ((value * 8) * 1048576) 
	case BitRateGibibytePerSecond:
		return ((value * 8) * 1073741824) 
	case BitRateTebibytePerSecond:
		return ((value * 8) * 1099511627776) 
	case BitRatePebibytePerSecond:
		return ((value * 8) * 1125899906842624) 
	case BitRateExbibytePerSecond:
		return ((value * 8) * 1152921504606846976) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the BitRate in the default unit (BitPerSecond),
// formatted to two decimal places.
func (a BitRate) String() string {
	return a.ToString(BitRateBitPerSecond, 2)
}

// ToString formats the BitRate value as a string with the specified unit and fractional digits.
// It converts the BitRate to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the BitRate value will be converted (e.g., BitPerSecond))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the BitRate with the unit abbreviation.
func (a *BitRate) ToString(unit BitRateUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetBitRateAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetBitRateAbbreviation(unit))
}

// Equals checks if the given BitRate is equal to the current BitRate.
//
// Parameters:
//    other: The BitRate to compare against.
//
// Returns:
//    bool: Returns true if both BitRate are equal, false otherwise.
func (a *BitRate) Equals(other *BitRate) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current BitRate with another BitRate.
// It returns -1 if the current BitRate is less than the other BitRate, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The BitRate to compare against.
//
// Returns:
//    int: -1 if the current BitRate is less, 1 if greater, and 0 if equal.
func (a *BitRate) CompareTo(other *BitRate) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given BitRate to the current BitRate and returns the result.
// The result is a new BitRate instance with the sum of the values.
//
// Parameters:
//    other: The BitRate to add to the current BitRate.
//
// Returns:
//    *BitRate: A new BitRate instance representing the sum of both BitRate.
func (a *BitRate) Add(other *BitRate) *BitRate {
	return &BitRate{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given BitRate from the current BitRate and returns the result.
// The result is a new BitRate instance with the difference of the values.
//
// Parameters:
//    other: The BitRate to subtract from the current BitRate.
//
// Returns:
//    *BitRate: A new BitRate instance representing the difference of both BitRate.
func (a *BitRate) Subtract(other *BitRate) *BitRate {
	return &BitRate{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current BitRate by the given BitRate and returns the result.
// The result is a new BitRate instance with the product of the values.
//
// Parameters:
//    other: The BitRate to multiply with the current BitRate.
//
// Returns:
//    *BitRate: A new BitRate instance representing the product of both BitRate.
func (a *BitRate) Multiply(other *BitRate) *BitRate {
	return &BitRate{value: a.value * other.BaseValue()}
}

// Divide divides the current BitRate by the given BitRate and returns the result.
// The result is a new BitRate instance with the quotient of the values.
//
// Parameters:
//    other: The BitRate to divide the current BitRate by.
//
// Returns:
//    *BitRate: A new BitRate instance representing the quotient of both BitRate.
func (a *BitRate) Divide(other *BitRate) *BitRate {
	return &BitRate{value: a.value / other.BaseValue()}
}

// GetBitRateAbbreviation gets the unit abbreviation.
func GetBitRateAbbreviation(unit BitRateUnits) string {
	switch unit { 
	case BitRateBitPerSecond:
		return "bit/s" 
	case BitRateBytePerSecond:
		return "B/s" 
	case BitRateKilobitPerSecond:
		return "kbit/s" 
	case BitRateMegabitPerSecond:
		return "Mbit/s" 
	case BitRateGigabitPerSecond:
		return "Gbit/s" 
	case BitRateTerabitPerSecond:
		return "Tbit/s" 
	case BitRatePetabitPerSecond:
		return "Pbit/s" 
	case BitRateExabitPerSecond:
		return "Ebit/s" 
	case BitRateKibibitPerSecond:
		return "KiBbit/s" 
	case BitRateMebibitPerSecond:
		return "MiBbit/s" 
	case BitRateGibibitPerSecond:
		return "GiBbit/s" 
	case BitRateTebibitPerSecond:
		return "TiBbit/s" 
	case BitRatePebibitPerSecond:
		return "PiBbit/s" 
	case BitRateExbibitPerSecond:
		return "EiBbit/s" 
	case BitRateKilobytePerSecond:
		return "kB/s" 
	case BitRateMegabytePerSecond:
		return "MB/s" 
	case BitRateGigabytePerSecond:
		return "GB/s" 
	case BitRateTerabytePerSecond:
		return "TB/s" 
	case BitRatePetabytePerSecond:
		return "PB/s" 
	case BitRateExabytePerSecond:
		return "EB/s" 
	case BitRateKibibytePerSecond:
		return "KiBB/s" 
	case BitRateMebibytePerSecond:
		return "MiBB/s" 
	case BitRateGibibytePerSecond:
		return "GiBB/s" 
	case BitRateTebibytePerSecond:
		return "TiBB/s" 
	case BitRatePebibytePerSecond:
		return "PiBB/s" 
	case BitRateExbibytePerSecond:
		return "EiBB/s" 
	default:
		return ""
	}
}