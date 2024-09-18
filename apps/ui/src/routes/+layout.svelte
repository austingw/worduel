<script lang="ts">
	import type { Snippet } from 'svelte';
	import '../app.css';
	let { children }: { children: Snippet } = $props()
</script>

<div class="w-full flex flex-col items-center justify-center">
	{@render children()}
</div>
