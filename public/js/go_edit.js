const editBtn = document.getElementById("editBtn");
editBtn.addEventListener("click", function() {
    // Отправляем GET-запрос на сервер для перехода на страницу
    window.location.href = "/account_edit";
});