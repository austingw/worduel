<script lang="ts">
	import { PUBLIC_API_URL } from '$env/static/public';
	import establishWebSocket from '../lib/establishWebSocket.ts';
	import Keyboard from '../lib/components/Keyboard.svelte';
	let webSocketEstablished = false;
	let ws: WebSocket | null = null;
	let letters: string[] = $state([]);
	let answer: string = $derived(letters.join(''));

	function addLetter(letter: string) {
		letters.push(letter);
	}

	function removeLetter(letter: string) {
		letters.pop(letter);
	}

	function submitAnswer() {
		ws.send(answer);
	}

	$effect(() => {
		ws = establishWebSocket(webSocketEstablished, ws);
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
	<Keyboard {addLetter} {removeLetter} {submitAnswer} />
</div>
