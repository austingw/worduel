<script lang="ts">
	import { env } from '$env/dynamic/public';
	import { createQuery } from '@tanstack/svelte-query';

	const query = createQuery(() => ({
		queryKey: ['rooms'],
		queryFn: async () => {
			const response = await fetch(`${env.PUBLIC_API_URL}/rooms`);
			if (!response.ok) {
				throw new Error('Failed to fetch rooms');
			}
			return await response.json();
		}
	}));
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
