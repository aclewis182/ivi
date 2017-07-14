// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

/*
Package srsds345 implements the IVI driver for the Stanford Research System
DS345 function generator.

State Caching: Not implemented
*/
package srsds345

import "github.com/gotmc/ivi"

// Required to implement the Inherent Capabilities & Attributes
const (
	classSpecMajorVersion = 4
	classSpecMinorVersion = 3
	classSpecRevision     = "5.2"
	groupCapabilities     = "IviFgenBase,IviFgenStdfunc,IviFgenTrigger,IviFgenInternalTrigger,IviFgenBurst"
)

// TODO(mdr): Seems like groupCapabilities should be a []string instead of
// string

var supportedInstrumentModels = []string{
	"DS345",
}

// Constants used for FGen
const (
	outputCount = 1
)

// SRSDS345 provides the IVI driver for an Agilent 33220A or 33210A
// function generator.
type SRSDS345 struct {
	inst        ivi.Instrument
	outputCount int
	Channels    []Channel
	ivi.Inherent
}

// New creates a new SRSDS345 IVI Instrument.
func New(inst ivi.Instrument, reset bool) (*SRSDS345, error) {
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
	fgen := SRSDS345{
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

// Channel represents a repeated capability of an output channel for the
// function generator.
type Channel struct {
	id   int
	inst ivi.Instrument
}