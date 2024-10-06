<script lang="ts">
	import { QueryClient, QueryClientProvider } from '@tanstack/svelte-query';
	import type { Snippet } from 'svelte';
	import '../app.css';
	import { browser } from '$app/environment';

	let { children }: { children: Snippet } = $props();
	const queryClient = new QueryClient({
		defaultOptions: {
			queries: {
				enabled: browser
			}
		}
	});
</script>

<QueryClientProvider client={queryClient}>
	<div class="w-screen h-screen flex flex-col items-center justify-center">
		{@render children()}
	</div>
</QueryClientProvider>
