<script lang="ts">
	import { getScrumContext as getScrumContext } from '$lib/models/scrum.svelte';
	import { getUserContext, User } from '$lib/models/user.svelte';
	import { getDrawerStore, type DrawerSettings, type DrawerStore } from '@skeletonlabs/skeleton';
	import { CalendarClock, Plus } from 'lucide-svelte';
	import { onMount } from 'svelte';

	let user: User;
	let drawerStore: DrawerStore;

	onMount(() => {
		user = getUserContext();
		drawerStore = getDrawerStore();
	});

	// contexts
	let scrum = getScrumContext();

	function onAdd() {
		const settings: DrawerSettings = { id: 'example-1' };
		drawerStore.open(settings);
	}
</script>

<div class="h-full">
	{#if scrum.meetings.length > 0}
		<div>
			<dl class="list-dl">
				{#each scrum.meetings as s}
					<div>
						<span class="badge"><CalendarClock></CalendarClock></span>
						<span class="flex-auto">
							<dt>{s.name}</dt>
							<dd>{s.minutes} minutes</dd>
						</span>
					</div>
				{/each}
			</dl>
		</div>
	{:else}
		<div class="flex flex-col items-center h-full">
			<div class="text-center p-6 text-2xl">
				Ahh Snap!!! <br />Login to view your scrums or create a new meeting.
			</div>
		</div>
	{/if}

	<button
		onclick={onAdd}
		type="button"
		class="add-button bg-secondary-500 m-3 btn-icon btn-icon-xl"
	>
		<Plus></Plus>
	</button>
</div>

<style>
	.add-button {
		position: fixed;
		right: 0;
		bottom: 0;
	}
</style>
