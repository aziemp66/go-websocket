let socket = null

console.log('Loaded chat.js');

document.addEventListener('DOMContentLoaded', () => {
	socket = new WebSocket('ws://127.0.0.1:8080/websocket')

	socket.onopen = () => {
		console.log('Connected to websocket')
	}

	socket.onclose = () => {
		console.log('Disconnected from websocket')
	}

	socket.onerror = (error) => {
		console.log('Error: ', error)
	}

	socket.onmessage = (msg) => {
		console.log('Message: ', msg)

		const data = JSON.parse(msg.data)

		console.log('Data: ', data);
	}
})