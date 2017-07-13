// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

/*
Package agilent33220 implements the IVI driver for the Agilent 33220A function
generator.

State Caching: Not implemented
*/
package agilent33220

import "github.com/gotmc/ivi"

// Required to implement the Inherent Capabilities & Attributes
const (
	classSpecMajorVersion = 4
	classSpecMinorVersion = 3
	classSpecRevision     = "5.2"
	groupCapabilities     = "FgenBase,FgenStdfunc"
)

var supportedInstrumentModels = []string{
	"33220A",
	"33210A",
}

// Constants used for FGen
const (
	outputCount = 1
)

// Agilent33220 provides the IVI driver for an Agilent 33220A or 33210A
// function generator.
type Agilent33220 struct {
	inst        ivi.Instrument
	outputCount int
	Channels    []Channel
	ivi.Inherent
}

// New creates a new Agilent33220 IVI Instrument.
func New(inst ivi.Instrument, reset bool) (*Agilent33220, error) {
	ch := Channel{
		id:   0,
		inst: inst,
	}
	channels := make([]Channel, outputCount)
	channels[0] = ch
	inherentBase := ivi.InherentBase{
		ClassSpecMajorVersion:     classSpecMajorVersion,
		ClassSpecMinorVersion:     classSpecMinorVersion,
		ClassSpecRevision:         classSpecRevision,
		GroupCapabilities:         groupCapabilities,
		SupportedInstrumentModels: supportedInstrumentModels,
	}
	inherent := ivi.NewInherent(inst, inherentBase)
	fgen := Agilent33220{
		inst:        inst,
		outputCount: outputCount,
		Channels:    channels,
		Inherent:    inherent,
	}
	if reset {
		err := fgen.Reset()
		return &fgen, err
	}
	return &fgen, nil
}

// OutputCount returns the number of available output channels. OutputCount is
// the getter for the read-only IviFgenBase Attribute Output Count described in
// Section 4.2.1 of IVI-4.3: IviFgen Class Specification.
func (fgen *Agilent33220) OutputCount() int {
	return fgen.outputCount
}
