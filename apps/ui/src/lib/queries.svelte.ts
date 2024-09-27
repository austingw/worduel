import { env } from "$env/dynamic/public";
import { createQuery } from "@tanstack/svelte-query";

export const getRooms = () => {
return createQuery(() => ({
		queryKey: ['rooms'],
		queryFn: async () => {
			const response = await fetch(`${env.PUBLIC_API_URL}/rooms`);
			if (!response.ok) {
				throw new Error('Failed to fetch rooms');
			}
			return await response.json();
		}
	}));

}
