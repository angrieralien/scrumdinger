import type { ScrumMeeting } from '$lib/models/scrum.svelte';
import { getUserContext, User } from '$lib/models/user.svelte';

let version = 'v1';

class ScrumApi {
	constructor() {}
	async GET() {
		let response = await fetch('/' + version + '/scrums', {
			method: 'GET',
			headers: {
				Authorization: 'Bearer ' + localStorage.getItem('token')
			}
		});

		if (response.status != 200) {
			let data = await response.json();
			return Promise.reject(data);
		} else {
			return response.json();
		}
	}

	async POST(data: string) {
		let response = await fetch('/' + version + '/scrums', {
			method: 'POST',
			body: data,
			headers: {
				Authorization: 'Bearer ' + localStorage.getItem('token')
			}
		});

		if (response.status != 200) {
			let data = await response.json();
			return Promise.reject(data);
		} else {
			return response.json();
		}
	}
}

export const scrumAPI = new ScrumApi();
