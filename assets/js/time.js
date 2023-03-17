function getCurrentTimeString() {
    const now = new Date();
    const y = now.getFullYear();
    const m = now.getMonth();
    const d = now.getDate();
    const h = now.getHours();
    const mi = now.getMinutes();
    const s = now.getSeconds();
    return y + '-' + m + '-' + d + ' ' + h + ':' + m + ':' + s;
}
function updateCurrentTime() {
    const txtTime = document.getElementById('currentTime');
    txtTime.innerText = getCurrentTimeString()
}

setInterval(updateCurrentTime, 1000);
window.alert("Hello Welcome")