import { env } from '$env/dynamic/public';
import { createQuery } from '@tanstack/svelte-query';

export const getRooms = () => {
	return createQuery(() => ({
		queryKey: ['rooms'],
		queryFn: async () => {
			const response = await fetch(`${env.PUBLIC_API_URL}/rooms`);
			if (!response.ok) {
				throw new Error('Failed to fetch rooms');
			}
			const data = await response.json();
			console.log(data.rooms);
			return Object.values(data.rooms);
		}
	}));
};
