<script lang="ts">
	import type { PageData } from '../$types';
	import LinkIcon from '~icons/fa6-solid/arrow-up-right-from-square';

	export let data: PageData;
	const paragraphs = data.bill.text.split('\n');
	new Date().toLocaleDateString;
</script>

<div class="flex flex-col items-center">
	<div class="w-full max-w-xl flex flex-col gap-4">
		<h1 class="text-xl font-bold max-w-xl">{data.bill.title}</h1>

		<section class="flex flex-row justify-end">
			<a href={data.bill.url} target="_blank">
				<LinkIcon class="text-purple-700" />
			</a>
		</section>

		<section class="flex flex-row w-full items-center justify-between">
			<span>Issued: {data.bill.issued.toLocaleDateString()}</span>
			<span>Updated: {data.bill.updated.toLocaleDateString()}</span>
		</section>

		<hr />

		<section>
			<h2 class="font-medium text-lg">Actions</h2>
			<ul class="flex flex-col gap-2">
				{#each data.bill.actions as action}
					<li class="p-2 border rounded flex flex-row items-center gap-2 justify-between">
						<p>{action.text}</p>
						<p>{action.date.toLocaleDateString()}</p>
					</li>
				{/each}
			</ul>
		</section>

		<hr />

		<section class="flex flex-col items-center gap-2 text-center w-full">
			{#each paragraphs as p}
				<p class="w-full max-w-xl">{p}</p>
			{/each}
		</section>
	</div>
</div>
