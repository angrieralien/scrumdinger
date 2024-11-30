<script lang="ts">
	import { createEventDispatcher, onDestroy } from 'svelte';
	import { tweened } from 'svelte/motion';
	import { linear as easing } from 'svelte/easing';
	import { fly } from 'svelte/transition';
	import { Play, Pause, RotateCcw } from 'lucide-svelte';

	const dispatch = createEventDispatcher();

	//export let countdown;

	let { countdown, attendees } = $props();

	let attendeeCountdown = Math.floor(countdown / attendees);
	let attendeesRemaining = $state(attendees);
	let now = $state(Date.now());
	let end = $state(now + attendeeCountdown * 1000);

	const count = $derived(Math.round((end - now) / 1000));
	const h = $derived(Math.floor(count / 3600));
	const m = $derived(Math.floor((count - h * 3600) / 60));
	const s = $derived(count - h * 3600 - m * 60);
	const remainingTime = $derived(count + (attendeesRemaining - 1) * attendeeCountdown);
	const remainingTimePercent = $derived(Math.floor((remainingTime / countdown) * 100 * 100) / 100);

	function updateTimer() {
		now = Date.now();
	}

	function playDing() {
		var audio = new Audio('/media/ding.mp3');
		audio.play();
	}

	let interval = setInterval(updateTimer, 1000);
	$effect(() => {
		if (count === 0 && attendeesRemaining - 1 === 0) {
			clearInterval(interval);
			playDing();
		} else if (count == 0) {
			playDing();

			now = Date.now();
			end = now + attendeeCountdown * 1000;

			attendeesRemaining = attendeesRemaining - 1;
		}
	});

	let isPaused = $state(false);
	let isResetting = $state(false);
	const duration = 1000;

	let offset = tweened(1, { duration, easing });
	let rotation = tweened(360, { duration, easing });

	$effect(() => {
		offset.set(Math.max(count - 1, 0) / attendeeCountdown);
	});
	$effect(() => {
		rotation.set((Math.max(count - 1, 0) / attendeeCountdown) * 360);
	});

	function handleNew() {
		clearInterval(interval);
		dispatch('new');
	}

	function handleStart() {
		let tmpCount = count;
		now = Date.now();
		end = now + tmpCount * 1000;
		interval = setInterval(updateTimer, 1000);
		offset.set(Math.max(count - 1, 0) / attendeeCountdown);
		rotation.set((Math.max(count - 1, 0) / attendeeCountdown) * 360);
		isPaused = false;
	}

	function handlePause() {
		offset.set(count / attendeeCountdown);
		rotation.set((count / attendeeCountdown) * 360);
		clearInterval(interval);
		isPaused = true;
	}

	function handleReset() {
		clearInterval(interval);
		isResetting = true;
		isPaused = false;
		Promise.all([offset.set(1), rotation.set(360)]).then(() => {
			isResetting = false;
			now = Date.now();
			end = now + attendeeCountdown * 1000;
			interval = setInterval(updateTimer, 1000);
		});
	}

	function padValue(value: any, length = 2, char = '0') {
		const { length: currentLength } = value.toString();
		if (currentLength >= length) return value.toString();
		return `${char.repeat(length - currentLength)}${value}`;
	}

	onDestroy(() => {
		clearInterval(interval);
	});
</script>

<div class="grid grid-cols-none">
	<div class="grid-span-1 mx-auto">
		<svg
			class="max-w-[250px]"
			in:fly={{ y: -5 }}
			viewBox="-50 -50 100 100"
			width="250"
			height="250"
		>
			<title>Remaining seconds: {count}</title>
			<!-- <g fill="none" stroke="currentColor" stroke-width="2">
				<circle stroke="currentColor" r="46" />
				<path
					stroke="#3b82f6"
					d="M 0 -46 a 46 46 0 0 0 0 92 46 46 0 0 0 0 -92"
					pathLength="1"
					stroke-dasharray="1"
					stroke-dashoffset={$offset}
				/>
			</g> -->

			<g fill="none" stroke="currentColor" stroke-width="2">
				<circle stroke="currentColor" r="46" />
				<path
					stroke="#3b82f6"
					d="M 0 -46 a 46 46 0 0 0 0 92 46 46 0 0 0 0 -92"
					pathLength="1"
					stroke-dasharray="1"
					stroke-dashoffset={$offset}
				/>
			</g>

			<g fill="#3b82f6" stroke="none">
				<g transform="rotate({$rotation})">
					<g transform="translate(0 -46)">
						<circle r="4" />
					</g>
				</g>
			</g>

			<g fill="currentColor" text-anchor="middle" dominant-baseline="baseline" font-size="13">
				<text x="-3" y="6.5">
					{#each Object.entries({ h, m, s }) as [key, value], i}
						{#if attendeeCountdown >= 60 ** (2 - i)}
							<tspan dx="3" font-weight="bold">{padValue(value)}</tspan><tspan
								dx="0.5"
								font-size="7">{key}</tspan
							>
						{/if}
					{/each}
				</text>
			</g>
		</svg>
	</div>

	<div class="grid-span-1 mx-auto max-w-[250px]" in:fly={{ y: -10, delay: 120 }}>
		{#if isPaused}
			<button disabled={isResetting || count === 0} onclick={handleStart}>
				<Play />
			</button>
		{:else}
			<button disabled={isResetting || count === 0} onclick={handlePause}>
				<Pause />
			</button>
		{/if}
		<button class="pl-3" onclick={handleReset}><RotateCcw /></button>
	</div>

	<div class="grid-span-1 mx-auto max-w-[250px]">
		<div class="flex flex-row items-center">
			<svg class="mx-auto" width="100%" height="10">
				<!-- Background track -->
				<rect x="0" y="0" width="100%" height="10" fill="#ccc" />

				<!-- Progress bar -->
				<rect
					class="progress"
					x="0"
					y="0"
					width="{remainingTimePercent}%"
					height="10"
					stroke-width="1"
					fill="#3b82f6"
				/>
			</svg>

			<div class="pl-3">
				{remainingTime}
			</div>
		</div>
	</div>
</div>

<style>
	.progress {
		transition: width 1s linear;
	}
</style>
