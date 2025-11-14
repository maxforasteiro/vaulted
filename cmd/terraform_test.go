// Copyright 2018 SumUp Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bytes"
	"testing"

	"github.com/maxforasteiro/vaulted/pkg/aes"
	"github.com/maxforasteiro/vaulted/pkg/base64"
	"github.com/maxforasteiro/vaulted/pkg/hcl"
	"github.com/maxforasteiro/vaulted/pkg/os/ostest"
	"github.com/maxforasteiro/vaulted/pkg/pkcs7"
	"github.com/maxforasteiro/vaulted/pkg/rsa"
	"github.com/maxforasteiro/vaulted/pkg/testutils"
	"github.com/stretchr/testify/assert"
)

func TestNewTerraformCmd(t *testing.T) {
	t.Parallel()

	osExecutor := ostest.NewFakeOsExecutor(t)
	b64Svc := base64.NewBase64Service()
	rsaSvc := rsa.NewRsaService(osExecutor)
	aesSvc := aes.NewAesService(pkcs7.NewPkcs7Service())
	hclSvc := hcl.NewHclService()

	actual := NewTerraformCmd(
		osExecutor,
		rsaSvc,
		hclSvc,
		b64Svc,
		aesSvc,
	)

	assert.Equal(t, "terraform", actual.Use)
	assert.Equal(t, "Terraform resources related commands", actual.Short)
	assert.Equal(t, "Terraform resources related commands", actual.Long)
}

func TestTerraformCmd_Execute(t *testing.T) {
	t.Parallel()

	outputBuff := &bytes.Buffer{}

	osExecutor := ostest.NewFakeOsExecutor(t)
	b64Svc := base64.NewBase64Service()
	rsaSvc := rsa.NewRsaService(osExecutor)
	aesSvc := aes.NewAesService(pkcs7.NewPkcs7Service())
	hclSvc := hcl.NewHclService()

	cmdInstance := NewTerraformCmd(
		osExecutor,
		rsaSvc,
		hclSvc,
		b64Svc,
		aesSvc,
	)

	_, err := testutils.RunCommandInSameProcess(
		cmdInstance,
		[]string{},
		outputBuff,
	)

	assert.Equal(
		t,
		`Terraform resources related commands

Usage:
  terraform [flags]
  terraform [command]

Available Commands:
  help        Help about any command
  vault       github.com/sumup-oss/terraform-provider-vaulted resources related commands

Flags:
  -h, --help   help for terraform

Use "terraform [command] --help" for more information about a command.
`,
		outputBuff.String(),
	)
	assert.Nil(t, err)

	osExecutor.AssertExpectations(t)
}
