import { getContext, setContext } from 'svelte';

let scrumsKey = Symbol('scrum');

export function setScrumsContext(scrums: ScrumMeeting[]) {
	setContext(scrumsKey, scrums);
}

export function getScrumsContext(): ScrumMeeting[] {
	return getContext(scrumsKey) as ScrumMeeting[];
}

export class ScrumMeeting {
	name: string = $state('Project A');
	time: number = $state(5);
	attendees: string[] = $state([]);
}
