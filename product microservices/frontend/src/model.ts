export interface Product {
	id: number;
	name: string;
	description: string;
	sku: string;
	price: number;
}

export enum ActionType {
	GET,
	ADD,
	DELETE,
	TOGGLE,
	UPDATE,
}

export type Actions = { type: ActionType.GET; payload: object };
