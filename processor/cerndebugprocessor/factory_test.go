// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package cerndebugprocessor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFactory(t *testing.T) {
	factory := NewFactory()
	assert.EqualValues(t, "cerndebug", factory.Type())
	config := factory.CreateDefaultConfig()
	assert.NotNil(t, config)
}
