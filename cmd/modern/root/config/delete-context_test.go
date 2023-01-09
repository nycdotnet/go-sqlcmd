// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	"github.com/microsoft/go-sqlcmd/internal/cmdparser"
	"github.com/microsoft/go-sqlcmd/internal/test"
	"testing"
)

func TestDeleteContext(t *testing.T) {
	cmdparser.TestSetup(t)
	cmdparser.TestCmd[*AddUser]("--username user --auth-type other")
	cmdparser.TestCmd[*AddEndpoint]()
	cmdparser.TestCmd[*AddContext]("--endpoint endpoint --user user")
	cmdparser.TestCmd[*DeleteContext]("--name context")
}

func TestNegDeleteContext(t *testing.T) {
	defer func() { test.CatchExpectedError(recover(), t) }()

	cmdparser.TestSetup(t)
	cmdparser.TestCmd[*DeleteContext]()
}

func TestNegDeleteContext2(t *testing.T) {
	defer func() { test.CatchExpectedError(recover(), t) }()

	cmdparser.TestSetup(t)
	cmdparser.TestCmd[*DeleteContext]("--name does-not-exist")
}