<script lang="ts">
	import { getRooms } from '$lib/queries.svelte';

	const query = getRooms();
</script>

<div class="flex flex-col items-center justify-center">
	{#if query.isLoading}
		<span class="loading loading-spinner loading-lg text-secondary"></span>
	{:else if query.data}
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
				{#each query.data.rooms as room, index}
					<tr>
						<td>{index + 1}</td>
						<td>{room.name}</td>
						<td>{room.players.length}</td>
						<td>
							<button
								class={`btn btn-secondary ${room.players.length > 1 && 'btn-disabled'}`}
								onclick={() => console.log('joins room')}>Join</button
							>
						</td>
					</tr>
				{/each}
			</tbody>
		</table>
	{/if}
</div>
