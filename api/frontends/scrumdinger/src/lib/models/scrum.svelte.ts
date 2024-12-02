import { getContext, setContext } from 'svelte';

let scrumsKey = Symbol('scrum');

export function setScrumContext(scrums: Scrum) {
	setContext(scrumsKey, scrums);
}

export function getScrumContext(): Scrum {
	return getContext(scrumsKey) as Scrum;
}

export class APIScrumMeetings {
	items: Array<APIScrumMeeting> = [];
	total: number = 0;
	page: number = 1;
	rowsPerPage: number = 1;
}

export class APIScrumMeeting {
	id: string = '';
	name: string = '';
	time: number = 0;
	attendees: string[] = [];
	color: string = '';
}

export function toAppScrumMeetings(asms: APIScrumMeeting[]): ScrumMeeting[] {
	let meetings: ScrumMeeting[] = [];
	asms.forEach((item) => {
		let meeting = toAppScrumMeeting(item);
		meetings.push(meeting);
	});
	return meetings;
}

export function toAppScrumMeeting(asm: APIScrumMeeting) {
	let scrum = new ScrumMeeting();
	scrum.id = asm.id;
	scrum.name = asm.name;
	scrum.time = asm.time;
	scrum.color = asm.color;
	scrum.attendees = asm.attendees;

	return scrum;
}

export class ScrumMeeting {
	id: string = '';
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

	fromJSON(obj: any) {
		this.id = obj.id;
		this.name = obj.name;
		this.time = obj.time;
		this.color = obj.color;
		this.attendees = obj.attendees;
	}
}

export class Scrum {
	meetings: ScrumMeeting[] = $state([]);
}
