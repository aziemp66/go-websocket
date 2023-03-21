let socket = null

console.log('Loaded chat.js');

document.addEventListener('DOMContentLoaded', () => {
	socket = new WebSocket('ws://127.0.0.1:8080/websocket')

	socket.onopen = () => {
		console.log('Connected to websocket')
	}


})