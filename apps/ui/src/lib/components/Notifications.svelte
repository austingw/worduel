<script lang="ts">
	import { getMessages, getShowNotification } from '$lib/websocket.svelte';
	let msgs = $state(getMessages());

	$effect(() => {
		msgs = getMessages();
	});
</script>

{#if getShowNotification()}
	<div role="alert" class="alert bg-secondary absolute bottom-4 right-4 w-fit h-fit animate-bounce">
		<p>{msgs[0]}</p>
	</div>
{/if}

<div class="flex flex-col bg-base-200 items-center justify-center w-11/12 sm:w-2/5">
	<div class="collapse collapse-plus">
		<input type="checkbox" />
		<div class="collapse-title text-xl font-medium">View Notifications</div>
		<div class="collapse-content">
			{#each msgs as msg}
				<div class="divider"></div>
				<p>{msg}</p>
			{/each}
		</div>
	</div>
</div>
