<script lang="ts">
	import type { View } from '$lib/types';

	type StartProps = {
		name: string;
		changeView: (newView: View) => void;
	};

	let { name, changeView }: StartProps = $props();
</script>

<div class="flex flex-col h-full items-center justify-start gap-4">
	<p>Welcome! Please enter a name to get started:</p>
	<div class="flex flex-row items-center justify-center gap-4">
		<input type="text" bind:value={name} placeholder="Enter name..." class="input input-bordered" />
		<button
			class="btn btn-secondary"
			onclick={() => {
				if (name === '') {
					alert('Please enter your name');
					return;
				}
				changeView('list');
			}}>Confirm</button
		>
	</div>
</div>
