package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// BitRateUnits enumeration
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

// BitRateDto represents an BitRate
type BitRateDto struct {
	Value float64
	Unit  BitRateUnits
}

// BitRateDtoFactory struct to group related functions
type BitRateDtoFactory struct{}

func (udf BitRateDtoFactory) FromJSON(data []byte) (*BitRateDto, error) {
	a := BitRateDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a BitRateDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// BitRate struct
type BitRate struct {
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

// BitRateFactory struct to group related functions
type BitRateFactory struct{}

func (uf BitRateFactory) CreateBitRate(value float64, unit BitRateUnits) (*BitRate, error) {
	return newBitRate(value, unit)
}

func (uf BitRateFactory) FromDto(dto BitRateDto) (*BitRate, error) {
	return newBitRate(dto.Value, dto.Unit)
}

func (uf BitRateFactory) FromDtoJSON(data []byte) (*BitRate, error) {
	unitDto, err := BitRateDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return BitRateFactory{}.FromDto(*unitDto)
}


// FromBitPerSecond creates a new BitRate instance from BitPerSecond.
func (uf BitRateFactory) FromBitsPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateBitPerSecond)
}

// FromBytePerSecond creates a new BitRate instance from BytePerSecond.
func (uf BitRateFactory) FromBytesPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateBytePerSecond)
}

// FromKilobitPerSecond creates a new BitRate instance from KilobitPerSecond.
func (uf BitRateFactory) FromKilobitsPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateKilobitPerSecond)
}

// FromMegabitPerSecond creates a new BitRate instance from MegabitPerSecond.
func (uf BitRateFactory) FromMegabitsPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateMegabitPerSecond)
}

// FromGigabitPerSecond creates a new BitRate instance from GigabitPerSecond.
func (uf BitRateFactory) FromGigabitsPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateGigabitPerSecond)
}

// FromTerabitPerSecond creates a new BitRate instance from TerabitPerSecond.
func (uf BitRateFactory) FromTerabitsPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateTerabitPerSecond)
}

// FromPetabitPerSecond creates a new BitRate instance from PetabitPerSecond.
func (uf BitRateFactory) FromPetabitsPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRatePetabitPerSecond)
}

// FromExabitPerSecond creates a new BitRate instance from ExabitPerSecond.
func (uf BitRateFactory) FromExabitsPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateExabitPerSecond)
}

// FromKibibitPerSecond creates a new BitRate instance from KibibitPerSecond.
func (uf BitRateFactory) FromKibibitsPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateKibibitPerSecond)
}

// FromMebibitPerSecond creates a new BitRate instance from MebibitPerSecond.
func (uf BitRateFactory) FromMebibitsPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateMebibitPerSecond)
}

// FromGibibitPerSecond creates a new BitRate instance from GibibitPerSecond.
func (uf BitRateFactory) FromGibibitsPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateGibibitPerSecond)
}

// FromTebibitPerSecond creates a new BitRate instance from TebibitPerSecond.
func (uf BitRateFactory) FromTebibitsPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateTebibitPerSecond)
}

// FromPebibitPerSecond creates a new BitRate instance from PebibitPerSecond.
func (uf BitRateFactory) FromPebibitsPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRatePebibitPerSecond)
}

// FromExbibitPerSecond creates a new BitRate instance from ExbibitPerSecond.
func (uf BitRateFactory) FromExbibitsPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateExbibitPerSecond)
}

// FromKilobytePerSecond creates a new BitRate instance from KilobytePerSecond.
func (uf BitRateFactory) FromKilobytesPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateKilobytePerSecond)
}

// FromMegabytePerSecond creates a new BitRate instance from MegabytePerSecond.
func (uf BitRateFactory) FromMegabytesPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateMegabytePerSecond)
}

// FromGigabytePerSecond creates a new BitRate instance from GigabytePerSecond.
func (uf BitRateFactory) FromGigabytesPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateGigabytePerSecond)
}

// FromTerabytePerSecond creates a new BitRate instance from TerabytePerSecond.
func (uf BitRateFactory) FromTerabytesPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateTerabytePerSecond)
}

// FromPetabytePerSecond creates a new BitRate instance from PetabytePerSecond.
func (uf BitRateFactory) FromPetabytesPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRatePetabytePerSecond)
}

// FromExabytePerSecond creates a new BitRate instance from ExabytePerSecond.
func (uf BitRateFactory) FromExabytesPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateExabytePerSecond)
}

