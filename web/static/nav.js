      const sidebar = document.getElementById("sidebar");
      const overlay = document.getElementById("overlay");
      const openBtn = document.getElementById("openSidebar");
      const closeBtn = document.getElementById("closeSidebar");

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