import { env } from '$env/dynamic/public';

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
	newWs.addEventListener('close', (event) => {
		console.log('[websocket] connection closed', event);
	});
	newWs.addEventListener('message', (event) => {
		console.log('[websocket] message received', event);
	});
	return newWs;
}

export function sendJoin({ ws, room, user }: { ws: WebSocket; room: string; user: string }) {
	ws.send(JSON.stringify({ type: 'join', content: room, user }));
}

export function sendAnswer({ ws, answer, user }: { ws: WebSocket; answer: string; user: string }) {
	ws.send(JSON.stringify({ type: 'attend', content: answer, user }));
}

export function sendLeave({ ws, user }: { ws: WebSocket; user: string }) {
	ws.send(JSON.stringify({ type: 'leave', user }));
}
