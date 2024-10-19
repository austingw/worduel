<script lang="ts">
	import { createRoom } from '$lib/mutations.svelte';
	import { getRooms } from '$lib/queries.svelte';
	import type { View } from '$lib/types';
	import { useQueryClient } from '@tanstack/svelte-query';
	import { setWs, sendJoin, setCurrentRoom } from '$lib/websocket.svelte';

	type RoomListProps = {
		name: string;
		changeView: (newView: View) => void;
	};

	let { name, changeView }: RoomListProps = $props();
	const create = createRoom();
	const query = getRooms();
	const queryClient = useQueryClient();

	function handleCreateRoom() {
		create.mutateAsync(name).then((res) => {
			if (res.newRoom) {
				queryClient.invalidateQueries({
					queryKey: ['rooms']
				});
				handleJoinRoom(res.newRoom.name);
			}
		});
	}

	async function handleJoinRoom(room: string) {
		const ws = await setWs();
		if (ws?.readyState === WebSocket.OPEN) {
			sendJoin({
				ws,
				room,
				username: name
			});
			setCurrentRoom(room);
			changeView('room');
		}
	}

	$effect(() => {
		setCurrentRoom('');
	});
</script>

<div class="flex flex-col h-full w-1/2 items-center justify-between gap-4 pt-8">
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
						<td>{room?.name}</td>
						<td>{room?.users?.filter((u) => u?.name?.length > 0)?.length} / 2</td>
						<td>
							<button
								class={`btn btn-accent btn-xs ${
									room?.users[0]?.name !== '' && room?.users[1]?.name !== '' && 'btn-disabled'
								}`}
								onclick={() => handleJoinRoom(room?.name)}
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
		<div class="flex flex-row w-full items-center justify-center gap-4 pt-4 mb-8">
			<button class="btn btn-secondary" onclick={handleCreateRoom}>Create Room</button>
		</div>
	{/if}
</div>
