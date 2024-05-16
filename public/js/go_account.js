document.addEventListener('DOMContentLoaded', function() {
    const accountButton = document.getElementById('account-link');
    accountButton.addEventListener('click', function(event) {
        event.preventDefault();
        window.location.href = accountButton.getAttribute('href');
    });
});