// FromKibibytePerSecond creates a new BitRate instance from KibibytePerSecond.
func (uf BitRateFactory) FromKibibytesPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateKibibytePerSecond)
}

// FromMebibytePerSecond creates a new BitRate instance from MebibytePerSecond.
func (uf BitRateFactory) FromMebibytesPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateMebibytePerSecond)
}

// FromGibibytePerSecond creates a new BitRate instance from GibibytePerSecond.
func (uf BitRateFactory) FromGibibytesPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateGibibytePerSecond)
}

// FromTebibytePerSecond creates a new BitRate instance from TebibytePerSecond.
func (uf BitRateFactory) FromTebibytesPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRateTebibytePerSecond)
}

// FromPebibytePerSecond creates a new BitRate instance from PebibytePerSecond.
func (uf BitRateFactory) FromPebibytesPerSecond(value float64) (*BitRate, error) {
	return newBitRate(value, BitRatePebibytePerSecond)
}

// FromExbibytePerSecond creates a new BitRate instance from ExbibytePerSecond.
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

// BaseValue returns the base value of BitRate in BitPerSecond.
func (a *BitRate) BaseValue() float64 {
	return a.value
}


// BitPerSecond returns the value in BitPerSecond.
func (a *BitRate) BitsPerSecond() float64 {
	if a.bits_per_secondLazy != nil {
		return *a.bits_per_secondLazy
	}
	bits_per_second := a.convertFromBase(BitRateBitPerSecond)
	a.bits_per_secondLazy = &bits_per_second
	return bits_per_second
}

// BytePerSecond returns the value in BytePerSecond.
func (a *BitRate) BytesPerSecond() float64 {
	if a.bytes_per_secondLazy != nil {
		return *a.bytes_per_secondLazy
	}
	bytes_per_second := a.convertFromBase(BitRateBytePerSecond)
	a.bytes_per_secondLazy = &bytes_per_second
	return bytes_per_second
}

// KilobitPerSecond returns the value in KilobitPerSecond.
func (a *BitRate) KilobitsPerSecond() float64 {
	if a.kilobits_per_secondLazy != nil {
		return *a.kilobits_per_secondLazy
	}
	kilobits_per_second := a.convertFromBase(BitRateKilobitPerSecond)
	a.kilobits_per_secondLazy = &kilobits_per_second
	return kilobits_per_second
}

// MegabitPerSecond returns the value in MegabitPerSecond.
func (a *BitRate) MegabitsPerSecond() float64 {
	if a.megabits_per_secondLazy != nil {
		return *a.megabits_per_secondLazy
	}
	megabits_per_second := a.convertFromBase(BitRateMegabitPerSecond)
	a.megabits_per_secondLazy = &megabits_per_second
	return megabits_per_second
}

// GigabitPerSecond returns the value in GigabitPerSecond.
func (a *BitRate) GigabitsPerSecond() float64 {
	if a.gigabits_per_secondLazy != nil {
		return *a.gigabits_per_secondLazy
	}
	gigabits_per_second := a.convertFromBase(BitRateGigabitPerSecond)
	a.gigabits_per_secondLazy = &gigabits_per_second
	return gigabits_per_second
}

// TerabitPerSecond returns the value in TerabitPerSecond.
func (a *BitRate) TerabitsPerSecond() float64 {
	if a.terabits_per_secondLazy != nil {
		return *a.terabits_per_secondLazy
	}
	terabits_per_second := a.convertFromBase(BitRateTerabitPerSecond)
	a.terabits_per_secondLazy = &terabits_per_second
	return terabits_per_second
}

// PetabitPerSecond returns the value in PetabitPerSecond.
func (a *BitRate) PetabitsPerSecond() float64 {
	if a.petabits_per_secondLazy != nil {
		return *a.petabits_per_secondLazy
	}
	petabits_per_second := a.convertFromBase(BitRatePetabitPerSecond)
	a.petabits_per_secondLazy = &petabits_per_second
	return petabits_per_second
}

// ExabitPerSecond returns the value in ExabitPerSecond.
func (a *BitRate) ExabitsPerSecond() float64 {
	if a.exabits_per_secondLazy != nil {
		return *a.exabits_per_secondLazy
	}
	exabits_per_second := a.convertFromBase(BitRateExabitPerSecond)
	a.exabits_per_secondLazy = &exabits_per_second
	return exabits_per_second
}

