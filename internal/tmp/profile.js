document.querySelector('form').addEventListener('submit', function(event) {
    event.preventDefault();

    const age = document.getElementById('age').value;
    const gender = document.getElementById('gender').value;
    const primeTime = document.getElementById('prime-time').value;
    const contact = document.getElementById('contact').value;
    const games = document.getElementById('games').value;

    if (age && gender && primeTime && contact && games) {
        alert('Данные сохранены успешно!');
    } else {
        alert('Пожалуйста, заполните все поля.');
    }
});
