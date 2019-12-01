<template>
	<v-app>
		<AlertMessage ref="AlertMessage"></AlertMessage>
		<BaseHeader
			v-bind:connected="session.user.isConnected()"
			v-bind:pseudo="session.user.pseudo"
			v-bind:buildMenu="menuBuilder()"
			v-bind:disconnectFunction="disconnect"
			v-bind:loginFunction="loginView"
			v-bind:registerFunction="registerView"
		></BaseHeader>
		<v-content>
			<router-view />
		</v-content>
	</v-app>
</template>
<script src="./App.ts"></script>
