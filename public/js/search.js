const searchInput = document.getElementById('search_input');
searchInput.addEventListener('input', handleSearch);
function handleSearch() {
    const searchTerm = searchInput.value.toLowerCase();
    const ordersContainer = document.getElementById('ordersContainer');
    const orders = ordersContainer.getElementsByClassName('order');

    for (const order of orders) {
        const titleElement = order.querySelector('.title_order');
        const authorElement = order.querySelector('.author');
        const textElement = order.querySelector('.text_order');

        const title = titleElement.textContent.toLowerCase();
        const author = authorElement.textContent.toLowerCase();
        const text = textElement.textContent.toLowerCase();

        if (title.includes(searchTerm) || author.includes(searchTerm) || text.includes(searchTerm)) {
            order.style.display = 'block';
        } else {
            order.style.display = 'none';
        }
    }
}
const searchForm = document.querySelector('.search_str');
searchForm.addEventListener('submit', (event) => {
    event.preventDefault();
});