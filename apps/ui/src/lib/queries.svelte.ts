import { env } from '$env/dynamic/public';
import { createQuery, type CreateQueryResult } from '@tanstack/svelte-query';
import type { Room } from './types';

export const getRooms = (): CreateQueryResult<Room[], Error> => {
	return createQuery(() => ({
		queryKey: ['rooms'],
		refetchInterval: 1000,
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

export const getRoomData = (roomName: string): CreateQueryResult<Room, Error> => {
	return createQuery(() => ({
		queryKey: ['room', roomName],
		queryFn: async () => {
			const response = await fetch(`${env.PUBLIC_API_URL}/rooms/${roomName}`);
			if (!response.ok) {
				throw new Error('Failed to fetch room data');
			}
			const room = await response.json();
			console.log('sss', room);
			return room.room;
		}
	}));
};
