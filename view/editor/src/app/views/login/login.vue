<template>
  <v-container>
      <h1 class="text-center">Login</h1>
      <p class="text-center">Use your pseudo and password to connect on your HereosBook account<a @click="goRegister()">, no have account ? </a></p>
      <v-form ref="form" v-model="validForm">
        <v-text-field  v-model="pseudo" :rules="pseudoRules" label="Pseudo" placeholder="Pseudo" required solo></v-text-field>
        <v-text-field  v-model="password" :rules="passwordRules" label="Password" placeholder="Password"  required  solo
          :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
          :type="showPassword ? 'text' : 'password'"
          @click:append="showPassword = !showPassword"
        ></v-text-field>
        <div class="text-right">
          <v-btn :disabled="!isValid()" color="success" @click="connection()">
            Connection
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

<script src="./login.ts"></script>
