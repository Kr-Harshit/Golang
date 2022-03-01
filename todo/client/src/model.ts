export interface Todo {
	id: number;
	todo: string;
	deadline: Date;
	isDone: boolean;
}

export enum ActionsType {
	GET,
	ADD,
	DELETE,
	TOGGLE,
	UPDATE,
}

export type Actions =
	| { type: ActionsType.GET; payload: object }
	| { type: ActionsType.ADD; payload: { todo: string; deadline: Date } }
	| { type: ActionsType.DELETE; payload: number }
	| { type: ActionsType.TOGGLE; payload: number }
	| { type: ActionsType.UPDATE; payload: { id: Number; todo: string } };
