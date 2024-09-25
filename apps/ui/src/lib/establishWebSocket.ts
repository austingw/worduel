import { env } from '$env/dynamic/public';

export default function establishWebSocket(webSocketEstablished: boolean) {
const parsedUrl = env.PUBLIC_API_URL.split(":").slice(1).join(":");

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
