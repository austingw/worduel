import { env } from '$env/dynamic/public';

let wsMessages = $state<string[]>(['test']);
let showNotification = $state(false);
let notification = $derived<string>(wsMessages[0]);

export function establishWebSocket(webSocketEstablished: boolean) {
	const parsedUrl = env.PUBLIC_API_URL.split(':').slice(1).join(':');

	if (webSocketEstablished) {
		return null;
	}
	const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
	const newWs = new WebSocket(`${protocol}${parsedUrl}/ws`);

	webSocketEstablished = true;

	newWs.addEventListener('open', (event) => {
		webSocketEstablished = true;
		console.log('[websocket] connection open', event);
	});
	newWs.addEventListener('close', () => {
		sendLeave(newWs);
		wsMessages = [];
	});
	newWs.addEventListener('message', (message: MessageEvent) => {
		const data: { message: string } = JSON.parse(message.data);
		wsMessages = [data.message, ...wsMessages];
		showNotification = true;
		setTimeout(() => {
			showNotification = false;
		}, 5000);
	});
	return newWs;
}

export function sendJoin({ ws, room, user }: { ws: WebSocket; room: string; user: string }) {
	ws.send(JSON.stringify({ type: 'join', content: room, user }));
}

export function sendAnswer({ ws, answer, user }: { ws: WebSocket; answer: string; user: string }) {
	ws.send(JSON.stringify({ type: 'attend', content: answer, user }));
}

export function sendLeave(ws: WebSocket) {
	ws.send(JSON.stringify({ type: 'leave' }));
}

export function getMessages() {
	return wsMessages;
}

export function getNotification() {
	return notification;
}

export function getShowNotification() {
	return showNotification;
}
