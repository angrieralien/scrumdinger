<script lang="ts">
	import { getScrumContext, ScrumMeeting } from '$lib/models/scrum.svelte';
	import { getUserContext } from '$lib/models/user.svelte';
	import { RangeSlider } from '@skeletonlabs/skeleton';
	import { InputChip } from '@skeletonlabs/skeleton';
	import { getDrawerStore } from '@skeletonlabs/skeleton';

	const drawerStore = getDrawerStore();
	let scrums = getScrumContext();
	let scrum = new ScrumMeeting();

	let user = getUserContext();

	function submit() {
		if (user.isLoggedIn) {
			//POST
		} else {
			scrums.meetings.push(scrum);
		}

		drawerStore.close();
	}

	function cancel() {
		drawerStore.close();
	}
</script>

<div class="m-5">
	<h1 class="mt-3 h-[32px]">{scrum.name}</h1>

	<form onsubmit={submit}>
		<div class="grid gap-6 mb-6 md:grid-cols-2">
			<div class="mt-6">
				<input
					type="input"
					bind:value={scrum.name}
					class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
					placeholder="ScrumDinger"
					required
				/>
			</div>

			<div class="w-full flex">
				<RangeSlider
					class="grow pr-3"
					accent="accent-surface-500"
					name="range-slider"
					bind:value={scrum.minutes}
					max={60}
					min={5}
					step={5}
					ticked
				></RangeSlider>
				{scrum.minutes} minutes
			</div>
			<InputChip bind:value={scrum.attendees} name="chips" placeholder="Add attendee names" />

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
