<template>
  <v-container>
      <h1 class="text-center">Register</h1>
      <p class="text-center">
        Create new HereosBook account ! <a @click="goLogin()">I already have an account !</a>
      </p>
      <v-form ref="form" v-model="validForm">
        <v-text-field v-model="pseudo" :rules="pseudoRules"
          label="Pseudo" placeholder="Pseudo" required solo></v-text-field>
        <v-text-field v-model="email" :rules="emailRules" label="Email"
          placeholder="Email" required solo></v-text-field>
        <v-text-field v-model="password" :rules="passwordRules" label="Password"
         placeholder="Password"  required  solo
          :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
          :type="showPassword ? 'text' : 'password'"
          @click:append="showPassword = !showPassword"
        ></v-text-field>
        <v-text-field v-model="lastName" :rules="lastNameRules"
           label="LastName" placeholder="LastName" required solo></v-text-field>
        <v-text-field v-model="firstName" :rules="firstNameRules" label="FirstName"
           placeholder="FirstName" required solo></v-text-field>
        <v-text-field v-model.number="age" type="number" :rules="ageRules" label="Age"
           placeholder="Age" required solo></v-text-field>
        <v-select v-model="genre" :items="genres" :rules="genreRules" label="Genre" required solo>
        </v-select>
        <div class="text-right">
          <v-btn :disabled="!isValid()" color="success" @click="register()">
            Register !
          </v-btn>
        </div>
      </v-form>
      <v-alert class="mt-4" prominent v-if="displayAlert" type="error">
        <v-row align="center">
          <v-col class="grow">Error connection to server : <em>{{ getURL() }}</em> </v-col>
          <v-col class="shrink">
            <v-icon @click="justPing()">mdi-reload</v-icon>
          </v-col>
        </v-row>
      </v-alert>
  </v-container>
</template>

<script src="./register.ts"></script>
