// Copyright 2025 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package objects

import (
	"testing"

	gc "gopkg.in/check.v1"
)

//go:generate go run go.uber.org/mock/mockgen -typed -package objects -destination application_service_mock_test.go github.com/juju/juju/apiserver/internal/handlers/objects ApplicationServiceGetter,ApplicationService,StateGetter,ModelState,Model

func TestPackage(t *testing.T) {
	gc.TestingT(t)
}
