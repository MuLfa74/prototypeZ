// Храним сообщения для каждого пользователя
const messagesData = {
    user1: [
        { from: 'user1', text: 'Привет!' },
        { from: 'Ты', text: 'Привет, как дела?' },
    ],
    user2: [
        { from: 'user2', text: 'Как ты?' },
        { from: 'Ты', text: 'Все хорошо!' },
    ],
    user3: [
        { from: 'user3', text: 'Давно не общались!' },
        { from: 'Ты', text: 'Да, согласен!' },
    ],
    user4: [
        { from: 'user4', text: 'Привет, чем занимаешься?' },
        { from: 'Ты', text: 'Работаю!' },
    ]
};

// Функция для отрисовки сообщений
function renderMessages(user) {
    const messagesContainer = document.getElementById('messages');
    messagesContainer.innerHTML = ''; // Очищаем текущие сообщения

    // Отображаем все сообщения для выбранного пользователя
    messagesData[user].forEach(message => {
        const messageElement = document.createElement('div');
        messageElement.classList.add('message');
        messageElement.innerHTML = `
            <span class="user-name">${message.from}:</span>
            <span class="message-text">${message.text}</span>
        `;
        messagesContainer.appendChild(messageElement);
    });

    // Прокручиваем в самый низ
    messagesContainer.scrollTop = messagesContainer.scrollHeight;
}

// Слушаем клик по пользователю для переключения чата
document.querySelectorAll('.users-list ul li').forEach(userElement => {
    userElement.addEventListener('click', () => {
        const user = userElement.getAttribute('data-user');
        renderMessages(user);
    });
});

// Отправка нового сообщения
document.getElementById('send-btn').addEventListener('click', function() {
    const messageInput = document.getElementById('message-input');
    const messageText = messageInput.value.trim();

    if (messageText !== '') {
        // Добавляем новое сообщение для текущего пользователя
        const currentUser = document.querySelector('.users-list ul li.active');
        const user = currentUser ? currentUser.getAttribute('data-user') : 'user1';

        messagesData[user].push({ from: 'Ты', text: messageText });

        // Отображаем новые сообщения
        renderMessages(user);

        // Очищаем поле ввода
        messageInput.value = '';
    }
});

// Изначально отображаем чат для первого пользователя
renderMessages('user1');
