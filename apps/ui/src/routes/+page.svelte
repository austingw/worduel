<script lang="ts">
	import { PUBLIC_API_URL } from '$env/static/public';
	import establishWebSocket from '$lib/establishWebSocket';
	import Grid from '$lib/components/Grid.svelte';
	import Keyboard from '$lib/components/Keyboard.svelte';
	let webSocketEstablished = false;
	let ws: WebSocket | null = null;
	let letters: string[] = $state(['', '', '', '', '']);
	let cursor: number = $state(0);
	let answer: string = $derived(letters.join(''));
	let attempts: string[] = $state([]);

	function addLetter(letter: string) {
		letters[cursor] = letter;
		cursor = Math.min(Math.max(0, cursor + 1), 4);
	}

	function removeLetter(letter: string) {
		letters[cursor] = '';
		if (cursor > 0) {
			cursor = Math.min(Math.max(0, cursor - 1), 4);
		}
	}

	function submitAnswer() {
		if (letters.contains('')) {
			return;
		}
		ws?.send(answer);
		attempts.push(answer);
		letters = ['', '', '', '', ''];
		cursor = 0;
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
	<Grid {letters} {attempts} />
	<Keyboard {addLetter} {removeLetter} {submitAnswer} />
</div>
