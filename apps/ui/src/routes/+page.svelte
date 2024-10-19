<script lang="ts">
	import Room from '$lib/components/Room.svelte';
	import RoomList from '$lib/components/RoomList.svelte';
	import Start from '$lib/components/Start.svelte';
	import Footer from '$lib/components/Footer.svelte';
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

<div
	class="w-full max-w-screen max-full min-h-screen h-full flex flex-col items-center justify-between pt-12"
>
	<div class="flex flex-col w-full h-full items-center justify-start">
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
	<Footer />
</div>
