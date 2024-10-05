<script lang="ts">
	import type { Screen } from '$lib/types';

	type StartScreenProps = {
		name: string;
		changeScreen: (newScreen: Screen) => void;
	};

	let { name, changeScreen }: StartScreenProps = $props();
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
			changeScreen('list');
		}}>Confirm</button
	>
</div>