// KibibitPerSecond returns the value in KibibitPerSecond.
func (a *BitRate) KibibitsPerSecond() float64 {
	if a.kibibits_per_secondLazy != nil {
		return *a.kibibits_per_secondLazy
	}
	kibibits_per_second := a.convertFromBase(BitRateKibibitPerSecond)
	a.kibibits_per_secondLazy = &kibibits_per_second
	return kibibits_per_second
}

// MebibitPerSecond returns the value in MebibitPerSecond.
func (a *BitRate) MebibitsPerSecond() float64 {
	if a.mebibits_per_secondLazy != nil {
		return *a.mebibits_per_secondLazy
	}
	mebibits_per_second := a.convertFromBase(BitRateMebibitPerSecond)
	a.mebibits_per_secondLazy = &mebibits_per_second
	return mebibits_per_second
}

// GibibitPerSecond returns the value in GibibitPerSecond.
func (a *BitRate) GibibitsPerSecond() float64 {
	if a.gibibits_per_secondLazy != nil {
		return *a.gibibits_per_secondLazy
	}
	gibibits_per_second := a.convertFromBase(BitRateGibibitPerSecond)
	a.gibibits_per_secondLazy = &gibibits_per_second
	return gibibits_per_second
}

// TebibitPerSecond returns the value in TebibitPerSecond.
func (a *BitRate) TebibitsPerSecond() float64 {
	if a.tebibits_per_secondLazy != nil {
		return *a.tebibits_per_secondLazy
	}
	tebibits_per_second := a.convertFromBase(BitRateTebibitPerSecond)
	a.tebibits_per_secondLazy = &tebibits_per_second
	return tebibits_per_second
}

// PebibitPerSecond returns the value in PebibitPerSecond.
func (a *BitRate) PebibitsPerSecond() float64 {
	if a.pebibits_per_secondLazy != nil {
		return *a.pebibits_per_secondLazy
	}
	pebibits_per_second := a.convertFromBase(BitRatePebibitPerSecond)
	a.pebibits_per_secondLazy = &pebibits_per_second
	return pebibits_per_second
}

// ExbibitPerSecond returns the value in ExbibitPerSecond.
func (a *BitRate) ExbibitsPerSecond() float64 {
	if a.exbibits_per_secondLazy != nil {
		return *a.exbibits_per_secondLazy
	}
	exbibits_per_second := a.convertFromBase(BitRateExbibitPerSecond)
	a.exbibits_per_secondLazy = &exbibits_per_second
	return exbibits_per_second
}

// KilobytePerSecond returns the value in KilobytePerSecond.
func (a *BitRate) KilobytesPerSecond() float64 {
	if a.kilobytes_per_secondLazy != nil {
		return *a.kilobytes_per_secondLazy
	}
	kilobytes_per_second := a.convertFromBase(BitRateKilobytePerSecond)
	a.kilobytes_per_secondLazy = &kilobytes_per_second
	return kilobytes_per_second
}

// MegabytePerSecond returns the value in MegabytePerSecond.
func (a *BitRate) MegabytesPerSecond() float64 {
	if a.megabytes_per_secondLazy != nil {
		return *a.megabytes_per_secondLazy
	}
	megabytes_per_second := a.convertFromBase(BitRateMegabytePerSecond)
	a.megabytes_per_secondLazy = &megabytes_per_second
	return megabytes_per_second
}

// GigabytePerSecond returns the value in GigabytePerSecond.
func (a *BitRate) GigabytesPerSecond() float64 {
	if a.gigabytes_per_secondLazy != nil {
		return *a.gigabytes_per_secondLazy
	}
	gigabytes_per_second := a.convertFromBase(BitRateGigabytePerSecond)
	a.gigabytes_per_secondLazy = &gigabytes_per_second
	return gigabytes_per_second
}

// TerabytePerSecond returns the value in TerabytePerSecond.
func (a *BitRate) TerabytesPerSecond() float64 {
	if a.terabytes_per_secondLazy != nil {
		return *a.terabytes_per_secondLazy
	}
	terabytes_per_second := a.convertFromBase(BitRateTerabytePerSecond)
	a.terabytes_per_secondLazy = &terabytes_per_second
	return terabytes_per_second
}

// PetabytePerSecond returns the value in PetabytePerSecond.
func (a *BitRate) PetabytesPerSecond() float64 {
	if a.petabytes_per_secondLazy != nil {
		return *a.petabytes_per_secondLazy
	}
	petabytes_per_second := a.convertFromBase(BitRatePetabytePerSecond)
	a.petabytes_per_secondLazy = &petabytes_per_second
	return petabytes_per_second
}

