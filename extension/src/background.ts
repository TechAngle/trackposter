/**
 * Copyright TechAngle 2026. All rights reserved.
 * Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
 *
 * Author: https://github.com/TechAngle
 */

import { TrackposterClient } from "./client/client";

const client = new TrackposterClient();

interface TrackInfo {
  title: string;
  author: string;
  url: string;
}

const bodyElement = document.body;
if (bodyElement == null) {
  throw new Error("cannot find body element");
}

const SOUND_ACTIONS_QUERY_SELECTOR = ".soundActions > .sc-button-group";

const BUTTON_CLASS = "tp__button";
const BUTTON_TEXT = "[TP] Add To Queue";

const mutationConfig: MutationObserverInit = {
  childList: true,
  subtree: true,
};
const mutationObserver = new MutationObserver(mutationCallback);

/**
 * Callback for mutations
 *
 * @param mutations Mutations list
 */
function mutationCallback(mutations: MutationRecord[]) {
  for (const mutation of mutations) {
    // checking for every added nodes
    if (mutation.addedNodes.length === 0) return;

    mutation.addedNodes.forEach((el) => {
      if (el.nodeType !== Node.ELEMENT_NODE) return;

      const element = el as Element;
      if (element.matches(SOUND_ACTIONS_QUERY_SELECTOR)) addButton(element);
    });
  }
}

function hasTpButton(soundActionsNode: Element): boolean {
  return !!soundActionsNode.querySelector(`.${BUTTON_CLASS}`);
}

/**
 * Send add track request to the server
 */
async function addTrack(info: TrackInfo) {
  try {
    const id = await client.addTrack({
      trackAuthor: info.author,
      trackTitle: info.title,
      trackUrl: info.url,
    });

    // if error haven't occurred
    console.log(`Track ID: ${id}`);
  } catch (err) {
    console.error(`Failed to add track: ${err}`);
  }
}

/**
 * Find track URL for target
 */
function getTrackUrl(target: Element): string | null {
  // checking if it's full page
  if (target.closest(".l-listen-hero")) {
    return window.location.origin + window.location.pathname;
  }

  // if it is recommended
  const closestLink = target.closest("a") as HTMLAnchorElement | null;
  const trackRoot = target.closest(
    ".sound, .trackItem, .relatedList__item, .playableItem",
  );
  const link = closestLink || trackRoot?.querySelector('a[href*="/"]');

  if (link && (link as HTMLAnchorElement).href) {
    const url = new URL((link as HTMLAnchorElement).href);
    const parts = url.pathname.split("/").filter(Boolean);

    if (parts.length >= 2) {
      return url.origin + url.pathname;
    }
  }

  return null;
}

/**
 * Extract information from target
 */
function extractInformation(target: Element): TrackInfo | null {
  const contextRoot = target.closest(
    ".sound__content, .trackItem, .playableItem",
  );

  let titleEl = contextRoot?.querySelector(".soundTitle__title");
  let authorEl = contextRoot?.querySelector(".soundTitle__username");

  if (!titleEl || !authorEl) {
    titleEl = document.querySelector(
      ".soundTitle__titleHeroContainer .soundTitle__title",
    );
    authorEl = document.querySelector(
      ".soundTitle__usernameHeroContainer .soundTitle__username",
    );
  }

  if (!titleEl || !authorEl) {
    return null;
  }

  const url =
    getTrackUrl(target) || window.location.origin + window.location.pathname;

  return {
    title: titleEl.textContent?.trim() || "Unknown Title",
    author: authorEl.textContent?.trim() || "Unknown Author",
    url: url,
  };
}

/**
 * Add button if not exists on current elements
 */
function addButton(target: Element) {
  if (target == null) return;
  if (hasTpButton(target)) return;

  // get information from target
  const info = extractInformation(target);
  if (info == null) throw new Error("failed to extract information");
  console.log("Track info for button:", info);

  // creating sc-like button
  const tpButton = document.createElement("button");
  tpButton.classList.add(
    BUTTON_CLASS,
    "sc-button",
    "sc-button-secondary",
    "sc-button-medium",
    "sc-button-responsive",
  );
  tpButton.textContent = BUTTON_TEXT;
  tpButton.title = "Add track to the download queue";
  tpButton.ariaLabel = "Add track to the download queue";
  tpButton.type = "button";
  tpButton.tabIndex = 0;
  tpButton.onclick = async () => await addTrack(info);

  target.appendChild(tpButton);

  console.log("Added button");
}

// starting observing
mutationObserver.observe(bodyElement, mutationConfig);

let lastUrl = location.href;
// checking for url updates and others
setInterval(() => {
  document.querySelectorAll(SOUND_ACTIONS_QUERY_SELECTOR).forEach(addButton);
  if (location.href === lastUrl) {
    return;
  }

  lastUrl = location.href;
  setTimeout(() => {
    // removing old buttons
    document.querySelectorAll(`.${BUTTON_CLASS}`).forEach((el) => el.remove());
  }, 1000);
}, 1500);

// calling first update
document.querySelectorAll(SOUND_ACTIONS_QUERY_SELECTOR).forEach(addButton);
