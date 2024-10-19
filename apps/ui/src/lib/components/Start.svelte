<script lang="ts">
	import type { View } from '$lib/types';

	type StartProps = {
		inputName: (name: string) => void;
		changeView: (newView: View) => void;
	};

	let { inputName, changeView }: StartProps = $props();
	let newName = $state<string>('');

	function handleNameInput(name: string) {
		if (name === '') {
			alert('Please enter your name');
			return;
		}
		inputName(name);
		localStorage.setItem('name', name);
		changeView('list');
	}
</script>

<div class="flex flex-col h-full w-9/12 items-center justify-start gap-4 pt-8">
	<p>Welcome! Please enter a name to get started:</p>
	<div class="flex flex-row items-center justify-center gap-4">
		<input
			type="text"
			bind:value={newName}
			placeholder="Enter name..."
			class="input input-bordered"
		/>
		<button class="btn btn-secondary" onclick={() => handleNameInput(newName)}>Confirm</button>
	</div>
</div>
