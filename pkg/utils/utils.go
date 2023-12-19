/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package utils

import (
	"fmt"
	"strings"
)

const (
	DefaultModelPathSeparator = "|"
	ArrayFieldMarker          = "[]"
	AllMapKeysFieldMarker     = "*"
)

func BuildModelPath(modelPathSeparator string, paths ...string) string {
	return strings.Join(paths, modelPathSeparator)
}

func BuildDefaultModelPath(paths ...string) string {
	return BuildModelPath(DefaultModelPathSeparator, paths...)
}

func BuildArrayField(field string) string {
	return fmt.Sprintf("%s%s", field, ArrayFieldMarker)
}
