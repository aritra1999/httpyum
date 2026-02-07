import adapter from "@sveltejs/adapter-static";

const dev = process.env.NODE_ENV === "development";

/** @type {import('@sveltejs/kit').Config} */
const config = {
	kit: {
		adapter: adapter({
			pages: "build",
			assets: "build",
			fallback: undefined,
			precompress: false,
			strict: true,
		}),
		paths: {
			base: dev ? "" : process.env.BASE_PATH || "",
		},
		alias: {
			"@/*": "./path/to/lib/*",
		},
	},
};

export default config;
