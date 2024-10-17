<script lang="ts">
	import Grid from '$lib/components/Grid.svelte';
	import Notifications from '$lib/components/Notifications.svelte';
	import Keyboard from '$lib/components/Keyboard.svelte';
	import {
		getWs,
		getCurrentRoom,
		getReturnUser,
		setReturnUser,
		setCurrentRoom,
		sendAnswer,
		sendLeave
	} from '$lib/websocket.svelte';
	import type { View } from '$lib/types';

	type RoomProps = {
		name: string;
		changeView: (newView: View) => void;
	};
	let { name, changeView }: RoomProps = $props();

	const ws = getWs();
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
		if (ws !== null) {
			sendAnswer({ ws, room: getCurrentRoom(), answer: currentAttempt, username: name });
		}
		attempts.push(currentAttempt);
		letters = [];
	}

	$effect(() => {
		if (getReturnUser()) {
			changeView('list' as View);
			setReturnUser(false);
		}
	});
</script>

<div>
	<button
		onclick={() => {
			if (ws) {
				sendLeave({ ws, room: getCurrentRoom(), username: name });
			}
			setCurrentRoom('');
			changeView('list' as View);
		}}>Leave</button
	>
	<Grid {letters} {attempts} {answer} />
	<Keyboard {addLetter} {removeLetter} {submitAnswer} />
	<Notifications />
</div>
