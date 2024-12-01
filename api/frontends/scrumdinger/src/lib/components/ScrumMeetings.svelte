<script lang="ts">
	import { getScrumContext as getScrumContext } from '$lib/models/scrum.svelte';
	import { getUserContext, User } from '$lib/models/user.svelte';
	import { getDrawerStore, type DrawerSettings, type DrawerStore } from '@skeletonlabs/skeleton';
	import { CalendarClock, Ellipsis, EllipsisVertical, Plus } from 'lucide-svelte';
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

<div class="h-full p-4">
	<img class="rounded-3xl" src="/images/scrum_team_rec.svg" alt="scrum team" />

	{#if scrum.meetings.length > 0}
		<div class="pb-6">
			<dl class="list-dl">
				{#each scrum.meetings as s}
					<div class="border-b-2 bg-on-surface">
						<div class="rounded-full h-10 w-10 ring-1 ring-gray-500 m-3 {s.color}">
							<CalendarClock></CalendarClock>
						</div>
						<span class="flex-auto">
							<dt>{s.name}</dt>
							<dd>{s.minutes} minutes</dd>
						</span>
						<span class="full"></span>
						<button onclick={() => {}} type="button" class="btn-icon ring-2 m-3">
							<EllipsisVertical></EllipsisVertical>
						</button>
					</div>
				{/each}
			</dl>
		</div>
	{:else}
		<div class="flex flex-col items-center">
			<div class="text-center p-6 text-2xl">
				Ahh Snap!!! <br />Login to view your scrums or create a new meeting.
			</div>
		</div>
	{/if}

	<button onclick={onAdd} type="button" class="add-button bg-primary-500 m-3 btn-icon btn-icon-xl">
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
