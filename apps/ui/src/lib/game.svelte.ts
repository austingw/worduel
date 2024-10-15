let gameStart = $state<boolean>(false);

export function getGameStart() {
	return gameStart;
}

export function setGameStart(bool: boolean) {
	gameStart = bool;
}
