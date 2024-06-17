// Получаем все радиокнопки с именем "filterType"
const filterRadios = document.querySelectorAll('input[name="filterType"]');

// Получаем контейнеры для каждого типа откликов
const outgoingRespondsContainer = document.getElementById('outgoing-responds');
const incomingRespondsContainer = document.getElementById('incoming-responds');

// Добавляем обработчик событий на изменение выбранной радиокнопки
filterRadios.forEach(radio => {
    radio.addEventListener('change', () => {
        // Скрываем оба контейнера
        outgoingRespondsContainer.style.display = 'none';
        incomingRespondsContainer.style.display = 'none';

        // Показываем только соответствующий контейнер
        if (radio.value === 'outgoing') {
            outgoingRespondsContainer.style.display = 'block';
        } else if (radio.value === 'incoming') {
            incomingRespondsContainer.style.display = 'block';
        }
    });
});