/**
 * Copyright TechAngle 2026. All rights reserved.
 * Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
 *
 * Author: https://github.com/TechAngle
 */

// @ts-ignore: fuck it
import Alpine from "@alpinejs/csp";

declare global {
  interface Window {
    Alpine: typeof Alpine;
  }
}

export {};
