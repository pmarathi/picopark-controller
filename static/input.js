// sends requests from HTML to inputHandler
function sendInput(button, action) {
    fetch("/input", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
            button: button,
            action: action
        })
    });
}

// handle button input
document.querySelectorAll("button").forEach(btn => {
    const name = btn.dataset.btn;

    btn.addEventListener("mousedown", () => sendInput(name, "press"));
    btn.addEventListener("mouseup", () => sendInput(name, "release"));
});
