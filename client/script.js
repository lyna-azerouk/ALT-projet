document.getElementById("loginForm").addEventListener("submit", function(event) {
    event.preventDefault();
    var formData = new FormData(this);
    var jsonData = {};
    formData.forEach(function(value, key) {
        jsonData[key] = value;
    });
    
    fetch('/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(jsonData)
    })
    .then(response => response.json())
    .then(console.log(response))
    .then(data => {
        document.getElementById("message").textContent = data.message;
    })
    .catch(error => {
        console.error('Error:', error);
    });
});
