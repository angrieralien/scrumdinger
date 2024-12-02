import { getContext, setContext } from 'svelte';

let scrumsKey = Symbol('scrum');

export function setScrumContext(scrums: Scrum) {
	setContext(scrumsKey, scrums);
}

export function getScrumContext(): Scrum {
	return getContext(scrumsKey) as Scrum;
}

export class ScrumMeeting {
	name: string = $state('Project A');
	time: number = $state(5);
	attendees: string[] = $state([]);
	color: string = $state('bg-primary-500');

	toJson() {
		return JSON.stringify({
			name: this.name,
			time: this.time,
			attendees: this.attendees,
			color: this.color
		});
	}
}

export class Scrum {
	meetings: ScrumMeeting[] = $state([]);
}
