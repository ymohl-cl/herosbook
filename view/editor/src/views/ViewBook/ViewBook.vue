<template>
	<v-container fill-height fluid class="grey lighten-5 pa-0">
		<v-dialog v-model="openCategory" persistent max-width="1200px">
			<v-card>
				<v-card-title>
					<span class="headline">New category</span>
				</v-card-title>
				<v-card-text>
					<v-container>
						<v-row>
							<v-col cols="12" sm="6" md="4">
								<v-text-field v-model="category.title" label="Title" placeholder="Title" required solo></v-text-field>
							</v-col>
						</v-row>
						<v-row>
							<v-col cols="12" sm="6" md="4">
								<v-textarea v-model="category.description" solo placeholder="Description"></v-textarea>
							</v-col>
						</v-row>
					</v-container>
					<small>*indicates required field</small>
				</v-card-text>
				<v-card-actions>
					<v-spacer></v-spacer>
						<v-btn color="blue darken-1" text @click="closeDialog()">Cancel</v-btn>
						<v-btn color="blue darken-1" text @click="validDialog()">Save</v-btn>
				</v-card-actions>
			</v-card>
		</v-dialog>
		<v-layout row fill-height no-gutters fluid>
			<v-col cols="2" outlined>
				<v-card height="100%">
					<v-expansion-panels
						v-model="panel"
						multiple
					>
						<v-expansion-panel>
							<v-expansion-panel-header>Persons
								<template v-slot:actions>
									<v-btn @click.native.stop @click="openDialog('person')" icon small text>
										<v-icon dark color="orange">mdi-plus</v-icon>
									</v-btn>
								</template>
							</v-expansion-panel-header>
							<v-expansion-panel-content>
								<v-list three-line>
									<template v-for="(category, index) in book.categories.persons">
										<v-divider
											v-if="index > 0"
											:key="category.identifier-divider"
										></v-divider>
										<v-list-item :key="category.identifier-list">
											<v-list-item-avatar>
												<v-img :src="category.img"></v-img>
											</v-list-item-avatar>
											<v-list-item-content>
												<v-list-item-title v-html="category.title"></v-list-item-title>
											</v-list-item-content>
										</v-list-item>
									</template>
								</v-list>
							</v-expansion-panel-content>
						</v-expansion-panel>
						<v-expansion-panel>
							<v-expansion-panel-header>Location
								<template v-slot:actions>
									<v-btn @click.native.stop @click="openDialog('location')" icon small text>
										<v-icon dark color="orange">mdi-plus</v-icon>
									</v-btn>
								</template>
							</v-expansion-panel-header>
							<v-expansion-panel-content>
								<v-list three-line>
									<template v-for="(category, index) in book.categories.locations">
										<v-divider
											v-if="index > 0"
											:key="category.identifier+'divider'"
										></v-divider>
										<v-list-item :key="category.identifier+'list'">
											<v-list-item-avatar>
												<v-img :src="category.img"></v-img>
											</v-list-item-avatar>
											<v-list-item-content>
												<v-list-item-title v-html="category.title"></v-list-item-title>
											</v-list-item-content>
										</v-list-item>
									</template>
								</v-list>
							</v-expansion-panel-content>
						</v-expansion-panel>
						<v-expansion-panel>
							<v-expansion-panel-header>Tags
								<template v-slot:actions>
									<v-btn @click.native.stop @click="openDialog('custom')" icon small text>
										<v-icon dark color="orange">mdi-plus</v-icon>
									</v-btn>
								</template>
							</v-expansion-panel-header>
							<v-expansion-panel-content>
								<v-list three-line>
									<template v-for="(category, index) in book.categories.customs">
										<v-divider
											v-if="index > 0"
											:key="category.identifier"
										></v-divider>
										<v-list-item :key="category.identifier">
											<v-list-item-avatar>
												<v-img :src="category.img"></v-img>
											</v-list-item-avatar>
											<v-list-item-content>
												<v-list-item-title v-html="category.title"></v-list-item-title>
											</v-list-item-content>
										</v-list-item>
									</template>
								</v-list>
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
					<!--<v-expansion-panels
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
					</v-expansion-panels>-->
				</v-card>
			</v-col>
		</v-layout>
	</v-container>
</template>
<script src="./ViewBook.ts"></script>
