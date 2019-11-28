<template>
  <v-container>
    <v-dialog v-model="open" max-width="290">
      <v-card>
        <v-card-title class="headline">{{ title }}</v-card-title>
        <v-card-text>{{ description }}</v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn  color="green darken-1" text @click="dismiss()">
            Cancel
          </v-btn>
          <v-btn color="green darken-1" text @click="confirm()">
            OK
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>
<script src="./confirmDialog.ts"></script>
