/**
 * Copyright TechAngle 2026. All rights reserved.
 * Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
 *
 * Author: https://codeberg.com/TechAngle
 */

import { TrackInfo } from "./trackInfo";
import { Track } from "./track";

export interface ITrackposterClient {
  getQueue(): Promise<TrackInfo[] | null>;
  addTrack(track: TrackInfo): Promise<string | null>;
  isServerAlive(): Promise<boolean>;
  ping(): Promise<number>;
}
