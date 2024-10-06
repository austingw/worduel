<script lang="ts">
	import { createRoom } from '$lib/mutations.svelte';
	import { getRooms } from '$lib/queries.svelte';

	type RoomListProps = {
		name: string;
	};

	let { name }: RoomListProps = $props();
	const create = createRoom();
	const query = getRooms();

	function handleCreateRoom() {
		console.log(name);
		create.mutateAsync(name).then((res) => console.log(res));
	}
	console.log(name);
</script>

<div class="flex flex-col h-full w-1/2 items-center justify-start gap-4">
	{#if query.isLoading}
		<span class="loading loading-spinner loading-lg text-secondary"></span>
	{:else if query?.data}
		<table class="table table-zebra">
			<thead>
				<tr>
					<th>#</th>
					<th class="flex flex-grow">Room</th>
					<th># of Players</th>
					<th>Join?</th>
				</tr>
			</thead>
			<tbody>
				{#each query.data as room, index}
					<tr>
						<td>{index + 1}</td>
						<td>{room.name}</td>
						<td>1 / 2</td>
						<td>
							<button
								class={`btn btn-accent btn-xs ${
									room?.users[0]?.name !== '' && room?.users[1]?.name !== '' && 'btn-disabled'
								}`}
								onclick={() => console.log('joins room')}
								>{room?.users[0]?.name !== '' && room?.users[1]?.name !== ''
									? 'Full'
									: 'Join'}</button
							>
						</td>
					</tr>
				{/each}
			</tbody>
		</table>
	{:else}
		<p>No rooms available</p>
	{/if}
	{#if !query.isLoading}
		<div class="flex flex-row w-full items-center justify-center gap-4">
			<button class="btn btn-secondary" onclick={handleCreateRoom}>Create Room</button>
		</div>
	{/if}
</div>
