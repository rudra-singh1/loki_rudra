// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package internal

import (
	otlpresource "go.opentelemetry.io/collector/pdata/internal/data/protogen/resource/v1"
)

type Resource struct {
	orig  *otlpresource.Resource
	state *State
}

func GetOrigResource(ms Resource) *otlpresource.Resource {
	return ms.orig
}

func GetResourceState(ms Resource) *State {
	return ms.state
}

func NewResource(orig *otlpresource.Resource, state *State) Resource {
	return Resource{orig: orig, state: state}
}

func CopyOrigResource(dest, src *otlpresource.Resource) {
	dest.Attributes = CopyOrigMap(dest.Attributes, src.Attributes)
	dest.DroppedAttributesCount = src.DroppedAttributesCount
}

func GenerateTestResource() Resource {
	orig := otlpresource.Resource{}
	state := StateMutable
	tv := NewResource(&orig, &state)
	FillTestResource(tv)
	return tv
}

func FillTestResource(tv Resource) {
	FillTestMap(NewMap(&tv.orig.Attributes, tv.state))
	tv.orig.DroppedAttributesCount = uint32(17)
}
