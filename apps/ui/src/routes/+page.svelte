<script lang="ts">
	import { PUBLIC_API_URL } from '$env/static/public';
	import establishWebSocket from '$lib/establishWebSocket';
	import Grid from '$lib/components/Grid.svelte';
	import Keyboard from '$lib/components/Keyboard.svelte';
	let webSocketEstablished = false;
	let ws: WebSocket | null = null;
	let letters = $state<string[]>([]);
	let currentAttempt = $derived<string>(letters.join(''));
	let attempts = $state<string[]>([]);
	let answer = $state<string>('tares');

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
		ws?.send(currentAttempt);
		attempts.push(currentAttempt);
		letters = [];
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
	<Grid {letters} {attempts} {answer} />
	<Keyboard {addLetter} {removeLetter} {submitAnswer} />
</div>
