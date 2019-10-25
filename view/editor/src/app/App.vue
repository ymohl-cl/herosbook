<template>
  <v-app>
    <v-app-bar v-if="isDisplayAppBar()" app>
      <v-toolbar-title class="headline text-uppercase">
        <span class="font-weight-light">Hereos Book - {{ getPseudo() }}</span>
      </v-toolbar-title>
      <v-spacer></v-spacer>
    </v-app-bar>
    <v-content>
      <router-view />
    </v-content>
  </v-app>
</template>
<script src="./App.ts"></script>
