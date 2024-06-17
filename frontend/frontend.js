document.getElementById('noteForm').addEventListener('submit', async function(event) {
    event.preventDefault();

    const title = document.getElementById('title').value;
    const content = document.getElementById('content').value;

    const response = await fetch('http://localhost:8080/notes', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ title, content })
    });

    const data = await response.json();
    if (response.ok) {
        document.getElementById('message').innerText = `Note created with ID: ${data.id}`;
    } else {
        document.getElementById('message').innerText = `Error: ${data.error}`;
    }
});
