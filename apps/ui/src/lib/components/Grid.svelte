<script lang="ts">
	type GridProps = {
		letters: string[];
		attempts: string[];
		answer: string;
	};

	type parsedAttempt = {
		val: string;
		class: string;
	};

	let { letters, attempts, answer }: GridProps = $props();
	let blanks = $derived.by<string[]>(() => {
		if (letters.length < 5) {
			return Array(5 - letters.length).fill('');
		}
		return [];
	});

	let parsedAttempts = $derived.by<parsedAttempt[][]>(() => {
		console.log('letters', answer);
		const answerMap = new Map<string, number>();
		const parsedAttempts: parsedAttempt[][] = [];

		for (let i = 0; i < answer.length; i++) {
			answerMap.set(answer[i], (answerMap.get(answer[i]) || 0) + 1);
		}

		for (let attempt of attempts) {
			const clonedAnswerMap = new Map(answerMap);
			const parsedAttempt: parsedAttempt[] = [];

			for (let i = 0; i < attempt.length; i++) {
				parsedAttempt.push({ val: attempt[i], class: 'badge-neutral' });
			}

			for (let i = 0; i < attempt.length; i++) {
				let count = clonedAnswerMap.get(attempt[i]) || 0;
				if (answer[i] == attempt[i]) {
					parsedAttempt[i].class = 'badge-success';
					clonedAnswerMap.set(attempt[i], count - 1);
				}
			}

			for (let i = 0; i < attempt.length; i++) {
				let count = clonedAnswerMap.get(attempt[i]) || 0;
				if (answer.includes(attempt[i]) && count > 0) {
					parsedAttempt[i].class = 'badge-warning';
					clonedAnswerMap.set(attempt[i], count - 1);
				}
			}
			parsedAttempts.push(parsedAttempt);
		}

		return parsedAttempts;
	});
</script>

<div class="flex flex-col items-center justify-start w-full h-fit gap-2">
	{#each parsedAttempts as attempt}
		<div class="flex flex-row items-center justify-center gap-2">
			{#each attempt as letter}
				<div class={`badge badge-lg ${letter.class} text-lg w-12 h-12`}>
					{letter.val}
				</div>
			{/each}
		</div>
	{/each}
	<div class="flex flow-row items-center justify-center gap-2">
		{#each letters as letter}
			<div class="badge badge-lg text-lg w-12 h-12">{letter}</div>
		{/each}
		{#each blanks as blank}
			<div class="badge badge-lg text-lg w-12 h-12">{blank}</div>
		{/each}
	</div>
</div>
