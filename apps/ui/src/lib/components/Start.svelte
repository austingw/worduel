<script lang="ts">
	import type { View } from '$lib/types';

	type StartProps = {
		name: string;
		changeView: (newView: View) => void;
	};

	let { name, changeView }: StartProps = $props();
</script>

<div class="flex flex-row items-center justify-between">
	<input
		type="text"
		bind:value={name}
		placeholder="Enter your name..."
		class="input input-bordered"
	/>
	<button
		class="btn btn-secondary"
		onclick={() => {
			changeView('list');
		}}>Confirm</button
	>
</div>
