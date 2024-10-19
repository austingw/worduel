<script lang="ts">
	import Room from '$lib/components/Room.svelte';
	import RoomList from '$lib/components/RoomList.svelte';
	import Start from '$lib/components/Start.svelte';
	import type { View } from '$lib/types';

	let name = $state<string>('');
	let view = $state<View>('start');

	function changeView(newView: View) {
		view = newView;
	}

	function inputName(newName: string) {
		name = newName;
	}

	$effect(() => {
		const localName = localStorage.getItem('name');
		if (localName && localName?.length > 0) {
			name = localName;
			changeView('list');
		}
	});
</script>

<div class="w-full max-w-screen h-screen flex flex-col items-center justify-start pt-12">
	<article class="prose text-center">
		<h1 class="underline">Worduel</h1>
	</article>

	{#if view === 'start'}
		<Start {inputName} {changeView} />
	{:else if view === 'list'}
		<RoomList {name} {changeView} />
	{:else if view === 'room'}
		<Room {name} {changeView} />
	{/if}
</div>
