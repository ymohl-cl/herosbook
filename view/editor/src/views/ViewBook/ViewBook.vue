<template>
	<v-container fill-height fluid class="grey lighten-5 pa-0">
		<v-layout row fill-height no-gutters fluid>
			<v-col cols="2" outlined>
				<v-card height="100%">
					<v-expansion-panels
						v-model="panel"
						multiple
					>
						<v-expansion-panel>
							<v-expansion-panel-header>Persons</v-expansion-panel-header>
							<v-expansion-panel-content>
								Marion
							</v-expansion-panel-content>
						</v-expansion-panel>
						<v-expansion-panel>
							<v-expansion-panel-header>Location</v-expansion-panel-header>
							<v-expansion-panel-content>
								Blackwood
							</v-expansion-panel-content>
						</v-expansion-panel>
						<v-expansion-panel>
							<v-expansion-panel-header>Tags</v-expansion-panel-header>
							<v-expansion-panel-content>
								Bloodest
							</v-expansion-panel-content>
						</v-expansion-panel>
					</v-expansion-panels>
				</v-card>	
			</v-col>
			<v-col>
				<v-row v-for="(nodes, i) in treeNodes" :key="i" justify="center">
					<v-col v-for="(node, j) in nodes" :key="j" cols="2" md="2">
						<v-card class="pa-2" color="#385F73" dark>
							<v-card-title class="headline">{{ node.title }}</v-card-title>
							<v-card-subtitle>{{ node.description }}</v-card-subtitle>
							<v-card-actions>
								<v-btn text>Edit Now</v-btn>
							</v-card-actions>
						</v-card>
					</v-col>
				</v-row>
			</v-col>
			<v-col cols="2" outlined>
				<v-card height="100%">
					<v-expansion-panels
						:accordion="accordion"
						:popout="popout"
						:inset="inset"
						:multiple="multiple"
						:focusable="focusable"
						:disabled="disabled"
						:readonly="readonly"
						:flat="flat"
						:hover="hover"
						:tile="tile"
					>
						<v-expansion-panel
							v-for="(node, i) in orphanNodes"
							:key="i"
						>
						<v-expansion-panel-header>{{ node.title }}</v-expansion-panel-header>
						<v-expansion-panel-content>
							{{ node.description }}
						</v-expansion-panel-content>
						</v-expansion-panel>
					</v-expansion-panels>
				</v-card>
			</v-col>
		</v-layout>
	</v-container>
</template>
<script src="./ViewBook.ts"></script>
