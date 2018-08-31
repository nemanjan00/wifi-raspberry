<template>
	<div>
		<b-navbar toggleable="md" type="dark" variant="dark">

			<b-navbar-toggle target="nav_collapse"></b-navbar-toggle>

			<b-navbar-brand href="#">WiFi Raspberry</b-navbar-brand>

			<b-collapse is-nav id="nav_collapse">

				<b-navbar-nav>
				</b-navbar-nav>

				<!-- Right aligned nav items -->
				<b-navbar-nav class="ml-auto">
					<b-nav-form>
						<b-form-input size="sm" class="mr-sm-2" type="text" placeholder="Search"/>
						<b-button size="sm" class="my-2 my-sm-0" type="submit">Search</b-button>
					</b-nav-form>
					<b-nav-item href="https://github.com/nemanjan00/wifi-raspberry" target="blank"><i class="fab fa-github"></i> GitHub</b-nav-item>
				</b-navbar-nav>

			</b-collapse>
		</b-navbar>

		<div class="row">
			<div class="col-2" style="border-right: 1px solid rgba(0,0,0,.1);height: calc(100vh - 4rem);">
				<b-nav vertical>
					<b-nav-item active><i class="fas fa-cogs"></i> Advanced</b-nav-item>
					<b-nav-item><i class="fas fa-users"></i> Clients</b-nav-item>
					<b-nav-item><i class="fas fa-wrench"></i> Configuration</b-nav-item>
					<b-nav-item><i class="fas fa-columns"></i> Dashboard</b-nav-item>
					<b-nav-item><i class="fas fa-filter"></i> Filters</b-nav-item>
					<b-nav-item><i class="fas fa-question-circle"></i> Help</b-nav-item>
					<b-nav-item><i class="fas fa-list-ul"></i> Logging</b-nav-item>
					<b-nav-item><i class="fas fa-cubes"></i> Modules</b-nav-item>
					<b-nav-item><i class="fas fa-globe-asia"></i> Networking</b-nav-item>
					<b-nav-item><i class="fas fa-book"></i> Notes</b-nav-item>
					<b-nav-item><i class="fas fa-wifi"></i> PineAP</b-nav-item>
					<b-nav-item><i class="fas fa-search"></i> Recon</b-nav-item>
					<b-nav-item><i class="far fa-newspaper"></i> Reporting</b-nav-item>
					<b-nav-item><i class="fas fa-map"></i> Tracking</b-nav-item>
				</b-nav>
			</div>
		</div>
	</div>
</template>
