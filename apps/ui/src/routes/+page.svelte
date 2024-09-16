<script lang="ts">
	import { PUBLIC_API_URL } from '$env/static/public';
	import Keyboard from '../lib/components/Keyboard.svelte';
	let webSocketEstablished = false;
	let ws: WebSocket | null = null;
	let log: string[] = [];

	const logEvent = (str: string) => {
		log = [...log, str];
	};

	const establishWebSocket = () => {
		if (webSocketEstablished) return;
		const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
		ws = new WebSocket(`${protocol}//localhost:4000/ws`);
		webSocketEstablished = true;

		ws.addEventListener('open', (event) => {
			webSocketEstablished = true;
			console.log('[websocket] connection open', event);
			logEvent('[websocket] connection open');
		});
		ws.addEventListener('close', (event) => {
			console.log('[websocket] connection closed', event);
			logEvent('[websocket] connection closed');
		});
		ws.addEventListener('message', (event) => {
			console.log('[websocket] message received', event);
			logEvent(`[websocket] message received: ${event.data}`);
		});
	};

	$effect(() => {
		establishWebSocket();
		return () => {
			if (ws) {
				ws.close();
			}
		};
	});
</script>

<div class="w-full flex flex-col items-center justify-center">
	<h1 class="text-3xl font-bold underline">Worduel</h1>
	<button class="btn btn-primary" onclick={() => ws.send('Testing')}>Click me</button>
	<Keyboard />
</div>
