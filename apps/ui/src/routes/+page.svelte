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
		ws?.send(answer);
		attempts.push(answer);
		letters = ['', '', '', '', ''];
		cursor = 0;
	}

	console.log('state', letters, cursor, answer, attempts);

	$effect(() => {
		ws = establishWebSocket(webSocketEstablished);
		return () => {
			if (ws) {
				ws.close();
			}
		};
	});
</script>

<div class="w-full flex flex-col items-center justify-center">
	<h1 class="text-3xl font-bold underline pt-20">🖋️Worduel🤺</h1>
	<h2 class="text-2xl font-bold underline pt-20">Attempts</h2>
	<Grid {letters} {attempts} />
	<ul>
		{#each attempts as attempt}
			<li>{attempt}</li>
		{/each}
	</ul>
	<button class="btn btn-primary" onclick={() => ws?.send('Testing')}>Click me</button>
	<Keyboard {addLetter} {removeLetter} {submitAnswer} />
</div>
