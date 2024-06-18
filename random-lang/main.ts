const languages = ["PHP", "JavaScript", "TypeScript", "Java", "Swift", "Go", "C++"];

// Learn more at https://deno.land/manual/examples/module_metadata#concepts
if (import.meta.main) {
	const randomIndex = Math.floor(Math.random() * languages.length);
	console.log(languages[randomIndex])
}
