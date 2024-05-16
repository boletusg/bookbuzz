const cancelBtn = document.getElementById("cancelBtn");
cancelBtn.addEventListener("click", function() {
    // Отправляем GET-запрос на сервер для перехода на страницу
    window.location.href = "/account_page";
});