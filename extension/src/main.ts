/**
 * Copyright TechAngle 2026. All rights reserved.
 * Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
 *
 * Author: https://codeberg.com/TechAngle
 */

// alpine module for window
import "./lib/alpinejs.mjs";

import { Track } from "./types/track";
import { TrackposterClient } from "./client/client";
import { secondsToDuration } from "./utils/time_utils";
import trackInfoComponent from "./components/trackInfoComponent";
import { TrackInfo } from "./types/trackInfo";
import { HOST_PING_DELAY } from "./config/config";

const client = new TrackposterClient();

// checks connection and rewrites it state in store
async function checkConnection() {
  let serverAlive = await client.isServerAlive();
  window.Alpine.store("serverOk", serverAlive);
}

// updates tracks list state
async function updateTracksList() {
  try {
    const list = await client.getQueue();
    if (list == null) {
      return;
    }

    window.Alpine.store("tracksList", list);
  } catch (err) {
    console.error(`Failed to update tracks list: ${err}`);
  }
}

document.addEventListener("alpine:init", () => {
  // initial values
  window.Alpine.store("tracksList", []);
  window.Alpine.store("serverOk", false);
  window.Alpine.store("currentMenu", "");

  // modal
  window.Alpine.store("modal", {
    secondsToDuration: secondsToDuration,
    setCurrentMenu: (menu: string) => {
      window.Alpine.store("currentMenu", menu);
    },
    closeMenu: () => window.Alpine.store("currentMenu", ""),
  });

  // checking server connection
  checkConnection();

  updateTracksList();
});

window.Alpine.start();

// set server update interval
setInterval(checkConnection, HOST_PING_DELAY);
