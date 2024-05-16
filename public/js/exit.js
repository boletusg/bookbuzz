const exitBtn = document.getElementById("exitBtn");
exitBtn.addEventListener("click", function() {
    // Отправляем GET-запрос на сервер для перехода на страницу
    window.location.href = "/login";

});