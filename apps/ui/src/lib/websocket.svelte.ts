import { env } from '$env/dynamic/public';
import { setGameStart } from './game.svelte';

let ws = $state<WebSocket | null>(null);

let wsMessages = $state<string[]>(['test']);
let showNotification = $state(false);
let notification = $state<string>('');
let currentRoom = $state<string>('');

export function getWs() {
	return ws;
}
export async function setWs() {
	ws = await establishWs();
	return ws;
}

export function getCurrentRoom() {
	return currentRoom;
}
export function setCurrentRoom(room: string) {
	currentRoom = room;
}

export async function establishWs(): Promise<WebSocket | null> {
	const parsedUrl = env.PUBLIC_API_URL.split(':').slice(1).join(':');

	if (getWs() !== null) {
		return null;
	}
	const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
	const newWs = new WebSocket(`${protocol}${parsedUrl}/ws`);

	return new Promise((resolve, reject) => {
		newWs.addEventListener('open', (event) => {
			console.log('[websocket] connection open', event);
			resolve(newWs);
		});

		newWs.addEventListener('close', () => {
			sendLeave({ ws: newWs, room: currentRoom, username: 'test' });
			wsMessages = [];
		});

		newWs.addEventListener('error', (error) => {
			console.error('[websocket] connection error', error);
			reject(error);
		});

		newWs.addEventListener('message', (message: MessageEvent) => {
			const data: { type: string; message: string } = JSON.parse(message.data);

			if (data.message === '' || data.message === undefined) {
				return;
			}

			if (data.type === 'start') {
				setTimeout(() => {
					setGameStart(true);
				}, 3000);
			} else if (data.type === 'end') {
				setGameStart(false);
			}

			wsMessages = [data.message, ...wsMessages];
			notification = data.message;
			showNotification = true;
			setTimeout(() => {
				showNotification = false;
			}, 5000);
		});
	});
}

export function sendJoin({
	ws,
	room,
	username
}: {
	ws: WebSocket;
	room: string;
	username: string;
}) {
	ws.send(JSON.stringify({ type: 'join', room, username }));
}

export function sendAnswer({
	ws,
	answer,
	room,
	username
}: {
	ws: WebSocket;
	answer: string;
	room: string;
	username: string;
}) {
	ws.send(JSON.stringify({ type: 'attempt', content: answer, room, username }));
}

export function sendLeave({
	ws,
	room,
	username
}: {
	ws: WebSocket;
	room: string;
	username: string;
}) {
	ws.send(JSON.stringify({ type: 'leave', room, username }));
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
