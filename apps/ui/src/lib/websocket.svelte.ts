import { env } from '$env/dynamic/public';

let webSocketEstablished = false;
let ws = $state<WebSocket | null>(null);

let wsMessages = $state<string[]>(['test']);
let showNotification = $state(false);
let notification = $state<string>('');

export function getWs() {
	return ws;
}

export function setWs() {
	ws = establishWs(webSocketEstablished);
}

export function establishWs(webSocketEstablished: boolean) {
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
		if (data.message === '' || data.message === undefined) {
			return;
		}
		wsMessages = [data.message, ...wsMessages];
		notification = data.message;
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
	ws.send(JSON.stringify({ type: 'attempt', content: answer, user }));
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
