document.addEventListener('DOMContentLoaded', function() {
    const respondButton = document.getElementById('responds-link');
    respondButton.addEventListener('click', function(event) {
        event.preventDefault();
        window.location.href = respondButton.getAttribute('href');
    });
});