/**
 * Copyright TechAngle 2026. All rights reserved.
 * Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
 *
 * Author: https://codeberg.com/TechAngle
 */

import type { Alpine as AlpineType } from "alpinejs";

declare global {
  interface Window {
    Alpine: AlpineType;
    secondsToDuration: (v: number) => string;
  }
}
