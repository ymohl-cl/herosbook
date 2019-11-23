<template>
  <v-container>
    <div v-if="!createNewBookDisplay">
      <confirmDialog
        v-bind:open="wantDeleteBookId!=''"
        v-bind:confirmCallback="deleteBookAction"
        v-bind:dismissCallback="() => { wantDeleteBookId='' }"
        v-bind:description="'Êtes-vous sur de vouloir supprimer ce projet ?'">

      </confirmDialog>
      <div v-if="books.length > 0">
        <h1>Books list</h1>
        <v-card class="mx-auto" max-width="344" outlined v-for="book in books"
         v-bind:key="book.identifier">
          <v-list-item three-line>
            <v-list-item-content>
              <!-- <div class="overline mb-4">TEST BOOK ?</div> //-->
              <v-list-item-title class="headline mb-1">{{ book.title }}</v-list-item-title>
              <v-list-item-subtitle>{{ book.description }}</v-list-item-subtitle>
            </v-list-item-content>
          </v-list-item>
          <v-card-actions>
            <v-btn outlined color="primary" @click="goBook(book.identifier)">Edit</v-btn>
            <v-btn outlined color="error" @click="openDeleteDialog(book.identifier)">Delete</v-btn>
          </v-card-actions>
        </v-card>
      </div>
      <div class="container pa-3" :elevation="1" v-if="books.length == 0">
        <div class="flex">
          <h3 class="display-2">Bienvenue dans HereosBook</h3>
          <span class="subheading mb-2">
            Vous allez pouvoir donner vie à votre imagination en
            écrivant des livres dont vous êtes le héros !
          </span>
          <v-divider></v-divider>
          <div class="title mb-3">Vous êtes prêt ?</div>
          <div class="my-2">
            <v-btn large color="primary" @click="createNewBookDisplay = true">
              Commencer l'écriture
            </v-btn>
          </div>
        </div>
      </div>
    </div>
    <div v-if="createNewBookDisplay">
      <h1>Votre nouvelle histoire</h1>
      <v-form ref="form" v-model="validForm">
        <v-text-field  v-model="formCreateBook.title" :rules="titleRules"
           label="Titre" placeholder="Titre" required solo></v-text-field>
        <v-textarea  v-model="formCreateBook.description"
          label="Description" placeholder="Description" solo></v-textarea>
        <v-select v-model="formCreateBook.genre"
        :rules="genreRules"
        :items="genres"
        item-text="name"
        item-value="key"
        label="Genre" solo></v-select>
        <div class="text-right">
          <v-btn :disabled="!isValid()" color="success" @click="createBook()">
            Créer le projet
          </v-btn>
        </div>
      </v-form>
    </div>
  </v-container>
</template>

<script src="./home.ts"></script>
