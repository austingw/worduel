<script lang="ts">
	import Grid from '$lib/components/Grid.svelte';
	import Notifications from '$lib/components/Notifications.svelte';
	import Keyboard from '$lib/components/Keyboard.svelte';

	import { establishWebSocket } from '$lib/websocket.svelte';

	let webSocketEstablished = false;
	let ws = $state<WebSocket | null>(null);
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
		ws?.send(
			JSON.stringify({
				type: 'join',
				content: 'fef',
				user: {
					name: 'jeff',
					score: 1
				}
			})
		);
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

<Grid {letters} {attempts} {answer} />
<Keyboard {addLetter} {removeLetter} {submitAnswer} />
<Notifications />
