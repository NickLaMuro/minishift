/*
Copyright (C) 2016 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package addon

import (
	"bytes"
	"fmt"
	"github.com/minishift/minishift/pkg/testing/cli"
	"github.com/minishift/minishift/pkg/util/os/atexit"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func Test_addon_name_must_be_specified_for_disable_command(t *testing.T) {
	tmpMinishiftHomeDir := cli.SetupTmpMinishiftHome(t)
	origStdout, stdOutWriter, stdOutReader := cli.CaptureStdOut(t)
	defer cli.TearDown(tmpMinishiftHomeDir, origStdout)

	atexit.RegisterExitHandler(cli.CreateExitHandlerFunc(t, stdOutWriter, stdOutReader, 1, emptyDisableError))

	runDisableAddon(nil, nil)
}

func Test_unkown_name_for_disable_command_returns_error(t *testing.T) {
	tmpMinishiftHomeDir := cli.SetupTmpMinishiftHome(t)
	os.Mkdir(filepath.Join(tmpMinishiftHomeDir, "addons"), 0777)

	origStdout, stdOutWriter, stdOutReader := cli.CaptureStdOut(t)
	defer cli.TearDown(tmpMinishiftHomeDir, origStdout)

	testAddOnName := "foo"
	runDisableAddon(nil, []string{testAddOnName})

	stdOutWriter.Close()
	var buffer bytes.Buffer
	io.Copy(&buffer, stdOutReader)

	expectedOut := fmt.Sprintf(noAddOnToDisableMessage+"\n", testAddOnName)
	if expectedOut != buffer.String() {
		t.Fatalf("Expected output '%s'. Got '%s'.", expectedOut, buffer.String())
	}
}
