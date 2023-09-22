<script>
	import { GeneratePassword } from './lib/wailsjs/go/main/App';
	let wordCount = 4; // Default word count
	let separator = '-';
	let includeNumbers = false;
	let capitalizeLetters = false;
	let passphrase = '';

	function generatePassphrase() {
		GeneratePassword(
			wordCount,
			separator,
			includeNumbers,
			capitalizeLetters
		).then((result) => {
			passphrase = result;
		});
	}

	function copyPassphraseToClipboard() {
		const passphraseElement = document.getElementById('passphrase');
		const passphrase = passphraseElement.innerText;
		navigator.clipboard.writeText(passphrase);
		alert('Copied to clipboard!');
	}
</script>

<main class="max-w-md mx-auto mt-10 p-6 border rounded-lg shadow-lg bg-white">
	<h1 class="text-xl font-semibold mb-4">Passphrase Generator</h1>

	<label class="block mb-2">
		Number of Words:
		<input
			type="number"
			bind:value={wordCount}
			class="border rounded-lg px-2 py-1"
			min="1"
			max="10"
		/>
	</label>

	<label class="block mb-2">
		Word Separator Symbol:
		<input
			type="text"
			bind:value={separator}
			class="border rounded-lg px-2 py-1"
		/>
	</label>

	<label class="block mb-2">
		<input type="checkbox" bind:checked={includeNumbers} class="mr-2" />
		Include Numbers
	</label>

	<label class="block mb-4">
		<input type="checkbox" bind:checked={capitalizeLetters} class="mr-2" />
		Capitalize Letters
	</label>

	<button
		on:click={generatePassphrase}
		class="bg-blue-500 text-white px-4 py-2 rounded-lg hover:bg-blue-600"
	>
		Generate Passphrase
	</button>

	<div class="mt-6">
		<strong>Generated Passphrase:</strong>
		<p
			id="passphrase"
			class="border rounded-lg px-2 py-1 mt-2 bg-gray-100 h-auto"
			on:click={copyPassphraseToClipboard}
			style="cursor: pointer;"
		>
			{passphrase}
		</p>
	</div>
</main>

<style>
	/* You can add more Tailwind CSS styles or customize as needed */
</style>
