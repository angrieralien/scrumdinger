<script lang="ts">
	import { getDrawerContext } from '$lib/models/drawer.svelte';
	import { getScrumContext, ScrumMeeting } from '$lib/models/scrum.svelte';
	import { getUserContext } from '$lib/models/user.svelte';
	import { RangeSlider } from '@skeletonlabs/skeleton';
	import { InputChip } from '@skeletonlabs/skeleton';
	import { getDrawerStore } from '@skeletonlabs/skeleton';
	import { Check, Trash2 } from 'lucide-svelte';
	import { onDestroy, onMount } from 'svelte';

	const drawerStore = getDrawerStore();

	let drawerMeta = getDrawerContext();
	let scrums = getScrumContext();
	let scrum = new ScrumMeeting();
	let editScrum: ScrumMeeting;
	let editIdx: number;

	let user = getUserContext();

	let state = 'create';

	onMount(() => {
		if (drawerMeta && drawerMeta.data && drawerMeta.data['editScrum']) {
			editScrum = drawerMeta.data['editScrum']['scrum'];
			editIdx = drawerMeta.data['editScrum']['idx'];
			if (editScrum !== null) {
				copyScrum(editScrum, scrum);
				state = 'edit';
			}
		}
	});

	onDestroy(() => {
		delete drawerMeta.data['editScrum'];
	});

	function copyScrum(src: ScrumMeeting, dst: ScrumMeeting) {
		dst.attendees.length = 0;
		src.attendees.forEach((a) => {
			dst.attendees.push(a);
		});

		dst.color = src.color;
		dst.minutes = src.minutes;
		dst.name = src.name;
	}

	function deleteMeeting() {
		scrums.meetings.splice(editIdx, 1);
		drawerStore.close();
	}

	function submit() {
		if (user.isLoggedIn) {
			console.log('here');
			//POST
		} else {
			if (state === 'edit') {
				copyScrum(scrum, editScrum);
			} else {
				scrums.meetings.push(scrum);
			}
		}

		drawerStore.close();
	}

	function cancel() {
		drawerStore.close();
	}

	function colorSelected(c: string) {
		scrum.color = c;
	}
</script>

<div class="m-5">
	<div class="mt-3 h-[32px] flex flex-row">
		<h1 class="">{scrum.name}</h1>
		{#if state === 'edit'}
			<span class="grow"></span>

			<Trash2 onclick={deleteMeeting} />
		{/if}
	</div>
	<form onsubmit={submit}>
		<div class="flex flex-col my-6">
			<div class="">
				<input
					type="input"
					bind:value={scrum.name}
					class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
					placeholder="ScrumDinger"
					required
				/>
			</div>

			<div class="w-full flex my-6">
				<RangeSlider
					class="grow pr-3"
					accent="accent-surface-500"
					name="range-slider"
					bind:value={scrum.minutes}
					max={60}
					min={1}
					step={1}
					ticked
				></RangeSlider>
				{scrum.minutes} minutes
			</div>

			<div class="flex flex-row justify-center my-3">
				{#each ['bg-primary-500', 'bg-secondary-500', 'bg-white', 'bg-success-500', 'bg-warning-500'] as c}
					<button
						onclick={() => {
							colorSelected(c);
						}}
						type="button"
						class="btn-icon ring-2 ring-gray-500 m-3 {c}"
					>
						{#if scrum.color === c}
							<Check></Check>
						{/if}
					</button>
				{/each}
			</div>

			<InputChip
				class="my-6"
				bind:value={scrum.attendees}
				name="chips"
				placeholder="Add attendee names"
			/>

			<div class="flex flex-row justify-end w-full">
				<button
					onclick={() => cancel()}
					type="button"
					class="btn variant-filled rounded btn-md bg-secondary-500 m-2">Cancel</button
				>
				<button type="submit" class="btn variant-filled rounded btn-md bg-secondary-500 m-2"
					>Submit</button
				>
			</div>
		</div>
	</form>
</div>
