<template>
	<v-container>
		<v-container>
			Use your pseudo and you password to connect on your account
		</v-container>
		<v-container>
			<v-form ref="form" v-model="validForm" lazy-validation>
				<v-text-field  v-model="pseudo" :rules="pseudoRules" :counter="42"
					label="Pseudo" placeholder="Pseudo" required solo
				></v-text-field>
				<v-text-field  v-model="password" :rules="passwordRules" label="Password"
					placeholder="Password"  required  solo
					:append-icon="showPass ? 'mdi-eye' : 'mdi-eye-off'"
					:type="showPass ? 'text' : 'password'"
					@click:append="showPass = !showPass"
				></v-text-field>
				<v-btn :disabled="!validForm" color="success" class="mr-4"
					@click="connect()"
				>Connect</v-btn>
			</v-form>
		</v-container>
		<v-container>
			You don't have an account yet ?
			<a @click="register()">become a Heroes and create my account
			</a>
		</v-container>
	</v-container>
</template>
<script src="./ViewLogin.ts"></script>
