<script lang="ts">
	import type { BillSummary } from '$lib';
	import Arrow from '~icons/fa6-solid/angle-down';
	import LinkIcon from '~icons/fa6-solid/arrow-up-right-from-square';

	let actionsExpanded = false;
	export let bill: BillSummary;
</script>

<a href={`/${bill.bill_id}`}>
	<div class="p-2 rounded bg-gray-200 flex flex-col gap-4">
		<section class="flex flex-row items-start gap-4 justify-between">
			<p>{bill.short_title}</p>
			<a href={bill.url} target="_blank">
				<LinkIcon class="text-purple-800" />
			</a>
		</section>

		<section class="flex flex-row flex-wrap justify-between gap-4 text-sm">
			<p>Issued: {bill.issued.toLocaleDateString()}</p>
			<p>Updated: {bill.updated.toLocaleDateString()}</p>
		</section>

		<section class="rounded-lg bg-gray-300 p-2 flex flex-col gap-4">
			<button
				on:click={() => (actionsExpanded = !actionsExpanded)}
				class="w-full cursor-pointer flex flex-row items-center justify-between"
			>
				<span class="">
					Actions ({bill.actions.length})
				</span>
				<Arrow class={`${actionsExpanded ? 'rotate-180' : ''}`} />
			</button>
			<ul class={`${actionsExpanded ? '' : 'hidden'} flex flex-col gap-2`}>
				{#each bill.actions as action}
					<li class="flex flex-row items-center justify-between gap-2">
						<p>
							{action.text}
						</p>
						<span>{action.date.toLocaleDateString()}</span>
					</li>
				{/each}
			</ul>
		</section>
	</div>
</a>
