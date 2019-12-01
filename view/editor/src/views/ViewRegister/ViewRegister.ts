import Vue from "vue"
import Component from "vue-class-component"
import navigation from "@/services/ServiceNavigation"
import * as route from "@/router"
import session, { Session } from "@/services/ServiceSession"

@Component
export default class ViewRegister extends Vue {
	session:Session = session
	validForm:boolean = false
	password:string = ""
	showPass:boolean = false
	pseudoRules:any = [(v:string) => !!v || "Pseudo is required"]
	passwordRules:any = [(v:string) => !!v || "Password is required"]
	lastNameRules:any = [(v:string) => !!v || "Last name is required"]
	firstNameRules:any = [(v:string) => !!v || "First name is required"]
	ageRules:any = [(v:number) => v > 0 || "Age not null or negative"]
	genreRules:any = [(v:string) => !!v || "Genre is required"]
	emailRules:any = [
		(v:string) => !!v || "Email is required",
		(v:string) => /.+@.+\..+/.test(v) || "E-mail must be valid",
	]

	genres:string[] = [
		"Male",
		"Female",
		"Anothers",
	]

	login() {
		navigation.replaceView(route.loginPagePath)
	}

	register() {
		this.session.user.record(this.password, () => {})
	}
}
