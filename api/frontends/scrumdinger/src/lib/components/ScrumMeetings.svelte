<script lang="ts">
	import { goto } from '$app/navigation';
	import {
		getScrumContext as getScrumContext,
		Scrum,
		ScrumMeeting
	} from '$lib/models/scrum.svelte';
	import { getUserContext, User } from '$lib/models/user.svelte';
	import { getDrawerStore, type DrawerSettings, type DrawerStore } from '@skeletonlabs/skeleton';
	import { CalendarClock, Ellipsis, EllipsisVertical, Plus } from 'lucide-svelte';
	import { onMount } from 'svelte';
	import ScrumTimer from './ScrumTimer.svelte';
	import { stopPropagation } from 'svelte/legacy';

	let user: User;
	let drawerStore: DrawerStore;
	let showTimer = $state(false);
	let selectedScrum: ScrumMeeting = $state(new ScrumMeeting());

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

	function onselect(s: ScrumMeeting) {
		selectedScrum = s;
		showTimer = true;
	}
</script>

{#if !showTimer}
	<div class="h-full p-4">
		<!-- <img class="rounded-3xl" src="/images/scrum_team_rec.svg" alt="scrum team" /> -->

		{#if scrum.meetings.length > 0}
			<div class="pb-6">
				<dl class="list-dl">
					{#each scrum.meetings as s}
						<!-- svelte-ignore a11y_click_events_have_key_events -->
						<!-- svelte-ignore a11y_no_static_element_interactions -->
						<div class="border-b-2" onclick={() => onselect(s)}>
							<div class="rounded-full h-10 w-10 ring-1 ring-gray-500 m-3 {s.color}">
								<CalendarClock></CalendarClock>
							</div>
							<span class="flex-auto">
								<dt>{s.name}</dt>
								<dd>{s.minutes} minutes</dd>
							</span>
							<span class="full"></span>
							<button
								onclick={(event) => {
									event.stopPropagation();
								}}
								type="button"
								class="btn-icon ring-gray-500 ring-2 m-3"
							>
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

		<button
			onclick={onAdd}
			type="button"
			class="add-button bg-secondary-500 m-3 btn-icon btn-icon-xl"
		>
			<Plus></Plus>
		</button>
	</div>
{/if}
{#if showTimer}
	<ScrumTimer
		countdown={selectedScrum.minutes * 60}
		attendees={selectedScrum.attendees}
		done={() => {
			selectedScrum = new ScrumMeeting();
			showTimer = false;
		}}
	></ScrumTimer>
{/if}

<style>
	.add-button {
		position: fixed;
		right: 0;
		bottom: 0;
	}
</style>
