function shortenURL() {
    const urlInput = document.getElementById("url-input").value;
    const resultDiv = document.getElementById("result");

    fetch("/shorten", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({ url: urlInput })
    })
    .then(response => response.json())
    .then(data => {
        if (data.short_url) {
            resultDiv.innerHTML = `Shortened URL: <a href="${data.short_url}" target="_blank">${data.short_url}</a>`;
        } else {
            resultDiv.innerHTML = `Error: ${data.error}`;
        }
    })
    .catch(error => {
        resultDiv.innerHTML = `Error: ${error.message}`;
    });
}
