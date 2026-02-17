/**
 * Copyright TechAngle 2026. All rights reserved.
 * Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
 *
 * Author: https://codeberg.com/TechAngle
 */

import dayjs from "dayjs";
import utc from "dayjs/plugin/utc";

dayjs.extend(utc);

function isDurationHours(seconds: number): boolean {
  return seconds / 60 > 60;
}

const secondsToDuration = (seconds: number): string => {
  let format = isDurationHours(seconds) ? "HH:mm:ss" : "mm:ss";

  return dayjs.unix(seconds).utc().format(format);
};

export { secondsToDuration };
