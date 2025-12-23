document.querySelector('form').addEventListener('submit', function(event) {
    event.preventDefault();

    const rating = document.getElementById('rating').value;
    const age = document.getElementById('age').value;
    const gender = document.getElementById('gender').value;
    const primeTime = document.getElementById('prime-time').value;
    const goal = document.getElementById('goal').value;
    const contact = document.getElementById('contact').value;

    if (rating && age && gender && primeTime && goal && contact) {
        const newRequest = { rating, age, gender, primeTime, goal, contact };

        // Получаем существующие заявки или создаем новый массив, если их нет
        let requests = JSON.parse(localStorage.getItem('requests')) || [];

        // Добавляем новую заявку в массив
        requests.push(newRequest);

        // Сохраняем обновленный массив в localStorage
        localStorage.setItem('requests', JSON.stringify(requests));

        alert('Заявка успешно создана!');
        window.location.href = 'requests.html';  // Переход на страницу с заявками
    } else {
        alert('Пожалуйста, заполните все поля.');
    }
});
