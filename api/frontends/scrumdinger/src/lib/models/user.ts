interface user {
	sub: string;
	roles: string;
	name: string;
}

import { jwtDecode } from 'jwt-decode';

export class User {
	isLoggedIn = false;
	constructor() {
		let token = localStorage.getItem('token');
		if (token) {
			const payload = jwtDecode(token);
			let exp = payload['exp'];

			if (exp && Date.now() < exp * 1000) {
				this.isLoggedIn = true;
			}
		}
	}
}

export let user = new User();
