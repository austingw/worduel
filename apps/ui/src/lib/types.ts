export type View = 'start' | 'list' | 'room';

export type User = {
	name: string;
	score: number;
};

export type Room = {
	name: string;
	users: User[];
	currentWord: string;
};
