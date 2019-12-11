<template>
	<div>
		<v-container>
			<v-banner single-line>
				Books
				<template v-slot:actions>
					<v-btn color="orange" dark @click="createBook()">New book</v-btn>
				</template>
			</v-banner>
		</v-container>
		<v-container v-if="session.books.length > 0">
			<ConfirmDialog
				v-bind:open="deleteBookId!=''"
				v-bind:confirmCallback="deleteBook"
				v-bind:dismissCallback="() => { deleteBookId='' }"
				v-bind:description="'Êtes-vous sur de vouloir supprimer ce projet ?'"
			></ConfirmDialog>
			<v-container>
				<v-row>
					<v-col cols="auto"
							v-for="book in session.books"
							v-bind:key="book.identifier"
						>
						<v-card class="mx-auto" max-width="344" outlined>
						<v-img height="250" src="@/assets/book-no-cover.png"></v-img>
						<v-card-title>{{ book.title }}</v-card-title>
						<v-card-text>
							<div>{{ book.description }}</div>
						</v-card-text>
						<v-divider class="mx-4"></v-divider>
						<v-card-text>
							<p  style="display: inline" class="font-weight-black">Genre: </p>
							<p style="display: inline">{{ book.genre }}</p>
							<br />
							<p  style="display: inline" class="font-weight-black">Creation date: </p>
							<p style="display: inline">Fictive date</p>
						</v-card-text>
						<v-card-actions>
							<v-btn color="orange" @click="showBook(book.identifier)">Edit</v-btn>
							<v-btn color="error" @click="() => { deleteBookId = book.identifier }">Delete</v-btn>
						</v-card-actions>
						</v-card>
					</v-col>
				</v-row>
			</v-container>
		</v-container>
		<v-container v-else>
			<v-container>
			<p class="text-center">
				You don't started a book yet. You can give a life on your imagination to write a Heroes book !
			</p>
			</v-container>
			<v-container class="text-center">
				<v-btn large color="primary" @click="createBook()">
					Start my first book
				</v-btn>
			</v-container>
		</v-container>
	</div>
</template>
<script src="./ResumeHome.ts"></script>
