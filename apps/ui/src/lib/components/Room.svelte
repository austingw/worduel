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
	import { getRoomData } from '$lib/queries.svelte';
	import { getGameStart } from '$lib/game.svelte';
	import wordList from '$lib/wordList';

	type RoomProps = {
		name: string;
		changeView: (newView: View) => void;
	};
	let { name, changeView }: RoomProps = $props();

	const ws = getWs();
	let letters = $state<string[]>([]);
	let currentAttempt = $derived<string>(letters.join(''));
	let attempts = $state<string[]>([]);
	let alert = $state<string>('');
	const query = getRoomData(getCurrentRoom());
	const answer = $derived<string>(query?.data?.currentWord || '');

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
			alert = 'Word must be 5 letters long';
			setTimeout(() => {
				alert = '';
			}, 3000);
			return;
		}
		if (!wordList.includes(currentAttempt)) {
			alert = 'Word not in list';
			setTimeout(() => {
				alert = '';
			}, 3000);
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

		if (getGameStart()) {
			query.refetch();
			attempts = [];
			letters = [];
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
	<h1>Room: {getCurrentRoom()}</h1>
	<h2>Scores: {query?.data?.users[0].score ?? 0} - {query?.data?.users[1].score ?? 1}</h2>
	<Grid {letters} {attempts} {answer} />
	{#if alert}
		<div
			role="alert"
			class="toast toast-top toast-center bg-warning absolute top-8 w-fit h-fit animate-none"
		>
			<p>{alert}</p>
		</div>
	{/if}
	<Keyboard {addLetter} {removeLetter} {submitAnswer} />
	<Notifications />
</div>
