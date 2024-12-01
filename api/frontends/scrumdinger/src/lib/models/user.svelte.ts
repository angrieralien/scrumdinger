import { getContext, setContext } from 'svelte';
import { jwtDecode } from 'jwt-decode';

let userKey = Symbol('user');

export function setUserContext(user: User) {
	setContext(userKey, user);
}

export function getUserContext(): User {
	return getContext(userKey) as User;
}

interface user {
	sub: string;
	roles: string;
	name: string;
}

export class User {
	isLoggedIn = $state(false);

	token: string | null = $state('');
	date = $state(Date.now());
	constructor() {
		this.token = localStorage.getItem('token');
		if (this.token) {
			const payload = jwtDecode(this.token);
			let exp = payload['exp'];

			if (exp && Date.now() < exp * 1000) {
				this.isLoggedIn = true;
			}
		}
	}
}
