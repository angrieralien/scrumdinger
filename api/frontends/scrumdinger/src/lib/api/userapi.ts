import { PUBLIC_KID } from '$env/static/public';

let version = 'v1';

class UserApi {
	constructor() {}

	/**
	 * login uses basic auth to login to scrumdinger.
	 * Successful a jwt token is returned on successful login.
	 *
	 * @param email users email address
	 * @param password users password
	 * @returns Promise<string> json
	 */
	async login(email: string, password: string) {
		let response = await fetch('/' + version + '/auth/token/' + PUBLIC_KID, {
			method: 'GET',
			headers: {
				authorization: 'Basic ' + btoa(email + ':' + password)
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

export const userAPI = new UserApi();
