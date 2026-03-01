/**
 * Copyright TechAngle 2026. All rights reserved.
 * Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
 *
 * Author: https://github.com/TechAngle
 */

import dayjs from "dayjs";
import utc from "dayjs/plugin/utc";

dayjs.extend(utc);

const secondsToDuration = (seconds: number): string => {
  if (!seconds && seconds !== 0) return "00:00";

  const isHours = seconds >= 3600;
  let format = isHours ? "HH:mm:ss" : "mm:ss";

  return dayjs.unix(seconds).utc().format(format);
};

export { secondsToDuration };
