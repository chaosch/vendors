// Copyright 2014 Eric Zhou. All Rights Reserved.
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

package base64Captcha

import (
	"strings"
	"testing"
)

func TestGenerateCaptchaPngBase64String(t *testing.T) {

	pngString := GenerateCaptchaPngBase64String(RandomId(), PngHeight, PngHeight, DotCount, DefaultLen, MaxSkew)

	if strings.Contains(pngString, "data:image/png;base64,iVBO") {
		t.Log("test success")

	} else {
		t.Error("test failed")

	}
}

func TestVerifyCaptcha(t *testing.T) {

	digits := RandomDigits(DefaultLen)
	identifier := RandomId()
	globalStore.Set(identifier, digits)
	NewImage(identifier, digits, PngWidth, PngHeight).WriteToBase64String()
	if VerifyCaptcha(identifier, string(digits)) {
		t.Log("test success")
	} else {
		t.Log("test fail")
	}
}

func TestGenerateCaptchaPngBase64StringDefault(t *testing.T) {
	pngString := GenerateCaptchaPngBase64StringDefault(RandomId())

	if strings.Contains(pngString, "data:image/png;base64,iVBO") {
		t.Log("test success")

	} else {
		t.Error("test failed")

	}
}
