<script lang="ts">
	import { getScrumContext, ScrumMeeting } from '$lib/models/scrum.svelte';
	import { getUserContext, User } from '$lib/models/user.svelte';
	import {
		getDrawerStore,
		getToastStore,
		type DrawerSettings,
		type DrawerStore,
		type ToastSettings
	} from '@skeletonlabs/skeleton';
	import { CalendarClock, History, Pencil, Plus } from 'lucide-svelte';
	import { onMount } from 'svelte';
	import ScrumTimer from './ScrumTimer.svelte';
	import { DrawerMeta, getDrawerContext, setDrawerContext } from '$lib/models/drawer.svelte';

	const toastStore = getToastStore();

	let user: User;
	let drawerStore: DrawerStore;
	let drawerMeta: DrawerMeta;
	let showTimer = $state(false);
	let selectedScrum: ScrumMeeting = $state(new ScrumMeeting());

	onMount(() => {
		user = getUserContext();
		drawerStore = getDrawerStore();
		drawerMeta = getDrawerContext();
	});

	// contexts
	let scrum = getScrumContext();

	function onAdd() {
		const settings: DrawerSettings = {};
		delete drawerMeta.data['editScrum'];
		drawerMeta.data['component'] = 'CRUDScrum';

		drawerStore.open(settings);
	}

	function onEdit(event: any, s: ScrumMeeting, idx: number) {
		event.stopPropagation();
		const settings: DrawerSettings = {};
		drawerMeta.data['component'] = 'CRUDScrum';
		drawerMeta.data['editScrum'] = {
			scrum: s,
			idx: idx
		};
		drawerStore.open(settings);
	}

	function onHistory(event: any, s: ScrumMeeting) {
		event.stopPropagation();
		const settings: DrawerSettings = {};
		drawerMeta.data['component'] = 'HistoryScrum';
		drawerMeta.data['historyScrum'] = {
			scrum: s
		};
		drawerStore.open(settings);
	}

	function onselect(s: ScrumMeeting) {
		if (s.attendees.length < 1) {
			let t: ToastSettings = {
				message: 'No Attendees assigned to this meeting!',
				timeout: 5000,
				background: 'bg-error-500'
			};
			toastStore.trigger(t);
		} else {
			selectedScrum = s;
			showTimer = true;
		}
	}
</script>

{#if !showTimer}
	<div class="h-full p-4">
		<!-- <img class="rounded-3xl" src="/images/scrum_team_rec.svg" alt="scrum team" /> -->

		{#if scrum.meetings.length > 0}
			<div class="pb-6">
				<dl class="list-dl">
					{#each scrum.meetings as s, idx}
						<!-- svelte-ignore a11y_click_events_have_key_events -->
						<!-- svelte-ignore a11y_no_static_element_interactions -->
						<div class="border-b-2 border-white cursor-pointer" onclick={() => onselect(s)}>
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
									onHistory(event, s);
								}}
								type="button"
								class="btn-icon ring-gray-500 ring-2 m-3"
							>
								<History />
							</button>
							<button
								onclick={(event) => {
									onEdit(event, s, idx);
								}}
								type="button"
								class="btn-icon ring-gray-500 ring-2 m-3"
							>
								<Pencil />
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
