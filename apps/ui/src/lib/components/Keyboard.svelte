<div class="w-full flex flex-col justify-end items-center gap-1">
	{#each ['qwertyuiop', 'asdfghjkl'] as row}
		<div class="flex flex-row justify-center items-center gap-1">
			{#each row as letter}
				<button class="btn btn-neutral">{letter}</button>
			{/each}
		</div>
	{/each}

	<div class="flex flex-row justify-center items-center gap-1">
		<button class="btn btn-neutral">Enter</button>
		{#each 'zxcvbnm' as letter}
			<button class="btn btn-neutral">{letter}</button>
		{/each}

		<button class="btn btn-neutral">← Back</button>
	</div>
</div>
