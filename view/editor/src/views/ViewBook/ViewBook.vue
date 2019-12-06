<template>
	<v-container>
		<div>
			<h1>{{ book.title }}</h1>
			<p>{{ book.description }}</p>
		</div>
		<div>
			<v-list v-if="nodes.length > 0">
				<v-subheader>Node list</v-subheader>
				<v-list-item-group color="primary">
					<v-list-item v-for="(node, i) in nodes" :key="i">
						<v-list-item-content>
							<v-list-item-title>{{ node.title }} (TODO : cliquez pour éditer)</v-list-item-title>
							<v-list-item-action>
								<v-btn outlined color="success" @click="currentNode = node">Edit</v-btn>
							</v-list-item-action>
						</v-list-item-content>
					</v-list-item>
					<v-list-item>
						<v-list-item-content>
							<v-list-item-title>
								<v-btn outlined color="success" @click="createNode()">Créer noeud</v-btn>
							</v-list-item-title>
						</v-list-item-content>
					</v-list-item>
				</v-list-item-group>
			</v-list>
			<div v-if="nodes.length == 0">
				Ce projet n'a pas encore de noeud, créez en un
				<v-btn outlined color="success" @click="createNode()">Créer noeud</v-btn>
			</div>
		</div>
		<!-- FOR-EXAMPLE //-->
		<div v-if="currentNode.identifier != ''" style="padding:10px; background-color:orange;">
			<v-text-field v-model="currentNode.title" label="Pseudo" placeholder="Pseudo" required solo
			></v-text-field>
			<TextEditor v-bind:text="currentNode.content" v-bind:textChangedFunction="updateTextNode"></TextEditor>
			<v-btn outlined color="success" @click="saveCurrentNode()">Sauvegarder</v-btn>
		</div>
	</v-container>
</template>
<script src="./ViewBook.ts"></script>
