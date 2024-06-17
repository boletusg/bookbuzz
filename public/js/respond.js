// Получаем кнопку "Откликнуться"
const respondBtn = document.getElementById('respond');

respondBtn.addEventListener('click', function() {
    // Получаем идентификатор заказа из URL
    const orderID = new URLSearchParams(window.location.search).get('id');

    // Делаем AJAX запрос на сервер, чтобы записать отклик в базу данных
    fetch('/respond', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ orderID: orderID })
    })
        .then(response => response.json())
        .then(data => {
            // Обработка данных JSON
            console.log(data.message);

        })
        .catch(error => {
            console.error('Ошибка:', error);
        });
});