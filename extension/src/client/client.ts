/**
 * Copyright TechAngle 2026. All rights reserved.
 * Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
 *
 * Author: https://codeberg.com/TechAngle
 */

import { DEFAULT_HOST } from "../config/config";
import { TrackInfo } from "../types/trackInfo";
import { Track } from "../types/track";
import { GetQueueError, TrackAdditionError } from "../types/errors";
import { ITrackposterClient } from "../types/ITrackposterClient";

interface PingResponse {
  delta
}

/**
 * Client for communicating with server.
 *
 * Uses DEFAULT_HOST if server address was not set
 */
export class TrackposterClient implements ITrackposterClient {
  private serverAddress: string;

  constructor(serverAddress?: string) {
    this.serverAddress = serverAddress || DEFAULT_HOST;
  }

  /**
   * Get host url with endpoint
   */
  private buildHostUrl(endpoint: string): string {
    // adding '/' if endpoint don't have one in the beginning
    if (!endpoint.startsWith("/")) endpoint = `${endpoint}`;

    return `${this.serverAddress}${endpoint}`;
  }

  /**
   * Get queue from server
   */
  public async getQueue(): Promise<TrackInfo[] | null> {
    try {
      const res = await fetch(this.buildHostUrl("/api/tracks/queue"));
      if (!res.ok) {
        throw new Error(`Bad queue status (${res.status} ${res.statusText})`);
      }

      const data = await res.json();
      return data as TrackInfo[];
    } catch (err) {
      throw new GetQueueError(`Failed to get queue: ${err}`);
    }
  }

  /**
   * @param track Track info
   */
  public async addTrack(track: TrackInfo): Promise<string | null> {
    try {
      const res = await fetch(this.buildHostUrl("/api/tracks/addTrack"), {
        method: "POST",
        body: JSON.stringify(track),
        headers: {
          "Content-Type": "application/json",
        },
      });

      if (!res.ok) {
        throw new Error(
          `failed to add track (${res.status} ${res.statusText})`,
        );
      }

      const data = await res.json();
      return data.trackId ? data.trackId : null;
    } catch (err) {
      throw new TrackAdditionError(`Failed to add track: ${err}`);
    }
  }

  /**
   * Check if server online.
   *
   * @returns Is server online
   */
  public async isServerAlive(): Promise<boolean> {
    try {
      const res = await fetch(this.buildHostUrl("/api/status"));
      return res.ok;
    } catch (err) {
      console.error(`API connection bad: ${err}`);
      return false;
    }
  }

  /**
   * Get ping from server
   */
  public async ping(): Promise<number | null> {
    try {
      const timestamp = new Date().getTime() / 1000;

      const res = await fetch(this.buildHostUrl("/api/ping"));
      if (!res.ok) {
      }

      const data:  = res.json();
      return data.delta || null;
    } catch (err) {
      console.error(`Ping error: ${err}`);
      throw err;
    }
  }
}
