window.addEventListener('load', function() {
    const requests = JSON.parse(localStorage.getItem('requests')) || [];
    const requestsContainer = document.querySelector('nav ul');  // Контейнер для заявок

    // Логируем содержимое localStorage
    console.log("Загруженные заявки:", requests);


    if (requests.length === 0) {
        console.log("Нет заявок для отображения");
    }

    // Добавляем заявки в DOM
    requests.forEach(function(request) {
        const li = document.createElement('li');
        li.classList.add('request-item'); // Добавляем класс для стилизации

        li.innerHTML = `
            <a href="game.html">
                <div class="request-block">
                    <p><strong>Рейтинг:</strong> ${request.rating}</p>
                    <p><strong>Возраст:</strong> ${request.age}</p>
                    <p><strong>Пол:</strong> ${request.gender}</p>
                    <p><strong>Прайм-тайм:</strong> ${request.primeTime}</p>
                    <p><strong>Цель:</strong> ${request.goal}</p>
                    <p><strong>Контакты:</strong> ${request.contact}</p>
                </div>
            </a>
        `;

        requestsContainer.appendChild(li);  // Добавляем заявку в контейнер
    });
});
