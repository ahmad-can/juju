// Copyright 2020 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package state

import (
	"github.com/juju/errors"
	"github.com/juju/mgo/v3/txn"

	stateerrors "github.com/juju/juju/state/errors"
)

var (
	newProviderIDNotUniqueError     = stateerrors.NewProviderIDNotUniqueError
	newParentDeviceHasChildrenError = stateerrors.NewParentDeviceHasChildrenError
	newErrCharmAlreadyUploaded      = stateerrors.NewErrCharmAlreadyUploaded

	IsCharmAlreadyUploadedError    = stateerrors.IsCharmAlreadyUploadedError
	IsProviderIDNotUniqueError     = stateerrors.IsProviderIDNotUniqueError
	IsParentDeviceHasChildrenError = stateerrors.IsParentDeviceHasChildrenError
	IsNotAlive                     = stateerrors.IsNotAlive
)

var (
	// State package internal errors.
	machineNotAliveErr     = stateerrors.NewNotAliveError("machine")
	applicationNotAliveErr = stateerrors.NewNotAliveError("application")
	unitNotAliveErr        = stateerrors.NewNotAliveError("unit")
	notAliveErr            = stateerrors.NewNotAliveError("")
)

func onAbort(txnErr, err error) error {
	if txnErr == txn.ErrAborted ||
		errors.Cause(txnErr) == txn.ErrAborted {
		return errors.Trace(err)
	}
	return errors.Trace(txnErr)
}
