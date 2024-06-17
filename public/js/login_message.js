document.getElementById('respond').addEventListener('submit', function(event) {
    event.preventDefault(); // Предотвращаем отправку формы по умолчанию

    var formData = new FormData(this);

    fetch('http://localhost:8080/order_page', {
        method: 'POST',
        body: formData
    })
        .then(response => response.json())
        .then(data => {
            // Обработка данных JSON
            console.log(data.message);
            // Вывод сообщения во всплывающем окне
            alert(data.message);
        })
        .catch(error => {
            console.error('Ошибка:', error);
        });
});