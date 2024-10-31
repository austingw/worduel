<script lang="ts">
	import type { ParsedAttempt } from '$lib/types';

	type GridProps = {
		letters: string[];
		parsedAttempts: ParsedAttempt[][];
	};

	let { letters, parsedAttempts }: GridProps = $props();
	let blanks = $derived.by<string[]>(() => {
		if (letters.length < 5) {
			return Array(5 - letters.length).fill('');
		}
		return [];
	});
</script>

<div class="flex flex-col items-center justify-start w-full h-fit gap-2">
	{#each parsedAttempts as attempt}
		<div class="flex flex-row items-center justify-center gap-2">
			{#each attempt as letter}
				<div class={`badge badge-lg ${letter.class} text-lg w-12 h-12`}>
					{letter.val}
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
