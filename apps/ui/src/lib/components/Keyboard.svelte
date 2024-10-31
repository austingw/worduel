<script lang="ts">
	import { getGameStart } from '$lib/game.svelte';
	import type { ParsedAttempt } from '$lib/types';

	type KeyboardProps = {
		addLetter: (letter: string) => void;
		removeLetter: () => void;
		submitAnswer: () => void;
		parsedAttempts: ParsedAttempt[][];
	};

	let { addLetter, removeLetter, submitAnswer, parsedAttempts }: KeyboardProps = $props();

	let incorrectLetters = $derived.by<string[]>(() => {
		const incorrectLetters = parsedAttempts.flat();
		const letterSet = new Set(incorrectLetters.map((letter) => letter.val));
		for (let letter of incorrectLetters) {
			if (letter.class === 'badge-success' || letter.class === 'badge-warning') {
				letterSet.delete(letter.val);
			}
		}
		return [...letterSet];
	});
</script>

<div class="w-full flex flex-col justify-end items-center gap-1">
	{#each ['qwertyuiop', 'asdfghjkl'] as row}
		<div class="flex flex-row justify-center items-center gap-1">
			{#each row as letter}
				<button
					class={`btn btn-circle ${!incorrectLetters.includes(letter) && 'btn-neutral'} btn-sm sm:btn-md`}
					disabled={!getGameStart()}
					onclick={() => addLetter(letter)}>{letter}</button
				>
			{/each}
		</div>
	{/each}

	<div class="flex flex-row justify-center items-center gap-1">
		<button
			class="btn btn-circle btn-secondary btn-sm sm:btn-md text-[.5rem] sm:text-sm"
			disabled={!getGameStart()}
			onclick={() => submitAnswer()}>Enter</button
		>
		{#each 'zxcvbnm' as letter}
			<button
				class={`btn btn-circle ${!incorrectLetters.includes(letter) && 'btn-neutral'} btn-sm sm:btn-md`}
				disabled={!getGameStart()}
				onclick={() => addLetter(letter)}>{letter}</button
			>
		{/each}

		<button
			class="btn btn-circle text-lg btn-accent btn-sm sm:btn-md"
			disabled={!getGameStart()}
			onclick={() => removeLetter()}
			>←
		</button>
	</div>
</div>
