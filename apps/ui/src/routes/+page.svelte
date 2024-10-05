<script lang="ts">
	import { establishWebSocket } from '$lib/websocket.svelte';
	import Grid from '$lib/components/Grid.svelte';
	import Notifications from '$lib/components/Notifications.svelte';
	import Keyboard from '$lib/components/Keyboard.svelte';
	import RoomList from '$lib/components/RoomList.svelte';
	import StartScreen from '$lib/components/StartScreen.svelte';
	import type { Screen } from '$lib/types';
	import { createRoom } from '$lib/mutations.svelte';

	let webSocketEstablished = false;
	let ws = $state<WebSocket | null>(null);
	let letters = $state<string[]>([]);
	let currentAttempt = $derived<string>(letters.join(''));
	let attempts = $state<string[]>([]);
	let answer = $state<string>('tares');
	let name = $state<string>('');
	let screen = $state<Screen>('start');

	function changeScreen(newScreen: Screen) {
		screen = newScreen;
	}

	export function getName() {
		return name;
	}

	const create = createRoom();

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

<div class="w-full flex flex-col items-center justify-center gap-8">
	<StartScreen {name} {changeScreen} />
	<RoomList />

	<button
		class="btn btn-secondary"
		onclick={async () => create.mutateAsync(getName()).then((res) => console.log(res))}
		>Confirm</button
	>
	<Grid {letters} {attempts} {answer} />
	<Keyboard {addLetter} {removeLetter} {submitAnswer} />
	<Notifications />
</div>
