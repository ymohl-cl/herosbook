<template>
	<v-container>
		<v-container>
			Seriously ? You are you a Heroes !? Tell me how your book's name
		</v-container>
		<v-container>
			<v-form ref="form" v-model="validForm">
				<v-text-field v-model="title" :rules="titleRules"
					label="Title" placeholder="Title" required solo></v-text-field>
				<v-text-field v-model="description" :rules="descriptionRules" label="Description"
					placeholder="Description" required solo></v-text-field>
				<v-select v-model="genre" :items="genres" menu-props="auto" label="Genre" single-line
				></v-select>
				<div class="text-right">
					<v-btn class="ma-2 orange--text" color="white" dark @click="cancel()">Cancel</v-btn>
					<v-btn class="ma-2 orange--text" color="white" dark :disabled="!validForm" @click="create()">
						Create
					</v-btn>
				</div>
			</v-form>
		</v-container>
	</v-container>
</template>
<script src="./CreateBook.ts"></script>
