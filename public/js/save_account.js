document.getElementById("accountInfo").addEventListener("submit", function(event) {
    event.preventDefault(); // Предотвращаем отправку формы по умолчанию
    // Выполняем AJAX-запрос для обработки регистрации
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "http://localhost:8080/account_edit");
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.onload = function() {
        if (xhr.status === 200) {
            const response = JSON.parse(xhr.responseText);
            if (response.success) {
                // изменения успешно внесены, перенаправляем на другую страницу
                window.location.href = "/account_page";
            } else {
                // Обработка ошибок регистрации
                alert(response.message);
            }
        } else {
            // Обработка ошибок AJAX-запроса
            alert("Ошибка при выполнении AJAX-запроса");
        }
    };
    const formData = new FormData(document.getElementById("accountInfo"));
    xhr.send(new URLSearchParams(formData));
});