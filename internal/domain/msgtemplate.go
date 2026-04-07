// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package domain

type MessageTemplate string

func (t MessageTemplate) String() string {
	return string(t)
}
