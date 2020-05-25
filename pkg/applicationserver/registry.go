// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package applicationserver

import (
	"context"

	pbtypes "github.com/gogo/protobuf/types"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

var (
	errDuplicateIdentifiers = errors.DefineAlreadyExists("duplicate_identifiers", "identifiers already exists")
)

// DeviceRegistry is a store for end devices.
type DeviceRegistry interface {
	// Get returns the end device by its identifiers.
	Get(ctx context.Context, ids ttnpb.EndDeviceIdentifiers, paths []string) (*ttnpb.EndDevice, error)
	// Set creates, updates or deletes the end device by its identifiers.
	Set(ctx context.Context, ids ttnpb.EndDeviceIdentifiers, paths []string, f func(*ttnpb.EndDevice) (*ttnpb.EndDevice, []string, error)) (*ttnpb.EndDevice, error)
}

type deprecatedDeviceField struct {
	Old          string
	New          string
	GetTransform func(dev *ttnpb.EndDevice)
	SetTransform func(dev *ttnpb.EndDevice, useOld, useNew bool) error
}

type deprecatedDeviceFieldMatch struct {
	deprecatedDeviceField
	MatchedOld bool
	MatchedNew bool
}

type deprecatedDeviceFieldRegistryWrapper struct {
	fields   []deprecatedDeviceField
	registry DeviceRegistry
}

func matchDeprecatedDeviceFields(paths []string, deprecated []deprecatedDeviceField) ([]string, []deprecatedDeviceFieldMatch) {
	var matched []deprecatedDeviceFieldMatch
	for _, f := range deprecated {
		hasOld, hasNew := ttnpb.HasAnyField(paths, f.Old), ttnpb.HasAnyField(paths, f.New)
		switch {
		case !hasOld && !hasNew:
			continue
		case hasOld && hasNew:
		case hasOld:
			paths = ttnpb.AddFields(paths, f.New)
		case hasNew:
			paths = ttnpb.AddFields(paths, f.Old)
		}
		matched = append(matched, deprecatedDeviceFieldMatch{
			deprecatedDeviceField: f,
			MatchedOld:            hasOld,
			MatchedNew:            hasNew,
		})
	}
	return paths, matched
}

func (w deprecatedDeviceFieldRegistryWrapper) Get(ctx context.Context, ids ttnpb.EndDeviceIdentifiers, paths []string) (*ttnpb.EndDevice, error) {
	paths, deprecated := matchDeprecatedDeviceFields(paths, w.fields)
	dev, err := w.registry.Get(ctx, ids, paths)
	if err != nil || dev == nil {
		return dev, err
	}
	for _, d := range deprecated {
		d.GetTransform(dev)
	}
	return dev, nil
}

func (w deprecatedDeviceFieldRegistryWrapper) Set(ctx context.Context, ids ttnpb.EndDeviceIdentifiers, paths []string, f func(*ttnpb.EndDevice) (*ttnpb.EndDevice, []string, error)) (*ttnpb.EndDevice, error) {
	paths, deprecated := matchDeprecatedDeviceFields(paths, w.fields)
	dev, err := w.registry.Set(ctx, ids, paths, func(dev *ttnpb.EndDevice) (*ttnpb.EndDevice, []string, error) {
		if dev != nil {
			for _, d := range deprecated {
				d.GetTransform(dev)
			}
		}
		dev, paths, err := f(dev)
		if err != nil || dev == nil {
			return dev, paths, err
		}
		for _, d := range deprecated {
			d.SetTransform(dev, d.MatchedOld, d.MatchedNew)
		}
		return dev, paths, nil
	})
	if err != nil || dev == nil {
		return dev, err
	}
	for _, d := range deprecated {
		d.GetTransform(dev)
	}
	return dev, nil
}

func wrapDeviceRegistryWithDeprecatedFields(r DeviceRegistry, fields ...deprecatedDeviceField) DeviceRegistry {
	return deprecatedDeviceFieldRegistryWrapper{
		fields:   fields,
		registry: r,
	}
}

var deprecatedDeviceFields = []deprecatedDeviceField{
	{
		Old: "skip_payload_crypto",
		New: "skip_payload_crypto_override",
		GetTransform: func(dev *ttnpb.EndDevice) {
			if dev.SkipPayloadCryptoOverride == nil && dev.SkipPayloadCrypto {
				dev.SkipPayloadCryptoOverride = &pbtypes.BoolValue{Value: true}
			} else {
				dev.SkipPayloadCrypto = dev.SkipPayloadCryptoOverride.GetValue()
			}
		},
		SetTransform: func(dev *ttnpb.EndDevice, _, _ bool) error {
			dev.SkipPayloadCrypto = dev.SkipPayloadCryptoOverride.GetValue()
			return nil
		},
	},
}

// LinkRegistry is a store for application links.
type LinkRegistry interface {
	// Get returns the link by the application identifiers.
	Get(ctx context.Context, ids ttnpb.ApplicationIdentifiers, paths []string) (*ttnpb.ApplicationLink, error)
	// Range ranges the links and calls the callback function, until false is returned.
	Range(ctx context.Context, paths []string, f func(context.Context, ttnpb.ApplicationIdentifiers, *ttnpb.ApplicationLink) bool) error
	// Set creates, updates or deletes the link by the application identifiers.
	Set(ctx context.Context, ids ttnpb.ApplicationIdentifiers, paths []string, f func(*ttnpb.ApplicationLink) (*ttnpb.ApplicationLink, []string, error)) (*ttnpb.ApplicationLink, error)
}
