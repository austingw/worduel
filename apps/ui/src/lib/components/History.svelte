<script lang="ts">
	import { getMessages } from '$lib/websocket.svelte';
	let msgs = $state(getMessages());

	$effect(() => {
		msgs = getMessages();
	});
</script>

<div class="flex flex-col bg-base-200 items-center justify-center w-11/12 sm:w-2/5">
	<div class="collapse collapse-plus">
		<input type="checkbox" />
		<div class="collapse-title text-xl font-medium">Room History</div>
		<div class="collapse-content">
			{#each msgs as msg}
				<p>{msg}</p>
			{/each}
		</div>
	</div>
</div>
