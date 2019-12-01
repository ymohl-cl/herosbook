<template>
	<div>
	<v-layout>
	<v-app-bar color="orange" dark absolute>
		<v-toolbar-title>Hereos Book</v-toolbar-title>
		<v-spacer></v-spacer>
		<div v-if="connected">
			<v-toolbar-title class="d-inline">{{ pseudo }}</v-toolbar-title>
			<v-menu left bottom>
				<template v-slot:activator="{ on }">
					<v-btn icon v-on="on">
						<v-icon>mdi-dots-vertical</v-icon>
					</v-btn>
				</template>
				<v-list>
					<v-list-item v-for="item in menu" :key="item.title" @click="item.f()">
						<v-list-item-title>{{ item.title }}</v-list-item-title>
					</v-list-item>
				</v-list>
			</v-menu>
		</div>
		<div v-else>
				<v-btn class="ma-2 orange--text" color="white" dark @click="register()">Register</v-btn>
				<v-btn class="ma-2 orange--text" color="white" dark @click="login()">login</v-btn>
		</div>
	</v-app-bar>
	</v-layout>
	<v-container></v-container><!--TODO: replace all v-container by spacer system with size fixed -->
	<v-container></v-container>
	<v-container></v-container>
	<v-container></v-container>
	</div>
</template>
<script src="./BaseHeader.ts"></script>
