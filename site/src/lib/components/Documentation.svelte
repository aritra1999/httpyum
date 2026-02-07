<script lang="ts">
	import { Copy, Check } from "@lucide/svelte";
	import { Button } from "$lib/components/ui/button/index.js";
	import * as Tabs from "$lib/components/ui/tabs/index.js";

	interface Shortcut {
		keys: string[];
		description: string;
	}

	interface Props {
		shortcuts: Shortcut[];
		copied?: boolean;
		onCopy?: () => void;
	}

	let { shortcuts, copied = false, onCopy }: Props = $props();
</script>

<section
	class="py-16 px-8 bg-white border-b border-[#e5e5e5] max-md:py-12 max-md:px-4"
	id="docs"
>
	<div class="container max-w-[1100px] mx-auto px-8 max-md:px-4">
		<h2
			class="text-[2.5rem] font-bold m-0 mb-10 -tracking-[0.02em] lowercase text-black relative inline-block after:content-[''] after:absolute after:-bottom-2 after:left-0 after:w-1/2 after:h-1 after:bg-black max-md:text-[1.75rem] max-md:mb-6 max-md:after:w-2/5 max-[480px]:text-2xl"
		>
			Documentation
		</h2>

		<Tabs.Root value="getting-started" class="docs-interactive">
			<div
				class="w-full overflow-hidden bg-[#f5f5f5] border border-[#e5e5e5] rounded-lg p-1 shadow-[0_1px_3px_rgba(0,0,0,0.05)] max-md:m-0 max-md:overflow-x-auto max-md:-webkit-overflow-scrolling-touch"
			>
				<Tabs.List
					class="docs-tabs overflow-x-auto overflow-y-hidden -webkit-overflow-scrolling-touch scrollbar-thin flex-nowrap !w-full !max-w-full !bg-transparent !p-0 [&::-webkit-scrollbar]:h-[3px] [&::-webkit-scrollbar-track]:bg-[#f5f5f5] [&::-webkit-scrollbar-thumb]:bg-[#999] [&::-webkit-scrollbar-thumb]:rounded max-md:!justify-start max-md:!w-auto max-md:!min-w-full"
				>
					<Tabs.Trigger
						value="getting-started"
						class="doc-tab !rounded-lg !border-none !bg-transparent !font-medium !transition-all !duration-200 !text-sm !py-8 !px-7 whitespace-nowrap !flex-shrink-0 !min-w-fit !text-[#666] data-[state=active]:!bg-white data-[state=active]:!text-black data-[state=active]:!font-semibold data-[state=active]:!shadow-[0_1px_2px_rgba(0,0,0,0.05)] hover:!bg-white/50 hover:!text-black max-md:!text-[0.8125rem] max-md:!py-3 max-md:!px-4 max-[480px]:!text-xs max-[480px]:!py-2 max-[480px]:!px-2.5"
					>
						Getting Started
					</Tabs.Trigger>
					<Tabs.Trigger
						value="file-format"
						class="doc-tab !rounded-lg !border-none !bg-transparent !font-medium !transition-all !duration-200 !text-sm !py-8 !px-7 whitespace-nowrap !flex-shrink-0 !min-w-fit !text-[#666] data-[state=active]:!bg-white data-[state=active]:!text-black data-[state=active]:!font-semibold data-[state=active]:!shadow-[0_1px_2px_rgba(0,0,0,0.05)] hover:!bg-white/50 hover:!text-black max-md:!text-[0.8125rem] max-md:!py-3 max-md:!px-4 max-[480px]:!text-xs max-[480px]:!py-2 max-[480px]:!px-2.5"
					>
						File Format
					</Tabs.Trigger>
					<Tabs.Trigger
						value="variables"
						class="doc-tab !rounded-lg !border-none !bg-transparent !font-medium !transition-all !duration-200 !text-sm !py-8 !px-7 whitespace-nowrap !flex-shrink-0 !min-w-fit !text-[#666] data-[state=active]:!bg-white data-[state=active]:!text-black data-[state=active]:!font-semibold data-[state=active]:!shadow-[0_1px_2px_rgba(0,0,0,0.05)] hover:!bg-white/50 hover:!text-black max-md:!text-[0.8125rem] max-md:!py-3 max-md:!px-4 max-[480px]:!text-xs max-[480px]:!py-2 max-[480px]:!px-2.5"
					>
						Variables
					</Tabs.Trigger>
					<Tabs.Trigger
						value="json-viewer"
						class="doc-tab !rounded-lg !border-none !bg-transparent !font-medium !transition-all !duration-200 !text-sm !py-8 !px-7 whitespace-nowrap !flex-shrink-0 !min-w-fit !text-[#666] data-[state=active]:!bg-white data-[state=active]:!text-black data-[state=active]:!font-semibold data-[state=active]:!shadow-[0_1px_2px_rgba(0,0,0,0.05)] hover:!bg-white/50 hover:!text-black max-md:!text-[0.8125rem] max-md:!py-3 max-md:!px-4 max-[480px]:!text-xs max-[480px]:!py-2 max-[480px]:!px-2.5"
					>
						JSON Viewer
					</Tabs.Trigger>
					<Tabs.Trigger
						value="shortcuts"
						class="doc-tab !rounded-lg !border-none !bg-transparent !font-medium !transition-all !duration-200 !text-sm !py-8 !px-7 whitespace-nowrap !flex-shrink-0 !min-w-fit !text-[#666] data-[state=active]:!bg-white data-[state=active]:!text-black data-[state=active]:!font-semibold data-[state=active]:!shadow-[0_1px_2px_rgba(0,0,0,0.05)] hover:!bg-white/50 hover:!text-black max-md:!text-[0.8125rem] max-md:!py-3 max-md:!px-4 max-[480px]:!text-xs max-[480px]:!py-2 max-[480px]:!px-2.5"
					>
						Keyboard Shortcuts
					</Tabs.Trigger>
				</Tabs.List>
			</div>

			<Tabs.Content
				value="getting-started"
				class="doc-panel !bg-[#f5f5f5] !border !border-[#e5e5e5] !rounded-lg !p-8 animate-[fadeIn_0.3s_ease-in-out]"
			>
				<h3 class="text-2xl font-bold m-0 mb-6 max-md:text-xl">
					Getting Started
				</h3>
				<div
					class="doc-content [&_p]:m-0 [&_p]:mb-4 [&_p]:text-[#666] [&_p]:text-[0.9375rem]"
				>
					<h4
						class="text-lg font-bold my-8 mt-0 text-black max-md:text-base max-md:my-6 max-md:mt-0"
					>
						Installation
					</h4>
					<p>Install httpyum using the installation script:</p>
					<div
						class="bg-white border border-[#e5e5e5] rounded p-6 my-4 relative max-md:p-3"
					>
						<pre
							class="font-['JetBrains_Mono',monospace] text-sm leading-relaxed text-[#333] m-0 overflow-x-auto whitespace-pre-wrap break-all max-md:text-[0.8125rem]">curl -fsSL https://raw.githubusercontent.com/aritra1999/httpyum/main/scripts/install.sh | bash</pre>
						<Button
							variant="ghost"
							size="sm"
							class="copy-btn !absolute !top-4 !right-4 !bg-white !text-black !border !border-[#e5e5e5] !rounded max-md:!top-3 max-md:!right-3 hover:!bg-[#f5f5f5] hover:!border-[#999]"
							onclick={onCopy}
						>
							{#if copied}
								<Check size={16} />
							{:else}
								<Copy size={16} />
							{/if}
						</Button>
					</div>
					<div class="flex gap-3 my-4 mb-6 max-md:flex-wrap">
						<span
							class="py-1.5 px-3 bg-[#f5f5f5] border border-[#e5e5e5] rounded text-[0.8125rem] font-medium text-[#666]"
							>macOS</span
						>
						<span
							class="py-1.5 px-3 bg-[#f5f5f5] border border-[#e5e5e5] rounded text-[0.8125rem] font-medium text-[#666]"
							>Linux</span
						>
						<span
							class="py-1.5 px-3 bg-[#f5f5f5] border border-[#e5e5e5] rounded text-[0.8125rem] font-medium text-[#666]"
							>Windows</span
						>
					</div>

					<h4
						class="text-lg font-bold my-8 text-black max-md:text-base max-md:my-6"
					>
						Usage
					</h4>
					<p>Run httpyum with any .http file:</p>
					<pre
						class="font-['JetBrains_Mono',monospace] text-sm leading-relaxed bg-white text-[#333] border border-[#e5e5e5] rounded p-4 my-4 overflow-x-auto block max-md:text-[0.8125rem] max-md:p-3">httpyum example.http</pre>
					<p>Command-line options:</p>
					<ul
						class="my-4 pl-6 text-[#666] text-[0.9375rem] [&_li]:my-2 [&_code]:font-['JetBrains_Mono',monospace] [&_code]:text-[0.875em] [&_code]:bg-white [&_code]:text-[#333] [&_code]:py-0.5 [&_code]:px-1.5 [&_code]:rounded-sm [&_code]:border [&_code]:border-[#e5e5e5]"
					>
						<li>
							<code>--no-headers</code> - Hide response headers
						</li>
						<li>
							<code>-h, --help</code> - Show help message
						</li>
						<li>
							<code>-v, --version</code> - Show version
						</li>
					</ul>
				</div>
			</Tabs.Content>

			<Tabs.Content
				value="file-format"
				class="doc-panel !bg-[#f5f5f5] !border !border-[#e5e5e5] !rounded-lg !p-8 animate-[fadeIn_0.3s_ease-in-out]"
			>
				<h3 class="text-2xl font-bold m-0 mb-6 max-md:text-xl">
					File Format
				</h3>
				<div
					class="doc-content [&_p]:m-0 [&_p]:mb-4 [&_p]:text-[#666] [&_p]:text-[0.9375rem]"
				>
					<p>Standard .http file format with variables:</p>
					<pre
						class="font-['JetBrains_Mono',monospace] text-sm leading-relaxed bg-white text-[#333] border border-[#e5e5e5] rounded p-4 my-4 overflow-x-auto block max-md:text-[0.8125rem] max-md:p-3">@variable = value

### Request Name
METHOD url
Header: value

body</pre>
				</div>
			</Tabs.Content>

			<Tabs.Content
				value="variables"
				class="doc-panel !bg-[#f5f5f5] !border !border-[#e5e5e5] !rounded-lg !p-8 animate-[fadeIn_0.3s_ease-in-out]"
			>
				<h3 class="text-2xl font-bold m-0 mb-6 max-md:text-xl">
					Variables
				</h3>
				<div
					class="doc-content [&_p]:m-0 [&_p]:mb-4 [&_p]:text-[#666] [&_p]:text-[0.9375rem] [&_code]:font-['JetBrains_Mono',monospace] [&_code]:text-[0.875em] [&_code]:bg-white [&_code]:text-[#333] [&_code]:py-0.5 [&_code]:px-1.5 [&_code]:rounded-sm [&_code]:border [&_code]:border-[#e5e5e5]"
				>
					<p>
						Define variables with <code>@</code> and use with
						<code>{"{{"}{"}}"}</code>:
					</p>
					<pre
						class="font-['JetBrains_Mono',monospace] text-sm leading-relaxed bg-white text-[#333] border border-[#e5e5e5] rounded p-4 my-4 overflow-x-auto block max-md:text-[0.8125rem] max-md:p-3">@baseUrl = https://api.example.com
GET {"{{"}baseUrl{"}}"}/users</pre>
					<p>Load from environment:</p>
					<pre
						class="font-['JetBrains_Mono',monospace] text-sm leading-relaxed bg-white text-[#333] border border-[#e5e5e5] rounded p-4 my-4 overflow-x-auto block max-md:text-[0.8125rem] max-md:p-3">@token = {"{{"}$dotenv JWT{"}}"}</pre>
				</div>
			</Tabs.Content>

			<Tabs.Content
				value="json-viewer"
				class="doc-panel !bg-[#f5f5f5] !border !border-[#e5e5e5] !rounded-lg !p-8 animate-[fadeIn_0.3s_ease-in-out]"
			>
				<h3 class="text-2xl font-bold m-0 mb-6 max-md:text-xl">
					JSON Viewer
				</h3>
				<div
					class="doc-content [&_p]:m-0 [&_p]:mb-4 [&_p]:text-[#666] [&_p]:text-[0.9375rem]"
				>
					<p>
						Press <kbd
							class="inline-flex items-center justify-center min-w-[28px] py-1.5 px-2 bg-[#f5f5f5] text-[#333] border border-[#e5e5e5] rounded font-['JetBrains_Mono',monospace] text-[0.8125rem] font-semibold not-italic"
							>f</kbd
						> to open JSON responses in jless for interactive exploration.
					</p>
					<p>Install jless for the best experience:</p>
					<pre
						class="font-['JetBrains_Mono',monospace] text-sm leading-relaxed bg-white text-[#333] border border-[#e5e5e5] rounded p-4 my-4 overflow-x-auto block max-md:text-[0.8125rem] max-md:p-3"># macOS
brew install jless

# Linux
cargo install jless</pre>
				</div>
			</Tabs.Content>

			<Tabs.Content
				value="shortcuts"
				class="doc-panel !bg-[#f5f5f5] !border !border-[#e5e5e5] !rounded-lg !p-8 animate-[fadeIn_0.3s_ease-in-out]"
			>
				<h3 class="text-2xl font-bold m-0 mb-6 max-md:text-xl">
					Keyboard Shortcuts
				</h3>
				<div
					class="grid grid-cols-[repeat(auto-fit,minmax(250px,1fr))] gap-6 max-md:grid-cols-1"
				>
					{#each shortcuts as shortcut}
						<div
							class="flex items-center gap-4 p-4 bg-white border border-[#e5e5e5] rounded-md transition-all duration-200"
						>
							<div class="flex gap-2">
								{#each shortcut.keys as key}
									<kbd
										class="inline-flex items-center justify-center min-w-[28px] py-1 px-2 bg-[#f5f5f5] text-[#333] border border-[#e5e5e5] rounded font-['JetBrains_Mono',monospace] text-[0.8125rem] font-semibold not-italic"
										>{key}</kbd
									>
								{/each}
							</div>
							<div class="text-[0.9375rem] text-[#666]">
								{shortcut.description}
							</div>
						</div>
					{/each}
				</div>
			</Tabs.Content>
		</Tabs.Root>
	</div>
</section>

<style>
	@keyframes fadeIn {
		from {
			opacity: 0;
			transform: translateY(10px);
		}
		to {
			opacity: 1;
			transform: translateY(0);
		}
	}
</style>