// ExabytePerSecond returns the value in ExabytePerSecond.
func (a *BitRate) ExabytesPerSecond() float64 {
	if a.exabytes_per_secondLazy != nil {
		return *a.exabytes_per_secondLazy
	}
	exabytes_per_second := a.convertFromBase(BitRateExabytePerSecond)
	a.exabytes_per_secondLazy = &exabytes_per_second
	return exabytes_per_second
}

// KibibytePerSecond returns the value in KibibytePerSecond.
func (a *BitRate) KibibytesPerSecond() float64 {
	if a.kibibytes_per_secondLazy != nil {
		return *a.kibibytes_per_secondLazy
	}
	kibibytes_per_second := a.convertFromBase(BitRateKibibytePerSecond)
	a.kibibytes_per_secondLazy = &kibibytes_per_second
	return kibibytes_per_second
}

// MebibytePerSecond returns the value in MebibytePerSecond.
func (a *BitRate) MebibytesPerSecond() float64 {
	if a.mebibytes_per_secondLazy != nil {
		return *a.mebibytes_per_secondLazy
	}
	mebibytes_per_second := a.convertFromBase(BitRateMebibytePerSecond)
	a.mebibytes_per_secondLazy = &mebibytes_per_second
	return mebibytes_per_second
}

// GibibytePerSecond returns the value in GibibytePerSecond.
func (a *BitRate) GibibytesPerSecond() float64 {
	if a.gibibytes_per_secondLazy != nil {
		return *a.gibibytes_per_secondLazy
	}
	gibibytes_per_second := a.convertFromBase(BitRateGibibytePerSecond)
	a.gibibytes_per_secondLazy = &gibibytes_per_second
	return gibibytes_per_second
}

// TebibytePerSecond returns the value in TebibytePerSecond.
func (a *BitRate) TebibytesPerSecond() float64 {
	if a.tebibytes_per_secondLazy != nil {
		return *a.tebibytes_per_secondLazy
	}
	tebibytes_per_second := a.convertFromBase(BitRateTebibytePerSecond)
	a.tebibytes_per_secondLazy = &tebibytes_per_second
	return tebibytes_per_second
}

// PebibytePerSecond returns the value in PebibytePerSecond.
func (a *BitRate) PebibytesPerSecond() float64 {
	if a.pebibytes_per_secondLazy != nil {
		return *a.pebibytes_per_secondLazy
	}
	pebibytes_per_second := a.convertFromBase(BitRatePebibytePerSecond)
	a.pebibytes_per_secondLazy = &pebibytes_per_second
	return pebibytes_per_second
}

// ExbibytePerSecond returns the value in ExbibytePerSecond.
func (a *BitRate) ExbibytesPerSecond() float64 {
	if a.exbibytes_per_secondLazy != nil {
		return *a.exbibytes_per_secondLazy
	}
	exbibytes_per_second := a.convertFromBase(BitRateExbibytePerSecond)
	a.exbibytes_per_secondLazy = &exbibytes_per_second
	return exbibytes_per_second
}


// ToDto creates an BitRateDto representation.
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

// ToDtoJSON creates an BitRateDto representation.
func (a *BitRate) ToDtoJSON(holdInUnit *BitRateUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts BitRate to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a BitRate) String() string {
	return a.ToString(BitRateBitPerSecond, 2)
}

// ToString formats the BitRate to string.
// fractionalDigits -1 for not mention
func (a *BitRate) ToString(unit BitRateUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *BitRate) getUnitAbbreviation(unit BitRateUnits) string {
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

// Check if the given BitRate are equals to the current BitRate
func (a *BitRate) Equals(other *BitRate) bool {
	return a.value == other.BaseValue()
}

// Check if the given BitRate are equals to the current BitRate
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

// Add the given BitRate to the current BitRate.
func (a *BitRate) Add(other *BitRate) *BitRate {
	return &BitRate{value: a.value + other.BaseValue()}
}

// Subtract the given BitRate to the current BitRate.
func (a *BitRate) Subtract(other *BitRate) *BitRate {
	return &BitRate{value: a.value - other.BaseValue()}
}

// Multiply the given BitRate to the current BitRate.
func (a *BitRate) Multiply(other *BitRate) *BitRate {
	return &BitRate{value: a.value * other.BaseValue()}
}

// Divide the given BitRate to the current BitRate.
func (a *BitRate) Divide(other *BitRate) *BitRate {
	return &BitRate{value: a.value / other.BaseValue()}
}