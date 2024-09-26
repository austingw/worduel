<script lang="ts">
	import { getRooms } from '$lib/queries.svelte';

	const query = getRooms();
</script>

<table class="table table-zebra">
	<thead>
		<tr>
			<th></th>
			<th class="flex flex-grow">Room</th>
			<th># of Players</th>
			<th>Join?</th>
		</tr>
	</thead>
	<tbody> </tbody>
</table>

<h3>{query.data?.hello}</h3>
