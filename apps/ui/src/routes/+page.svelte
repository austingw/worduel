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

	export function getName() {
		return name;
	}
</script>

<div class="w-full h-full flex flex-col items-center justify-start gap-8 pt-8">
	<article class="prose text-center">
		<h1 class="underline align">Worduel</h1>
	</article>
	{#if view === 'start'}
		<Start {name} {changeView} />
	{:else if view === 'list'}
		<RoomList {getName} />
	{:else if view === 'room'}
		<Room />
	{/if}
</div>
