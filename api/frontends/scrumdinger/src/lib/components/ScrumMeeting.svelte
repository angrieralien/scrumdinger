<script>
	import ScrumTimer from '$lib/components/ScrumTimer.svelte';

	let minutes = $state(5);
	let countdown = $derived(minutes * 60);
	let attendeesNames = $state({
		names: ['Stephen', 'Mark', 'John']
	});

	let attendees = $derived(attendeesNames.names);
</script>

<div class="flex h-full">
	<ScrumTimer {countdown} {attendees} />
</div>
