<script lang="ts">
	type KeyboardProps = {
		addLetter: (letter: string) => void;
		removeLetter: () => void;
		submitAnswer: () => void;
	};

	let { addLetter, removeLetter, submitAnswer }: KeyboardProps = $props();
</script>

<div class="w-full flex flex-col justify-end items-center gap-1">
	{#each ['qwertyuiop', 'asdfghjkl'] as row}
		<div class="flex flex-row justify-center items-center gap-1">
			{#each row as letter}
				<button class="btn btn-neutral btn-sm sm:btn-md" onclick={() => addLetter(letter)}
					>{letter}</button
				>
			{/each}
		</div>
	{/each}

	<div class="flex flex-row justify-center items-center gap-1">
		<button class="btn btn-primary btn-sm sm:btn-md" onclick={() => submitAnswer()}>Enter</button>
		{#each 'zxcvbnm' as letter}
			<button class="btn btn-neutral btn-sm sm:btn-md" onclick={() => addLetter(letter)}
				>{letter}</button
			>
		{/each}

		<button class="btn btn-secondary btn-sm sm:btn-md" onclick={() => removeLetter()}>← Back</button
		>
	</div>
</div>
