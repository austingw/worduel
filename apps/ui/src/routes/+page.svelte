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
</script>

<div class="w-full h-full flex flex-col items-center justify-start gap-8 pt-12">
	<article class="prose text-center">
		<h1 class="underline">Worduel {name}</h1>
	</article>

	{#if view === 'start'}
		<Start {inputName} {changeView} />
	{:else if view === 'list'}
		<RoomList {name} {changeView} />
	{:else if view === 'room'}
		<Room {name} {changeView} />
	{/if}
	<button class="btn btn-secondary" onclick={() => console.log(name)}>Jm</button>
</div>
