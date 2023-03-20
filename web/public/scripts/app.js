const socket = io("/socket.io");

socket.on("connect", (text)=> {
	// const el = document.createElement("li")
	// el.innerHTML = text;

	// document.querySelector(".messages").appendChild(el);

	console.log("connected");
});

console.log(socket);