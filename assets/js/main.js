// Fetch version

fetch("https://gov.theprocedural.com/version")
    .then((response) => response.text())
    .then((version) => {
        const versionElement = document.getElementById("version");
        versionElement.textContent = "v" + version;
    })
    .catch((error) => console.error("Error fetching version:", error));

// Fetch and render README.md
fetch("/README.md")
    .then((response) => response.text())
    .then((markdownText) => {
        const readmeContent = document.getElementById("readmeContent");
        readmeContent.innerHTML = marked.parse(markdownText);

        // Apply Tailwind CSS classes to Markdown elements

        // Headings
        readmeContent.querySelectorAll("h1").forEach((element) => {
            element.classList.add(
                "font-bold",
                "border-b",
                "text-2xl",
                "my-6",
                "pb-6"
            );
        });
        readmeContent.querySelectorAll("h2").forEach((element) => {
            element.classList.add(
                "border-gray-300",
                "border-b",
                "text-2xl",
                "mt-6",
                "mb-4",
                "pb-2"
            );
        });
        readmeContent.querySelectorAll("h3").forEach((element) => {
            element.classList.add("text-lg", "font-bold", "mt-3", "mb-1");
        });
        readmeContent
            .querySelectorAll("h1, h2, h3, h4, h5, h6")
            .forEach((element) => {
                element.id = element.textContent
                    .toLowerCase()
                    .replace(/[^\w]+/g, "-");
            });

        // Links
        readmeContent.querySelectorAll("a").forEach((element) => {
            element.classList.add("text-blue-500", "hover:underline");
        });

        // Paragraphs
        readmeContent.querySelectorAll("p").forEach((element) => {
            element.classList.add("mb-4");
        });

        // Code blocks
        readmeContent.querySelectorAll("pre").forEach((element) => {
            element.classList.add(
                "mb-4",
                "p-4",
                "!bg-gray-200",
                "dark:!bg-gray-900",
                "dark:text-gray-100",
                "rounded-lg",
                "language-bash",
                "relative"
            );
        });
        readmeContent.querySelectorAll("code").forEach((element) => {
            element.classList.add("text-sm", "font-mono", "dark:text-gray-100");
        });

        // Lists
        readmeContent.querySelectorAll("ul").forEach((element) => {
            element.classList.add("list-disc", "list-inside", "ml-5");
        });
        readmeContent.querySelectorAll("ol").forEach((element) => {
            element.classList.add("list-decimal", "pl-5");
            element.style.marginLeft = "1.25rem";
        });
        readmeContent.querySelectorAll("li").forEach((element) => {
            element.classList.add("mb-2");
        });

        addCopyButtons();

        // Smooth scrolling
        const tocLinks = readmeContent.querySelectorAll('a[href^="#"]');
        tocLinks.forEach((link) => {
            link.addEventListener("click", async function (e) {
                e.preventDefault();

                const targetId = this.getAttribute("href");
                const targetElement = document.querySelector(targetId);
                const targetOffset = targetElement.getBoundingClientRect().top;
                const initialOffset = window.scrollY;
                const duration = 500;

                function smoothScroll(timestamp) {
                    const elapsed = timestamp - startTime;
                    const progress = Math.min(elapsed / duration, 1);
                    const ease = (progress) =>
                        0.5 - Math.cos(progress * Math.PI) / 2;
                    const scrollTop =
                        initialOffset + (targetOffset - 75) * ease(progress);

                    window.scrollTo({
                        top: scrollTop,
                        behavior: "smooth",
                    });

                    if (elapsed < duration) {
                        requestAnimationFrame(smoothScroll);
                    }
                }

                const startTime = performance.now();
                requestAnimationFrame(smoothScroll);

                // I cannot use this code below because scrollIntoView() doesn't support offset
                // const targetId = this.getAttribute("href");
                // const targetElement = document.querySelector(targetId);

                // if (targetElement) {
                //     await targetElement.scrollIntoView({
                //         behavior: "smooth",
                //     });

                //     window.scrollBy(0, -75);
                // }
            });
        });
    })
    .catch((error) => console.error("Error fetching README.md:", error));

function addCopyButtons() {
    const codeBlocks = document.querySelectorAll("pre code");
    codeBlocks.forEach((codeBlock, index) => {
        const button = document.createElement("button");
        button.className =
            "copy-button absolute top-0 right-0 mt-3 mr-2 bg-white dark:bg-black text-gray-600 dark:text-gray-300 border border-gray-300 dark:border-gray-700 py-1 px-2 rounded-lg text-sm cursor-pointer";
        button.innerHTML = '<i class="fa-regular fa-copy"></i>';
        button.addEventListener("click", () => {
            const textArea = document.createElement("textarea");
            textArea.value = codeBlock.textContent;
            document.body.appendChild(textArea);
            textArea.select();
            navigator.clipboard.writeText(codeBlock.textContent);
            document.body.removeChild(textArea);
            button.textContent = "Copied!";
            setTimeout(() => {
                button.innerHTML = '<i class="fa-regular fa-copy"></i>';
            }, 1000);
        });
        codeBlock.parentNode.classList.add("relative");
        codeBlock.parentNode.appendChild(button);
    });
}

// Dark mode toggle button
const darkModeToggle = document.getElementById("darkModeToggle");
darkModeToggle.addEventListener("click", toggleDarkMode);

function toggleDarkMode() {
    const isDarkMode = document.body.classList.toggle("dark");
    if (isDarkMode) {
        darkModeToggle.innerHTML = '<i class="fa-regular fa-sun"></i>';
        localStorage.setItem("darkMode", "enabled");
    } else {
        darkModeToggle.innerHTML = '<i class="fa-regular fa-moon"></i>';
        localStorage.setItem("darkMode", "disabled");
    }
}

// Check local storage for dark mode preference
const savedDarkMode = localStorage.getItem("darkMode");
if (savedDarkMode === "enabled") {
    document.body.classList.add("dark");
    darkModeToggle.innerHTML = '<i class="fa-regular fa-sun"></i>';
}

// Check for system dark mode preference
if (
    window.matchMedia &&
    window.matchMedia("(prefers-color-scheme: dark)").matches &&
    !savedDarkMode
) {
    darkModeToggle.innerHTML = '<i class="fa-regular fa-sun"></i>';
    document.body.classList.add("dark");
}

// Toggle home and docs
const home = document.getElementById("home");
const docs = document.getElementById("readmeContent");
const homeButton = document.getElementById("homeLink");
const docsButton = document.getElementById("docsLink");

homeButton.addEventListener("click", () => toggleView(true));
docsButton.addEventListener("click", () => toggleView(false));

function toggleView(isHome) {
    home.classList.toggle("hidden", !isHome);
    docs.classList.toggle("hidden", isHome);
    homeButton.classList.toggle("block", isHome);
    docsButton.classList.toggle("block", !isHome);
}
