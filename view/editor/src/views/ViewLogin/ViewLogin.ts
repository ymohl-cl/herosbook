import Vue from "vue"
import Component from "vue-class-component"
import navigation from "@/services/ServiceNavigation"
import session, { Session } from "@/services/ServiceSession"
import * as route from "@/router"
import * as alert from "@/components/AlertMessage/AlertMessage.vue"

@Component
export default class ViewLogin extends Vue {
	session:Session = session
	validForm:boolean = false
	pseudo:string = ""
	password:string = ""
	showPass:boolean = false
	pseudoRules:any = [(v:string) => !!v || "Pseudo is required"]
	passwordRules:any = [(v:string) => !!v || "Password is required"]

	connect() {
		session.user.login(this.pseudo, this.password, () => {
			navigation.replaceView(route.resumePagePath)
		})
		// TODO: implement catch error to add the component alert message
	}

	register() {
		navigation.replaceView(route.registerPagePath)
		// TODO: refact section to add alert
		// this.session.alert.addAlert("Error occurred :D", alert.InfoMessage)
	}
}
