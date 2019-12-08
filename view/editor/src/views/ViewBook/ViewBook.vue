<template>
	<v-container>
		<div>
			<h1>{{ book.title }}</h1>
			<p>{{ book.description }}</p>
		</div>
		<v-row>
			<v-col class="col-4">
				<v-list v-if="nodes.length > 0">
					<v-subheader>Node list</v-subheader>
					<v-list-item-group color="primary">
						<v-list-item v-for="(node, i) in nodes" :key="i" @click="setCurrentNode(node)">
							<v-list-item-content>
								<v-list-item-title>{{ node.title }} (TODO : cliquez pour éditer)</v-list-item-title>
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
			</v-col>
			<v-col class="col-8" v-if="currentNode.identifier != ''">
				<v-container v-if="!editNodeMode">
					<h1>
						{{ currentNode.title }}
						<v-btn outlined color="success" @click="setEditNodeMode(true)">Edit</v-btn>
					</h1>
					<div>{{ currentNode.content }}</div>
				</v-container >
				<!-- FOR-EXAMPLE //-->
				<v-container v-else style="padding:10px; background-color:orange;">
					<v-text-field v-model="currentNode.title" label="Pseudo" placeholder="Pseudo" required solo
					></v-text-field>
					<TextEditor v-bind:text="currentNode.content"
						v-bind:textChangedFunction="updateTextNode"></TextEditor>
					<v-btn outlined color="success" @click="saveCurrentNode()">Sauvegarder</v-btn>
					<v-btn outlined color="danger" @click="setEditNodeMode(false)">Annuler</v-btn>
				</v-container>
			</v-col>
		</v-row>
	</v-container>
</template>
<script src="./ViewBook.ts"></script>
