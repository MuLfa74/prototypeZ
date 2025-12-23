document.querySelector('form').addEventListener('submit', function(event) {
    event.preventDefault();

    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;

    if (username && password) {
        alert('Вход выполнен успешно!');
        window.location.href = 'index.html';  // Переход в профиль
    } else {
        alert('Пожалуйста, заполните все поля.');
    }
});
