import { getContext, setContext } from 'svelte';

let drawerKey = Symbol('drawer');

export function setDrawerContext(meta: DrawerMeta) {
	setContext(drawerKey, meta);
}

export function getDrawerContext(): DrawerMeta {
	return getContext(drawerKey) as DrawerMeta;
}

export class DrawerMeta {
	data: any = {};
}
