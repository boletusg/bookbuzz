document.getElementById("login").addEventListener("submit", function(event) {
    event.preventDefault(); // Предотвращаем отправку формы по умолчанию
    var xhr = new XMLHttpRequest();
    xhr.open("POST", "http://localhost:8080/login");
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.onload = function() {
        if (xhr.status === 200) {
            var response = JSON.parse(xhr.responseText);
            if (response.success) {
                // Аутентификация прошла успешно, перенаправляем на другую страницу
                window.location.href = "/home";
            } else {
                // Обработка ошибок аутентификации
                alert(response.message);
            }
        } else {
            // Обработка ошибок AJAX-запроса
            alert("Ошибка при выполнении AJAX-запроса");
        }
    };
    var formData = new FormData(document.getElementById("login"));
    xhr.send(new URLSearchParams(formData));
});