<script lang="ts">
	let isNavMenuOpen = $state(false);

	function toggleNavMenu() {
		isNavMenuOpen = !isNavMenuOpen;
	}
</script>

<header class="bg-bg fixed shadow-xl top-0 left-0 w-full flex items-center z-10">
	<div class="container text-2xl">
		<div class="flex justify-between items-center relative">
			<div class="px-4">
				<a href="#home" class="text-3xl font-bold text-primary block my-6">Portfolio</a>
			</div>

			<div class="flex items-center px-4">
				<button on:click={toggleNavMenu} id="hamburger" name="hamburger" type="button" class="absolute right-4 block lg:hidden {isNavMenuOpen ? 'hamburger-active' : ''}">
					<span class="hamburger-line origin-top-left transition duration-200 ease-in-out" />
					<span class="hamburger-line transition duration-200 ease-in-out" />
					<span class="hamburger-line origin-bottom-left transition duration-200 ease-in-out" />
				</button>

				<nav
					id="nav-menu"
					class="bg-bg absolute py-5 rounded-xl shadow-lg max-w-[250px] w-full right-4 top-full lg:static lg:block lg:max-w-full lg:rounded-none lg:shadow-none {isNavMenuOpen ? '' : 'hidden'}"
				>
					<ul class="block lg:flex">
						<li class="group">
							<a href="#home" class="text-lg font-medium text-dark py-2 mx-8 flex group-hover:text-primary">Home</a>
						</li>
						<li class="group">
							<a href="#about" class="text-lg font-medium text-dark py-2 mx-8 flex group-hover:text-primary">About</a>
						</li>
						<li class="group">
							<a href="#services" class="text-lg font-medium text-dark py-2 mx-8 flex group-hover:text-primary">Services</a>
						</li>
						<li class="group">
							<a href="#portfolio" class="text-lg font-medium text-dark py-2 mx-8 flex group-hover:text-primary">Portfolio</a>
						</li>
					</ul>
				</nav>
			</div>
		</div>
	</div>
</header>