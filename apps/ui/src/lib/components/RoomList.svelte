<script lang="ts">
	import { getRooms } from '$lib/queries.svelte';

	const query = getRooms();
</script>

<div class="flex flex-col items-center justify-center">
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
	{/if}
</div>
