<script lang="ts">
	import establishWebSocket from '$lib/establishWebSocket';
	import Grid from '$lib/components/Grid.svelte';
	import Keyboard from '$lib/components/Keyboard.svelte';
	import { json } from '@sveltejs/kit';
	let webSocketEstablished = false;
	let ws = $state<WebSocket | null>(null);
	let letters = $state<string[]>([]);
	let currentAttempt = $derived<string>(letters.join(''));
	let attempts = $state<string[]>([]);
	let answer = $state<string>('tares');
	let name = $state<string>('');

	function addLetter(letter: string) {
		if (letters.length < 5) {
			letters.push(letter);
		}
	}

	function removeLetter() {
		if (letters.length > 0) {
			letters.pop();
		}
	}

	function submitAnswer() {
		if (currentAttempt.length < 5) {
			return;
		}
		ws?.send(JSON.stringify({ type: 'answer', content: currentAttempt }));
		attempts.push(currentAttempt);
		letters = [];
	}

	function sendName() {
		if (name.length > 0) {
			ws?.send(JSON.stringify({ type: 'name', content: name }));
		}
	}

	$effect(() => {
		ws = establishWebSocket(webSocketEstablished);
		return () => {
			if (ws) {
				ws.close();
			}
		};
	});
</script>

<div class="w-full flex flex-col items-center justify-center gap-8">
	<h1 class="text-3xl font-bold underline pt-20">🖋️Worduel🤺</h1>
	<input
		type="text"
		bind:value={name}
		placeholder="Enter your name..."
		class="input input-bordered"
	/>
	<button class="btn btn-secondary" onclick={sendName}>Confirm</button>
	<Grid {letters} {attempts} {answer} />
	<Keyboard {addLetter} {removeLetter} {submitAnswer} />
</div>
