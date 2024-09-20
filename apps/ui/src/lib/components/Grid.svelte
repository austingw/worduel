<script lang="ts">
	type GridProps = {
		letters: string[];
		attempts: string[];
		answer: string;
	};

	let { letters, attempts, answer }: GridProps = $props();
	let blanks = $derived.by<string[]>(() => {
		if (letters.length < 5) {
			return Array(5 - letters.length).fill('');
		}
		return [];
	});

	function styleLetter(attemptChar: string, answerChar: string) {
		if (attemptChar === answerChar) {
			return 'badge-success';
		} else if (answer.includes(attemptChar)) {
			return 'badge-warning';
		} else {
			return 'badge-neutral';
		}
	}
</script>

<div class="flex flex-col items-center justify-center w-full h-full gap-2">
	{#each attempts as attempt}
		<div class="flex flex-row items-center justify-center gap-2">
			{#each attempt as letter, index}
				<div class={`badge badge-lg ${styleLetter(letter, answer[index])} text-lg w-12 h-12`}>
					{letter}
				</div>
			{/each}
		</div>
	{/each}
	<div class="flex flow-row items-center justify-center gap-2">
		{#each letters as letter}
			<div class="badge badge-lg text-lg w-12 h-12">{letter}</div>
		{/each}
		{#each blanks as blank}
			<div class="badge badge-lg text-lg w-12 h-12">{blank}</div>
		{/each}
	</div>
</div>
