<script>
	import ScrumTimer from '$lib/components/ScrumTimer.svelte';

	let minutes = $state(1);
	let countdown = $derived(minutes * 60);
	let attendeesNames = $state({
		names: ['Stephen', 'Mark', 'John']
	});

	let attendees = $derived(attendeesNames.names);
</script>

<div class="flex h-full">
	<ScrumTimer {countdown} {attendees} />
	<!-- <div class="flex flex-row items-center">
		<label for="minutes-text" class="block mb-2 text-sm font-medium dark:text-white pr-3"
			>Minutes</label
		>
		<input
			type="number"
			id="minutes-text"
            required
			aria-describedby="minutes-text-explanation"
			class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
			placeholder="5"
            value={minutes}
		/>
	</div> -->
</div>
