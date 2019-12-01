<template>
	<v-container>
		<v-container>
			Create new HereosBook account ! <a @click="login()">I already have an account !</a>
		</v-container>
		<v-container>
			<v-form ref="form" v-model="validForm">
				<v-text-field v-model="session.user.pseudo" :rules="pseudoRules"
					label="Pseudo" placeholder="Pseudo" required solo></v-text-field>
				<v-text-field v-model="session.user.email" :rules="emailRules" label="Email"
					placeholder="Email" required solo></v-text-field>
				<v-text-field v-model="session.user.password" :rules="passwordRules" label="Password"
					placeholder="Password"  required  solo
					:append-icon="showPass ? 'mdi-eye' : 'mdi-eye-off'"
					:type="showPass ? 'text' : 'password'"
					@click:append="showPass = !showPass"
				></v-text-field>
				<v-text-field v-model="session.user.lastName" :rules="lastNameRules"
					label="LastName" placeholder="LastName" required solo></v-text-field>
				<v-text-field v-model="session.user.firstName" :rules="firstNameRules" label="FirstName"
					placeholder="FirstName" required solo></v-text-field>
				<v-text-field v-model.number="session.user.age" type="number" :rules="ageRules" label="Age"
					placeholder="Age" required solo></v-text-field>
				<v-select v-model="session.user.genre" :items="genres" :rules="genreRules" label="Genre" required solo>
				</v-select>
				<div class="text-right">
					<v-btn :disabled="!validForm" color="success" @click="register()">
						Register !
					</v-btn>
				</div>
			</v-form>
		</v-container>
	</v-container>
</template>
<script src="./ViewRegister.ts"></script>
