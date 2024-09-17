<script>
	let { letters, attempts } = $props();
</script>

<div class="flex flex-col items-center justify-center w-full h-full gap-2">
	{#each attempts as attempt}
		<div class="flex flex-row items-center justify-center gap-2">
			{#each attempt as letter}
				<div class="badge badge-lg badge-ghost text-lg w-12 h-12">{letter}</div>
			{/each}
		</div>
	{/each}
	<div class="flex flow-row items-center justify-center gap-2">
		{#each letters as letter}
			<div class="badge badge-lg text-lg w-12 h-12">{letter}</div>
		{/each}
	</div>
</div>
