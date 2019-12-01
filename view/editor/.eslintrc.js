module.exports = {
	root: true,
	env: {
		node: true,
	},
	extends: [
		"plugin:vue/essential",
		"@vue/airbnb",
		"@vue/typescript",
	],
	rules: {
		"no-console": process.env.NODE_ENV === "production" ? "error" : "off",
		"no-debugger": process.env.NODE_ENV === "production" ? "error" : "off",
		"class-methods-use-this": "off",
		"max-len": ["error", { code: 120 }],
		indent: ["error", "tab"],
		"no-tabs": 0,
		"no-underscore-dangle": "off",
		"no-multi-spaces": ["error"],
		semi: ["error", "never"],
		"no-plusplus": ["error", { allowForLoopAfterthoughts: true }],
		quotes: ["error", "double", { avoidEscape: true }],
		"lines-between-class-members": ["error", "always", { exceptAfterSingleLine: true }],
		"padding-line-between-statements": ["error",
			{ blankLine: "always", prev: ["const", "let", "var"], next: "*" },
			{ blankLine: "any", prev: ["const", "let", "var"], next: ["const", "let", "var"] }],
		"vue/no-tabs": 0,
	},
	parserOptions: {
		parser: "@typescript-eslint/parser",
	},
}
