import { env } from '$env/dynamic/public';
import { createMutation } from '@tanstack/svelte-query';

export const createRoom = () => {
	return createMutation(() => ({
		mutationKey: ['createRoom'],
		mutationFn: async (name: string) => {
			const response = await fetch(`${env.PUBLIC_API_URL}/rooms`, {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify({ name })
			});
			if (!response.ok) {
				throw new Error('Failed to create room');
			}
			return await response.json();
		}
	}));
};
