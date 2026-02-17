// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://codeberg.com/TechAngle

package server

import "testing"

func TestPortValidation(t *testing.T) {
	invalidPort := 24
	validPort := 443

	if isPortValid(invalidPort) {
		t.Fatalf("invalid port is valid! (%d is %v)", invalidPort, isPortValid(invalidPort))
	}

	if !isPortValid(validPort) {
		t.Fatalf("valid port is invalid (%d is %v)", validPort, isPortValid(validPort))
	}
}

func TestSoundсloudValidLinkWithoutArguments(t *testing.T) {
	link := "https://soundcloud.com/xtrullor/paracosm"
	isValid := isSoundcloudLink(link)

	if !isValid {
		t.Fatalf("Valid link without arguments was defined as invalid! (%s != %v)", link, isValid)
	}
}

func TestSoundcloudValidLinkWithArguments(t *testing.T) {
	link := "https://soundcloud.com/xtrullor/paracosm?si=d45a2b0bec464eb0a7fcb2b15e55985e&utm_source=clipboard&utm_medium=text&utm_campaign=social_sharing"
	isValid := isSoundcloudLink(link)

	if !isValid {
		t.Fatalf("Valid link with arguments was defined as invalid! (%s != %v)", link, isValid)
	}
}

func TestSoundcloudValidMobileLink(t *testing.T) {
	link := "https://on.soundcloud.com/2hvQCqRgHjJ5UveMLY"
	isValid := isSoundcloudLink(link)

	if !isValid {
		t.Fatalf("Valid mobile link was defined as invalid! (%s != %v)", link, isValid)
	}
}

func TestSoundcloudInvalidLink(t *testing.T) {
	link := "https://soundcloud.org/xtrullor/paracosm" // soundcloud.com -> .org
	isValid := isSoundcloudLink(link)

	if isValid {
		t.Fatalf("Invalid link was defined as valid!")
	}
}
