/**
 * Copyright TechAngle 2026. All rights reserved.
 * Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
 *
 * Author: https://github.com/TechAngle
 */

const bodyElement = document.body;
if (bodyElement == null) {
  throw new Error("cannot find body element");
}

const SOUND_ACTIONS_QUERY_SELECTOR = ".soundActions > .sc-button-group";

const BUTTON_CLASS = "tp__button";
const BUTTON_TEXT = "[TP] Save To Queue";

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
 * Add button if not exists on current elements
 */
function addButton(target: Element) {
  if (target == null) return;
  if (hasTpButton(target)) return;

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

  target.appendChild(tpButton);

  console.log("Added button");
}

// starting observing
mutationObserver.observe(bodyElement, mutationConfig);

// calling first update
document.querySelectorAll(SOUND_ACTIONS_QUERY_SELECTOR).forEach(addButton);
