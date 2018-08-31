<template>
	<div>
		<b-navbar toggleable="md" type="dark" variant="dark">

			<b-navbar-toggle target="nav_collapse"></b-navbar-toggle>

			<b-navbar-brand href="#">WiFi Raspberry</b-navbar-brand>

			<b-collapse is-nav id="nav_collapse">

				<b-navbar-nav>
					<b-nav-item href="#">Link</b-nav-item>
					<b-nav-item href="#" disabled>Disabled</b-nav-item>
				</b-navbar-nav>

				<!-- Right aligned nav items -->
				<b-navbar-nav class="ml-auto">

					<b-nav-form>
						<b-form-input size="sm" class="mr-sm-2" type="text" placeholder="Search"/>
						<b-button size="sm" class="my-2 my-sm-0" type="submit">Search</b-button>
					</b-nav-form>

					<b-nav-item-dropdown text="Lang" right>
						<b-dropdown-item href="#">EN</b-dropdown-item>
						<b-dropdown-item href="#">ES</b-dropdown-item>
						<b-dropdown-item href="#">RU</b-dropdown-item>
						<b-dropdown-item href="#">FA</b-dropdown-item>
					</b-nav-item-dropdown>

					<b-nav-item-dropdown right>
						<!-- Using button-content slot -->
						<template slot="button-content">
							<em>User</em>
						</template>
						<b-dropdown-item href="#">Profile</b-dropdown-item>
						<b-dropdown-item href="#">Signout</b-dropdown-item>
					</b-nav-item-dropdown>
				</b-navbar-nav>

			</b-collapse>
		</b-navbar>

		<div class="row">
			<div class="col-2" style="border-right: 1px solid rgba(255,255,255,.1);height: calc(100vh - 4rem);">
				<b-nav vertical>
					<b-nav-item active>Advanced</b-nav-item>
					<b-nav-item>Clients</b-nav-item>
					<b-nav-item>Configuration</b-nav-item>
					<b-nav-item>Dashboard</b-nav-item>
					<b-nav-item>Filters</b-nav-item>
					<b-nav-item>Help</b-nav-item>
					<b-nav-item>Logging</b-nav-item>
					<b-nav-item>Modules</b-nav-item>
					<b-nav-item>Networking</b-nav-item>
					<b-nav-item>Notes</b-nav-item>
					<b-nav-item>PineAP</b-nav-item>
					<b-nav-item>Recon</b-nav-item>
					<b-nav-item>Reporting</b-nav-item>
					<b-nav-item>Tracking</b-nav-item>
				</b-nav>
			</div>
		</div>
	</div>
</template>
