document.querySelector('form').addEventListener('submit', function(event) {
    event.preventDefault();

    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;
    const confirmPassword = document.getElementById('confirm-password').value;

    if (password !== confirmPassword) {
        alert('Пароли не совпадают!');
        return;
    }

    if (username && password) {
        alert('Регистрация успешна!');
        window.location.href = 'index.html';  // Переход на страницу входа
    } else {
        alert('Пожалуйста, заполните все поля.');
    }
});
