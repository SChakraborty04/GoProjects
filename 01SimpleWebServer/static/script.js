function showGores(){
    if (window.location.pathname === '/success') {
        document.getElementById('gores').style.display = 'block';
    }
}
window.onload=showGores