import { env } from '$env/dynamic/public';
import { createQuery, type CreateQueryResult } from '@tanstack/svelte-query';
import type { Room } from './types';

export const getRooms = (): CreateQueryResult<Room[], Error> => {
	return createQuery(() => ({
		queryKey: ['rooms'],
		queryFn: async () => {
			const response = await fetch(`${env.PUBLIC_API_URL}/rooms`);
			if (!response.ok) {
				throw new Error('Failed to fetch rooms');
			}
			const data = await response.json();
			return Object.values(data.rooms);
		}
	}));
};
