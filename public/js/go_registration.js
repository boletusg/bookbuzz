const registerBtn = document.getElementById("registerBtn");
registerBtn.addEventListener("click", function() {
    // Отправляем GET-запрос на сервер для перехода на страницу
    window.location.href = "/registration_page";
});