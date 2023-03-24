const socket = new WebSocket('ws://127.0.0.1:8080/websocket')

let o = document.getElementById("output")
let userField = document.getElementById("username")
let messageField = document.getElementById("message")

window.onbeforeunload = () => { 
	console.log("Leaving page");
	let jsonData = {}

	jsonData.action = "left"

	socket.send(JSON.stringify(jsonData))
}

console.log('Loaded chat.js');

document.addEventListener('DOMContentLoaded', () => {
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
		// console.log('Message: ', msg)
		// const data = JSON.parse(msg.data)
		// console.log('Data: ', data);

		const data = JSON.parse(msg.data)
		console.log("Action: ", data.action);

		switch (data.action) {
			case "list_users":
				let ul = document.getElementById("online_users")
				while (ul.firstChild) {
					ul.removeChild(ul.firstChild)
				}

				if (data.connected_users.length > 0) {
					data.connected_users.forEach((user) => { 
						const li = document.createElement("li")
						li.appendChild(document.createTextNode(user))
						ul.appendChild(li)
					})
				}
			case "broadcast":
				o.innerHTML = o.innerHTML + data.message + "<br>"
		}
	}

	let usernameInput = document.getElementById("username")

	usernameInput.addEventListener("change", (e) => {
		let jsonData = {}

		jsonData.action = "username"
		jsonData.username = e.target.value
		
		socket.send(JSON.stringify(jsonData))
	})

	document.getElementById("message").addEventListener("keydown", (e) => { 
		if (e.code == "Enter" && (e.target.value != "" && userField.value != "")) {
			if (!socket) {
				console.log("No socket connection")
				return false
			}
			e.preventDefault()
			e.stopPropagation
			sendMessage()
		}
	})

	document.getElementById("sendBtn").addEventListener("click", (e) => {
		if ((userField.value === "") || (messageField.value === "")) {
			alert("No username or message")
			return false
		} else {
			sendMessage()
		}
	})
})

function sendMessage() {
	let jsonData = {}

	jsonData.action = "broadcast"
	jsonData.username = userField.value
	jsonData.message = messageField.value

	socket.send(JSON.stringify(jsonData))
}