// Получение значения параметра "id" из URL
const urlParams = new URLSearchParams(window.location.search);
const orderId = urlParams.get('id');

// Выполнение запроса на сервер для получения информации о заказе
fetch(`/api/order?id=${orderId}`)
    .then(response => response.json())
    .then(data => {
        // Обработка полученных данных о заказе
        // Например, можно обновить элементы страницы с информацией о заказе
        // Например, document.getElementById('order-title').textContent = data.title;
    })
    .catch(error => {
        console.error('Ошибка при выполнении запроса:', error);
    });
