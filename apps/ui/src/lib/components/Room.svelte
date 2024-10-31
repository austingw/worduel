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
	import type { View, ParsedAttempt } from '$lib/types';
	import { getRoomData } from '$lib/queries.svelte';
	import { getGameStart, setGameStart } from '$lib/game.svelte';
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

	let parsedAttempts = $derived.by<ParsedAttempt[][]>(() => {
		const answerMap = new Map<string, number>();
		const parsedAttempts: ParsedAttempt[][] = [];

		for (let i = 0; i < answer.length; i++) {
			answerMap.set(answer[i], (answerMap.get(answer[i]) || 0) + 1);
		}

		for (let attempt of attempts) {
			const clonedAnswerMap = new Map(answerMap);
			const parsedAttempt: ParsedAttempt[] = [];

			for (let i = 0; i < attempt.length; i++) {
				parsedAttempt.push({ val: attempt[i], class: 'badge-neutral' });
			}

			for (let i = 0; i < attempt.length; i++) {
				let count = clonedAnswerMap.get(attempt[i]) || 0;
				if (answer[i] == attempt[i]) {
					parsedAttempt[i].class = 'badge-success';
					clonedAnswerMap.set(attempt[i], count - 1);
				}
			}

			for (let i = 0; i < attempt.length; i++) {
				let count = clonedAnswerMap.get(attempt[i]) || 0;
				if (answer.includes(attempt[i]) && count > 0) {
					parsedAttempt[i].class = 'badge-warning';
					clonedAnswerMap.set(attempt[i], count - 1);
				}
			}
			parsedAttempts.push(parsedAttempt);
		}

		return parsedAttempts;
	});

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

<div class="flex flex-col w-full h-full items-center justify-between gap-4 pt-4">
	<div class="flex flex-col w-full h-full items-center justify-start gap-12 pt-8">
		<div class="flex flex-row w-9/12 md:w-1/3 items-center justify-between gap-4">
			{#if query?.data?.users?.[1] && query?.data?.users[1]?.name?.length > 0}
				<div class="shadow-black">
					<div class="flex flex-col items-start justify-center">
						<h2>
							{query?.data?.users[0].score ?? 0} - {query?.data?.users[1].score ?? 1}
						</h2>
						<h2>
							{query?.data?.users[0].name} vs. {query?.data?.users[1].name}
						</h2>
					</div>
				</div>
			{:else}
				<p>Waiting for another player...</p>
			{/if}
			<button
				class="btn btn-accent btn-xs"
				onclick={() => {
					if (ws) {
						sendLeave({ ws, room: getCurrentRoom(), username: name });
					}
					setGameStart(false);
					setCurrentRoom('');
					changeView('list' as View);
				}}>Leave</button
			>
		</div>
		<Grid {letters} {parsedAttempts} />
		{#if alert}
			<div
				role="alert"
				class="toast toast-top toast-center bg-warning absolute top-8 w-fit h-fit animate-none"
			>
				<p>{alert}</p>
			</div>
		{/if}
		<Keyboard {addLetter} {removeLetter} {submitAnswer} />
	</div>
	<Notifications />
</div>
