<script lang="ts">
	import { getScrumsContext, ScrumMeeting } from '$lib/models/scrum.svelte';
	import { RangeSlider } from '@skeletonlabs/skeleton';
	import { InputChip } from '@skeletonlabs/skeleton';
	let scrums = getScrumsContext();
	let scrum = new ScrumMeeting();

	function done() {}
</script>

<div>
	<h1>{scrum.name}</h1>
	<form>
		<div class="grid gap-6 mb-6 md:grid-cols-2">
			<div class="mb-6">
				<label for="info" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
					>Meeting Info</label
				>
				<input
					type="input"
					bind:value={scrum.name}
					class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
					placeholder="john.doe@company.com"
					required
				/>
			</div>

			<RangeSlider
				accent="accent-surface-500"
				name="range-slider"
				bind:value={scrum.time}
				max={60}
				step={5}
				ticked
			>
				<div class="flex justify-between items-center">
					<div class="font-bold">Minutes</div>
					<div class="text-xs">{scrum.time}</div>
				</div></RangeSlider
			>
			<label for="attendees" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
				>Attendees</label
			>

			<InputChip bind:value={scrum.attendees} name="chips" placeholder="Enter a name..." />

			<button
				type="submit"
				class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm w-full sm:w-auto px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
				>Submit</button
			>
		</div>
	</form>
	{#each scrum.attendees as a}
		<p>{a}</p>
	{/each}
</div>
