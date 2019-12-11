<template>
	<v-container>
		<ConfirmDialog
				v-bind:open="displayConfirm"
				v-bind:confirmCallback="callbackSuccessConfirm"
				v-bind:dismissCallback="callbackCancelConfirm"
				v-bind:description="textConfirm"
		></ConfirmDialog>
		<v-container v-if="!editBookMode">
			<h1>
				{{ book.title }}
				<v-btn outlined color="success" @click="setEditBookMode(true)">Edit</v-btn>
			</h1>
			<p>{{ book.description }}</p>
		</v-container>
		<!-- FOR-EXAMPLE //-->
		<v-container v-else class="pa-2 orange">
			<v-text-field v-model="book.title" label="Title" placeholder="Title" required solo
			></v-text-field>
			<v-textarea v-model="book.description" solo placeholder="Description"></v-textarea>
			<v-btn outlined color="success" @click="saveCurrentBook()">Sauvegarder</v-btn>
			<v-btn outlined class="ml-3" color="primary" @click="setEditBookMode(false)">Annuler</v-btn>
		</v-container>
		<v-row>
			<v-col>
				<v-list>
					<v-subheader>Category list</v-subheader>
					<v-list-item-group color="primary">
						<v-list-item v-for="(category, i) in book.categories" :key="i">
							<v-list-item-content>
								<v-list-item-title>
									{{ category.title }}
									<v-btn outlined color="error"
										@click="openDeleteCategory(category.identifier)">Delete</v-btn>
								</v-list-item-title>
							</v-list-item-content>
						</v-list-item>
						<v-list-item>
							<v-list-item-content>
								<v-list-item-title>
									<v-row>
										<v-col>
											<v-text-field v-model="titleNewCategory" label="Title" placeholder="Title"
											required solo></v-text-field>
										</v-col>
										<v-col>
											<v-btn outlined color="success" :disabled="titleNewCategory === ''"
												@click="createCategory()">Create category</v-btn>
										</v-col>
									</v-row>
								</v-list-item-title>
							</v-list-item-content>
						</v-list-item>
					</v-list-item-group>
				</v-list>
			</v-col>
		</v-row>
		<v-row>
			<v-col>
				<h4>Affiche la catégorie :</h4>
				<v-select v-model="currentCategoryIdentifier" :items="getCategoriesdentifier()"
					solo label="Catégorie" @input="generateDisplayNode()"></v-select>
			</v-col>
		</v-row>
		<v-row>
			<v-col class="col-4">
				<v-list v-if="nodes.length > 0">
					<v-subheader>Node list</v-subheader>
					<v-list-item-group color="primary">
						<v-list-item v-for="(node, i) in nodesDisplayed" :key="i" @click="setCurrentNode(node)">
							<v-list-item-content>
								<v-list-item-title>{{ node.title }}</v-list-item-title>
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
			<v-col class="col-8 pa-2" v-if="currentNode.identifier != ''">
				<v-container class="white" v-if="!editNodeMode">
					<h1>
						{{ currentNode.title }}
						<v-btn outlined color="success" @click="setEditNodeMode(true)">Edit</v-btn>
						<v-btn outlined class="ml-3" color="error"
							@click="openDeleteCurrentNode(currentNode.identifier)">Delete</v-btn>
					</h1>
					<div v-html="currentNode.content"></div>
				</v-container >
				<!-- FOR-EXAMPLE //-->
				<v-container v-else class="pa-2 orange">
					<v-text-field v-model="currentNode.title" label="Title" placeholder="Title" required solo
					></v-text-field>
					<TextEditor v-bind:text="currentNode.content"
						v-bind:textChangedFunction="updateTextNode"></TextEditor>
					Catégories :
					<span v-if="currentNode.categories.length == 0">Pas de catégorie définie<br ></span>
					<div v-else>
						<div v-for="(catIdentifier, i) in currentNode.categories" :key="i">
							{{ getCategoryByIdentifier(catIdentifier).title }}
								<v-btn color="error" @click="currentNode.removeCategoryByIndex(i)">X</v-btn>
						</div>
					</div>
					<v-row>
						<v-col>
							<v-select v-model="categoryToAdd" :items="book.categories" item-text="title"
								item-value="identifier" solo label="Catégorie"></v-select>
						</v-col>
						<v-col>
							<v-btn outlined color="success" :disabled="categoryToAdd === ''"
								@click="currentNode.addCategory(categoryToAdd)">Ajouter categorie</v-btn>
						</v-col>
					</v-row>
					<v-btn outlined color="success" @click="saveCurrentNode()">Sauvegarder</v-btn>
					<v-btn outlined class="ml-3" color="primary"
						@click="setEditNodeMode(false)">Annuler</v-btn>
				</v-container>
			</v-col>
		</v-row>
	</v-container>
</template>
<script src="./ViewBook.ts"></script>
