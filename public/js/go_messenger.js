// Получаем все ссылки на страницу messenger_page
var messengerLinks = document.querySelectorAll(".user_link");

// Добавляем обработчик события click для каждой ссылки
messengerLinks.forEach(function(link) {
    link.addEventListener("click", function(event) {
        event.preventDefault(); // Отменяем стандартное поведение ссылки
        window.location.href = "/messenger_page"; // Переходим на страницу messenger_page
    });
});