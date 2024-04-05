<script lang="ts">
	import BillSummary from '$lib/bill-summary.svelte';
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<h1 class="text-xl font-bold">Bills</h1>
<hr class="my-8" />

<div class="flex flex-col items-center gap-8">
	<ol class="flex flex-col gap-4 max-w-xl">
		{#each data.bills as bill}
			<li>
				<BillSummary {bill} />
			</li>
		{/each}
	</ol>

	<section class="flex flex-row items-center gap-8 w-full max-w-xl justify-between">
		{#if data.page > 0}
			<a href={`/?page=${Number(data.page ?? 0) - 1}`}>
				<button class="bg-purple-200 px-4 py-2 rounded">Prev</button>
			</a>
		{:else}
			<span></span>
		{/if}

		<p>Page {data.page + 1}</p>

		{#if data.bills.length > 0}
			<a href={`/?page=${Number(data.page ?? 0) + 1}`}>
				<button class="bg-blue-200 px-4 py-2 rounded">Next</button>
			</a>
		{:else}
			<span></span>
		{/if}
	</section>
</div>
