const sidebar = document.getElementById("sidebar");
const overlay = document.getElementById("overlay");
const openBtn = document.getElementById("openSidebar");
const closeBtn = document.getElementById("closeSidebar");
let lastTarget = document.getElementById("start-link");

const focusedClassName = "flex items-center gap-3 rounded-lg px-3 py-2 bg-gray-800 text-white font-medium cursor-pointer"
const unfocusedClassName = "flex items-center gap-3 rounded-lg px-3 py-2 text-gray-300 hover:bg-gray-800 hover:text-white cursor-pointer"

const focusedSpanClassName = "h-2 w-2 rounded-full bg-purple-400"
const unfocusedSpanClassName =  "h-2 w-2 rounded-full bg-gray-500"

function openSidebar() {
  sidebar.classList.remove("-translate-x-full");
  overlay.classList.remove("hidden");
}

function closeSidebar() {
  sidebar.classList.add("-translate-x-full");
  overlay.classList.add("hidden");
}

openBtn?.addEventListener("click", openSidebar);
closeBtn?.addEventListener("click", closeSidebar);
overlay?.addEventListener("click", closeSidebar);

window.addEventListener("resize", () => {
  if (window.innerWidth >= 768) closeSidebar();
});

sidebar.addEventListener("click", (event) => {
  const topicLink = event.target.closest("a");
  if (!topicLink) return;

  lastTarget?.setAttribute("class", unfocusedClassName);
  topicLink.setAttribute("class", focusedClassName);

  const lastSpan = lastTarget?.querySelector("span");
  if (lastSpan) lastSpan.setAttribute("class", unfocusedSpanClassName);

  const currentSpan = topicLink.querySelector("span");
  if (currentSpan) currentSpan.setAttribute("class", focusedSpanClassName);
  
  lastTarget = topicLink;
});